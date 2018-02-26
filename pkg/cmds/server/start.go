package server

import (
	"net/http"
	"strings"

	stringz "github.com/appscode/go/strings"
	utilerrors "github.com/appscode/go/util/errors"
	grpc_cors "github.com/appscode/grpc-go-addons/cors"
	"github.com/appscode/grpc-go-addons/endpoints"
	grpc_security "github.com/appscode/grpc-go-addons/security"
	"github.com/appscode/grpc-go-addons/server"
	"github.com/appscode/grpc-go-addons/server/options"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	GRPCEndpoints    = endpoints.GRPCRegistry{}
	GatewayEndpoints = endpoints.ProxyRegistry{}
	CorsPatterns     = grpc_cors.PatternRegistry{}
)

type ServerOptions struct {
	RecommendedOptions *options.RecommendedOptions
	LogRPC             bool
}

func NewServerOptions() *ServerOptions {
	o := &ServerOptions{
		RecommendedOptions: options.NewRecommendedOptions(),
	}
	return o
}

func (o ServerOptions) Validate(args []string) error {
	var errors []error
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

func (o *ServerOptions) Complete() error {
	return nil
}

func (o ServerOptions) Config() (*server.Config, error) {
	config := server.NewConfig()
	if err := o.RecommendedOptions.ApplyTo(config); err != nil {
		return nil, err
	}

	config.SetGRPCRegistry(GRPCEndpoints)
	config.SetProxyRegistry(GatewayEndpoints)
	config.SetCORSRegistry(CorsPatterns)

	optsGLog := []grpc_glog.Option{
		grpc_glog.WithDecider(func(methodFullName string, err error) bool {
			// will not log gRPC calls if it was a call to healthcheck and no error was raised
			if err == nil && methodFullName == "/github.com.appscode.hellogrpc.apis.status.StatusService/Status" {
				return false
			}

			// by default you will log all calls
			return o.LogRPC
		}),
	}
	payloadDecider := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
		// will not log gRPC calls if it was a call to healthcheck and no error was raised
		if fullMethodName == "/github.com.appscode.hellogrpc.apis.status.StatusService/Status" {
			return false
		}

		// by default you will log all calls
		return o.LogRPC
	}
	glogEntry := ctx_glog.NewEntry(ctx_glog.Logger)
	grpc_glog.ReplaceGrpcLogger()

	config.GRPCServerOption(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_glog.PayloadStreamServerInterceptor(glogEntry, payloadDecider),
			grpc_glog.StreamServerInterceptor(glogEntry, optsGLog...),
			grpc_cors.StreamServerInterceptor(grpc_cors.OriginHost(config.CORSOriginHost), grpc_cors.AllowSubdomain(config.CORSAllowSubdomain)),
			grpc_security.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_glog.PayloadUnaryServerInterceptor(glogEntry, payloadDecider),
			grpc_glog.UnaryServerInterceptor(glogEntry, optsGLog...),
			grpc_cors.UnaryServerInterceptor(grpc_cors.OriginHost(config.CORSOriginHost), grpc_cors.AllowSubdomain(config.CORSAllowSubdomain)),
			grpc_security.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	config.GatewayMuxOption(gwrt.WithIncomingHeaderMatcher(func(h string) (string, bool) {
		if stringz.PrefixFold(h, "access-control-request-") ||
			stringz.PrefixFold(h, "k8s-") ||
			strings.EqualFold(h, "Origin") ||
			strings.EqualFold(h, "Cookie") ||
			strings.EqualFold(h, "X-Phabricator-Csrf") {
			return h, true
		}
		return "", false
	}),
		gwrt.WithOutgoingHeaderMatcher(func(h string) (string, bool) {
			if stringz.PrefixFold(h, "access-control-allow-") ||
				strings.EqualFold(h, "Set-Cookie") ||
				strings.EqualFold(h, "vary") ||
				strings.EqualFold(h, "x-content-type-options") ||
				stringz.PrefixFold(h, "x-ratelimit-") {
				return h, true
			}
			return "", false
		}),
		gwrt.WithMetadata(func(c context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("x-forwarded-method", req.Method)
		}),
		gwrt.WithProtoErrorHandler(gwrt.DefaultHTTPProtoErrorHandler))

	return config, nil
}

func (o ServerOptions) RunServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.New()
	if err != nil {
		return err
	}

	return server.Run(stopCh)
}
