package main

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/garyburd/redigo/redis"
)

func TestRunner(t *testing.T) {
	r, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}

	defer r.Close()
	addr := r.Addr()

	c, err := redis.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	str := `
SET a 1
RET OK

GET b
RET nil

GET a
RET_LEN 1
RET "1"

MGET a b
RET ["1", nil]

SET a 1.2
RET OK

GET a
RET "1.2"

SET a 2
INCR a
RET 3

SET a -1
GET a
RET "-1"

SET a -1.2
GET a
RET "-1.2"
`

	s := &Scanner{}
	s.Init([]byte(str))

	runner := &ScriptRunner{}
	err = runner.Run(c, s)
	if err != nil {
		t.Fatal(err)
	}
}
