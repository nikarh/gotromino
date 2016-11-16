package tetris

import (
	"fmt"
	"image"
)

type Tetromino struct {
	Color  uint8
	Dim    uint8
	Points []uint8
}

func (t Tetromino) Rotate(d image.Point) Tetromino {
	rotated := make([]uint8, len(t.Points))
	switch d {
	case Left:
		for i, p := range t.Points {
			rotated[i] = p<<4 + (t.Dim - 1 - p>>4) // x, y = y, d-1-x
		}
		return Tetromino{t.Color, t.Dim, rotated}
	case Right:
		for i, p := range t.Points {
			rotated[i] = (t.Dim-1-p&0x0F)<<4 + p>>4 // x, y = d-1-y, x
		}
		return Tetromino{t.Color, t.Dim, rotated}
	default:
		return t
	}

}

func (t Tetromino) String() string {
	str := "{ "
	for _, c := range t.Points {
		str = str + fmt.Sprintf("%02x ", c)
	}
	return str + "}"
}

var tetrominos = []Tetromino{
	{2, 4, []uint8{0x01, 0x11, 0x21, 0x31}}, // I 0000_1111_0000_0000
	{3, 2, []uint8{0x00, 0x01, 0x11, 0x10}}, // O 11_11
	{4, 3, []uint8{0x10, 0x01, 0x11, 0x21}}, // T 010_111_000
	{5, 3, []uint8{0x01, 0x11, 0x21, 0x20}}, // L 001_111_000
	{6, 3, []uint8{0x00, 0x01, 0x11, 0x21}}, // J 100_111_000
	{7, 3, []uint8{0x00, 0x10, 0x11, 0x21}}, // Z 110_011_000
	{8, 3, []uint8{0x10, 0x20, 0x01, 0x11}}, // S 011_110_000
}
