package main

import (
	"fmt"
	"sort"

	"github.com/flopp/aoc2021/helpers"
)

type XY struct {
	x int
	y int
}

var Directions []XY = []XY{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

func (xy XY) D(d XY) XY {
	return XY{xy.x + d.x, xy.y + d.y}
}

type HeightMap struct {
	sizeX int
	sizeY int
	data  []int
}

func CreateHeightMap(sizeX int, sizeY int, data []int) *HeightMap {
	if sizeX <= 0 || sizeY <= 0 || len(data) != sizeX*sizeY {
		panic("bad bimensions")
	}
	h := HeightMap{sizeX, sizeY, data}
	return &h
}

func (h *HeightMap) Get(xy XY) int {
	if xy.x < 0 || xy.x >= h.sizeX || xy.y < 0 || xy.y >= h.sizeY {
		return 9
	}
	return h.data[xy.x+h.sizeX*xy.y]
}

func (h *HeightMap) Set(xy XY, value int) {
	if xy.x < 0 || xy.x >= h.sizeX || xy.y < 0 || xy.y >= h.sizeY {
		return
	}
	h.data[xy.x+h.sizeX*xy.y] = value
}

func flood_fill(h *HeightMap, start XY, basin int) int {
	basin_size := 0
	pending := make([]XY, 0)
	pending = append(pending, start)
	for pending_pos := 0; pending_pos < len(pending); pending_pos++ {
		xy := pending[pending_pos]
		if h.Get(xy) == basin {
			continue
		}
		h.Set(xy, basin)
		basin_size++
		for _, d := range Directions {
			if h.Get(xy.D(d)) < 9 {
				pending = append(pending, xy.D(d))
			}
		}
	}
	return basin_size
}

func main() {
	size_x := 0
	size_y := 0
	data := make([]int, 0)

	for _, line := range helpers.ReadStdin() {
		if size_y == 0 {
			size_x = len(line)
		}
		size_y++
		if len(line) != size_x {
			panic(fmt.Errorf("bad line length (%d, expected %d): %s", len(line), size_x, line))
		}

		for _, c := range line {
			data = append(data, helpers.MustParseInt(string(c)))
		}
	}
	height_map := CreateHeightMap(size_x, size_y, data)

	if helpers.Part1() {
		risk_level_sum := 0
		xy := XY{}
		for xy.x = 0; xy.x < size_x; xy.x++ {
			for xy.y = 0; xy.y < size_y; xy.y++ {
				value := height_map.Get(xy)
				ok := true
				for _, d := range Directions {
					if height_map.Get(xy.D(d)) <= value {
						ok = false
						break
					}
				}
				if ok {
					risk_level := value + 1
					risk_level_sum += risk_level
				}
			}
		}

		fmt.Printf("%d\n", risk_level_sum)
	} else {
		basin := 10
		basin_sizes := make([]int, 0)
		xy := XY{}
		for xy.x = 0; xy.x < size_x; xy.x++ {
			for xy.y = 0; xy.y < size_y; xy.y++ {
				value := height_map.Get(xy)
				if value >= 9 {
					// 9 => wall; 10... => already some other basin
					continue
				}
				basin_sizes = append(basin_sizes, flood_fill(height_map, xy, basin))
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
