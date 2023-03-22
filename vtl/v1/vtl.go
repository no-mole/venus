/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package v1

import (
	"context"
	"github.com/no-mole/venus/vtl/v1/command"
	"github.com/spf13/viper"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:        "vtl",
	Short:      "A command line tools client for [venus]",
	SuggestFor: []string{"vtl"},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	rootCmd.SetContext(ctx)
	rootCmd.PersistentFlags().String(command.Endpoint, "", "node endpoint,separate with commands,default is 127.0.0.1:3333;example:[127.0.0.1:3333 or 127.0.0.1:3333,127.0.0.1:3334]")
	rootCmd.PersistentFlags().String(command.UserName, "", "username")
	rootCmd.PersistentFlags().String(command.RootPassword, "", "password,must with username")
	rootCmd.PersistentFlags().String(command.PeerToken, "", "peer-token")
	rootCmd.PersistentFlags().String(command.AccessKey, "", "access-key")
	rootCmd.PersistentFlags().String(command.AccessKeySecret, "", "access-key-secret")

	rootCmd.PersistentFlags().Duration(command.DialTimeout, time.Second, "dail server timeout,default is 1s")
	rootCmd.PersistentFlags().Duration(command.DialKeepaliveTimeout, time.Second, "dail server keepalive timeout,default is 1s")
	rootCmd.PersistentFlags().Duration(command.DialKeepaliveTime, 10*time.Second, "dail server keepalive time,default is 10s")

	rootCmd.PersistentFlags().Int(command.MaxCallSendMsgSize, 0, "max call send msg size")
	rootCmd.PersistentFlags().Int(command.MaxCallRecvMsgSize, 0, "max call recv msg size")

	viper.BindPFlags(rootCmd.PersistentFlags())
	rootCmd.AddCommand(command.NewClusterCommand())
	rootCmd.AddCommand(command.NewAccessKeyCommand())
	rootCmd.AddCommand(command.NewKvCommand())
	rootCmd.AddCommand(command.NewLeaseCommand())
	rootCmd.AddCommand(command.NewUserCommand())
	rootCmd.AddCommand(command.NewMicroServiceCommand())
	rootCmd.AddCommand(command.NewNamespaceCommand())
}
