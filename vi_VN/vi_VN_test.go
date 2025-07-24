package vi_VN

import (
	"testing"
	"time"

	"github.com/EverlongProject/locales"
	"github.com/EverlongProject/locales/currency"
)

func TestLocale(t *testing.T) {

	trans := New()
	expected := "vi_VN"

	if trans.Locale() != expected {
		t.Errorf("Expected '%s' Got '%s'", expected, trans.Locale())
	}
}

func TestPluralsRange(t *testing.T) {

	trans := New()

	tests := []struct {
		expected locales.PluralRule
	}{
		{
			expected: locales.PluralRuleOther,
		},
	}

	rules := trans.PluralsRange()
	expected := 1
	if len(rules) != expected {
		t.Errorf("Expected '%d' Got '%d'", expected, len(rules))
	}

	for _, tt := range tests {

		r := locales.PluralRuleUnknown

		for i := 0; i < len(rules); i++ {
			if rules[i] == tt.expected {
				r = rules[i]
				break
			}
		}
		if r == locales.PluralRuleUnknown {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, r)
		}
	}
}

func TestPluralsOrdinal(t *testing.T) {

	trans := New()

	tests := []struct {
		expected locales.PluralRule
	}{
		{
			expected: locales.PluralRuleOther,
		},
	}

	rules := trans.PluralsOrdinal()
	expected := 1
	if len(rules) != expected {
		t.Errorf("Expected '%d' Got '%d'", expected, len(rules))
	}

	for _, tt := range tests {

		r := locales.PluralRuleUnknown

		for i := 0; i < len(rules); i++ {
			if rules[i] == tt.expected {
				r = rules[i]
				break
			}
		}
		if r == locales.PluralRuleUnknown {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, r)
		}
	}
}

func TestPluralsCardinal(t *testing.T) {

	trans := New()

	tests := []struct {
		expected locales.PluralRule
	}{
		{
			expected: locales.PluralRuleOther,
		},
	}

	rules := trans.PluralsCardinal()
	expected := 1
	if len(rules) != expected {
		t.Errorf("Expected '%d' Got '%d'", expected, len(rules))
	}

	for _, tt := range tests {

		r := locales.PluralRuleUnknown

		for i := 0; i < len(rules); i++ {
			if rules[i] == tt.expected {
				r = rules[i]
				break
			}
		}
		if r == locales.PluralRuleUnknown {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, r)
		}
	}
}

func TestRangePlurals(t *testing.T) {

	trans := New()

	tests := []struct {
		num1     float64
		v1       uint64
		num2     float64
		v2       uint64
		expected locales.PluralRule
	}{
		{
			num1:     1,
			v1:       1,
			num2:     2,
			v2:       2,
			expected: locales.PluralRuleOther,
		},
	}

	for _, tt := range tests {
		rule := trans.RangePluralRule(tt.num1, tt.v1, tt.num2, tt.v2)
		if rule != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, rule)
		}
	}
}

func TestOrdinalPlurals(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		expected locales.PluralRule
	}{
		{
			num:      1,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      2,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      3,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      4,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      5,
			v:        0,
			expected: locales.PluralRuleOther,
		},
	}

	for _, tt := range tests {
		rule := trans.OrdinalPluralRule(tt.num, tt.v)
		if rule != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, rule)
		}
	}
}

func TestCardinalPlurals(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		expected locales.PluralRule
	}{
		{
			num:      1,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      4,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      21,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      24,
			v:        0,
			expected: locales.PluralRuleOther,
		},
		{
			num:      1.2,
			v:        1,
			expected: locales.PluralRuleOther,
		},
		{
			num:      2.07,
			v:        2,
			expected: locales.PluralRuleOther,
		},
		{
			num:      10.94,
			v:        2,
			expected: locales.PluralRuleOther,
		},
	}

	for _, tt := range tests {
		rule := trans.CardinalPluralRule(tt.num, tt.v)
		if rule != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, rule)
		}
	}
}

