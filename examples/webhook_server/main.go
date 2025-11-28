package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"

    g "github.com/GeliverApp/geliver-go/pkg/geliver"
)

func main() {
    http.HandleFunc("/webhooks/geliver", func(w http.ResponseWriter, r *http.Request) {
        body, _ := io.ReadAll(r.Body)
        // TODO: verify signature (disabled for now)
        var evt g.WebhookUpdateTrackingRequest
        if err := json.Unmarshal(body, &evt); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        if evt.Event == "TRACK_UPDATED" {
            log.Printf("Tracking update: %s %s\n", evt.Shipment.TrackingURL, evt.Shipment.TrackingNumber)
        }
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ok"))
    })
    log.Println("listening :3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
