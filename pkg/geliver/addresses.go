package geliver

import (
    "context"
    "errors"
    "net/url"
)

// CreateAddress creates a new address and returns the typed Address.
func (c *Client) CreateAddress(ctx context.Context, req CreateAddressRequest) (*Address, error) {
    var out Address
    if err := c.do(ctx, "POST", "/addresses", nil, req, &out); err != nil { return nil, err }
    return &out, nil
}

// AddressesListResponse represents the paginated list of addresses.
type AddressesListResponse struct {
    Result            bool      `json:"result,omitempty"`
    AdditionalMessage string    `json:"additionalMessage,omitempty"`
    Limit             int       `json:"limit,omitempty,string"`
    Page              int       `json:"page,omitempty,string"`
    TotalRows         int       `json:"totalRows,omitempty,string"`
    TotalPages        int       `json:"totalPages,omitempty,string"`
    Data              []Address `json:"data,omitempty"`
}

// ListAddresses returns a typed paginated list of addresses.
func (c *Client) ListAddresses(ctx context.Context, isRecipient *bool, limit, page *int) (*AddressesListResponse, error) {
    q := url.Values{}
    if isRecipient != nil { q.Set("isRecipientAddress", btoa(*isRecipient)) }
    if limit != nil { q.Set("limit", itoa(*limit)) }
    if page != nil { q.Set("page", itoa(*page)) }
    var out AddressesListResponse
    if err := c.do(ctx, "GET", "/addresses", q, nil, &out); err != nil { return nil, err }
    return &out, nil
}

// CreateSenderAddress helper sets isRecipientAddress=false.
func (c *Client) CreateSenderAddress(ctx context.Context, req CreateAddressRequest) (*Address, error) {
    f := false
    req.IsRecipientAddress = &f
    if req.Zip == nil || *req.Zip == "" {
        return nil, errors.New("zip is required for sender addresses")
    }
    return c.CreateAddress(ctx, req)
}

// CreateRecipientAddress helper sets isRecipientAddress=true.
func (c *Client) CreateRecipientAddress(ctx context.Context, req CreateAddressRequest) (*Address, error) {
    t := true
    req.IsRecipientAddress = &t
    return c.CreateAddress(ctx, req)
}

// GetAddressTyped returns a typed Address by ID.
func (c *Client) GetAddressTyped(ctx context.Context, addressID string) (*Address, error) {
    var out Address
    if err := c.do(ctx, "GET", "/addresses/"+addressID, nil, nil, &out); err != nil { return nil, err }
    return &out, nil
}
