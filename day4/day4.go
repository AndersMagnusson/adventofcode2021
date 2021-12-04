package day4

import (
	"2021/pkg/inputparsing"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	lines, err := inputparsing.ReadLine(p + "/day4/input.txt")
	if err != nil {
		return err
	}
	draws := strings.Split(lines[0], ",")
	boards := getBoards(lines)
	for _, d := range draws {
		for _, b := range boards {
			res, won := b.Draw(d)
			if won {
				fmt.Println(res)
				return nil
			}
		}
	}
	return nil
}

func Run2() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	lines, err := inputparsing.ReadLine(p + "/day4/input.txt")
	if err != nil {
		return err
	}
	draws := strings.Split(lines[0], ",")
	boards := getBoards(lines)
	wins := 0
	for _, d := range draws {
		for _, b := range boards {
			res, won := b.Draw(d)
			if won {
				wins++
				if wins == len(boards) {
					fmt.Println(res)
					return nil
				}
			}
		}
	}
	return nil
}

func getBoards(lines []string) []*Board {
	items := make([][]string, 5)
	boardIndex := 0
	boards := make([]*Board, 0)

	for i := 1; i < len(lines); i++ {
		if len(lines[i]) <= 1 {
			boardIndex = 0
			items = make([][]string, 5)
			continue
		}
		temp := strings.Split(lines[i], " ")
		numbers := make([]string, 0)
		for _, n := range temp {
			trimmed := strings.TrimSpace(n)
			if len(trimmed) > 0 {
				numbers = append(numbers, trimmed)
			}
		}
		items[boardIndex] = numbers
		boardIndex++
		if boardIndex == 5 {
			board := NewBoard()
			board.Add(items)
			boards = append(boards, board)
		}
	}
	return boards
}

func NewBoard() *Board {
	b := &Board{
		rows:    make([][]*Item, 5),
		columns: make([][]*Item, 5),
		items:   map[string]*Item{},
	}
	return b
}

type Board struct {
	rows    [][]*Item
	columns [][]*Item
	items   map[string]*Item
	won     bool
}

func (b *Board) Add(items [][]string) {
	for rowI, v := range items {
		for columnI, number := range v {
			item := &Item{
				Number: number,
				RowPos: rowI,
				ColPos: columnI,
			}
			b.items[item.Number] = item

		}
	}
}

func (b *Board) Draw(number string) (int, bool) {
	if b.won {
		return 0, false
	}
	item, found := b.items[number]
	if found {
		item.Marked = true
		b.columns[item.ColPos] = append(b.columns[item.ColPos], item)
		b.rows[item.RowPos] = append(b.rows[item.RowPos], item)
		if len(b.columns[item.ColPos]) == 5 || len(b.rows[item.RowPos]) == 5 {
			sum := 0
			drawNumber, err := strconv.Atoi(number)
			if err != nil {
				panic("wrong draw number")
			}
			for _, item := range b.items {
				if !item.Marked {
					num, err := strconv.Atoi(item.Number)
					if err != nil {
						panic("wrong number")
					}
					sum = sum + num
				}
			}
			b.won = true
			return sum * drawNumber, true
		}
	}
	return 0, false
}

type Item struct {
	Number string
	Marked bool
	RowPos int
	ColPos int
}
