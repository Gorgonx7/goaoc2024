package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	UP_RIGHT
	UP_LEFT
	DOWN_RIGHT
	DOWN_LEFT
)
const (
	target = "XMAS"
)

func main() {
	data := readData("data.txt")
	lines := strings.Split(data, "\n")
	total := 0
	for line := range lines {
		for char := range len(lines[line]) {
			if SolveForDirection(target, UP, char, line, lines) {
				total += 1
			}
			if SolveForDirection(target, DOWN, char, line, lines) {
				total += 1
			}
			if SolveForDirection(target, LEFT, char, line, lines) {
				total += 1
			}
			if SolveForDirection(target, RIGHT, char, line, lines) {
				total += 1
			}
			if SolveForDirection(target, UP_LEFT, char, line, lines) {
				total += 1
			}
			if SolveForDirection(target, UP_RIGHT, char, line, lines) {
				total += 1
			}
			if SolveForDirection(target, DOWN_LEFT, char, line, lines) {
				total += 1
			}
			if SolveForDirection(target, DOWN_RIGHT, char, line, lines) {
				total += 1
			}
		}
	}
	fmt.Printf("total: %v\n", total)
}

func SolveForDirection(target string, direction Direction, char, line int, lines []string) bool {
	if len(target) == 1 {
		return lines[line][char] == target[0]
	}
	switch direction {
	case RIGHT:
		if char+1 > len(lines[line])-1 {
			return false
		}
		if lines[line][char] == target[0] {
			return SolveForDirection(target[1:], RIGHT, char+1, line, lines)
		}
		return false
	case LEFT:
		if lines[line][char] == target[0] {
			if char-1 < 0 {
				return false // Edge handling
			}
			return SolveForDirection(target[1:], LEFT, char-1, line, lines)
		}
		return false
	case UP:
		if lines[line][char] == target[0] {
			if line-1 < 0 {
				return false // Edge handling
			}
			return SolveForDirection(target[1:], UP, char, line-1, lines)
		}
		return false
	case DOWN:
		if lines[line][char] == target[0] {
			if line+1 > len(lines)-1 {
				return false // Edge handling
			}
			return SolveForDirection(target[1:], DOWN, char, line+1, lines)
		}
		return false
	case UP_LEFT:
		if lines[line][char] == target[0] {
			if line-1 < 0 || char-1 < 0 {
				return false // Edge handling
			}
			return SolveForDirection(target[1:], UP_LEFT, char-1, line-1, lines)
		}
		return false
	case UP_RIGHT:
		if lines[line][char] == target[0] {
			if line-1 < 0 || char+1 > len(lines[line])-1 {
				return false // Edge handling
			}
			return SolveForDirection(target[1:], UP_RIGHT, char+1, line-1, lines)
		}
		return false
	case DOWN_RIGHT:
		if lines[line][char] == target[0] {
			if line+1 > len(lines)-1 || char+1 > len(lines[line])-1 {
				return false // Edge handling
			}
			return SolveForDirection(target[1:], DOWN_RIGHT, char+1, line+1, lines)
		}
		return false
	case DOWN_LEFT:
		if lines[line][char] == target[0] {
			if line+1 > len(lines)-1 || char-1 < 0 {
				return false
			}
			return SolveForDirection(target[1:], DOWN_LEFT, char-1, line+1, lines)
		}
		return false
	default:
		panic(fmt.Errorf("invalid case"))
	}
}

func readData(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
