// +build ignore

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Create MurmurRPC/
	err := os.MkdirAll("MurmurRPC", 0755)
	if err != nil {
		panic(err)
	}

	// Fetch MurmurRPC.proto
	resp, err := http.Get("https://raw.githubusercontent.com/mumble-voip/mumble/master/src/murmur/MurmurRPC.proto")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath.Join("MurmurRPC", "MurmurRPC.proto"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		panic(err)
	}

	// Generate protobuf
	cmd := exec.Command("protoc", "-IMurmurRPC", "--go_out=plugins=grpc:MurmurRPC", "MurmurRPC/MurmurRPC.proto")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	// Fix import
	contents, err := ioutil.ReadFile("MurmurRPC/MurmurRPC.pb.go")
	if err != nil {
		panic(err)
	}
	find := []byte("\npackage MurmurRPC\n")
	replace := []byte("\npackage MurmurRPC " + `// import "layeh.com/murmur-cli/MurmurRPC"` + "\n")
	contents = bytes.Replace(contents, find, replace, 1)
	if err := ioutil.WriteFile("MurmurRPC/MurmurRPC.pb.go", contents, 0755); err != nil {
		panic(err)
	}
}
