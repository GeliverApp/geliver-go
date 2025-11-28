package main

import (
	"context"
	"fmt"
	"os"

	g "github.com/GeliverApp/geliver-go/pkg/geliver"
)

func ptrs(s string) *string { return &s }

func main() {
	token := os.Getenv("GELIVER_TOKEN")
	if token == "" {
		fmt.Println("GELIVER_TOKEN is required")
		return
	}
	c := g.NewClient(token)
	ctx := context.Background()

	// Minimal sender address (reuse your existing ID in real world)
	sender, err := c.CreateSenderAddress(ctx, g.CreateAddressRequest{
		Name: "OneStep Sender", Email: "sender@example.com", Phone: "+905000000099",
		Address1: "Hasan Mahallesi", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
		DistrictName: "Esenyurt", Zip: ptrs("34020"),
	})
	if err != nil || sender == nil {
		fmt.Println("sender error:", err)
		return
	}

	// One-step: create transaction directly without accept flow
	length, width, height, weight := "10.0", "10.0", "10.0", "1.0"
	req := g.CreateShipmentWithRecipientAddress{
		CreateShipmentRequestBase: g.CreateShipmentRequestBase{
			SenderAddressID: sender.ID,
			Length:          &length, Width: &width, Height: &height, DistanceUnit: ptrs("cm"),
			Weight: &weight, MassUnit: ptrs("kg"),
		},
		RecipientAddress: g.Address{
			Name: "OneStep Recipient", Phone: "+905000000000", Address1: "Atatürk Mahallesi", CountryCode: "TR",
			CityName: "Istanbul", CityCode: "34", DistrictName: "Esenyurt",
		},
	}
	tx, err := c.CreateTransactionWithRecipientAddress(ctx, req)
	if err != nil {
		fmt.Println("create transaction error:", err)
		return
	}
	fmt.Println("transaction id:", tx.ID)
}
