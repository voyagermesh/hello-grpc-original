package main

import (
	"os"

	"gomodules.xyz/kglog"
	"voyagermesh.dev/hello-grpc/pkg/cmds"
)

func main() {
	kglog.InitLogs()
	defer kglog.FlushLogs()

	if err := cmds.NewRootCmd(Version).Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