func TestDaysAbbreviated(t *testing.T) {

	trans := New()
	days := trans.WeekdaysAbbreviated()

	for i, day := range days {
		s := trans.WeekdayAbbreviated(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", day, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "CN",
		},
		{
			idx:      1,
			expected: "Th 2",
		},
		{
			idx:      2,
			expected: "Th 3",
		},
		{
			idx:      3,
			expected: "Th 4",
		},
		{
			idx:      4,
			expected: "Th 5",
		},
		{
			idx:      5,
			expected: "Th 6",
		},
		{
			idx:      6,
			expected: "Th 7",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayAbbreviated(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysNarrow(t *testing.T) {

	trans := New()
	days := trans.WeekdaysNarrow()

	for i, day := range days {
		s := trans.WeekdayNarrow(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", day, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "CN",
		},
		{
			idx:      1,
			expected: "T2",
		},
		{
			idx:      2,
			expected: "T3",
		},
		{
			idx:      3,
			expected: "T4",
		},
		{
			idx:      4,
			expected: "T5",
		},
		{
			idx:      5,
			expected: "T6",
		},
		{
			idx:      6,
			expected: "T7",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayNarrow(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysShort(t *testing.T) {

	trans := New()
	days := trans.WeekdaysShort()

	for i, day := range days {
		s := trans.WeekdayShort(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", day, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "CN",
		},
		{
			idx:      1,
			expected: "Th 2",
		},
		{
			idx:      2,
			expected: "Th 3",
		},
		{
			idx:      3,
			expected: "Th 4",
		},
		{
			idx:      4,
			expected: "Th 5",
		},
		{
			idx:      5,
			expected: "Th 6",
		},
		{
			idx:      6,
			expected: "Th 7",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayShort(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysWide(t *testing.T) {

	trans := New()
	days := trans.WeekdaysWide()

	for i, day := range days {
		s := trans.WeekdayWide(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", day, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "Chủ Nhật",
		},
		{
			idx:      1,
			expected: "Thứ Hai",
		},
		{
			idx:      2,
			expected: "Thứ Ba",
		},
		{
			idx:      3,
			expected: "Thứ Tư",
		},
		{
			idx:      4,
			expected: "Thứ Năm",
		},
		{
			idx:      5,
			expected: "Thứ Sáu",
		},
		{
			idx:      6,
			expected: "Thứ Bảy",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayWide(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsAbbreviated(t *testing.T) {

	trans := New()
	months := trans.MonthsAbbreviated()

	for i, month := range months {
		s := trans.MonthAbbreviated(time.Month(i + 1))
		if s != month {
			t.Errorf("Expected '%s' Got '%s'", month, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "thg 1",
		},
		{
			idx:      2,
			expected: "thg 2",
		},
		{
			idx:      3,
			expected: "thg 3",
		},
		{
			idx:      4,
			expected: "thg 4",
		},
		{
			idx:      5,
			expected: "thg 5",
		},
		{
			idx:      6,
			expected: "thg 6",
		},
		{
			idx:      7,
			expected: "thg 7",
		},
		{
			idx:      8,
			expected: "thg 8",
		},
		{
			idx:      9,
			expected: "thg 9",
		},
		{
			idx:      10,
			expected: "thg 10",
		},
		{
			idx:      11,
			expected: "thg 11",
		},
		{
			idx:      12,
			expected: "thg 12",
		},
	}

	for _, tt := range tests {
		s := trans.MonthAbbreviated(time.Month(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsNarrow(t *testing.T) {

	trans := New()
	months := trans.MonthsNarrow()

	for i, month := range months {
		s := trans.MonthNarrow(time.Month(i + 1))
		if s != month {
			t.Errorf("Expected '%s' Got '%s'", month, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "1",
		},
		{
			idx:      2,
			expected: "2",
		},
		{
			idx:      3,
			expected: "3",
		},
		{
			idx:      4,
			expected: "4",
		},
		{
			idx:      5,
			expected: "5",
		},
		{
			idx:      6,
			expected: "6",
		},
		{
			idx:      7,
			expected: "7",
		},
		{
			idx:      8,
			expected: "8",
		},
		{
			idx:      9,
			expected: "9",
		},
		{
			idx:      10,
			expected: "10",
		},
		{
			idx:      11,
			expected: "11",
		},
		{
			idx:      12,
			expected: "12",
		},
	}

	for _, tt := range tests {
		s := trans.MonthNarrow(time.Month(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsWide(t *testing.T) {

	trans := New()
	months := trans.MonthsWide()

	for i, month := range months {
		s := trans.MonthWide(time.Month(i + 1))
		if s != month {
			t.Errorf("Expected '%s' Got '%s'", month, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "tháng 1",
		},
		{
			idx:      2,
			expected: "tháng 2",
		},
		{
			idx:      3,
			expected: "tháng 3",
		},
		{
			idx:      4,
			expected: "tháng 4",
		},
		{
			idx:      5,
			expected: "tháng 5",
		},
		{
			idx:      6,
			expected: "tháng 6",
		},
		{
			idx:      7,
			expected: "tháng 7",
		},
		{
			idx:      8,
			expected: "tháng 8",
		},
		{
			idx:      9,
			expected: "tháng 9",
		},
		{
			idx:      10,
			expected: "tháng 10",
		},
		{
			idx:      11,
			expected: "tháng 11",
		},
		{
			idx:      12,
			expected: "tháng 12",
		},
	}

	for _, tt := range tests {
		s := trans.MonthWide(time.Month(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeShort(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "09:00",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 3, 21, 0, time.UTC),
			expected: "20:03",
		},
	}

	for _, tt := range tests {
		s := trans.FmtTimeShort(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeMedium(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "09:00:01",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 3, 21, 0, time.UTC),
			expected: "20:03:21",
		},
	}

	for _, tt := range tests {
		s := trans.FmtTimeMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeLong(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "09:00:01 UTC",
		},
		{
			t:        time.Date(2016, 02, 03, 23, 3, 21, 0, time.UTC),
			expected: "23:03:21 UTC",
		},
	}

	for _, tt := range tests {
		s := trans.FmtTimeLong(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeFull(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "09:00:01 Giờ Phối hợp Quốc tế",
		},
		{
			t:        time.Date(2016, 02, 03, 23, 3, 21, 0, time.UTC),
			expected: "23:03:21 Giờ Phối hợp Quốc tế",
		},
	}

	for _, tt := range tests {
		s := trans.FmtTimeFull(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateShort(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "03/02/2016",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 3, 21, 0, time.UTC),
			expected: "03/02/2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29/02/2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29/02/2016",
		},
		{
			t:        time.Date(-500, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29/02/0500",
		},
	}

	for _, tt := range tests {
		s := trans.FmtDateShort(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateMedium(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "3 thg 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 3, 21, 0, time.UTC),
			expected: "3 thg 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29 thg 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29 thg 2, 2016",
		},
		{
			t:        time.Date(-500, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29 thg 2, 0500",
		},
	}

	for _, tt := range tests {
		s := trans.FmtDateMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateLong(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "3 tháng 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 3, 21, 0, time.UTC),
			expected: "3 tháng 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29 tháng 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29 tháng 2, 2016",
		},
		{
			t:        time.Date(-500, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "29 tháng 2, 0500",
		},
	}

	for _, tt := range tests {
		s := trans.FmtDateLong(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateFull(t *testing.T) {

	trans := New()

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "Thứ Tư, 3 tháng 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 3, 21, 0, time.UTC),
			expected: "Thứ Tư, 3 tháng 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "Thứ Hai, 29 tháng 2, 2016",
		},
		{
			t:        time.Date(2016, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "Thứ Hai, 29 tháng 2, 2016",
		},
		{
			t:        time.Date(-500, 02, 29, 4, 3, 21, 0, time.UTC),
			expected: "Thứ Hai, 29 tháng 2, 0500",
		},
	}

	for _, tt := range tests {
		s := trans.FmtDateFull(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtNumber(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      1123456.789,
			v:        2,
			expected: "1.123.456,79",
		},
		{
			num:      1123456.789,
			v:        1,
			expected: "1.123.456,8",
		},
		{
			num:      221123456.789,
			v:        2,
			expected: "221.123.456,79",
		},
		{
			num:      -221123456.789,
			v:        2,
			expected: "-221.123.456,79",
		},
		{
			num:      -221123456.789,
			v:        2,
			expected: "-221.123.456,79",
		},
		{
			num:      0,
			v:        2,
			expected: "0,00",
		},
		{
			num:      -0,
			v:        2,
			expected: "0,00",
		},
		{
			num:      -0,
			v:        2,
			expected: "0,00",
		},
	}

	for _, tt := range tests {
		s := trans.FmtNumber(tt.num, tt.v)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtCurrency(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		currency currency.Type
		expected string
	}{
		{
			num:      1123456.789,
			v:        2,
			currency: currency.VND,
			expected: "1.123.457 ₫",
		},
		{
			num:      1123456.789,
			v:        1,
			currency: currency.VND,
			expected: "1.123.456,8 ₫",
		},
		{
			num:      221123456.789,
			v:        2,
			currency: currency.VND,
			expected: "221.123.457 ₫",
		},
		{
			num:      -221123456.789,
			v:        2,
			currency: currency.VND,
			expected: "-221.123.457 ₫",
		},
		{
			num:      -221123456.789,
			v:        2,
			currency: currency.VND,
			expected: "-221.123.457 ₫",
		},
		{
			num:      0,
			v:        2,
			currency: currency.VND,
			expected: "0 ₫",
		},
		{
			num:      -0,
			v:        2,
			currency: currency.VND,
			expected: "0 ₫",
		},
		{
			num:      -0,
			v:        2,
			currency: currency.VND,
			expected: "0 ₫",
		},
		{
			num:      1.23,
			v:        0,
			currency: currency.VND,
			expected: "1 ₫",
		},
	}

	for _, tt := range tests {
		s := trans.FmtCurrency(tt.num, tt.v, tt.currency)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtAccounting(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		currency currency.Type
		expected string
	}{
		{
			num:      1123456.789,
			v:        2,
			currency: currency.VND,
			expected: "1.123.457 ₫",
		},
		{
			num:      1123456.789,
			v:        1,
			currency: currency.VND,
			expected: "1.123.456,8 ₫",
		},
		{
			num:      221123456.789,
			v:        2,
			currency: currency.VND,
			expected: "221.123.457 ₫",
		},
		{
			num:      -221123456.789,
			v:        2,
			currency: currency.VND,
			expected: "(221.123.457 ₫)",
		},
		{
			num:      -221123456.789,
			v:        2,
			currency: currency.VND,
			expected: "(221.123.457 ₫)",
		},
		{
			num:      0,
			v:        2,
			currency: currency.VND,
			expected: "0 ₫",
		},
		{
			num:      -0,
			v:        2,
			currency: currency.VND,
			expected: "0 ₫",
		},
		{
			num:      -0,
			v:        2,
			currency: currency.VND,
			expected: "0 ₫",
		},
		{
			num:      1.23,
			v:        0,
			currency: currency.VND,
			expected: "1 ₫",
		},
	}

	for _, tt := range tests {
		s := trans.FmtAccounting(tt.num, tt.v, tt.currency)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtPercent(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      15,
			v:        0,
			expected: "1.500%",
		},
		{
			num:      15,
			v:        2,
			expected: "1.500,00%",
		},
		{
			num:      434.45,
			v:        0,
			expected: "43.445%",
		},
		{
			num:      34.41,
			v:        2,
			expected: "3.441,00%",
		},
		{
			num:      -34,
			v:        0,
			expected: "-3.400%",
		},
	}

	for _, tt := range tests {
		s := trans.FmtPercent(tt.num, tt.v)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtMonthDayMedium(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "3 thg 2",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtMonthDayMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtMonthYearMedium(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "thg 2 2016",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtMonthYearMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}
