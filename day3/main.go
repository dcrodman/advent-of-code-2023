package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		return
	}

	part1(input)
	part2(input)
}

// What is the sum of all of the part numbers in the engine schematic?
func part1(input []byte) {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	// Rune-ify the entire schematic.
	engine := make([][]rune, 0)
	for scanner.Scan() {
		engineLine := make([]rune, 0)
		for _, r := range scanner.Text() {
			engineLine = append(engineLine, r)
		}
		engine = append(engine, engineLine)
	}

	var sumPartNumbers int64 = 0

	// Now walk it again looking for integers, checking each digit character
	// to see if there are any adjacent symbols. If so, parse and add it.
	for row := 0; row < len(engine); row++ {
		for col := 0; col < len(engine[row]); {
			var (
				count bool
				sb    strings.Builder
			)
			for col < len(engine[row]) && unicode.IsDigit(engine[row][col]) {
				sb.WriteRune(engine[row][col])

				// Check in a box around the digit if we haven't found a symbol yet.
				if !count {
					for i := row - 1; i <= row+1; i++ {
						if i < 0 || i >= len(engine[row]) {
							continue
						}

						for j := col - 1; j <= col+1; j++ {
							if j < 0 || j >= len(engine) {
								continue
							}
							count = count || (engine[i][j] != '.' && !unicode.IsDigit(engine[i][j]))
						}
					}
				}
				col += 1
			}

			if count {
				val, _ := strconv.ParseInt(sb.String(), 10, 64)
				sumPartNumbers += val
			}
			col += 1
		}
	}

	fmt.Println("sum of the part numbers in the engine schematic: ", sumPartNumbers)
}

type gear struct {
	row int
	col int
}

func part2(input []byte) {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	// The first key is x, the second key is y, and the value is the integer
	// that overlaps those coordinates.
	engine := make(map[int]map[int]int64)
	gears := make([]gear, 0)

	// Walk through the schematic and build up our insane little map of maps.
	row := 0
	for scanner.Scan() {
		engine[row] = make(map[int]int64)
		digitsStartCol := 0
		digits := make([]string, 0)

		col := 0
		stringRunes := []rune(scanner.Text())
		for i := range stringRunes {
			r := stringRunes[i]
			if unicode.IsDigit(r) {
				if digitsStartCol == 0 {
					digitsStartCol = col
				}
				digits = append(digits, string(r))
				col++

				// Stupid hack, but fall through on the last iteration so we don't drop a number.
				if i < len(stringRunes)-1 {
					continue
				}
			}

			if r == '*' {
				gears = append(gears, gear{row, col})
			}
			if len(digits) > 0 {
				num := strings.Join(digits, "")
				val, _ := strconv.ParseInt(num, 10, 64)

				for i := digitsStartCol; i <= col; i++ {
					engine[row][i] = val
				}
				digitsStartCol = 0
				digits = nil
			}
			engine[row][col] = -1
			col++
		}
		row += 1
	}

	var sumRatios uint64 = 0

	// Now, run through the gears again and reference the map to check if
	// any diagonal coordinates correspond to parts.
	for _, gear := range gears {
		adjacentParts := make([]int64, 0)
		for i := gear.row - 1; i <= gear.row+1; i++ {
			engineRow, ok := engine[i]
			if !ok {
				continue
			}
			for j := gear.col - 1; j <= gear.col+1; j++ {
				part, ok := engineRow[j]
				if gear.row == 112 && gear.col == 136 {
					fmt.Printf("looking at [%d][%d] ok=%v, part=%v\n", i, j, ok, part)
				}
				if !ok || part < 0 {
					continue
				}
				if len(adjacentParts) == 0 || part != adjacentParts[len(adjacentParts)-1] {
					adjacentParts = append(adjacentParts, part)
				}
			}
		}

		if len(adjacentParts) == 2 {
			sumRatios += uint64(adjacentParts[0]) * uint64(adjacentParts[1])
		} else {
			fmt.Printf("gear at [%d][%d] was not added (adjacent parts: %v)\n", gear.row, gear.col, adjacentParts)
		}
	}
	fmt.Println("the sum of all gear ratios is: ", sumRatios)
}
