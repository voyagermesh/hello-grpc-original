package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	hello "github.com/appscode/hello-grpc/pkg/apis/hello/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func doGRPC(address, crtPath string) error {
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
	result, err := client.Intro(context.Background(), &hello.IntroRequest{Name: "Voyager"})
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func main() {
	address := flag.String("address", "", "server address")
	crt := flag.String("crt", "", "path to cert file")
	flag.Parse()

	if err := doGRPC(*address, *crt); err != nil {
		log.Fatal(err)
	}
}
