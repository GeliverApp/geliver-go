package geliver

import (
    "context"
    "encoding/json"
    "fmt"
)

// AcceptOffer purchases a label using offerID and returns a typed Transaction.
func (c *Client) AcceptOffer(ctx context.Context, offerID string) (*Transaction, error) {
    body := map[string]any{"offerID": offerID}
    var out Transaction
    if err := c.do(ctx, "POST", "/transactions", nil, body, &out); err != nil { return nil, err }
    return &out, nil
}

// CreateTransaction performs one-step label purchase by posting shipment details directly to /transactions.
// The body is similar to CreateShipment; server creates shipment and returns a Transaction.
func (c *Client) CreateTransaction(ctx context.Context, body map[string]any) (*Transaction, error) {
    if body == nil { body = map[string]any{} }
    // default order.sourceCode
    if ov, ok := body["order"].(map[string]any); ok {
        if ov["sourceCode"] == nil || ov["sourceCode"] == "" { ov["sourceCode"] = "API"; body["order"] = ov }
    }
    if ra, ok := body["recipientAddress"].(map[string]any); ok {
        if ph, okp := ra["phone"].(string); !okp || ph == "" {
            return nil, fmt.Errorf("phone is required for recipientAddress")
        }
    }
    // normalize numeric-to-string for dimension/weight
    for _, k := range []string{"length","width","height","weight"} {
        if v, ok := body[k]; ok && v != nil { body[k] = toString(v) }
    }
    var out Transaction
    if err := c.do(ctx, "POST", "/transactions", nil, body, &out); err != nil { return nil, err }
    return &out, nil
}

// CreateTransactionWithRecipientAddress typed helper for one-step purchase.
func (c *Client) CreateTransactionWithRecipientAddress(ctx context.Context, req CreateShipmentWithRecipientAddress) (*Transaction, error) {
    b, _ := json.Marshal(req)
    var m map[string]any
    _ = json.Unmarshal(b, &m)
    return c.CreateTransaction(ctx, m)
}

// CreateTransactionWithRecipientID typed helper for one-step purchase.
func (c *Client) CreateTransactionWithRecipientID(ctx context.Context, req CreateShipmentWithRecipientID) (*Transaction, error) {
    b, _ := json.Marshal(req)
    var m map[string]any
    _ = json.Unmarshal(b, &m)
    return c.CreateTransaction(ctx, m)
}

func toString(v any) string { return fmt.Sprintf("%v", v) }
