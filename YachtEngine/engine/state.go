package engine

type State struct {
	*InnerState `json:"state"`
}
type InnerState struct {
	Turn       int                   `json:"turn,omitempty"`
	Player     string                `json:"player,omitempty"`
	Trial      int                   `json:"trial,omitempty"`
	Dices      [5]int                `json:"dices"`
	ScoreBoard map[string]ScoreBoard `json:"scoreBoard"`
}

func createState(g *Game, round, trial int, p *Player, diceSet *DiceSet) *State {
	playerID := ""
	if p != nil {
		playerID = p.ID
	}
	state := &InnerState{
		Turn:       round,
		Player:     playerID,
		Trial:      trial,
		Dices:      diceSet.ToInts(),
		ScoreBoard: map[string]ScoreBoard{},
	}
	for _, player := range g.Players {
		state.ScoreBoard[player.ID] = *player.ScoreBoard
	}
	return &State{InnerState: state}
}

type LogStateElem struct {
	InnerState    `json:"state"`
	InnerDecision `json:"decision"`
}

type LogState struct {
	Log    []LogStateElem `json:"log"`
	Result InnerState     `json:"final"`
}
