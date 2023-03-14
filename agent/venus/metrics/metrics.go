package metrics

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uber-go/tally/v4"
	"google.golang.org/grpc"
	"time"
)

const (
	HTTP_REQS_NAME    = "http_requests_total"
	HTTP_LATENCY_NAME = "http_request_duration_seconds"

	RPC_REQS_NAME    = "rpc_requests_total"
	RPC_LATENCY_NAME = "rpc_request_duration_seconds"

	RPC_STREAM_REQS_NAME    = "rpc_stream_requests_total"
	RPC_STREAM_LATENCY_NAME = "rpc_stream_request_duration_seconds"

	HAS_LEADER = "has_leader"
)

// HttpRequestTotal Record http request total
func (c *PrometheusCollector) HttpRequestTotal() func(*gin.Context) {
	return func(ctx *gin.Context) {
		c.count(HTTP_REQS_NAME, map[string]string{
			"path":      ctx.FullPath(),
			"http_code": fmt.Sprintf("%d", ctx.Writer.Status()),
			"method":    ctx.Request.Method,
		})
	}
}

// HttpRequestDurationTime Record http duration time
func (c *PrometheusCollector) HttpRequestDurationTime() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		histogram := c.histogram(HTTP_LATENCY_NAME, map[string]string{
			"path":      ctx.FullPath(),
			"http_code": fmt.Sprintf("%d", ctx.Writer.Status()),
			"method":    ctx.Request.Method,
		}, tally.DurationBuckets{
			0 * time.Millisecond,
			50 * time.Millisecond,
			250 * time.Millisecond,
			1000 * time.Millisecond,
			2500 * time.Millisecond,
			10000 * time.Millisecond,
		})
		hsw := histogram.Start()
		ctx.Next()
		hsw.Stop()

	}
}

// RpcRequestTotal Record rpc request total
func (c *PrometheusCollector) RpcRequestTotal() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		c.count(RPC_REQS_NAME, map[string]string{
			"method": method,
			"req":    fmt.Sprintf("%v", req),
		})
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// RpcRequestDurationTime Record rpc duration time
func (c *PrometheusCollector) RpcRequestDurationTime() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		histogram := c.histogram(RPC_LATENCY_NAME, map[string]string{
			"method": method,
			"req":    fmt.Sprintf("%v", req),
		}, tally.DurationBuckets{
			0 * time.Millisecond,
			50 * time.Millisecond,
			250 * time.Millisecond,
			1000 * time.Millisecond,
			2500 * time.Millisecond,
			10000 * time.Millisecond,
		})
		hsw := histogram.Start()
		err = invoker(ctx, method, req, reply, cc, opts...)
		hsw.Stop()
		return err
	}
}

func (c *PrometheusCollector) RpcStreamRequestTotal() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (clientStream grpc.ClientStream, err error) {
		c.count(RPC_STREAM_REQS_NAME, map[string]string{
			"method": method,
			"desc":   desc.StreamName,
		})
		return streamer(ctx, desc, cc, method)
	}
}

func (c *PrometheusCollector) RpcStreamRequestDurationTime() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (clientStream grpc.ClientStream, err error) {
		histogram := c.histogram(RPC_STREAM_LATENCY_NAME, map[string]string{
			"method": method,
			"desc":   desc.StreamName,
		}, tally.DurationBuckets{
			0 * time.Millisecond,
			50 * time.Millisecond,
			250 * time.Millisecond,
			1000 * time.Millisecond,
			2500 * time.Millisecond,
			10000 * time.Millisecond,
		})
		hsw := histogram.Start()
		clientStream, err = streamer(ctx, desc, cc, method)
		hsw.Stop()
		return
	}
}

func (c *PrometheusCollector) HasLeader(hasLeader bool, address, id string) {
	exist := float64(0)
	if hasLeader {
		exist = float64(1)
	}
	c.gauge(HAS_LEADER, exist, map[string]string{
		"leader_address": address,
		"leader_id":      id,
	})
}
