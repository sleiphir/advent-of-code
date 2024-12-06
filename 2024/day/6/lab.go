package main

import "fmt"

type Lab struct {
	Board Board
	Guard Guard
}

// Print the lab's board with the guard
func (lab Lab) String() string {
	str := ""
	for y, row := range lab.Board.GetCells() {
		for x, cell := range row {
			if lab.Guard.Pos.Equal(x, y) {
				str += fmt.Sprint(lab.Guard)
			} else {
				str += string(cell)
			}
		}
		str += "\n"
	}
	return str
}

func NewLab(board Board, guard Guard) Lab {
	return Lab{
		Guard: guard,
		Board: board,
	}
}

func (lab *Lab) Reset() {
	lab.Board.Reset()
	lab.Guard.Reset()
}

// Count the number of cells marked as visited (X)
func (lab Lab) CountVisited() (sum int) {
	for _, row := range lab.Board.GetCells() {
		for _, cell := range row {
			if cell == Visited {
				sum += 1
			}
		}
	}
	return
}

// Make the guard walk one cell towards the direction they are currently facing.
// If a wall obstruct the guard's path they will turn right.
// Returns true if the guard is still inside the lab, false otherwise.
func (lab *Lab) MoveGuard() bool {
	nextPos := lab.Guard.NextPos()
	nextCell := lab.Board.At(nextPos)
	// Reached the outside of the board
	if nextCell == nil {
		return false
	}
	switch *nextCell {
	case Wall:
		// Guard faces a wall
		lab.Guard.TurnRight()
	case Visited:
		if lab.Guard.Record[nextPos] == lab.Guard.Dir {
			// Guard has already been on this cell pointing in the same direction
			lab.Guard.IsLooping = true
		}
		fallthrough
	default:
		lab.Guard.StepForward()
		lab.Board.Set(lab.Guard.Pos, Visited)
		// Record the direction the guard is pointing to based on its position
		lab.Guard.Record[lab.Guard.Pos] = lab.Guard.Dir
	}
	return true
}
