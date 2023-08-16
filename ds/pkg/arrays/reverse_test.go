package arrays

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"Hello", "olleH"},
		{"", ""},
		{"Hello World", "dlroW olleH"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := reverse(tt.input); got != tt.output {
				t.Errorf("For input %s, want %s, got %s", tt.input, tt.output, got)
			}
		})
	}
}
