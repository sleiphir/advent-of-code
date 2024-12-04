package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input := parseInput(input)
	fmt.Println("Part 1:", solve1(input))
	fmt.Println("Part 2:", solve2(input))
}

func solve1(matrix []string) (sum int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'X' {
				sum += checkPos(i, j, matrix)
			}
		}
	}
	return
}

func solve2(matrix []string) (sum int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'A' {
				if checkCross(i, j, matrix) {
					sum += 1
				}
			}
		}
	}
	return
}

func inBounds(i, j int, matrix []string) bool {
	return i >= 0 && i < len(matrix[0]) && j >= 0 && j < len(matrix)
}

func checkPos(i, j int, matrix []string) (sum int) {
	target := "MAS"
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, dir := range dirs {
		if checkDir(i, j, dir[0], dir[1], matrix, target) {
			sum += 1
		}
	}
	return
}

func checkCross(i, j int, matrix []string) bool {
	if inBounds(i-1, j-1, matrix) && inBounds(i+1, j+1, matrix) {
		a := fmt.Sprintf("%c%c%c", matrix[i-1][j-1], matrix[i][j], matrix[i+1][j+1])
		b := fmt.Sprintf("%c%c%c", matrix[i-1][j+1], matrix[i][j], matrix[i+1][j-1])
		return (a == "MAS" || a == "SAM") && (b == "MAS" || b == "SAM")
	}
	return false
}

func checkDir(i, j int, dir_x, dir_y int, matrix []string, target string) bool {
	pos := 0
	start_x := j
	start_y := i
	x := start_x
	y := start_y
	for k := 1; k <= len(target); k++ {
		x = start_x + (dir_x * k)
		y = start_y + (dir_y * k)
		if !inBounds(y, x, matrix) || matrix[y][x] != target[pos] {
			return false
		}
		pos += 1
		if pos == len(target) {
			return true
		}
	}
	return false
}

func parseInput(input string) (result []string) {
	result = strings.Split(input, "\n")
	return result[:len(result)-1]
}
