package engine

import "fmt"

type GameConfig struct {
	NumPlayers int `json:"num_players"`

	AlgorithmPorts []int `json:"algorithm_ports"`
}

func InitGame(gameConfig GameConfig) *Game {
	gf := initGameFactory()
	for i := 0; i < gameConfig.NumPlayers; i++ {
		p := &Player{
			ID:            fmt.Sprintf("%d", i),
			ScoreBoard:    &ScoreBoard{},
			AlgorithmPort: gameConfig.AlgorithmPorts[i],
		}
		gf = gf.withPlayer(p)
	}

	return gf.build()
}
