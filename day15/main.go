package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/flopp/aoc2021/helpers"
)

type Grid struct {
	width   int
	height  int
	data    []int
	visited []bool
}

func (grid *Grid) Multiply(factor int) Grid {
	w2 := factor * grid.width
	h2 := factor * grid.height
	grid2 := Grid{w2, h2, make([]int, w2*h2), make([]bool, w2*h2)}
	for i, v := range grid.data {
		x := i % grid.width
		y := i / grid.width
		for xx := 0; xx < factor; xx++ {
			for yy := 0; yy < factor; yy++ {
				c := v + xx + yy
				for c > 9 {
					c -= 9
				}
				grid2.data[grid.width*xx+x+w2*(grid.height*yy+y)] = c
			}
		}
	}
	return grid2
}

type PosCost struct {
	pos  int
	cost int
}

func (grid *Grid) next(pos int) []PosCost {
	result := make([]PosCost, 0, 4)
	x := pos % grid.width
	y := pos / grid.width
	if x > 0 && !grid.visited[pos-1] {
		result = append(result, PosCost{pos - 1, grid.data[pos-1]})
	}
	if x+1 < grid.width && !grid.visited[pos+1] {
		result = append(result, PosCost{pos + 1, grid.data[pos+1]})
	}
	if y > 0 && !grid.visited[pos-grid.width] {
		result = append(result, PosCost{pos - grid.width, grid.data[pos-grid.width]})
	}
	if y+1 < grid.height && !grid.visited[pos+grid.width] {
		result = append(result, PosCost{pos + grid.width, grid.data[pos+grid.width]})
	}
	return result
}

type PQItem struct {
	pos   int
	cost  int
	index int
}
type PQ []*PQItem

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	grid := Grid{}
	helpers.ReadStdin(func(line string) {
		if grid.width == 0 {
			grid.width = len(line)
		}
		grid.height++
		for _, c := range line {
			grid.data = append(grid.data, helpers.MustParseInt(string(c)))
		}
	})
	grid.visited = make([]bool, len(grid.data))

	if !helpers.Part1() {
		grid = grid.Multiply(5)
	}

	costItems := make(map[int]*PQItem)
	costPQ := make(PQ, grid.width*grid.height)
	for i := range costPQ {
		item := PQItem{i, math.MaxInt32, i}
		costPQ[i] = &item
		costItems[i] = &item
	}
	// cost of start position is 0
	costPQ[0].cost = 0
	heap.Init(&costPQ)

	for found := false; !found; /**/ {
		minItem := heap.Pop(&costPQ).(*PQItem)
		grid.visited[minItem.pos] = true
		// we've reached the target
		if minItem.pos == len(grid.data)-1 {
			found = true
			break
		}
		for _, next := range grid.next(minItem.pos) {
			nextItem := costItems[next.pos]
			if next.cost+minItem.cost >= nextItem.cost {
				continue
			}
			nextItem.cost = next.cost + minItem.cost
			heap.Fix(&costPQ, nextItem.index)
		}
	}

	fmt.Println(costItems[len(grid.data)-1].cost)
}
