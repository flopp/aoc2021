package main

import (
	"regexp"

	"github.com/flopp/aoc2021/helpers"
)

type XYZ struct {
	x int
	y int
	z int
}

func (p XYZ) rotations() []XYZ {
	r := make([]XYZ, 0, 24)
	r = append(r, XYZ{p.x, p.y, p.z})
	r = append(r, XYZ{p.x, -p.y, -p.z})
	r = append(r, XYZ{-p.x, p.y, -p.z})
	r = append(r, XYZ{-p.x, -p.y, p.z})
	r = append(r, XYZ{p.x, p.z, -p.y})
	r = append(r, XYZ{p.x, -p.z, p.y})
	r = append(r, XYZ{-p.x, p.z, p.y})
	r = append(r, XYZ{-p.x, -p.z, -p.y})
	r = append(r, XYZ{p.y, p.z, p.x})
	r = append(r, XYZ{p.y, -p.z, -p.x})
	r = append(r, XYZ{-p.y, p.z, -p.x})
	r = append(r, XYZ{-p.y, -p.z, p.x})
	r = append(r, XYZ{p.y, p.x, -p.z})
	r = append(r, XYZ{p.y, -p.x, p.z})
	r = append(r, XYZ{-p.y, p.x, p.z})
	r = append(r, XYZ{-p.y, -p.x, -p.z})
	r = append(r, XYZ{p.z, p.x, p.y})
	r = append(r, XYZ{p.z, -p.x, -p.y})
	r = append(r, XYZ{-p.z, p.x, -p.y})
	r = append(r, XYZ{-p.z, -p.x, p.y})
	r = append(r, XYZ{p.z, p.y, -p.x})
	r = append(r, XYZ{p.z, -p.y, p.x})
	r = append(r, XYZ{-p.z, p.y, p.x})
	r = append(r, XYZ{-p.z, -p.y, -p.x})
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p0, p1 XYZ) int {
	return abs(p0.x-p1.x) + abs(p0.y-p1.y) + abs(p0.z-p1.z)
}

type Scanner struct {
	beacons []XYZ
}

func main() {
	re_scanner := regexp.MustCompile(`^--- scanner (\d+) ---$`)
	re_xyz := regexp.MustCompile(`^(-?\d+),(-?\d+),(-?\d+)$`)
	scanner := (*Scanner)(nil)
	scanners := []Scanner{}
	helpers.ReadStdin(func(line string) {
		if match := re_scanner.FindStringSubmatch(line); match != nil {
			scanners = append(scanners, Scanner{})
			scanner = &(scanners[len(scanners)-1])
		} else if match := re_xyz.FindStringSubmatch(line); match != nil {
			scanner.beacons = append(scanner.beacons, p.xYZ{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2]), helpers.MustParseInt(match[3])})
		}
	})

}
