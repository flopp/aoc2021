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

	rebootSteps := make([]RebootStep, 0)
	helpers.ReadStdin(func(line string) {
		if match := re_on.FindStringSubmatch(line); match != nil {
			r := Range3D{Range{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2]) + 1}, Range{helpers.MustParseInt(match[3]), helpers.MustParseInt(match[4]) + 1}, Range{helpers.MustParseInt(match[5]), helpers.MustParseInt(match[6]) + 1}}
			if !helpers.Part1() || r50.contains(r) {
				rebootSteps = append(rebootSteps, RebootStep{true, r})
			}
		} else if match := re_off.FindStringSubmatch(line); match != nil {
			r := Range3D{Range{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2]) + 1}, Range{helpers.MustParseInt(match[3]), helpers.MustParseInt(match[4]) + 1}, Range{helpers.MustParseInt(match[5]), helpers.MustParseInt(match[6]) + 1}}
			if !helpers.Part1() || r50.contains(r) {
				rebootSteps = append(rebootSteps, RebootStep{false, r})
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
	count := 0
	r := Range3D{}
	for ix, x := range xPlanes {
		if ix == 0 {
			r.x.max = x
			continue
		}
		r.x.min = r.x.max
		r.x.max = x
		for iy, y := range yPlanes {
			if iy == 0 {
				r.y.max = y
				continue
			}
			r.y.min = r.y.max
			r.y.max = y
			for iz, z := range zPlanes {
				if iz == 0 {
					r.z.max = z
					continue
				}
				r.z.min = r.z.max
				r.z.max = z
				on := false
				for _, s := range rebootSteps {
					if s.r.contains(r) {
						on = s.on
					}
				}
				if on {
					count += r.size()
				}
			}
		}
	}

	fmt.Println(count)
}
