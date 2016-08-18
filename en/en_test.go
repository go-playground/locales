package en

import (
	"testing"
	"time"

	"github.com/go-playground/locales/currency"
)

func TestDaysAbbreviated(t *testing.T) {

	trans := New()
	days := trans.WeekdaysAbbreviated()

	for i, day := range days {
		s := string(trans.WeekdayAbbreviated(time.Weekday(i)))
		if s != string(day) {
			t.Errorf("Expected '%s' Got '%s'", string(day), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "Sun",
		},
		{
			idx:      1,
			expected: "Mon",
		},
		{
			idx:      2,
			expected: "Tue",
		},
		{
			idx:      3,
			expected: "Wed",
		},
		{
			idx:      4,
			expected: "Thu",
		},
		{
			idx:      5,
			expected: "Fri",
		},
		{
			idx:      6,
			expected: "Sat",
		},
	}

	for _, tt := range tests {
		s := string(trans.WeekdayAbbreviated(time.Weekday(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysNarrow(t *testing.T) {

	trans := New()
	days := trans.WeekdaysNarrow()

	for i, day := range days {
		s := string(trans.WeekdayNarrow(time.Weekday(i)))
		if s != string(day) {
			t.Errorf("Expected '%s' Got '%s'", string(day), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "S",
		},
		{
			idx:      1,
			expected: "M",
		},
		{
			idx:      2,
			expected: "T",
		},
		{
			idx:      3,
			expected: "W",
		},
		{
			idx:      4,
			expected: "T",
		},
		{
			idx:      5,
			expected: "F",
		},
		{
			idx:      6,
			expected: "S",
		},
	}

	for _, tt := range tests {
		s := string(trans.WeekdayNarrow(time.Weekday(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysShort(t *testing.T) {

	trans := New()
	days := trans.WeekdaysShort()

	for i, day := range days {
		s := string(trans.WeekdayShort(time.Weekday(i)))
		if s != string(day) {
			t.Errorf("Expected '%s' Got '%s'", string(day), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "Su",
		},
		{
			idx:      1,
			expected: "Mo",
		},
		{
			idx:      2,
			expected: "Tu",
		},
		{
			idx:      3,
			expected: "We",
		},
		{
			idx:      4,
			expected: "Th",
		},
		{
			idx:      5,
			expected: "Fr",
		},
		{
			idx:      6,
			expected: "Sa",
		},
	}

	for _, tt := range tests {
		s := string(trans.WeekdayShort(time.Weekday(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysWide(t *testing.T) {

	trans := New()
	days := trans.WeekdaysWide()

	for i, day := range days {
		s := string(trans.WeekdayWide(time.Weekday(i)))
		if s != string(day) {
			t.Errorf("Expected '%s' Got '%s'", string(day), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "Sunday",
		},
		{
			idx:      1,
			expected: "Monday",
		},
		{
			idx:      2,
			expected: "Tuesday",
		},
		{
			idx:      3,
			expected: "Wednesday",
		},
		{
			idx:      4,
			expected: "Thursday",
		},
		{
			idx:      5,
			expected: "Friday",
		},
		{
			idx:      6,
			expected: "Saturday",
		},
	}

	for _, tt := range tests {
		s := string(trans.WeekdayWide(time.Weekday(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsAbbreviated(t *testing.T) {

	trans := New()
	months := trans.MonthsAbbreviated()

	for i, month := range months {
		s := string(trans.MonthAbbreviated(time.Month(i + 1)))
		if s != string(month) {
			t.Errorf("Expected '%s' Got '%s'", string(month), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "Jan",
		},
		{
			idx:      2,
			expected: "Feb",
		},
		{
			idx:      3,
			expected: "Mar",
		},
		{
			idx:      4,
			expected: "Apr",
		},
		{
			idx:      5,
			expected: "May",
		},
		{
			idx:      6,
			expected: "Jun",
		},
		{
			idx:      7,
			expected: "Jul",
		},
		{
			idx:      8,
			expected: "Aug",
		},
		{
			idx:      9,
			expected: "Sep",
		},
		{
			idx:      10,
			expected: "Oct",
		},
		{
			idx:      11,
			expected: "Nov",
		},
		{
			idx:      12,
			expected: "Dec",
		},
	}

	for _, tt := range tests {
		s := string(trans.MonthAbbreviated(time.Month(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsNarrow(t *testing.T) {

	trans := New()
	months := trans.MonthsNarrow()

	for i, month := range months {
		s := string(trans.MonthNarrow(time.Month(i + 1)))
		if s != string(month) {
			t.Errorf("Expected '%s' Got '%s'", string(month), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "J",
		},
		{
			idx:      2,
			expected: "F",
		},
		{
			idx:      3,
			expected: "M",
		},
		{
			idx:      4,
			expected: "A",
		},
		{
			idx:      5,
			expected: "M",
		},
		{
			idx:      6,
			expected: "J",
		},
		{
			idx:      7,
			expected: "J",
		},
		{
			idx:      8,
			expected: "A",
		},
		{
			idx:      9,
			expected: "S",
		},
		{
			idx:      10,
			expected: "O",
		},
		{
			idx:      11,
			expected: "N",
		},
		{
			idx:      12,
			expected: "D",
		},
	}

	for _, tt := range tests {
		s := string(trans.MonthNarrow(time.Month(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsWide(t *testing.T) {

	trans := New()
	months := trans.MonthsWide()

	for i, month := range months {
		s := string(trans.MonthWide(time.Month(i + 1)))
		if s != string(month) {
			t.Errorf("Expected '%s' Got '%s'", string(month), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "January",
		},
		{
			idx:      2,
			expected: "February",
		},
		{
			idx:      3,
			expected: "March",
		},
		{
			idx:      4,
			expected: "April",
		},
		{
			idx:      5,
			expected: "May",
		},
		{
			idx:      6,
			expected: "June",
		},
		{
			idx:      7,
			expected: "July",
		},
		{
			idx:      8,
			expected: "August",
		},
		{
			idx:      9,
			expected: "September",
		},
		{
			idx:      10,
			expected: "October",
		},
		{
			idx:      11,
			expected: "November",
		},
		{
			idx:      12,
			expected: "December",
		},
	}

	for _, tt := range tests {
		s := string(trans.MonthWide(time.Month(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

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
