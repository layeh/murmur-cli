package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("tree")

	cmd.Add("query", func(args Args) {
		server := args.MustServer(0)
		Output(client.TreeQuery(ctx, &MurmurRPC.Tree_Query{
			Server: server,
		}))
	})
}
