package main

import (
	"fmt"
	"sort"

	"github.com/flopp/aoc2021/helpers"
)

type stack struct {
	top  int
	data []rune
}

func CreateStack() *stack {
	s := stack{-1, []rune{}}
	return &s
}

func (s *stack) IsEmpty() bool {
	return s.top == -1
}

func (s *stack) Push(v rune) {
	if s.top+1 == len(s.data) {
		s.data = append(s.data, v)
	} else {
		s.data[s.top+1] = v
	}
	s.top++
}

func (s *stack) Pop() (bool, rune) {
	if s.top == -1 {
		return false, rune(0)
	}
	s.top--
	return true, s.data[s.top+1]
}

func is_corrupted_or_incomplete(line string) (int, int) {
	s := CreateStack()
	for _, c := range line {
		switch c {
		case '(':
			s.Push(c)
		case '[':
			s.Push(c)
		case '{':
			s.Push(c)
		case '<':
			s.Push(c)
		case ')':
			if ok, r := s.Pop(); !ok || r != '(' {
				return 3, 0
			}
		case ']':
			if ok, r := s.Pop(); !ok || r != '[' {
				return 57, 0
			}
		case '}':
			if ok, r := s.Pop(); !ok || r != '{' {
				return 1197, 0
			}
		case '>':
			if ok, r := s.Pop(); !ok || r != '<' {
				return 25137, 0
			}
		default:
			panic(fmt.Errorf("bad character '%s' in line '%s'", string(c), line))
		}
	}

	completion_score := 0
	for !s.IsEmpty() {
		_, c := s.Pop()
		switch c {
		case '(':
			completion_score = 5*completion_score + 1
		case '[':
			completion_score = 5*completion_score + 2
		case '{':
			completion_score = 5*completion_score + 3
		case '<':
			completion_score = 5*completion_score + 4
		}
	}
	return 0, completion_score
}

func main() {
	total_syntax_error_score := 0
	completion_scores := make([]int, 0)
	for _, line := range helpers.ReadStdin() {
		corruption_score, completion_score := is_corrupted_or_incomplete(line)
		if corruption_score != 0 {
			total_syntax_error_score += corruption_score
		} else if completion_score != 0 {
			completion_scores = append(completion_scores, completion_score)
		} else {
			panic(fmt.Errorf("line is valid: %s", line))
		}
	}
	if helpers.Part1() {
		fmt.Printf("%d\n", total_syntax_error_score)
	} else {
		sort.Ints((completion_scores))
		middle_score := completion_scores[len(completion_scores)/2]
		fmt.Printf("%d\n", middle_score)
	}
}
