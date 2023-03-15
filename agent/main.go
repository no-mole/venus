package main

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strings"

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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
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

			fmt.Printf("%+v", conf)

			s, err := venus.NewServer(ctx, conf)
			if err != nil {
				panic(err)
			}
			err = s.Wait()
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
			println("v0.0.7")
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path,find in[/etc/venus/venus.yaml|$HOME/venus.yaml]")
	rootCmd.PersistentFlags().StringVar(&nodeID, "node-id", "node1", "node name")
	rootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", "data", "data dir")
	rootCmd.PersistentFlags().StringVar(&localAddr, "local-addr", "", "local addr  for peer communication,default is 'grpc-endpoint'")
	rootCmd.PersistentFlags().StringVar(&grpcEndpoint, "grpc-endpoint", "127.0.0.1:6233", "grpc endpoint")
	rootCmd.PersistentFlags().StringVar(&httpEndpoint, "http-endpoint", "127.0.0.1:7233", "grpc endpoint")
	rootCmd.PersistentFlags().BoolVar(&bootstrap, "boot", false, "bootstrap cluster,only works on new cluster")
	rootCmd.PersistentFlags().StringVar(&joinAddr, "join", "", "join exist cluster addr")
	rootCmd.PersistentFlags().StringVar(&logLevel, "level", "info", "log level[debug|info|warn|err]")
	rootCmd.PersistentFlags().StringVar(&peerToken, "peer-token", "", "cluster peers certification token,string of length 8-16")
	err := rootCmd.Execute()
	if err != nil {
		println(err.Error())
	}
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	if cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	} else {
		v.AddConfigPath("/etc/venus")

		// Find home directory.
		home, _ := os.UserHomeDir()
		if home != "" {
			v.AddConfigPath(home)
		}

		v.SetConfigName("venus")
		v.SetConfigType("yaml")
	}

	if err := v.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	v.SetEnvPrefix("VENUS")
	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				panic(err)
			}
		}
	})
}
