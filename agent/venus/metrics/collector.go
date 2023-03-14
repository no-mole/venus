package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/uber-go/tally/v4"
	promreporter "github.com/uber-go/tally/v4/prometheus"
	"io"
	"time"
)

var (
	r         promreporter.Reporter
	Collector *PrometheusCollector
)

func init() {
	Collector = NewMetricsCollector("venus", 1*time.Second)
}

type PrometheusCollector struct {
	scope tally.Scope
	io.Closer
}

func NewMetricsCollector(prefix string, interval time.Duration) *PrometheusCollector {
	r = promreporter.NewReporter(promreporter.Options{})
	// Note: `promreporter.DefaultSeparator` is "_".
	// Prometheus doesn't like metrics with "." or "-" in them.
	scope, closer := tally.NewRootScope(tally.ScopeOptions{
		Prefix:         prefix,
		Tags:           map[string]string{},
		CachedReporter: r,
		Separator:      promreporter.DefaultSeparator,
	}, interval)

	return &PrometheusCollector{
		scope:  scope,
		Closer: closer,
	}
}

func (c *PrometheusCollector) Close() {
	_ = c.Closer.Close()
}

func (c *PrometheusCollector) count(name string, tags map[string]string) {
	counter := c.scope.Tagged(tags).Counter(name)
	counter.Inc(1)
}

func (c *PrometheusCollector) gauge(name string, value float64, tags map[string]string) {
	gauge := c.scope.Tagged(tags).Gauge(name)
	gauge.Update(value)
}

func (c *PrometheusCollector) histogram(name string, tags map[string]string, Buckets tally.Buckets) tally.Histogram {
	if Buckets == nil {
		Buckets = tally.DefaultBuckets
	}

	return c.scope.Tagged(tags).Histogram(name, Buckets)
}

func (c *PrometheusCollector) timer(name string, tags map[string]string) tally.Timer {
	return c.scope.Tagged(tags).Timer(name)
}

func (c *PrometheusCollector) HttpHandler() func(*gin.Context) {
	return func(ctx *gin.Context) {
		r.HTTPHandler().ServeHTTP(ctx.Writer, ctx.Request)
	}
}
