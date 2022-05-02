package main

import (
	"github.com/spf13/cobra"
	"gitlab.com/ondrej.smola/ctcgrpc/cmd/client"
	"gitlab.com/ondrej.smola/ctcgrpc/cmd/server"
	"gitlab.com/ondrej.smola/ctcgrpc/pkg/util"
)

func main() {
	cmd := &cobra.Command{
		Use: "ctcgrpc",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cmd.AddCommand(server.Cmd(), client.Cmd())

	util.ExitOnError(cmd.Execute())
}
