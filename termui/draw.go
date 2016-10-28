package termui

import (
	"image"

	"github.com/nikarh/gotromino/tetris"
	"github.com/nsf/termbox-go"
)

func tbprintString(msg string, offset image.Point) {
	for i, c := range msg {
		termbox.SetCell(offset.X+i, offset.Y, c, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func tbprintField(field [][]uint8, offset image.Point) {
	w, h := len(field[0]), len(field)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			tbprintBlock(image.Pt(x*2+offset.X, y+offset.Y), field[y][x])
		}
	}
}

func tbprintPiece(p tetris.Piece, offset image.Point) {
	for _, pt := range p.Tile.Points {
		x, y := int(pt>>4)+p.Pos.X, int(pt&0x0F)+p.Pos.Y
		tbprintBlock(image.Pt(x*2+offset.X, y+offset.Y), p.Color)
	}
}

func tbprintPieceNoOffset(p tetris.Piece, offset image.Point) {
	for _, pt := range p.Tile.Points {
		x, y := int(pt>>4), int(pt&0x0F)
		tbprintBlock(image.Pt(x*2+offset.X, y+offset.Y), p.Color)
	}
}

func tbprintBlock(pos image.Point, color uint8) {
	termbox.SetCell(pos.X, pos.Y, '█', termbox.Attribute(color), termbox.ColorDefault)
	termbox.SetCell(pos.X+1, pos.Y, '█', termbox.Attribute(color), termbox.ColorDefault)
}

func tbfill(rect image.Rectangle, color termbox.Attribute) {
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			termbox.SetCell(x, y, '█', color, color)
		}
	}
}

func tbprintRect(rect image.Rectangle) {
	termbox.SetCell(rect.Min.X, rect.Min.Y, '┌', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(rect.Max.X, rect.Min.Y, '┐', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(rect.Min.X, rect.Max.Y, '└', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(rect.Max.X, rect.Max.Y, '┘', termbox.ColorDefault, termbox.ColorDefault)

	for x := rect.Min.X + 1; x < rect.Max.X; x++ {
		termbox.SetCell(x, rect.Min.Y, '─', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x, rect.Max.Y, '─', termbox.ColorDefault, termbox.ColorDefault)
	}
	for y := rect.Min.Y + 1; y < rect.Max.Y; y++ {
		termbox.SetCell(rect.Min.X, y, '│', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(rect.Max.X, y, '│', termbox.ColorDefault, termbox.ColorDefault)
	}
}
