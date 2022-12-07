package main

import (
	"fmt"
	"os"
	"strings"
)

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

func scoreRoundA(round Round) int {
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

func scoreRoundB(round Round) int {
	score := 0
	switch round.me {
	case "X":
		switch round.you {
		case "A":
			score += 3
		case "B":
			score += 1
		case "C":
			score += 2
		}
	case "Y":
		score += 3
		switch round.you {
		case "A":
			score += 1
		case "B":
			score += 2
		case "C":
			score += 3
		}
	case "Z":
		score += 6
		switch round.you {
		case "A":
			score += 2
		case "B":
			score += 3
		case "C":
			score += 1
		}
	}
	return score
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")
	rounds := getRounds(lines)

	totalScoreA := 0
	for _, round := range rounds {
		totalScoreA += scoreRoundA(round)
	}
	fmt.Println("Solution to problem 1:", totalScoreA)

	totalScoreB := 0
	for _, round := range rounds {
		totalScoreB += scoreRoundB(round)
	}
	fmt.Println("Solution to problem 1:", totalScoreB)
}
