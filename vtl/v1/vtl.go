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
	rootCmd.PersistentFlags().String("endpoint", "", "node endpoint,separate with commands,default is 127.0.0.1:3333;example:[127.0.0.1:3333 or 127.0.0.1:3333,127.0.0.1:3334]")
	viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))
	rootCmd.PersistentFlags().String("username", "", "username")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	rootCmd.PersistentFlags().String("root-password", "", "password,must with username")
	viper.BindPFlag("root-password", rootCmd.PersistentFlags().Lookup("root-password"))
	rootCmd.PersistentFlags().String("peer-token", "", "peer-token")
	viper.BindPFlag("peer-token", rootCmd.PersistentFlags().Lookup("peer-token"))
	rootCmd.PersistentFlags().String("access-key", "", "access-key")
	viper.BindPFlag("access-key", rootCmd.PersistentFlags().Lookup("access-key"))
	rootCmd.PersistentFlags().String("access-key-secret", "", "access-key-secret")
	viper.BindPFlag("access-key-secret", rootCmd.PersistentFlags().Lookup("access-key-secret"))

	rootCmd.PersistentFlags().Duration("dial-timeout", time.Second, "dail server timeout,default is 1s")
	rootCmd.PersistentFlags().Duration("dial-keepalive-timeout", 10*time.Second, "dail server keepalive timeout,default is 10s")

	rootCmd.PersistentFlags().Int("max-call-send-msg-size", 0, "max call send msg size")
	viper.BindPFlag("max-call-send-msg-size", rootCmd.PersistentFlags().Lookup("max-call-send-msg-size"))
	rootCmd.PersistentFlags().Int("max-call-recv-msg-size", 0, "max call recv msg size")
	viper.BindPFlag("max-call-recv-msg-size", rootCmd.PersistentFlags().Lookup("max-call-recv-msg-size"))

	rootCmd.AddCommand(command.NewClusterCommand())
	rootCmd.AddCommand(command.NewAccessKeyCommand())
	rootCmd.AddCommand(command.NewKvCommand())
	rootCmd.AddCommand(command.NewLeaseCommand())
	rootCmd.AddCommand(command.NewUserCommand())
	rootCmd.AddCommand(command.NewMicroServiceCommand())
	rootCmd.AddCommand(command.NewNamespaceCommand())
}
