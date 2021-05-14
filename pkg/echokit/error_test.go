package echokit_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/immanoj16/taskbook/pkg/echokit"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestHTTPError(t *testing.T) {
	testCases := []struct {
		TestName string
		err      error
		expected *echo.HTTPError
	}{
		{
			err:      errors.New("some-error"),
			expected: echo.NewHTTPError(http.StatusInternalServerError, "some-error"),
		},
		{
			err:      echo.NewHTTPError(99, "some-message"),
			expected: echo.NewHTTPError(99, "some-message"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.expected, echokit.HTTPError(tt.err))
		})
	}
}

func TestNewValidErr(t *testing.T) {
	require.Equal(t, echo.NewHTTPError(422, "some-message"), echokit.NewValidErr("some-message"))
}
