package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/GeliverApp/sdk-go/pkg/geliver"
)

func ptrs(s string) *string { return &s }
func ptrb(b bool) *bool     { return &b }

func main() {
	token := os.Getenv("GELIVER_TOKEN")
	if token == "" {
		fmt.Println("GELIVER_TOKEN is required")
		return
	}
	c := geliver.NewClient(token)
	ctx := context.Background()
	senderPhone := "+905051234567"
	sender, _ := c.CreateSenderAddress(ctx, geliver.CreateAddressRequest{
		Name: "ACME Inc.", Email: "ops@acme.test", Phone: &senderPhone,
		Address1: "Street 1", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
		DistrictName: "Esenyurt", DistrictID: 107605, Zip: "34020",
	})

	// Inline alıcı adresi (kayıt oluşturmadan)
	recipientPhone := "+905051234568"
	length, width, height, weight := 10.0, 10.0, 10.0, 1.0
	req := geliver.CreateShipmentWithRecipientAddress{
		CreateShipmentRequestBase: geliver.CreateShipmentRequestBase{
			SourceCode: "API", SenderAddressID: sender.ID,
			Length: &length, Width: &width, Height: &height, DistanceUnit: ptrs("cm"), Weight: &weight, MassUnit: ptrs("kg"), Test: ptrb(true),
		},
		RecipientAddress: geliver.Address{
			Name: "John Doe", Email: "john@example.com", Phone: recipientPhone,
			Address1: "Dest St 2", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
			DistrictName: "Esenyurt", DistrictID: 107605, Zip: "34020",
		},
	}
	s, err := c.CreateShipmentWithRecipientAddress(ctx, req)
	if err != nil || s == nil {
		fmt.Println("create shipment error", err)
		return
	}

	// Teklifler create yanıtında hazır olabilir; önce onu kontrol edin
	offers := s.Offers
	if !(offers.PercentageCompleted >= 99 || offers.Cheapest != nil) {
		for {
			gs, err := c.GetShipment(ctx, s.ID)
			if err != nil || gs == nil {
				fmt.Println("fetch shipment error", err)
				return
			}
			if gs.Offers.PercentageCompleted >= 99 && gs.Offers.Cheapest != nil {
				offers = gs.Offers
				break
			}
			time.Sleep(time.Second)
		}
	}
	// Accept offer
	if offers.Cheapest != nil {
		_, _ = c.AcceptOffer(ctx, offers.Cheapest.ID)
	}
	// fetch latest shipment to print details
	latestAccepted, _ := c.GetShipment(ctx, s.ID)
	if latestAccepted.Barcode != "" {
		fmt.Println("barcode:", latestAccepted.Barcode)
	}
	if latestAccepted.TrackingNumber != "" {
		fmt.Println("tracking number:", latestAccepted.TrackingNumber)
	}
	if latestAccepted.LabelURL != "" {
		fmt.Println("label:", latestAccepted.LabelURL)
	}
	if latestAccepted.TrackingURL != "" {
		fmt.Println("tracking:", latestAccepted.TrackingURL)
	}

	// Test gönderilerinde her GET /shipments isteği kargo durumunu bir adım ilerletir; prod'da webhook önerilir.
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		_, _ = c.GetShipment(ctx, s.ID)
	}
	tracked, _ := c.GetShipment(ctx, s.ID)
	fmt.Println("tracking number (refresh):", tracked.TrackingNumber)
	if tracked.TrackingStatus != nil {
		fmt.Println("final status:", tracked.TrackingStatus.TrackingStatusCode, tracked.TrackingStatus.TrackingSubStatusCode)
	}
	latest, _ := c.GetShipment(ctx, s.ID)
	if latest.TrackingStatus != nil {
		fmt.Println("status:", latest.TrackingStatus.TrackingStatusCode, latest.TrackingStatus.TrackingSubStatusCode)
	}
	// download labels
	b, _ := c.DownloadShipmentLabel(ctx, s.ID)
	_ = os.WriteFile("label.pdf", b, 0644)
	// fetch latest to get responsive label URL
	latest2, _ := c.GetShipment(ctx, s.ID)
	if latest2.ResponsiveLabelURL != "" {
		html, _ := c.DownloadResponsiveURL(ctx, latest2.ResponsiveLabelURL)
		_ = os.WriteFile("label.html", []byte(html), 0644)
	}
}
