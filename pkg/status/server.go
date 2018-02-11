package status

import (
	v "github.com/appscode/go/version"
	proto "github.com/appscode/hello-grpc/pkg/apis/status"
	"github.com/appscode/hello-grpc/pkg/server/endpoints"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(proto.RegisterStatusServiceServer, &Server{})
	endpoints.ProxyServerEndpoints.Register(proto.RegisterStatusServiceHandlerFromEndpoint)
	endpoints.ProxyServerCorsPattern.Register(proto.ExportStatusServiceCorsPatterns())
}

type Server struct {
}

var _ proto.StatusServiceServer = &Server{}

func (s *Server) Status(ctx context.Context, req *proto.StatusRequest) (*proto.StatusResponse, error) {
	return &proto.StatusResponse{
		Version: &proto.Version{
			Version:         v.Version.Version,
			VersionStrategy: v.Version.VersionStrategy,
			Os:              v.Version.Os,
			Arch:            v.Version.Arch,
			CommitHash:      v.Version.CommitHash,
			GitBranch:       v.Version.GitBranch,
			GitTag:          v.Version.GitTag,
			CommitTimestamp: v.Version.CommitTimestamp,
			BuildTimestamp:  v.Version.BuildTimestamp,
			BuildHost:       v.Version.BuildHost,
			BuildHostOs:     v.Version.BuildHostOs,
			BuildHostArch:   v.Version.BuildHostArch,
		},
	}, nil
}
