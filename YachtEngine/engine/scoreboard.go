package engine

import (
	"fmt"
	"strings"
)

var ErrDuplicatedScore = fmt.Errorf("duplicated score")

type ScoreBoard struct {
	Ones          *int `json:"aces"`
	Twos          *int `json:"deuces"`
	Threes        *int `json:"threes"`
	Fours         *int `json:"fours"`
	Fives         *int `json:"fives"`
	Sixs          *int `json:"sixes"`
	Choice        *int `json:"choice"`
	FourOfAKind   *int `json:"fourKind"`
	FullHouse     *int `json:"fullHouse"`
	SmallStraight *int `json:"smallStraight"`
	LargeStraight *int `json:"largeStraight"`
	Yacht         *int `json:"yacht"`
}

func (s *ScoreBoard) SetScore(name string, dices *DiceSet) error {
	switch name {
	case "aces":
		if s.Ones != nil {
			return ErrDuplicatedScore
		}
		s.Ones = s.calculateOnes(dices)
		return nil
	case "deuces":
		if s.Twos != nil {
			return ErrDuplicatedScore
		}
		s.Twos = s.calculateTwos(dices)
		return nil
	case "threes":
		if s.Threes != nil {
			return ErrDuplicatedScore
		}
		s.Threes = s.calculateThrees(dices)
		return nil
	case "fours":
		if s.Fours != nil {
			return ErrDuplicatedScore
		}
		s.Fours = s.calculateFours(dices)
		return nil
	case "fives":
		if s.Fives != nil {
			return ErrDuplicatedScore
		}
		s.Fives = s.calculateFives(dices)
		return nil
	case "sixes":
		if s.Sixs != nil {
			return ErrDuplicatedScore
		}
		s.Sixs = s.calculateSixs(dices)
		return nil
	case "choice":
		if s.Choice != nil {
			return ErrDuplicatedScore
		}
		s.Choice = s.calculateChoice(dices)
		return nil
	case "fourKind":
		if s.FourOfAKind != nil {
			return ErrDuplicatedScore
		}
		s.FourOfAKind = s.calculateFourOfAKind(dices)
		return nil
	case "fullHouse":
		if s.FullHouse != nil {
			return ErrDuplicatedScore
		}
		s.FullHouse = s.calculateFullHouse(dices)
		return nil
	case "smallStraight":
		if s.SmallStraight != nil {
			return ErrDuplicatedScore
		}
		s.SmallStraight = s.calculateSmallStraight(dices)
		return nil
	case "largeStraight":
		if s.LargeStraight != nil {
			return ErrDuplicatedScore
		}
		s.LargeStraight = s.calculateLargeStraight(dices)
		return nil
	case "yacht":
		if s.Yacht != nil {
			return ErrDuplicatedScore
		}
		s.Yacht = s.calculateYacht(dices)
		return nil
	default:
		return fmt.Errorf("invalid score")
	}
}

func (s *ScoreBoard) calculateOnes(dices *DiceSet) *int {
	score := 0
	for _, d := range dices {
		if d.Value == 1 {
			score += d.Value
		}
	}
	return &score
}

func (s *ScoreBoard) calculateTwos(dices *DiceSet) *int {
	score := 0
	for _, d := range dices {
		if d.Value == 2 {
			score += d.Value
		}
	}
	return &score
}

func (s *ScoreBoard) calculateThrees(dices *DiceSet) *int {
	score := 0
	for _, d := range dices {
		if d.Value == 3 {
			score += d.Value
		}
	}
	return &score
}

func (s *ScoreBoard) calculateFours(dices *DiceSet) *int {
	score := 0
	for _, d := range dices {
		if d.Value == 4 {
			score += d.Value
		}
	}
	return &score
}

func (s *ScoreBoard) calculateFives(dices *DiceSet) *int {
	score := 0
	for _, d := range dices {
		if d.Value == 5 {
			score += d.Value
		}
	}
	return &score
}

func (s *ScoreBoard) calculateSixs(dices *DiceSet) *int {
	score := 0
	for _, d := range dices {
		if d.Value == 6 {
			score += d.Value
		}
	}
	return &score
}

func (s *ScoreBoard) calculateChoice(dices *DiceSet) *int {
	score := 0
	for _, d := range dices {
		score += d.Value
	}
	return &score
}

func (s *ScoreBoard) calculateFullHouse(dices *DiceSet) *int {
	diceMap := dices.ToMap()
	isAllLargerThanTwo := func() bool {
		for _, num := range diceMap {
			if num >= 2 {
				continue
			}
			return false
		}
		return true
	}()

	isYacht := func() bool {
		return *s.calculateYacht(dices) > 0
	}()

	score := 0
	if isYacht || (len(diceMap) == 2 && isAllLargerThanTwo) {
		score = *s.calculateChoice(dices)
	}
	return &score
}

