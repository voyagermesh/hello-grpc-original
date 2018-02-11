package cmds

import (
	_ "net/http/pprof"

	"github.com/appscode/go/hold"
	apiCmd "github.com/appscode/hello-grpc/pkg/server/cmd"
	"github.com/appscode/hello-grpc/pkg/server/cmd/options"
	"github.com/spf13/cobra"
)

func NewCmdRun() *cobra.Command {
	opt := options.New()
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Run hello apis",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			apiCmd.Run(opt)
			hold.Hold()
		},
	}

	opt.AddFlags(cmd.Flags())
	return cmd
}
