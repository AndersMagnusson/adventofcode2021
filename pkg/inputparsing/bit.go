package inputparsing

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
)

func ReadLineAsBits(filepath string) ([][]int, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)

	reader := bufio.NewReader(buf)
	res := make([][]int, 0)
	for {
		line, prefix, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		if prefix {
			continue
		}

		lineVals := make([]int, 0)
		for _, r := range string(line) {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			lineVals = append(lineVals, v)
		}
		res = append(res, lineVals)
	}
	return res, nil
}
