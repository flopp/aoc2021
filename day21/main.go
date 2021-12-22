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

func main() {
	players := [2]Player{}
	re_player := regexp.MustCompile(`^Player (1|2) starting position: (\d+)`)
	helpers.ReadStdin(func(line string) {
		if match := re_player.FindStringSubmatch(line); match != nil {
			players[helpers.MustParseInt(match[1])-1] = Player{helpers.MustParseInt(match[2]) - 1, 0}
		}
	})

	playerIndex := 0
	die := createDeterministicDie(100)
	for true {
		if players[playerIndex].play(&die) {
			fmt.Println(players[1-playerIndex].score * die.numberOfRolls())
			break
		}
		playerIndex = 1 - playerIndex
	}
}
