package engine

type State struct {
	*InnerState `json:"state"`
}
type InnerState struct {
	Winner     string                      `json:"winner,omitempty"`
	Turn       int                         `json:"turn,omitempty"`
	Player     string                      `json:"player,omitempty"`
	Trial      int                         `json:"trial,omitempty"`
	Dices      [5]int                      `json:"dices"`
	ScoreBoard map[string]ScoreBoardResult `json:"scoreBoard"`
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
		ScoreBoard: map[string]ScoreBoardResult{},
	}
	for _, player := range g.Players {
		state.ScoreBoard[player.ID] = player.ScoreBoard.ToScoreBoardResult()
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
