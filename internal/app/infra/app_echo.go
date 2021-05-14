package infra

import (
	"time"

	"github.com/immanoj16/taskbook/pkg/logruskit"
	"github.com/labstack/echo/v4"
)

type (
	// AppCfg application configuration
	AppCfg struct {
		Address      string        `envconfig:"ADDRESS" default:":8089" required:"true"`
		ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
		WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
		Debug        bool          `envconfig:"DEBUG" default:"true"`
	}
)

// NewEcho returns  a new instance of server
func NewEcho(cfg *AppCfg) *echo.Echo {
	e := echo.New()
	logger := SetLogger(cfg.Debug)

	e.HideBanner = true
	e.Debug = cfg.Debug
	e.Logger = logruskit.EchoLogger(logger)
	return e
}
