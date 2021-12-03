package day3

import (
	"2021/pkg/inputparsing"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Run() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	res, err := inputparsing.ReadLineAsBits(p + "/day3/input.txt")
	if err != nil {
		return err
	}
	vals, err := getVals(res)
	if err != nil {
		return err
	}

	gamma := ""
	epsilon := ""
	for _, v := range vals {
		if v < 0 {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else if v > 0 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			return fmt.Errorf("got 0")
		}
	}

	gv, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return err
	}
	ev, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return err
	}
	fmt.Println(gv * ev)
	return nil
}

func Run2() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	res, err := inputparsing.ReadLineAsBits(p + "/day3/input.txt")
	if err != nil {
		return err
	}
	// oxygen
	oxygen := res
	for i := 0; i < len(res[0]); i++ {
		v, err := getVals(oxygen)
		if err != nil {
			return err
		}
		if v[i] < 0 {
			oxygen = startsWith(oxygen, 0, i)
		} else {
			oxygen = startsWith(oxygen, 1, i)
		}
		if len(oxygen) == 1 {
			break
		}
	}
	// co2 scrubber
	scrubber := res
	for i := 0; i < len(res[0]); i++ {
		v, err := getVals(scrubber)
		if err != nil {
			return err
		}
		if v[i] >= 0 {
			scrubber = startsWith(scrubber, 0, i)
		} else {
			scrubber = startsWith(scrubber, 1, i)
		}
		if len(scrubber) == 1 {
			break
		}
	}
	if len(oxygen) != 1 || len(scrubber) != 1 {
		return errors.New("expected just one element")
	}

	o, err := getIntFromBits(oxygen[0])
	if err != nil {
		return err
	}
	s, err := getIntFromBits(scrubber[0])
	if err != nil {
		return err
	}
	fmt.Println(o * s)
	return nil
}

func startsWith(numbers [][]int, start int, index int) [][]int {
	res := make([][]int, 0)
	for _, v := range numbers {
		if v[index] == start {
			res = append(res, v)
		}
	}
	return res
}
func getVals(res [][]int) ([]int, error) {
	l := len(res[0])
	vals := make([]int, l)
	for _, line := range res {
		for i, v := range line {
			switch v {
			case 0:
				vals[i] = vals[i] - 1
			case 1:
				vals[i] = vals[i] + 1
			default:
				return nil, fmt.Errorf("expected 0 or 1 but got %d", v)
			}
		}
	}
	return vals, nil
}

func getIntFromBits(vals []int) (int64, error) {
	sv := ""

	for _, v := range vals {
		sv = sv + strconv.Itoa(v)
	}

	return strconv.ParseInt(sv, 2, 64)
}
