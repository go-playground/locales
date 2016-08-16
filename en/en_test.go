package en

import (
	"testing"

	"github.com/go-playground/locales/currency"
)

func TestCurrency(t *testing.T) {

	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      1123456.5643,
			v:        2,
			expected: "$1,123,456.56",
		},
		{
			num:      1123456.5643,
			v:        1,
			expected: "$1,123,456.60",
		},
		{
			num:      221123456.5643,
			v:        3,
			expected: "$221,123,456.564",
		},
		{
			num:      -221123456.5643,
			v:        3,
			expected: "-$221,123,456.564",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtCurrency(tt.num, tt.v, currency.USD))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s", tt.expected, s)
		}
	}
}
