// Package config provides access to relevant runtime settings.
package config

import (
	"os"
)

var (
	c Config
)

// Config holds the current runtime configuration options.
type Config struct {
	StaticRoot string
	ViewRoot   string
}

func init() {
	setupConfig()
}

func setupConfig() {
	c.StaticRoot = os.Getenv("STATIC_ROOT")
	c.ViewRoot = os.Getenv("VIEW_ROOT")
}

// Get returns a struct with the current runtime configuration.
func Get() Config {
	return c
}
