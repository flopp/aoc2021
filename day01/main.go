package main

import (
	"fmt"

	"github.com/flopp/aoc2021/helpers"
)

func main() {
	var window_size int
	if helpers.Part1() {
		window_size = 1
	} else {
		window_size = 3
	}

	depths := make([]int, 0)
	for _, line := range helpers.ReadStdin() {
		depths = append(depths, helpers.MustParseInt(line))
	}

	deeper := 0
	last_window := 0
	for index, depth := range depths {
		if index < window_size {
			last_window += depth
		} else {
			window := 0
			for _, d := range depths[index+1-window_size : index+1] {
				window += d
			}
			if window > last_window {
				deeper++
			}
			last_window = window
		}
	}

	fmt.Printf("%d\n", deeper)
}
