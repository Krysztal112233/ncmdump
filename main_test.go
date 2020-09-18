package main

import (
	"testing"
)

func TestMakePool(t *testing.T) {
	args := make([]string, 700)
	for k := range args {
		args[k] = "111111111"
	}
	pool, _ := MakePool(args)
	for k, v := range pool {
		t.Logf("%d\t%s", k, v)
	}
}
