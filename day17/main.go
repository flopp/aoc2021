package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2021/helpers"
)

type IntRange struct {
	min int
	max int
}

func (r IntRange) inRange(v int) bool {
	return r.min <= v && v <= r.max
}

type XY struct {
	x int
	y int
}

func cannotReachTarget(p XY, v XY, targetX IntRange, targetY IntRange) bool {
	if p.y < targetY.min && v.y <= 0 {
		return true
	}
	if p.x < targetX.min && v.x <= 0 {
		return true
	}
	if p.x > targetX.max && v.x >= 0 {
		return true
	}
	return false
}

func main() {
	var target_x IntRange
	var target_y IntRange
	re_target := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	helpers.ReadStdin(func(line string) {
		if match := re_target.FindStringSubmatch(line); match != nil {
			target_x = IntRange{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2])}
			target_y = IntRange{helpers.MustParseInt(match[3]), helpers.MustParseInt(match[4])}
		}
	})

	totalMaxY := 0
	hits := 0
	// brute force loops :(
	for vy := -1000; vy <= 1000; vy++ {
		for vx := 1; vx <= target_x.max+1; vx++ {
			maxY := 0
			targetReached := false
			p := XY{0, 0}
			v := XY{vx, vy}
			for !cannotReachTarget(p, v, target_x, target_y) {
				// check if in target
				if target_x.inRange(p.x) && target_y.inRange(p.y) {
					targetReached = true
				}

				// update max y
				if p.y > maxY {
					maxY = p.y
				}

				// advance
				p.x += v.x
				p.y += v.y
				if v.x > 0 {
					v.x--
				} else if v.x < 0 {
					v.x++
				}
				v.y--
			}
			if targetReached {
				hits++
				if maxY > totalMaxY {
					totalMaxY = maxY
				}
			}
		}
	}

	if helpers.Part1() {
		fmt.Println(totalMaxY)
	} else {
		fmt.Println(hits)
	}
}
