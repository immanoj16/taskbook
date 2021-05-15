package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	// Config application configuration
	Config struct {
		Address      string        `envconfig:"ADDRESS" default:":8089" required:"true"`
		ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
		WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
		Debug        bool          `envconfig:"DEBUG" default:"true"`
	}
)

// FromEnv load configuration from env vars
func FromEnv(prefix string) (Config, error) {
	c := Config{}
	return c, envconfig.Process(prefix, &c)
}
