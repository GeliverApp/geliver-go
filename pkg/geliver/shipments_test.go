package geliver

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func jsonResp(status int, v any) *http.Response {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(v)
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(&buf),
	}
}

func TestListShipments(t *testing.T) {
	c := NewClient("test")
	c.BaseURL = "https://example.test"
	c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if ua := r.Header.Get("User-Agent"); ua != DefaultUserAgent {
			t.Fatalf("expected User-Agent=%q, got %q", DefaultUserAgent, ua)
		}
		if r.Method == "GET" && r.URL.Path == "/shipments" {
			return jsonResp(200, map[string]any{
				"result":     true,
				"limit":      2,
				"page":       1,
				"totalRows":  2,
				"totalPages": 1,
				"data":       []map[string]any{{"id": "s1"}, {"id": "s2"}},
			}), nil
		}
		return jsonResp(404, map[string]any{}), nil
	})}
	out, err := c.ListShipments(context.Background(), &ListParams{Limit: ptrInt(2)})
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Data) != 2 {
		t.Fatalf("expected 2, got %d", len(out.Data))
	}
}

func TestCreateIncomingShipmentKeepsAddressDirection(t *testing.T) {
	c := NewClient("test")
	c.BaseURL = "https://example.test"
	c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if r.Method != "POST" || r.URL.Path != "/shipments" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["isReturn"] != true || body["recipientAddressID"] != "merchant-1" {
			t.Fatalf("unexpected body: %v", body)
		}
		sender, ok := body["senderAddress"].(map[string]any)
		if !ok || sender["name"] != "Customer" {
			t.Fatalf("unexpected sender address: %v", body["senderAddress"])
		}
		if _, exists := body["senderAddressID"]; exists {
			t.Fatalf("unexpected senderAddressID: %v", body["senderAddressID"])
		}
		return jsonResp(200, map[string]any{"result": true, "data": map[string]any{"id": "incoming-1"}}), nil
	})}
	isReturn := true
	shipment, err := c.CreateShipmentWithRecipientID(context.Background(), CreateShipmentWithRecipientID{
		CreateShipmentRequestBase: CreateShipmentRequestBase{
			IsReturn: &isReturn, SenderAddress: &Address{Name: "Customer", Phone: "+905051234567", Address1: "A"},
		},
		RecipientAddressID: "merchant-1",
	})
	if err != nil || shipment.ID != "incoming-1" {
		t.Fatalf("unexpected shipment: %v, %v", shipment, err)
	}
}

func TestCreateIncomingShipmentAddressCombinations(t *testing.T) {
	isReturn := true
	sender := Address{Name: "Customer", Phone: "+905051234567", Address1: "A"}
	recipient := Address{Name: "Merchant", Phone: "+905051234568", Address1: "B"}
	cases := []struct {
		name         string
		body         any
		senderKey    string
		recipientKey string
	}{
		{"id-id", CreateShipmentWithRecipientID{CreateShipmentRequestBase: CreateShipmentRequestBase{IsReturn: &isReturn, SenderAddressID: "customer-1"}, RecipientAddressID: "merchant-1"}, "senderAddressID", "recipientAddressID"},
		{"id-inline", CreateShipmentWithRecipientAddress{CreateShipmentRequestBase: CreateShipmentRequestBase{IsReturn: &isReturn, SenderAddressID: "customer-1"}, RecipientAddress: recipient}, "senderAddressID", "recipientAddress"},
		{"inline-id", CreateShipmentWithRecipientID{CreateShipmentRequestBase: CreateShipmentRequestBase{IsReturn: &isReturn, SenderAddress: &sender}, RecipientAddressID: "merchant-1"}, "senderAddress", "recipientAddressID"},
		{"inline-inline", CreateShipmentWithRecipientAddress{CreateShipmentRequestBase: CreateShipmentRequestBase{IsReturn: &isReturn, SenderAddress: &sender}, RecipientAddress: recipient}, "senderAddress", "recipientAddress"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient("test")
			c.BaseURL = "https://example.test"
			c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
				var body map[string]any
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Fatal(err)
				}
				if r.Method != "POST" || r.URL.Path != "/shipments" || body["isReturn"] != true {
					t.Fatalf("unexpected request: %s %s %v", r.Method, r.URL.Path, body)
				}
				if _, ok := body[tc.senderKey]; !ok {
					t.Fatalf("missing sender selection: %v", body)
				}
				if _, ok := body[tc.recipientKey]; !ok {
					t.Fatalf("missing recipient selection: %v", body)
				}
				for _, key := range []string{"senderAddressID", "senderAddress", "recipientAddressID", "recipientAddress"} {
					if key != tc.senderKey && key != tc.recipientKey {
						if _, ok := body[key]; ok {
							t.Fatalf("unexpected address field %s: %v", key, body)
						}
					}
				}
				return jsonResp(200, map[string]any{"result": true, "data": map[string]any{"id": "incoming-1"}}), nil
			})}
			shipment, err := c.CreateShipmentTyped(context.Background(), tc.body)
			if err != nil || shipment.ID != "incoming-1" {
				t.Fatalf("unexpected shipment: %v, %v", shipment, err)
			}
		})
	}
}

