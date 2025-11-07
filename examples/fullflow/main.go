package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/GeliverApp/geliver-go/pkg/geliver"
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
	sender, err := c.CreateSenderAddress(ctx, geliver.CreateAddressRequest{
		Name: "ACME Inc.", Email: "ops@acme.test", Phone: &senderPhone,
		Address1: "Street 1", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
		DistrictName: "Esenyurt", DistrictID: 107605, Zip: "34020",
	})

	if err != nil || sender == nil {
		fmt.Println("create sender error", err)
		return
	}

	fmt.Println("sender address ID:", sender.ID)
	// Inline alıcı adresi (kayıt oluşturmadan)
	recipientPhone := "+905051234568"
	length, width, height, weight := "10.0", "10.0", "10.0", "1.0"
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
		// Print detailed API error if available
		if apiErr, ok := err.(*geliver.APIError); ok {
			fmt.Println("create shipment API error:", apiErr.Status, apiErr.Code, apiErr.Body)
		} else {
			fmt.Println("create shipment error", err)
		}
		return
	}

	// Etiket indirme: Teklif kabulünden sonra (Transaction) gelen URL'leri kullanabilirsiniz de; URL'lere her shipment nesnesinin içinden ulaşılır.

	// Teklifler create yanıtında hazır olabilir; önce onu kontrol edin
	offers := s.Offers
	if !(offers.PercentageCompleted == 100 || offers.Cheapest != nil) {
		for {
			gs, err := c.GetShipment(ctx, s.ID)
			if err != nil || gs == nil {
				fmt.Println("fetch shipment error", err)
				return
			}
			if gs.Offers.PercentageCompleted >= 100 && gs.Offers.Cheapest != nil {
				offers = gs.Offers
				break
			}
			time.Sleep(time.Second)
		}
	}
	// Accept offer

	var trx *geliver.Transaction
	if offers.Cheapest != nil {
		trx, _ = c.AcceptOffer(ctx, offers.Cheapest.ID)
	}

	if trx != nil && trx.Shipment != nil && trx.Shipment.LabelURL != "" {
		b, _ := c.DownloadURL(ctx, trx.Shipment.LabelURL)
		_ = os.WriteFile("label.pdf", b, 0644)
	}
	if trx != nil && trx.Shipment != nil && trx.Shipment.ResponsiveLabelURL != "" {
		html, _ := c.DownloadResponsiveURL(ctx, trx.Shipment.ResponsiveLabelURL)
		_ = os.WriteFile("label.html", []byte(html), 0644)
	}

	if trx.Shipment.Barcode != "" {
		fmt.Println("barcode:", trx.Shipment.Barcode)
	}
	if trx.Shipment.TrackingNumber != "" {
		fmt.Println("tracking number:", trx.Shipment.TrackingNumber)
	}
	if trx.Shipment.LabelURL != "" {
		fmt.Println("label:", trx.Shipment.LabelURL)
	}
	if trx.Shipment.TrackingURL != "" {
		fmt.Println("tracking:", trx.Shipment.TrackingURL)
	}

	// Test gönderilerinde her GET /shipments isteği kargo durumunu bir adım ilerletir; prod'da webhook önerilir.
	/*

		tracked, _ := c.GetShipment(ctx, s.ID)
		fmt.Println("tracking number (refresh):", tracked.TrackingNumber)
		if tracked.TrackingStatus != nil {
			fmt.Println("final status:", tracked.TrackingStatus.TrackingStatusCode, tracked.TrackingStatus.TrackingSubStatusCode)
		}
		latest, _ := c.GetShipment(ctx, s.ID)
		if latest.TrackingStatus != nil {
			fmt.Println("status:", latest.TrackingStatus.TrackingStatusCode, latest.TrackingStatus.TrackingSubStatusCode)
		}*/

}
