package main

import (
	"github.com/gilwong00/snake/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
