package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initConfig(conn *grpc.ClientConn) {
	meta := MurmurRPC.NewConfigServiceClient(conn)

	cmd := root.Add("config")

	cmd.Add("get", func(args []string) {
		server := MustServer(args)
		Output(meta.Get(ctx, server))
	})

	cmd.Add("get-field", func(args []string) {
		server := MustServer(args)
		key := MustString(args, 1)
		Output(meta.GetField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
		}))
	})

	cmd.Add("set-field", func(args []string) {
		server := MustServer(args)
		key := MustString(args, 1)
		value := MustString(args, 2)
		Output(meta.SetField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
			Value:  &value,
		}))
	})

	cmd.Add("get-defaults", func(args []string) {
		Output(meta.GetDefaults(ctx, void))
	})
}
