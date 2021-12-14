package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2021/helpers"
)

type RuleKey struct {
	start byte
	end   byte
}

type CacheKey struct {
	start byte
	end   byte
	steps int
}

type Counts map[byte]int

func recurse(start byte, end byte, steps int, rules map[RuleKey]byte, cache map[CacheKey]Counts) Counts {
	key := CacheKey{start, end, steps}
	if cached, found := cache[key]; found {
		return cached
	}

	counts := make(Counts)
	ruleKey := RuleKey{start, end}
	if insertion, found := rules[ruleKey]; found {
		counts[insertion]++
		for k, v := range recurse(start, insertion, steps-1, rules, cache) {
			counts[k] += v
		}
		for k, v := range recurse(insertion, end, steps-1, rules, cache) {
			counts[k] += v
		}
	}

	cache[key] = counts

	return counts
}

func main() {
	re_template := regexp.MustCompile(`^([A-Z]+)$`)
	re_rule := regexp.MustCompile(`^([A-Z])([A-Z])\s+->\s+([A-Z])$`)

	template := ""
	rules := make(map[RuleKey]byte)

	helpers.ReadStdin(func(line string) {
		if match := re_template.FindStringSubmatch(line); match != nil {
			template = match[1]
		} else if match := re_rule.FindStringSubmatch(line); match != nil {
			rules[RuleKey{match[1][0], match[2][0]}] = match[3][0]
		}
	})

	var steps int
	if helpers.Part1() {
		steps = 10
	} else {
		steps = 40
	}

	cache := make(map[CacheKey]Counts)
	for k, v := range rules {
		counts := make(Counts)
		counts[v] = 1
		cache[CacheKey{k.start, k.end, 1}] = counts
	}

	counts := make(Counts)
	counts[template[0]]++
	for i := 1; i < len(template); i++ {
		counts[template[i]]++
		for k, v := range recurse(template[i-1], template[i], steps, rules, cache) {
			counts[k] += v
		}
	}

	min := -1
	max := -1
	for _, value := range counts {
		if min < 0 {
			min = value
			max = value
		}
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Println(max - min)
}
