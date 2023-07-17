package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_GiveYearlySalary(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	u := User{Username: "Test", Budget: 50, Salary: 100}
	for i := 0; i < 3; i++ {
		Wg.Add(1)
		go GiveYearlySalary(&u, i, &Wg, &Mutex)
		Wg.Wait()
	}
	//DON'T FORGET IT -----------------------
	w.Close()
	//DON'T FORGET IT -----------------------

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "Final Budget is : 350") {
		fmt.Errorf("Incorrect")
	}
}
