package command

import (
	"fmt"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"github.com/spf13/cobra"
)

func NewMicroServiceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service <subcommand>",
		Short: "service commands",
	}
	cmd.AddCommand(NewServiceRegisterCommand())
	cmd.AddCommand(NewServiceDiscoveryCommand())
	cmd.AddCommand(NewServiceDescCommand())
	cmd.AddCommand(NewServiceListCommand())
	cmd.AddCommand(NewServiceVersionsListCommand())
	return cmd
}

var (
	serviceName     string
	serviceVersion  string
	serviceEndpoint string
)

func NewServiceRegisterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register <namespace> <service-name> <service-version> <service-endpoint> <lease-id> [options]",
		Short: "register service",
		Run:   serviceRegisterCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for register service")
	cmd.Flags().StringVar(&serviceName, "service-name", "", "service-name for register service")
	cmd.Flags().StringVar(&serviceVersion, "service-version", "", "service-version for register service")
	cmd.Flags().StringVar(&serviceEndpoint, "service-endpoint", "", "service-endpoint for register service")
	cmd.Flags().Int64Var(&leaseId, "lease-id", 0, "lease-id for register service")
	return cmd
}

func serviceRegisterCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.Register(cmd.Context(), &pbmicroservice.ServiceInfo{
		Namespace:       namespace,
		ServiceName:     serviceName,
		ServiceVersion:  serviceVersion,
		ServiceEndpoint: serviceEndpoint,
	}, leaseId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s register success", serviceName)
}

func NewServiceDiscoveryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discovery <namespace> <service-name> <service-version> <service-endpoint> [options]",
		Short: "discovery service",
		Run:   serviceDiscoveryCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for discovery service")
	cmd.Flags().StringVar(&serviceName, "service-name", "", "service-name for discovery service")
	cmd.Flags().StringVar(&serviceVersion, "service-version", "", "service-version for discovery service")
	cmd.Flags().StringVar(&serviceEndpoint, "service-endpoint", "", "service-endpoint for discovery service")
	return cmd
}

func serviceDiscoveryCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.Discovery(cmd.Context(), &pbmicroservice.ServiceInfo{
		Namespace:       namespace,
		ServiceName:     serviceName,
		ServiceVersion:  serviceVersion,
		ServiceEndpoint: serviceEndpoint,
	})
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewServiceDescCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "desc <namespace> <service-name> <service-version> <service-endpoint> [options]",
		Short: "service desc",
		Run:   serviceDescCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for service desc")
	cmd.Flags().StringVar(&serviceName, "service-name", "", "service-name for service desc")
	cmd.Flags().StringVar(&serviceVersion, "service-version", "", "service-version for service desc")
	cmd.Flags().StringVar(&serviceEndpoint, "service-endpoint", "", "service-endpoint for service desc")
	return cmd
}

func serviceDescCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.ServiceDesc(cmd.Context(), &pbmicroservice.ServiceInfo{
		Namespace:       namespace,
		ServiceName:     serviceName,
		ServiceVersion:  serviceVersion,
		ServiceEndpoint: serviceEndpoint,
	})
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewServiceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <namespace> [options]",
		Short: "service list",
		Run:   serviceListCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for service list")
	return cmd
}

func serviceListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.ListServices(cmd.Context(), namespace)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewServiceVersionsListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "versionList <namespace> <service-name> [options]",
		Short: "service versions list",
		Run:   serviceVersionsListCommandFunc,
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "namespace for service version list")
	cmd.Flags().StringVar(&serviceName, "service-name", "", "service-name for service version list")
	return cmd
}

func serviceVersionsListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.ListServiceVersions(cmd.Context(), namespace, serviceName)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}
