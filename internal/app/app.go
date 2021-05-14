package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/immanoj16/taskbook/internal/app/infra"
	"github.com/immanoj16/taskbook/pkg/echokit"
	"github.com/immanoj16/taskbook/pkg/errkit"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Config *infra.AppCfg
	Echo   *echo.Echo
	DB     *sql.DB `name:"pg"`
}

// NewServer returns new server
func NewServer() (*Server, error) {
	var cfg infra.AppCfg
	prefix := "APP"
	if err := envconfig.Process(prefix, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	e := infra.NewEcho(&cfg)
	return &Server{
		Config: &cfg,
		Echo:   e,
	}, nil
}

// Start starts the echo server
func (s *Server) Start() error {
	setServer(s.Echo)
	SetProfiler(s.Echo, HealthCheck{})
	if s.Config.Debug {
		routes := echokit.DumpEcho(s.Echo)
		logrus.Debugf("Print routes:\n  %s\n\n", strings.Join(routes, "\n  "))
	}
	return s.Echo.StartServer(&http.Server{
		Addr:         s.Config.Address,
		ReadTimeout:  s.Config.ReadTimeout,
		WriteTimeout: s.Config.WriteTimeout,
	})
}

func setServer(e *echo.Echo) {
	e.Use(infra.LogMiddleware)
	e.Use(middleware.Recover())
}

// Shutdown shutdown all services
func (s *Server) Shutdown() error {
	fmt.Printf("Shutdown at %s", time.Now().String())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	errs := errkit.Errors{
		//s.DB.Close(),
		s.Echo.Shutdown(ctx),
	}

	return errs.Unwrap()
}
