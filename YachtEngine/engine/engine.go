package engine

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayerArg struct {
	Name string
	Port int
}

func (pa *PlayerArg) String() string {
	return fmt.Sprintf("(%s) :%d", pa.Name, pa.Port)
}
func (pa *PlayerArg) Set(arg string) error {
	splitted := strings.Split(arg, ",")
	if len(splitted) != 2 {
		return fmt.Errorf("invalid format of player(%s). should be {name},{port} (e.g., jmsim,9099)", arg)
	}

	port, err := strconv.Atoi(splitted[1])
	if err != nil {
		return err
	}

	pa.Name = splitted[0]
	pa.Port = port
	return nil
}

type GameConfig struct {
	PlayerArgs []PlayerArg

	OutputPath string
}

func InitGame(gameConfig GameConfig) *Game {
	gf := initGameFactory()
	for _, pa := range gameConfig.PlayerArgs {
		p := &Player{
			ID:            pa.Name,
			ScoreBoard:    &ScoreBoard{},
			AlgorithmPort: pa.Port,
		}
		gf = gf.withPlayer(p)
	}
	gf = gf.withOutputPath(gameConfig.OutputPath)
	return gf.build()
}
