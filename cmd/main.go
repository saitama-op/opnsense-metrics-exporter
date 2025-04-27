package main

import (
    "log"
    "opnsense-metrics-exporter/pkg/server"
)

func main() {
    log.Println("Starting OPNsense metrics exporter...")
    server.Start()
}
