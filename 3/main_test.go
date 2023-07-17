package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_changeMsg(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(2)
	go changeMsg("ABC", &wg, &mutex)
	go changeMsg("EFG", &wg, &mutex)
	wg.Wait()
	if Msg != "ABC" && Msg != "EFG" {
		assert.Error(t, errors.New("Incorrect"))
	}
}
