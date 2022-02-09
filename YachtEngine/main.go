package main

import (
	"engine"
	"fmt"
)

func main() {
	fmt.Println("### hello world ###")

	game := engine.InitGame(engine.GameConfig{
		NumPlayers:     2,
		AlgorithmPaths: []string{"", ""},
	})
	game.Play()
}
