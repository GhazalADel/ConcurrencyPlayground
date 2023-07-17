package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"sync"
	"testing"
)

func Test_changeMsg(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go changeMsg("Hello Test", &wg)
	wg.Wait()
	assert.Equal(t, "Hello Test", Msg)
}

func Test_main(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout
	//because of unknown result for section 1 , comment it for testing
	assert.Equal(t, "Bye\n", output)
}