func TestNewShipmentFlowsRequireOneAddressPerSide(t *testing.T) {
	c := NewClient("test")
	invalid := []map[string]any{
		{"recipientAddressID": "merchant-1"},
		{"senderAddressID": "customer-1", "senderAddress": map[string]any{"address1": "A"}, "recipientAddressID": "merchant-1"},
		{"senderAddressID": "customer-1"},
		{"senderAddressID": "customer-1", "recipientAddressID": "merchant-1", "recipientAddress": map[string]any{"address1": "B"}},
	}
	for _, body := range invalid {
		if _, err := c.CreateShipment(context.Background(), body); err == nil {
			t.Fatalf("CreateShipment accepted invalid addresses: %v", body)
		}
		if _, err := c.CreateTransaction(context.Background(), map[string]any{"shipment": body}); err == nil {
			t.Fatalf("CreateTransaction accepted invalid addresses: %v", body)
		}
	}
}

func TestAcceptOffer(t *testing.T) {
	c := NewClient("test")
	c.BaseURL = "https://example.test"
	c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if ua := r.Header.Get("User-Agent"); ua != DefaultUserAgent {
			t.Fatalf("expected User-Agent=%q, got %q", DefaultUserAgent, ua)
		}
		if r.Method == "POST" && r.URL.Path == "/transactions" {
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body["offerID"] != "offer-123" {
				t.Fatalf("unexpected offerID: %v", body["offerID"])
			}
			return jsonResp(200, map[string]any{
				"result": true,
				"data":   map[string]any{"id": "tx1", "offerID": "offer-123"},
			}), nil
		}
		return jsonResp(404, map[string]any{}), nil
	})}
	tx, err := c.AcceptOffer(context.Background(), "offer-123")
	if err != nil {
		t.Fatal(err)
	}
	if tx.ID != "tx1" {
		t.Fatalf("unexpected id: %s", tx.ID)
	}
}

func TestCreateReturnShipment_UsesPostAndDefaults(t *testing.T) {
	c := NewClient("test")
	c.BaseURL = "https://example.test"
	c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if ua := r.Header.Get("User-Agent"); ua != DefaultUserAgent {
			t.Fatalf("expected User-Agent=%q, got %q", DefaultUserAgent, ua)
		}
		if r.Method != "POST" {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/shipments/shp-1" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}

		if body["isReturn"] != true {
			t.Fatalf("expected isReturn=true, got %v", body["isReturn"])
		}
		if body["willAccept"] != false {
			t.Fatalf("expected willAccept=false, got %v", body["willAccept"])
		}

		// json.Decoder decodes numbers as float64 by default
		if body["count"] != float64(1) {
			t.Fatalf("expected count=1, got %v", body["count"])
		}

		return jsonResp(200, map[string]any{
			"result": true,
			"data":   map[string]any{"id": "ret-1"},
		}), nil
	})}

	ret, err := c.CreateReturnShipment(context.Background(), "shp-1", ReturnShipmentRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if ret.ID != "ret-1" {
		t.Fatalf("unexpected id: %s", ret.ID)
	}
}

func TestCreateReturnShipment_WithWillAccept_ReturnsError(t *testing.T) {
	c := NewClient("test")
	if _, err := c.CreateReturnShipment(context.Background(), "shp-1", ReturnShipmentRequest{WillAccept: true}); err == nil {
		t.Fatal("expected error when WillAccept=true")
	}
}

func ptrInt(v int) *int { return &v }
