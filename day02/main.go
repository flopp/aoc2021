package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1 := true
	switch {
	case os.Args[1] == "part1":
		part1 = true
	case os.Args[1] == "part2":
		part1 = false
	default:
		panic(fmt.Errorf("bad part option: <%s>", os.Args[1]))
	}

	aim := int64(0)
	horizontal_position := int64(0)
	depth := int64(0)
	re := regexp.MustCompile(`^(forward|down|up) (\d+)$`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Errorf("bad line: <%s>", line))
		}

		value, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			panic(err)
		}

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
