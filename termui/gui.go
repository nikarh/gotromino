package termui

import (
	"image"

	"time"

	"os"

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

func Game() {
	g := tetris.NewGame(image.Pt(10, 22))
	t := time.NewTicker(time.Second)

	for {
		select {
		case e := <-eventQueue:
			if e.Type != termbox.EventKey {
				continue
			}
			switch e.Key {
			case termbox.KeyArrowLeft:
				g.Move(tetris.Left)
			case termbox.KeyArrowRight:
				g.Move(tetris.Right)
			case termbox.KeyArrowDown:
				g.Move(tetris.Down)
			case termbox.KeyArrowUp:
				g.Rotate()
			case termbox.KeySpace:
				tt := time.NewTicker(time.Millisecond * 10)
				for g.Move(tetris.Down) {
					draw(g)
					<-tt.C
				}
				tt.Stop()
				g.Tick()
			case termbox.KeyEsc:
				os.Exit(0)
			}
		case <-t.C:
			g.Tick()
		default:
			draw(g)
		}
	}
}

func draw(g *tetris.Game) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	w, _ := termbox.Size()
	fx, fy := g.Field.Size.X*2, g.Field.Size.Y
	sysw := 19

	offset := image.Pt((w-fx-sysw-4)/2, 2)

	tbprintRect(image.Rect(offset.X, offset.Y, offset.X+fx+sysw+1, offset.Y+fy+2+1))
	tbprintRect(image.Rect(offset.X, offset.Y+2, offset.X+fx+sysw+1, offset.Y+fy+2+1))
	tbprintString("Gotromino", offset.Add(image.Pt((fx+sysw)/2-3, 1)))

	sm := offset.Add(image.Pt(2+fx+sysw/2, 3))
	tbprintString("Next piece", sm.Add(image.Pt(-5, 0)))
	tbprintRect(image.Rect(sm.X-5, sm.Y+1, sm.X+4, sm.Y+1+5))
	tbfill(image.Rect(sm.X-4, sm.Y+2, sm.X+4, sm.Y+2+4), termbox.ColorWhite)

	d := int(g.Next.Tile.Dim)
	_ = d
	tbprintPieceNoOffset(g.Next, sm.Add(image.Pt(-d, 3)))

	tbprintRect(image.Rect(offset.X, offset.Y+2, offset.X+fx+1, offset.Y+fy+2+1))

	offset = offset.Add(image.Pt(1, 3))
	tbprintField(g.Field.Raw, offset)
	tbprintPiece(g.Piece, offset)

	termbox.Flush()
}
