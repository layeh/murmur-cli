package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("config")

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.GetConfig(ctx, server))
	})

	cmd.Add("get-field", func(args Args) {
		server := args.MustServer(0)
		key := args.MustString(1)
		Output(client.GetConfigField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
		}))
	})

	cmd.Add("set-field", func(args Args) {
		server := args.MustServer(0)
		key := args.MustString(1)
		value := args.MustString(2)
		Output(client.SetConfigField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
			Value:  &value,
		}))
	})

	cmd.Add("get-default", func(args Args) {
		Output(client.GetDefaultConfig(ctx, void))
	})
}
