package main

import (
	"bufio"
	"fmt"
	"github.com/flopp/aoc2021/helpers"
	"os"
	"strings"
)

func insert(timers *map[int]int, timer int, count int) {
	v, ok := (*timers)[timer]
	if ok {
		(*timers)[timer] = v + count
	} else {
		(*timers)[timer] = count
	}
}

func main() {
	var days int
	if helpers.Part1() {
		days = 80
	} else {
		days = 256
	}

	timers := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, timer_s := range strings.Split(line, ",") {
			timer := helpers.MustParseInt(timer_s)
			insert(&timers, timer, 1)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for day := 0; day < days; day++ {
		new_timers := make(map[int]int)
		for timer, count := range timers {
			if timer == 0 {
				insert(&new_timers, 6, count)
				insert(&new_timers, 8, count)
			} else {
				insert(&new_timers, timer-1, count)
			}
		}
		timers = new_timers
	}

	total := 0
	for _, count := range timers {
		total += count
	}
	fmt.Printf("%d\n", total)
}
