package main

import (
	"github.com/pradlerfilip/CTC_2022/ctcgrpc/cmd/client"
	"github.com/pradlerfilip/CTC_2022/ctcgrpc/cmd/server"
	"github.com/pradlerfilip/CTC_2022/ctcgrpc/pkg/util"
	"github.com/spf13/cobra"
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
