package main

import (
	"io"

	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initServers(conn *grpc.ClientConn) {
	servers := MurmurRPC.NewServerServiceClient(conn)

	cmd := root.Add("server")

	cmd.Add("create", func(args []string) {
		Output(servers.Create(ctx, void))
	})

	cmd.Add("query", func(args []string) {
		Output(servers.Query(ctx, &MurmurRPC.Server_Query{}))
	})

	cmd.Add("get", func(args []string) {
		server := MustServer(args)
		Output(servers.Get(ctx, server))
	})

	cmd.Add("start", func(args []string) {
		server := MustServer(args)
		Output(servers.Start(ctx, server))
	})

	cmd.Add("stop", func(args []string) {
		server := MustServer(args)
		Output(servers.Stop(ctx, server))
	})

	cmd.Add("remove", func(args []string) {
		server := MustServer(args)
		Output(servers.Remove(ctx, server))
	})

	cmd.Add("events", func(args []string) {
		server := MustServer(args)
		stream, err := servers.Events(ctx, server)
		if err != nil {
			Output(nil, err)
			return
		}
		for {
			msg, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					Output(nil, err)
				}
				return
			}
			Output(msg, nil)
		}
	})
}
