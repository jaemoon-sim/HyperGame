package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Player struct {
	ID          string `json:"id"`
	*ScoreBoard `json:"scoreboard"`

	AlgorithmPort int `json:"-"`
}

type InnerDecision struct {
	Keep       []int   `json:"keep"`
	ScoreTitle *string `json:"choice"`
}

type Decision struct {
	InnerDecision `json:"decision"`
}

func (d Decision) String() string {
	b, _ := json.MarshalIndent(d, "", " ")
	return string(b)
}

func (d Decision) keepDices(diceSet *DiceSet) {
	for _, k := range d.InnerDecision.Keep {
		diceSet[k].Kept = true
	}
}

func (p *Player) Play(state *State, diceSet *DiceSet) (decision Decision, nextTrial bool) {
	decision = p.GetDecision(state)
	if decision.InnerDecision.ScoreTitle != nil {
		decision.keepDices(diceSet)
		err := p.ScoreBoard.SetScore(*decision.InnerDecision.ScoreTitle, diceSet)
		if err != nil {
			log.Fatalf("player(%s) invalid score title %s: %v\n", p.ID, *decision.InnerDecision.ScoreTitle, err)
		}
		return decision, false
	}

	return decision, true
}

func (p *Player) GetDecision(state *State) Decision {
	b, err := json.Marshal(state)
	if err != nil {
		log.Fatalf("failed to marshal state")
	}
	reqBody := bytes.NewReader(b)
	resp, err := http.Post(fmt.Sprintf("http://localhost:%d/decide", p.AlgorithmPort), "application/json", reqBody)
	if err != nil {
		log.Fatalf("failed to request to player(%s): %v", p.ID, err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to get response from player(%s): %v", p.ID, err)
	}

	decision := Decision{}
	if err := json.Unmarshal(respBody, &decision); err != nil {
		log.Fatalf("failed to parse response from player(%s): %v", p.ID, err)
	}
	return decision
}
