package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		return
	}

	part1(input)
	// part2(input)

	inputstr := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		inputstr = append(inputstr, scanner.Text())
	}
	fmt.Println(Solve(inputstr))
}

// How many points are they worth in total?
func part1(input []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	totalScore := 0
	for scanner.Scan() {
		cardParts := strings.Split(scanner.Text(), ":")
		numbers := strings.Split(cardParts[1], "|")

		winningNumbers := strings.TrimSpace(numbers[0])
		ourNumbers := strings.TrimSpace(numbers[1])

		// Build up our set of winning numbers.
		winnersSet := make(map[int64]struct{})
		for _, winningNumber := range strings.Split(winningNumbers, " ") {
			wn := strings.TrimSpace(winningNumber)
			val, _ := strconv.ParseInt(wn, 10, 64)
			winnersSet[val] = struct{}{}
		}

		cardScore := 0
		// Roll our own set intersection to walk through the set of numbers we
		// have and count up the number of matches.
		for _, ourNumber := range strings.Split(ourNumbers, " ") {
			on := strings.TrimSpace(ourNumber)
			val, err := strconv.ParseInt(on, 10, 64)
			if err != nil {
				// Single digits are padded with a space, so the split may leave us with
				// empty strings. It's safe to just skip those.
				continue
			}

			if _, ok := winnersSet[val]; ok {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}

		totalScore += cardScore
	}

	fmt.Println("total score for the card:", totalScore)
}

// func part2(input []byte) {
// 	scanner := bufio.NewScanner(bytes.NewReader(input))

// 	totalCards := make(map[int]int)

// 	currentCard := 1
// 	for scanner.Scan() {
// 		card := scanner.Text()
// 		wins := getNumCardsWon(card)

// 		totalCards[currentCard] = totalCards[currentCard] + 1
// 		for i := currentCard; i <= currentCard+wins; i++ {
// 			totalCards[i] = totalCards[i] + 1
// 		}

// 		currentCard++
// 	}

// 	sum := 0
// 	for _, total := range totalCards {
// 		sum += total
// 	}

// 	fmt.Println(totalCards)
// 	fmt.Println(sum)
// }

// Stolen from https://www.reddit.com/r/adventofcode/comments/18actmy/comment/kc21aq1/?utm_source=share&utm_medium=web2x&context=3
// to try to figure out how the hell this works.
func Solve(lines []string) (part2 int) {
	counts := make(map[int]int)
	for l, line := range lines {
		// We've seen a card, increment the total
		part2++
		// Get the card id/value
		card := l + 1
		// Increase the number of times we've seen this card again I guess?
		counts[card]++
		// Now get the number of times we've hit this card
		count := counts[card]

		copies := getNumCardsWon(line)
		// Now traverse the copies of the card, increasing their totals by the
		// number of times we've seen this card
		for x := 1; x <= copies; x++ {
			counts[card+x] += count
			part2 += count
		}
		delete(counts, card)
	}
	return part2
}

// func part2(input []byte) {
// 	scanner := bufio.NewScanner(bytes.NewReader(input))
//	cardRegex := regexp.MustCompile(`Card\s+(?P<num>\d).*`)

// 	// Read in all of the cards so that we can access them by index.
// 	var cardNumber int = 1
// 	copiesPerCard := make(map[int][]int)

// 	for scanner.Scan() {
// 		numWins := getNumCardsWon(scanner.Text())

// 		// Store the copies that row wins. In the example, this
// 		// would be {1: [2, 3, 4, 5], ...}
// 		copiesPerCard[cardNumber] = make([]int, numWins)
// 		i := 0
// 		for cardCopy := cardNumber + 1; cardCopy <= cardNumber+int(numWins); cardCopy++ {
// 			copiesPerCard[cardNumber][cardCopy] = cardCopy
// 			i++
// 		}

// 		cardNumber++
// 	}

// 	totalCards := make(map[int]int)
// 	for card, copies := range copiesPerCard {
// 		totalCards[card] = totalCards[card] + 1

// 		for _, copy := range copies {
// 			q := make([]int, 1)
// 			q[0] = copy

// 			for len(q) > 0 {
// 				var next int
// 				next, q = q[0], q[1:]

// 				totalCards[next] = totalCards[next] + 1

// 			}
// 		}
// 	}
// }

func getNumCardsWon(card string) int {
	cardParts := strings.Split(card, ":")
	numbers := strings.Split(cardParts[1], "|")

	winningNumbers := strings.TrimSpace(numbers[0])
	ourNumbers := strings.TrimSpace(numbers[1])

	// Build up our set of winning numbers.
	winnersSet := make(map[int64]struct{})
	for _, winningNumber := range strings.Split(winningNumbers, " ") {
		wn := strings.TrimSpace(winningNumber)
		val, _ := strconv.ParseInt(wn, 10, 64)
		winnersSet[val] = struct{}{}
	}

	numbersWon := 0
	// Walk through the set of numbers we have and count how many we won.
	for _, ourNumber := range strings.Split(ourNumbers, " ") {
		on := strings.TrimSpace(ourNumber)
		val, err := strconv.ParseInt(on, 10, 64)
		if err != nil {
			// Single digits are padded with a space, so the split may leave us with
			// empty strings. It's safe to just skip those.
			continue
		}
		if _, ok := winnersSet[val]; ok {
			numbersWon++
		}
	}

	return numbersWon
}
