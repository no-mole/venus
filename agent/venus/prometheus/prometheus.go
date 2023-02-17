package prometheus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
	"time"
)

var (
// DefaultBuckets prometheus buckets in seconds.
//DefaultBuckets = []float64{0.3, 1.2, 5.0}
)

const (
	reqsName              = "http_requests_total"
	latencyName           = "http_request_duration_seconds"
	httpClientReqsName    = "http_client_requests_total"
	httpClientLatencyName = "http_client_duration_seconds"
)

// Prometheus is a handler that exposes prometheus metrics for the number of requests,
// the latency and the response size, partitioned by status code, method and HTTP path.
//
// Usage: pass its `ServeHTTP` to a route or globally.
type Prometheus struct {
	reqs    *prometheus.CounterVec
	latency *prometheus.HistogramVec
	listen  string

	httpClientReqs    *prometheus.CounterVec
	httpClientLatency *prometheus.HistogramVec
	redisErrorReqs    *prometheus.CounterVec
}

type log interface {
	Info(v ...interface{})
}

// NewPrometheus  returns a new prometheus middleware.
// If buckets are empty then `DefaultBuckets` are set.
func NewPrometheus(name, listen string) *Prometheus {
	p := Prometheus{}
	if listen == "" {
		listen = ":9090"
	}
	p.listen = listen
	p.reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        reqsName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"service": name},
		},
		[]string{"http_code", "code", "method", "path"},
	)
	prometheus.MustRegister(p.reqs)
	p.latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        latencyName,
		Help:        "How long it took to process the request, partitioned by status code, method and HTTP path.",
		ConstLabels: prometheus.Labels{"service": name},
	},
		[]string{"http_code", "code", "method", "path"},
	)
	prometheus.MustRegister(p.latency)

	p.httpClientReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        httpClientReqsName,
			Help:        "",
			ConstLabels: prometheus.Labels{"service": name},
		},
		[]string{"domain", "http_code", "protocol", "method"},
	)
	prometheus.MustRegister(p.httpClientReqs)

	p.httpClientLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        httpClientLatencyName,
		Help:        "",
		ConstLabels: prometheus.Labels{"service": name},
	},
		[]string{"domain", "http_code", "protocol", "method"},
	)
	prometheus.MustRegister(p.httpClientLatency)

	return &p
}

// NewPrometheusHandle .
func NewPrometheusHandle(p *Prometheus) func(*gin.Context) {
	http.Handle("/", promhttp.Handler())
	go func() {
		_ = http.ListenAndServe(p.listen, nil)
	}()

	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		httpCode := ctx.Writer.Status()
		code := "0"
		if value, exists := ctx.Get("code"); exists {
			code = fmt.Sprint(value)
		}

		path := ctx.FullPath()
		p.reqs.WithLabelValues(strconv.Itoa(httpCode), code, ctx.Request.Method, path).
			Inc()

		p.latency.WithLabelValues(strconv.Itoa(httpCode), code, ctx.Request.Method, path).
			Observe(float64(time.Since(start).Nanoseconds()) / 1000000000)
	}
}

func (p *Prometheus) HttpClientWithLabelValues(domain, httpCode, protocol, method string, starTime time.Time) {
	p.httpClientReqs.WithLabelValues(domain, httpCode, protocol, method).Inc()
	p.httpClientLatency.WithLabelValues(domain, httpCode, protocol, method).Observe(float64(time.Since(starTime).Nanoseconds()) / 1000000000)
}
