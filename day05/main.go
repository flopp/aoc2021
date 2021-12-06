package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func mustParseInt(s string) int {
	x, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(x)
}

type xy struct {
	x int
	y int
}

func insertXY(board *map[xy]int, x, y int) {
	p := xy{x, y}
	v, ok := (*board)[p]
	if ok {
		(*board)[p] = v + 1
	} else {
		(*board)[p] = 1
	}
}

func insertLine(board *map[xy]int, x1, y1, x2, y2 int) {
	if x1 == x2 {
		dy := 1
		if y1 > y2 {
			dy = -1
		}
		for y := y1; y != y2; y += dy {
			insertXY(board, x1, y)
		}
		insertXY(board, x2, y2)
	} else if y1 == y2 {
		dx := 1
		if x1 > x2 {
			dx = -1
		}
		for x := x1; x != x2; x += dx {
			insertXY(board, x, y1)
		}
		insertXY(board, x2, y2)
	} else {
		dx := 1
		if x1 > x2 {
			dx = -1
		}
		dy := 1
		if y1 > y2 {
			dy = -1
		}
		for x, y := x1, y1; x != x2 && y != y2; /**/ {
			insertXY(board, x, y)
			x += dx
			y += dy
		}
		insertXY(board, x2, y2)
	}
}

func countCrossings(board *map[xy]int) int {
	sum := 0
	for _, count := range *board {
		if count > 1 {
			sum += 1
		}
	}
	return sum
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

	board := make(map[xy]int)

	re := regexp.MustCompile(`^(\d+),(\d+)\s*->\s*(\d+),(\d+)$`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		x1 := mustParseInt(match[1])
		y1 := mustParseInt(match[2])
		x2 := mustParseInt(match[3])
		y2 := mustParseInt(match[4])

		if part1 {
			if x1 == x2 || y1 == y2 {
				insertLine(&board, x1, y1, x2, y2)
			}
		} else {
			insertLine(&board, x1, y1, x2, y2)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", countCrossings(&board))
}
