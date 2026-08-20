package locales

import (
	"testing"
)

func TestPluralRuleString(t *testing.T) {
	tests := []struct {
		rule     PluralRule
		expected string
	}{
		{PluralRuleUnknown, "Unknown"},
		{PluralRuleZero, "Zero"},
		{PluralRuleOne, "One"},
		{PluralRuleTwo, "Two"},
		{PluralRuleFew, "Few"},
		{PluralRuleMany, "Many"},
		{PluralRuleOther, "Other"},
		{PluralRule(999), "Unknown"},
	}

	for _, tt := range tests {
		if got := tt.rule.String(); got != tt.expected {
			t.Errorf("PluralRule(%d).String() = %q, want %q", tt.rule, got, tt.expected)
		}
	}
}

func TestW(t *testing.T) {
	tests := []struct {
		n        float64
		v        uint64
		expected int64
	}{
		{1.0, 0, 0},
		{1.0, 2, 0},
		{1.2, 1, 1},
		{1.20, 2, 1},
		{1.23, 2, 2},
		{1.2300, 4, 2},
		{1.2345, 4, 4},
		{0.0, 0, 0},
	}

	for _, tt := range tests {
		if got := W(tt.n, tt.v); got != tt.expected {
			t.Errorf("W(%v, %d) = %d, want %d", tt.n, tt.v, got, tt.expected)
		}
	}
}

func TestF(t *testing.T) {
	tests := []struct {
		n        float64
		v        uint64
		expected int64
	}{
		{1.0, 0, 0},
		{1.0, 2, 0},
		{1.2, 1, 2},
		{1.20, 2, 20},
		{1.23, 2, 23},
		{1.2300, 4, 2300},
	}

	for _, tt := range tests {
		if got := F(tt.n, tt.v); got != tt.expected {
			t.Errorf("F(%v, %d) = %d, want %d", tt.n, tt.v, got, tt.expected)
		}
	}
}

func TestT(t *testing.T) {
	tests := []struct {
		n        float64
		v        uint64
		expected int64
	}{
		{1.0, 0, 0},
		{1.0, 2, 0},
		{1.2, 1, 2},
		{1.20, 2, 2},
		{1.23, 2, 23},
		{1.2300, 4, 23},
		{1.0450, 4, 45},
	}

	for _, tt := range tests {
		if got := T(tt.n, tt.v); got != tt.expected {
			t.Errorf("T(%v, %d) = %d, want %d", tt.n, tt.v, got, tt.expected)
		}
	}
}
