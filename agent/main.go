package main

import (
	"context"
	"fmt"
	"os"

	"github.com/no-mole/venus/agent/venus"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile      string
	nodeID       string
	localAddr    string
	dataDir      string
	grpcEndpoint string
	httpEndpoint string
	joinAddr     string
	logLevel     string
	bootstrap    bool
	peerToken    string

	rootCmd = &cobra.Command{
		Use:   "venus",
		Short: "配置中心、注册中心,使用raft保证节点数据的一致性",
		Run: func(cmd *cobra.Command, args []string) {
			if localAddr == "" {
				localAddr = grpcEndpoint
			}
			ctx := context.Background()
			conf := config.GetDefaultConfig()
			conf.NodeID = nodeID
			conf.DaftDir = dataDir
			conf.GrpcEndpoint = grpcEndpoint
			conf.HttpEndpoint = httpEndpoint
			conf.BootstrapCluster = bootstrap
			conf.JoinAddr = joinAddr
			conf.LoggerLevel = config.LoggerLevel(logLevel)
			conf.PeerToken = peerToken
			conf.LocalAddr = localAddr
			s, err := venus.NewServer(ctx, conf)
			if err != nil {
				panic(err)
			}
			err = s.Start()
			if err != nil {
				println(err.Error())
			}
			println("server stopped!")
		},
	}
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "print server version",
		Run: func(cmd *cobra.Command, args []string) {
			println("v0.0.1")
		},
	})
}

// @title venus
// @version v1.0
// @description 统一对外输出的接口层,返回参数标准位json,结构为{"code":err code,"msg":"提示信息","data":object"}，文档中只展示data的结构

// @schemes https http
// @host 127.0.0.1:7233
// @BasePath /api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path,find in[/etc/venus/venus.yaml|$HOME/venus.yaml]")
	rootCmd.Flags().StringVar(&nodeID, "node-id", "node1", "node name")
	rootCmd.Flags().StringVar(&dataDir, "data-dir", "data", "data dir")
	rootCmd.Flags().StringVar(&localAddr, "local-addr", "", "local addr  for peer communication,default is 'grpc-endpoint'")
	rootCmd.Flags().StringVar(&grpcEndpoint, "grpc-endpoint", "127.0.0.1:6233", "grpc endpoint")
	rootCmd.Flags().StringVar(&httpEndpoint, "http-endpoint", "127.0.0.1:7233", "grpc endpoint")
	rootCmd.Flags().BoolVar(&bootstrap, "boot", false, "bootstrap cluster,only works on new cluster")
	rootCmd.Flags().StringVar(&joinAddr, "join", "", "join exist cluster addr")
	rootCmd.Flags().StringVar(&logLevel, "level", "info", "log level[debug|info|warn|err]")
	rootCmd.Flags().StringVar(&peerToken, "peer-token", "", "cluster peers certification token,string of length 8-16")

	_ = viper.BindPFlag("node_id", rootCmd.PersistentFlags().Lookup("node-id"))
	_ = viper.BindPFlag("data_dir", rootCmd.PersistentFlags().Lookup("data-dir"))
	_ = viper.BindPFlag("grpc_endpoint", rootCmd.PersistentFlags().Lookup("grpc-endpoint"))
	_ = viper.BindPFlag("http_endpoint", rootCmd.PersistentFlags().Lookup("http-endpoint"))
	_ = viper.BindPFlag("boot", rootCmd.PersistentFlags().Lookup("boot"))
	_ = viper.BindPFlag("join", rootCmd.PersistentFlags().Lookup("join"))
	_ = viper.BindPFlag("level", rootCmd.PersistentFlags().Lookup("level"))
	_ = viper.BindPFlag("peer_token", rootCmd.PersistentFlags().Lookup("peer-token"))

	err := rootCmd.Execute()
	if err != nil {
		println(err.Error())
	}
}

func initConfig() {
	viper.SetEnvPrefix("VENUS")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath("/etc/venus")
		viper.AddConfigPath(home)
		viper.SetConfigName("venus")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
