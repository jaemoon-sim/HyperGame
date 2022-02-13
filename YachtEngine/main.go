package main

import (
	"engine"
	"fmt"
)

func main() {
	fmt.Println("### hello world ###")

	game := engine.InitGame(engine.GameConfig{
		NumPlayers:     1,
		AlgorithmPorts: []int{9910},
	})
	game.Play()
}
