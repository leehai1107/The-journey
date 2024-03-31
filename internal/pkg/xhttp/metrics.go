package xhttp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type outgoingMetrics struct {
	duration    *prometheus.HistogramVec
	requests    *prometheus.CounterVec
	dnsDuration *prometheus.HistogramVec
	tlsDuration *prometheus.HistogramVec
	inflight    prometheus.Gauge
}

// Describe implements prometheus.Collector interface.
func (i *outgoingMetrics) Describe(in chan<- *prometheus.Desc) {
	i.duration.Describe(in)
	i.requests.Describe(in)
	i.dnsDuration.Describe(in)
	i.tlsDuration.Describe(in)
	i.inflight.Describe(in)
}

// Collect implements prometheus.Collector interface.
func (i *outgoingMetrics) Collect(in chan<- prometheus.Metric) {
	i.duration.Collect(in)
	i.requests.Collect(in)
	i.dnsDuration.Collect(in)
	i.tlsDuration.Collect(in)
	i.inflight.Collect(in)
}

func NewOutgoingMetrics(subsystem string, constLabels map[string]string) *outgoingMetrics {
	namespace := defaultNamespace
	return &outgoingMetrics{
		requests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "http_outgoing_requests_total",
				Help:        "A counter for outgoing requests from the wrapped client.",
				ConstLabels: constLabels,
			},
			[]string{"code", "method", "path"},
		),
		duration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "http_outgoing_request_duration_histogram_seconds",
				Help:        "A histogram of outgoing request latencies.",
				Buckets:     prometheus.DefBuckets,
				ConstLabels: constLabels,
			},
			[]string{"method", "path"},
		),
		dnsDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "http_outgoing_dns_duration_histogram_seconds",
				Help:        "Trace dns latency histogram.",
				Buckets:     []float64{.005, .01, .025, .05},
				ConstLabels: constLabels,
			},
			[]string{"event"},
		),
		tlsDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "http_outgoing_tls_duration_histogram_seconds",
				Help:        "Trace tls latency histogram.",
				Buckets:     []float64{.05, .1, .25, .5},
				ConstLabels: constLabels,
			},
			[]string{"event"},
		),
		inflight: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace:   namespace,
			Subsystem:   subsystem,
			Name:        "http_outgoing_in_flight_requests",
			Help:        "A gauge of in-flight outgoing requests for the wrapped client.",
			ConstLabels: constLabels,
		}),
	}
}

func buildTraceTransport(transport http.RoundTripper, metrics *outgoingMetrics) http.RoundTripper {
	trace := &promhttp.InstrumentTrace{
		DNSStart: func(t float64) {
			metrics.dnsDuration.WithLabelValues("dns_start").Observe(t)
		},
		DNSDone: func(t float64) {
			metrics.dnsDuration.WithLabelValues("dns_done").Observe(t)
		},
		TLSHandshakeStart: func(t float64) {
			metrics.tlsDuration.WithLabelValues("tls_handshake_start").Observe(t)
		},
		TLSHandshakeDone: func(t float64) {
			metrics.tlsDuration.WithLabelValues("tls_handshake_done").Observe(t)
		},
	}
	roundTripper := promhttp.InstrumentRoundTripperInFlight(metrics.inflight,
		promhttp.InstrumentRoundTripperTrace(trace,
			instrumentRoundTripperCounter(metrics.requests,
				instrumentRoundTripperDuration(metrics.duration, transport))))
	return roundTripper
}

func instrumentRoundTripperCounter(requestCounter *prometheus.CounterVec,
	next http.RoundTripper) promhttp.RoundTripperFunc {
	return func(req *http.Request) (*http.Response, error) {
		rsp, err := next.RoundTrip(req)
		if err == nil {
			groupPath := extractGroupPath(req)
			requestCounter.WithLabelValues(fmt.Sprint(rsp.StatusCode), req.Method, groupPath).Inc()
		}
		return rsp, err
	}
}

func instrumentRoundTripperDuration(requestDuration *prometheus.HistogramVec,
	next http.RoundTripper) promhttp.RoundTripperFunc {
	return func(req *http.Request) (*http.Response, error) {
		start := time.Now()
		resp, err := next.RoundTrip(req)
		if err == nil {
			groupPath := extractGroupPath(req)
			requestDuration.WithLabelValues(req.Method, groupPath).Observe(time.Since(start).Seconds())
		}
		return resp, err
	}
}

func extractGroupPath(req *http.Request) string {
	groupPathHeader := req.Header.Get(GroupPathHeader)
	if groupPathHeader != "" {
		return groupPathHeader
	}
	if groupPath, ok := req.Context().Value(groupPath).(string); ok {
		return groupPath
	}
	return "unknown"
}
