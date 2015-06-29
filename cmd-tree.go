package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initTree(conn *grpc.ClientConn) {
	client := MurmurRPC.NewTreeServiceClient(conn)

	cmd := root.Add("tree")

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.Get(ctx, server))
	})
}
