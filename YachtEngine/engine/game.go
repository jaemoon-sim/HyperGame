package engine

import (
	"encoding/json"
	"fmt"
)

type GameFactory struct {
	Players []*Player
}

func initGameFactory() *GameFactory {
	return &GameFactory{}
}

func (gf *GameFactory) withPlayer(p *Player) *GameFactory {
	gf.Players = append(gf.Players, p)
	return gf
}

func (gf *GameFactory) build() *Game {
	return &Game{
		Players: gf.Players,
	}
}

type Game struct {
	Players []*Player `json:"players"`
}

func (g *Game) Play() {
	for round := 1; round <= 12; round++ {
		fmt.Printf("###### round %d #######\n", round)
		for _, player := range g.Players {
			fmt.Printf("   ### player %s\n", player.ID)

			diceSet := NewDiceSet()
			for trial := 1; trial <= 3; trial++ {
				diceSet.Roll()

				state := createState(g, round, trial, player, diceSet)
				b, _ := json.MarshalIndent(state, "", " ")
				fmt.Printf("%s\n", b)
				if next := player.Play(state, diceSet); !next {
					break
				}
			}
		}
	}

	fmt.Printf("FINISHED !!! 🎉\n")
	b, _ := json.MarshalIndent(g, "", " ")
	fmt.Printf("%s\n", b)
}
