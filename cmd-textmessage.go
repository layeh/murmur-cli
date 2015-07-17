package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"
)

func init() {
	cmd := root.Add("textmessage")

	cmd.Add("send", func(args Args) {
		server := args.MustServer(0)
		if len(args) <= 1 {
			panic("not enough arguments")
		}
		text := args.MustString(len(args) - 1)

		tm := &MurmurRPC.TextMessage{
			Server: server,
			Text:   &text,
		}

		for i := 1; i < len(args)-1; i++ {
			if session, ok := args.PrefixedUint32("sender:", i); ok && tm.Actor == nil {
				tm.Actor = &MurmurRPC.User{
					Server:  server,
					Session: &session,
				}
			} else if session, ok := args.PrefixedUint32("user:", i); ok {
				tm.Users = append(tm.Users, &MurmurRPC.User{
					Server:  server,
					Session: &session,
				})
			} else if id, ok := args.PrefixedUint32("channel:", i); ok {
				tm.Channels = append(tm.Channels, &MurmurRPC.Channel{
					Server: server,
					Id:     &id,
				})
			} else if id, ok := args.PrefixedUint32("tree:", i); ok {
				tm.Trees = append(tm.Trees, &MurmurRPC.Channel{
					Server: server,
					Id:     &id,
				})
			} else {
				panic("unknown argument")
			}
		}

		Output(client.TextMessageSend(ctx, tm))
	})
}
