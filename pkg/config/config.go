package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       int    `mapstructure:"server_port"`
	APIUrl           string `mapstructure:"api_url"`
	APIKey           string `mapstructure:"api_key"`
	APISecret        string `mapstructure:"api_secret"`
	PrometheusPort   int    `mapstructure:"prometheus_port"`
	LogLevel         string `mapstructure:"log_level"`
	Timeout          int    `mapstructure:"timeout"`
	RetryCount       int    `mapstructure:"retry_count"`
	RetryDelay       int    `mapstructure:"retry_delay"`
	IgnoreSSLWarning bool   `mapstructure:"ignore_ssl_warning"`

	// Add more config fields here
}

var Cfg Config

func LoadConfig(path string) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")   // or "json", "toml", etc
	viper.AddConfigPath(path)     // look for config in the given path
	viper.AddConfigPath(".")      // and also in the current directory

	viper.AutomaticEnv() // override with env vars if available

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
