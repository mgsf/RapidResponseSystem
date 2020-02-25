// Package config provides access to relevant runtime settings.
package config

import (
	"os"
)

const (
	// EnvironmentDev is the value of "Environment" that will be set when the application is running in development mode.
	EnvironmentDev = "dev"
)

var (
	c Config
)

// Config holds the current runtime configuration options.
type Config struct {
	Environment string
	StaticRoot  string
	ViewRoot    string
}

func init() {
	setupConfig()
}

func setupConfig() {
	c.Environment = os.Getenv("ENV")
	c.StaticRoot = os.Getenv("STATIC_ROOT")
	c.ViewRoot = os.Getenv("VIEW_ROOT")
}

// Get returns a struct with the current runtime configuration.
func Get() Config {
	return c
}
