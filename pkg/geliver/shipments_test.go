package geliver

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestListShipments(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/shipments" {
            w.Header().Set("Content-Type", "application/json")
            _ = json.NewEncoder(w).Encode(map[string]any{
                "result": true,
                "limit": 2,
                "page": 1,
                "totalRows": 2,
                "totalPages": 1,
                "data": []map[string]any{{"id": "s1"}, {"id": "s2"}},
            })
            return
        }
        w.WriteHeader(404)
    }))
    defer srv.Close()

    c := NewClient("test")
    c.BaseURL = srv.URL
    out, err := c.ListShipments(context.Background(), &ListParams{Limit: ptrInt(2)})
    if err != nil { t.Fatal(err) }
    if len(out.Data) != 2 { t.Fatalf("expected 2, got %d", len(out.Data)) }
}

func TestAcceptOffer(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" && r.URL.Path == "/transactions" {
            _ = json.NewEncoder(w).Encode(map[string]any{
                "result": true,
                "data": map[string]any{"id": "tx1", "offerID": "offer-123"},
            })
            return
        }
        w.WriteHeader(404)
    }))
    defer srv.Close()

    c := NewClient("test")
    c.BaseURL = srv.URL
    tx, err := c.AcceptOffer(context.Background(), "offer-123")
    if err != nil { t.Fatal(err) }
    if tx.ID != "tx1" { t.Fatalf("unexpected id: %s", tx.ID) }
}

func ptrInt(v int) *int { return &v }

