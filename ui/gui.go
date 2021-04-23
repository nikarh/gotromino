package ui

import (
	"fmt"
	"github.com/gdamore/tcell"
	"image"

	"github.com/nikarh/gotromino/game"
)

var options = struct {
	shadows bool
}{
	shadows: true,
}

func Init() (*Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	return &Screen{screen}, nil
}

func NewGame(screen *Screen) bool {
	redraw := make(chan struct{})
	g := game.New(image.Pt(10, 22), redraw)

	go func() {
		screen.draw(g)
		for range redraw {
			screen.draw(g)
		}
	}()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Rune() {
			case 'p':
				if !g.End {
					g.Paused = !g.Paused
					redraw <- struct{}{}
				}
			case 's':
				options.shadows = !options.shadows
				redraw <- struct{}{}
			case 'z':
				g.Rotate(game.Left)
			case 'x':
				g.Rotate(game.Right)
			case ' ':
				g.HardDrop()
			}

			switch ev.Key() {
			case tcell.KeyLeft:
				g.Move(game.Left)
			case tcell.KeyRight:
				g.Move(game.Right)
			case tcell.KeyDown:
				g.SoftDrop()
			case tcell.KeyUp:
				g.Rotate(game.Right)
			case tcell.KeyEnter:
				if g.End {
					close(redraw)
					return true
				}
			case tcell.KeyEscape:
				return false
			case tcell.KeyCtrlC:
				return false
			}
		case *tcell.EventResize:
			screen.Sync()
			redraw <- struct{}{}
		}
	}
}

func (s Screen) draw(g *game.Game) {
	s.SetStyle(tcell.StyleDefault)
	s.Clear()

	w, _ := s.Size()
	fx, fy := g.Matrix.Size.X*2, g.Matrix.Size.Y-2
	sysw, queuew := 25, 13

	offset := image.Pt((w-fx-sysw-queuew-4)/2, 2)

	s.tbPrintRect(image.Rect(offset.X, offset.Y, offset.X+fx+sysw+queuew+1, offset.Y+fy+2+1))
	s.tbPrintRect(image.Rect(offset.X, offset.Y+2, offset.X+fx+sysw+queuew+1, offset.Y+fy+2+1))
	s.tbPrintString("Gotromino", offset.Add(image.Pt((fx+sysw+queuew)/2-3, 1)))

	// Queue
	qo := offset.Add(image.Pt(fx+1, 2))
	s.tbPrintRect(image.Rect(qo.X, qo.Y, qo.X+queuew, qo.Y+fy+1))
	s.tbPrintString("Next", qo.Add(image.Pt(5, 1)))


	next := g.Queue.Polyominoes
	s.tbPrintRect(image.Rect(qo.X+2, qo.Y+2, qo.X+9+2, qo.Y+2+5))
	s.tbPrintPolyomino(next[0], qo.Add(image.Pt(7-int(next[0].Dim), 4)))

	for i := 1; i < len(next); i++ {
		s.tbPrintPolyomino(next[i], qo.Add(image.Pt(7-int(next[i].Dim), 6+i*3)))
	}

	sm := offset.Add(image.Pt(2+fx+sysw/2+queuew, 3))
	// Score
	dx, dy := -10, 1
	s.tbPrintString(fmt.Sprintf("Level: %d", g.Level), sm.Add(image.Pt(dx, dy)))
	s.tbPrintString(fmt.Sprintf("Lines: %d", g.Lines), sm.Add(image.Pt(dx, dy+2)))
	s.tbPrintString(fmt.Sprintf("Score: %d", g.Score), sm.Add(image.Pt(dx, dy+3)))

	// Controls
	dx, dy = -11, 10
	s.tbPrintRect(image.Rect(sm.X-13, sm.Y+dy-1, sm.X+12, sm.Y+dy+10))
	s.tbPrintString("Esc   - exit", sm.Add(image.Pt(dx, dy)))
	s.tbPrintString("p     - pause", sm.Add(image.Pt(dx, dy+1)))

	s.tbPrintString("←, →  - move", sm.Add(image.Pt(dx, dy+3)))
	s.tbPrintString("z     - turn CCW", sm.Add(image.Pt(dx, dy+4)))
	s.tbPrintString("x, ↑  - turn CW", sm.Add(image.Pt(dx, dy+5)))
	s.tbPrintString("↓     - soft drop", sm.Add(image.Pt(dx, dy+6)))
	s.tbPrintString("space - hard drop", sm.Add(image.Pt(dx, dy+7)))

	s.tbPrintString("s     - toggle shadows", sm.Add(image.Pt(dx, dy+9)))

	// Game matrix
	s.tbPrintRect(image.Rect(offset.X, offset.Y+2, offset.X+fx+1, offset.Y+fy+2+1))
	offset = offset.Add(image.Pt(1, 3))
	s.tbPrintMatrix(g.Matrix, offset)
	if options.shadows {
		s.tbPrintShadow(g.Shadow, offset)
	}
	s.tbPrintPiece(g.Piece, offset)

	dx = w / 2
	if g.Paused {
		s.tbInfo("Press p to unpause", image.Rect(dx-11, 12, dx+11, 16))
	}
	if g.End {
		s.tbInfo(
			fmt.Sprintf("Topped out! Score: %d\n\nPress Enter to restart", g.Score),
			image.Rect(dx-16, 11, dx+16, 17))
	}

	s.Show()
}
