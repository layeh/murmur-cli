package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initBan(conn *grpc.ClientConn) {
	client := MurmurRPC.NewBanServiceClient(conn)

	cmd := root.Add("ban")

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.Get(ctx, &MurmurRPC.Ban_Query{
			Server: server,
		}))
	})
}
