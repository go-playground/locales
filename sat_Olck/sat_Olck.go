package sat_Olck

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sat_Olck struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'sat_Olck' locale
func New() locales.Translator {
	return &sat_Olck{
		locale:            "sat_Olck",
		pluralsCardinal:   []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:    nil,
		pluralsRange:      nil,
		decimal:           ".",
		group:             ",",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "ᱡᱟᱱ", "ᱯᱷᱟ", "ᱢᱟᱨ", "ᱟᱯᱨ", "ᱢᱮ", "ᱡᱩᱱ", "ᱡᱩᱞ", "ᱟᱜᱟ", "ᱥᱮᱯ", "ᱚᱠᱴ", "ᱱᱟᱣ", "ᱫᱤᱥ"},
		monthsNarrow:      []string{"", "ᱡ", "ᱯ", "ᱢ", "ᱟ", "ᱢ", "ᱡ", "ᱡ", "ᱟ", "ᱥ", "ᱚ", "ᱱ", "ᱫ"},
		monthsWide:        []string{"", "ᱡᱟᱱᱣᱟᱨᱤ", "ᱯᱷᱟᱨᱣᱟᱨᱤ", "ᱢᱟᱨᱪ", "ᱟᱯᱨᱮᱞ", "ᱢᱮ", "ᱡᱩᱱ", "ᱡᱩᱞᱟᱭ", "ᱟᱜᱟᱥᱛ", "ᱥᱮᱯᱴᱮᱢᱵᱟᱨ", "ᱚᱠᱴᱚᱵᱟᱨ", "ᱱᱟᱣᱟᱢᱵᱟᱨ", "ᱫᱤᱥᱟᱢᱵᱟᱨ"},
		daysAbbreviated:   []string{"ᱥᱤᱸ", "ᱚᱛ", "ᱵᱟ", "ᱥᱟᱹ", "ᱥᱟᱹᱨ", "ᱡᱟᱹ", "ᱧᱩ"},
		daysNarrow:        []string{"ᱥ", "ᱚ", "ᱵ", "ᱥ", "ᱥ", "ᱡ", "ᱧ"},
		daysWide:          []string{"ᱥᱤᱸᱜᱮ", "ᱚᱛᱮ", "ᱵᱟᱞᱮ", "ᱥᱟᱹᱜᱩᱱ", "ᱥᱟᱹᱨᱫᱤ", "ᱡᱟᱹᱨᱩᱢ", "ᱧᱩᱦᱩᱢ"},
		periodsWide:       []string{"ᱥᱮᱛᱟᱜ", "ᱧᱤᱫᱟᱹ"},
		erasAbbreviated:   []string{"ᱥᱮᱨᱢᱟ ᱞᱟᱦᱟ", "ᱤᱥᱣᱤ"},
		erasNarrow:        []string{"", ""},
		erasWide:          []string{"", ""},
		timezones:         map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "ᱮᱴᱞᱟᱱᱴᱤᱠ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "AEDT": "AEDT", "AEST": "AEST", "AKDT": "AKDT", "AKST": "AKST", "ARST": "ARST", "ART": "ART", "AST": "ᱮᱴᱞᱟᱱᱴᱤᱠ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "AWDT": "AWDT", "AWST": "AWST", "BOT": "BOT", "BT": "BT", "CAT": "CAT", "CDT": "ᱛᱟᱱᱟᱞᱟ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "CHADT": "CHADT", "CHAST": "CHAST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "ᱛᱟᱱᱟᱞᱟ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "ChST": "ChST", "EAT": "EAT", "ECT": "ECT", "EDT": "ᱤᱥᱴᱟᱨᱱ ᱥᱤᱧᱟᱜ ᱵᱚᱠᱛᱚ", "EST": "ᱤᱥᱴᱟᱨᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "GFT": "GFT", "GMT": "ᱜᱨᱤᱱᱣᱤᱪ ᱢᱤᱱ ᱚᱠᱛᱚ", "GST": "GST", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HAT": "HAT", "HECU": "HECU", "HEEG": "HEEG", "HENOMX": "HENOMX", "HEOG": "HEOG", "HEPM": "HEPM", "HEPMX": "HEPMX", "HKST": "HKST", "HKT": "HKT", "HNCU": "HNCU", "HNEG": "HNEG", "HNNOMX": "HNNOMX", "HNOG": "HNOG", "HNPM": "HNPM", "HNPMX": "HNPMX", "HNT": "HNT", "IST": "IST", "JDT": "JDT", "JST": "JST", "LHDT": "LHDT", "LHST": "LHST", "MDT": "ᱢᱟᱩᱱᱴᱮᱱ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "MESZ": "ᱥᱮᱱᱴᱨᱟᱞ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "MEZ": "ᱥᱮᱱᱴᱨᱟᱞ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "MST": "ᱢᱟᱩᱱᱴᱮᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "MYT": "MYT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "ᱤᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "OEZ": "ᱤᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "PDT": "ᱯᱮᱥᱤᱯᱷᱤᱠ ᱥᱤᱧᱟᱜ ᱚᱠᱛᱚ", "PST": "ᱯᱮᱥᱤᱯᱷᱤᱠ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "SAST": "SAST", "SGT": "SGT", "SRT": "SRT", "TMST": "TMST", "TMT": "TMT", "UYST": "UYST", "UYT": "UYT", "VET": "VET", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "ᱣᱮᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱥᱟᱢᱟᱨ ᱚᱠᱛᱚ", "WEZ": "ᱣᱮᱥᱴᱟᱨᱱ ᱩᱨᱚᱯᱤᱭᱟᱱ ᱮᱥᱴᱮᱱᱰᱟᱨᱰ ᱚᱠᱛᱚ", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (sat *sat_Olck) Locale() string {
	return sat.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sat_Olck'
func (sat *sat_Olck) PluralsCardinal() []locales.PluralRule {
	return sat.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sat_Olck'
func (sat *sat_Olck) PluralsOrdinal() []locales.PluralRule {
	return sat.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sat_Olck'
func (sat *sat_Olck) PluralsRange() []locales.PluralRule {
	return sat.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sat_Olck'
func (sat *sat_Olck) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sat_Olck'
func (sat *sat_Olck) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sat_Olck'
func (sat *sat_Olck) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sat *sat_Olck) MonthAbbreviated(month time.Month) string {
	return sat.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sat *sat_Olck) MonthsAbbreviated() []string {
	return sat.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sat *sat_Olck) MonthNarrow(month time.Month) string {
	return sat.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sat *sat_Olck) MonthsNarrow() []string {
	return sat.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sat *sat_Olck) MonthWide(month time.Month) string {
	return sat.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sat *sat_Olck) MonthsWide() []string {
	return sat.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sat *sat_Olck) WeekdayAbbreviated(weekday time.Weekday) string {
	return sat.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sat *sat_Olck) WeekdaysAbbreviated() []string {
	return sat.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sat *sat_Olck) WeekdayNarrow(weekday time.Weekday) string {
	return sat.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sat *sat_Olck) WeekdaysNarrow() []string {
	return sat.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sat *sat_Olck) WeekdayShort(weekday time.Weekday) string {
	return sat.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sat *sat_Olck) WeekdaysShort() []string {
	return sat.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sat *sat_Olck) WeekdayWide(weekday time.Weekday) string {
	return sat.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sat *sat_Olck) WeekdaysWide() []string {
	return sat.daysWide
}

// Decimal returns the decimal point of number
func (sat *sat_Olck) Decimal() string {
	return sat.decimal
}

// Group returns the group of number
func (sat *sat_Olck) Group() string {
	return sat.group
}

// Group returns the minus sign of number
func (sat *sat_Olck) Minus() string {
	return sat.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sat_Olck' and handles both Whole and Real numbers based on 'v'
func (sat *sat_Olck) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sat_Olck' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sat *sat_Olck) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sat_Olck'
func (sat *sat_Olck) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sat.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sat_Olck'
// in accounting notation.
func (sat *sat_Olck) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sat.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sat.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sat.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sat.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sat.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sat.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sat.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sat_Olck'
func (sat *sat_Olck) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sat.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sat.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sat.periodsAbbreviated[0]...)
	} else {
		b = append(b, sat.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sat.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
