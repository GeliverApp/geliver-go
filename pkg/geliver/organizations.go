package geliver

import "context"

// GetBalance returns the organization balance as a typed structure.
func (c *Client) GetBalance(ctx context.Context, organizationID string) (*OrganizationBalance, error) {
    var out OrganizationBalance
    if err := c.do(ctx, "GET", "/organizations/"+organizationID+"/balance", nil, nil, &out); err != nil { return nil, err }
    return &out, nil
}
