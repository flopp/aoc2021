package helpers

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	a := strings.Split(s, "")
	sort.Strings(a)
	return strings.Join(a, "")
}
