package sw_CD

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sw_CD struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'sw_CD' locale
func New() locales.Translator {
	return &sw_CD{
		locale:                 "sw_CD",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "FC", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "mkw", "mpi", "mtu", "min", "mtn", "mst", "msb", "mun", "mts", "mku", "mkm", "mkb"},
		monthsNarrow:           []string{"", "k", "p", "t", "i", "t", "s", "s", "m", "t", "k", "m", "m"},
		monthsWide:             []string{"", "mwezi ya kwanja", "mwezi ya pili", "mwezi ya tatu", "mwezi ya ine", "mwezi ya tanu", "mwezi ya sita", "mwezi ya saba", "mwezi ya munane", "mwezi ya tisa", "mwezi ya kumi", "mwezi ya kumi na moya", "mwezi ya kumi ya mbili"},
		daysAbbreviated:        []string{"yen", "kwa", "pil", "tat", "ine", "tan", "sit"},
		daysNarrow:             []string{"y", "k", "p", "t", "i", "t", "s"},
		daysShort:              []string{"Jumapili", "Jumatatu", "Jumanne", "Jumatano", "Alhamisi", "Ijumaa", "Jumamosi"},
		daysWide:               []string{"siku ya yenga", "siku ya kwanza", "siku ya pili", "siku ya tatu", "siku ya ine", "siku ya tanu", "siku ya sita"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"am", "pm"},
		periodsWide:            []string{"ya asubuyi", "ya muchana"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Kabla ya Kristo", "Baada ya Kristo"},
		timezones:              map[string]string{"GMT": "Saa za Greenwich", "∅∅∅": "Saa za Majira ya joto za Peru", "JDT": "Saa za Mchana za Japani", "ARST": "Saa za Majira ya joto za Ajentina", "CAT": "Saa za Afrika ya Kati", "CST": "Saa za Wastani za Kati", "PDT": "Saa za Mchana za Pasifiki", "OEZ": "Saa Wastani za Mashariki mwa Ulaya", "MDT": "Saa za Mchana za Mountain", "HNT": "Saa za Wastani za Newfoundland", "TMST": "Saa za Majira ya joto za Turkmenistan", "ADT": "Saa za Mchana za Atlantiki", "ART": "Saa Wastani za Ajentina", "HKT": "Saa Wastani za Hong Kong", "MYT": "Saa za Malaysia", "COT": "Saa Wastani za Kolombia", "EDT": "Saa za Mchana za Mashariki", "HAST": "Saa za Wastani za Hawaii-Aleutian", "ACWST": "Saa Wastani za Magharibi ya Kati ya Australia", "LHST": "Saa Wastani za Lord Howe", "LHDT": "Saa za Mchana za Lord Howe", "MEZ": "Saa Wastani za Ulaya ya kati", "IST": "Saa Wastani za India", "AEST": "Saa Wastani za Mashariki mwa Australia", "GFT": "Saa za Guiana ya Ufaransa", "ACST": "Saa Wastani za Australia ya Kati", "PST": "Saa za Wastani za Pasifiki", "NZST": "Saa Wastani za Nyuzilandi", "CLST": "Saa za Majira ya joto za Chile", "CDT": "Saa za Mchana za Kati", "ECT": "Saa za Ekwado", "GYT": "Saa za Guyana", "BOT": "Saa za Bolivia", "HAT": "Saa za Mchana za Newfoundland", "TMT": "Saa Wastani za Turkmenistan", "WESZ": "Saa za Majira ya joto za Magharibi mwa Ulaya", "AKDT": "Saa za Mchana za Alaska", "NZDT": "Saa za Mchana za Nyuzilandi", "OESZ": "Saa za Majira ya joto za Mashariki mwa Ulaya", "AEDT": "Saa za Mchana za Mashariki mwa Australia", "SAST": "Saa Wastani za Afrika Kusini", "EST": "Saa za Wastani za Mashariki", "ACWDT": "Saa za Mchana za Magharibi ya Kati ya Australia", "UYST": "Saa za Majira ya joto za Urugwai", "SGT": "Saa Wastani za Singapore", "WIB": "Saa za Magharibi mwa Indonesia", "WART": "Saa Wastani za Magharibi mwa Ajentina", "SRT": "Saa za Suriname", "JST": "Saa Wastani za Japani", "EAT": "Saa za Afrika Mashariki", "WAST": "Saa za Majira ya joto za Afrika Magharibi", "AWDT": "Saa za Mchana za Australia Magharibi", "WAT": "Saa Wastani za Afrika Magharibi", "AWST": "Saa Wastani za Australia Magharibi", "CHAST": "Saa Wastani za Chatham", "HADT": "Saa za Mchana za Hawaii-Aleutian", "ACDT": "Saa za Mchana za Australia ya Kati", "AST": "Saa za Wastani za Atlantiki", "WITA": "Saa za Indonesia ya Kati", "BT": "Saa za Bhutan", "MESZ": "Saa za Majira ya joto za Ulaya ya Kati", "ChST": "Saa Wastani za Chamorro", "COST": "Saa za Majira ya joto za Kolombia", "WEZ": "Saa Wastani za Magharibi mwa Ulaya", "CHADT": "Saa za Mchana za Chatham", "WARST": "Saa za Majira ya joto za Magharibi mwa Ajentina", "VET": "Saa za Venezuela", "UYT": "Saa Wastani za Urugwai", "AKST": "Saa za Wastani za Alaska", "HKST": "Saa za Majira ya joto za Hong Kong", "MST": "Saa za Wastani za Mountain", "WIT": "Saa za Mashariki mwa Indonesia", "CLT": "Saa Wastani za Chile"},
	}
}

