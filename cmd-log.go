package main // import "layeh.com/murmur-cli"

import (
	"layeh.com/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("log")

	cmd.Add("query", func(args Args) {
		server := args.MustServer(0)
		query := &MurmurRPC.Log_Query{
			Server: server,
		}
		if len(args) > 1 {
			min := args.MustUint32(1)
			max := args.MustUint32(2)
			query.Min = &min
			query.Max = &max
		}
		Output(client.LogQuery(ctx, query))
	})
}
