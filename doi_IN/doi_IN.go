package doi_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type doi_IN struct {
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

// New returns a new instance of translator for the 'doi_IN' locale
func New() locales.Translator {
	return &doi_IN{
		locale:             "doi_IN",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "जन.", "फर.", "मार्च", "अप्रैल", "मेई", "जून", "जुलाई", "अग.", "सित.", "अक्तू.", "नव.", "दिस."},
		monthsNarrow:       []string{"", "ज", "फ", "मा", "अ", "मे", "जू", "जु", "अ", "सि", "अ", "न", "दि"},
		monthsWide:         []string{"", "जनवरी", "फरवरी", "मार्च", "अप्रैल", "मेई", "जून", "जुलाई", "अगस्त", "सितंबर", "अत्तूबर", "नवंबर", "दिसंबर"},
		daysAbbreviated:    []string{"ऐत", "सोम", "मंगल", "बुध", "बीर", "शुक्र", "शनि"},
		daysNarrow:         []string{"ऐ.", "सो.", "म.", "बु.", "बी.", "शु.", "श."},
		daysShort:          []string{"ऐत", "सोम", "मंगल", "बुध", "बीर", "शुक्र", "शनि"},
		daysWide:           []string{"ऐतबार", "सोमबार", "मंगलबार", "बुधबार", "बीरबार", "शुक्रबार", "शनीबार"},
		periodsAbbreviated: []string{"सवेर", "स’ञ"},
		periodsNarrow:      []string{"सवेर", "स’ञ"},
		periodsWide:        []string{"सवेर", "बाद दपैहर"},
		erasAbbreviated:    []string{"", ""},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"", ""},
		timezones:          map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "अटलांटिक डेलाइट समां", "AEDT": "AEDT", "AEST": "AEST", "AKDT": "AKDT", "AKST": "AKST", "ARST": "ARST", "ART": "ART", "AST": "अटलांटिक मानक समां", "AWDT": "AWDT", "AWST": "AWST", "BOT": "BOT", "BT": "BT", "CAT": "CAT", "CDT": "उत्तरी अमरीकी डेलाइट केंदरी समां", "CHADT": "CHADT", "CHAST": "CHAST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "उत्तरी अमरीकी मानक केंदरी समां", "ChST": "ChST", "EAT": "EAT", "ECT": "ECT", "EDT": "उत्तरी अमरीकी डेलाइट पूर्वी समां", "EST": "उत्तरी अमरीकी मानक पूर्वी समां", "GFT": "GFT", "GMT": "ग्रीनविच मीन टाइम", "GST": "GST", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HAT": "HAT", "HECU": "HECU", "HEEG": "HEEG", "HENOMX": "HENOMX", "HEOG": "HEOG", "HEPM": "HEPM", "HEPMX": "HEPMX", "HKST": "HKST", "HKT": "HKT", "HNCU": "HNCU", "HNEG": "HNEG", "HNNOMX": "HNNOMX", "HNOG": "HNOG", "HNPM": "HNPM", "HNPMX": "HNPMX", "HNT": "HNT", "IST": "IST", "JDT": "JDT", "JST": "JST", "LHDT": "LHDT", "LHST": "LHST", "MDT": "उत्तरी अमरीकी डेलाइट माउंटेन समां", "MESZ": "केंदरी यूरोपी गर्मियें दा समां", "MEZ": "केंदरी यूरोपी मानक समां", "MST": "उत्तरी अमरीकी मानक माउंटेन समां", "MYT": "MYT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "उत्तरी यूरोपी गर्मियें दा समां", "OEZ": "उत्तरी यूरोपी मानक समां", "PDT": "उत्तरी अमरीकी डेलाइट प्रशांत समां", "PST": "उत्तरी अमरीकी मानक प्रशांत समां", "SAST": "SAST", "SGT": "SGT", "SRT": "SRT", "TMST": "TMST", "TMT": "TMT", "UYST": "UYST", "UYT": "UYT", "VET": "VET", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "पच्छमी यूरोपी गर्मियें दा समां", "WEZ": "पच्छमी यूरोपी मानक समां", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (doi *doi_IN) Locale() string {
	return doi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'doi_IN'
func (doi *doi_IN) PluralsCardinal() []locales.PluralRule {
	return doi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'doi_IN'
func (doi *doi_IN) PluralsOrdinal() []locales.PluralRule {
	return doi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'doi_IN'
func (doi *doi_IN) PluralsRange() []locales.PluralRule {
	return doi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'doi_IN'
func (doi *doi_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'doi_IN'
func (doi *doi_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'doi_IN'
func (doi *doi_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (doi *doi_IN) MonthAbbreviated(month time.Month) string {
	return doi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (doi *doi_IN) MonthsAbbreviated() []string {
	return doi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (doi *doi_IN) MonthNarrow(month time.Month) string {
	return doi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (doi *doi_IN) MonthsNarrow() []string {
	return doi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (doi *doi_IN) MonthWide(month time.Month) string {
	return doi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (doi *doi_IN) MonthsWide() []string {
	return doi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (doi *doi_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return doi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (doi *doi_IN) WeekdaysAbbreviated() []string {
	return doi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (doi *doi_IN) WeekdayNarrow(weekday time.Weekday) string {
	return doi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (doi *doi_IN) WeekdaysNarrow() []string {
	return doi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (doi *doi_IN) WeekdayShort(weekday time.Weekday) string {
	return doi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (doi *doi_IN) WeekdaysShort() []string {
	return doi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (doi *doi_IN) WeekdayWide(weekday time.Weekday) string {
	return doi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (doi *doi_IN) WeekdaysWide() []string {
	return doi.daysWide
}

// Decimal returns the decimal point of number
func (doi *doi_IN) Decimal() string {
	return doi.decimal
}

// Group returns the group of number
func (doi *doi_IN) Group() string {
	return doi.group
}

// Group returns the minus sign of number
func (doi *doi_IN) Minus() string {
	return doi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'doi_IN' and handles both Whole and Real numbers based on 'v'
func (doi *doi_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, doi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, doi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, doi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'doi_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (doi *doi_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, doi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, doi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, doi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'doi_IN'
func (doi *doi_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := doi.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, doi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, doi.group[0])
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
		b = append(b, doi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, doi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'doi_IN'
// in accounting notation.
func (doi *doi_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := doi.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, doi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, doi.group[0])
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

		b = append(b, doi.minus[0])

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
			b = append(b, doi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, doi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, doi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, doi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, doi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, doi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, doi.periodsAbbreviated[0]...)
	} else {
		b = append(b, doi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, doi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, doi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, doi.periodsAbbreviated[0]...)
	} else {
		b = append(b, doi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, doi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, doi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, doi.periodsAbbreviated[0]...)
	} else {
		b = append(b, doi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'doi_IN'
func (doi *doi_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, doi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, doi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, doi.periodsAbbreviated[0]...)
	} else {
		b = append(b, doi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := doi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
