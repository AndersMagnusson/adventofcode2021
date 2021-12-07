package main

import (
	"2021/day7"
	"fmt"
)

func main() {
	err := day7.Run2()
	if err != nil {
		fmt.Printf("Error1: %s", err)
	}
	// err = day6.Run(256)
	// if err != nil {
	// 	fmt.Printf("Error2: %s", err)
	// }
}
