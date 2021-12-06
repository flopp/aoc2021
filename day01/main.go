package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	window_size, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	depths := make([]int64, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		depth, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		depths = append(depths, depth)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	deeper := 0
	last_window := int64(0)
	for index, depth := range depths {
		if int64(index) < window_size {
			last_window += depth
		} else {
			window := int64(0)
			for _, d := range depths[int64(index)+1-window_size : int64(index)+1] {
				window += d
			}
			if window > last_window {
				deeper += 1
			}
			last_window = window
		}
	}

	fmt.Printf("%d\n", deeper)
}
