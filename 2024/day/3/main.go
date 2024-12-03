package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", solve1(parseInput(input, regexp.MustCompile(`mul\(\d+,\d+\)`))))
	fmt.Println("Part 2:", solve2(parseInput(input, regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`))))
}

func solve1(matches []string) (sum int) {
	for _, match := range matches {
		numbers := strings.Split(strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")"), ",")
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])
		sum += a * b
	}
	return
}

func solve2(matches []string) (sum int) {
	shouldCount := true
	for _, match := range matches {
		switch match {
		case "do()":
			shouldCount = true
		case "don't()":
			shouldCount = false
		default:
			if shouldCount {
				numbers := strings.Split(strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")"), ",")
				a, _ := strconv.Atoi(numbers[0])
				b, _ := strconv.Atoi(numbers[1])
				sum += a * b
			}
		}
	}
	return
}

func parseInput(input string, regex *regexp.Regexp) (result []string) {
	var rows []string = strings.Split(input, "\n")
	content := strings.Join(rows, "")
	result = regex.FindAllString(content, -1)
	return
}
