package main

import (
	"io"

	"layeh.com/murmur-cli/MurmurRPC"

	"github.com/golang/protobuf/proto"
)

func init() {
	cmd := root.Add("contextaction")

	cmd.Add("add", func(args Args) {
		server := args.MustServer(0)
		context := args.MustBitmask(1, MurmurRPC.ContextAction_Context_value, false)
		action := args.MustString(2)
		text := args.MustString(3)
		session := args.MustUint32(4)
		Output(client.ContextActionAdd(ctx, &MurmurRPC.ContextAction{
			Server:  server,
			Context: proto.Uint32(uint32(context)),
			Action:  &action,
			Text:    &text,
			User: &MurmurRPC.User{
				Session: &session,
			},
		}))
	})

	cmd.Add("remove", func(args Args) {
		server := args.MustServer(0)
		action := args.MustString(1)
		contextAction := &MurmurRPC.ContextAction{
			Server: server,
			Action: &action,
		}
		if session, ok := args.Uint32(2); ok {
			contextAction.User = &MurmurRPC.User{
				Session: &session,
			}
		}
		Output(client.ContextActionRemove(ctx, contextAction))
	})

	cmd.Add("events", func(args Args) {
		server := args.MustServer(0)
		action := args.MustString(1)
		stream, err := client.ContextActionEvents(ctx, &MurmurRPC.ContextAction{
			Server: server,
			Action: &action,
		})
		if err != nil {
			panic(err)
		}
		for {
			msg, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					panic(err)
				}
				return
			}
			Output(msg, nil)
		}
	})
}
