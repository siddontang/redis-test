package main

import (
	"reflect"
	"testing"
)

func TestScanner(t *testing.T) {
	s := &Scanner{}

	str := `
# abc
SET a 1
RET OK

SET b [1, 2, "abc"]
RET OK

GET b
RET [1, 2, "abc"]

RET ["a", "b", ["a", "b", ["a"]]]
`

	s.Init([]byte(str))

	checkScanCommands(t, s, "SET", "a", int64(1))
	checkScanCommands(t, s, "RET", "OK")
	checkScanCommands(t, s, "SET", "b", []interface{}{int64(1), int64(2), "abc"})
	checkScanCommands(t, s, "RET", "OK")
	checkScanCommands(t, s, "GET", "b")
	checkScanCommands(t, s, "RET", []interface{}{int64(1), int64(2), "abc"})
	checkScanCommands(t, s, "RET", []interface{}{"a", "b", []interface{}{
		"a", "b", []interface{}{"a"},
	}})
}

func checkScanCommands(t *testing.T, s *Scanner, expected ...interface{}) {
	got := s.ScanCommand()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("%v != %v", expected, got)
	}
}
