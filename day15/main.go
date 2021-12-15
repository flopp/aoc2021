package main

import (
	"fmt"
	"math"

	"github.com/flopp/aoc2021/helpers"
)

type Grid struct {
	Width  int
	Height int
	data   []int
}

func (grid *Grid) appendRow(row []int) {
	if grid.Width == 0 {
		grid.Width = len(row)
	} else if grid.Width != len(row) {
		panic("bad input")
	}
	grid.Height++
	for _, cell := range row {
		grid.data = append(grid.data, cell)
	}
}

func (grid *Grid) Pos(x int, y int) int {
	return x + y*grid.Width
}

func (grid *Grid) Multiply(factor int) Grid {
	grid2 := Grid{factor * grid.Width, factor * grid.Height, make([]int, factor*factor*grid.Width*grid.Height)}
	for i, v := range grid.data {
		x := i % grid.Width
		y := i / grid.Width
		for xx := 0; xx < factor; xx++ {
			for yy := 0; yy < factor; yy++ {
				c := v + xx + yy
				for c > 9 {
					c -= 9
				}
				grid2.data[grid2.Pos(grid.Width*xx+x, grid.Height*yy+y)] = c
			}
		}
	}
	return grid2
}

type PosCost struct {
	Pos  int
	Cost int
}

func (grid *Grid) next(pos int) []PosCost {
	result := make([]PosCost, 0, 4)
	x := pos % grid.Width
	y := pos / grid.Width
	if x > 0 {
		result = append(result, PosCost{pos - 1, grid.data[pos-1]})
	}
	if x+1 < grid.Width {
		result = append(result, PosCost{pos + 1, grid.data[pos+1]})
	}
	if y > 0 {
		result = append(result, PosCost{pos - grid.Width, grid.data[pos-grid.Width]})
	}
	if y+1 < grid.Height {
		result = append(result, PosCost{pos + grid.Width, grid.data[pos+grid.Width]})
	}
	return result
}

func main() {
	grid := Grid{}
	helpers.ReadStdin(func(line string) {
		row := make([]int, 0, len(line))
		for _, c := range line {
			row = append(row, helpers.MustParseInt(string(c)))
		}
		grid.appendRow(row)
	})

	if !helpers.Part1() {
		grid = grid.Multiply(5)
	}

	start := grid.Pos(0, 0)
	target := grid.Pos(grid.Width-1, grid.Height-1)
	visited := make([]bool, grid.Width*grid.Height)

	cost := make([]int, grid.Width*grid.Height)
	for i := range cost {
		cost[i] = math.MaxInt32
	}
	cost[start] = 0

	found := false
	for !found {
		minCost := math.MaxInt32
		minPos := -1
		for i, c := range cost {
			if !visited[i] && c < minCost {
				minCost = c
				minPos = i
			}
		}
		visited[minPos] = true
		if minPos == target {
			found = true
			break
		}
		for _, next := range grid.next(minPos) {
			if !visited[next.Pos] && cost[next.Pos] > next.Cost+minCost {
				cost[next.Pos] = next.Cost + minCost
			}
		}
	}

	fmt.Println(cost[target])
}
