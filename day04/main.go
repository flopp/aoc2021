package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/flopp/aoc2021/helpers"
)

type board []int

func createBoard() *board {
	b := board(make([]int, 25))
	return &b
}

func (b *board) setValue(row int, col int, val int) {
	(*b)[row*5+col] = val
}

func (b *board) value(row int, col int) int {
	return (*b)[row*5+col]
}

func (b *board) win(numbers map[int]bool) bool {
	for i := 0; i < 5; i++ {
		row := true
		col := true
		for j := 0; j < 5; j++ {
			if _, exists := numbers[b.value(i, j)]; !exists {
				row = false
			}
			if _, exists := numbers[b.value(j, i)]; !exists {
				col = false
			}
		}
		if row || col {
			return true
		}
	}
	return false
}

func (b *board) score(numbers map[int]bool) int {
	s := 0
	for _, v := range *b {
		if _, exists := numbers[v]; !exists {
			s += v
		}
	}
	return s
}

func main() {
	re := regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s*$`)
	boards := make([]*board, 0)
	numbers := make([]int, 0)
	state := 0
	current_board := (*board)(nil)
	for _, line := range helpers.ReadStdin() {
		switch state {
		case 0:
			for _, v := range strings.Split(line, ",") {
				numbers = append(numbers, helpers.MustParseInt(v))
			}
			state++
		case 1:
			if len(line) != 0 {
				panic(fmt.Errorf("bad line (expected empty line): <%s>", line))
			}
			state++
		default:
			if state == 2 {
				current_board = createBoard()
			}
			match := re.FindStringSubmatch(line)
			if match == nil {
				panic(fmt.Errorf("bad line: <%s>", line))
			}
			for i := 0; i < 5; i++ {
				current_board.setValue(state-2, i, helpers.MustParseInt(match[1+i]))
			}
			if state == 6 {
				boards = append(boards, current_board)
				state = 1
			} else {
				state++
			}
		}
	}

	numbers_set := make(map[int]bool)
	if helpers.Part1() {
		for _, number := range numbers {
			numbers_set[number] = true
			for _, b := range boards {
				if b.win(numbers_set) {
					fmt.Printf("%d\n", number*b.score(numbers_set))
					return
				}
			}
		}
	} else {
		last_score := -1
		already_won := make([]bool, len(boards))
		for _, number := range numbers {
			numbers_set[number] = true
			for index, b := range boards {
				if already_won[index] {
					continue
				}
				if b.win(numbers_set) {
					already_won[index] = true
					last_score = number * b.score(numbers_set)
				}
			}
		}
		fmt.Printf("%d\n", last_score)
	}
}
