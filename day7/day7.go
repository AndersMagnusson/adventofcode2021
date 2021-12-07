package day7

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
	pos, err := inputparsing.ReadLineWithSepAsInts(p+"/day7/test.txt", ",")
	if err != nil {
		return err
	}
	total := 0
	for _, v := range pos {
		total += v
	}
	start := total / len(pos)
	bestFuel := avg(start, pos)
	i := 1
	for {
		if start-i < 0 || start+1 >= len(pos) {
			break
		}
		down := avg(start-i, pos)
		up := avg(start+i, pos)
		if down < bestFuel {
			bestFuel = down
		} else if up < bestFuel {
			bestFuel = up
		}
		i++
	}
	fmt.Println(bestFuel)
	return nil
}
func Run2() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	pos, err := inputparsing.ReadLineWithSepAsInts(p+"/day7/input.txt", ",")
	if err != nil {
		return err
	}
	total := 0
	for _, v := range pos {
		total += v
	}
	start := total / len(pos)
	bestFuel := avgInc(start, pos)

	i := 1
	for {
		if start-i < 0 || start+1 >= len(pos) {
			break
		}
		down := avgInc(start-i, pos)
		up := avgInc(start+i, pos)
		if down < bestFuel {
			bestFuel = down
		} else if up < bestFuel {
			bestFuel = up
		}
		i++
	}
	fmt.Println(bestFuel)
	return nil
}

func avg(avg int, pos []int) int {
	fuel := 0
	for _, steps := range pos {
		if steps < avg {
			fuel = fuel + avg - steps
		} else {
			fuel = fuel + steps - avg
		}
	}
	return fuel
}
func avgInc(avg int, pos []int) int {
	fuel := 0
	for _, steps := range pos {
		if steps < avg {
			diff := avg - steps
			sum := 0
			for i := 1; i < diff; i++ {
				sum += i
			}
			fuel = fuel + avg - steps + sum
		} else {
			diff := steps - avg
			sum := 0
			for i := 1; i < diff; i++ {
				sum += i
			}
			fuel = fuel + steps - avg + sum
		}
	}
	return fuel
}
