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
	apiUrl := config.Cfg.APIUrl
	apiKey := config.Cfg.APIKey
	apiSecret := config.Cfg.APISecret
	data, err := api.FetchOPNSenseData(apiUrl, apiKey, apiSecret)
	if err != nil {
		log.Println("Error fetching OPNsense data: ", err)
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
	log.Println("Exporter running on port :", config.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Cfg.ServerPort), nil))
}
