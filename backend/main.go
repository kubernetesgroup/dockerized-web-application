package main

import (
"encoding/json"
"log"
"net/http"
"os"
"time"
)

type StatusResponse struct {
Service   string `json:"service"`
Status    string `json:"status"`
Timestamp string `json:"timestamp"`
Version   string `json:"version"`
}

func main() {
port := os.Getenv("PORT")
if port == "" {
port = "5000"
}

http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(StatusResponse{
Service:   "Kubexis Core API",
Status:    "healthy",
Timestamp: time.Now().UTC().Format(time.RFC3339),
Version:   "2.4.0-prod",
})
})

log.Printf(" API Server running on port %s", port)
log.Fatal(http.ListenAndServe(":"+port, nil))
}
