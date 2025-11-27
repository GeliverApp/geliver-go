package main

import (
	"context"
	"fmt"
	"os"

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
    Name: "Şirketim A.Ş", Email: "ops@sirketim.net", Phone: &senderPhone,
    Address1: "Hasan mahallesi", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
    DistrictName: "Esenyurt", Zip: ptrs("34020"),
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
            SenderAddressID: sender.ID,
            Length:          &length, Width: &width, Height: &height, DistanceUnit: ptrs("cm"), Weight: &weight, MassUnit: ptrs("kg"),
            // Normal flow: explicitly set to false; set true for kapıda ödeme (Payment on Delivery)
            ProductPaymentOnDelivery: ptrb(false),
            Test: ptrb(true),
            Order: &geliver.OrderRequest{OrderNumber: "ABC12333322", SourceIdentifier: ptrs("https://magazaadresiniz.com"), TotalAmount: ptrs("150"), TotalAmountCurrency: ptrs("TL")},
        },
        RecipientAddress: geliver.Address{
            Name: "Ahmet Mehmet", Email: "ahmetmehmet@ornek.com", Phone: recipientPhone,
            Address1: "Hasan mahallesi", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
            DistrictName: "Esenyurt",
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

	// Teklifler create yanıtındaki offers alanında gelir
	offers := s.Offers
	if offers.Cheapest == nil {
		fmt.Println("No cheapest offer available (henüz hazır değil)")
		return
	}

	trx, err := c.AcceptOffer(ctx, offers.Cheapest.ID)
	if err != nil {
		if apiErr, ok := err.(*geliver.APIError); ok {
			fmt.Println("accept offer API error:", apiErr.Status, apiErr.Code, apiErr.Body)
		} else {
			fmt.Println("accept offer error:", err)
		}
		return
	}

	// Etiket indirme: LabelFileType kontrolü
	// Eğer LabelFileType "PROVIDER_PDF" ise, LabelURL'den indirilen PDF etiket kullanılmalıdır.
	// Eğer LabelFileType "PDF" ise, responsiveLabelURL (HTML) dosyası kullanılabilir.
	if trx.Shipment != nil {
		if trx.Shipment.LabelFileType == "PROVIDER_PDF" {
			// PROVIDER_PDF: Sadece PDF etiket kullanılmalı
			if trx.Shipment.LabelURL != "" {
				b, _ := c.DownloadURL(ctx, trx.Shipment.LabelURL)
				_ = os.WriteFile("label.pdf", b, 0644)
				fmt.Println("PDF etiket indirildi (PROVIDER_PDF)")
			}
		} else if trx.Shipment.LabelFileType == "PDF" {
			// PDF: ResponsiveLabel (HTML) kullanılabilir
			if trx.Shipment.ResponsiveLabelURL != "" {
				html, _ := c.DownloadResponsiveURL(ctx, trx.Shipment.ResponsiveLabelURL)
				_ = os.WriteFile("label.html", []byte(html), 0644)
				fmt.Println("HTML etiket indirildi (PDF)")
			}
			// İsteğe bağlı olarak PDF de indirilebilir
			if trx.Shipment.LabelURL != "" {
				b, _ := c.DownloadURL(ctx, trx.Shipment.LabelURL)
				_ = os.WriteFile("label.pdf", b, 0644)
			}
		}
	}

	if trx.Shipment != nil {
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
