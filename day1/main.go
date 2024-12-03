package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := readData("data.txt")
	lines := strings.Split(data, "\n")
	leftSide := []int{}
	rightSide := []int{}
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
		rightSide = append(rightSide, right)
	}
	slices.Sort(leftSide)
	slices.Sort(rightSide)
	totalDiff := 0
	for i := 0; i < len(leftSide); i++ {
		totalDiff += int(math.Abs(float64(leftSide[i]) - float64(rightSide[i])))

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
