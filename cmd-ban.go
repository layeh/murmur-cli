package main // import "layeh.com/murmur-cli"

import (
	"layeh.com/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("ban")

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.BansGet(ctx, &MurmurRPC.Ban_Query{
			Server: server,
		}))
	})
}
