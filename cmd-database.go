package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("database")

	cmd.Add("query", func(args Args) {
		server := args.MustServer(0)
		query := &MurmurRPC.DatabaseUser_Query{
			Server: server,
		}
		if filter, ok := args.String(1); ok {
			query.Filter = &filter
		}
		Output(client.QueryDatabaseUsers(ctx, query))
	})

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		id := args.MustUint32(1)
		Output(client.GetDatabaseUser(ctx, &MurmurRPC.DatabaseUser{
			Server: server,
			Id:     &id,
		}))
	})
}
