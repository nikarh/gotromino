package tetris

import (
	"image"
	"math"
	"math/rand"
	"time"
)

var (
	Left  = image.Pt(-1, 0)
	Right = image.Pt(1, 0)
	Down  = image.Pt(0, 1)
)

type Piece struct {
	Pos       image.Point
	Tetromino Tetromino
	Color     uint8
}

func NewPiece(fieldWidth int) Piece {
	i := rand.Intn(len(tetrominos))
	tetromino := tetrominos[i]
	return Piece{
		Pos:       image.Pt((fieldWidth-int(tetromino.Dim))/2, 0),
		Tetromino: tetromino,
		Color:     uint8(i + 2),
	}
}

type Game struct {
	Piece Piece
	Next  Piece
	Field Field

	Level uint8
	Lines uint32
	Score uint32

	Paused bool
	End    bool

	Refresh chan struct{}
	Actions chan func()
}

func newTicker(level uint8) *time.Ticker {
	speed := 725*math.Pow(0.85, float64(level)) + float64(level)
	return time.NewTicker(time.Duration(speed) * time.Millisecond)
}

func NewGame(size image.Point) *Game {
	g := &Game{
		Field: NewField(size),
		Piece: NewPiece(size.X),
		Next:  NewPiece(size.X),

		Refresh: make(chan struct{}),
		Actions: make(chan func()),
	}

	ticker := newTicker(g.Level)
	go (func() {
		for {
			select {
			case <-ticker.C:
				if g.Paused {
					continue
				}
				if g.End = !g.tick(); g.End {
					ticker.Stop()
				} else if lvl := uint8(g.Lines) / 10; lvl != g.Level {
					g.Level = lvl
					ticker.Stop()
					ticker = newTicker(g.Level)
				}
				g.Refresh <- struct{}{}
			case action := <-g.Actions:
				if g.End || g.Paused {
					continue
				}
				action()
			}
		}
	})()

	return g
}

func (g *Game) tick() bool {
	if !g.Move(Down) {
		g.Field.Put(g.Piece)

		c := g.Field.FindCompleted()
		g.addScore(uint32(len(c)))
		g.Field.Clear(c)

		if len(c) < 2 && g.Field.Full() {
			return false
		}

		if !g.Field.Fits(g.Next.Tetromino, g.Next.Pos) {
			return false
		}

		g.Piece = g.Next
		g.Next = NewPiece(g.Field.Size.X)
	}

	return true
}

func (g *Game) addScore(lines uint32) {
	g.Lines += lines
	switch lines {
	case 0:
	case 1:
		g.Score += 40 * uint32(g.Level)
	case 2:
		g.Score += 100 * uint32(g.Level)
	case 3:
		g.Score += 300 * uint32(g.Level)
	default:
		g.Score += lines * 300 * uint32(g.Level)
	}
}

func (g *Game) Rotate() bool {
	t := g.Piece.Tetromino.Rotate()
	if !g.Field.Fits(t, g.Piece.Pos) {
		return false
	}
	g.Piece.Tetromino = t
	return true
}

func (g *Game) Move(d image.Point) bool {
	pos := g.Piece.Pos.Add(d)
	if !g.Field.Fits(g.Piece.Tetromino, pos) {
		return false
	}
	g.Piece.Pos = pos

	return true
}

func (g *Game) SoftDrop() {
	if g.Move(Down) {
		g.Score += 1
	}
}

func (g *Game) HardDrop() {
	t := time.NewTicker(time.Millisecond * 10)
	for g.Move(Down) {
		g.Score += 2
		g.Refresh <- struct{}{}
		<-t.C
	}
	t.Stop()
	g.tick()
	g.Refresh <- struct{}{}
}
