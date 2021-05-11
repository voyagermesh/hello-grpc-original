package hello

import (
	"fmt"
	"time"

	proto "voyagermesh.dev/hello-grpc/pkg/apis/hello/v1alpha1"
	"voyagermesh.dev/hello-grpc/pkg/cmds/server"
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

func (s *Server) Stream(req *proto.IntroRequest, stream proto.HelloService_StreamServer) error {
	for i := 0; i < 60; i++ {
		intro := fmt.Sprintf("%d: hello, %s!", i, req.Name)
		if err := stream.Send(&proto.IntroResponse{Intro: intro}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
