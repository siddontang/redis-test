package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `
Usage:
    redis-test FILE [options]

Options:
    -h <hostname>      Server hostname (default: 127.0.0.1).
    -p <port>          Server port (default: 6379).
    -s <socket>        Server socket (overrides hostname and port).
    -a <password>      Password to use when connecting to the server.
`
	d, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		fmt.Printf("Parse arguments failed, err: %v\n", err)
		return
	}

	fmt.Printf("hello %v\n", d)
}
