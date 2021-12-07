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

	for _, line := range helpers.ReadStdin() {
		match := re.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Errorf("bad line: <%s>", line))
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
	}

	fmt.Printf("%d\n", horizontal_position*depth)
}
