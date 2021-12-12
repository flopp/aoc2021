package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2021/helpers"
)

func permutateArray(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func computePermutations(numbers []string) [][]string {
	permutations := make([][]string, 0)

	indexes := make([]int, 7)
	for i := range indexes {
		indexes[i] = i
	}
	for _, p := range permutateArray(indexes) {
		permuted_numbers := make([]string, 0, 10)
		for _, number := range numbers {
			permuted_number := ""
			for _, c := range number {
				i := int(c - 'a')
				permuted_number += string(rune(int('a') + p[i]))
			}
			permuted_numbers = append(permuted_numbers, helpers.SortString(permuted_number))
		}
		permutations = append(permutations, permuted_numbers)
	}

	return permutations
}

func main() {
	re_pattern := regexp.MustCompile(`\b([a-g]+)\b`)
	patterns_list := make([][]string, 0)
	helpers.ReadStdin(func(line string) {
		matches := re_pattern.FindAllStringSubmatch(line, -1)
		if matches == nil {
			panic("bad input")
		}
		patterns := make([]string, 0, len(matches))
		for _, v := range matches {
			patterns = append(patterns, helpers.SortString(v[1]))
		}
		patterns_list = append(patterns_list, patterns)
	})

	if helpers.Part1() {
		re1478 := regexp.MustCompile(`^([a-g]{2}|[a-g]{3}|[a-g]{4}|[a-g]{7})$`)
		count1478 := 0
		for _, patterns := range patterns_list {
			for _, pattern := range patterns[10:] {
				if re1478.FindStringSubmatch(pattern) != nil {
					count1478++
				}
			}
		}

		fmt.Println(count1478)
	} else {
		numbers := make([]string, 0, 10)
		numbers = append(numbers, "abcefg")
		numbers = append(numbers, "cf")
		numbers = append(numbers, "acdeg")
		numbers = append(numbers, "acdfg")
		numbers = append(numbers, "bcdf")
		numbers = append(numbers, "abdfg")
		numbers = append(numbers, "abdefg")
		numbers = append(numbers, "acf")
		numbers = append(numbers, "abcdefg")
		numbers = append(numbers, "abcdfg")

		sum := 0
		permutations := computePermutations(numbers)
		for _, patterns := range patterns_list {
			ok := false
			for _, permutation := range permutations {
				ok = true
				for _, pattern := range patterns {
					found := false
					for _, p := range permutation {
						if p == pattern {
							found = true
							break
						}
					}
					if !found {
						ok = false
						break
					}
				}
				if ok {
					value := 0
					for _, pattern := range patterns[10:] {
						for pi, p := range permutation {
							if p == pattern {
								value = (10 * value) + pi
								break
							}
						}
					}
					sum += value
					break
				}
			}
			if !ok {
				panic("no permutation found")
			}
		}

		fmt.Println(sum)
	}
}
