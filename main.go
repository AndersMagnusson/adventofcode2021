package main

import (
	"2021/day1"
	"fmt"
)

func main() {
	err := day1.Run()
	if err != nil {
		fmt.Printf("Error1: %s", err)
	}
	err = day1.Run2()
	if err != nil {
		fmt.Printf("Error2: %s", err)
	}
}
