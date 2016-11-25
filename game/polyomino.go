package game

import (
	"image"
)

type Polyomino struct {
	Id     rune
	Dim    uint8
	Points []uint8
}

func (t Polyomino) Rotate(d image.Point) Polyomino {
	rotated := make([]uint8, len(t.Points))
	switch d {
	case Left:
		for i, p := range t.Points {
			rotated[i] = p<<4 + (t.Dim - 1 - p>>4) // x, y = y, d-1-x
		}
		return Polyomino{t.Id, t.Dim, rotated}
	case Right:
		for i, p := range t.Points {
			rotated[i] = (t.Dim-1-p&0x0F)<<4 + p>>4 // x, y = d-1-y, x
		}
		return Polyomino{t.Id, t.Dim, rotated}
	default:
		return t
	}

}

var tetrominoes = []Polyomino{
	{'I', 4, []uint8{0x01, 0x11, 0x21, 0x31}}, // 0000_1111_0000_0000
	{'J', 3, []uint8{0x00, 0x01, 0x11, 0x21}}, // 100_111_000
	{'L', 3, []uint8{0x01, 0x11, 0x21, 0x20}}, // 001_111_000
	{'O', 2, []uint8{0x00, 0x01, 0x11, 0x10}}, // 11_11
	{'S', 3, []uint8{0x10, 0x20, 0x01, 0x11}}, // 011_110_000
	{'T', 3, []uint8{0x10, 0x01, 0x11, 0x21}}, // 010_111_000
	{'Z', 3, []uint8{0x00, 0x10, 0x11, 0x21}}, // 110_011_000
}
