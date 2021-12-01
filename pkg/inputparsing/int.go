package inputparsing

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
)

func ReadLineAsInts(filepath string) ([]int, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)

	reader := bufio.NewReader(buf)
	res := make([]int, 0)
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
		v, err := strconv.Atoi(string(line))
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	return res, nil
}
