package main

import (
	"flag"
	"log"

	"github.com/saitama-op/opnsense-metrics-exporter/pkg/config"
	"github.com/saitama-op/opnsense-metrics-exporter/pkg/server"
)

func main() {
	configfile := flag.String("config_path", "../", "Path to the configuration file")
	// Load configuration from the specified path
	config.LoadConfig(*configfile)
	log.Println("Starting OPNsense metrics exporter...")
	server.Start()
}
