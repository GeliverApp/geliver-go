package geliver

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateTransactionWrapsShipment(t *testing.T) {
	c := NewClient("test")
	c.BaseURL = "https://example.test"
	c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if r.Method == "POST" && r.URL.Path == "/transactions" {
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if _, ok := body["test"]; ok {
				t.Fatalf("test must be nested under shipment")
			}
			if body["providerServiceCode"] != "SURAT_STANDART" {
				t.Fatalf("expected providerServiceCode at root")
			}
			if body["providerAccountID"] != "acc-1" {
				t.Fatalf("expected providerAccountID at root")
			}
			shipment, ok := body["shipment"].(map[string]any)
			if !ok {
				t.Fatalf("expected shipment object")
			}
			if shipment["test"] != true {
				t.Fatalf("expected shipment.test=true")
			}
			if _, ok := shipment["providerServiceCode"]; ok {
				t.Fatalf("providerServiceCode must be at root")
			}
			if _, ok := shipment["providerAccountID"]; ok {
				t.Fatalf("providerAccountID must be at root")
			}
			if _, ok := shipment["length"].(string); !ok {
				t.Fatalf("expected shipment.length to be string")
			}
			if _, ok := shipment["weight"].(string); !ok {
				t.Fatalf("expected shipment.weight to be string")
			}
			order, ok := shipment["order"].(map[string]any)
			if !ok || order["sourceCode"] != "SDK" {
				t.Fatalf("expected order.sourceCode=SDK")
			}

			return jsonResp(200, map[string]any{
				"result": true,
				"data":   map[string]any{"id": "tx1", "offerID": "offer-123"},
			}), nil
		}
		return jsonResp(404, map[string]any{}), nil
	})}
	ctx := context.Background()
	tx, err := c.CreateTransaction(ctx, map[string]any{
		"senderAddressID": "sender-1",
		"recipientAddress": map[string]any{
			"name":  "R",
			"phone": "+905000000000",
		},
		"length":              10.5,
		"weight":              1.25,
		"distanceUnit":        "cm",
		"massUnit":            "kg",
		"test":                true,
		"providerServiceCode": "SURAT_STANDART",
		"providerAccountID":   "acc-1",
		"order":               map[string]any{"orderNumber": "ORDER-1"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil || tx.ID != "tx1" {
		t.Fatalf("unexpected tx: %+v", tx)
	}
}

func TestCreateReturnTransaction_UsesShipmentEndpointAndForcesWillAccept(t *testing.T) {
	c := NewClient("test")
	c.BaseURL = "https://example.test"
	c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if r.Method != "POST" || r.URL.Path != "/shipments/shp-1" {
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}

		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["isReturn"] != true {
			t.Fatalf("expected isReturn=true, got %v", body["isReturn"])
		}
		if body["willAccept"] != true {
			t.Fatalf("expected willAccept=true, got %v", body["willAccept"])
		}
		if body["count"] != float64(1) {
			t.Fatalf("expected count=1, got %v", body["count"])
		}
		if body["providerServiceCode"] != "SURAT_STANDART" {
			t.Fatalf("unexpected providerServiceCode: %v", body["providerServiceCode"])
		}

		return jsonResp(200, map[string]any{
			"result": true,
			"data":   map[string]any{"id": "tx-1", "offerID": "offer-1", "transactionType": "CREATE_SHIPMENT"},
		}), nil
	})}

	tx, err := c.CreateReturnTransaction(context.Background(), "shp-1", ReturnShipmentRequest{
		ProviderServiceCode: ptr("SURAT_STANDART"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil || tx.ID != "tx-1" {
		t.Fatalf("unexpected tx: %+v", tx)
	}
}

func ptr(v string) *string { return &v }
