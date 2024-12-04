package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := readData("data.txt")
	data = fmt.Sprintf("do()%sdon't()", data)
	dosDonts := regexp.MustCompile(`do\(\)(.*?)don't\(\)|do\(\)(.*?)$|^(.*?)do\(\)`)
	doDontMatches := dosDonts.FindAllString(data, -1)
	total := 0
	for _, doMatch := range doDontMatches {
		rx := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		matches := rx.FindAllString(doMatch, -1)

		for _, match := range matches {
			parts := strings.Split(match, ",")
			part1, err := strconv.Atoi(strings.TrimPrefix(parts[0], "mul("))
			if err != nil {
				panic(err)
			}
			part2, err := strconv.Atoi(strings.TrimSuffix(parts[1], ")"))
			if err != nil {
				panic(err)
			}
			total += part1 * part2
		}
	}

	fmt.Printf("part 2: %v\n", total)

}

func readData(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
