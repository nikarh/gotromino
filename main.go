package main

import (
	"math/rand"
	"time"

	"github.com/nikarh/gotromino/termui"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	destroy := termui.Init()
	defer destroy()

	for termui.NewGame() {
	}
}
