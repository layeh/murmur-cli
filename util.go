package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/layeh/murmur-cli/MurmurRPC"

	"github.com/golang/protobuf/proto"
)

func Output(data interface{}, err error) {
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}

func MustServer(args []string) *MurmurRPC.Server {
	if len(args) <= 0 {
		panic(errors.New("missing server ID argument"))
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		panic(errors.New("invalid server ID"))
	}
	return &MurmurRPC.Server{
		Id: proto.Uint32(uint32(id)),
	}
}

func MustString(args []string, index int) string {
	if len(args) < index {
		panic(errors.New("missing string value"))
	}
	return args[index]
}
