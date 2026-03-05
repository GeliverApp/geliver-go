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

func ptrInt(v int) *int { return &v }
