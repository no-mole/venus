/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package v1

import (
	"context"
	"github.com/no-mole/venus/vtl/v1/command"
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
	rootCmd.PersistentFlags().String("username", "", "username")
	rootCmd.PersistentFlags().String("password", "", "password,must with username")

	rootCmd.PersistentFlags().Duration("dial-timeout", time.Second, "dail server timeout,default is 1s")
	rootCmd.PersistentFlags().Duration("dial-keepalive-timeout", 10*time.Second, "dail server keepalive timeout,default is 10s")

	rootCmd.PersistentFlags().Int("max-call-send-msg-size", 0, "max call send msg size")
	rootCmd.PersistentFlags().Int("max-call-recv-msg-size", 0, "max call recv msg size")

	rootCmd.AddCommand(command.NewClusterCommand())
}
