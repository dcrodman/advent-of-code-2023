// I sincerely hope nobody looks at this but if you do then let me take a second
// to just apologize because this is a last ditch I-need-to-catch-up effort.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type step struct {
	row       int
	col       int
	direction Direction
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1(input)
	part2(input)
}

func parseInput(input []byte) [][]rune {
	contraption := make([][]rune, 0)

	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	for scanner.Scan() {
		row := make([]rune, 0)
		for _, r := range scanner.Text() {
			row = append(row, r)
		}
		contraption = append(contraption, row)
	}

	return contraption
}

func part1(input []byte) {
	contraption := parseInput(input)
	solution := make([][]rune, len(contraption))
	for i := range contraption {
		solution[i] = make([]rune, len(contraption[i]))
		for j := range solution[i] {
			solution[i][j] = '.'
		}
	}
	visited := make(map[step]struct{})

	beam(contraption, solution, visited, step{0, 0, Right})
	fmt.Println("part1 total energized:", sumEnergized(solution))
}

func part2(input []byte) {
	contraption := parseInput(input)

	maxNumEnergized := 0
	for row := range contraption {
		for col := range contraption[row] {
			// Create a new solution space to work in.
			solution := make([][]rune, len(contraption))
			for i := range contraption {
				solution[i] = make([]rune, len(contraption[i]))
				for j := range solution[i] {
					solution[i][j] = '.'
				}
			}

			visited := make(map[step]struct{})
			for _, startDirection := range determineStartDirection(contraption, row, col) {
				start := step{row: row, col: col, direction: startDirection}
				beam(contraption, solution, visited, start)

				numEnergized := sumEnergized(solution)
				if numEnergized > maxNumEnergized {
					maxNumEnergized = numEnergized
				}
			}
		}
	}

	fmt.Println("part2 highest number of energized:", maxNumEnergized)
}

func determineStartDirection(contraption [][]rune, row, col int) []Direction {
	maxRow := len(contraption) - 1
	maxCol := len(contraption[0]) - 1
	directions := make([]Direction, 0)

	switch {
	case row == 0 && col == 0:
		// Top left corner.
		directions = append(directions, Right)
		directions = append(directions, Down)
	case col == 0:
		// Left column.
		directions = append(directions, Right)
	case col == 0 && row == maxRow:
		// Bottom left corner.
		directions = append(directions, Right)
		directions = append(directions, Up)
	case col == maxCol:
		// Bottom row.
		directions = append(directions, Up)
	case row == maxRow && col == maxCol:
		// Bottom right corner.
		directions = append(directions, Left)
		directions = append(directions, Up)
	case col == maxCol:
		// Right column.
		directions = append(directions, Left)
	case row == 0 && col == maxCol:
		// Top right corner.
		directions = append(directions, Left)
		directions = append(directions, Down)
	case row == 0:
		// Top row.
		directions = append(directions, Down)
	}
	return directions
}

func beam(contraption, solution [][]rune, visited map[step]struct{}, s step) {
	if s.row < 0 || s.row > len(contraption)-1 || s.col < 0 || s.col > len(contraption[0])-1 {
		return
	}

	if _, ok := visited[s]; ok {
		return
	}
	visited[s] = struct{}{}

	solution[s.row][s.col] = '#'

	nextRow, nextCol := s.row, s.col
	nextDirection := s.direction
	switch contraption[s.row][s.col] {
	case '.':
		switch s.direction {
		case Up:
			nextRow -= 1
		case Down:
			nextRow += 1
		case Left:
			nextCol -= 1
		case Right:
			nextCol += 1
		}
	case '|':
		switch s.direction {
		case Up:
			nextRow -= 1
		case Down:
			nextRow += 1
		case Left, Right:
			beam(contraption, solution, visited, step{row: s.row + 1, col: s.col, direction: Down})
			beam(contraption, solution, visited, step{row: s.row - 1, col: s.col, direction: Up})
			return
		}
	case '-':
		switch s.direction {
		case Up, Down:
			beam(contraption, solution, visited, step{row: s.row, col: s.col + 1, direction: Right})
			beam(contraption, solution, visited, step{row: s.row, col: s.col - 1, direction: Left})
			return
		case Left:
			nextCol -= 1
		case Right:
			nextCol += 1
		}
	case '\\':
		switch s.direction {
		case Up:
			beam(contraption, solution, visited, step{row: s.row, col: s.col - 1, direction: Left})
		case Down:
			beam(contraption, solution, visited, step{row: s.row, col: s.col + 1, direction: Right})
		case Left:
			beam(contraption, solution, visited, step{row: s.row - 1, col: s.col, direction: Up})
		case Right:
			beam(contraption, solution, visited, step{row: s.row + 1, col: s.col, direction: Down})
		}
		return
	case '/':
		switch s.direction {
		case Up:
			beam(contraption, solution, visited, step{row: s.row, col: s.col + 1, direction: Right})
		case Down:
			beam(contraption, solution, visited, step{row: s.row, col: s.col - 1, direction: Left})
		case Left:
			beam(contraption, solution, visited, step{row: s.row + 1, col: s.col, direction: Down})
		case Right:
			beam(contraption, solution, visited, step{row: s.row - 1, col: s.col, direction: Up})
		}
		return
	}

	beam(contraption, solution, visited, step{row: nextRow, col: nextCol, direction: nextDirection})
}

func sumEnergized(solution [][]rune) int {
	totalEnergized := 0
	for row := range solution {
		for col := range solution[row] {
			if solution[row][col] == '#' {
				totalEnergized += 1
			}
		}
	}
	return totalEnergized
}
