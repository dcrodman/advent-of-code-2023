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
	times, distances := parseInputPart1("input.txt")
	fmt.Println("part1 # possibilities:", calculatePossibilities(times, distances))

	times, distances = parseInputPart2("input.txt")
	fmt.Println("part2 # possibilities:", calculatePossibilities(times, distances))
}

func parseInputPart1(filename string) (times []int, distances []int) {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		for _, v := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			val, err := strconv.Atoi(v)
			if err != nil {
				continue
			}

			if strings.Contains(scanner.Text(), "Time") {
				times = append(times, val)
			} else {
				distances = append(distances, val)
			}
		}
	}
	return
}

func parseInputPart2(filename string) (times []int, distances []int) {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		smushed := strings.ReplaceAll(parts[1], " ", "")

		val, err := strconv.Atoi(smushed)
		if err != nil {
			continue
		}

		if strings.Contains(scanner.Text(), "Time") {
			times = append(times, val)
		} else {
			distances = append(distances, val)
		}
	}
	return
}

func calculatePossibilities(times []int, distances []int) int {
	totalPossibilities := 1
	for i := 0; i < len(times); i++ {
		raceTime := times[i]
		raceDistance := distances[i]

		possibilities := 0
		for t := 1; t < raceTime; t++ {
			traveled := t * (raceTime - t)
			if traveled > raceDistance {
				possibilities++
			}
		}
		totalPossibilities *= possibilities
	}
	return totalPossibilities
}
