package main

import (
    "log"
    "github.com/saitama-op/opnsense-metrics-exporter/pkg/server"
)

func main() {
    log.Println("Starting OPNsense metrics exporter...")
    server.Start()
}
