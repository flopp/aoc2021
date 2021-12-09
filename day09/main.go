package main

import (
	"fmt"
	"sort"

	"github.com/flopp/aoc2021/helpers"
)

func flood_fill(height_map []int, size_x int, size_y int, p int, basin int) int {
	basin_size := 0
	pending := make([]int, 0)
	pending = append(pending, p)
	for pending_pos := 0; pending_pos < len(pending); pending_pos++ {
		i := pending[pending_pos]
		if height_map[p] == basin {
			continue
		}
		height_map[i] = basin
		basin_size++
		x := p % size_x
		y := p / size_x
		if x > 0 && height_map[i-1] < 9 {
			pending = append(pending, i-1)
		}
		if x+1 < size_x && height_map[i+1] < 9 {
			pending = append(pending, i+1)
		}
		if y > 0 && height_map[i-size_x] < 9 {
			pending = append(pending, i-size_x)
		}
		if y+1 < size_y && height_map[i+size_x] < 9 {
			pending = append(pending, i+size_x)
		}
	}
	return basin_size
}

func main() {
	size_x := 0
	size_y := 0
	height_map := make([]int, 0)

	for _, line := range helpers.ReadStdin() {
		if size_y == 0 {
			size_x = len(line)
		}
		size_y++
		if len(line) != size_x {
			panic(fmt.Errorf("bad line length (%d, expected %d): %s", len(line), size_x, line))
		}

		for _, c := range line {
			height_map = append(height_map, helpers.MustParseInt(string(c)))
		}
	}

	if helpers.Part1() {
		risk_level_sum := 0
		for i, value := range height_map {
			x := i % size_x
			y := i / size_x

			if x > 0 && height_map[i-1] <= value {
				continue
			}
			if x+1 < size_x && height_map[i+1] <= value {
				continue
			}
			if y > 0 && height_map[i-size_x] <= value {
				continue
			}
			if y+1 < size_y && height_map[i+size_x] <= value {
				continue
			}

			risk_level := value + 1
			risk_level_sum += risk_level
		}

		fmt.Printf("%d\n", risk_level_sum)
	} else {
		basin := 10
		basin_sizes := make([]int, 0)
		for i, value := range height_map {
			if value < 9 {
				basin_sizes = append(basin_sizes, flood_fill(height_map, size_x, size_y, i, basin))
				basin++
			}
		}
		sort.Ints(basin_sizes)
		mult := 1
		for _, basin_size := range basin_sizes[len(basin_sizes)-3:] {
			mult *= basin_size
		}

		fmt.Printf("%d\n", mult)
	}
}
