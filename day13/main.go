package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2021/helpers"
)

type XY struct {
	x int
	y int
}

func print(coordinates []XY) {
	max_x := 0
	max_y := 0
	for _, xy := range coordinates {
		if xy.x < 0 || xy.y < 0 {
			panic(xy)
		}
		max_x = helpers.Max(max_x, xy.x)
		max_y = helpers.Max(max_y, xy.y)
	}
	max_x++
	max_y++

	grid := make([]bool, max_x*max_y)
	for _, xy := range coordinates {
		grid[xy.x+xy.y*max_x] = true
	}

	for i, b := range grid {
		if !b {
			fmt.Printf(".")
		} else {
			fmt.Printf("#")
		}
		if i%max_x == max_x-1 {
			fmt.Println()
		}
	}
}

func fold(coordinates []XY, folding XY) []XY {
	folded_coordinates := make([]XY, 0, len(coordinates))
	if folding.y == 0 {
		// fold along x
		for _, xy := range coordinates {
			if xy.x <= folding.x {
				folded_coordinates = append(folded_coordinates, xy)
			} else {
				folded_coordinates = append(folded_coordinates, XY{2*folding.x - xy.x, xy.y})
			}
		}
	} else {
		// fold along y
		for _, xy := range coordinates {
			if xy.y <= folding.y {
				folded_coordinates = append(folded_coordinates, xy)
			} else {
				folded_coordinates = append(folded_coordinates, XY{xy.x, 2*folding.y - xy.y})
			}
		}
	}

	return folded_coordinates
}

func main() {
	re_xy := regexp.MustCompile(`^(\d+),(\d+)$`)
	re_fold_x := regexp.MustCompile(`^fold along x=(\d+)$`)
	re_fold_y := regexp.MustCompile(`^fold along y=(\d+)$`)

	coordinates := []XY{}
	foldings := []XY{}

	helpers.ReadStdin(func(line string) {
		if match := re_xy.FindStringSubmatch(line); match != nil {
			coordinates = append(coordinates, XY{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2])})
		} else if match := re_fold_x.FindStringSubmatch(line); match != nil {
			foldings = append(foldings, XY{helpers.MustParseInt(match[1]), 0})
		} else if match := re_fold_y.FindStringSubmatch(line); match != nil {
			foldings = append(foldings, XY{0, helpers.MustParseInt(match[1])})
		}
	})

	if helpers.Part1() {
		folded_coordinates := fold(coordinates, foldings[0])
		count := make(map[XY]bool)
		for _, xy := range folded_coordinates {
			count[xy] = true
		}
		fmt.Println(len(count))
	} else {
		for _, folding := range foldings {
			coordinates = fold(coordinates, folding)
		}
		print(coordinates)
	}

}
