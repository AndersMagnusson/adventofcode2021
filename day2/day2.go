package day2

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
	res, err := inputparsing.ReadLineAsStringInt(p + "/day2/input.txt")
	if err != nil {
		return err
	}
	h := 0
	d := 0
	for _, v := range res {
		switch v.Val {
		case "forward":
			h = h + v.ValInt
		case "down":
			d = d + v.ValInt
		case "up":
			d = d - v.ValInt
		}
	}
	fmt.Println(h * d)
	return nil

}

func Run2() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	res, err := inputparsing.ReadLineAsStringInt(p + "/day2/input.txt")
	if err != nil {
		return err
	}
	h := 0
	d := 0
	a := 0
	for _, v := range res {
		switch v.Val {
		case "forward":
			h = h + v.ValInt
			if a > 0 {
				d = d + (a * v.ValInt)
			}
		case "down":
			a = a + v.ValInt
		case "up":
			a = a - v.ValInt
		}
	}
	fmt.Println(h * d)
	return nil

}
