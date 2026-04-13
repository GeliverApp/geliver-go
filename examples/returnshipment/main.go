package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/GeliverApp/geliver-go/pkg/geliver"
)

func main() {
	token := os.Getenv("GELIVER_TOKEN")
	shipmentID := os.Getenv("GELIVER_RETURN_SHIPMENT_ID")
	if shipmentID == "" && len(os.Args) > 1 {
		shipmentID = os.Args[1]
	}
	if token == "" || shipmentID == "" {
		fmt.Println("Set GELIVER_TOKEN and GELIVER_RETURN_SHIPMENT_ID, or pass the shipment ID as the first argument.")
		return
	}

	c := geliver.NewClient(token)
	ctx := context.Background()

	returned, err := c.CreateReturnShipment(ctx, shipmentID, geliver.ReturnShipmentRequest{})
	if err != nil || returned == nil {
		fmt.Println("create return shipment error:", err)
		return
	}

	fmt.Println("return shipment:", returned.ID)
	fmt.Println("Label is not purchased yet. This example waits for offers and buys it with AcceptOffer.")

	current := returned
	deadline := time.Now().Add(60 * time.Second)
	for current.Offers == nil || current.Offers.Cheapest == nil || current.Offers.Cheapest.ID == "" {
		if time.Now().After(deadline) {
			fmt.Println("timed out waiting for return offers")
			return
		}
		progress := 0.0
		if current.Offers != nil {
			progress = current.Offers.PercentageCompleted
		}
		fmt.Println("waiting offers...", progress, "%")
		time.Sleep(time.Second)
		current, err = c.GetShipment(ctx, returned.ID)
		if err != nil || current == nil {
			fmt.Println("get return shipment error:", err)
			return
		}
	}

	tx, err := c.AcceptOffer(ctx, current.Offers.Cheapest.ID)
	if err != nil || tx == nil {
		fmt.Println("accept offer error:", err)
		return
	}

	fmt.Println("transaction:", tx.ID)
	if tx.Shipment != nil {
		fmt.Println("purchased return shipment:", tx.Shipment.ID)
	}
}
