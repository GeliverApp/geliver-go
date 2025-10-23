package geliver

import (
    "context"
    "net/url"
)

// ListCities returns cities for a given country code.
func (c *Client) ListCities(ctx context.Context, countryCode string) ([]City, error) {
    q := url.Values{"countryCode": []string{countryCode}}
    var out []City
    if err := c.do(ctx, "GET", "/cities", q, nil, &out); err != nil { return nil, err }
    return out, nil
}

// ListDistricts returns districts for a given country and city code.
func (c *Client) ListDistricts(ctx context.Context, countryCode, cityCode string) ([]District, error) {
    q := url.Values{"countryCode": []string{countryCode}, "cityCode": []string{cityCode}}
    var out []District
    if err := c.do(ctx, "GET", "/districts", q, nil, &out); err != nil { return nil, err }
    return out, nil
}

// ListCitiesTyped returns typed slice of City.
func (c *Client) ListCitiesTyped(ctx context.Context, countryCode string) ([]City, error) {
    var out []City
    if err := c.do(ctx, "GET", "/cities", url.Values{"countryCode": []string{countryCode}}, nil, &out); err != nil { return nil, err }
    return out, nil
}

// ListDistrictsTyped returns typed slice of District.
func (c *Client) ListDistrictsTyped(ctx context.Context, countryCode, cityCode string) ([]District, error) {
    var out []District
    if err := c.do(ctx, "GET", "/districts", url.Values{"countryCode": []string{countryCode}, "cityCode": []string{cityCode}}, nil, &out); err != nil { return nil, err }
    return out, nil
}
