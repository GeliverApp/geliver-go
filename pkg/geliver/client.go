package geliver

import (
    "bytes"
    "context"
    "encoding/json"
    "io"
    "net/http"
    "net/url"
    "strings"
    "time"
)

const DefaultBaseURL = "https://api.geliver.io/api/v1"

type Client struct {
    BaseURL    string
    Token      string
    HTTP       *http.Client
    MaxRetries int
}

type envelope[T any] struct {
    Result           *bool  `json:"result,omitempty"`
    Message          string `json:"message,omitempty"`
    AdditionalMessage string `json:"additionalMessage,omitempty"`
    Limit            *int   `json:"limit,omitempty"`
    Page             *int   `json:"page,omitempty"`
    TotalRows        *int   `json:"totalRows,omitempty"`
    TotalPages       *int   `json:"totalPages,omitempty"`
    Data             T      `json:"data"`
}

type APIError struct {
    Status int
    Code   string
    Body   any
}

func (e *APIError) Error() string { return "geliver: api error" }

func NewClient(token string) *Client {
    return &Client{
        BaseURL:    DefaultBaseURL,
        Token:      token,
        HTTP:       &http.Client{Timeout: 30 * time.Second},
        MaxRetries: 2,
    }
}

func (c *Client) do(ctx context.Context, method, path string, q url.Values, body any, out any) error {
    base := strings.TrimRight(c.BaseURL, "/")
    u, _ := url.Parse(base + path)
    if q != nil {
        u.RawQuery = q.Encode()
    }
    var rdr io.Reader
    if body != nil {
        b, _ := json.Marshal(body)
        rdr = bytes.NewReader(b)
    }
    req, _ := http.NewRequestWithContext(ctx, method, u.String(), rdr)
    req.Header.Set("Authorization", "Bearer "+c.Token)
    req.Header.Set("Content-Type", "application/json")

    attempt := 0
    for {
        res, err := c.HTTP.Do(req)
        if err != nil {
            if attempt >= c.MaxRetries { return err }
            attempt++
            backoff(attempt)
            continue
        }
        defer res.Body.Close()
        b, _ := io.ReadAll(res.Body)
        if res.StatusCode >= 400 {
            var parsed map[string]any
            _ = json.Unmarshal(b, &parsed)
            apiErr := &APIError{Status: res.StatusCode}
            if code, _ := parsed["code"].(string); code != "" { apiErr.Code = code }
            apiErr.Body = parsed
            if shouldRetry(res.StatusCode) && attempt < c.MaxRetries {
                attempt++
                backoff(attempt)
                continue
            }
            return apiErr
        }
        if out == nil { return nil }
        // Try envelope
        type anyEnvelope struct{ Data json.RawMessage `json:"data"` }
        var env anyEnvelope
        if err := json.Unmarshal(b, &env); err == nil && env.Data != nil {
            return json.Unmarshal(env.Data, out)
        }
        return json.Unmarshal(b, out)
    }
}

func shouldRetry(status int) bool { return status == 429 || status >= 500 }

func backoff(attempt int) {
    base := time.Duration(200*(1<<(attempt-1))) * time.Millisecond
    if base > 2*time.Second { base = 2 * time.Second }
    time.Sleep(base)
}
