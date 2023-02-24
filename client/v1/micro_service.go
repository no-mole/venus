package clientv1

import (
	"context"

	"github.com/no-mole/venus/proto/pbmicroservice"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type MicroService interface {
	Register(ctx context.Context, services []*pbmicroservice.ServiceInfo, leaseId int64) error
	Discovery(ctx context.Context, info *pbmicroservice.ServiceInfo, fn func(eps []string) error) error
	DiscoveryOnce(ctx context.Context, info *pbmicroservice.ServiceInfo) (*pbmicroservice.DiscoveryServiceResponse, error)
	ListServices(ctx context.Context, namespace string) (*pbmicroservice.ListServicesResponse, error)
	ListServiceVersions(ctx context.Context, namespace, serviceName string) (*pbmicroservice.ListServiceVersionsResponse, error)
}

func NewMicroService(c *Client, logger *zap.Logger) MicroService {
	return &microService{
		remote:   pbmicroservice.NewMicroServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("micro_service"),
	}
}

var _ MicroService = &microService{}

type microService struct {
	remote   pbmicroservice.MicroServiceClient
	callOpts []grpc.CallOption
	logger   *zap.Logger
}

func (m *microService) Register(ctx context.Context, services []*pbmicroservice.ServiceInfo, leaseId int64) error {
	m.logger.Debug("Register", zap.Any("info", services), zap.Int64("leaseId", leaseId))
	_, err := m.remote.Register(ctx, &pbmicroservice.RegisterServicesRequest{
		Services: services,
		LeaseId:  leaseId,
	}, m.callOpts...)
	return err
}

func (m *microService) Discovery(ctx context.Context, info *pbmicroservice.ServiceInfo, fn func(eps []string) error) error {
	client, err := m.remote.Discovery(ctx, info, m.callOpts...)
	if err != nil {
		return err
	}
	for {
		resp, err := client.Recv()
		if err != nil {
			return err
		}
		err = fn(resp.Endpoints)
		if err != nil {
			return err
		}
	}
}

func (m *microService) DiscoveryOnce(ctx context.Context, info *pbmicroservice.ServiceInfo) (*pbmicroservice.DiscoveryServiceResponse, error) {
	return m.remote.DiscoveryOnce(ctx, info, m.callOpts...)
}

func (m *microService) ListServices(ctx context.Context, namespace string) (*pbmicroservice.ListServicesResponse, error) {
	return m.remote.ListServices(ctx, &pbmicroservice.ListServicesRequest{Namespace: namespace}, m.callOpts...)

}

func (m *microService) ListServiceVersions(ctx context.Context, namespace, serviceName string) (*pbmicroservice.ListServiceVersionsResponse, error) {
	return m.remote.ListServiceVersions(ctx, &pbmicroservice.ListServiceVersionsRequest{Namespace: namespace, ServiceName: serviceName}, m.callOpts...)
}
