package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      bool
	}{
		{"1 + 2", 3, false},
		{"2 - 3", -1, false},
		{"4 * 5", 20, false},
		{"10 / 2", 5, false},
		{"(3 + 5) * 2", 16, false},
		{"---1", -1, false},
		{"5 + (-3 * 2)", -1, false},
		{"3 / (1 + 2)", 1, false},
		{"1 / 0", 0, true},   // Expected to fail
		{"2 ** 3", 0, true},  // Unsupported operator
		{"3 +", 0, true},     // Incomplete expression
		{"abc + 1", 0, true}, // Invalid characters
	}

	for _, test := range tests {
		defer func() {
			if r := recover(); r != nil {
				if !test.err {
					t.Errorf("Unexpected error for %q: %v", test.input, r)
				}
			}
		}()

		postfix := InfixToPostfix(test.input)
		tree := BuildExpressionTree(postfix)
		result := tree.Evaluate()

		if !test.err && result != test.expected {
			t.Errorf("For input %q, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
