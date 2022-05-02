package server

import (
	"github.com/pradlerfilip/CTC_2022/ctcgrpc/pkg/api"
	"github.com/pradlerfilip/CTC_2022/ctcgrpc/pkg/store"
	"github.com/pradlerfilip/CTC_2022/ctcgrpc/pkg/util"
	"github.com/spf13/cobra"
	etcd "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"net"
	"time"
)

func Cmd() *cobra.Command {
	var etcdEndpoints []string
	var listen string

	cmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			etcdCli, err := etcd.New(etcd.Config{
				Endpoints:   etcdEndpoints,
				DialTimeout: 3 * time.Second,
				DialOptions: []grpc.DialOption{grpc.WithReturnConnectionError()},
			})
			util.ExitOnError(err)

			svr := api.NewServer(store.NewEtcd(etcdCli))

			lis, err := net.Listen("tcp", listen)
			util.ExitOnError(err)

			grpcServer := grpc.NewServer()
			api.RegisterApiServer(grpcServer, svr)
			util.ExitOnError(grpcServer.Serve(lis))
		},
	}

	f := cmd.Flags()
	f.StringSliceVar(&etcdEndpoints, "etcd", []string{"localhost:2379"}, "")
	f.StringVar(&listen, "listen", ":8080", "listen address")

	return cmd
}
