package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("user")

	cmd.Add("query", func(args Args) {
		server := args.MustServer(0)
		Output(client.UserQuery(ctx, &MurmurRPC.User_Query{
			Server: server,
		}))
	})

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		session := args.MustUint32(1)
		Output(client.UserGet(ctx, &MurmurRPC.User{
			Server:  server,
			Session: &session,
		}))
	})

	cmd.Add("kick", func(args Args) {
		server := args.MustServer(0)
		session := args.MustUint32(1)
		kick := &MurmurRPC.User_Kick{
			Server: server,
			User: &MurmurRPC.User{
				Session: &session,
			},
		}
		if reason, ok := args.String(2); ok {
			kick.Reason = &reason
		}
		Output(client.UserKick(ctx, kick))
	})
}
