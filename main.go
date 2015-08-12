package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/layeh/murmur-cli/MurmurRPC"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	root   = NewCommand()
	void   = &MurmurRPC.Void{}
	ctx    = context.Background()
	client MurmurRPC.V1Client
)

const usage = `murmur-cli provides an interface to a grpc-enabled murmur server.
usage: murmur-cli [flags] [command... [arguments...]]

Flags:
  --address="127.0.0.1:50051"   address and port of murmur's grpc endpoint
                                (can also be set via $MURMUR_ADDRESS).
  --timeout="10s"               duration to wait for connection.
  --template=""                 Go text/template template to use when outputing
                                data. By default, JSON objects are printed.
  --insecure=false              Disable TLS encryption.
  --help                        Print command list.
`

const usageCommands = `
Commands:
  acl get <server id> <channel id>
  acl get-effective-permissions <server id> <session> <channel id>

  ban get <server id>

  channel query <server id>
  channel get <server id> <channel id>
  channel add <server id> <parent channel id> <name>
  channel remove <server id> <channel id>

  config get <server id>
  config get-field <server id> <key>
  config set-field <server id> <key> <value>
  config get-defaults

  contextaction add <server id> <context> <action> <text> <session>
    Context is a comma seperated list of the following:
      Server
      Channel
      User
  contextaction remove <server id> <action> [session]
  contextaction events <server id> <action>

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
  textmessage filter <server id> <program> [args...]

  tree query <server id>

  user query <server id>
  user get <server id> <session>
  user kick <server id> <session> [reason]
`

var outputTemplate *template.Template

func main() {
	// Flags
	defaultAddress := "127.0.0.1:50051"
	if envAddress := os.Getenv("MURMUR_ADDRESS"); envAddress != "" {
		defaultAddress = envAddress
	}

	address := flag.String("address", defaultAddress, "")
	timeout := flag.Duration("timeout", time.Second*10, "")
	templateText := flag.String("template", "", "")
	insecure := flag.Bool("insecure", false, "")

	help := flag.Bool("help", false, "")
	helpShort := flag.Bool("h", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
		if *help || *helpShort {
			fmt.Fprintf(os.Stderr, usageCommands)
		}
	}

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
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(*timeout),
	}
	if !*insecure {
		var tlsConfig tls.Config
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tlsConfig)))
	}
	conn, err := grpc.Dial(*address, opts...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	client = MurmurRPC.NewV1Client(conn)

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
