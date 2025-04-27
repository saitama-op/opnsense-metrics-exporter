package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/saitama-op/opnsense-metrics-exporter/pkg/api"
	"github.com/saitama-op/opnsense-metrics-exporter/pkg/config"
	"github.com/saitama-op/opnsense-metrics-exporter/pkg/metrics"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /metrics")

	// Fetch data from the OPNsense API
	data, err := api.FetchOPNSenseData(config.Cfg.APIUrl, config.Cfg.APIKey, config.Cfg.APISecret, config.Cfg.IgnoreSSLWarning)
	if err != nil {
		log.Printf("Error fetching OPNsense data: %v\n", err)
		http.Error(w, "Failed to fetch data from OPNsense API", http.StatusInternalServerError)
		return
	}

	// Validate and extract the metric value
	metricValue, ok := data["api_version"].(string) // Assuming the value is a string
	if !ok {
		log.Println("Invalid or missing 'api_version' in OPNsense data")
		http.Error(w, "Invalid data from OPNsense", http.StatusInternalServerError)
		return
	}
	log.Printf("Fetched Metric Value: %s\n", metricValue)

	// Update metrics dynamically (replace 1 with actual value if applicable)
	metrics.UpdateMetrics(1)

	// Serve Prometheus metrics
	promhttp.Handler().ServeHTTP(w, r)
}

func Start() {
	// Validate server port configuration
	if config.Cfg.ServerPort <= 0 || config.Cfg.ServerPort > 65535 {
		log.Fatalf("Invalid server port: %d. Please configure a valid port in the range 1-65535.", config.Cfg.ServerPort)
	}

	http.HandleFunc("/metrics", metricsHandler)
	log.Printf("Exporter running on port: %d\n", config.Cfg.ServerPort)

	// Start the HTTP server
	err := http.ListenAndServe(":"+strconv.Itoa(config.Cfg.ServerPort), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
