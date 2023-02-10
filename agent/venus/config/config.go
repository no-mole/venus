package config

import "time"

type Config struct {
	NodeID           string `json:"node_id" yaml:"node_id"`
	RaftDir          string
	GrpcEndpoint     string
	HttpEndpoint     string
	BootstrapCluster bool
	ApplyTimeout     time.Duration
	JoinAddr         string
}
