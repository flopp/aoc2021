package main

import (
	"fmt"

	"github.com/flopp/aoc2021/helpers"
)

type Image struct {
	background byte
	data       []string
	width      int
	height     int
}

func (image *Image) get(x int, y int) byte {
	if x < 0 || x >= image.width || y < 0 || y >= image.height {
		return image.background
	}
	return image.data[y][x]
}

func (image *Image) neighbors(x int, y int) int {
	n := 0
	for dy := -1; dy <= +1; dy++ {
		for dx := -1; dx <= +1; dx++ {
			n = n << 1
			if image.get(x+dx, y+dy) == '#' {
				n++
			}
		}
	}
	return n
}

func (image *Image) apply(image_enhancement_algorithm string) {
	data := []string{}
	for y := -1; y <= image.height; y++ {
		line := ""
		for x := -1; x <= image.width; x++ {
			line = line + string(image_enhancement_algorithm[image.neighbors(x, y)])
		}
		data = append(data, line)
	}
	image.data = data
	image.background = image_enhancement_algorithm[image.neighbors(-10, -10)]
	image.width += 2
	image.height += 2
}

func (image *Image) lit_pixels() int {
	count := 0
	for _, line := range image.data {
		for _, c := range line {
			if c == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	image := Image{'.', []string{}, 0, 0}
	var image_enhancement_algorithm = ""
	helpers.ReadStdin(func(line string) {
		if len(image_enhancement_algorithm) == 0 {
			image_enhancement_algorithm = line
		} else if len(line) != 0 {
			image.height++
			image.width = len(line)
			image.data = append(image.data, line)
		}
	})

	var steps int
	if helpers.Part1() {
		steps = 2
	} else {
		steps = 50
	}

	for step := 0; step < steps; step++ {
		image.apply(image_enhancement_algorithm)
	}

	fmt.Println(image.lit_pixels())
}
