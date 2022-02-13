package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type GameFactory struct {
	Players []*Player

	OutputPath string
}

func initGameFactory() *GameFactory {
	return &GameFactory{}
}

func (gf *GameFactory) withPlayer(p *Player) *GameFactory {
	gf.Players = append(gf.Players, p)
	return gf
}

func (gf *GameFactory) withOutputPath(path string) *GameFactory {
	gf.OutputPath = path
	return gf
}

func (gf *GameFactory) build() *Game {
	return &Game{
		Players:    gf.Players,
		OutputPath: gf.OutputPath,
	}
}

type Game struct {
	Players []*Player `json:"players"`

	OutputPath string `json:"output_path"`
}

func (g *Game) Play() {
	logState := &LogState{}

	for round := 1; round <= 12; round++ {
		for _, player := range g.Players {
			diceSet := NewDiceSet()
			for trial := 1; trial <= 3; trial++ {
				diceSet.Roll()
				state := createState(g, round, trial, player, diceSet)

				decision, next := player.Play(state, diceSet)

				logState.Log = append(logState.Log, LogStateElem{*state.InnerState, decision.InnerDecision})
				if !next {
					break
				}
			}
		}
	}

	logState.Result = *createState(g, 0, 0, nil, nil).InnerState

	g.PrintLog(logState)
}

func (g *Game) PrintLog(logState *LogState) {
	b, _ := json.MarshalIndent(logState, "", " ")

	if g.OutputPath == "" {
		fmt.Printf("%s\n", b)
	} else {
		f, err := os.Create(g.OutputPath)
		if err != nil {
			log.Fatalf("failed to create file(%s): %v", g.OutputPath, err)
		}
		defer f.Close()
		f.Write(b)
	}

}
