package server

import (
    "fmt"
    "log"
    "net/http"
    "opnsense-metrics-exporter/pkg/api"
    "opnsense-metrics-exporter/pkg/metrics"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    apiUrl := "https://your-opnsense-url/api/v1/status"
    apiKey := "your-api-key"
    data, err := api.FetchOPNSenseData(apiUrl, apiKey)
    if err != nil {
        log.Println("Error fetching OPNsense data:", err)
        http.Error(w, "Error fetching data", http.StatusInternalServerError)
        return
    }

    // Assuming "status" field in the response contains the metric value
    metricValue, ok := data["status"].(float64)
    if !ok {
        http.Error(w, "Invalid data from OPNsense", http.StatusInternalServerError)
        return
    }

    metrics.UpdateMetrics(metricValue)
    promhttp.Handler().ServeHTTP(w, r)
}

func Start() {
    http.HandleFunc("/metrics", metricsHandler)
    log.Println("Exporter running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