func (s *ScoreBoard) calculateFourOfAKind(dices *DiceSet) *int {
	str := dices.String()
	score := 0
	if strings.Contains(str, "1111") ||
		strings.Contains(str, "2222") ||
		strings.Contains(str, "3333") ||
		strings.Contains(str, "4444") ||
		strings.Contains(str, "5555") ||
		strings.Contains(str, "6666") {
		score = *s.calculateChoice(dices)
	}
	return &score
}

func (s *ScoreBoard) calculateSmallStraight(dices *DiceSet) *int {
	diceMap := dices.ToMap()
	score := 0
	if (diceMap[1] > 0 && diceMap[2] > 0 && diceMap[3] > 0 && diceMap[4] > 0) ||
		(diceMap[2] > 0 && diceMap[3] > 0 && diceMap[4] > 0 && diceMap[5] > 0) ||
		(diceMap[3] > 0 && diceMap[4] > 0 && diceMap[5] > 0 && diceMap[6] > 0) {
		score = 15
	}
	return &score
}

func (s *ScoreBoard) calculateLargeStraight(dices *DiceSet) *int {
	diceMap := dices.ToMap()
	score := 0
	if (diceMap[1] > 0 && diceMap[2] > 0 && diceMap[3] > 0 && diceMap[4] > 0 && diceMap[5] > 0) ||
		(diceMap[2] > 0 && diceMap[3] > 0 && diceMap[4] > 0 && diceMap[5] > 0 && diceMap[6] > 0) {
		score = 30
	}

	return &score
}

func (s *ScoreBoard) calculateYacht(dices *DiceSet) *int {
	str := dices.String()
	score := 0
	if strings.Contains(str, "11111") ||
		strings.Contains(str, "22222") ||
		strings.Contains(str, "33333") ||
		strings.Contains(str, "44444") ||
		strings.Contains(str, "55555") ||
		strings.Contains(str, "66666") {
		score = 50
	}
	return &score
}

type ScoreBoardResult struct {
	Ones          *int `json:"aces"`
	Twos          *int `json:"deuces"`
	Threes        *int `json:"threes"`
	Fours         *int `json:"fours"`
	Fives         *int `json:"fives"`
	Sixs          *int `json:"sixes"`
	SubTotal      *int `json:"subtotal"`
	Choice        *int `json:"choice"`
	FourOfAKind   *int `json:"fourKind"`
	FullHouse     *int `json:"fullHouse"`
	SmallStraight *int `json:"smallStraight"`
	LargeStraight *int `json:"largeStraight"`
	Yacht         *int `json:"yacht"`
	Total         *int `json:"total"`
}

func (s *ScoreBoard) calculateSubTotal() int {
	result := 0
	if s.Ones != nil {
		result += *s.Ones
	}
	if s.Twos != nil {
		result += *s.Twos
	}
	if s.Threes != nil {
		result += *s.Threes
	}
	if s.Fours != nil {
		result += *s.Fours
	}
	if s.Fives != nil {
		result += *s.Fives
	}
	if s.Sixs != nil {
		result += *s.Sixs
	}

	if result >= 63 {
		result += 35
	}
	return result
}

func (s *ScoreBoard) calculateTotal() int {
	result := s.calculateSubTotal()
	if s.Choice != nil {
		result += *s.Choice
	}
	if s.FourOfAKind != nil {
		result += *s.FourOfAKind
	}
	if s.FullHouse != nil {
		result += *s.FullHouse
	}
	if s.SmallStraight != nil {
		result += *s.SmallStraight
	}
	if s.LargeStraight != nil {
		result += *s.LargeStraight
	}
	if s.Yacht != nil {
		result += *s.Yacht
	}
	return result
}

func (s *ScoreBoard) ToScoreBoardResult() ScoreBoardResult {
	subTotal := s.calculateSubTotal()
	total := s.calculateTotal()

	return ScoreBoardResult{
		Ones:     s.Ones,
		Twos:     s.Twos,
		Threes:   s.Threes,
		Fours:    s.Fours,
		Fives:    s.Fives,
		Sixs:     s.Sixs,
		SubTotal: &subTotal,

		Choice:        s.Choice,
		FourOfAKind:   s.FourOfAKind,
		FullHouse:     s.FullHouse,
		SmallStraight: s.SmallStraight,
		LargeStraight: s.LargeStraight,
		Yacht:         s.Yacht,
		Total:         &total,
	}
}
