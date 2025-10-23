package geliver

import (
    "context"
    "net/url"
)

// ListProviderAccounts lists provider accounts (typed array from envelope).
func (c *Client) ListProviderAccounts(ctx context.Context) ([]ProviderAccount, error) {
    var out []ProviderAccount
    if err := c.do(ctx, "GET", "/provideraccounts", nil, nil, &out); err != nil { return nil, err }
    return out, nil
}

// NOTE: Typed variant removed due to schema changes; use ListProviderAccounts instead.

// CreateProviderAccount creates a provider account connection.
func (c *Client) CreateProviderAccount(ctx context.Context, req ProviderAccountRequest) (*ProviderAccount, error) {
    var out ProviderAccount
    if err := c.do(ctx, "POST", "/provideraccounts", nil, req, &out); err != nil { return nil, err }
    return &out, nil
}

// DeleteProviderAccount deletes a provider account connection. Returns deleted resource if available.
func (c *Client) DeleteProviderAccount(ctx context.Context, providerAccountID string, isDeleteAccountConnection *bool) (*ProviderAccount, error) {
    q := url.Values{}
    if isDeleteAccountConnection != nil { q.Set("isDeleteAccountConnection", btoa(*isDeleteAccountConnection)) }
    var out ProviderAccount
    if err := c.do(ctx, "DELETE", "/provideraccounts/"+providerAccountID, q, nil, &out); err != nil { return nil, err }
    return &out, nil
}
