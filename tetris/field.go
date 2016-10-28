package tetris

import (
	"image"
)

type Field struct {
	Size image.Point
	Raw  [][]uint8
}

func NewField(size image.Point) Field {
	raw := make([][]uint8, size.Y)
	for y := 0; y < size.Y; y++ {
		raw[y] = make([]uint8, size.X)
	}

	return Field{
		Size: size,
		Raw:  raw,
	}
}

func (t Field) Fits(ti Tile, pos image.Point) bool {
	for _, pt := range ti.Points {
		x, y := pos.X+int(pt>>4), pos.Y+int(pt&0x0F)
		if x < 0 || y < 0 || x >= t.Size.X || y >= t.Size.Y {
			return false
		}
		if t.Raw[y][x] > 0 {
			return false
		}
	}
	return true
}

func (t Field) Put(p Piece) {
	for _, pt := range p.Tile.Points {
		x, y := p.Pos.X+int(pt>>4), p.Pos.Y+int(pt&0x0F)
		t.Raw[y][x] = p.Color
	}
}

func (t Field) FindCompleted() []int {
	result := make([]int, 0)

	for y := 0; y < t.Size.Y; y++ {
		completed := true
		for x := 0; x < t.Size.X; x++ {
			if t.Raw[y][x] == 0 {
				completed = false
				break
			}
		}
		if completed {
			result = append(result, y)
		}
	}

	return result
}

func (t Field) Clear(lines []int) {
	for i, cy := range lines {
		for y := cy; y > i; y-- {
			t.Raw[y] = t.Raw[y-1]
		}
		t.Raw[i] = make([]uint8, t.Size.X)
	}
}