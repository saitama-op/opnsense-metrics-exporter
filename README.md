# OPNsense Metrics Exporter

A custom metrics exporter for OPNsense, written in Go, to collect and expose system metrics for monitoring and visualization with Prometheus.

## Features

- Fetchs data from OPNsense API.
- Exposes Prometheus-compatible metrics at `/metrics` endpoint.
- Easy to configure for different OPNsense installations.

## Requirements

- Go 1.20+
- OPNsense API credentials

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/saitama-op/opnsense-metrics-exporter.git
   ```

2. Change to the project directory:
   ```bash
   cd opnsense-metrics-exporter
   ```

3. Build the project:
   ```bash
   go build -o opnsense-metrics-exporter cmd/main.go
   ```

4. Run the exporter:
   ```bash
   ./opnsense-metrics-exporter
   ```

5. Access the metrics at:
   ```bash
   http://localhost:8080/metrics
   ```

## License

MIT License
