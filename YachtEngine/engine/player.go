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

func (p *Player) Play(state *State, diceSet *DiceSet) (decision Decision, nextTrial bool, err error) {
	decision, err = p.GetDecision(state)
	if err != nil {
		return Decision{}, nextTrial, err
	}
	if decision.InnerDecision.ScoreTitle != nil && *decision.InnerDecision.ScoreTitle != "" {
		err := p.ScoreBoard.SetScore(*decision.InnerDecision.ScoreTitle, diceSet)
		if err != nil {
			return Decision{}, false, fmt.Errorf("invalid score title %s: %w", *decision.InnerDecision.ScoreTitle, err)
		}
		return decision, false, nil
	}
	if state.Trial == 3 {
		return Decision{}, false, fmt.Errorf("no choice")
	}
	decision.keepDices(diceSet)

	return decision, true, nil
}

func (p *Player) GetDecision(state *State) (Decision, error) {
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
		log.Printf("%s\n", string(respBody))
		return Decision{}, fmt.Errorf("invalid decision json")
	}
	return decision, nil
}
