package main

import (
	"fmt"
	"sync"
)

func squareNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num * num)
}
func main() {
	var wg sync.WaitGroup
	numbers := []int{
		2,
		8,
		100,
		76,
	}
	//if we add more than length of numbers slice, deadlock will heappen
	wg.Add(len(numbers))
	for _, n := range numbers {
		//order is not regular
		go squareNumber(n, &wg)
	}
	wg.Wait()

	//if we don't add 1 , in the squareNumber function the value will become -1 --> PANIC
	wg.Add(1)
	//it prints after array elements always
	squareNumber(200, &wg)
	//we do not need wg.Wait() because we called this function inside main goroutine
}
