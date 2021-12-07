package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/flopp/aoc2021/helpers"
)

func main() {
	part1 := helpers.Part1()

	crabs := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, crab_s := range strings.Split(line, ",") {
			crabs = append(crabs, helpers.MustParseInt(crab_s))
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	min := crabs[0]
	max := min
	for _, crab := range crabs {
		min = helpers.Min(min, crab)
		max = helpers.Max(max, crab)
	}

	minFuel := -1
	for pos := min; pos <= max; pos++ {
		fuel := 0
		for _, crab := range crabs {
			distance := helpers.Abs(crab - pos)
			if part1 {
				fuel += distance
			} else {
				fuel += (distance * (distance + 1)) / 2
			}
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Printf("%d\n", minFuel)
}
