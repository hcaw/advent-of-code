package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items   []int
	opArith string
	operand string
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

	inspections := make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[j] += 1
				num, err := strconv.Atoi(monkey.operand)
				if monkey.opArith == "+" {
					if err != nil {
						item += item
					} else {
						item += num
					}
				} else {
					if err != nil {
						item *= item
					} else {
						item *= num
					}
				}
				item = int(item / 3)
				if item%monkey.testDiv == 0 {
					monkeys[monkey.testT].items = append(monkeys[monkey.testT].items, item)
				} else {
					monkeys[monkey.testF].items = append(monkeys[monkey.testF].items, item)
				}
			}
			monkeys[j].items = []int{}
		}
	}
	sort.Ints(inspections)
	monkeyBiz := inspections[len(inspections)-2] * inspections[len(inspections)-1]
	fmt.Println("Solution to problem 1 is", monkeyBiz)
}
