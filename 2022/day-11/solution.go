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
	operand   string
	testDiv int
	testT   int
	testF   int
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	monkeyGroups := strings.Split(string(input), "\n\n")

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
		monkey.opArith = strings.Fields(lines[2])[4]
		monkey.operand = strings.Fields(lines[2])[5]
		testDiv, _ := strconv.Atoi(strings.Fields(lines[3])[3])
		monkey.testDiv = testDiv
		testT, _ := strconv.Atoi(strings.Fields(lines[4])[5])
		testF, _ := strconv.Atoi(strings.Fields(lines[5])[5])
		monkey.testT = testT
		monkey.testF = testF
		monkeys[i] = monkey
	}
	fmt.Println(monkeys)

	// loop 20 rounds
	for i := 0; i < 20; i++ {

	}
}
