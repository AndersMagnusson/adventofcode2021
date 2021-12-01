package day1

import (
	"2021/pkg/inputparsing"
	"fmt"
	"os"
)

func Run() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	res, err := inputparsing.ReadLineAsInts(p + "/day1/input.txt")
	if err != nil {
		return err
	}
	prev := res[0]
	increased := 0
	for i := 1; i < len(res); i++ {
		if res[i] > prev {
			increased++
		}
		prev = res[i]
	}
	fmt.Println(increased)
	return nil
}

func Run2() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	res, err := inputparsing.ReadLineAsInts(p + "/day1/input.txt")

	prev := res[0] + res[1] + res[2]
	increased := 0
	for i := 1; i < len(res)-2; i++ {
		v := res[i] + res[i+1] + res[i+2]
		if v > prev {
			increased++
		}
		prev = v
	}
	fmt.Println(increased)
	return nil
}
