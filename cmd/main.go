package main

import (
	"github.com/gilwong00/snake/internal/game"
)

func main() {
	game := game.NewGame()
	if err := game.RunGame(); err != nil {
		panic(err)
	}
}
