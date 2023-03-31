package clientv1

import (
	"context"

	"github.com/no-mole/venus/proto/pbmicroservice"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type MicroService interface {
	Register(ctx context.Context, info *pbmicroservice.ServiceInfo, leaseId int64) error
	Discovery(ctx context.Context, info *pbmicroservice.ServiceInfo) (*pbmicroservice.DiscoveryServiceResponse, error)
	ServiceDesc(ctx context.Context, info *pbmicroservice.ServiceInfo) (*pbmicroservice.ServiceEndpointInfo, error)
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

func (m *microService) ServiceDesc(ctx context.Context, info *pbmicroservice.ServiceInfo) (*pbmicroservice.ServiceEndpointInfo, error) {
	return m.remote.ServiceDesc(ctx, info, m.callOpts...)
}

func (m *microService) Register(ctx context.Context, info *pbmicroservice.ServiceInfo, leaseId int64) error {
	m.logger.Debug("Register", zap.Any("info", info), zap.Int64("leaseId", leaseId))
	_, err := m.remote.Register(ctx, &pbmicroservice.RegisterServicesRequest{
		ServiceDesc: info,
		LeaseId:     leaseId,
	}, m.callOpts...)
	return err
}

func (m *microService) Discovery(ctx context.Context, info *pbmicroservice.ServiceInfo) (*pbmicroservice.DiscoveryServiceResponse, error) {
	return m.remote.Discovery(ctx, info, m.callOpts...)
}

func (m *microService) ListServices(ctx context.Context, namespace string) (*pbmicroservice.ListServicesResponse, error) {
	return m.remote.ListServices(ctx, &pbmicroservice.ListServicesRequest{Namespace: namespace}, m.callOpts...)

}

func (m *microService) ListServiceVersions(ctx context.Context, namespace, serviceName string) (*pbmicroservice.ListServiceVersionsResponse, error) {
	return m.remote.ListServiceVersions(ctx, &pbmicroservice.ListServiceVersionsRequest{Namespace: namespace, ServiceName: serviceName}, m.callOpts...)
}
