package main

import (
	"bufio"
	"fmt"
	"github.com/flopp/aoc2021/helpers"
	"os"
	"regexp"
)

func main() {
	part1 := helpers.Part1()

	aim := 0
	horizontal_position := 0
	depth := 0
	re := regexp.MustCompile(`^(forward|down|up) (\d+)$`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Errorf("bad line: <%s>", line))
		}

		value := helpers.MustParseInt(match[2])

		switch {
		case match[1] == "forward":
			horizontal_position += value
			if !part1 {
				depth += value * aim
			}
		case match[1] == "up":
			if part1 {
				depth -= value
			} else {
				aim -= value
			}
		case match[1] == "down":
			if part1 {
				depth += value
			} else {
				aim += value
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", horizontal_position*depth)
}
