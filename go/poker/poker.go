package poker

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const testVersion = 3

var legalSuits = map[string]bool{"♡": true, "♧": true, "♤": true, "♢": true}
var values = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

type card struct {
	suit string
	rank int
}

type hand struct {
	handType, bestRank int
	cards              []card
}

func BestHand(handStrings []string) ([]string, error) {
	hands := []hand{}
	for _, handString := range handStrings {
		h, err := stringToHand(handString) // also analyzes the hand
		if err != nil {
			return nil, err
		}
		hands = append(hands, h)
	}

	winners := []string{handStrings[0]}
	winner := 0
	for i := 1; i < len(hands); i++ {
		comparison := better(hands[winner], hands[i])
		if comparison == 1 {
			winners = []string{handStrings[i]}
			winner = i
		} else if comparison == 0 {
			winners = append(winners, handStrings[i])
		}
	}
	return winners, nil
}

// stringToHand translates handString into a hand and analyzes the hand,
// setting handType and bestRank.
func stringToHand(handString string) (hand, error) {
	h := hand{}
	fields := strings.Fields(handString)
	for _, field := range fields {
		rankStr := field[0:1]
		suit := field[1:]
		if rankStr == "1" && (field[1:2] == "0" || field[1:2] == "1") {
			rankStr = field[0:2]
			suit = field[2:]
		}

		rank, ok := values[rankStr]
		if !ok {
			return hand{}, errors.New(fmt.Sprintf("no such card value: %s", field))
		}

		_, ok = legalSuits[suit]
		if !ok {
			return hand{}, errors.New(fmt.Sprintf("no such suit: %s", field))
		}
		h.cards = append(h.cards, card{suit, rank})
	}
	if len(h.cards) != 5 {
		return hand{}, errors.New(fmt.Sprintf("only %d cards in hand; need 5", len(h.cards)))
	}
	handType, bestRank := analyze(h)
	h.handType = handType
	h.bestRank = bestRank
	return h, nil
}

// better returns 1 if that hand is better than this hand, -1 if this hand
// is better than that hand, and 0 if they are equal in every way.
func better(this hand, that hand) int {
	if that.handType > this.handType {
		return 1
	} else if that.handType < this.handType {
		return -1
	}
	if that.bestRank > this.bestRank {
		return 1
	}
	if that.bestRank < this.bestRank {
		return -1
	}

	for i := 4; i >= 0; i-- {
		if that.cards[i].rank > this.cards[i].rank {
			return 1
		} else if that.cards[i].rank < this.cards[i].rank {
			return -1
		}
	}
	return 0
}

// analyze calculates and returns the handType and bestRank of a hand.
func analyze(h hand) (int, int) {
	sortHandByRank(h)
	if rank, matches := straightFlush(h); matches {
		return 9, rank
	}
	if rank, matches := fourOfAKind(h); matches {
		return 8, rank
	}
	if rank, matches := fullHouse(h); matches {
		return 7, rank
	}
	if rank, matches := flush(h); matches {
		return 6, rank
	}
	if rank, matches := straight(h); matches {
		return 5, rank
	}
	if rank, matches := threeOfAKind(h); matches {
		return 4, rank
	}
	if rank, matches := twoPair(h); matches {
		return 3, rank
	}
	if rank, matches := pair(h); matches {
		return 2, rank
	}
	return 1, highestCard(h).rank
}

// highestCard returns the card with the highest rank.
func highestCard(h hand) card {
	return h.cards[4]
}

// **************** hands ****************

// - Straight flush: All cards in the same suit, and in sequence
func straightFlush(h hand) (int, bool) {
	if rank, isStraight := straight(h); isStraight {
		if _, isFlush := flush(h); isFlush {
			return rank, true
		}
	}
	return 0, false
}

// - Four of a kind: Four of the cards have the same rank
func fourOfAKind(h hand) (int, bool) {
	return ofAKind(h, 4)
}

// - Full House: Three cards of one rank, the other two of another rank.
// Value returned is rank of set of three.
func fullHouse(h hand) (int, bool) {
	m := rankMap(h)
	if len(m) != 2 {
		return 0, false
	}
	var twoRank, threeRank int
	for _, cards := range m {
		if len(cards) == 2 {
			twoRank = cards[0].rank
		} else if len(cards) == 3 {
			threeRank = cards[0].rank
		}
	}
	if twoRank == 0 || threeRank == 0 {
		return 0, false
	}
	return threeRank, true
}

// - Flush: All cards in the same suit
func flush(h hand) (int, bool) {
	for i := 1; i < 5; i++ {
		if h.cards[0].suit != h.cards[i].suit {
			return 0, false
		}
	}
	return h.cards[4].rank, true
}

// - Straight: All cards in sequence (aces can be high or low, but not both at once)
func straight(h hand) (int, bool) {
	if h.cards[4].rank == 14 && h.cards[0].rank == 2 { // try ace low
		if h.cards[1].rank == 3 && h.cards[2].rank == 4 && h.cards[3].rank == 5 {
			return 5, true
		}
	}
	for i := 0; i < 4; i++ {
		if h.cards[i].rank+1 != h.cards[i+1].rank {
			return 0, false
		}
	}
	return h.cards[4].rank, true
}

// - Three of a kind: Three of the cards have the same rank
func threeOfAKind(h hand) (int, bool) {
	return ofAKind(h, 3)
}

// - Two pair: Two pairs of cards have the same rank
func twoPair(h hand) (int, bool) {
	for i := 0; i < 4; i++ {
		if h.cards[i].rank == h.cards[i+1].rank {
			for j := i + 2; j < 4; j++ {
				if h.cards[j].rank == h.cards[j+1].rank {
					return max(h.cards[i].rank, h.cards[j].rank), true
				}
			}
		}
	}
	return 0, false
}

// - Pair: Two cards have the same rank
func pair(h hand) (int, bool) {
	for i := 0; i < 3; i++ {
		if h.cards[i].rank == h.cards[i+1].rank {
			return h.cards[i].rank, true
		}
	}
	return 0, false
}

// - High card: None of the above conditions are met
func highCard(h hand) (int, bool) {
	return highestCard(h).rank, true
}

func ofAKind(h hand, num int) (int, bool) {
	m := rankMap(h)
	for _, cards := range m {
		if len(cards) == num {
			return highestCard(h).rank, true
		}
	}
	return 0, false
}

// rankMap returns a map from rank to cards in the hand.
func rankMap(h hand) map[int][]card {
	m := map[int][]card{}
	for i := 0; i < 5; i++ {
		_, found := m[h.cards[i].rank]
		if !found {
			m[h.cards[i].rank] = []card{h.cards[i]}
		} else {
			m[h.cards[i].rank] = append(m[h.cards[i].rank], h.cards[i])
		}
	}
	return m
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// **************** sorting ****************

type cardRankSort hand

func (cs cardRankSort) Len() int { return len(cs.cards) }

func (cs cardRankSort) Swap(i, j int) { cs.cards[i], cs.cards[j] = cs.cards[j], cs.cards[i] }

func (cs cardRankSort) Less(i, j int) bool {
	return cs.cards[i].rank < cs.cards[j].rank
}

func sortHandByRank(hand hand) {
	sort.Sort(cardRankSort(hand))
}
