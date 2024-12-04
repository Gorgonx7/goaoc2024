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
			if lines[line][char] == "A"[0] {
				rightDiagonalOK := false
				if SolveForDirection("M", UP_LEFT, char, line, lines) {
					if SolveForDirection("S", DOWN_RIGHT, char, line, lines) {
						rightDiagonalOK = true
					}
				} else if SolveForDirection("S", UP_LEFT, char, line, lines) {
					if SolveForDirection("M", DOWN_RIGHT, char, line, lines) {
						rightDiagonalOK = true
					}
				}
				leftDiagonalOK := false
				if SolveForDirection("M", UP_RIGHT, char, line, lines) {
					if SolveForDirection("S", DOWN_LEFT, char, line, lines) {
						leftDiagonalOK = true
					}
				} else if SolveForDirection("S", UP_RIGHT, char, line, lines) {
					if SolveForDirection("M", DOWN_LEFT, char, line, lines) {
						leftDiagonalOK = true
					}
				}
				if rightDiagonalOK && leftDiagonalOK {
					total += 1
				}
			}
			// if SolveForDirection(target, UP, char, line, lines) {
			// 	total += 1
			// }
			// if SolveForDirection(target, DOWN, char, line, lines) {
			// 	total += 1
			// }
			// if SolveForDirection(target, LEFT, char, line, lines) {
			// 	total += 1
			// }
			// if SolveForDirection(target, RIGHT, char, line, lines) {
			// 	total += 1
			// }
			// if SolveForDirection(target, UP_LEFT, char, line, lines) {
			// 	total += 1
			// }
			// if SolveForDirection(target, UP_RIGHT, char, line, lines) {
			// 	total += 1
			// }
			// if SolveForDirection(target, DOWN_LEFT, char, line, lines) {
			// 	total += 1
			// }
			// if SolveForDirection(target, DOWN_RIGHT, char, line, lines) {
			// 	total += 1
			// }
		}
	}
	fmt.Printf("total: %v\n", total)
}

func SolveForDirection(target string, direction Direction, char, line int, lines []string) bool {
	switch direction {
	case RIGHT:
		if char+1 > len(lines[line])-1 {
			return false
		}
		if lines[line][char+1] == target[0] {
			return true
		}
		return false
	case LEFT:
		if char-1 < 0 {
			return false // Edge handling
		}
		if lines[line][char-1] == target[0] {
			return true
		}
		return false
	case UP:
		if line-1 < 0 {
			return false // Edge handling
		}
		if lines[line-1][char] == target[0] {

			return true
		}
		return false
	case DOWN:
		if line+1 > len(lines)-1 {
			return false // Edge handling
		}
		if lines[line+1][char] == target[0] {
			return true
		}
		return false
	case UP_LEFT:
		if line-1 < 0 || char-1 < 0 {
			return false // Edge handling
		}
		if lines[line-1][char-1] == target[0] {
			return true
		}
		return false
	case UP_RIGHT:
		if line-1 < 0 || char+1 > len(lines[line])-1 {
			return false // Edge handling
		}
		if lines[line-1][char+1] == target[0] {
			return true
		}
		return false
	case DOWN_RIGHT:
		if line+1 > len(lines)-1 || char+1 > len(lines[line])-1 {
			return false // Edge handling
		}
		if lines[line+1][char+1] == target[0] {
			return true
		}
		return false
	case DOWN_LEFT:
		if line+1 > len(lines)-1 || char-1 < 0 {
			return false
		}
		if lines[line+1][char-1] == target[0] {
			return true
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
