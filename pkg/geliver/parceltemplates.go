package geliver

import (
    "context"
)

// CreateParcelTemplate creates a parcel template.
func (c *Client) CreateParcelTemplate(ctx context.Context, req CreateParcelTemplateRequest) (*ParcelTemplate, error) {
    var out ParcelTemplate
    if err := c.do(ctx, "POST", "/parceltemplates", nil, req, &out); err != nil { return nil, err }
    return &out, nil
}

// ListParcelTemplates lists parcel templates.
func (c *Client) ListParcelTemplates(ctx context.Context) ([]ParcelTemplate, error) {
    var out []ParcelTemplate
    if err := c.do(ctx, "GET", "/parceltemplates", nil, nil, &out); err != nil { return nil, err }
    return out, nil
}

// NOTE: Typed variant removed due to schema changes; use ListParcelTemplates instead.

// DeleteParcelTemplate removes a parcel template.
func (c *Client) DeleteParcelTemplate(ctx context.Context, templateID string) (*ParcelTemplate, error) {
    var out ParcelTemplate
    if err := c.do(ctx, "DELETE", "/parceltemplates/"+templateID, nil, nil, &out); err != nil { return nil, err }
    return &out, nil
}
