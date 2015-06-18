package main

import (
	"./MurmurRPC"

	"google.golang.org/grpc"
)

func initServers(conn *grpc.ClientConn) {
	servers := MurmurRPC.NewServerServiceClient(conn)

	serverCmd := root.Add("server")

	serverCmd.Add("create", func(args []string) {
		Output(servers.Create(ctx, void))
	})

	serverCmd.Add("query", func(args []string) {
		Output(servers.Query(ctx, &MurmurRPC.Server_Query{}))
	})

	serverCmd.Add("get", func(args []string) {
		server := MustServer(args)
		Output(servers.Get(ctx, server))
	})

	serverCmd.Add("start", func(args []string) {
		server := MustServer(args)
		Output(servers.Start(ctx, server))
	})

	serverCmd.Add("stop", func(args []string) {
		server := MustServer(args)
		Output(servers.Stop(ctx, server))
	})

	serverCmd.Add("remove", func(args []string) {
		server := MustServer(args)
		Output(servers.Remove(ctx, server))
	})
}
