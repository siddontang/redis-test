package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/garyburd/redigo/redis"
)

func exitf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

func main() {
	usage := `
Usage:
    redis-test FILE [options]

Options:
    -h <hostname>      Server hostname (default: 127.0.0.1).
    -p <port>          Server port (default: 6379).
    -a <password>      Password to use when connecting to the server.
`
	d, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		exitf("Parse arguments failed, err: %v\n", err)
	}

	fileName := d["FILE"].(string)
	f, err := os.Open(fileName)
	if err != nil {
		exitf("Open script %s err: %v\n", fileName, err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		exitf("Read file %s err: %v\n", fileName, err)
	}

	host := "127.0.0.1"
	port := "6379"

	if s, ok := d["-h"].(string); ok && len(s) != 0 {
		host = s
	}

	if s, ok := d["-p"].(string); ok && len(s) != 0 {
		port = s
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		exitf("Dial Redis %s err %v\n", addr, err)
	}
	defer c.Close()

	s := &Scanner{}
	s.Init(data)

	r := &ScriptRunner{}
	err = r.Run(c, s)
	if err != nil {
		exitf("Run script %s err :%v\n", fileName, err)
	}
}
