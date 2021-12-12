package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2021/helpers"
)

func main() {
	part1 := helpers.Part1()

	aim := 0
	horizontal_position := 0
	depth := 0
	re := regexp.MustCompile(`^(forward|down|up) (\d+)$`)

	helpers.ReadStdin(func(line string) {
		match := re.FindStringSubmatch(line)
		if match == nil {
			panic("bad input")
		}

		value := helpers.MustParseInt(match[2])

		switch match[1] {
		case "forward":
			horizontal_position += value
			if !part1 {
				depth += value * aim
			}
		case "up":
			if part1 {
				depth -= value
			} else {
				aim -= value
			}
		case "down":
			if part1 {
				depth += value
			} else {
				aim += value
			}
		}
	})

	fmt.Println(horizontal_position * depth)
}
