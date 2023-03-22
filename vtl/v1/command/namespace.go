package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewNamespaceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "namespace <subcommand>",
		Short: "namespace commands",
	}
	cmd.AddCommand(NewNamespaceAddCommand())
	cmd.AddCommand(NewNamespaceDelCommand())
	cmd.AddCommand(NewNamespaceListCommand())
	cmd.AddCommand(NewNamespaceAddUserCommand())
	cmd.AddCommand(NewNamespaceDelUserCommand())
	cmd.AddCommand(NewNamespaceUserListCommand())
	cmd.AddCommand(NewNamespaceAddAccessKeyCommand())
	cmd.AddCommand(NewNamespaceDelAccessKeyCommand())
	cmd.AddCommand(NewNamespaceAccessKeyListCommand())
	return cmd
}

var (
	namespaceAlias string
	namespaceUid   string
	role           string
)

func NewNamespaceAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <namespace-uid> [options]",
		Short: "namespace add",
		Run:   namespaceAddCommandFunc,
	}
	cmd.Flags().StringVar(&namespaceUid, "namespace-uid", "", "namespace-uid for namespace add")
	cmd.Flags().StringVar(&namespaceAlias, "namespace-alias", "", "namespace-alias for namespace add")
	return cmd
}

func namespaceAddCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.NamespaceAdd(cmd.Context(), namespaceAlias, namespaceUid)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewNamespaceDelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del <namespace> [options]",
		Short: "namespace del",
		Run:   namespaceDelCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for namespace del")
	return cmd
}

func namespaceDelCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.NamespaceDel(cmd.Context(), namespace)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has deleted", namespace)
}

func NewNamespaceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [options]",
		Short: "namespace list",
		Run:   namespaceListCommandFunc,
	}
	return cmd
}

func namespaceListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.NamespacesList(cmd.Context())
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewNamespaceAddUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "userAdd <namespace> <uid> [options]",
		Short: "namespace add user",
		Run:   namespaceAddUserCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for namespace add user")
	cmd.Flags().StringVar(&uid, "uid", "", "uid for namespace add user")
	return cmd
}

func namespaceAddUserCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.NamespaceAddUser(cmd.Context(), namespace, uid, role)
	if err != nil {
		panic(err)
	}
	fmt.Printf("namespace:%s,uid:%s add success", namespace, uid)
}

func NewNamespaceDelUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "userDel <namespace> <uid> [options]",
		Short: "namespace del user",
		Run:   namespaceDelUserCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for namespace del user")
	cmd.Flags().StringVar(&uid, "uid", "", "uid for namespace del user")
	return cmd
}

func namespaceDelUserCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.NamespaceDelUser(cmd.Context(), namespace, uid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("namespace:%s,uid:%s has deleted", namespace, uid)
}

func NewNamespaceUserListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "userList <namespace> [options]",
		Short: "namespace user list",
		Run:   namespaceUserListCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for namespace user list")
	return cmd
}

func namespaceUserListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.NamespaceUserList(cmd.Context(), namespace)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewNamespaceAddAccessKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "akAdd <namespace> <ak> [options]",
		Short: "namespace add accessKey",
		Run:   namespaceAddAccessKeyCommandFunc,
	}
	cmd.Flags().StringVar(&ak, "ak", "", "ak for namespace accessKey add")
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for namespace accessKey add")
	return cmd
}

func namespaceAddAccessKeyCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.NamespaceAddAccessKey(cmd.Context(), ak, namespace)
	if err != nil {
		panic(err)
	}
	fmt.Printf("namespace:%s,ak:%s add success", namespace, ak)
}

func NewNamespaceDelAccessKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "akDel <namespace> <ak> [options]",
		Short: "namespace del accessKey",
		Run:   namespaceDelAccessKeyCommandFunc,
	}
	cmd.Flags().StringVar(&ak, "ak", "", "ak for namespace accessKey del")
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for namespace accessKey del")
	return cmd
}

func namespaceDelAccessKeyCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.NamespaceDelAccessKey(cmd.Context(), ak, namespace)
	if err != nil {
		panic(err)
	}
	fmt.Printf("namespace:%s,ak:%s has deleted", namespace, ak)
}

func NewNamespaceAccessKeyListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "akList <namespace> [options]",
		Short: "namespace accessKey list",
		Run:   namespaceAccessKeyListCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for namespace accessKey list")
	return cmd
}

func namespaceAccessKeyListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.NamespaceAccessKeyList(cmd.Context(), namespace)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}
