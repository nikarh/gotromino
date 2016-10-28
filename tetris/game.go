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
	Pos   image.Point
	Tile  Tile
	Color uint8
}

func NewPiece(fieldWidth int) Piece {
	i := rand.Intn(len(tiles))
	tile := tiles[i]
	return Piece{
		Pos:   image.Pt((fieldWidth-int(tile.Dim))/2, 0),
		Tile:  tile,
		Color: uint8(i + 1),
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
	tile := g.Piece.Tile.Rotate()
	if !g.Field.Fits(tile, g.Piece.Pos) {
		return false
	}
	g.Piece.Tile = tile
	return true
}

func (g *Game) Move(d image.Point) bool {
	pos := g.Piece.Pos.Add(d)
	if !g.Field.Fits(g.Piece.Tile, pos) {
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

		g.Piece = g.Next
		g.Next = NewPiece(g.Field.Size.X)
	}

	return true
}
