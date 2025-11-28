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

	sender, err := c.CreateSenderAddress(ctx, g.CreateAddressRequest{
		Name: "OwnAg Sender", Email: "sender@example.com", Phone: "+905000000097",
		Address1: "Hasan Mahallesi", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
		DistrictName: "Esenyurt", Zip: ptrs("34020"),
	})
	if err != nil || sender == nil {
		fmt.Println("sender error:", err)
		return
	}

	length, width, height, weight := "10.0", "10.0", "10.0", "1.0"
	provider := "SURAT_STANDART"
	account := "c0dfdb42-012d-438c-9d49-98d13b4d4a2b"
	req := g.CreateShipmentWithRecipientAddress{
		CreateShipmentRequestBase: g.CreateShipmentRequestBase{
			SenderAddressID: sender.ID,
			Length:          &length, Width: &width, Height: &height, DistanceUnit: ptrs("cm"),
			Weight: &weight, MassUnit: ptrs("kg"),
			ProviderServiceCode: &provider,
			ProviderAccountID:   &account,
		},
		RecipientAddress: g.Address{
			Name: "OwnAg Recipient", Phone: "+905000000002", Address1: "Dest 2", CountryCode: "TR",
			CityName: "Istanbul", CityCode: "34", DistrictName: "Esenyurt",
		},
	}
	tx, err := c.CreateTransactionWithRecipientAddress(ctx, req)
	if err != nil {
		fmt.Println("create transaction (own agreement) error:", err)
		return
	}
	fmt.Println("transaction id:", tx.ID)
}
