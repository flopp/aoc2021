package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	for i := 0; i < 5; i += 1 {
		row := true
		col := true
		for j := 0; j < 5; j += 1 {
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
	part1 := true
	switch {
	case os.Args[1] == "part1":
		part1 = true
	case os.Args[1] == "part2":
		part1 = false
	default:
		panic(fmt.Errorf("bad part option: <%s>", os.Args[1]))
	}

	re := regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s*$`)
	boards := make([]*board, 0)
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	state := 0
	current_board := (*board)(nil)
	for scanner.Scan() {
		line := scanner.Text()

		switch state {
		case 0:
			for _, v := range strings.Split(line, ",") {
				number, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, int(number))
			}
			state += 1
		case 1:
			if len(line) != 0 {
				panic(fmt.Errorf("bad line (expected empty line): <%s>", line))
			}
			state += 1
		case 2:
			current_board = createBoard()
			match := re.FindStringSubmatch(line)
			if match == nil {
				panic(fmt.Errorf("bad line: <%s>", line))
			}
			for i := 0; i < 5; i += 1 {
				number, err := strconv.ParseInt(match[1+i], 10, 32)
				if err != nil {
					panic(err)
				}
				current_board.setValue(state-2, i, int(number))
			}
			state += 1
		case 3:
			match := re.FindStringSubmatch(line)
			if match == nil {
				panic(fmt.Errorf("bad line: <%s>", line))
			}
			for i := 0; i < 5; i += 1 {
				number, err := strconv.ParseInt(match[1+i], 10, 32)
				if err != nil {
					panic(err)
				}
				current_board.setValue(state-2, i, int(number))
			}
			state += 1
		case 4:
			match := re.FindStringSubmatch(line)
			if match == nil {
				panic(fmt.Errorf("bad line: <%s>", line))
			}
			for i := 0; i < 5; i += 1 {
				number, err := strconv.ParseInt(match[1+i], 10, 32)
				if err != nil {
					panic(err)
				}
				current_board.setValue(state-2, i, int(number))
			}
			state += 1
		case 5:
			match := re.FindStringSubmatch(line)
			if match == nil {
				panic(fmt.Errorf("bad line: <%s>", line))
			}
			for i := 0; i < 5; i += 1 {
				number, err := strconv.ParseInt(match[1+i], 10, 32)
				if err != nil {
					panic(err)
				}
				current_board.setValue(state-2, i, int(number))
			}
			state += 1
		case 6:
			match := re.FindStringSubmatch(line)
			if match == nil {
				panic(fmt.Errorf("bad line: <%s>", line))
			}
			for i := 0; i < 5; i += 1 {
				number, err := strconv.ParseInt(match[1+i], 10, 32)
				if err != nil {
					panic(err)
				}
				current_board.setValue(state-2, i, int(number))
			}
			boards = append(boards, current_board)
			state = 1
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	numbers_set := make(map[int]bool)
	if part1 {
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
