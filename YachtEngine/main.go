package main

import (
	"engine"
	"flag"
	"log"
)

var outputPath = flag.String("output", "", "game result path")

func main() {
	flag.Parse()
	args := flag.Args()

	playerArgs := []engine.PlayerArg{}
	for _, arg := range args {
		pa := &engine.PlayerArg{}
		if err := pa.Set(arg); err != nil {
			log.Fatalf("failed to parse flags: %v", err)
		}
		playerArgs = append(playerArgs, *pa)
	}

	game := engine.InitGame(engine.GameConfig{
		PlayerArgs: playerArgs,
		OutputPath: *outputPath,
	})
	game.Play()
}
