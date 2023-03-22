package command

import (
	"fmt"
	"github.com/no-mole/venus/proto/pbaccesskey"
	"github.com/spf13/cobra"
)

func NewAccessKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accessKey <subcommand>",
		Short: "accessKey commands",
	}
	cmd.AddCommand(NewAccessKeyGenCommand())
	cmd.AddCommand(NewAccessKeyDelCommand())
	cmd.AddCommand(NewAccessKeyChangeStatusCommand())
	cmd.AddCommand(NewAccessKeyLoginCommand())
	cmd.AddCommand(NewAccessKeyListCommand())
	cmd.AddCommand(NewAccessKeyNamespaceListCommand())

	return cmd
}

var (
	alias  string
	ak     string
	status int32
	secret string
)

func NewAccessKeyGenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen <alias> [options]",
		Short: "gen a access key ",
		Run:   accessKeyGenCommandFunc,
	}
	cmd.Flags().StringVar(&alias, "alias", "", "alias for access key gen")
	return cmd
}

func accessKeyGenCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	info, err := client.AccessKeyGen(cmd.Context(), alias)
	if err != nil {
		panic(err)
	}
	println(info.String())
}

func NewAccessKeyDelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del <ak> [options]",
		Short: "delete a access key",
		Run:   accessKeyDelCommandFunc,
	}
	cmd.Flags().StringVar(&ak, "ak", "", "access key for delete")
	return cmd
}

func accessKeyDelCommandFunc(cmd *cobra.Command, _ []string) {
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.AccessKeyDel(cmd.Context(), ak)
	if err != nil {
		panic(err)
	}
	fmt.Printf("access key %s deleted", ak)
}

func NewAccessKeyChangeStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change <ak> <status> [options]",
		Short: "change status of access key",
		Run:   accessKeyChangeStatusCommandFunc,
	}
	cmd.Flags().StringVar(&ak, "ak", "", "access key for change status")
	cmd.Flags().Int32Var(&status, "status", 0, "status for change status")
	return cmd
}

func accessKeyChangeStatusCommandFunc(cmd *cobra.Command, _ []string) {
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.AccessKeyChangeStatus(cmd.Context(), ak, pbaccesskey.AccessKeyStatus(status))
	if err != nil {
		panic(err)
	}
	fmt.Printf("access key %s status changed, status is %d", ak, status)
}

func NewAccessKeyLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login <ak> <secret> [options]",
		Short: "access key login",
		Run:   accessKeyLoginCommandFunc,
	}
	cmd.Flags().StringVar(&ak, "ak", "", "access key login")
	cmd.Flags().StringVar(&secret, "secret", "", "access key login")
	return cmd
}

func accessKeyLoginCommandFunc(cmd *cobra.Command, _ []string) {
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	info, err := client.AccessKeyLogin(cmd.Context(), ak, secret)
	if err != nil {
		panic(err)
	}
	println(info.String())
}

func NewAccessKeyListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [options]",
		Short: "access key list",
		Run:   accessKeyListCommandFunc,
	}
	return cmd
}

func accessKeyListCommandFunc(cmd *cobra.Command, _ []string) {
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.AccessKeyList(cmd.Context())
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewAccessKeyNamespaceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "namespace list [options]",
		Short: "access key namespace list",
		Run:   accessKeyNamespaceListCommandFunc,
	}
	cmd.Flags().StringVar(&ak, "ak", "", "access key for namespace list")
	return cmd
}

func accessKeyNamespaceListCommandFunc(cmd *cobra.Command, _ []string) {
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.AccessKeyNamespaceList(cmd.Context(), ak)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}
