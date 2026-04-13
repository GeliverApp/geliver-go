package main

import (
	"context"
	"fmt"
	"os"

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

	tx, err := c.CreateReturnTransaction(ctx, shipmentID, geliver.ReturnShipmentRequest{})
	if err != nil || tx == nil {
		fmt.Println("create return transaction error:", err)
		return
	}

	fmt.Println("transaction:", tx.ID)
	if tx.Shipment != nil {
		fmt.Println("purchased return shipment:", tx.Shipment.ID)
	}
}
