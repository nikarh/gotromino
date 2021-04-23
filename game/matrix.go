package game

import (
	"image"
)

type Matrix struct {
	Size image.Point
	Raw  [][]rune
}

func newMatrix(size image.Point) Matrix {
	raw := make([][]rune, size.Y)
	for y := 0; y < size.Y; y++ {
		raw[y] = make([]rune, size.X)
	}

	return Matrix{
		Size: size,
		Raw:  raw,
	}
}

func (m Matrix) Fits(p Polyomino, pos image.Point) bool {
	for _, pt := range p.Points {
		x, y := pos.X+int(pt>>4), pos.Y+int(pt&0x0F)
		if x < 0 || y < 0 || x >= m.Size.X || y >= m.Size.Y {
			return false
		}
		if m.Raw[y][x] > 0 {
			return false
		}
	}
	return true
}

func (m Matrix) Put(p Polyomino, pos image.Point) {
	for _, pt := range p.Points {
		x, y := pos.X+int(pt>>4), pos.Y+int(pt&0x0F)
		m.Raw[y][x] = p.Id
	}
}

func (m Matrix) FullLines() []int {
	result := make([]int, 0)

	for y := 0; y < m.Size.Y; y++ {
		completed := true
		for x := 0; x < m.Size.X; x++ {
			if m.Raw[y][x] == 0 {
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

func (m Matrix) Clear(lines []int) {
	for i, cy := range lines {
		for y := cy; y > i; y-- {
			m.Raw[y] = m.Raw[y-1]
		}
		m.Raw[i] = make([]rune, m.Size.X)
	}
}

func (m Matrix) Full() bool {
	for y := 1; y >= 0; y-- {
		for x := 0; x < m.Size.X; x++ {
			if m.Raw[y][x] > 0 {
				return true
			}
		}
	}
	return false
}
