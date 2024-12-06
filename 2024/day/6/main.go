package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed _input.txt
var input string

func main() {
	lab := parseInput(input)
	fmt.Println("Part 1:", solve1(lab))
	lab.Reset()
	fmt.Println("Part 2:", solve2(lab))
}

func solve1(lab Lab) int {
	// Guard the lab until the guard exit the board
	for lab.MoveGuard() {
		// fmt.Println(lab)
	}
	return lab.CountVisited()
}

// A little slow but gets the job done
func solve2(lab Lab) (sum int) {
	// Start by running it once to get the recorded positions
	for lab.MoveGuard() {
		// fmt.Println(lab)
	}
	// Try to create a loop by putting a wall at each recorded positions
	for _, pos := range lab.Guard.GetRecordedPositions() {
		lab.Reset()
		if lab.Board.Set(pos, Wall) {
			for lab.MoveGuard() && !lab.Guard.IsLooping {
				// fmt.Println(lab)
			}
			if lab.Guard.IsLooping {
				// Loop found
				sum += 1
			}
		}
	}
	return
}

func parseInput(input string) Lab {
	var rows []string = strings.Split(input, "\n")
	guardPos := findGuard(rows[:len(rows)-1])
	cells := [][]Cell{}
	for _, line := range rows[:len(rows)-1] {
		cells = append(cells, parseLine(line))
	}
	return NewLab(NewBoard(cells), NewGuard(guardPos, Up))
}

// Locate the guard within the input file
func findGuard(rows []string) Pos {
	for y, line := range rows[:len(rows)-1] {
		for x, cell := range line {
			if cell == '^' {
				return Pos{X: x, Y: y}
			}
		}
	}
	return Pos{}
}

func parseLine(line string) (row []Cell) {
	for _, c := range line {
		if string(c) == "." {
			row = append(row, Empty)
		} else if string(c) == "#" {
			row = append(row, Wall)
		} else if c == '^' {
			row = append(row, Visited)
		}
	}
	return
}
