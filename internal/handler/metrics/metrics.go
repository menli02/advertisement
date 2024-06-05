package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "status_code"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)

}
func RegisterStandardMetricsHandler() {
	http.Handle("/metrics", promhttp.Handler())
}

func IncrementHTTPRequestTotal(method, statusCode string) {
	httpRequestsTotal.WithLabelValues(method, statusCode).Inc()
}
