package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/layeh/murmur-cli/MurmurRPC"

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
  --template=""                 Go text/template to use when outputing data.
                                By default, JSON objects are generated.

Commands:
  ban get <server id>

  channel query <server id>
  channel get <server id> <channel id>
  channel add <server id> <parent channel id> <name>
  channel remove <server id> <channel id>

  config get <server id>
  config get-field <server id> <key>
  config set-field <server id> <key> <value>
  config get-defaults

  database query <server id> [filter]
  database get <server id> <user id>

  log query <server id> (<min> <max>)

  meta uptime
  meta version
  meta events

  server create
  server query
  server get <server id>
  server start <server id>
  server stop <server id>
  server remove <server id>
  server events <server id>

  textmessage send <server id> [sender:<session>] [targets...] <text>
    Valid targets:
      user:<session>
      channel:<id>
      tree:<id>

  tree query <server id>
`

var outputTemplate *template.Template

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
	templateText := flag.String("template", "", "")
	flag.Parse()

	if *templateText != "" {
		var err error
		outputTemplate, err = template.New("output").Parse(*templateText)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// grpc connection
	conn, err := grpc.Dial(*address, grpc.WithTimeout(*timeout))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	// Services
	initBan(conn)
	initChannel(conn)
	initConfig(conn)
	initDatabase(conn)
	initLog(conn)
	initMeta(conn)
	initServers(conn)
	initTextMessage(conn)
	initTree(conn)

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
