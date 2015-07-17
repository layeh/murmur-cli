package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("acl")

	cmd.Add("get", func(args Args) {
		server := args.MustServer(0)
		channelID := args.MustUint32(1)
		Output(client.ACLGet(ctx, &MurmurRPC.Channel{
			Server: server,
			Id:     &channelID,
		}))
	})

	cmd.Add("get-effective-permissions", func(args Args) {
		server := args.MustServer(0)
		session := args.MustUint32(1)
		channelID := args.MustUint32(2)
		Output(client.ACLGetEffectivePermissions(ctx, &MurmurRPC.ACL_Query{
			Server: server,
			User: &MurmurRPC.User{
				Session: &session,
			},
			Channel: &MurmurRPC.Channel{
				Id: &channelID,
			},
		}))
	})

}
