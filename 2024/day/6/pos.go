package main

type Pos struct {
	X int
	Y int
}

func (p Pos) Equal(x, y int) bool {
	return p.X == x && p.Y == y
}

func (p Pos) Add(p2 Pos) Pos {
	return Pos{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

func (p *Pos) MoveToward(d Dir) {
	p.X += d.X
	p.Y += d.Y
}
