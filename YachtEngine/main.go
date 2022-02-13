package main

import (
	"engine"
)

func main() {
	game := engine.InitGame(engine.GameConfig{
		NumPlayers:     2,
		AlgorithmPorts: []int{9910, 9910},
		OutputPath:     "./log.json",
	})
	game.Play()
}
