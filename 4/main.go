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

var Wg sync.WaitGroup
var Mutex sync.Mutex

func GiveYearlySalary(u *User, i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	u.Budget += u.Salary
	fmt.Printf("In month %d, Badget is : %d\n", i+1, u.Budget)
	mutex.Unlock()
}
func main() {

	u := User{Username: "Ghazal", Budget: 100, Salary: 2000}
	for i := 0; i < 12; i++ {
		Wg.Add(1)
		go GiveYearlySalary(&u, i, &Wg, &Mutex)
		Wg.Wait()
	}
	fmt.Printf("Final Budget is : %d", u.Budget)
}
