package tetris

import (
	"image"
	"math/rand"
)

var (
	Left  = image.Pt(-1, 0)
	Right = image.Pt(1, 0)
	Down  = image.Pt(0, 1)
)

type Piece struct {
	Pos       image.Point
	Tetrimino Tetrimino
	Color     uint8
}

func NewPiece(fieldWidth int) Piece {
	i := rand.Intn(len(tetriminos))
	tetrimino := tetriminos[i]
	return Piece{
		Pos:       image.Pt((fieldWidth-int(tetrimino.Dim))/2, 0),
		Tetrimino: tetrimino,
		Color:     uint8(i + 2),
	}
}

type Game struct {
	Piece Piece
	Next  Piece
	Field Field
}

func NewGame(size image.Point) *Game {
	return &Game{
		Field: NewField(size),
		Piece: NewPiece(size.X),
		Next:  NewPiece(size.X),
	}
}

func (g *Game) Rotate() bool {
	t := g.Piece.Tetrimino.Rotate()
	if !g.Field.Fits(t, g.Piece.Pos) {
		return false
	}
	g.Piece.Tetrimino = t
	return true
}

func (g *Game) Move(d image.Point) bool {
	pos := g.Piece.Pos.Add(d)
	if !g.Field.Fits(g.Piece.Tetrimino, pos) {
		return false
	}
	g.Piece.Pos = pos
	return true
}

func (g *Game) Tick() bool {
	if !g.Move(Down) {
		g.Field.Put(g.Piece)

		c := g.Field.FindCompleted()
		g.Field.Clear(c)

		if len(c) < 2 && g.Field.EndGame() {
			return false
		}

		if !g.Field.Fits(g.Next.Tetrimino, g.Next.Pos) {
			return false
		}

		g.Piece = g.Next
		g.Next = NewPiece(g.Field.Size.X)
	}

	return true
}
