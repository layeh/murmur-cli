package main

import (
	"layeh.com/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("config")

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		Output(client.ConfigGet(ctx, server))
	})

	cmd.Add("get-field", func(args Args) {
		server := args.MustServer(0)
		key := args.MustString(1)
		Output(client.ConfigGetField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
		}))
	})

	cmd.Add("set-field", func(args Args) {
		server := args.MustServer(0)
		key := args.MustString(1)
		value := args.MustString(2)
		Output(client.ConfigSetField(ctx, &MurmurRPC.Config_Field{
			Server: server,
			Key:    &key,
			Value:  &value,
		}))
	})

	cmd.Add("get-default", func(args Args) {
		Output(client.ConfigGetDefault(ctx, void))
	})
}
