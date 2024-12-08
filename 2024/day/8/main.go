package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	freqs, matrix := parseInput(input)
	fmt.Println("Part 1:", solve1(freqs, matrix))
	fmt.Println("Part 2:", solve2(freqs, matrix))
}

func solve1(freqs map[rune][]Vec2, matrix [][]rune) int {
	solutions := map[Vec2]bool{}
	min := Vec2{X: 0, Y: 0}
	max := Vec2{X: len(matrix[0]), Y: len(matrix)}
	for _, coords := range freqs {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				nodeA := coords[i]
				nodeB := coords[j]
				d := nodeB.Sub(nodeA)

				antinodeA := nodeA.Add(d.Mul(Vec2{X: 2, Y: 2}))
				antinodeB := nodeB.Sub(d.Mul(Vec2{X: 2, Y: 2}))

				if antinodeA.InBounds(min, max) {
					solutions[antinodeA] = true
				}
				if antinodeB.InBounds(min, max) {
					solutions[antinodeB] = true
				}
			}
		}
	}
	return len(solutions)
}

func solve2(freqs map[rune][]Vec2, matrix [][]rune) int {
	solutions := map[Vec2]bool{}
	min := Vec2{X: 0, Y: 0}
	max := Vec2{X: len(matrix[0]), Y: len(matrix)}
	for _, coords := range freqs {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				nodeA := coords[i]
				nodeB := coords[j]
				d := nodeB.Sub(nodeA)

				// Count nodes as antinodes as well
				solutions[nodeA] = true
				solutions[nodeB] = true
				k := 2
				for true {
					antinodeA := nodeA.Add(d.Mul(Vec2{X: k, Y: k}))
					antinodeB := nodeB.Sub(d.Mul(Vec2{X: k, Y: k}))
					k += 1
					aOut := !antinodeA.InBounds(min, max)
					bOut := !antinodeB.InBounds(min, max)
					if !aOut {
						solutions[antinodeA] = true
					}
					if !bOut {
						solutions[antinodeB] = true
					}
					if aOut && bOut {
						break
					}
				}
			}
		}
	}
	return len(solutions)
}

type Vec2 struct {
	X int
	Y int
}

// Check that v is contained by the vectors min and max with min inclusive & max exclusive
func (v Vec2) InBounds(min, max Vec2) bool {
	return v.X >= min.X && v.Y >= min.Y && v.X < max.X && v.Y < max.Y
}

func (v Vec2) Equal(v2 Vec2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vec2) Mul(v2 Vec2) Vec2 {
	return Vec2{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func parseInput(input string) (map[rune][]Vec2, [][]rune) {
	var rows []string = strings.Split(input, "\n")
	freqs := map[rune][]Vec2{}
	matrix := [][]rune{}
	var line string
	var c rune
	var x, y int
	for y, line = range rows[:len(rows)-1] {
		row := []rune{}
		for x, c = range line {
			row = append(row, c)
			if c == '.' {
				continue
			}
			freqs[c] = append(freqs[c], Vec2{X: x, Y: y})
		}
		matrix = append(matrix, row)
	}
	return freqs, matrix
}
