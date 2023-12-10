package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		return
	}

	seeds, maps := parseMap(input)
	part1(seeds, maps)
	part2(seeds, maps)
}

// What is the lowest location number that corresponds to any of the initial seed numbers?
func part1(seeds []int64, maps [][]mapEntry) {
	var minLocation int64 = math.MaxInt64
	for _, seed := range seeds {
		location := descend(seed, maps)
		if location < minLocation {
			minLocation = location
		}
	}

	fmt.Println("part1: minimum value is:", minLocation)
}

// Consider all of the initial seed numbers listed in the ranges on the first line of the almanac.
// What is the lowest location number that corresponds to any of the initial seed numbers?
func part2(seeds []int64, maps [][]mapEntry) {
	var minLocation int64 = math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		for s := seeds[i]; s < seeds[i]+seeds[i+1]; s++ {
			location := descend(s, maps)
			if location < minLocation {
				minLocation = location
			}
		}
	}

	fmt.Println("part2: minimum value is:", minLocation)
}

type mapEntry struct {
	destRangeStart    int64
	sourceRangeStart  int64
	sourceRangeLength int64
}

func parseMap(input []byte) ([]int64, [][]mapEntry) {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	seeds := make([]int64, 0)
	maps := make([][]mapEntry, 0)

	readMapLine := false
	var mapLine []mapEntry
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "seeds") {
			seedLine := strings.Split(scanner.Text(), ":")
			seedStrs := strings.Split(seedLine[1], " ")

			for _, s := range seedStrs {
				if s != "" {
					val, _ := strconv.ParseInt(s, 10, 64)
					seeds = append(seeds, val)
				}
			}
			continue
		}

		if strings.Contains(scanner.Text(), "map") {
			readMapLine = true
			continue
		} else if scanner.Text() == "" {
			readMapLine = false
			if len(mapLine) > 0 {
				maps = append(maps, mapLine)
			}
			mapLine = nil
		}

		if readMapLine {
			parsed := make([]int64, 3)
			for i, m := range strings.Split(scanner.Text(), " ") {
				val, _ := strconv.ParseInt(m, 10, 64)
				parsed[i] = val
			}
			mapLine = append(mapLine, mapEntry{parsed[0], parsed[1], parsed[2]})
		}
	}
	maps = append(maps, mapLine)

	return seeds, maps
}

func descend(currentValue int64, maps [][]mapEntry) int64 {
	// Bottom out after the final location map.
	if len(maps) == 0 {
		return currentValue
	}
	entries, maps := maps[0], maps[1:]

	for _, entry := range entries {
		rangeEnd := entry.sourceRangeStart + entry.sourceRangeLength
		if entry.sourceRangeStart <= currentValue && currentValue < rangeEnd {
			nextValue := entry.destRangeStart + (currentValue - entry.sourceRangeStart)
			return descend(nextValue, maps)
		}
	}
	// No matches; descend to the next map with the current value.
	return descend(currentValue, maps)
}
