package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/nikarh/gotromino/ui"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	screen, err := ui.Init()
	if err != nil {
		log.Panic(err)
	}

	if err := screen.Init(); err != nil {
		log.Panic(err)
	}

	defer screen.Fini()

	for ui.NewGame(screen) {
	}
}
