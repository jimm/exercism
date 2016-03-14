package poker

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const testVersion = 3

var legalSuits = map[string]bool {"♡": true, "♧": true, "♤": true, "♢": true}
var values = map[string]int {
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"10": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type card struct {
	suit string
	val int
}

type hand []card

// - Straight flush: All cards in the same suit, and in sequence
// - Four of a kind: Four of the cards have the same rank
// - Full House: Three cards of one rank, the other two of another rank
// - Flush: All cards in the same suit
// - Straight: All cards in sequence (aces can be high or low, but not both at once)
// - Three of a kind: Three of the cards have the same rank
// - Two pair: Two pairs of cards have the same rank
// - Pair: Two cards have the same rank
// - High card: None of the above conditions are met

func BestHand(handStrings []string) ([]string, error) {
	hands := []hand{}
	for _, handString := range handStrings {
		hand, err := stringToHand(handString)
		if err != nil {
			return nil, err
		}
		hands = append(hands, hand)
	}

	winner := 0
	for i := 1; i < len(hands); i++ {
		if better(hands[winner], hands[i]) {
			winner = i
		}
	}
	return []string{handStrings[winner]}, nil
}

func stringToHand(handString string) (hand, error) {
	hand := []card{}
	fields := strings.Fields(handString)
	for _, field := range fields {
		fmt.Println("field", field)
		valStr := field[:len(field)-2]
		fmt.Println("valStr", valStr)
		val, ok := values[valStr]
		if !ok {
			return nil, errors.New(fmt.Sprintf("no such card value: %s", field))
		} else {				// DEBUG
			fmt.Println("OK card value:", valStr)
		}

		suitStr := field[len(field)-1:]
		_, ok = legalSuits[suitStr]
		if !ok {
			return nil, errors.New(fmt.Sprintf("no such suit: %s", field))
		}
		hand = append(hand, card{suitStr, val})
	}
	return hand, nil
}

// better returns true if that hand is better than this hand.
func better(this hand, that hand) bool {
	rThis, thisVal := rank(this)
	rThat, thatVal := rank(that)
	if rThat > rThis {
		return true
	} else if rThat < rThis {
		return false
	}
	return thatVal > thisVal
}

// rank returns two ints: the hand ranking (higher numbers are
// better) and the high value for that rank (for example, the high card in a
// straight or the higher pair of a two-pair hand).
func rank(h hand) (int, int) {
	sortHandByValue(h)
	if val, matches := straightFlush(h); matches {
		return 9, val
	}
	if val, matches := fourOfAKind(h); matches {
		return 8, val
	}
	if val, matches := fullHouse(h); matches {
		return 7, val
	}
	if val, matches := flush(h); matches {
		return 6, val
	}
	if val, matches := straight(h); matches {
		return 5, val
	}
	if val, matches := threeOfAKind(h); matches {
		return 4, val
	}
	if val, matches := twoPair(h); matches {
		return 3, val
	}
	if val, matches := pair(h); matches {
		return 2, val
	}
	return 1, highCard(h).val
}


func straightFlush(h hand) (int, bool) {
    val := 0
	return val, true
}

func fourOfAKind(h hand) (int, bool) {
    val := 0
	return val, true
}

func fullHouse(h hand) (int, bool) {
    val := 0
	return val, true
}

func flush(h hand) (int, bool) {
    val := 0
	return val, true
}

func straight(h hand) (int, bool) {
    val := 0
	return val, true
}

func threeOfAKind(h hand) (int, bool) {
    val := 0
	return val, true
}

func twoPair(h hand) (int, bool) {
    val := 0
	return val, true
}

func pair(h hand) (int, bool) {
    val := 0
	return val, true
}

func highCard(h hand) card {
	hc := card{val: -1}
	for _, card := range h {
		if card.val > hc.val {
			hc = card
		}
	}
	return hc
}

// **************** sorting ****************

type cardValueSort []card

func (cs cardValueSort) Len() int { return len(cs) }

func (cs cardValueSort) Swap(i, j int) { cs[i], cs[j] = cs[j], cs[i] }

func (card cardValueSort) Less(i, j int) bool {
	return card[i].val < card[j].val
}

func sortHandByValue(hand hand) {
	sort.Sort(cardValueSort(hand))
}
