package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func NewKvCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kv <subcommand>",
		Short: "kv commands",
	}
	cmd.AddCommand(NewKvAddCommand())
	cmd.AddCommand(NewKvFetchCommand())
	cmd.AddCommand(NewKvDelCommand())
	cmd.AddCommand(NewKvListCommand())
	cmd.AddCommand(NewKvWatchCommand())

	return cmd
}

var (
	namespace string
	key       string
	dataType  string
	value     string
)

func NewKvAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <namespace> <key> <alias> <data-type> <value> [options]",
		Short: "add a kv",
		Run:   kvAddCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for add kv in this namespace")
	cmd.Flags().StringVar(&key, "key", "", "key for add kv in this namespace")
	cmd.Flags().StringVar(&alias, "alias", "", "alias for add kv in this namespace")
	cmd.Flags().StringVar(&dataType, "data-type", "", "data-type for add kv in this namespace")
	cmd.Flags().StringVar(&value, "value", "", "value for add kv in this namespace")
	return cmd
}

func kvAddCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	info, err := client.AddKV(cmd.Context(), namespace, key, dataType, value, alias)
	if err != nil {
		panic(err)
	}
	println(info.String())
}

func NewKvFetchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch <namespace> <key> [options]",
		Short: "fetch a kv in this namespace",
		Run:   kvFetchCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for fetch kv in this namespace")
	cmd.Flags().StringVar(&key, "key", "", "key for fetch kv in this namespace")
	return cmd
}

func kvFetchCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	info, err := client.FetchKey(cmd.Context(), namespace, key)
	if err != nil {
		panic(err)
	}
	println(info.String())
}

func NewKvDelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del <namespace> <key> [options]",
		Short: "del a kv in this namespace",
		Run:   kvDelCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for delete kv in this namespace")
	cmd.Flags().StringVar(&key, "key", "", "key for delete kv in this namespace")
	return cmd
}

func kvDelCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.DelKey(cmd.Context(), namespace, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("this key %s has deleted", key)
}

func NewKvListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <namespace> [options]",
		Short: "kv list in this namespace",
		Run:   kvListCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "kv list in this namespace")
	return cmd
}

func kvListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.ListKeys(cmd.Context(), namespace)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewKvWatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch <namespace> <key> [options]",
		Short: "watch kv in this namespace",
		Run:   kvWatchCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "watch kv in this namespace")
	cmd.Flags().StringVar(&key, "key", "", "watch kv in this namespace")
	return cmd
}

func kvWatchCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.WatchKey(cmd.Context(), namespace, key, func(namespace, key string) error {
		fmt.Printf("%s, namespace: %s, key: %s\n", time.Now().Format(time.RFC3339), namespace, key)
		item, err := client.FetchKey(cmd.Context(), namespace, key)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", item)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
