package termui

import (
	"image"
	"fmt"

	"github.com/nikarh/gotromino/tetris"
	"github.com/nsf/termbox-go"
)

var eventQueue = make(chan termbox.Event)

func Init() func() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	return func() {
		termbox.Close()
	}
}

func NewGame() bool {
	g := tetris.NewGame(image.Pt(10, 22))
	for {
		draw(g)
		select {
		case e := <-eventQueue:
			if e.Type != termbox.EventKey {
				continue
			}

			if !g.End && e.Ch == 'p' {
				g.Paused = !g.Paused
			}

			switch e.Key {
			case termbox.KeyArrowLeft:
				g.Actions <- func() {
					g.Move(tetris.Left)
				}
			case termbox.KeyArrowRight:
				g.Actions <- func() {
					g.Move(tetris.Right)
				}
			case termbox.KeyArrowDown:
				g.Actions <- func() {
					g.SoftDrop()
				}
			case termbox.KeyArrowUp:
				g.Actions <- func() {
					g.Rotate()
				}
			case termbox.KeySpace:
				g.Actions <- func() {
					g.HardDrop()
				}
			case termbox.KeyEsc:
				return false
			case termbox.KeyEnter:
				if g.End {
					return true
				}
			}
		case <-g.Refresh:
		}
	}
}

func draw(g *tetris.Game) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	w, _ := termbox.Size()
	fx, fy := g.Field.Size.X*2, g.Field.Size.Y-2
	sysw := 25

	offset := image.Pt((w-fx-sysw-4)/2, 2)

	tbprintRect(image.Rect(offset.X, offset.Y, offset.X+fx+sysw+1, offset.Y+fy+2+1))
	tbprintRect(image.Rect(offset.X, offset.Y+2, offset.X+fx+sysw+1, offset.Y+fy+2+1))
	tbprintString("Gotromino", offset.Add(image.Pt((fx+sysw)/2-3, 1)))

	// Next piece info
	sm := offset.Add(image.Pt(2+fx+sysw/2, 3))
	tbprintString("Next tetromino", sm.Add(image.Pt(-8, 0)))
	tbprintRect(image.Rect(sm.X-5, sm.Y+1, sm.X+4, sm.Y+1+5))
	tbfill(image.Rect(sm.X-4, sm.Y+2, sm.X+4, sm.Y+2+4), termbox.ColorDefault)
	tbprintPieceNoOffset(g.Next, sm.Add(image.Pt(-int(g.Next.Tetromino.Dim), 3)))

	// Score
	dx, dy := -11, 8
	tbprintString(fmt.Sprintf("Level: %d", g.Level), sm.Add(image.Pt(dx, dy)))
	tbprintString(fmt.Sprintf("Lines: %d", g.Lines), sm.Add(image.Pt(dx, dy+1)))
	tbprintString(fmt.Sprintf("Score: %d", g.Score), sm.Add(image.Pt(dx, dy+3)))

	// Controls
	dx, dy = -11, 13
	tbprintRect(image.Rect(sm.X-13, sm.Y+dy-1, sm.X+12, sm.Y+dy+7))
	tbprintString("Esc   - exit", sm.Add(image.Pt(dx, dy)))
	tbprintString("p     - pause", sm.Add(image.Pt(dx, dy+1)))
	tbprintString("←     - move left", sm.Add(image.Pt(dx, dy+2)))
	tbprintString("→     - move right", sm.Add(image.Pt(dx, dy+3)))
	tbprintString("↑     - turn clockwise", sm.Add(image.Pt(dx, dy+4)))
	tbprintString("↓     - soft drop", sm.Add(image.Pt(dx, dy+5)))
	tbprintString("space - hard drop", sm.Add(image.Pt(dx, dy+6)))

	// Game field
	tbprintRect(image.Rect(offset.X, offset.Y+2, offset.X+fx+1, offset.Y+fy+2+1))
	offset = offset.Add(image.Pt(1, 3))
	tbprintField(g.Field, offset)
	tbprintPiece(g.Piece, offset)

	dx = w / 2
	if g.Paused {
		tbInfo("Press p to unpause", image.Rect(dx-11, 12, dx+11, 16))
	}
	if g.End {
		tbInfo(
			fmt.Sprintf("Topped out! Score: %d\n\nPress Enter to restart", g.Score),
			image.Rect(dx-16, 11, dx+16, 17))
	}

	termbox.Flush()
}
