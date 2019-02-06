package hello

import (
	"fmt"
	"io"
	"log"
	"time"

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

func (s *Server) Stream(stream proto.HelloService_StreamServer) error {
	requestChan := make(chan string)
	errChan := make(chan error)

	go receive(stream, requestChan, errChan)

	for {
		var intro string

		select {
		case err := <-errChan:
			return err
		case name := <-requestChan:
			intro = fmt.Sprintf("hello, %s!", name)
		default:
			intro = "hello, client!"
		}

		if err := stream.Send(&proto.IntroResponse{Intro: intro}); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func receive(stream proto.HelloService_StreamServer, requestChan chan string, errChan chan error) {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("Received EOF")
			break
		}
		if err != nil {
			errChan <- err
		}
		log.Println("Received", in.Name)
		requestChan <- in.Name
	}
}
