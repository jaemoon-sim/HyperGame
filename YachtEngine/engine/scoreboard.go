package engine

import (
	"fmt"
	"strings"
)

type Score struct {
	Value int  `json:"value"`
	Set   bool `json:"set"`
}

func (s Score) String() string {
	return string(s.Value)
}

type ScoreBoard struct {
	Ones          Score `json:"ones"`
	Twos          Score `json:"twos"`
	Threes        Score `json:"threes"`
	Fours         Score `json:"fours"`
	Fives         Score `json:"fives"`
	Sixs          Score `json:"sixs"`
	Choice        Score `json:"choice'`
	FourOfAKind   Score `json:"four_of_a_kind"`
	FullHouse     Score `json:"full_house`
	SmallStraight Score `json:"small_straight"`
	LargeStraight Score `json:"large_straight"`
	Yacht         Score `json:"yacht"`
}

func (s *ScoreBoard) SetScore(name string, dices *DiceSet) error {
	switch name {
	case "ones":
		s.Ones.Value = s.calculateOnes(dices)
		s.Ones.Set = true
		return nil
	case "twos":
		s.Twos.Value = s.calculateTwos(dices)
		s.Twos.Set = true
		return nil
	case "threes":
		s.Threes.Value = s.calculateThrees(dices)
		s.Threes.Set = true
		return nil
	case "fours":
		s.Fours.Value = s.calculateFours(dices)
		s.Fours.Set = true
		return nil
	case "fives":
		s.Fives.Value = s.calculateFives(dices)
		s.Fives.Set = true
		return nil
	case "sixs":
		s.Sixs.Value = s.calculateSixs(dices)
		s.Sixs.Set = true
		return nil
	case "choice":
		s.Choice.Value = s.calculateChoice(dices)
		s.Choice.Set = true
		return nil
	case "four_of_a_kind":
		s.FourOfAKind.Value = s.calculateFourOfAKind(dices)
		s.FourOfAKind.Set = true
		return nil
	case "full_house":
		s.FullHouse.Value = s.calculateFullHouse(dices)
		s.FullHouse.Set = true
		return nil
	case "small_straight":
		s.SmallStraight.Value = s.calculateSmallStraight(dices)
		s.SmallStraight.Set = true
		return nil
	case "large_straight":
		s.LargeStraight.Value = s.calculateLargeStraight(dices)
		s.LargeStraight.Set = true
		return nil
	case "yacht":
		s.Yacht.Value = s.calculateYacht(dices)
		s.Yacht.Set = true
		return nil
	default:
		return fmt.Errorf("invalid score")
	}
}

func (s *ScoreBoard) calculateOnes(dices *DiceSet) int {
	score := 0
	for d := range dices {
		if d == 1 {
			score += d
		}
	}
	return score
}

func (s *ScoreBoard) calculateTwos(dices *DiceSet) int {
	score := 0
	for d := range dices {
		if d == 2 {
			score += d
		}
	}
	return score
}

func (s *ScoreBoard) calculateThrees(dices *DiceSet) int {
	score := 0
	for d := range dices {
		if d == 3 {
			score += d
		}
	}
	return score
}

func (s *ScoreBoard) calculateFours(dices *DiceSet) int {
	score := 0
	for d := range dices {
		if d == 4 {
			score += d
		}
	}
	return score
}

func (s *ScoreBoard) calculateFives(dices *DiceSet) int {
	score := 0
	for d := range dices {
		if d == 5 {
			score += d
		}
	}
	return score
}

func (s *ScoreBoard) calculateSixs(dices *DiceSet) int {
	score := 0
	for d := range dices {
		if d == 6 {
			score += d
		}
	}
	return score
}

func (s *ScoreBoard) calculateChoice(dices *DiceSet) int {
	score := 0
	for d := range dices {
		score += d
	}
	return score
}

func (s *ScoreBoard) calculateFullHouse(dices *DiceSet) int {
	str := dices.String()
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
		return s.calculateChoice(dices)
	}
	return 0
}

func (s *ScoreBoard) calculateFourOfAKind(dices *DiceSet) int {
	str := dices.String()
	if strings.Contains(str, "1111") ||
		strings.Contains(str, "2222") ||
		strings.Contains(str, "3333") ||
		strings.Contains(str, "4444") ||
		strings.Contains(str, "5555") ||
		strings.Contains(str, "6666") {
		return s.calculateChoice(dices)
	}
	return 0
}

func (s *ScoreBoard) calculateSmallStraight(dices *DiceSet) int {
	str := dices.String()
	if strings.Contains(str, "1234") ||
		strings.Contains(str, "2345") ||
		strings.Contains(str, "3456") {
		return 15
	}
	return 0
}

func (s *ScoreBoard) calculateLargeStraight(dices *DiceSet) int {
	str := dices.String()
	if strings.Contains(str, "12345") ||
		strings.Contains(str, "23456") {
		return 30
	}
	return 0
}

func (s *ScoreBoard) calculateYacht(dices *DiceSet) int {
	str := dices.String()
	if strings.Contains(str, "11111") ||
		strings.Contains(str, "22222") ||
		strings.Contains(str, "33333") ||
		strings.Contains(str, "44444") ||
		strings.Contains(str, "55555") ||
		strings.Contains(str, "66666") {
		return 50
	}
	return 0
}
