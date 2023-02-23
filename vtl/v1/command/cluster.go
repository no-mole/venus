package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewClusterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster <subcommand>",
		Short: "Cluster membership related commands",
	}
	cmd.AddCommand(NewClusterMemberAddCommand())
	cmd.AddCommand(NewClusterMemberRemoveCommand())
	cmd.AddCommand(NewClusterMemberListCommand())
	return cmd
}

var (
	peerAddr string
	peerId   string
	peerRole string
)

func NewClusterMemberAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <memberAddr> [options]",
		Short: "Add a member into the cluster",
		Run:   memberAddCommandFunc,
	}
	cmd.Flags().StringVar(&peerId, "peer-id", "", "peer id for add")
	cmd.Flags().StringVar(&peerAddr, "peer-addr", "", "peer address for add")
	cmd.Flags().StringVar(&peerAddr, "peer-role", "voter", "peer role for add[voter|nonvoter],default is voter")
	return cmd
}

func memberAddCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	switch peerRole {
	case "voter":
		err = client.AddVoter(cmd.Context(), peerId, peerAddr, 0)
	case "nonvoter":
		err = client.AddNonvoter(cmd.Context(), peerId, peerAddr, 0)
	default:
		panic("not supported peer role")
	}
}

func NewClusterMemberRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove  [options]",
		Short: "Remove a member in the cluster",
		Run:   memberRemoveCommandFunc,
	}
	cmd.Flags().StringVar(&peerId, "peer-id", "", "peer id for add")
	return cmd
}

func memberRemoveCommandFunc(cmd *cobra.Command, args []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.RemoveServer(cmd.Context(), peerId, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("server %s removed", peerId)
}

func NewClusterMemberListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list members in the cluster",
		Run:   memberListCommandFunc,
	}
	return cmd
}

func memberListCommandFunc(cmd *cobra.Command, args []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.Nodes(cmd.Context())
	if err != nil {
		panic(err)
	}
	println(resp.String())
}
