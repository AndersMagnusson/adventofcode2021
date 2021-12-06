package day6

import (
	"2021/pkg/inputparsing"
	"fmt"
	"os"
)

func Run(days int) error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	fishes, err := inputparsing.ReadLineWithSepAsInts(p+"/day6/input.txt", ",")
	if err != nil {
		return err
	}
	start := [9]int{}
	for i := 0; i < len(fishes); i++ {
		start[fishes[i]]++
	}

	for i := 0; i < days; i++ {
		next := [9]int{}
		for ii := 0; ii < 9; ii++ {
			switch ii {
			case 0:
				next[6] += start[0]
				next[8] += start[0]
			default:
				next[ii-1] += start[ii]
			}
		}
		start = next
	}
	res := 0
	for _, v := range start {
		res += v
	}
	fmt.Println(res)
	return nil
}
