package config

import (
	"log"
	"os"
)

type Config struct {
	LinkedInEmail    string
	LinkedInPassword string
	Headless         bool
}

func LoadConfig() *Config {
	cfg := &Config{
		LinkedInEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword: os.Getenv("LINKEDIN_PASSWORD"),
		Headless:         false,
	}

	if cfg.LinkedInEmail == "" || cfg.LinkedInPassword == "" {
		log.Fatal("LINKEDIN_EMAIL or LINKEDIN_PASSWORD not set in environment variables")
	}

	return cfg
}
