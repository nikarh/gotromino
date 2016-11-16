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
}

func randomTetromino() Tetromino {
	return tetrominos[rand.Intn(len(tetrominos))]
}

type Game struct {
	Next  Tetromino
	Piece Piece
	Field Field

	Level uint8
	Lines uint32
	Score uint32

	Paused bool
	End    bool

	busy bool

	redraw chan struct{}
	action chan func()
}

func newTicker(level uint8) *time.Ticker {
	speed := 725*math.Pow(0.85, float64(level)) + float64(level)
	return time.NewTicker(time.Duration(speed) * time.Millisecond)
}

func NewGame(size image.Point, redraw chan struct{}) *Game {
	g := &Game{
		Field: newField(size),
		Next:  randomTetromino(),

		redraw: redraw,
		action: make(chan func()),
	}

	g.nextPiece()

	ticker := newTicker(g.Level)
	go (func() {
		defer ticker.Stop()
		defer close(g.action)

		for {
			select {
			case <-ticker.C:
				if g.Paused {
					continue
				}
				if g.End = !g.tick(); g.End {
					g.redraw <- struct{}{}
					return
				} else if lvl := uint8(g.Lines) / 10; lvl != g.Level {
					g.Level = lvl
					ticker.Stop()
					ticker = newTicker(g.Level)
				}
			case action := <-g.action:
				action()
			}
			g.redraw <- struct{}{}
		}
	})()

	return g
}

func (g *Game) nextPiece() {
	g.Piece = Piece{
		Pos:       image.Pt((g.Field.Size.X-int(g.Next.Dim))/2, 0),
		Tetromino: g.Next,
	}
	g.Next = randomTetromino()
}

func (g *Game) tick() bool {
	if !g.move(Down) {
		g.Field.Put(g.Piece)

		c := g.Field.FindCompleted()
		g.addScore(uint32(len(c)))
		g.Field.Clear(c)

		if len(c) < 2 && g.Field.Full() {
			return false
		}

		g.nextPiece()
		if !g.Field.Fits(g.Piece.Tetromino, g.Piece.Pos) {
			return false
		}
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

func (g *Game) move(d image.Point) bool {
	pos := g.Piece.Pos.Add(d)
	if !g.Field.Fits(g.Piece.Tetromino, pos) {
		return false
	}
	g.Piece.Pos = pos

	return true
}

func (g *Game) act(f func()) {
	if g.busy || g.Paused || g.End {
		return
	}
	g.action <- f
}

func (g *Game) Rotate(d image.Point) {
	g.act(func() {
		rotated := g.Piece.Tetromino.Rotate(d)
		if !g.Field.Fits(rotated, g.Piece.Pos) {
			return
		}
		g.Piece.Tetromino = rotated
	})
}

func (g *Game) Move(d image.Point) {
	g.act(func() {
		g.move(d)
	})
}

func (g *Game) SoftDrop() {
	g.act(func() {
		if g.move(Down) {
			g.Score += 1
		}
	})
}

func (g *Game) HardDrop() {
	g.act(func() {
		g.busy = true
		t := time.NewTicker(time.Millisecond * 10)
		for g.move(Down) {
			g.Score += 2
			g.redraw <- struct{}{}
			<-t.C
		}
		t.Stop()
		g.tick()
		g.busy = false
	})
}
