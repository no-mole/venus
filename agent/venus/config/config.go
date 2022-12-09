package config

type Config struct {
	NodeID           string `json:"node_id" yaml:"node_id"`
	RaftDir          string
	ServerAddr       string
	BootstrapCluster bool
}
