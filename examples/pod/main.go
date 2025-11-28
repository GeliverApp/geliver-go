package main

import (
	"context"
	"fmt"
	"os"

	g "github.com/GeliverApp/geliver-go/pkg/geliver"
)

func ptrs(s string) *string { return &s }
func ptrb(b bool) *bool     { return &b }

func main() {
	token := os.Getenv("GELIVER_TOKEN")
	if token == "" {
		fmt.Println("GELIVER_TOKEN is required")
		return
	}
	c := g.NewClient(token)
	ctx := context.Background()

	sender, err := c.CreateSenderAddress(ctx, g.CreateAddressRequest{
		Name: "POD Sender", Email: "sender@example.com", Phone: "+905000000098",
		Address1: "Hasan Mahallesi", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
		DistrictName: "Esenyurt", Zip: ptrs("34020"),
	})
	if err != nil || sender == nil {
		fmt.Println("sender error:", err)
		return
	}

	// Payment on Delivery: requires provider service supporting POD and order totals
	length, width, height, weight := "10.0", "10.0", "10.0", "1.0"
	prov := "PTT_KAPIDA_ODEME"
	total, currency := "150", "TL"
	order := g.OrderRequest{OrderNumber: "POD-12345", TotalAmount: &total, TotalAmountCurrency: &currency}
	req := g.CreateShipmentWithRecipientAddress{
		CreateShipmentRequestBase: g.CreateShipmentRequestBase{
			SenderAddressID: sender.ID,
			Length:          &length, Width: &width, Height: &height, DistanceUnit: ptrs("cm"),
			Weight: &weight, MassUnit: ptrs("kg"),
			ProviderServiceCode:      &prov,
			ProductPaymentOnDelivery: ptrb(true),
			Order:                    &order,
		},
		RecipientAddress: g.Address{
			Name: "POD Recipient", Phone: "+905000000001", Address1: "Atatürk Mahallesi", CountryCode: "TR",
			CityName: "Istanbul", CityCode: "34", DistrictName: "Esenyurt",
		},
	}
	tx, err := c.CreateTransactionWithRecipientAddress(ctx, req)
	if err != nil {
		fmt.Println("create transaction (POD) error:", err)
		return
	}
	fmt.Println("transaction id:", tx.ID)
}
