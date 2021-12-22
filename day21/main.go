package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2021/helpers"
)

type Die interface {
	roll() int
	numberOfRolls() int
}

type DeterministicDie struct {
	sides int
	rolls int
}

func createDeterministicDie(sides int) DeterministicDie {
	return DeterministicDie{sides, 0}
}

func (die *DeterministicDie) roll() int {
	value := (die.rolls % die.sides) + 1
	die.rolls++
	return value
}

func (die *DeterministicDie) numberOfRolls() int {
	return die.rolls
}

type Player struct {
	position int
	score    int
}

func (p *Player) play(die Die) bool {
	p.position = (p.position + die.roll() + die.roll() + die.roll()) % 10
	p.score += p.position + 1
	return p.score >= 1000
}

type GameState struct {
	next      int
	position1 int
	score1    int
	position2 int
	score2    int
}

type GameStateStack struct {
	top  int
	data []GameState
}

func createStack() *GameStateStack {
	s := GameStateStack{-1, []GameState{}}
	return &s
}

func (s *GameStateStack) IsEmpty() bool {
	return s.top == -1
}

func (s *GameStateStack) Push(v GameState) {
	if s.top+1 == len(s.data) {
		s.data = append(s.data, v)
	} else {
		s.data[s.top+1] = v
	}
	s.top++
}

func (s *GameStateStack) Pop() GameState {
	if s.top == -1 {
		panic("cannot pop from empty stack")
	}
	s.top--
	return s.data[s.top+1]
}

func (s *GameStateStack) Top() GameState {
	if s.top == -1 {
		panic("cannot get top from empty stack")
	}
	return s.data[s.top]
}

func playDirac(initialState GameState) int {
	wins := make(map[GameState]int)
	stack := createStack()
	stack.Push(initialState)
	for !stack.IsEmpty() {
		top := stack.Top()
		if _, found := wins[top]; found {
			stack.Pop()
			continue
		}
		if top.score1 >= 21 {
			wins[top] = 1
			stack.Pop()
			continue
		} else if top.score2 >= 21 {
			wins[top] = 0
			stack.Pop()
			continue
		}

		all := true
		allWins := 0
		if top.next == 0 {
			for roll1 := 1; roll1 <= 3; roll1++ {
				for roll2 := 1; roll2 <= 3; roll2++ {
					for roll3 := 1; roll3 <= 3; roll3++ {
						p := (top.position1 + roll1 + roll2 + roll3) % 10
						s := GameState{1, p, top.score1 + p + 1, top.position2, top.score2}
						if w, found := wins[s]; found {
							allWins += w
						} else {
							all = false
							stack.Push(s)
						}
					}
				}
			}
		} else {
			for roll1 := 1; roll1 <= 3; roll1++ {
				for roll2 := 1; roll2 <= 3; roll2++ {
					for roll3 := 1; roll3 <= 3; roll3++ {
						p := (top.position2 + roll1 + roll2 + roll3) % 10
						s := GameState{0, top.position1, top.score1, p, top.score2 + p + 1}
						if w, found := wins[s]; found {
							allWins += w
						} else {
							all = false
							stack.Push(s)
						}
					}
				}
			}
		}
		if all {
			wins[top] = allWins
			stack.Pop()
		}
	}

	return wins[initialState]
}

func main() {
	players := [2]Player{}
	re_player := regexp.MustCompile(`^Player (1|2) starting position: (\d+)`)
	helpers.ReadStdin(func(line string) {
		if match := re_player.FindStringSubmatch(line); match != nil {
			players[helpers.MustParseInt(match[1])-1] = Player{helpers.MustParseInt(match[2]) - 1, 0}
		}
	})
	if helpers.Part1() {
		playerIndex := 0
		die := createDeterministicDie(100)
		for true {
			if players[playerIndex].play(&die) {
				fmt.Println(players[1-playerIndex].score * die.numberOfRolls())
				break
			}
			playerIndex = 1 - playerIndex
		}
	} else {
		initialState := GameState{0, players[0].position, players[0].score, players[1].position, players[1].score}
		fmt.Println(playDirac(initialState))
	}
}
