package main

import (
	"fmt"
	"os"
	"strings"
)

// A = Rock, B = Paper, C = Scissors
// X = Rock, Y = Paper, Z = Scissors

// score for the shape you selected:
// 1 for Rock, 2 for Paper, 3 for Scissors
// score for the outcome of the round:
// 0 for lose, 3 for draw, 6 for win

type Round struct {
	me  string
	you string
}

func getRounds(lines []string) []Round {
	var rounds []Round
	for _, line := range lines {
		first := string(line[0])
		second := string(line[2])
		round := Round{second, first}
		rounds = append(rounds, round)
	}
	return rounds
}

func scoreRound(round Round) int {
	score := 0
	switch round.me {
	case "X":
		score += 1
		switch round.you {
		case "A":
			score += 3
		case "B":
		case "C":
			score += 6
		}
	case "Y":
		score += 2
		switch round.you {
		case "A":
			score += 6
		case "B":
			score += 3
		case "C":
		}
	case "Z":
		score += 3
		switch round.you {
		case "A":
		case "B":
			score += 6
		case "C":
			score += 3
		}
	}
	return score
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")
	rounds := getRounds(lines)

	totalScore := 0
	for _, round := range rounds {
		totalScore += scoreRound(round)
	}
	fmt.Println(totalScore)
}
