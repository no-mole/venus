package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net"
	"os"
	"strings"
)

var ClientIPKey = "client-ip"
var ClientHostnameKey = "client-hostname"

func WithMetadata(md metadata.MD) {
	md.Set(ClientIPKey, ip)
	md.Set(ClientHostnameKey, hostname)
}

func FromContext(ctx context.Context) (hostname string, ip string) {
	if h, ok := ctx.Value(ClientHostnameKey).(string); ok {
		hostname = h
	}
	if i, ok := ctx.Value(ClientIPKey).(string); ok {
		ip = i
	}
	return
}

func WithContext(ctx context.Context, md metadata.MD) context.Context {
	ips := md.Get(ClientIPKey)
	if len(ips) > 0 {
		ctx = context.WithValue(ctx, ClientIPKey, ips[0])
	}
	hostnames := md.Get(ClientHostnameKey)
	if len(hostnames) > 0 {
		ctx = context.WithValue(ctx, ClientHostnameKey, hostnames[0])
	}
	return ctx
}

var ip string
var hostname string

func init() {
	ip, _ = GetOutBoundIP()
	hostname, _ = os.Hostname()
}
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	defer conn.Close()
	if err != nil {
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
