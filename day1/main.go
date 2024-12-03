package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := readData("data.txt")
	lines := strings.Split(data, "\n")
	leftSide := []int{}
	rightSide := map[int]int{}
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		leftSide = append(leftSide, left)
		if _, ok := rightSide[right]; !ok {
			rightSide[right] = 1
			continue
		}
		rightSide[right] = rightSide[right] + 1
	}
	slices.Sort(leftSide)
	totalDiff := 0
	for i := 0; i < len(leftSide); i++ {
		totalDiff += leftSide[i] * rightSide[leftSide[i]]
	}
	fmt.Println("Result: ", totalDiff)
}

func readData(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
