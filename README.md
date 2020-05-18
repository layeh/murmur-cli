# murmur-cli

murmur-cli provides an interface to a grpc-enabled murmur server.

    usage: murmur-cli [flags] [command... [arguments...]]

    Flags:
      --address="127.0.0.1:50051"   address and port of murmur's grpc endpoint
                                    (can also be set via $MURMUR_ADDRESS).
      --timeout="10s"               duration to wait for connection.
      --template=""                 Go text/template template to use when outputting
                                    data. By default, JSON objects are printed.
      --insecure=false              Disable TLS encryption.
      --help                        Print command list.

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
      config get-default

      contextaction add <server id> <context> <action> <text> <session>
        Context is a comma separated list of the following:
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

## Installation

    go get -u layeh.com/murmur-cli

## Author

Tim Cooper (<tim.cooper@layeh.com>)
