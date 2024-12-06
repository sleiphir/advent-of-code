package main

type Dir struct {
	X int
	Y int
}

var (
	Up    Dir = Dir{X: 0, Y: -1}
	Down  Dir = Dir{X: 0, Y: 1}
	Left  Dir = Dir{X: -1, Y: 0}
	Right Dir = Dir{X: 1, Y: 0}
)

func (dir *Dir) TurnRight() {
	newDir := Dir{}
	switch *dir {
	case Up:
		newDir = Right
	case Right:
		newDir = Down
	case Down:
		newDir = Left
	case Left:
		newDir = Up
	}
	dir.X = newDir.X
	dir.Y = newDir.Y
}
