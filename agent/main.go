package main

import (
	"context"
	"fmt"
	"github.com/no-mole/venus/agent/venus"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	// Used for flags.
	cfgFile      string
	nodeID       string
	dataDir      string
	grpcEndpoint string
	httpEndpoint string
	joinAddr     string
	logLevel     string
	bootstrap    bool

	rootCmd = &cobra.Command{
		Use:   "venus",
		Short: "配置中心、注册中心,使用raft保证节点数据的一致性",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			conf := config.GetDefaultConfig()
			conf.NodeID = nodeID
			conf.DaftDir = dataDir
			conf.GrpcEndpoint = grpcEndpoint
			conf.HttpEndpoint = httpEndpoint
			conf.BootstrapCluster = bootstrap
			conf.JoinAddr = joinAddr
			conf.LoggerLevel = config.LoggerLevel(logLevel)
			s, err := venus.NewServer(ctx, conf)
			if err != nil {
				println(err.Error())
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

func main() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path,find in[/etc/venus/venus.yaml|$HOME/venus.yaml]")
	rootCmd.Flags().StringVar(&nodeID, "node-id", "node1", "node name")
	rootCmd.Flags().StringVar(&dataDir, "data-dir", "data", "data dir")
	rootCmd.Flags().StringVar(&grpcEndpoint, "grpc-endpoint", "127.0.0.1:6233", "grpc endpoint")
	rootCmd.Flags().StringVar(&httpEndpoint, "http-endpoint", "127.0.0.1:7233", "grpc endpoint")
	rootCmd.Flags().BoolVar(&bootstrap, "boot", false, "bootstrap cluster,only works on new cluster")
	rootCmd.Flags().StringVar(&joinAddr, "join", "", "join exist cluster addr")
	rootCmd.Flags().StringVar(&logLevel, "level", "info", "log level[debug|info|warn|err]")

	_ = viper.BindPFlag("nodeID", rootCmd.PersistentFlags().Lookup("node-id"))
	_ = viper.BindPFlag("dataDir", rootCmd.PersistentFlags().Lookup("data-dir"))
	_ = viper.BindPFlag("grpcEndpoint", rootCmd.PersistentFlags().Lookup("grpc-endpoint"))
	_ = viper.BindPFlag("httpEndpoint", rootCmd.PersistentFlags().Lookup("http-endpoint"))
	_ = viper.BindPFlag("boot", rootCmd.PersistentFlags().Lookup("boot"))
	_ = viper.BindPFlag("join", rootCmd.PersistentFlags().Lookup("join"))
	_ = viper.BindPFlag("level", rootCmd.PersistentFlags().Lookup("level"))

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
