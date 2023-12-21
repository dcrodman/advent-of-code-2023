package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1(input)
	part2(input)
}

func part1(input []byte) {
	totalHash := 0
	for _, step := range strings.Split(string(input), ",") {
		totalHash += computeHash(step)
	}
	fmt.Println("part1 hash:", totalHash)
}

func computeHash(label string) (hash int) {
	for _, r := range label {
		if string(r) == "\n" {
			continue
		}
		hash += int(r)
		hash *= 17
		hash %= 256
	}
	return
}

type lens struct {
	label string
	value int
}

func part2(input []byte) {
	boxes := make([][]lens, 256)

	for _, step := range strings.Split(string(input), ",") {
		step = strings.TrimSpace(step)

		if strings.Contains(step, "=") {
			parts := strings.Split(step, "=")
			label := parts[0]
			value, _ := strconv.Atoi(parts[1])

			exists := false
			b := computeHash(label)
			for i, l := range boxes[b] {
				if l.label == label {
					boxes[b][i] = lens{label: l.label, value: value}
					exists = true
					break
				}
			}
			if !exists {
				boxes[b] = append(boxes[b], lens{label: label, value: value})
			}

		} else {
			label := strings.TrimSuffix(step, "-")
			b := computeHash(label)

			for i, box := range boxes[b] {
				if box.label == label {
					boxes[b] = slices.Delete(boxes[b], i, i+1)
				}
			}
		}
	}

	focusingPower := 0
	for i, box := range boxes {
		for slot, lens := range box {
			focusingPower += (i + 1) * (slot + 1) * lens.value
		}
	}
	fmt.Println("part2 focusing power:", focusingPower)
}
