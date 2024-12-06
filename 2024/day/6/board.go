package main

type Board struct {
	board [][]Cell

	initialBoard [][]Cell
}

func NewBoard(board [][]Cell) Board {
	b := Board{}
	for y := range board {
		row := []Cell{}
		initialRow := []Cell{}
		for x := range board[y] {
			row = append(row, board[y][x])
			initialRow = append(initialRow, board[y][x])
		}
		b.board = append(b.board, row)
		b.initialBoard = append(b.initialBoard, initialRow)
	}
	return b
}

func (b *Board) Reset() {
	for y := range b.board {
		for x := range b.board[y] {
			b.board[y][x] = b.initialBoard[y][x]
		}
	}
}

func (b *Board) GetCells() [][]Cell {
	return b.board
}

// Return the cell at the given position or nothing
func (b *Board) At(p Pos) *Cell {
	if !b.InBounds(p) {
		return nil
	}
	return &b.board[p.Y][p.X]
}

// Set the cell at the given position to c or return false
func (b *Board) Set(p Pos, c Cell) bool {
	if b.InBounds(p) {
		b.board[p.Y][p.X] = c
		return true
	}
	return false
}

// Check wether the given position is inside the board
func (b *Board) InBounds(p Pos) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < len(b.board[0]) && p.Y < len(b.board)
}
