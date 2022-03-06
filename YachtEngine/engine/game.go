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
	playerIDs := []string{}
	for _, player := range g.Players {
		playerIDs = append(playerIDs, player.ID)
	}
	logState := &LogState{Players: playerIDs}
	defer g.PrintLog(logState)

	for round := 1; round <= 12; round++ {
		for _, player := range g.Players {
			diceSet := NewDiceSet()
			for trial := 1; trial <= 3; trial++ {
				diceSet.Roll()
				state := createState(g, round, trial, player, diceSet)

				decision, next, err := player.Play(state, diceSet)
				if err != nil {
					log.Printf("err %v\n", err)
					logState.Result = *createState(g, 0, 0, nil, nil).InnerState
					logState.Result.FoulBy = player.ID
					logState.Result.FoulDetail = fmt.Sprintf("%v", err)
					logState.Result.Winner = []string{}
					for _, p := range g.Players {
						if p.ID == player.ID {
							continue
						}
						logState.Result.Winner = append(logState.Result.Winner, p.ID)
					}
					return
				}

				logState.Log = append(logState.Log, LogStateElem{*state.InnerState, decision.InnerDecision})
				if !next {
					break
				}
			}
		}
	}

	logState.Result = *createState(g, 0, 0, nil, nil).InnerState
	logState.Result.Winner = calculateWinner(logState.Result)
}

func calculateWinner(innerState InnerState) []string {
	maxScore := -1
	winner := []string{}
	for playerID, scoreboard := range innerState.ScoreBoard {
		if *scoreboard.Total > maxScore {
			winner = []string{playerID}
			maxScore = *scoreboard.Total
		} else if *scoreboard.Total == maxScore {
			winner = append(winner, playerID)
		}
	}
	return winner
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
