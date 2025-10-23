package geliver

import (
    "context"
    "net/url"
)

type PriceListParams struct {
    ParamType    string
    Length       float64
    Width        float64
    Height       float64
    Weight       float64
    DistanceUnit *string
    MassUnit     *string
}

// ListPrices queries price list for given parcel dimensions/weight.
func (c *Client) ListPrices(ctx context.Context, p PriceListParams) (map[string]any, error) {
    q := url.Values{}
    q.Set("paramType", p.ParamType)
    q.Set("length", itoa(int(p.Length)))
    q.Set("width", itoa(int(p.Width)))
    q.Set("height", itoa(int(p.Height)))
    q.Set("weight", itoa(int(p.Weight)))
    if p.DistanceUnit != nil { q.Set("distanceUnit", *p.DistanceUnit) }
    if p.MassUnit != nil { q.Set("massUnit", *p.MassUnit) }
    var out map[string]any
    if err := c.do(ctx, "GET", "/priceList", q, nil, &out); err != nil { return nil, err }
    return out, nil
}
