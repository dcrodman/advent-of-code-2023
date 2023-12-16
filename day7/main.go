package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	bid   int
	rank  int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var part1HandRankings string = "23456789TJQKA"
	fmt.Println("part1 score:", play(input, part1HandRankings, evaluateStrength))
	var part2HandRankings string = "J23456789TQKA"
	fmt.Println("part2 score:", play(input, part2HandRankings, evaluateStrengthJokers))
}

func play(input []byte, cardRankings string, evaluator evaluator) int {
	hands := make([]hand, 0)

	// Read all of the cards and determine their rank.
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(parts[1])
		h := hand{
			cards: parts[0],
			bid:   bid,
			rank:  evaluator(parts[0]),
		}
		hands = append(hands, h)
	}

	// Now sort them by rank + the value of individual cards.
	slices.SortFunc(hands, func(a, b hand) int {
		if a.rank < b.rank {
			return -1
		} else if a.rank > b.rank {
			return 1
		}
		return compare(cardRankings, a, b)
	})

	totalWinnings := 0
	for i, h := range hands {
		totalWinnings += h.bid * (i + 1)
	}
	return totalWinnings
}

type evaluator func(string) int

func evaluateStrength(cards string) int {
	cardsByLabel := make(map[rune]int)
	for _, card := range cards {
		cardsByLabel[card] = cardsByLabel[card] + 1
	}
	distribution := make([]int, 0)
	for _, v := range cardsByLabel {
		distribution = append(distribution, v)
	}
	slices.Sort(distribution)

	switch {
	// five of a kind  = AAAAA -> labels = 1, dist = [5]
	case slices.Equal(distribution, []int{5}):
		return 7
	// four of a kind  = AAAAB -> labels = 2, dist = [4, 1]
	case slices.Equal(distribution, []int{1, 4}):
		return 6
	// full house 	   = AAABB -> labels = 2, dist = [3, 2]
	case slices.Equal(distribution, []int{2, 3}):
		return 5
	// three of a kind = AAABC -> labels = 3, dist = [3, 1, 1]
	case slices.Equal(distribution, []int{1, 1, 3}):
		return 4
	// two pair        = AABBC -> labels = 3, dist = [2, 2, 1]
	case slices.Equal(distribution, []int{1, 2, 2}):
		return 3
	// one pair        = AABCD -> labels = 4, dist = [2, 1, 1, 1]
	case slices.Equal(distribution, []int{1, 1, 1, 2}):
		return 2
	}
	// high card       = ABCDE -> labels = 5, dist = [1, 1, 1, 1, 1]
	return 1
}

func evaluateStrengthJokers(cards string) int {
	maxRanking := 0
	// This feels very brute force but...eh.
	for _, card := range "23456789TJQKA" {
		replaced := strings.ReplaceAll(cards, "J", string(card))
		ranking := evaluateStrength(replaced)
		if ranking > maxRanking {
			maxRanking = ranking
		}
	}
	return maxRanking
}

func compare(rankings string, hand1, hand2 hand) int {
	for i := range hand1.cards {
		hand1CardRank := strings.Index(rankings, string(hand1.cards[i]))
		hand2CardRank := strings.Index(rankings, string(hand2.cards[i]))

		if hand1CardRank < hand2CardRank {
			return -1
		} else if hand1CardRank > hand2CardRank {
			return 1
		}
	}
	return 0
}
