package main

import (
	"./MurmurRPC"

	"google.golang.org/grpc"
)

func initMeta(conn *grpc.ClientConn) {
	meta := MurmurRPC.NewMetaServiceClient(conn)

	metaCmd := root.Add("meta")

	metaCmd.Add("uptime", func(args []string) {
		Output(meta.GetUptime(ctx, void))
	})

	metaCmd.Add("version", func(args []string) {
		Output(meta.GetVersion(ctx, void))
	})
}
