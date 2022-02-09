package engine

import "fmt"

type GameConfig struct {
	NumPlayers int `json:"num_players"`

	AlgorithmPaths []string `json:"algorithm_paths"`
}

func InitGame(gameConfig GameConfig) *Game {
	gf := initGameFactory()
	for i := 0; i < gameConfig.NumPlayers; i++ {
		p := &Player{
			ID:            fmt.Sprintf("%d", i),
			ScoreBoard:    &ScoreBoard{},
			AlgorithmPath: gameConfig.AlgorithmPaths[i],
		}
		gf = gf.withPlayer(p)
	}

	return gf.build()
}
