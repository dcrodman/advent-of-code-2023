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

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		return
	}

	var part1Sum int64
	var part2Sum int64

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		sequenceParts := strings.Split(scanner.Text(), " ")

		baseSequence := make([]int64, len(sequenceParts))
		for i, s := range sequenceParts {
			val, _ := strconv.ParseInt(s, 10, 64)
			baseSequence[i] = val
		}

		sequences := make([][]int64, 0)
		sequences = append(sequences, baseSequence)

		part1Sum += extrapolate(sequences)
		slices.Reverse(sequences[0])
		part2Sum += extrapolate(sequences)
	}

	fmt.Println("part1 sum:", part1Sum)
	fmt.Println("part2 sum:", part2Sum)
}

func extrapolate(sequences [][]int64) int64 {
	// Build the sequences top-down.
	for {
		s := sequences[len(sequences)-1]
		if allZero(s) {
			break
		}

		// Compute the differences for the next sequence.
		nextSequence := make([]int64, len(s)-1)
		for i := 0; i < len(nextSequence); i++ {
			nextSequence[i] = s[i+1] - s[i]
		}
		sequences = append(sequences, nextSequence)
	}

	// Add the first placeholder value to the end.
	sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)

	// Now rebuild them bottom-up.
	for i := len(sequences) - 2; i >= 0; i-- {
		nextVal := sequences[i][len(sequences[i])-1] + sequences[i+1][len(sequences[i+1])-1]
		sequences[i] = append(sequences[i], nextVal)
	}

	return sequences[0][len(sequences[0])-1]
}

func allZero(s []int64) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
