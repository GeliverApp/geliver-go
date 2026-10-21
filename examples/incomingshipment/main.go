package main

import (
	"context"
	"fmt"
	"os"

	"github.com/GeliverApp/geliver-go/pkg/geliver"
)

func main() {
	token := os.Getenv("GELIVER_TOKEN")
	recipientAddressID := os.Getenv("GELIVER_MERCHANT_ADDRESS_ID")
	if token == "" || recipientAddressID == "" {
		fmt.Println("Set GELIVER_TOKEN and GELIVER_MERCHANT_ADDRESS_ID.")
		return
	}

	isReturn, isTest := true, true
	length, width, height, weight := "10", "10", "10", "1"
	distanceUnit, massUnit := "cm", "kg"
	c := geliver.NewClient(token)
	shipment, err := c.CreateShipmentWithRecipientID(context.Background(), geliver.CreateShipmentWithRecipientID{
		CreateShipmentRequestBase: geliver.CreateShipmentRequestBase{
			IsReturn: &isReturn,
			Test:     &isTest,
			SenderAddress: &geliver.Address{
				Name: "Müşteri Adı", Email: "customer@example.com", Phone: "+905051234567",
				Address1: "Örnek Mahallesi 1", CountryCode: "TR", CityCode: "34", DistrictName: "Kadıköy", Zip: "34000",
			},
			Length: &length, Width: &width, Height: &height, DistanceUnit: &distanceUnit,
			Weight: &weight, MassUnit: &massUnit,
		},
		RecipientAddressID: recipientAddressID,
	})
	if err != nil {
		fmt.Println("create shipment error:", err)
		return
	}
	fmt.Println("Gelecek gönderi:", shipment.ID)
}
