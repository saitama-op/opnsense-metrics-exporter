package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    CustomMetric = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "opnsense_custom_metric",
            Help: "Custom metric from OPNsense data",
        },
        []string{"status"},
    )
)

func init() {
    prometheus.MustRegister(CustomMetric)
}

func UpdateMetrics(value float64) {
    CustomMetric.WithLabelValues("active").Set(value)
}
