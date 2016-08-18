package en

import (
	"testing"
	"time"

	"github.com/go-playground/locales/currency"
)

func TestFullTime(t *testing.T) {

	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		t.Errorf("Expected '<nil>' Got '%s'", err)
	}

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, loc),
			expected: "9:05:01 am Eastern Standard Time",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtTimeFull(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestLongTime(t *testing.T) {

	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		t.Errorf("Expected '<nil>' Got '%s'", err)
	}

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, loc),
			expected: "9:05:01 am EST",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtTimeLong(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMediumTime(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, time.UTC),
			expected: "9:05:01 am",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtTimeMedium(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestShortTime(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, time.UTC),
			expected: "9:05 am",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtTimeShort(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFullDate(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "Wednesday, February 3, 2016",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtDateFull(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestLongDate(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "February 3, 2016",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtDateLong(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMediumDate(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "Feb 3, 2016",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtDateMedium(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestShortDate(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "2/3/16",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtDateShort(tt.t))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

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
		{
			num:      0,
			v:        2,
			expected: "$0.00",
		},
		{
			num:      -0,
			v:        2,
			expected: "$0.00",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtCurrency(tt.num, tt.v, currency.USD))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestAccounting(t *testing.T) {

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
			expected: "($221,123,456.564)",
		},
		{
			num:      -0,
			v:        2,
			expected: "$0.00",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := string(trans.FmtAccounting(tt.num, tt.v, currency.USD))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}
