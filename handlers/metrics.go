package handlers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func instrumentHandler(handleName string, next http.Handler) http.Handler {
	inFlightGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Subsystem:   "http",
		Name:        "in_flight_requests",
		Help:        "A gauge of requests currently being served",
		ConstLabels: prometheus.Labels{"handler": handleName},
	})

	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem:   "http",
			Name:        "requests_total",
			Help:        "A counter for requests served",
			ConstLabels: prometheus.Labels{"handler": handleName},
		},
		[]string{"code", "method"},
	)

	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem:   "http",
			Name:        "request_duration_seconds",
			Help:        "A histogram of latencies for requests",
			Buckets:     prometheus.ExponentialBuckets(0.01, 2, 11), // 10ms -> 10s
			ConstLabels: prometheus.Labels{"handler": handleName},
		},
		[]string{"code", "method"},
	)

	requestSize := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem:   "http",
			Name:        "request_size_bytes",
			Help:        "A histogram of request sizes for requests.",
			Buckets:     prometheus.ExponentialBuckets(2, 2, 15), // 2 bytes -> 32kb
			ConstLabels: prometheus.Labels{"handler": handleName},
		},
		[]string{"code", "method"},
	)

	responseSize := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem:   "http",
			Name:        "response_size_bytes",
			Help:        "A histogram of response sizes for requests.",
			Buckets:     prometheus.ExponentialBuckets(2, 2, 15), // 2 bytes -> 32kb
			ConstLabels: prometheus.Labels{"handler": handleName},
		},
		[]string{"code", "method"},
	)

	// Register all of the metrics in the standard registry.
	prometheus.MustRegister(
		inFlightGauge,
		requestCounter,
		requestDuration,
		requestSize,
		responseSize,
	)

	return promhttp.InstrumentHandlerInFlight(inFlightGauge,
		promhttp.InstrumentHandlerCounter(requestCounter,
			promhttp.InstrumentHandlerDuration(requestDuration,
				promhttp.InstrumentHandlerRequestSize(requestSize,
					promhttp.InstrumentHandlerResponseSize(responseSize, next),
				),
			),
		),
	)
}
