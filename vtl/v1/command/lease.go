package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewLeaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease <subcommand>",
		Short: "lease commands",
	}
	cmd.AddCommand(NewLeaseGrantCommand())
	cmd.AddCommand(NewLeaseTimeToLiveCommand())
	cmd.AddCommand(NewLeaseRevokeCommand())
	cmd.AddCommand(NewLeasesListCommand())
	cmd.AddCommand(NewLeaseKeepaliveOnceCommand())

	return cmd
}

var (
	ttl     int64
	leaseId int64
)

func NewLeaseGrantCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant <ttl> [options]",
		Short: "grant a lease ",
		Run:   leaseGrantCommandFunc,
	}
	cmd.Flags().Int64Var(&ttl, "ttl", 0, "ttl for grant of lease")
	return cmd
}

func leaseGrantCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	info, err := client.Grant(cmd.Context(), ttl)
	if err != nil {
		panic(err)
	}
	println(info.String())
}

func NewLeaseTimeToLiveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "timeToLive <lease-id> [options]",
		Short: "timeToLive of the lease",
		Run:   leaseTimeToLiveCommandFunc,
	}
	cmd.Flags().Int64Var(&leaseId, "lease-id", 0, "leaseId for the lease")
	return cmd
}

func leaseTimeToLiveCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.TimeToLive(cmd.Context(), leaseId)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewLeaseRevokeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke <lease-id> [options]",
		Short: "revoke the lease",
		Run:   leaseRevokeCommandFunc,
	}
	cmd.Flags().Int64Var(&leaseId, "lease-id", 0, "leaseId for revoke lease")
	return cmd
}

func leaseRevokeCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.Revoke(cmd.Context(), leaseId)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewLeasesListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [options]",
		Short: "lease list",
		Run:   leasesListCommandFunc,
	}
	return cmd
}

func leasesListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.Leases(cmd.Context())
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewLeaseKeepaliveOnceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keepalive <lease-id> [options]",
		Short: "keepalive to the lease",
		Run:   leaseKeepaliveOnceCommandFunc,
	}
	cmd.Flags().Int64Var(&leaseId, "lease-id", 0, "leaseId for keep alive")
	return cmd
}

func leaseKeepaliveOnceCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.KeepaliveOnce(cmd.Context(), leaseId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d has keeped alive", leaseId)
}
