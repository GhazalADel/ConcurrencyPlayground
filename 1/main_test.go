package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"sync"
	"testing"
)

func Test_squareNumber(t *testing.T) {
	//we have to keep stdout
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go squareNumber(11, &wg)
	wg.Wait()

	//if you forget it , a loop will create
	w.Close()

	res, _ := io.ReadAll(r)
	output := string(res)

	os.Stdout = stdout

	assert.Equal(t, "121\n", output)
}
