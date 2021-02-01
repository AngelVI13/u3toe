package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/AngelVI13/u3toe/game"
)



func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Ultimate 3Toe")

	gameInstance, err := game.NewGame()

	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}