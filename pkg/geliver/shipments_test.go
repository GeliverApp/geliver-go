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
        if r.Method == "GET" && r.URL.Path == "/shipments" {
            return jsonResp(200, map[string]any{
                "result": true,
                "limit": 2,
                "page": 1,
                "totalRows": 2,
                "totalPages": 1,
                "data": []map[string]any{{"id": "s1"}, {"id": "s2"}},
            }), nil
        }
        return jsonResp(404, map[string]any{}), nil
    })}
    out, err := c.ListShipments(context.Background(), &ListParams{Limit: ptrInt(2)})
    if err != nil { t.Fatal(err) }
    if len(out.Data) != 2 { t.Fatalf("expected 2, got %d", len(out.Data)) }
}

func TestAcceptOffer(t *testing.T) {
    c := NewClient("test")
    c.BaseURL = "https://example.test"
    c.HTTP = &http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
        if r.Method == "POST" && r.URL.Path == "/transactions" {
            var body map[string]any
            if err := json.NewDecoder(r.Body).Decode(&body); err != nil { t.Fatal(err) }
            if body["offerID"] != "offer-123" { t.Fatalf("unexpected offerID: %v", body["offerID"]) }
            return jsonResp(200, map[string]any{
                "result": true,
                "data": map[string]any{"id": "tx1", "offerID": "offer-123"},
            }), nil
        }
        return jsonResp(404, map[string]any{}), nil
    })}
    tx, err := c.AcceptOffer(context.Background(), "offer-123")
    if err != nil { t.Fatal(err) }
    if tx.ID != "tx1" { t.Fatalf("unexpected id: %s", tx.ID) }
}

func ptrInt(v int) *int { return &v }
