package mr_IN

import "testing"

func TestGrouping(t *testing.T) {

	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      1123456.5643,
			v:        2,
			expected: "11,23,456.56",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtNumber(tt.num, tt.v))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s", tt.expected, s)
		}
	}
}
