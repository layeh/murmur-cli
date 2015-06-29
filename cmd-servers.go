package main

import (
	"io"

	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initServers(conn *grpc.ClientConn) {
	client := MurmurRPC.NewServerServiceClient(conn)

	cmd := root.Add("server")

	cmd.Add("create", func(args Args) {
		Output(client.Create(ctx, void))
	})

	cmd.Add("query", func(args Args) {
		Output(client.Query(ctx, &MurmurRPC.Server_Query{}))
	})

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.Get(ctx, server))
	})

	cmd.Add("start", func(args Args) {
		server := args.MustServer(0)
		Output(client.Start(ctx, server))
	})

	cmd.Add("stop", func(args Args) {
		server := args.MustServer(0)
		Output(client.Stop(ctx, server))
	})

	cmd.Add("remove", func(args Args) {
		server := args.MustServer(0)
		Output(client.Remove(ctx, server))
	})

	cmd.Add("events", func(args Args) {
		server := args.MustServer(0)
		stream, err := client.Events(ctx, server)
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
