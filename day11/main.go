package main

import (
	"fmt"

	"github.com/flopp/aoc2021/helpers"
)

type XY struct {
	x int
	y int
}

func FromIndex(index int) XY {
	return XY{index % 10, index / 10}
}
func (xy XY) Valid() bool {
	return xy.x >= 0 && xy.x < 10 && xy.y >= 0 && xy.y < 10
}
func (xy XY) Index() int {
	return xy.x + xy.y*10
}
func (xy XY) Plus(offset XY) XY {
	return XY{xy.x + offset.x, xy.y + offset.y}
}
func (xy XY) Adjacent() []XY {
	result := make([]XY, 0, 8)
	for _, d := range []XY{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		a := xy.Plus(d)
		if a.Valid() {
			result = append(result, a)
		}
	}
	return result
}

func Step(energyLevels []int) int {
	flashes := 0

	// step 1: increment energy levels
	for i := range energyLevels {
		energyLevels[i]++
	}

	// step 2: flash
	again := true
	for again {
		again = false
		for i := range energyLevels {
			if energyLevels[i] <= 9 {
				continue
			}

			flashes++
			energyLevels[i] = 0

			xy := FromIndex(i)
			for _, a := range xy.Adjacent() {
				j := a.Index()
				if energyLevels[j] != 0 {
					again = true
					energyLevels[j]++
				}
			}
		}
	}

	return flashes
}

func main() {
	energyLevels := make([]int, 0, 100)
	for _, line := range helpers.ReadStdin() {
		for _, c := range line {
			energyLevels = append(energyLevels, helpers.MustParseInt(string(c)))
		}
	}
	if len(energyLevels) != 100 {
		panic("bad input size")
	}

	if helpers.Part1() {
		flashes := 0
		for loop := 1; loop <= 100; loop++ {
			flashes += Step(energyLevels)
		}
		fmt.Printf("%d\n", flashes)
	} else {
		for loop := 1; true; loop++ {
			if Step(energyLevels) == 100 {
				fmt.Printf("%d\n", loop)
				break
			}
		}
	}
}
