package main

import (
	"testing"
)

func Test_isReportSafe(t *testing.T) {
	tests := []struct {
		name         string
		levels       []int
		isDecreasing bool
		want         bool
	}{

		{
			name:   "example 1",
			levels: []int{7, 6, 4, 2, 1},
			want:   true,
		},
		{
			name:   "example 2",
			levels: []int{1, 2, 7, 8, 9},
			want:   false,
		},
		{
			name:   "example 3",
			levels: []int{9, 7, 6, 2, 1},
			want:   false,
		},
		{
			name:   "example 4",
			levels: []int{1, 3, 2, 4, 5},
			want:   false,
		},
		{
			name:   "example 5",
			levels: []int{8, 6, 4, 4, 1},
			want:   false,
		},
		{
			name:   "example 6",
			levels: []int{1, 3, 6, 7, 9},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReportSafe(tt.levels); got != tt.want {
				t.Errorf("isReportSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
