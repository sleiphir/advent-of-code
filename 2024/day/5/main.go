package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed rules.txt
var _rules string

//go:embed updates.txt
var _updates string

func main() {
	rules, updates := parseInput()
	fmt.Println("Part 1:", solve1(rules, updates))
	fmt.Println("Part 2:", solve2(rules, updates))
}

func solve1(rules Rules, updates []Update) (sum int) {
	for _, update := range updates {
		if update.Check(rules) {
			middle := update[len(update)/2]
			sum += middle
		}
	}
	return
}

func solve2(rules Rules, updates []Update) (sum int) {
	for _, update := range updates {
		if !update.Check(rules) {
			update.Fix(rules)
			middle := update[len(update)/2]
			sum += middle
		}
	}
	return
}

// Map from rules.txt, the right column is used as the key
// and its corresponding left values are stored in an array
type Rules map[int][]int

// An update is a line from updates.txt
type Update []int

// Check that an update is well ordered based on the rules map
func (update Update) Check(rules Rules) bool {
	for i := range update {
		if _, success := update.CheckIndex(rules, i); !success {
			return false
		}
	}
	return true
}

// Check that the value at the given index is well placed based on the rules map
// If it is not well placed the index of the value where it failed is returned
func (update Update) CheckIndex(rules Rules, index int) (int, bool) {
	current := update[index]
	for i := index + 1; i < len(update); i++ {
		for _, rule := range rules[current] {
			if update[i] == rule {
				return i, false
			}
		}
	}
	return -1, true
}

// Continuously check and swap invalid indexes around until the update is fixed
func (update Update) Fix(rules Rules) {
	for !update.Check(rules) {
		for currIdx := range update {
			if failIdx, success := update.CheckIndex(rules, currIdx); !success {
				update[currIdx], update[failIdx] = update[failIdx], update[currIdx]
			}
		}
	}
}

func parseLine(line string) (int, int) {
	elems := strings.Split(line, "|")
	lhs, _ := strconv.Atoi(elems[0])
	rhs, _ := strconv.Atoi(elems[1])
	return lhs, rhs
}

func parseInput() (Rules, []Update) {
	var rows []string = strings.Split(_rules, "\n")
	rules := Rules{}
	for _, line := range rows[:len(rows)-1] {
		lhs, rhs := parseLine(line)
		rules[rhs] = append(rules[rhs], lhs)
	}
	updates := []Update{}
	for _, line := range strings.Split(_updates, "\n") {
		if line == "" {
			continue
		}
		update := Update{}
		for _, value := range strings.Split(line, ",") {
			v, _ := strconv.Atoi(value)
			update = append(update, v)
		}
		updates = append(updates, update)
	}
	return rules, updates
}
