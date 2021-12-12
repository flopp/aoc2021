package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/flopp/aoc2021/helpers"
)

func search(node string, edges *map[string][]string, visited *map[string]int, allowSecondVisit bool) int {
	if node == "end" {
		return 1
	}

	paths := 0

	(*visited)[node]++
	for _, next := range (*edges)[node] {
		if unicode.IsLower(rune(next[0])) {
			if (*visited)[next] == 0 {
				paths += search(next, edges, visited, allowSecondVisit)
			} else if allowSecondVisit && (*visited)[next] == 1 && next != "start" {
				paths += search(next, edges, visited, false)
			}
		} else {
			paths += search(next, edges, visited, allowSecondVisit)
		}
	}
	(*visited)[node]--

	return paths
}

func main() {
	edges := make(map[string][]string)
	helpers.ReadStdin(func(line string) {
		a := strings.Split(line, "-")
		edges[a[0]] = append(edges[a[0]], a[1])
		edges[a[1]] = append(edges[a[1]], a[0])
	})

	visited := make(map[string]int)
	if helpers.Part1() {
		fmt.Println(search("start", &edges, &visited, false))
	} else {
		fmt.Println(search("start", &edges, &visited, true))
	}
}
