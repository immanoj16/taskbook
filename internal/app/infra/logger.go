package infra

import (
	"github.com/immanoj16/taskbook/pkg/logruskit"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

var _debug bool

// SetLogger sets logger
func SetLogger(debug bool) *logrus.Logger {
	_debug = debug
	logger := logrus.StandardLogger()
	if debug {
		logger.SetLevel(logrus.DebugLevel)
		logger.SetFormatter(&logrus.TextFormatter{})
	} else {
		logger.SetLevel(logrus.WarnLevel)
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	log.SetOutput(logger.Writer())
	return logger
}

// LogMiddleware to log every request
func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		start := time.Now()

		// generate request ID if not exist
		reqID := req.Header.Get(echo.HeaderXRequestID)
		if reqID == "" {
			reqID = generateRequestID()
		}
		res.Header().Set(echo.HeaderXRequestID, reqID)

		// Put fields in context
		ctx := req.Context()
		logruskit.PutField(&ctx, "req_id", reqID)

		// update request with new context
		req = req.WithContext(ctx)
		c.SetRequest(req)

		// current handler
		if err := next(c); err != nil {
			c.Error(err)
		}

		stop := time.Now()
		if _debug {
			logrus.WithFields(logrus.Fields{
				"exec_time":   stop.Sub(start).String(),
				"req_id":      reqID,
				"resp_status": res.Status,
				"resp_size":   strconv.FormatInt(res.Size, 10),
				"req_size":    req.Header.Get(echo.HeaderContentLength),
			}).Debugf("%s %s", req.Method, req.RequestURI)
		}
		return nil
	}
}

func generateRequestID() string {
	return xid.New().String()
}
