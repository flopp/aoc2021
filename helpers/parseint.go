package helpers

import (
	"strconv"
)

func MustParseInt(s string) int {
	x, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(x)
}
