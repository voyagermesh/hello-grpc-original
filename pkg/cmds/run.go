package cmds

import (
	_ "net/http/pprof"

	"voyagermesh.dev/hello-grpc/pkg/cmds/server"
	_ "voyagermesh.dev/hello-grpc/pkg/hello"
	_ "voyagermesh.dev/hello-grpc/pkg/status"
	"github.com/spf13/cobra"
)

func NewCmdRun(stopCh <-chan struct{}) *cobra.Command {
	o := server.NewServerOptions()

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Launch Hello GRPC server",
		Long:  "Launch Hello GRPC server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)
	flags.BoolVar(&o.LogRPC, "log-rpc", o.LogRPC, "log RPC request and response")

	return cmd
}