// Locale returns the current translators string locale
func (sw *sw_CD) Locale() string {
	return sw.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sw_CD'
func (sw *sw_CD) PluralsCardinal() []locales.PluralRule {
	return sw.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sw_CD'
func (sw *sw_CD) PluralsOrdinal() []locales.PluralRule {
	return sw.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sw_CD'
func (sw *sw_CD) PluralsRange() []locales.PluralRule {
	return sw.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sw_CD'
func (sw *sw_CD) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sw_CD'
func (sw *sw_CD) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sw_CD'
func (sw *sw_CD) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sw.CardinalPluralRule(num1, v1)
	end := sw.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sw *sw_CD) MonthAbbreviated(month time.Month) string {
	return sw.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sw *sw_CD) MonthsAbbreviated() []string {
	return sw.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sw *sw_CD) MonthNarrow(month time.Month) string {
	return sw.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sw *sw_CD) MonthsNarrow() []string {
	return sw.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sw *sw_CD) MonthWide(month time.Month) string {
	return sw.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sw *sw_CD) MonthsWide() []string {
	return sw.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sw *sw_CD) WeekdayAbbreviated(weekday time.Weekday) string {
	return sw.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sw *sw_CD) WeekdaysAbbreviated() []string {
	return sw.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sw *sw_CD) WeekdayNarrow(weekday time.Weekday) string {
	return sw.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sw *sw_CD) WeekdaysNarrow() []string {
	return sw.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sw *sw_CD) WeekdayShort(weekday time.Weekday) string {
	return sw.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sw *sw_CD) WeekdaysShort() []string {
	return sw.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sw *sw_CD) WeekdayWide(weekday time.Weekday) string {
	return sw.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sw *sw_CD) WeekdaysWide() []string {
	return sw.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sw_CD' and handles both Whole and Real numbers based on 'v'
func (sw *sw_CD) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sw_CD' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sw *sw_CD) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sw.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sw_CD'
func (sw *sw_CD) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sw.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sw.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sw_CD'
// in accounting notation.
func (sw *sw_CD) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sw.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, sw.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sw.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sw.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sw.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sw_CD'
func (sw *sw_CD) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sw.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
