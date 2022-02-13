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
	str := dices.String()
	score := 0
	if (strings.Contains(str, "11") &&
		(strings.Contains(str, "111") ||
			strings.Contains(str, "222") ||
			strings.Contains(str, "333") ||
			strings.Contains(str, "444") ||
			strings.Contains(str, "555") ||
			strings.Contains(str, "666"))) ||
		(strings.Contains(str, "22") &&
			(strings.Contains(str, "111") ||
				strings.Contains(str, "222") ||
				strings.Contains(str, "333") ||
				strings.Contains(str, "444") ||
				strings.Contains(str, "555") ||
				strings.Contains(str, "666"))) ||
		(strings.Contains(str, "33") &&
			(strings.Contains(str, "111") ||
				strings.Contains(str, "222") ||
				strings.Contains(str, "333") ||
				strings.Contains(str, "444") ||
				strings.Contains(str, "555") ||
				strings.Contains(str, "666"))) ||
		(strings.Contains(str, "44") &&
			(strings.Contains(str, "111") ||
				strings.Contains(str, "222") ||
				strings.Contains(str, "333") ||
				strings.Contains(str, "444") ||
				strings.Contains(str, "555") ||
				strings.Contains(str, "666"))) ||
		(strings.Contains(str, "55") &&
			(strings.Contains(str, "111") ||
				strings.Contains(str, "222") ||
				strings.Contains(str, "333") ||
				strings.Contains(str, "444") ||
				strings.Contains(str, "555") ||
				strings.Contains(str, "666"))) ||
		(strings.Contains(str, "66") &&
			(strings.Contains(str, "111") ||
				strings.Contains(str, "222") ||
				strings.Contains(str, "333") ||
				strings.Contains(str, "444") ||
				strings.Contains(str, "555") ||
				strings.Contains(str, "666"))) {
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
	str := dices.String()
	score := 0
	if strings.Contains(str, "1234") ||
		strings.Contains(str, "2345") ||
		strings.Contains(str, "3456") {
		score = 15
	}
	return &score
}

func (s *ScoreBoard) calculateLargeStraight(dices *DiceSet) *int {
	str := dices.String()
	score := 0
	if strings.Contains(str, "12345") ||
		strings.Contains(str, "23456") {
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
