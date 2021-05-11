module voyagermesh.dev/hello-grpc

go 1.16

require (
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/opentracing/opentracing-go v1.0.3-0.20180212003421-3999fca714c8 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/net v0.0.0-20191002035440-2ec189313ef0
	gomodules.xyz/errors v0.0.0-20201104190405-077f059979fd
	gomodules.xyz/grpc-go-addons v0.2.2-0.20210218145105-321b2e13985f
	gomodules.xyz/kglog v0.0.1
	gomodules.xyz/runtime v0.2.0
	gomodules.xyz/signals v0.0.0-20201104192641-f8f5c878d966
	gomodules.xyz/x v0.0.0-20201105065653-91c568df6331
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a
	google.golang.org/grpc v1.24.0
)

replace (
	github.com/grpc-ecosystem/go-grpc-middleware => github.com/tamalsaha/go-grpc-middleware v0.0.0-20180226223443-606e44dc6300
	github.com/grpc-ecosystem/grpc-gateway => github.com/appscode/grpc-gateway v1.3.1-ac
	gomodules.xyz/grpc-go-addons => gomodules.xyz/grpc-go-addons v0.2.2-0.20210218145105-321b2e13985f
)
