package main

import (
	"os"
	"testing"
)

func TestLibWasmExecJSPath(t *testing.T) {
	path, err := libWasmExecJSPath()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fi.Name())
}
