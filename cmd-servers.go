package main

import (
	"io"

	"github.com/layeh/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("server")

	cmd.Add("create", func(args Args) {
		Output(client.CreateServer(ctx, void))
	})

	cmd.Add("query", func(args Args) {
		Output(client.QueryServers(ctx, &MurmurRPC.Server_Query{}))
	})

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.GetServer(ctx, server))
	})

	cmd.Add("start", func(args Args) {
		server := args.MustServer(0)
		Output(client.StartServer(ctx, server))
	})

	cmd.Add("stop", func(args Args) {
		server := args.MustServer(0)
		Output(client.StopServer(ctx, server))
	})

	cmd.Add("remove", func(args Args) {
		server := args.MustServer(0)
		Output(client.RemoveServer(ctx, server))
	})

	cmd.Add("events", func(args Args) {
		server := args.MustServer(0)
		stream, err := client.ServerEvents(ctx, server)
		if err != nil {
			panic(err)
		}
		for {
			msg, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					panic(err)
				}
				return
			}
			Output(msg, nil)
		}
	})
}
