// https://adventofcode.com/2023/day/2
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var gameRegex = regexp.MustCompile(`Game (?P<id>\d+): (?P<game>.+)`)

var diceInBag = map[string]int64{
	"red": 12, "green": 13, "blue": 14,
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		return
	}

	part1(input)
	part2(input)
}

// Determine which games would have been possible if the bag had been loaded with only
// 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games
func part1(input []byte) {
	var validGameIdsSum int64 = 0

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		matches := gameRegex.FindStringSubmatch(scanner.Text())
		gameId, _ := strconv.ParseInt(matches[1], 10, 8)
		game := matches[2]
		valid := true

		for _, subset := range strings.Split(game, ";") {
			subset = strings.TrimSpace(subset)

			for _, dice := range strings.Split(subset, ",") {
				diceSplit := strings.Split(strings.TrimSpace(dice), " ")

				count, _ := strconv.ParseInt(diceSplit[0], 10, 8)
				color := diceSplit[1]
				if count > diceInBag[color] {
					valid = false
					break
				}
			}
		}

		if valid {
			fmt.Printf("Game %d is valid\n", gameId)
			validGameIdsSum += gameId
		}
	}

	fmt.Printf("sum of valid game IDs: %d\n", validGameIdsSum)
}

// For each game, find the minimum set of cubes that must have been present.
// What is the sum of the power of these sets?
func part2(input []byte) {
	var sumOfGamePowers int64

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		var (
			line            = scanner.Text()
			minDicePerColor = make(map[string]int64)
		)

		matches := gameRegex.FindStringSubmatch(line)
		for _, subset := range strings.Split(matches[2], ";") {
			subset = strings.TrimSpace(subset)

			for _, dice := range strings.Split(subset, ",") {
				diceSplit := strings.Split(strings.TrimSpace(dice), " ")

				count, _ := strconv.ParseInt(diceSplit[0], 10, 8)
				color := diceSplit[1]

				if min, ok := minDicePerColor[color]; !ok || count > min {
					minDicePerColor[color] = count
				}
			}
		}

		var power int64 = 1
		for _, count := range minDicePerColor {
			power *= count
		}

		sumOfGamePowers += power
	}

	fmt.Printf("sum of powers across valid game IDs: %d\n", sumOfGamePowers)
}
