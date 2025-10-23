package geliver

import (
    "context"
)

// AcceptOffer purchases a label using offerID and returns a typed Transaction.
func (c *Client) AcceptOffer(ctx context.Context, offerID string) (*Transaction, error) {
    body := map[string]any{"offerID": offerID}
    var out Transaction
    if err := c.do(ctx, "POST", "/transactions", nil, body, &out); err != nil { return nil, err }
    return &out, nil
}
