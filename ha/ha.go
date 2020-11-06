package ha

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ha struct {
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
	currencyPositivePrefix string
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'ha' locale
func New() locales.Translator {
	return &ha{
		locale:                 "ha",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "₦", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "D",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "D",
		monthsAbbreviated:      []string{"", "Jan", "Fab", "Mar", "Afi", "May", "Yun", "Yul", "Agu", "Sat", "Okt", "Nuw", "Dis"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "Y", "Y", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Janairu", "Faburairu", "Maris", "Afirilu", "Mayu", "Yuni", "Yuli", "Agusta", "Satumba", "Oktoba", "Nuwamba", "Disamba"},
		daysAbbreviated:        []string{"Lah", "Lit", "Tal", "Lar", "Alh", "Jum", "Asa"},
		daysNarrow:             []string{"L", "L", "T", "L", "A", "J", "A"},
		daysShort:              []string{"Lh", "Li", "Ta", "Lr", "Al", "Ju", "As"},
		daysWide:               []string{"Lahadi", "Litinin", "Talata", "Laraba", "Alhamis", "Jummaʼa", "Asabar"},
		periodsAbbreviated:     []string{"SF", "YM"},
		periodsNarrow:          []string{"SF", "YM"},
		periodsWide:            []string{"Safiya", "Yamma"},
		erasAbbreviated:        []string{"K.H", "BHAI"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "Lokacin Rana na Tsakiyar Austiraliya", "ACST": "Tsayayyen Lokacin Tsakiyar Austiraliya", "ACWDT": "Lokacin Rana na Yammacin Tsakiyar Austiraliya", "ACWST": "Tsayayyen Lokacin Yammacin Tsakiyar Austiraliya", "ADT": "Lokacin Rana na Kanada, Puerto Rico da Virgin Islands", "AEDT": "Lokacin Rana na Gabashin Austiraliya", "AEST": "Australian Eastern Standard Time", "AKDT": "Lokacin Rana na Alaska", "AKST": "Tsayayyen Lokacin Alaska", "ARST": "Lokacin Bazara na Argentina", "ART": "Tsayayyen Lokacin Argentina", "AST": "Tsayayyen Lokacin Kanada, Puerto Rico da Virgin Islands", "AWDT": "Lokacin Rana na Yammacin Austiralia", "AWST": "Tsayayyen Lokacin Yammacin Austiralia", "BOT": "Lokacin Bolivia", "BT": "Bhutan Time", "CAT": "Lokacin Afirka ta Tsakiya", "CDT": "Lokacin Rana dake Arewacin Amurika ta Tsakiya", "CHADT": "Lokacin Rana na Chatham", "CHAST": "Tsayayyen Lokacin Chatham", "CLST": "Lokacin Bazara na Chile", "CLT": "Tsayayyen Lokacin Chile", "COST": "Lokacin Bazara na Colombia", "COT": "Tsayayyen Lokacin Colombia", "CST": "Tsayayyen Lokaci dake Arewacin Amurika ta Tsakiya", "ChST": "Tsayayyen Lokacin Chamorro", "EAT": "Lokacin Gabashin Afirka", "ECT": "Lokacin Ecuador", "EDT": "Lokacin Rana ta Gabas dake Arewacin Amurika", "EST": "Tsayayyen Lokacin Gabas dake Arewacin Amurika", "GFT": "Lokacin French Guiana", "GMT": "Lokacin Greenwhich a London", "GST": "Lokacin Golf", "GYT": "Lokacin Guyana", "HADT": "Lokacin Rana na Hawaii-Aleutian", "HAST": "Tsayayyen Lokacin Hawaii-Aleutian", "HAT": "Lokaci rana ta Newfoundland", "HECU": "Lokacin Rana na Kuba", "HEEG": "Lokacin Rana na Gabashin Greenland", "HENOMX": "Lokacin Rana na Arewa Maso Yammacin Mekziko", "HEOG": "Lokacin Rana na Yammacin Greenland", "HEPM": "Lokacin Rana na St. Pierre da Miquelon", "HEPMX": "Lokacin Rana na Mekziko Pacific", "HKST": "Lokacin Bazara na Hong Kong", "HKT": "Tsayayyen Lokacin Hong Kong", "HNCU": "Tsayayyen Lokacin Kuba", "HNEG": "Tsayayyen Lokacin Gabashin Greenland", "HNNOMX": "Tsayayyen Lokacin Arewa Maso Yammacin Mekziko", "HNOG": "Tsayayyen Lokacin Yammacin Greenland", "HNPM": "Tsayayyen Lokacin St. Pierre da Miquelon", "HNPMX": "Tsayayyen Lokacin Mekziko Pacific", "HNT": "Lokaci Tsayayye ta Newfoundland", "IST": "India Standard Time", "JDT": "Japan Daylight Time", "JST": "Japan Standard Time", "LHDT": "Lokacin Rana na Vote Lord Howe", "LHST": "Tsayayyen Lokacin Lord Howe", "MDT": "Lokacin Rana na Tsaunin Arewacin Amurka", "MESZ": "Tsakiyar bazara a lokaci turai", "MEZ": "Ida Tsakiyar a Lokaci Turai", "MST": "Lokaci tsayayye na tsauni a Arewacin Amurica", "MYT": "Lokacin Malaysia", "NZDT": "Lokacin Rana na New Zealand", "NZST": "Tsayayyen Lokacin New Zealand", "OESZ": "Gabas a lokaci turai da bazara", "OEZ": "Ida lokaci a turai gabas", "PDT": "Lokacin Rana na Arewacin Amurka", "PST": "Lokaci Tsayayye na Arewacin Amurika", "SAST": "South Africa Standard Time", "SGT": "Tsayayyen Lokacin Singapore", "SRT": "Lokacin Suriname", "TMST": "Turkmenistan Summer Time", "TMT": "Tsayayyen Lokacin Turkmenistan", "UYST": "Lokacin Bazara na Uruguay", "UYT": "Tsayayyen Lokacin Uruguay", "VET": "Lokacin Venezuela", "WARST": "Lokacin Bazara na Yammacin Argentina", "WART": "Tsayayyen Lokacin Yammacin Argentina", "WAST": "Lokacin Bazara na Afirka ta Yamma", "WAT": "Tsayayyen Lokacin Afirka ta Yamma", "WESZ": "Ida lokaci ta yammacin turai da bazara", "WEZ": "Ida lokaci ta yammacin turai", "WIB": "Lokacin Yammacin Indonesia", "WIT": "Eastern Indonesia Time", "WITA": "Lokacin Indonesia ta Tsakiya", "∅∅∅": "Lokacin Bazara na Brasillia"},
	}
}

// Locale returns the current translators string locale
func (ha *ha) Locale() string {
	return ha.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ha'
func (ha *ha) PluralsCardinal() []locales.PluralRule {
	return ha.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ha'
func (ha *ha) PluralsOrdinal() []locales.PluralRule {
	return ha.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ha'
func (ha *ha) PluralsRange() []locales.PluralRule {
	return ha.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ha'
func (ha *ha) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ha'
func (ha *ha) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ha'
func (ha *ha) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ha *ha) MonthAbbreviated(month time.Month) string {
	return ha.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ha *ha) MonthsAbbreviated() []string {
	return ha.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ha *ha) MonthNarrow(month time.Month) string {
	return ha.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ha *ha) MonthsNarrow() []string {
	return ha.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ha *ha) MonthWide(month time.Month) string {
	return ha.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ha *ha) MonthsWide() []string {
	return ha.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ha *ha) WeekdayAbbreviated(weekday time.Weekday) string {
	return ha.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ha *ha) WeekdaysAbbreviated() []string {
	return ha.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ha *ha) WeekdayNarrow(weekday time.Weekday) string {
	return ha.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ha *ha) WeekdaysNarrow() []string {
	return ha.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ha *ha) WeekdayShort(weekday time.Weekday) string {
	return ha.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ha *ha) WeekdaysShort() []string {
	return ha.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ha *ha) WeekdayWide(weekday time.Weekday) string {
	return ha.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ha *ha) WeekdaysWide() []string {
	return ha.daysWide
}

// Decimal returns the decimal point of number
func (ha *ha) Decimal() string {
	return ha.decimal
}

// Group returns the group of number
func (ha *ha) Group() string {
	return ha.group
}

// Group returns the minus sign of number
func (ha *ha) Minus() string {
	return ha.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ha' and handles both Whole and Real numbers based on 'v'
func (ha *ha) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ha' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ha *ha) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ha'
func (ha *ha) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ha.currencies[currency]
	l := len(s) + len(symbol) + 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(ha.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ha.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ha.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ha.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ha'
// in accounting notation.
func (ha *ha) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ha.currencies[currency]
	l := len(s) + len(symbol) + 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ha.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ha.currencyNegativePrefix[j])
		}

		b = append(b, ha.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ha.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ha.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, ha.currencyNegativeSuffix...)
	} else {

		b = append(b, ha.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ha'
func (ha *ha) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ha'
func (ha *ha) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ha'
func (ha *ha) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ha'
func (ha *ha) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ha.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ha'
func (ha *ha) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ha'
func (ha *ha) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ha'
func (ha *ha) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ha'
func (ha *ha) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ha.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
