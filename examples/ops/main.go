package main

import (
    "context"
    "fmt"
    "os"
    "time"

    g "github.com/GeliverApp/geliver-go/pkg/geliver"
)

func ptrs(s string) *string { return &s }
func ptrb(b bool) *bool { return &b }

func main() {
    token := os.Getenv("GELIVER_TOKEN")
    if token == "" {
        fmt.Println("GELIVER_TOKEN is required")
        os.Exit(1)
    }
    c := g.NewClient(token)
    ctx := context.Background()

    // Create sender address once and reuse its ID
    phone := "+905051234567"
    sender, err := c.CreateSenderAddress(ctx, g.CreateAddressRequest{
        Name: "SDK Ops Sender", Email: "ops@example.com", Phone: &phone,
        Address1: "Hasan mah.", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
        DistrictName: "Esenyurt", Zip: ptrs("34020"),
    })
    if err != nil || sender == nil {
        fmt.Println("create sender error:", err)
        return
    }
    fmt.Println("sender:", sender.ID)

    // Create a test shipment with inline recipient
    length, width, height, weight := "10.0", "10.0", "10.0", "1.0"
    s, err := c.CreateShipmentWithRecipientAddress(ctx, g.CreateShipmentWithRecipientAddress{
        CreateShipmentRequestBase: g.CreateShipmentRequestBase{
            SenderAddressID: sender.ID,
            Length: &length, Width: &width, Height: &height, DistanceUnit: ptrs("cm"),
            Weight: &weight, MassUnit: ptrs("kg"),
            // Normal create shipment example: explicitly set to false. Set true for kapıda ödeme.
            ProductPaymentOnDelivery: ptrb(false),
            Test: ptrb(true),
        },
        RecipientAddress: g.Address{
            Name: "SDK Ops Recipient", Email: "recipient@example.com", Phone: "+905051234568",
            Address1: "A Mah.", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
            DistrictName: "Esenyurt",
        },
    })
    if err != nil || s == nil {
        if apiErr, ok := err.(*g.APIError); ok {
            fmt.Println("create shipment API error:", apiErr.Status, apiErr.Code, apiErr.Message, apiErr.AdditionalMessage)
        } else {
            fmt.Println("create shipment error:", err)
        }
        return
    }
    fmt.Println("shipment:", s.ID)

    // List shipments (pagination)
    limit := 5; page := 1
    list, err := c.ListShipments(ctx, &g.ListParams{ Limit: &limit, Page: &page })
    if err != nil {
        fmt.Println("list error:", err)
        return
    }
    fmt.Println("list count:", len(list.Data))

    // Get shipment
    got, err := c.GetShipment(ctx, s.ID)
    if err != nil {
        fmt.Println("get error:", err)
        return
    }
    fmt.Println("get status:", got.StatusCode)
    if got.TrackingStatus != nil {
        fmt.Println("tracking:", got.TrackingStatus.TrackingStatusCode, got.TrackingStatus.TrackingSubStatusCode)
    }

    // Update package
    nLen, nWid, nHei, nWei := "12.0", "12.0", "10.0", "1.2"
    upd := g.UpdatePackageRequest{ Length: &nLen, Width: &nWid, Height: &nHei, Weight: &nWei, DistanceUnit: ptrs("cm"), MassUnit: ptrs("kg") }
    upds, err := c.UpdatePackageTyped(ctx, s.ID, upd)
    if err != nil {
        fmt.Println("update error:", err)
        return
    }
    fmt.Println("updated length:", upds.Length)

    // Clone (Klonla)
    clone, err := c.CloneShipment(ctx, s.ID)
    if err != nil {
        fmt.Println("clone error:", err)
        return
    }
    fmt.Println("cloned shipment:", clone.ID)

    // Cancel original shipment
    canceled, err := c.CancelShipment(ctx, s.ID)
    if err != nil {
        fmt.Println("cancel error:", err)
        return
    }
    fmt.Println("canceled status:", canceled.StatusCode)

    // Small delay to avoid tight rate limits in examples
    time.Sleep(250 * time.Millisecond)
}
