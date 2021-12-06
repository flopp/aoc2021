package helpers

import (
	"fmt"
	"os"
)

func Part1() bool {
	switch {
	case os.Args[1] == "part1":
		return true
	case os.Args[1] == "part2":
		return false
	default:
		panic(fmt.Errorf("bad part option: <%s>", os.Args[1]))
	}
}
