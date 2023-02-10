package venus

import (
	"github.com/hashicorp/raft"
	"google.golang.org/grpc/resolver"
)

const scheme = "venus-leader"

var _ resolver.Builder = (*grpcResolver)(nil)

type grpcResolver struct {
	r  *raft.Raft
	cc resolver.ClientConn
}

func (g *grpcResolver) Build(_ resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	g.cc = cc
	g.ResolveNow(resolver.ResolveNowOptions{})
	return g, nil
}

func (g *grpcResolver) Scheme() string {
	return scheme
}

func (g *grpcResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	leaderAddr, _ := g.r.LeaderWithID()
	if leaderAddr == "" {
		//print no leader
		return
	}
	err := g.cc.UpdateState(resolver.State{
		Addresses: []resolver.Address{{Addr: string(leaderAddr)}},
	})
	if err != nil {
		//print err
		return
	}
}

func (g *grpcResolver) Close() {}
