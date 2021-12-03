package main

import (
	"2021/day3"
	"fmt"
)

func main() {
	err := day3.Run()
	if err != nil {
		fmt.Printf("Error1: %s", err)
	}
	err = day3.Run2()
	if err != nil {
		fmt.Printf("Error2: %s", err)
	}
}
