package main

import (
	"reflect"
	"testing"
)

func Test_buildRules(t *testing.T) {
	tests := []struct {
		name       string
		rulesBlock string
		want       map[int][]int
	}{
		{
			name:       "test parse",
			rulesBlock: "55|42\n55|81\n51|11\n51|22",
			want: map[int][]int{
				55: {42, 81},
				51: {11, 22},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildRules(tt.rulesBlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMiddleNumber(t *testing.T) {
	tests := []struct {
		name   string
		record string
		want   string
	}{
		{
			name:   "one",
			record: "75,47,61,53,29",
			want:   "61",
		},
		{
			name:   "two",
			record: "97,61,53,29,13",
			want:   "53",
		},
		{
			name:   "three",
			record: "75,29,13",
			want:   "29",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMiddleNumber(tt.record); got != tt.want {
				t.Errorf("getMiddleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recordIsGood(t *testing.T) {
	globalRules := map[string]bool{
		"47|53": true,
		"97|13": true,
		"97|61": true,
		"97|47": true,
		"75|29": true,
		"61|13": true,
		"75|53": true,
		"29|13": true,
		"97|29": true,
		"53|29": true,
		"61|53": true,
		"97|53": true,
		"61|29": true,
		"47|13": true,
		"75|47": true,
		"97|75": true,
		"47|61": true,
		"75|61": true,
		"47|29": true,
		"75|13": true,
		"53|13": true,
	}
	tests := []struct {
		name   string
		rules  map[string]bool
		record string
		want   bool
	}{
		{
			rules:  globalRules,
			record: "75,47,61,53,29",
			want:   true,
		},
		{
			rules:  globalRules,
			record: "97,61,53,29,13",
			want:   true,
		},
		{
			rules:  globalRules,
			record: "75,29,13",
			want:   true,
		},
		{
			rules:  globalRules,
			record: "75,97,47,61,53",
			want:   false,
		},
		{
			rules:  globalRules,
			record: "61,13,29",
			want:   false,
		},
		{
			rules:  globalRules,
			record: "97,13,75,29,47",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recordIsGood(tt.rules, tt.record); got != tt.want {
				t.Errorf("recordIsGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
