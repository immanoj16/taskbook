package echokit_test

import (
	"github.com/immanoj16/taskbook/pkg/echokit"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDumpEcho(t *testing.T) {
	e := echo.New()
	echokit.SetRoute(e, echokit.NewRouter(func(s echokit.Server) {
		s.Any("/any", nil)
		s.POST("/post", nil)
		s.GET("/get", nil)
		s.PATCH("/patch", nil)
		s.DELETE("/delete", nil)
		s.PUT("/put", nil)
	}))
	require.Equal(t, []string{
		"/any\tCONNECT,DELETE,GET,HEAD,OPTIONS,PATCH,POST,PROPFIND,PUT,REPORT,TRACE",
		"/delete\tDELETE",
		"/get\tGET",
		"/patch\tPATCH",
		"/post\tPOST",
		"/put\tPUT",
	}, echokit.DumpEcho(e))
}
