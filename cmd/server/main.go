package main

import (
	"fmt"
	"os"

	"github.com/mkyc/go-grpc-pb-test1/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
