package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/flopp/aoc2021/helpers"
)

type Range struct {
	min int
	max int
}

func (r Range) size() int {
	return r.max - r.min
}

func (r Range) contains(s Range) bool {
	return r.min <= s.min && s.max <= r.max
}

func (r Range) split(values []int) []Range {
	results := make([]Range, 0, 1)
	min := r.min
	for _, v := range values {
		if v <= min {
			continue
		}
		if v >= r.max {
			results = append(results, Range{min, r.max})
			break
		}
		results = append(results, Range{min, v})
		min = v
	}
	if min < r.max {
		results = append(results, Range{min, r.max})
	}
	return results
}

type Range3D struct {
	x Range
	y Range
	z Range
}

func (r Range3D) size() int {
	return r.x.size() * r.y.size() * r.z.size()
}

func (r Range3D) contains(s Range3D) bool {
	return r.x.contains(s.x) && r.y.contains(s.y) && r.z.contains(s.z)
}

func (r Range3D) splitX(values []int) []Range3D {
	results := make([]Range3D, 0, 1)
	for _, x := range r.x.split(values) {
		results = append(results, Range3D{x, r.y, r.z})
	}
	return results
}
func (r Range3D) splitY(values []int) []Range3D {
	results := make([]Range3D, 0, 1)
	for _, y := range r.y.split(values) {
		results = append(results, Range3D{r.x, y, r.z})
	}
	return results
}
func (r Range3D) splitZ(values []int) []Range3D {
	results := make([]Range3D, 0, 1)
	for _, z := range r.z.split(values) {
		results = append(results, Range3D{r.x, r.y, z})
	}
	return results
}

type RebootStep struct {
	on bool
	r  Range3D
}

func sortUnique(values []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range values {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	re_on := regexp.MustCompile(`^on x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)
	re_off := regexp.MustCompile(`^off x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

	r50 := Range3D{Range{-50, 51}, Range{-50, 51}, Range{-50, 51}}

	rebootSteps := make([]*RebootStep, 0)
	helpers.ReadStdin(func(line string) {
		if match := re_on.FindStringSubmatch(line); match != nil {
			r := Range3D{Range{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2]) + 1}, Range{helpers.MustParseInt(match[3]), helpers.MustParseInt(match[4]) + 1}, Range{helpers.MustParseInt(match[5]), helpers.MustParseInt(match[6]) + 1}}
			if !helpers.Part1() || r50.contains(r) {
				s := RebootStep{true, r}
				rebootSteps = append(rebootSteps, &s)
			}
		} else if match := re_off.FindStringSubmatch(line); match != nil {
			r := Range3D{Range{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2]) + 1}, Range{helpers.MustParseInt(match[3]), helpers.MustParseInt(match[4]) + 1}, Range{helpers.MustParseInt(match[5]), helpers.MustParseInt(match[6]) + 1}}
			if !helpers.Part1() || r50.contains(r) {
				s := RebootStep{false, r}
				rebootSteps = append(rebootSteps, &s)
			}
		}
	})

	xPlanes := make([]int, 0)
	yPlanes := make([]int, 0)
	zPlanes := make([]int, 0)
	for _, s := range rebootSteps {
		xPlanes = append(xPlanes, s.r.x.min, s.r.x.max)
		yPlanes = append(yPlanes, s.r.y.min, s.r.y.max)
		zPlanes = append(zPlanes, s.r.z.min, s.r.z.max)
	}
	xPlanes = sortUnique(xPlanes)
	yPlanes = sortUnique(yPlanes)
	zPlanes = sortUnique(zPlanes)
	fmt.Printf("regions = %d\n", len(xPlanes)*len(yPlanes)*len(zPlanes))

	newRebootSteps := make([]*RebootStep, 0, len(rebootSteps))
	for _, s := range rebootSteps {
		if s.r.x.max <= xPlanes[0] || xPlanes[len(xPlanes)-1] <= s.r.x.min {
			newRebootSteps = append(newRebootSteps, s)
		} else {
			for _, r := range s.r.splitX(xPlanes) {
				s2 := RebootStep{s.on, r}
				newRebootSteps = append(newRebootSteps, &s2)
			}
		}
	}
	fmt.Printf("split steps: %d\n", len(newRebootSteps))

	rebootSteps = rebootSteps[:0]
	for _, s := range newRebootSteps {
		if s.r.y.max <= yPlanes[0] || yPlanes[len(yPlanes)-1] <= s.r.y.min {
			rebootSteps = append(rebootSteps, s)
		} else {
			for _, r := range s.r.splitY(yPlanes) {
				s2 := RebootStep{s.on, r}
				rebootSteps = append(rebootSteps, &s2)
			}
		}
	}
	fmt.Printf("split steps: %d\n", len(rebootSteps))

	onRanges := make(map[Range3D]bool)
	for _, s := range rebootSteps {
		if s.r.z.max <= zPlanes[0] || zPlanes[len(zPlanes)-1] <= s.r.z.min {
			onRanges[s.r] = s.on
		} else {
			for _, r := range s.r.splitZ(zPlanes) {
				onRanges[r] = s.on
			}
		}
	}

	count := 0
	for r, isOn := range onRanges {
		if isOn {
			count += r.size()
		}
	}
	fmt.Println(count)
}
