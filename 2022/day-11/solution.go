package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items   []int
	opArith string
	opNum   int
	testDiv int
	testT   int
	testF   int
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	monkeyGroups := strings.Split(string(input), "\n\n")

	// initialise monkey data
	monkeys := make([]Monkey, len(monkeyGroups))
	for i, monkey := range monkeyGroups {
		lines := strings.Split(string(monkey), "\n")
		itemStr := strings.Split(strings.Split(lines[1], ":")[1], ",")
		monkey := Monkey{}
		items := []int{}
		for _, s := range itemStr {
			s = strings.TrimSpace(s)
			num, _ := strconv.Atoi(s)
			items = append(items, num)
		}
		monkey.items = items
		monkeys[i] = monkey
	}
	fmt.Println(monkeys)

	// loop 20 rounds
	for i := 0; i < 20; i++ {

	}
}
