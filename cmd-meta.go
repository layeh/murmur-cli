package main

import (
	"io"

	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initMeta(conn *grpc.ClientConn) {
	meta := MurmurRPC.NewMetaServiceClient(conn)

	cmd := root.Add("meta")

	cmd.Add("uptime", func(args []string) {
		Output(meta.GetUptime(ctx, void))
	})

	cmd.Add("version", func(args []string) {
		Output(meta.GetVersion(ctx, void))
	})

	cmd.Add("events", func(args []string) {
		stream, err := meta.Events(ctx, void)
		if err != nil {
			Output(nil, err)
			return
		}
		for {
			msg, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					Output(nil, err)
				}
				return
			}
			Output(msg, nil)
		}
	})
}
