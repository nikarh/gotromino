package tetris

import "testing"

func TestRotation(t *testing.T) {

	tile := tiles[3]

	t.Log(tile)

	t.Log(tile.Rotate())

}
