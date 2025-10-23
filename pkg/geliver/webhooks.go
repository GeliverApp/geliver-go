package geliver

import (
    "context"
)

func (c *Client) CreateWebhook(ctx context.Context, url string, typ *string) (map[string]any, error) {
    body := map[string]any{"url": url}
    if typ != nil { body["type"] = *typ }
    var out map[string]any
    if err := c.do(ctx, "POST", "/webhook", nil, body, &out); err != nil { return nil, err }
    return out, nil
}

func (c *Client) ListWebhooks(ctx context.Context) (map[string]any, error) {
    var out map[string]any
    if err := c.do(ctx, "GET", "/webhook", nil, nil, &out); err != nil { return nil, err }
    return out, nil
}

func (c *Client) DeleteWebhook(ctx context.Context, webhookID string) (map[string]any, error) {
    var out map[string]any
    if err := c.do(ctx, "DELETE", "/webhook/"+webhookID, nil, nil, &out); err != nil { return nil, err }
    return out, nil
}

func (c *Client) TestWebhook(ctx context.Context, typ, url string) (map[string]any, error) {
    body := map[string]any{"type": typ, "url": url}
    var out map[string]any
    if err := c.do(ctx, "PUT", "/webhook", nil, body, &out); err != nil { return nil, err }
    return out, nil
}

