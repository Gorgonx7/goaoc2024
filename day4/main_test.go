package main

import "testing"

func TestSolveForDirection(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		direction Direction
		char      int
		line      int
		lines     []string
		want      bool
	}{
		{
			name:      "Right",
			target:    "XMAS",
			direction: RIGHT,
			char:      0,
			line:      0,
			lines:     []string{"XMASABC"},
			want:      true,
		},
		{
			name:      "Left Edge",
			target:    "XMAS",
			direction: LEFT,
			char:      0,
			line:      0,
			lines:     []string{"XMASABC"},
			want:      false,
		},
		{
			name:      "LEFT",
			target:    "XMAS",
			direction: LEFT,
			char:      3,
			line:      0,
			lines:     []string{"SAMXABC"},
			want:      true,
		},
		{
			name:      "UP EDGE",
			target:    "XMAS",
			direction: UP,
			char:      3,
			line:      0,
			lines:     []string{"SAMXABC"},
			want:      false,
		},
		{
			name:      "UP",
			target:    "XMAS",
			direction: UP,
			char:      3,
			line:      3,
			lines: []string{
				"AAASAAR",
				"ABCASER",
				"BQEMCKR",
				"SAMXABC",
			},
			want: true,
		},
		{
			name:      "DOWN Edge",
			target:    "XMAS",
			direction: DOWN,
			char:      3,
			line:      3,
			lines: []string{
				"AAASAAR",
				"ABCASER",
				"BQEMCKR",
				"SAMXABC",
			},
			want: false,
		},
		{
			name:      "DOWN",
			target:    "XMAS",
			direction: DOWN,
			char:      3,
			line:      0,
			lines: []string{
				"AAAXAAR",
				"ABCMSER",
				"BQEACKR",
				"SAMSABC",
			},
			want: true,
		},
		{
			name:      "UP_LEFT Edge",
			target:    "XMAS",
			direction: UP_LEFT,
			char:      3,
			line:      0,
			lines: []string{
				"XAAXAAR",
				"AMCMSER",
				"BQAACKR",
				"SAMSSBC",
			},
			want: false,
		},
		{
			name:      "UP_LEFT",
			target:    "XMAS",
			direction: UP_LEFT,
			char:      3,
			line:      3,
			lines: []string{
				"SAAXAAR",
				"AACMSER",
				"BQMACKR",
				"SAMXSBC",
			},
			want: true,
		},
		{
			name:      "UP_RIGHT Edge",
			target:    "XMAS",
			direction: UP_RIGHT,
			char:      3,
			line:      0,
			lines: []string{
				"SAAXAAR",
				"AACMSER",
				"BQMACKR",
				"SAMXSBC",
			},
			want: false,
		},
		{
			name:      "UP_RIGHT",
			target:    "XMAS",
			direction: UP_RIGHT,
			char:      3,
			line:      3,
			lines: []string{
				"SAAXAAS",
				"AACMSAR",
				"BQMAMKR",
				"SAMXSBC",
			},
			want: true,
		},
		{
			name:      "DOWN_RIGHT edge",
			target:    "XMAS",
			direction: DOWN_RIGHT,
			char:      3,
			line:      3,
			lines: []string{
				"SAAXAAS",
				"AACMSAR",
				"BQMAMKR",
				"SAMXSBC",
			},
			want: false,
		},
		{
			name:      "DOWN_RIGHT",
			target:    "XMAS",
			direction: DOWN_RIGHT,
			char:      1,
			line:      0,
			lines: []string{
				"SXAXAAS",
				"AAMMSAR",
				"BQMAMKR",
				"SAMXSBC",
			},
			want: true,
		},
		{
			name:      "DOWN_LEFT edge",
			target:    "XMAS",
			direction: DOWN_LEFT,
			char:      3,
			line:      3,
			lines: []string{
				"SXAXAAS",
				"AAMMSAR",
				"BQMAMKR",
				"SAMXSBC",
			},
			want: false,
		},
		{
			name:      "DOWN_LEFT",
			target:    "XMAS",
			direction: DOWN_LEFT,
			char:      3,
			line:      0,
			lines: []string{
				"SXAXAAS",
				"AAMMSAR",
				"BAMAMKR",
				"SAMXSBC",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveForDirection(tt.target, tt.direction, tt.char, tt.line, tt.lines); got != tt.want {
				t.Errorf("SolveForDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
