package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initLog(conn *grpc.ClientConn) {
	log := MurmurRPC.NewLogServiceClient(conn)

	cmd := root.Add("log")

	cmd.Add("query", func(args []string) {
		server := MustServer(args)
		query := &MurmurRPC.Log_Query{
			Server: server,
		}
		if len(args) > 1 {
			min := MustUint32(args, 1)
			max := MustUint32(args, 2)
			query.Min = &min
			query.Max = &max
		}
		Output(log.Query(ctx, query))
	})
}
