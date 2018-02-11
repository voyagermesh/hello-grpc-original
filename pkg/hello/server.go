package hello

import (
	"fmt"

	proto "github.com/appscode/hello-grpc/pkg/apis/hello/v1alpha1"
	"github.com/appscode/hello-grpc/pkg/server/endpoints"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(proto.RegisterHelloServer, &Server{})
	endpoints.ProxyServerEndpoints.Register(proto.RegisterHelloHandlerFromEndpoint)
	endpoints.ProxyServerCorsPattern.Register(proto.ExportHelloCorsPatterns())
}

type Server struct {
}

var _ proto.HelloServer = &Server{}

func (s *Server) Intro(ctx context.Context, req *proto.IntroRequest) (*proto.IntroResponse, error) {
	return &proto.IntroResponse{
		Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}
