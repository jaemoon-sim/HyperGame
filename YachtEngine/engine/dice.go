package engine

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

type Dice struct {
	Value int  `json:"value"`
	Kept  bool `json:"fixed"`
}

type DiceSet [5]Dice

func NewDiceSet() *DiceSet {
	result := &DiceSet{}
	for i, _ := range result {
		result[i].Value = i + 1
	}
	return result
}

func (ds *DiceSet) String() string {
	slice := make([]int, 5)
	for i, d := range ds {
		slice[i] = int(d.Value)
	}
	sort.Ints(slice)

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%d%d%d%d%d", slice[0], slice[1], slice[2], slice[3], slice[4])
	return buf.String()
}

func (ds *DiceSet) Roll() {
	for _, dice := range ds {
		if dice.Kept {
			continue
		}
		dice.Value = rand.Intn(6) + 1
	}
}
