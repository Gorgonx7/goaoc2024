package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readData("data.txt")
	reports := strings.Split(data, "\n")
	totalSafe := 0
	for _, levelLine := range reports {
		levels := strings.Split(levelLine, " ")
		ilevels := convertStringSliceToint(levels)
		isSafe := isReportSafe(ilevels)
		if isSafe {
			totalSafe += 1
			continue
		}
		for i := range ilevels {
			toTest := make([]int, len(ilevels))
			copy(toTest, ilevels)
			toTest = append(toTest[:i], toTest[i+1:]...)
			isSafe := isReportSafe(toTest)
			if isSafe {
				totalSafe += 1
				break
			}
		}

	}
	fmt.Println(totalSafe)
}
func isDecreasing(levels []int) bool {
	return levels[0] > levels[1]

}
func convertStringSliceToint(levels []string) []int {
	var toReturn []int
	for _, level := range levels {
		first, err := strconv.Atoi(level)
		if err != nil {
			panic(err)
		}
		toReturn = append(toReturn, first)
	}
	return toReturn
}
func readData(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func isReportSafe(levels []int) bool {
	isSafe := true
	isDecreasing := isDecreasing(levels)
	for i, level := range levels {
		if i == 0 {
			continue
		}
		lastLevel := levels[i-1]
		if level == lastLevel {
			isSafe = false
			break
		}
		if isDecreasing && level > lastLevel {
			isSafe = false
			break
		} else if !isDecreasing && level < lastLevel {
			isSafe = false
			break
		}
		diff := math.Abs(float64(level) - float64(lastLevel))
		if diff > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}
