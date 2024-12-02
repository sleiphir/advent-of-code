package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	l, r := parseInput(input)
	fmt.Println("Part 1:", solve1(l, r))
	fmt.Println("Part 2:", solve2(l, r))
}

func solve1(l, r []int) int {
	sum := 0
	sort.Ints(l)
	sort.Ints(r)
	for i := range l {
		d := l[i] - r[i]
		if d < 0 {
			d *= -1
		}
		sum += d
	}
	return sum
}

func solve2(l, r []int) int {
	sum := 0
	for _, lhs := range l {
		count := 0
		for _, rhs := range r {
			if rhs == lhs {
				count += 1
			}
		}
		sum += lhs * count
	}
	return sum
}

func parseInput(input string) (l []int, r []int) {
	var cols []string = strings.Split(input, "\n")
	for _, line := range cols[:len(cols)-1] {
		lhs, rhs := parseLine(line)
		l = append(l, lhs)
		r = append(r, rhs)
	}
	return
}

func parseLine(line string) (int, int) {
	values := strings.Fields(line)
	lhs, _ := strconv.Atoi(values[0])
	rhs, _ := strconv.Atoi(values[1])
	return lhs, rhs
}
