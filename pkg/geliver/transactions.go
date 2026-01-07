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
    payload := map[string]any{}
    for k, v := range body { payload[k] = v }
    // default order.sourceCode
    if ov, ok := payload["order"].(map[string]any); ok {
        if ov["sourceCode"] == nil || ov["sourceCode"] == "" { ov["sourceCode"] = "API"; payload["order"] = ov }
    }
    if ra, ok := payload["recipientAddress"].(map[string]any); ok {
        if ph, okp := ra["phone"].(string); !okp || ph == "" {
            return nil, fmt.Errorf("phone is required for recipientAddress")
        }
    }
    // normalize numeric-to-string for dimension/weight
    for _, k := range []string{"length","width","height","weight"} {
        if v, ok := payload[k]; ok && v != nil { payload[k] = toString(v) }
    }
    var out Transaction
    if err := c.do(ctx, "POST", "/transactions", nil, map[string]any{"shipment": payload}, &out); err != nil { return nil, err }
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
