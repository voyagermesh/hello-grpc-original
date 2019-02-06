package cmds

import (
	"context"
	"fmt"
	"log"
	_ "net/http/pprof"
	"strings"

	hello "github.com/appscode/hello-grpc/pkg/apis/hello/v1alpha1"
	_ "github.com/appscode/hello-grpc/pkg/hello"
	_ "github.com/appscode/hello-grpc/pkg/status"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewCmdClient(stopCh <-chan struct{}) *cobra.Command {
	var address, certPath, name string

	cmd := &cobra.Command{
		Use:   "client",
		Short: "Launch Hello GRPC client",
		Long:  "Launch Hello GRPC client",
		RunE: func(c *cobra.Command, args []string) error {
			return doGRPC(address, certPath, name)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&address, "address", "", "address of the server")
	flags.StringVar(&certPath, "crt", "", "path to cert file")
	flags.StringVar(&name, "name", "appscode", "name field of hello-request")

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
