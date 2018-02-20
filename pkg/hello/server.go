package hello

import (
	"fmt"

	proto "github.com/appscode/hello-grpc/pkg/apis/hello/v1alpha1"
	"github.com/appscode/hello-grpc/pkg/cmds/server"
	"golang.org/x/net/context"
)

func init() {
	server.GRPCEndpoints.Register(proto.RegisterHelloServiceServer, &Server{})
	server.GatewayEndpoints.Register(proto.RegisterHelloServiceHandlerFromEndpoint)
	server.CorsPatterns.Register(proto.ExportHelloServiceCorsPatterns())
}

type Server struct {
}

var _ proto.HelloServiceServer = &Server{}

func (s *Server) Intro(ctx context.Context, req *proto.IntroRequest) (*proto.IntroResponse, error) {
	return &proto.IntroResponse{
		Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}
