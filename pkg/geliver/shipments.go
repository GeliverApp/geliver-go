package geliver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ListShipmentsResponse struct {
	Limit      *int       `json:"limit,omitempty"`
	Page       *int       `json:"page,omitempty"`
	TotalRows  *int       `json:"totalRows,omitempty"`
	TotalPages *int       `json:"totalPages,omitempty"`
	Data       []Shipment `json:"data"`
}

type ListParams struct {
	Limit               *int
	Page                *int
	SortBy              *string
	Filter              *string
	StartDate           *string
	EndDate             *string
	StatusFilter        *string
	InvoiceID           *string
	MerchantCode        *string
	OrderNumber         *string
	ProviderServiceCode *string
	StoreIdentifier     *string
	IsReturned          *bool
}

func (c *Client) CreateShipment(ctx context.Context, body map[string]any) (*Shipment, error) {
	if body == nil {
		body = map[string]any{}
	}
	if ord, ok := body["order"].(map[string]any); ok {
		if _, ok2 := ord["sourceCode"]; !ok2 || ord["sourceCode"] == "" {
			ord["sourceCode"] = "API"
			body["order"] = ord
		}
	}
	var out Shipment
	if err := c.do(ctx, "POST", "/shipments", nil, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CreateShipmentTyped(ctx context.Context, body any) (*Shipment, error) {
	// normalize to map and enforce default sourceCode
	b, _ := json.Marshal(body)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	if m == nil {
		m = map[string]any{}
	}
	if ov, ok := m["order"]; ok {
		if ord, ok2 := ov.(map[string]any); ok2 {
			if _, ok3 := ord["sourceCode"]; !ok3 || ord["sourceCode"] == "" {
				ord["sourceCode"] = "API"
				m["order"] = ord
			}
		}
	}
	return c.CreateShipment(ctx, m)
}

// CreateShipmentWithRecipientID creates a shipment using a recipient address ID (typed request).
func (c *Client) CreateShipmentWithRecipientID(ctx context.Context, req CreateShipmentWithRecipientID) (*Shipment, error) {
	return c.CreateShipmentTyped(ctx, req)
}

// CreateShipmentWithRecipientAddress creates a shipment using an inline recipient address (typed request).
func (c *Client) CreateShipmentWithRecipientAddress(ctx context.Context, req CreateShipmentWithRecipientAddress) (*Shipment, error) {
	return c.CreateShipmentTyped(ctx, req)
}

func (c *Client) GetShipment(ctx context.Context, shipmentID string) (*Shipment, error) {
	var out Shipment
	if err := c.do(ctx, "GET", "/shipments/"+url.PathEscape(shipmentID), nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) ListShipments(ctx context.Context, p *ListParams) (*ListShipmentsResponse, error) {
	q := url.Values{}
	if p != nil {
		if p.Limit != nil {
			q.Set("limit", itoa(*p.Limit))
		}
		if p.Page != nil {
			q.Set("page", itoa(*p.Page))
		}
		if p.SortBy != nil {
			q.Set("sortBy", *p.SortBy)
		}
		if p.Filter != nil {
			q.Set("filter", *p.Filter)
		}
		if p.StartDate != nil {
			q.Set("startDate", *p.StartDate)
		}
		if p.EndDate != nil {
			q.Set("endDate", *p.EndDate)
		}
		if p.StatusFilter != nil {
			q.Set("statusFilter", *p.StatusFilter)
		}
		if p.InvoiceID != nil {
			q.Set("invoiceID", *p.InvoiceID)
		}
		if p.MerchantCode != nil {
			q.Set("merchantCode", *p.MerchantCode)
		}
		if p.OrderNumber != nil {
			q.Set("orderNumber", *p.OrderNumber)
		}
		if p.ProviderServiceCode != nil {
			q.Set("providerServiceCode", *p.ProviderServiceCode)
		}
		if p.StoreIdentifier != nil {
			q.Set("storeIdentifier", *p.StoreIdentifier)
		}
		if p.IsReturned != nil {
			q.Set("isReturned", btoa(*p.IsReturned))
		}
	}
	var out ListShipmentsResponse
	if err := c.do(ctx, "GET", "/shipments", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdatePackage(ctx context.Context, shipmentID string, body map[string]any) (*Shipment, error) {
	var out Shipment
	if err := c.do(ctx, "PATCH", "/shipments/"+url.PathEscape(shipmentID), nil, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdatePackageTyped(ctx context.Context, shipmentID string, body UpdatePackageRequest) (*Shipment, error) {
	var out Shipment
	if err := c.do(ctx, "PATCH", "/shipments/"+url.PathEscape(shipmentID), nil, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CancelShipment(ctx context.Context, shipmentID string) (*Shipment, error) {
	var out Shipment
	if err := c.do(ctx, "DELETE", "/shipments/"+url.PathEscape(shipmentID), nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CloneShipment(ctx context.Context, shipmentID string) (*Shipment, error) {
	var out Shipment
	if err := c.do(ctx, "POST", "/shipments/"+url.PathEscape(shipmentID), nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Helpers
func (c *Client) WaitForOffers(ctx context.Context, shipmentID string, interval time.Duration, timeout time.Duration) (map[string]any, error) {
	start := time.Now()
	for {
		s, err := c.GetShipment(ctx, shipmentID)
		if err != nil {
			return nil, err
		}
		// Offers may be a struct type (OfferList) depending on schema
		// Try struct fields first
		if s.Offers.PercentageCompleted == 100 {
			// marshal to map for consistency
			var out map[string]any
			// cheap conversion (clients needing strong typing can use models.go directly)
			out = map[string]any{}
			out["percentageCompleted"] = s.Offers.PercentageCompleted
			out["cheapest"] = s.Offers.Cheapest
			out["fastest"] = s.Offers.Fastest
			out["list"] = s.Offers.List
			return out, nil
		}
		if time.Since(start) > timeout {
			return nil, errors.New("timeout waiting for offers")
		}
		time.Sleep(interval)
	}
}

func (c *Client) WaitForTrackingNumber(ctx context.Context, shipmentID string, interval time.Duration, timeout time.Duration) (*Shipment, error) {
	start := time.Now()
	for {
		s, err := c.GetShipment(ctx, shipmentID)
		if err != nil {
			return nil, err
		}
		if s.TrackingNumber != "" {
			return s, nil
		}
		if time.Since(start) > timeout {
			return nil, errors.New("timeout waiting for tracking number")
		}
		time.Sleep(interval)
	}
}

// CreateReturnShipment creates a return for given shipment ID (PATCH with isReturn=true).
func (c *Client) CreateReturnShipment(ctx context.Context, shipmentID string, body ReturnShipmentRequest) (*Shipment, error) {
	body.IsReturn = true
	var out Shipment
	if err := c.do(ctx, "PATCH", "/shipments/"+url.PathEscape(shipmentID), nil, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Label downloads
func (c *Client) DownloadURL(ctx context.Context, url string) ([]byte, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("download error: %d", res.StatusCode)
	}
	return io.ReadAll(res.Body)
}

func (c *Client) DownloadResponsiveURL(ctx context.Context, url string) (string, error) {
	b, err := c.DownloadURL(ctx, url)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Client) DownloadShipmentLabel(ctx context.Context, shipmentID string) ([]byte, error) {
	s, err := c.GetShipment(ctx, shipmentID)
	if err != nil {
		return nil, err
	}
	if s.LabelURL == "" {
		return nil, errors.New("shipment has no labelURL")
	}
	return c.DownloadURL(ctx, s.LabelURL)
}

func itoa(i int) string { return fmt.Sprintf("%d", i) }
func btoa(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
