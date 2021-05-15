package server

import (
	"fmt"
	"net/http"

	"github.com/immanoj16/taskbook/pkg/restkit"
	"github.com/labstack/echo/v4"
)

type HealthCheck struct {
}

const (
	healthCheckPath = "/application/health"
)

func SetProfiler(e *echo.Echo, hc HealthCheck) {
	e.GET(healthCheckPath, hc.Handle)
	e.HEAD(healthCheckPath, hc.Handle)
	e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
	e.GET("/debug/*/*", echo.WrapHandler(http.DefaultServeMux))
}

func (h *HealthCheck) Handle(ec echo.Context) error {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	health := restkit.HealthMap{
		"server": nil,
	}

	status, ok := health.Status()
	return ec.JSON(h.httpStatus(ok), h.response(status))
}

func (h *HealthCheck) httpStatus(ok bool) int {
	if ok {
		return http.StatusOK
	}
	return http.StatusServiceUnavailable
}

func (h *HealthCheck) response(status map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"name":   fmt.Sprintf("%s (%s)", "taskbook", "1.0.0"),
		"status": status,
	}
}
