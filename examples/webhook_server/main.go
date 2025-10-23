package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/webhooks/geliver", func(w http.ResponseWriter, r *http.Request) {
        body, _ := io.ReadAll(r.Body)
        // TODO: verify signature (disabled for now)
        var evt map[string]any
        _ = json.Unmarshal(body, &evt)
        w.WriteHeader(200)
        _, _ = w.Write([]byte("ok"))
    })
    log.Println("listening :3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}

