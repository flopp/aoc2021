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

func (s *stack) Pop() rune {
	if s.top == -1 {
		panic("cannot pop from empty stack")
	}
	s.top--
	return s.data[s.top+1]
}

var open_brace_map map[rune]rune = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var corruption_score_map map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completion_score_map map[rune]int = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func is_corrupted_or_incomplete(line string) (int, int) {
	s := CreateStack()
	for _, c := range line {
		switch {
		case c == '(' || c == '[' || c == '{' || c == '<':
			s.Push(c)
		case c == ')' || c == ']' || c == '}' || c == '>':
			if s.Pop() != open_brace_map[c] {
				return corruption_score_map[c], 0
			}
		default:
			panic("bad input")
		}
	}

	completion_score := 0
	for !s.IsEmpty() {
		c := s.Pop()
		completion_score = 5*completion_score + completion_score_map[c]
	}
	return 0, completion_score
}

func main() {
	total_syntax_error_score := 0
	completion_scores := make([]int, 0)
	helpers.ReadStdin(func(line string) {
		corruption_score, completion_score := is_corrupted_or_incomplete(line)
		if corruption_score != 0 {
			total_syntax_error_score += corruption_score
		} else if completion_score != 0 {
			completion_scores = append(completion_scores, completion_score)
		}
	})
	if helpers.Part1() {
		fmt.Println(total_syntax_error_score)
	} else {
		sort.Ints((completion_scores))
		middle_score := completion_scores[len(completion_scores)/2]
		fmt.Println(middle_score)
	}
}
