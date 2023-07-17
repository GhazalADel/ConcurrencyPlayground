package main

import (
	"fmt"
	"sync"
)

type User struct {
	Username string
	Budget   int
	Salary   int
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	u := User{Username: "Ghazal", Budget: 100, Salary: 2000}
	//wg.Add(12)
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			u.Budget += u.Salary
			fmt.Printf("In month %d, Badget is : %d\n", i+1, u.Budget)
			mutex.Unlock()
		}(i)
		wg.Wait()
	}
	//wg.Wait()

}
