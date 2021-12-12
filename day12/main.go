package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/flopp/aoc2021/helpers"
)

func isSmall(cave string) bool {
	return unicode.IsLower(rune(cave[0]))
}

func findPaths(cave string, connections map[string][]string, visits map[string]int, allowSecondVisit bool) int {
	if cave == "end" {
		return 1
	}

	paths := 0

	visits[cave]++
	for _, nextCave := range connections[cave] {
		if isSmall(nextCave) {
			switch visits[nextCave] {
			case 0:
				paths += findPaths(nextCave, connections, visits, allowSecondVisit)
			case 1:
				if allowSecondVisit && nextCave != "start" {
					paths += findPaths(nextCave, connections, visits, false)
				}
			}
		} else {
			paths += findPaths(nextCave, connections, visits, allowSecondVisit)
		}
	}
	visits[cave]--

	return paths
}

func main() {
	connections := make(map[string][]string)
	helpers.ReadStdin(func(line string) {
		a := strings.Split(line, "-")
		connections[a[0]] = append(connections[a[0]], a[1])
		connections[a[1]] = append(connections[a[1]], a[0])
	})

	visits := make(map[string]int)
	allowSecondVisit := !helpers.Part1()
	fmt.Println(findPaths("start", connections, visits, allowSecondVisit))
}
