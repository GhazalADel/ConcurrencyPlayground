package main

import (
	"fmt"
	"sync"
)

var Msg string

func changeMsg(msg2 string, wg *sync.WaitGroup) {
	defer wg.Done()
	Msg = msg2
}
func main() {
	var wg sync.WaitGroup

	//--------SECTION 1 --------
	wg.Add(3)
	go changeMsg("Hello2", &wg)
	go changeMsg("How Are You2", &wg)
	go changeMsg("Bye2", &wg)
	wg.Wait()

	fmt.Println(Msg) //UNKNOWN Result

	//--------SECTION 2 --------
	wg.Add(1)
	go changeMsg("Hello", &wg)
	wg.Wait()

	wg.Add(1)
	go changeMsg("How Are You", &wg)
	wg.Wait()

	wg.Add(1)
	go changeMsg("Bye", &wg)
	wg.Wait()

	fmt.Println(Msg) // Bye

}
