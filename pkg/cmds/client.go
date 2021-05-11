package cmds

import (
	"context"
	"fmt"
	"io"
	"log"
	_ "net/http/pprof"
	"strings"

	hello "voyagermesh.dev/hello-grpc/pkg/apis/hello/v1alpha1"
	_ "voyagermesh.dev/hello-grpc/pkg/hello"
	_ "voyagermesh.dev/hello-grpc/pkg/status"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewCmdClient(stopCh <-chan struct{}) *cobra.Command {
	var address, certPath, name string
	var stream bool

	cmd := &cobra.Command{
		Use:   "client",
		Short: "Launch Hello GRPC client",
		Long:  "Launch Hello GRPC client",
		RunE: func(c *cobra.Command, args []string) error {
			if stream {
				return doGRPCStream(address, certPath, name)
			}
			return doGRPC(address, certPath, name)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&address, "address", "", "address of the server")
	flags.StringVar(&certPath, "crt", "", "path to cert file")
	flags.StringVar(&name, "name", "appscode", "name field of hello-request")
	flags.BoolVar(&stream, "stream", false, "use stream API")

	return cmd
}

func doGRPC(address, crtPath, name string) error {
	address = strings.TrimPrefix(address, "http://")
	address = strings.TrimPrefix(address, "https://")

	log.Println(address)

	option := grpc.WithInsecure()
	if len(crtPath) > 0 {
		creds, err := credentials.NewClientTLSFromFile(crtPath, "")
		if err != nil {
			return fmt.Errorf("failed to load TLS certificate")
		}
		option = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial(address, option)
	if err != nil {
		return fmt.Errorf("did not connect, %v", err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)
	result, err := client.Intro(context.Background(), &hello.IntroRequest{Name: name})
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func doGRPCStream(address, crtPath, name string) error {
	address = strings.TrimPrefix(address, "http://")
	address = strings.TrimPrefix(address, "https://")

	log.Println(address)

	option := grpc.WithInsecure()
	if len(crtPath) > 0 {
		creds, err := credentials.NewClientTLSFromFile(crtPath, "")
		if err != nil {
			return fmt.Errorf("failed to load TLS certificate")
		}
		option = grpc.WithTransportCredentials(creds)

	}

	conn, err := grpc.Dial(address, option)
	if err != nil {
		return fmt.Errorf("did not connect, %v", err)
	}
	defer conn.Close()

	streamClient, err := hello.NewHelloServiceClient(conn).Stream(context.Background(), &hello.IntroRequest{Name: name})
	if err != nil {
		return err
	}

	for {
		result, err := streamClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println(result)
	}

	return nil
}
