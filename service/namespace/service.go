package namespace

import (
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/proto"
)

type namespaceService struct {
	pbnamespace.UnimplementedNamespaceServer
}

var encoder = proto.MarshalOptions{}
var decoder = proto.UnmarshalOptions{}

var (
	bucketName = []byte("namespaces")
)

func New() pbnamespace.NamespaceServer {
	return &namespaceService{}
}
