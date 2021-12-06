package day5

import (
	"2021/pkg/inputparsing"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(includeDiagonal bool) error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	lines, err := inputparsing.ReadLine(p + "/day5/input.txt")
	if err != nil {
		return err
	}
	items := make([]Line, 0)
	for _, l := range lines {
		values := strings.Split(strings.TrimSpace(strings.Replace(l, "->", ",", -1)), ",")
		if len(values) != 4 {
			return errors.New("expected 4 items")
		}
		line, err := NewLine(values[0], values[1], values[2], values[3])
		if err != nil {
			return err
		}
		items = append(items, line)
	}
	board := make(map[string]int)
	for _, pos := range items {
		pos.Mark(board, includeDiagonal)
	}
	count := 0
	for _, pos := range board {
		if pos >= 2 {
			count++
		}
	}
	fmt.Println(count)
	return nil
}

type Line struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
	x      int
	y      int
}

func (l *Line) LowestX() int {
	if l.StartX < l.EndX {
		return l.StartX
	}
	return l.EndX
}
func (l *Line) HigestX() int {
	if l.StartX > l.EndX {
		return l.StartX
	}
	return l.EndX
}
func (l *Line) LowestY() int {
	if l.StartY < l.EndY {
		return l.StartY
	}
	return l.EndY
}
func (l *Line) HighestY() int {
	if l.StartY > l.EndY {
		return l.StartY
	}
	return l.EndY
}
func (l *Line) IsDiagonal() bool {
	return l.StartX != l.EndX && l.StartY != l.EndY
}
func (l *Line) IsHorisontel() bool {
	return l.StartX != l.EndX
}
func (l *Line) IsVertical() bool {
	return l.StartY != l.EndY
}
func (l *Line) Pos() string {
	return fmt.Sprintf("%dx%d", l.x, l.y)
}
func (l *Line) Start() {
	l.x = l.StartX
	l.y = l.StartY
}
func (l *Line) NextDiagonal() bool {
	step := 1
	endX := l.EndX
	startX := l.StartX
	if l.EndX < l.StartX {
		step = -1
		endX = l.StartX
		startX = l.EndX
	}
	x := l.x + step
	if x < startX || x > endX {
		return false
	}
	l.x = x

	step = 1
	if l.EndY < l.StartY {
		step = -1
	}
	y := l.y + step
	l.y = y
	return true
}

func (l Line) Mark(board map[string]int, includeDiagonal bool) {
	if l.IsDiagonal() {
		if !includeDiagonal {
			return
		}
		l.Start()
		for {
			v, found := board[l.Pos()]
			if found {
				v++
			} else {
				v = 1
			}
			board[l.Pos()] = v
			if !l.NextDiagonal() {
				return
			}
		}
	}
	if l.IsHorisontel() {
		// Horisontel
		for i := l.LowestX(); i <= l.HigestX(); i++ {
			pos := fmt.Sprintf("%dx%d", i, l.StartY)
			v, found := board[pos]
			if found {
				v++
			} else {
				v = 1
			}
			board[pos] = v
		}
		return
	}
	if l.IsVertical() {
		// Vertical
		for i := l.LowestY(); i <= l.HighestY(); i++ {
			pos := fmt.Sprintf("%dx%d", l.StartX, i)
			v, found := board[pos]
			if found {
				v++
			} else {
				v = 1
			}
			board[pos] = v
		}
		return
	}
}

func NewLine(startX, startY, endX, endY string) (Line, error) {
	l := Line{}
	x, err := strconv.Atoi(strings.TrimSpace(startX))
	if err != nil {
		return l, err
	}
	y, err := strconv.Atoi(strings.TrimSpace(startY))
	if err != nil {
		return l, err
	}
	endx, err := strconv.Atoi(strings.TrimSpace(endX))
	if err != nil {
		return l, err
	}
	endyy, err := strconv.Atoi(strings.TrimSpace(endY))
	if err != nil {
		return l, err
	}
	l.StartX = x
	l.StartY = y
	l.EndX = endx
	l.EndY = endyy
	return l, nil
}
