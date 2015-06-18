package main

//go:generate mkdir -p MurmurRPC
//go:generate wget -q -O MurmurRPC/MurmurRPC.proto https://raw.githubusercontent.com/bontibon/mumble/grpc/src/murmur/MurmurRPC.proto
//go:generate protoc -IMurmurRPC --go_out=plugins=grpc:MurmurRPC MurmurRPC/MurmurRPC.proto
