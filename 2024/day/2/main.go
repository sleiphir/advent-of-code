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
	matrix := parseInput(input)
	fmt.Println("part 1:", solve1(matrix))
	fmt.Println("part 2:", solve2(matrix))
}

func solve1(matrix [][]int) (sum int) {
	for i := range matrix {
		if checkRow(matrix[i]) {
			sum += 1
		}
	}
	return
}

func solve2(matrix [][]int) (sum int) {
	for i := range matrix {
		safe := checkRow(matrix[i])
		if !safe {
			for j := range matrix[i] {
				if checkRow(removeIndex(matrix[i], j)) {
					safe = true
					break
				}
			}
		}
		if safe {
			sum += 1
		}
	}
	return
}

func removeIndex(arr []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}

type Dir int

const (
	Nil Dir = iota
	Asc
	Desc
	Equal
)

func GetDir(a int, b int) (dir Dir) {
	switch {
	case a == b:
		dir = Equal
	case a > b:
		dir = Desc
	case a < b:
		dir = Asc
	}
	return
}

// returns whether the row is correct
func checkRow(row []int) bool {
	lastDir := Nil
	for i := 1; i < len(row); i++ {
		currDir := GetDir(row[i], row[i-1])
		d := dist(row[i], row[i-1])
		if d == 0 || d > 3 || (lastDir != Nil && currDir != lastDir) {
			return false
		}
		lastDir = currDir
	}
	return true
}

func dist(a int, b int) int {
	r := a - b
	if r < 0 {
		r *= -1
	}
	return r
}

func parseInput(input string) (matrix [][]int) {
	var lines []string = strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		matrix = append(matrix, parseLine(line))
	}
	return
}

func parseLine(line string) (row []int) {
	for _, el := range strings.Fields(line) {
		value, _ := strconv.Atoi(el)
		row = append(row, value)
	}
	return
}
