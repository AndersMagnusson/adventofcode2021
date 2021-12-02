package main

import (
	"2021/day2"
	"fmt"
)

func main() {
	err := day2.Run()
	if err != nil {
		fmt.Printf("Error1: %s", err)
	}
	err = day2.Run2()
	if err != nil {
		fmt.Printf("Error2: %s", err)
	}
}
