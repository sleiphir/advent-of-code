package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	calculs := parseInput(input)
	fmt.Println("Part 1:", solve(calculs, []string{"+", "*"}))
	fmt.Println("Part 2:", solve(calculs, []string{"+", "*", "||"}))
}

func solve(calculs []Calcul, operators []string) (sum int) {
	for _, calc := range calculs {
		if canEvaluate(calc.Values[0], calc.Sum, calc.Values, operators, 1) {
			sum += calc.Sum
		}
	}
	return
}

// Recursively try all of the given operators until index reaches the length of the values array
func canEvaluate(currentValue, target int, values []int, operators []string, index int) bool {
	if index == len(values) {
		return currentValue == target
	}
	for _, op := range operators {
		switch op {
		case "+":
			currentValue += values[index]
		case "*":
			currentValue *= values[index]
		case "||":
			currentValue, _ = strconv.Atoi(fmt.Sprintf("%d%d", currentValue, values[index]))
		}
		if canEvaluate(currentValue, target, values, operators, index+1) {
			return true
		}
	}
	return false
}

func parseLine(line string) Calcul {
	arr := strings.Split(line, ": ")
	sum, _ := strconv.Atoi(arr[0])
	values := []int{}
	for _, val := range strings.Fields(arr[1]) {
		v, _ := strconv.Atoi(val)
		values = append(values, v)
	}
	return Calcul{Sum: sum, Values: values}
}

type Calcul struct {
	Sum    int
	Values []int
}

func parseInput(input string) (result []Calcul) {
	var rows []string = strings.Split(input, "\n")
	calculs := []Calcul{}
	for _, line := range rows[:len(rows)-1] {
		calcul := parseLine(line)
		calculs = append(calculs, calcul)
	}
	return calculs
}
