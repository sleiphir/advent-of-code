package main

type Guard struct {
	Pos Pos
	Dir Dir

	initialPos Pos
	initialDir Dir

	// Used for loop detection
	Record    map[Pos]Dir
	IsLooping bool
}

func NewGuard(pos Pos, dir Dir) Guard {
	return Guard{
		Pos: pos,
		Dir: dir,
		initialPos: Pos{
			X: pos.X,
			Y: pos.Y,
		},
		initialDir: Dir{
			X: dir.X,
			Y: dir.Y,
		},
		Record: map[Pos]Dir{},
	}
}

// Reset the guard to its original position and direction
func (g *Guard) Reset() {
	g.Pos = Pos{X: g.initialPos.X, Y: g.initialPos.Y}
	g.Dir = Dir{X: g.initialDir.X, Y: g.initialDir.Y}
	g.Record = map[Pos]Dir{}
	g.IsLooping = false
}

func (g Guard) String() string {
	switch g.Dir {
	case Up:
		return "^"
	case Right:
		return ">"
	case Down:
		return "v"
	case Left:
		return "<"
	}
	return "?"
}

func (g *Guard) StepForward() {
	g.Pos.MoveToward(g.Dir)
}

func (g *Guard) TurnRight() {
	g.Dir.TurnRight()
}

// Return the next position the guard will be by stepping forward once
func (g *Guard) NextPos() Pos {
	return g.Pos.Add(Pos(g.Dir))
}

// Return the list of positions the guard went to
func (g *Guard) GetRecordedPositions() []Pos {
	pos := []Pos{}
	for key := range g.Record {
		pos = append(pos, key)
	}
	return pos
}
