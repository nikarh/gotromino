package main

import (
	"math/rand"
	"time"

	"github.com/nikarh/gotromino/ui"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	destroy := ui.Init()
	defer destroy()

	for ui.NewGame() {
	}
}
