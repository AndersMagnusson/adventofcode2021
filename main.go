package main

import (
	"2021/day4"
	"fmt"
)

func main() {
	err := day4.Run()
	if err != nil {
		fmt.Printf("Error1: %s", err)
	}
	err = day4.Run2()
	if err != nil {
		fmt.Printf("Error2: %s", err)
	}
}
