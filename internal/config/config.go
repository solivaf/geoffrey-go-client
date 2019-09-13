package config

import (
	"flag"
	"os"
)

type Config struct {
	ClientTimeout string
	Url           string
}

func GetConfig() *Config {
	url := os.Getenv("GEOFFREY_SERVER_URL")
	if url == "" {
		url = *flag.String("geoffrey.server.url", "", "Geoffrey config server url")
		if url == "" {
			panic("error initializing geoffrey client. missing server url configuration.")
		}
	}

	clientTimeout := os.Getenv("GEOFFREY_CLIENT_TIMEOUT_MS")
	if clientTimeout == "" {
		clientTimeout = *flag.String("geoffrey.client.timeout.ms", "", "Geoffrey client timeout")
		if clientTimeout == "" {
			clientTimeout = "5000"
		}
	}

	return &Config{Url: url, ClientTimeout: clientTimeout}
}
