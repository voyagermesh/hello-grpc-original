package status

import (
	"fmt"

	proto "github.com/appscode/hello-grpc/pkg/apis/status"
	"github.com/appscode/hello-grpc/pkg/server/endpoints"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(proto.RegisterStatusServer, &Server{})
	endpoints.ProxyServerEndpoints.Register(proto.RegisterStatusHandlerFromEndpoint)
	endpoints.ProxyServerCorsPattern.Register(proto.ExportStatusCorsPatterns())
}

type Server struct {
}

var _ proto.StatusServer = &Server{}

func (s *Server) Intro(ctx context.Context, req *proto.IntroRequest) (*proto.IntroResponse, error) {
	return &proto.IntroResponse{
		Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}
