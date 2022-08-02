package middlewares

import (
	"github.com/fd239/gopher_keeper/config"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
)

var (
	httpTotalRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_microservice_total_requests",
		Help: "The total number of incoming HTTP requests",
	})
)

// MiddlewareManager http middlewares
type middlewareManager struct {
	log *zap.SugaredLogger
	cfg *config.Config
}

// MiddlewareManager interface
type MiddlewareManager interface {
	Metrics(next echo.HandlerFunc) echo.HandlerFunc
}

// NewMiddlewareManager constructor
func NewMiddlewareManager(log *zap.SugaredLogger, cfg *config.Config) *middlewareManager {
	return &middlewareManager{log: log, cfg: cfg}
}

// Metrics prometheus metrics
func (m *middlewareManager) Metrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		httpTotalRequests.Inc()
		return next(c)
	}
}
