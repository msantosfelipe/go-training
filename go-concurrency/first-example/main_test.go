package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	msg := "epsilon"
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething(msg, &wg)
	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, msg) {
		t.Errorf("expected to find %s but it is not there", msg)
	}
}
