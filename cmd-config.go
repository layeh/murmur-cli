package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initConfig(conn *grpc.ClientConn) {
	client := MurmurRPC.NewConfigServiceClient(conn)

	cmd := root.Add("config")

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.Get(ctx, server))
	})

	cmd.Add("get-field", func(args Args) {
		server := args.MustServer(0)
		key := args.MustString(1)
		Output(client.GetField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
		}))
	})

	cmd.Add("set-field", func(args Args) {
		server := args.MustServer(0)
		key := args.MustString(1)
		value := args.MustString(2)
		Output(client.SetField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
			Value:  &value,
		}))
	})

	cmd.Add("get-defaults", func(args Args) {
		Output(client.GetDefaults(ctx, void))
	})
}
