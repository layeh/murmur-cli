package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/layeh/murmur-cli/MurmurRPC"

	"github.com/golang/protobuf/proto"
)

func Output(data interface{}, err error) {
	if err != nil {
		panic(err)
	}
	if outputTemplate != nil {
		if err := outputTemplate.Execute(os.Stdout, data); err != nil {
			panic(err)
		}
		return
	}
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}

type Args []string

func (a Args) MustServer(i int) *MurmurRPC.Server {
	if len(a) <= i {
		panic(errors.New("missing server ID argument"))
	}
	id, err := strconv.Atoi(a[i])
	if err != nil {
		panic(errors.New("invalid server ID"))
	}
	return &MurmurRPC.Server{
		Id: proto.Uint32(uint32(id)),
	}
}

func (a Args) MustString(i int) string {
	if len(a) <= i {
		panic(errors.New("missing string value"))
	}
	return a[i]
}

func (a Args) String(i int) (string, bool) {
	if len(a) <= i {
		return "", false
	}
	return a[i], true
}

func (a Args) MustUint32(i int) uint32 {
	if len(a) <= i {
		panic(errors.New("missing uint32 value"))
	}
	n, err := strconv.ParseUint(a[i], 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(n)
}

func (a Args) PrefixedUint32(prefix string, i int) (uint32, bool) {
	if len(a) <= i {
		panic(errors.New("missing prefixed uint32 value"))
	}
	if !strings.HasPrefix(a[i], prefix) {
		return 0, false
	}
	n, err := strconv.ParseUint(a[i][len(prefix):], 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(n), true
}
