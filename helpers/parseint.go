package helpers

import (
	"strconv"
)

func MustParseInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(x)
}
