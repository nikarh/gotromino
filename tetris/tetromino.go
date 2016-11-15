package tetris

import "fmt"

type Tetromino struct {
	Dim    uint8
	Points []uint8
}

func (t Tetromino) Rotate() Tetromino {
	rotated := make([]uint8, len(t.Points))
	for i, p := range t.Points {
		rotated[i] = (t.Dim-1-p&0x0F)<<4 + p>>4
	}

	return Tetromino{t.Dim, rotated}
}

func (t Tetromino) String() string {
	str := "{ "
	for _, c := range t.Points {
		str = str + fmt.Sprintf("%02x ", c)
	}
	return str + "}"
}

var tetrominos = []Tetromino{
	{4, []uint8{0x01, 0x11, 0x21, 0x31}}, // I 0000_1111_0000_0000
	{2, []uint8{0x00, 0x01, 0x11, 0x10}}, // O 11_11
	{3, []uint8{0x10, 0x01, 0x11, 0x21}}, // T 010_111_000
	{3, []uint8{0x01, 0x11, 0x21, 0x20}}, // L 001_111_000
	{3, []uint8{0x00, 0x01, 0x11, 0x21}}, // J 100_111_000
	{3, []uint8{0x00, 0x10, 0x11, 0x21}}, // Z 110_011_000
	{3, []uint8{0x10, 0x20, 0x01, 0x11}}, // S 011_110_000
}
