package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var startX int
	var startY int

	tiles := make([][]string, 0)
	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	row := 0
	for scanner.Scan() {
		tileRow := make([]string, 0)
		for i, r := range scanner.Text() {
			if r == 'S' {
				startX = row
				startY = i
			}
			tileRow = append(tileRow, string(r))
		}
		tiles = append(tiles, tileRow)
		row++
	}

	part1(tiles, startX, startY)
}

func part1(tiles [][]string, startX, startY int) {
	paths := make([]int, 0)
	starting := location{row: startX, col: startY}

	northTile := tiles[startX-1][startY]
	fmt.Println("north")
	if northTile == "|" || northTile == "7" || northTile == "F" {
		next := location{row: startX - 1, col: startY}
		paths = append(paths, walk(tiles, starting, next))
	}

	fmt.Println("south")
	southTile := tiles[startX+1][startY]
	if southTile == "|" || southTile == "L" || southTile == "J" {
		next := location{row: startX + 1, col: startY}
		paths = append(paths, walk(tiles, starting, next))
	}

	fmt.Println("east")
	eastTile := tiles[startX][startY+1]
	if eastTile == "-" || eastTile == "7" || eastTile == "J" {
		next := location{row: startX, col: startY + 1}
		paths = append(paths, walk(tiles, starting, next))
	}

	fmt.Println("west")
	if startY-1 >= 0 {
		westTile := tiles[startX][startY-1]
		if westTile == "-" || westTile == "L" || westTile == "F" {
			next := location{row: startX, col: startY - 1}
			paths = append(paths, walk(tiles, starting, next))
		}
	}

	max := slices.Max(paths)
	fmt.Println(max / 2)
}

type location struct {
	row int
	col int
}

func walk(tiles [][]string, last, cur location) int {
	if tiles[cur.row][cur.col] == "S" {
		return 1
	}

	next := location{row: cur.row, col: cur.col}
	switch tiles[cur.row][cur.col] {
	//| is a vertical pipe connecting north and south.
	case "|":
		if cur.row < last.row {
			// Go south.
			next.row -= 1
		} else {
			// Go north.
			next.row += 1
		}
	// - is a horizontal pipe connecting east and west.
	case "-":
		if cur.col > last.col {
			next.col += 1
		} else {
			next.col -= 1
		}
	// L is a 90-degree bend connecting north and east.
	case "L":
		if cur.row-1 == last.row {
			// Go east.
			next.col += 1
		} else {
			// Go north.
			next.row -= 1
		}
	// J is a 90-degree bend connecting north and west.
	case "J":
		if cur.row-1 == last.row {
			// Go west.
			next.col -= 1
		} else {
			// Go north.
			next.row -= 1
		}
	// 7 is a 90-degree bend connecting south and west.
	case "7":
		if cur.row+1 == last.row {
			// Go west.
			next.col -= 1
		} else {
			// Go south.
			next.row += 1
		}
	// F is a 90-degree bend connecting south and east.
	case "F":
		if cur.row+1 == last.row {
			// Go east.
			next.col += 1
		} else {
			// Go south
			next.row += 1
		}
	// . is ground; there is no pipe in this tile.
	case ".":
		return 0
	}
	// fmt.Printf("%s (%v, %v) traveling to %s (%v, %v)\n", tiles[cur.row][cur.col], cur.row, cur.col, tiles[next.row][next.col], next.row, next.col)
	return 1 + walk(tiles, cur, next)
}
