package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"./MurmurRPC"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	root = NewCommand()
	void = &MurmurRPC.Void{}
	ctx  = context.Background()
)

const usage = `murmur-cli provides an interface to a grpc-enabled murmur server.
usage: murmur-cli [flags] [command... [arguments...]]

Flags:
  --address="127.0.0.1:50051"   address and port of murmur's grpc endpoint
                                (can also be set via $MURMUR_ADDRESS)
  --timeout="10s"               duration to wait for connection

Commands:
  meta uptime
  meta version

  server create
  server query
  server get <id>
  server start <id>
  server stop <id>
  server remove <id>
`

func main() {
	// Flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}

	defaultAddress := "127.0.0.1:50051"
	if envAddress := os.Getenv("MURMUR_ADDRESS"); envAddress != "" {
		defaultAddress = envAddress
	}

	address := flag.String("address", defaultAddress, "")
	timeout := flag.Duration("timeout", time.Second*10, "")
	flag.Parse()

	// grpc connection
	conn, err := grpc.Dial(*address, grpc.WithTimeout(*timeout))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	// Services
	initMeta(conn)
	initServers(conn)

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				jsonErr := struct {
					Error string `json:"error"`
				}{
					Error: err.Error(),
				}
				json.NewEncoder(os.Stderr).Encode(&jsonErr)
				os.Exit(3)
			}
		}
	}()

	if root.Do() != nil {
		flag.Usage()
		os.Exit(1)
	}
}
