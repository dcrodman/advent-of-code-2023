package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var instructionRegex = regexp.MustCompile(`(?P<node>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)`)

type mapNode struct {
	left  string
	right string
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	instructions, network := parseInput(input)
	part1(instructions, network)
	part2(instructions, network)
}

func parseInput(input []byte) ([]rune, map[string]mapNode) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Scan()

	instructions := []rune(scanner.Text())
	scanner.Scan()

	network := make(map[string]mapNode)
	for scanner.Scan() {
		matches := instructionRegex.FindStringSubmatch(scanner.Text())
		network[matches[1]] = mapNode{
			left:  matches[2],
			right: matches[3],
		}
	}

	return instructions, network
}

func part1(instructions []rune, network map[string]mapNode) {
	steps := 0
	instructionCtr := 0
	nextNode := "AAA"

	for nextNode != "ZZZ" {
		if instructions[instructionCtr] == 'L' {
			nextNode = network[nextNode].left
		} else {
			nextNode = network[nextNode].right
		}
		steps++
		instructionCtr++
		if instructionCtr == len(instructions) {
			instructionCtr = 0
		}
	}

	fmt.Println("part1 steps required:", steps)
}

func part2(instructions []rune, network map[string]mapNode) {
	nodes := make([]string, 0)
	for k := range network {
		if strings.HasSuffix(k, "A") {
			nodes = append(nodes, k)
		}
	}

	nodeSteps := make([]int, len(nodes))
	for i, n := range nodes {
		steps := 0
		instructionCtr := 0
		nextNode := n

		for !strings.HasSuffix(nextNode, "Z") {
			if instructions[instructionCtr] == 'L' {
				nextNode = network[nextNode].left
			} else {
				nextNode = network[nextNode].right
			}

			steps++
			instructionCtr++
			if instructionCtr == len(instructions) {
				instructionCtr = 0
			}
		}
		nodeSteps[i] = steps
	}

	// This feels stupid because this solution assumes that these cycles all
	// overlap at the same point but also the challenge seems designed for that?
	fmt.Println("part2 steps required:", lcm(nodeSteps[0], nodeSteps[1], nodeSteps[2:]...))
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
