package main

import (
	"io"
)

func init() {
	cmd := root.Add("meta")

	cmd.Add("uptime", func(args Args) {
		Output(client.GetUptime(ctx, void))
	})

	cmd.Add("version", func(args Args) {
		Output(client.GetVersion(ctx, void))
	})

	cmd.Add("events", func(args Args) {
		stream, err := client.Events(ctx, void)
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
