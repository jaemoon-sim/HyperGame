package engine

type State struct {
	*InnerState `json:"state"`
}
type InnerState struct {
	Turn       int                    `json:"turn"`
	Player     string                 `json:"player"`
	Trial      int                    `json:"trial"`
	Dices      [5]int                 `json:"dices"`
	ScoreBoard map[string]*ScoreBoard `json:"scoreBoard"`
}

func createState(g *Game, round, trial int, p *Player, diceSet *DiceSet) *State {
	state := &InnerState{
		Turn:       round,
		Player:     p.ID,
		Trial:      trial,
		Dices:      diceSet.ToInts(),
		ScoreBoard: map[string]*ScoreBoard{},
	}
	for _, player := range g.Players {
		state.ScoreBoard[player.ID] = player.ScoreBoard
	}
	return &State{InnerState: state}
}
