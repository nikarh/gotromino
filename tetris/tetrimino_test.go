package tetris

import "testing"

func TestRotation(t *testing.T) {

	tetrimino := tetriminos[3]

	t.Log(tetrimino)

	t.Log(tetrimino.Rotate())

}
