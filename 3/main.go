package main

import (
	"fmt"
	"sync"
)

var Msg string

func changeMsg(msg2 string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	Msg = msg2
	mutex.Unlock()
}

func main() {
	//for avoid race condition in last exercise, we can use mutex
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(3)
	go changeMsg("Hello2", &wg, &mutex)
	go changeMsg("How Are You2", &wg, &mutex)
	go changeMsg("Bye2", &wg, &mutex)
	wg.Wait()

	fmt.Println(Msg) //UNKNOWN Result

}
