package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readData("data.txt")
	split := strings.Split(data, "\n\n")
	rulesBlock := split[0]
	rules := buildRules(rulesBlock)
	instructions := split[1]
	total := 0
	for _, line := range strings.Split(instructions, "\n") {
		if recordIsGood(rules, line) {
			continue
		}
		line = fixRecord(rules, line)
		num, err := strconv.Atoi(getMiddleNumber(line))
		if err != nil {
			panic(err)
		}
		total += num
	}
	fmt.Println(total)
}
func getMiddleNumber(record string) string {
	nums := strings.Split(record, ",")
	if len(nums)%2 == 0 {
		panic(fmt.Errorf("even numbered list, what the *$|^?"))
	}
	return nums[len(nums)/2.0]
}
func fixRecord(rules map[string]bool, record string) string {
	splitRecord := strings.Split(record, ",")

RESTART:
	record = strings.Join(splitRecord, ",")
	if recordIsGood(rules, record) {
		return record
	}
	for i := range splitRecord {
		if i == 0 {
			continue
		}
		for j := range make([]int, i) {
			if ruleExists(rules, string(splitRecord[i]), string(splitRecord[j])) {
				store := make([]string, len(splitRecord))
				copy(store, splitRecord)
				result := make([]string, 0)
				result = append(result, store[:j]...)
				result = append(result, store[i])
				result = append(result, store[j])
				if i != len(store)-1 {
					store = append(store[:i], store[i+1:]...)
				} else {
					store = store[:i]
				}
				result = append(result, store[j+1:]...)

				copy(splitRecord, result)
				goto RESTART
			}
		}
	}
	panic(fmt.Errorf("reached end and shouldn't have"))
}
func recordIsGood(rules map[string]bool, record string) bool {
	splitRecord := strings.Split(record, ",")
	for i := range splitRecord {
		if i == 0 {
			continue
		}
		for j := range make([]int, i) {
			if ruleExists(rules, string(splitRecord[i]), string(splitRecord[j])) {
				return false
			}
		}
	}
	return true
}
func ruleExists(rules map[string]bool, numberUn, numberTu string) bool {
	_, ok := rules[fmt.Sprintf("%s|%s", numberUn, numberTu)]
	return ok
}

func buildRules(rulesBlock string) map[string]bool {
	rules := make(map[string]bool)
	for _, line := range strings.Split(rulesBlock, "\n") {
		rules[line] = true
	}
	return rules
}

func readData(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
