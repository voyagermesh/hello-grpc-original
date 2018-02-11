package hello

import (
	"fmt"

	proto "github.com/appscode/hello-grpc/pkg/apis/hello/v1alpha1"
	"github.com/appscode/hello-grpc/pkg/server/endpoints"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(proto.RegisterHelloServiceServer, &Server{})
	endpoints.ProxyServerEndpoints.Register(proto.RegisterHelloServiceHandlerFromEndpoint)
	endpoints.ProxyServerCorsPattern.Register(proto.ExportHelloServiceCorsPatterns())
}

type Server struct {
}

var _ proto.HelloServiceServer = &Server{}

func (s *Server) Intro(ctx context.Context, req *proto.IntroRequest) (*proto.IntroResponse, error) {
	return &proto.IntroResponse{
		Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}
