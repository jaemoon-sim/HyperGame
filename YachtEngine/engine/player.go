package engine

import "log"

type Player struct {
	ID          string `json:"id"`
	*ScoreBoard `json:"scoreboard"`

	AlgorithmPath string `json:"-"`
}

type Decision struct {
	Keep       []int  `json:"keep"`
	ScoreTitle string `json:"score_title"`
}

func contains(intSlice []int, num int) bool {
	for _, i := range intSlice {
		if i == num {
			return true
		}
	}
	return false
}

func (p *Player) Play(trial int, diceSet *DiceSet) (nextTrial bool) {
	decision := p.GetDecision(diceSet)
	if decision.ScoreTitle != "" {
		err := p.ScoreBoard.SetScore(decision.ScoreTitle, diceSet)
		if err != nil {
			log.Fatalf("player(%s) invalid score title %s: %v\n", p.ID, decision.ScoreTitle, err)
		}
		return false
	}

	for i := range diceSet {
		diceSet[i].Kept = contains(decision.Keep, i)
	}
	return true
}

func TODO() {}

func (p *Player) GetDecision(diceSet *DiceSet) Decision {
	TODO()
	return Decision{}
}
