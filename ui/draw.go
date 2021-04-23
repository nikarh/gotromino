package ui

import (
	"github.com/gdamore/tcell"
	"github.com/nikarh/gotromino/game"
	"image"

	"strings"
)

var colorTable = map[rune]tcell.Color{
	0: tcell.ColorDefault,
	'I': tcell.ColorAqua,
	'O': tcell.ColorYellow,
	'T': tcell.ColorFuchsia,
	'S': tcell.ColorGreen,
	'Z': tcell.ColorRed,
	'J': tcell.ColorBlue,
	'L': tcell.ColorWhite,
}

type Screen struct {
	tcell.Screen
}

func (s Screen) tbPrintString(msg string, offset image.Point) {
	i := 0
	for _, c := range msg {
		s.SetCell(offset.X+i, offset.Y, tcell.StyleDefault, c)
		i++
	}
}

func (s Screen) tbPrintMatrix(f game.Matrix, offset image.Point) {
	w, h := f.Size.X, f.Size.Y
	for y := 2; y < h; y++ {
		for x := 0; x < w; x++ {
			s.tbPrintBlock(image.Pt(x*2+offset.X, y+offset.Y-2), colorTable[f.Raw[y][x]])
		}
	}
}

func (s Screen) tbPrintPiece(p game.Piece, offset image.Point) {
	for _, pt := range p.Polyomino.Points {
		x, y := int(pt>>4)+p.Pos.X, int(pt&0x0F)+p.Pos.Y
		if y > 1 {
			s.tbPrintBlock(image.Pt(x*2+offset.X, y+offset.Y-2), colorTable[p.Polyomino.Id])
		}
	}
}

func (s Screen) tbPrintShadow(p game.Piece, offset image.Point) {
	for _, pt := range p.Polyomino.Points {
		x, y := int(pt>>4)+p.Pos.X, int(pt&0x0F)+p.Pos.Y
		if y > 1 {
			s.tbPrintShadowBlock(image.Pt(x*2+offset.X, y+offset.Y-2))
		}
	}
}

func (s Screen) tbPrintPolyomino(t game.Polyomino, offset image.Point) {
	for _, pt := range t.Points {
		x, y := int(pt>>4), int(pt&0x0F)
		s.tbPrintBlock(image.Pt(x*2+offset.X, y+offset.Y), colorTable[t.Id])
	}
}

func (s Screen) tbPrintBlock(pos image.Point, color tcell.Color) {
	s.SetCell(pos.X, pos.Y, tcell.StyleDefault.Background(color), ' ')
	s.SetCell(pos.X+1, pos.Y, tcell.StyleDefault.Background(color), ' ')
}

func (s Screen) tbPrintShadowBlock(pos image.Point) {
	s.SetCell(pos.X, pos.Y, tcell.StyleDefault, '╳')
	s.SetCell(pos.X+1, pos.Y, tcell.StyleDefault, '╳')
}

func (s Screen) tbFill(rect image.Rectangle) {
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			s.SetCell(x, y, tcell.StyleDefault, ' ')
		}
	}
}

func (s Screen) tbPrintRect(rect image.Rectangle) {
	s.SetCell(rect.Min.X, rect.Min.Y, tcell.StyleDefault, '┌')
	s.SetCell(rect.Max.X, rect.Min.Y, tcell.StyleDefault, '┐')
	s.SetCell(rect.Min.X, rect.Max.Y, tcell.StyleDefault, '└')
	s.SetCell(rect.Max.X, rect.Max.Y, tcell.StyleDefault, '┘')

	for x := rect.Min.X + 1; x < rect.Max.X; x++ {
		s.SetCell(x, rect.Min.Y, tcell.StyleDefault, '─')
		s.SetCell(x, rect.Max.Y, tcell.StyleDefault, '─')
	}
	for y := rect.Min.Y + 1; y < rect.Max.Y; y++ {
		s.SetCell(rect.Min.X, y, tcell.StyleDefault, '│')
		s.SetCell(rect.Max.X, y, tcell.StyleDefault, '│')
	}
}

func (s Screen) tbInfo(msg string, rect image.Rectangle) {
	s.tbFill(rect)

	s.SetCell(rect.Min.X, rect.Min.Y, tcell.StyleDefault, '╔')
	s.SetCell(rect.Max.X, rect.Min.Y, tcell.StyleDefault, '╗')
	s.SetCell(rect.Min.X, rect.Max.Y, tcell.StyleDefault, '╚')
	s.SetCell(rect.Max.X, rect.Max.Y, tcell.StyleDefault, '╝')

	for x := rect.Min.X + 1; x < rect.Max.X; x++ {
		s.SetCell(x, rect.Min.Y, tcell.StyleDefault, '═')
		s.SetCell(x, rect.Max.Y, tcell.StyleDefault, '═')
	}
	for y := rect.Min.Y + 1; y < rect.Max.Y; y++ {
		s.SetCell(rect.Min.X, y, tcell.StyleDefault, '║')
		s.SetCell(rect.Max.X, y, tcell.StyleDefault, '║')
	}

	lines := strings.Split(msg, "\n")
	for i, line := range lines {
		s.tbPrintString(line,
			image.Pt(rect.Min.X+(rect.Dx()-len(line))/2,
				rect.Min.Y+(rect.Dy()-len(lines)+1)/2+i))
	}
}

