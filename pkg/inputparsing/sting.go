package inputparsing

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type StringInt struct {
	Val    string
	ValInt int
}

func ReadLineAsStringInt(filepath string) ([]StringInt, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)

	reader := bufio.NewReader(buf)
	res := make([]StringInt, 0)
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
		vals := strings.Split(string(line), " ")
		if len(vals) != 2 {
			return nil, errors.New("expected two items")
		}

		vInt, err := strconv.Atoi(vals[1])
		if err != nil {
			return nil, err
		}
		v := StringInt{
			Val:    vals[0],
			ValInt: vInt,
		}
		res = append(res, v)
	}
	return res, nil
}
