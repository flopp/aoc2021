package main

import (
	"fmt"

	"github.com/flopp/aoc2021/helpers"
)

type SnailfishNumber struct {
	n0     *SnailfishNumber
	r0     int
	n1     *SnailfishNumber
	r1     int
	parent *SnailfishNumber
}

func (n *SnailfishNumber) magnitude() int {
	m := 0
	if n.n0 != nil {
		m += 3 * n.n0.magnitude()
	} else {
		m += 3 * n.r0
	}
	if n.n1 != nil {
		m += 2 * n.n1.magnitude()
	} else {
		m += 2 * n.r1
	}
	return m
}
func (n *SnailfishNumber) findDeep(depth int) *SnailfishNumber {
	if depth == 4 {
		return n
	}

	if n.n0 != nil {
		d := n.n0.findDeep(depth + 1)
		if d != nil {
			return d
		}
	}
	if n.n1 != nil {
		d := n.n1.findDeep(depth + 1)
		if d != nil {
			return d
		}
	}

	return nil
}

func goUpRight(n *SnailfishNumber, r int) {
	if n.n1 == nil {
		n.r1 += r
	} else {
		goUpRight(n.n1, r)
	}
}

func goUpLeft(n *SnailfishNumber, r int) {
	if n.n0 == nil {
		n.r0 += r
	} else {
		goUpLeft(n.n0, r)
	}
}

func goDownUntilLeft(n *SnailfishNumber, r int) {
	if n.parent == nil {
		return
	}
	if n.parent.n0 == n {
		goDownUntilLeft(n.parent, r)
	} else {
		if n.parent.n0 == nil {
			n.parent.r0 += r
		} else {
			goUpRight(n.parent.n0, r)
		}
	}
}

func goDownUntilRight(n *SnailfishNumber, r int) {
	if n.parent == nil {
		return
	}
	if n.parent.n1 == n {
		goDownUntilRight(n.parent, r)
	} else {
		if n.parent.n1 == nil {
			n.parent.r1 += r
		} else {
			goUpLeft(n.parent.n1, r)
		}
	}
}

func (n *SnailfishNumber) explode() bool {
	d := n.findDeep(0)
	if d == nil {
		return false
	}

	if d == d.parent.n0 {
		d.parent.n0 = nil
		d.parent.r0 = 0
		goDownUntilLeft(d.parent, d.r0)
		if d.parent.n1 == nil {
			d.parent.r1 += d.r1
		} else {
			goUpLeft(d.parent.n1, d.r1)
		}
	} else if d == d.parent.n1 {
		d.parent.n1 = nil
		d.parent.r1 = 0
		if d.parent.n0 == nil {
			d.parent.r0 += d.r0
		} else {
			goUpRight(d.parent.n0, d.r0)
		}
		goDownUntilRight(d.parent, d.r1)
	}

	return true
}

func (n *SnailfishNumber) split() bool {
	if n.n0 != nil {
		if n.n0.split() {
			return true
		}
	} else {
		if n.r0 >= 10 {
			n0 := SnailfishNumber{nil, n.r0 / 2, nil, n.r0 - (n.r0 / 2), n}
			n.n0 = &n0
			return true
		}
	}

	if n.n1 != nil {
		if n.n1.split() {
			return true
		}
	} else {
		if n.r1 >= 10 {
			n1 := SnailfishNumber{nil, n.r1 / 2, nil, n.r1 - (n.r1 / 2), n}
			n.n1 = &n1
			return true
		}
	}
	return false
}

func (n *SnailfishNumber) reduce() {
	for true {
		if n.explode() {
			continue
		}
		if n.split() {
			continue
		}
		break
	}
}

func add(n0 *SnailfishNumber, n1 *SnailfishNumber) *SnailfishNumber {
	if n0 == nil {
		return n1
	}
	if n1 == nil {
		return n0
	}
	n := SnailfishNumber{n0, 0, n1, 0, nil}
	n0.parent = &n
	n1.parent = &n
	n.reduce()
	return &n
}

func (n *SnailfishNumber) copy(parent *SnailfishNumber) *SnailfishNumber {
	c := SnailfishNumber{nil, n.r0, nil, n.r1, parent}
	if n.n0 != nil {
		c.n0 = n.n0.copy(&c)
	}
	if n.n1 != nil {
		c.n1 = n.n1.copy(&c)
	}
	return &c
}

func parse(line string, pos *int, parent *SnailfishNumber) (*SnailfishNumber, int) {
	if line[*pos] == '[' {
		n := SnailfishNumber{}
		n.parent = parent
		*pos++
		n.n0, n.r0 = parse(line, pos, &n)
		*pos++
		n.n1, n.r1 = parse(line, pos, &n)
		*pos++
		return &n, 0
	}

	r := 0
	for ; /**/ '0' <= line[*pos] && line[*pos] <= '9'; *pos++ {
		r = 10*r + int(line[*pos]-'0')
	}
	return nil, r
}

func main() {
	var numbers []*SnailfishNumber
	helpers.ReadStdin(func(line string) {
		pos := 0
		n, _ := parse(line, &pos, nil)
		numbers = append(numbers, n)
	})

	if helpers.Part1() {
		sum := (*SnailfishNumber)(nil)
		for _, n := range numbers {
			sum = add(sum, n)
		}
		fmt.Println(sum.magnitude())
	} else {
		maxMagnitude := -1
		for i0, n0 := range numbers {
			for i1, n1 := range numbers {
				if i0 == i1 {
					continue
				}

				magnitude := add(n0.copy(nil), n1.copy(nil)).magnitude()
				if magnitude > maxMagnitude {
					maxMagnitude = magnitude
				}
			}
		}
		fmt.Println(maxMagnitude)
	}
}
