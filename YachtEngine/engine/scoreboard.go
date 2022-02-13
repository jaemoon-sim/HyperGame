package engine

import (
	"fmt"
	"strings"
)

type ScoreBoard struct {
	Ones          *int `json:"ones"`
	Twos          *int `json:"twos"`
	Threes        *int `json:"threes"`
	Fours         *int `json:"fours"`
	Fives         *int `json:"fives"`
	Sixs          *int `json:"sixs"`
	Choice        *int `json:"choice'`
	FourOfAKind   *int `json:"four_of_a_kind"`
	FullHouse     *int `json:"full_house`
	SmallStraight *int `json:"small_straight"`
	LargeStraight *int `json:"large_straight"`
	Yacht         *int `json:"yacht"`
}

func (s *ScoreBoard) SetScore(name string, dices *DiceSet) error {
	switch name {
	case "ones":
		s.Ones = s.calculateOnes(dices)
		return nil
	case "twos":
		s.Twos = s.calculateTwos(dices)
		return nil
	case "threes":
		s.Threes = s.calculateThrees(dices)
		return nil
	case "fours":
		s.Fours = s.calculateFours(dices)
		return nil
	case "fives":
		s.Fives = s.calculateFives(dices)
		return nil
	case "sixs":
		s.Sixs = s.calculateSixs(dices)
		return nil
	case "choice":
		s.Choice = s.calculateChoice(dices)
		return nil
	case "four_of_a_kind":
		s.FourOfAKind = s.calculateFourOfAKind(dices)
		return nil
	case "full_house":
		s.FullHouse = s.calculateFullHouse(dices)
		return nil
	case "small_straight":
		s.SmallStraight = s.calculateSmallStraight(dices)
		return nil
	case "large_straight":
		s.LargeStraight = s.calculateLargeStraight(dices)
		return nil
	case "yacht":
		s.Yacht = s.calculateYacht(dices)
		return nil
	default:
		return fmt.Errorf("invalid score")
	}
}

func (s *ScoreBoard) calculateOnes(dices *DiceSet) *int {
	score := 0
	for d := range dices {
		if d == 1 {
			score += d
		}
	}
	return &score
}

func (s *ScoreBoard) calculateTwos(dices *DiceSet) *int {
	score := 0
	for d := range dices {
		if d == 2 {
			score += d
		}
	}
	return &score
}

func (s *ScoreBoard) calculateThrees(dices *DiceSet) *int {
	score := 0
	for d := range dices {
		if d == 3 {
			score += d
		}
	}
	return &score
}

func (s *ScoreBoard) calculateFours(dices *DiceSet) *int {
	score := 0
	for d := range dices {
		if d == 4 {
			score += d
		}
	}
	return &score
}

func (s *ScoreBoard) calculateFives(dices *DiceSet) *int {
	score := 0
	for d := range dices {
		if d == 5 {
			score += d
		}
	}
	return &score
}

func (s *ScoreBoard) calculateSixs(dices *DiceSet) *int {
	score := 0
	for d := range dices {
		if d == 6 {
			score += d
		}
	}
	return &score
}

func (s *ScoreBoard) calculateChoice(dices *DiceSet) *int {
	score := 0
	for d := range dices {
		score += d
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
