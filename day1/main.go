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

// The newly-improved calibration document consists of lines of text; each line originally
// contained a specific calibration value that the Elves now need to recover. On each line,
// the calibration value can be found by combining the first digit and the last digit (in
// that order) to form a single two-digit number.
//
// For example:
// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
//
// In this example, the calibration values of these four lines are 12, 38, 15, and
// 77. Adding these together produces 142.
//
// Consider your entire calibration document. What is the sum of all of the calibration values?
func part1(input []byte) {
	var total int64 = 0
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		digits := make([]rune, 0)
		for _, r := range line {
			if unicode.IsDigit(r) {
				digits = append(digits, r)
			}
		}

		calibration := fmt.Sprintf("%s%s", string(digits[0]), string(digits[len(digits)-1]))
		value, _ := strconv.ParseInt(calibration, 10, 8)
		total += value
	}

	fmt.Printf("part1 total calibration value: %d\n", total)
}

// Your calculation isn't quite right. It looks like some of the digits are actually
// spelled out with letters: one, two, three, four, five, six, seven, eight, and nine
// also count as valid "digits".

// Equipped with this new information, you now need to find the real first and last
// digit on each line. For example:
// two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen
//
// In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.
//
// What is the sum of all of the calibration values?

var spellings = map[string]int64{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func part2(input []byte) {
	var total int64 = 0
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		digits := make([]int64, 0)
		lineRunes := []rune(line)

		// I'm sure there's a way to do this O(n) time but I'm also a little hungover
		// from my birthday and going to settle for O(mn).
		for i := 0; i < len(lineRunes); i++ {
			if unicode.IsDigit(lineRunes[i]) {
				value, _ := strconv.ParseInt(string(lineRunes[i]), 10, 8)
				digits = append(digits, value)
			} else {
				for k, v := range spellings {
					if strings.HasPrefix(string(lineRunes[i:]), k) {
						digits = append(digits, v)
					}
				}
			}
		}

		calibration := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
		value, _ := strconv.ParseInt(calibration, 10, 8)
		total += value
	}

	fmt.Printf("part2 total calibration value: %d\n", total)
}
