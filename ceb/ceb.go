package ceb

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ceb struct {
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

// New returns a new instance of translator for the 'ceb' locale
func New() locales.Translator {
	return &ceb{
		locale:                 "ceb",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "₱", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "US $", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ene", "Peb", "Mar", "Abr", "May", "Hun", "Hul", "Ago", "Set", "Okt", "Nob", "Dis"},
		monthsNarrow:           []string{"", "E", "P", "M", "A", "M", "H", "H", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Enero", "Pebrero", "Marso", "Abril", "Mayo", "Hunyo", "Hulyo", "Agosto", "Setyembre", "Oktubre", "Nobyembre", "Disyembre"},
		daysAbbreviated:        []string{"Dom", "Lun", "Mar", "Miy", "Huw", "Biy", "Sab"},
		daysNarrow:             []string{"D", "L", "M", "M", "H", "B", "S"},
		daysShort:              []string{"Dom", "Lun", "Mar", "Miy", "Huw", "Biy", "Sab"},
		daysWide:               []string{"Domingo", "Lunes", "Martes", "Miyerkules", "Huwebes", "Biyernes", "Sabado"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Sa Wala Pa Si Kristo", "Anno Domini"},
		timezones:              map[string]string{"ACDT": "Oras sa Tag-init sa Central Australia", "ACST": "Tamdanang Oras sa Central Australia", "ACWDT": "Oras sa Tag-init sa Central Western Australia", "ACWST": "Tamdanang Oras sa Central Western Australia", "ADT": "Oras sa Tag-init sa Atlantiko", "AEDT": "Oras sa Tag-init sa Eastern Australia", "AEST": "Tamdanang Oras sa Eastern Australia", "AKDT": "Oras sa Tag-init sa Alaska", "AKST": "Tamdanang Oras sa Alaska", "ARST": "Oras sa Tag-init sa Argentina", "ART": "Tamdanang Oras sa Argentina", "AST": "Tamdanang Oras sa Atlantiko", "AWDT": "Oras sa Tag-init sa Western Australia", "AWST": "Tamdanang Oras sa Western Australia", "BOT": "Oras sa Bolivia", "BT": "Oras sa Bhutan", "CAT": "Oras sa Central Africa", "CDT": "Sentrong Oras sa Tag-init", "CHADT": "Oras sa Tag-init sa Chatham", "CHAST": "Tamdanang Oras sa Chatham", "CLST": "Oras sa Tag-init sa Chile", "CLT": "Tamdanang Oras sa Chile", "COST": "Oras sa Tag-init sa Colombia", "COT": "Tamdanang Oras sa Colombia", "CST": "Sentong Tamdanan nga Oras", "ChST": "Tamdanang Oras sa Chamorro", "EAT": "Oras sa East Africa", "ECT": "Oras sa Ecuador", "EDT": "Oras sa Tag-init sa Sidlakan", "EST": "Tamdanan nga Oras sa Sidlakan", "GFT": "Oras sa French Guiana", "GMT": "Oras sa Greenwich Mean", "GST": "Tamdanang Oras sa Gulf", "GYT": "Oras sa Guyana", "HADT": "Oras sa Tag-init sa Hawaii-Aleutian", "HAST": "Tamdanang Oras sa Hawaii-Aleutian", "HAT": "Oras sa Tag-init sa Newfoundland", "HECU": "Oras sa Tag-init sa Cuba", "HEEG": "Oras sa Tag-init sa East Greenland", "HENOMX": "Oras sa Tag-init sa Northwest Mexico", "HEOG": "Oras sa Tag-init sa West Greenland", "HEPM": "Oras sa Tag-init sa St.Pierre & Miquelon", "HEPMX": "Oras sa Tag-init sa Mexican Pacific", "HKST": "Oras sa Tag-init sa Hong Kong", "HKT": "Tamdanang Oras sa Hong Kong", "HNCU": "Tamdanang Oras sa Cuba", "HNEG": "Tamdanang Oras sa East Greenland", "HNNOMX": "Tamdanang Oras sa Northwest Mexico", "HNOG": "Tamdanang Oras sa West Greenland", "HNPM": "Tamdanang Oras sa St. Pierre & Miquelon", "HNPMX": "Tamdanang Oras sa Mexican Pacific", "HNT": "Tamdanang Oras sa Newfoundland", "IST": "Tamdanang Oras sa India", "JDT": "Oras sa Adlawan sa Japan", "JST": "Tamdanang Oras sa Japan", "LHDT": "Oras sa Tag-init sa Lord Howe", "LHST": "Tamdanang Oras sa Lord Howe", "MDT": "Oras sa Tag-init sa Kabukiran", "MESZ": "Oras sa Tag-init sa Central Europe", "MEZ": "Tamdanang Oras sa Central Europe", "MST": "Tamdanang Oras sa Kabukiran", "MYT": "Oras sa Malaysia", "NZDT": "Oras sa Tag-init sa New Zealand", "NZST": "Tamdanang Oras sa New Zealand", "OESZ": "Oras sa Tag-init sa Eastern Europe", "OEZ": "Tamdanang Oras sa Eastern Europe", "PDT": "Oras sa Tag-init sa Pasipiko", "PST": "Tamdanang Oras sa Pasipiko", "SAST": "Tamdanang Oras sa South Africa", "SGT": "Tamdanang Oras sa Singapore", "SRT": "Oras sa Suriname", "TMST": "Oras sa Tag-init sa Turkmenistan", "TMT": "Tamdanang Oras sa Turkmenistan", "UYST": "Oras sa Tag-init sa Uruguay", "UYT": "Tamdanang Oras sa Uruguay", "VET": "Oras sa Venezuela", "WARST": "Oras sa Tag-init sa Western Argentina", "WART": "Tamdanang Oras sa Western Argentina", "WAST": "Oras sa Tag-init sa West Africa", "WAT": "Tamdanang Oras sa West Africa", "WESZ": "Oras sa Tag-init sa Western Europe", "WEZ": "Tamdanang Oras sa Western Europe", "WIB": "Oras sa Western Indonesia", "WIT": "Oras sa Eastern Indonesia", "WITA": "Oras sa Central Indonesia", "∅∅∅": "Oras sa Tag-init sa Peru"},
	}
}

// Locale returns the current translators string locale
func (ceb *ceb) Locale() string {
	return ceb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ceb'
func (ceb *ceb) PluralsCardinal() []locales.PluralRule {
	return ceb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ceb'
func (ceb *ceb) PluralsOrdinal() []locales.PluralRule {
	return ceb.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ceb'
func (ceb *ceb) PluralsRange() []locales.PluralRule {
	return ceb.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ceb'
func (ceb *ceb) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	fMod10 := f % 10

	if (v == 0 && (i == 1 || i == 2 || i == 3)) || (v == 0 && (iMod10 != 4 && iMod10 != 6 && iMod10 != 9)) || (v != 0 && (fMod10 != 4 && fMod10 != 6 && fMod10 != 9)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ceb'
func (ceb *ceb) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ceb'
func (ceb *ceb) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ceb *ceb) MonthAbbreviated(month time.Month) string {
	return ceb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ceb *ceb) MonthsAbbreviated() []string {
	return ceb.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ceb *ceb) MonthNarrow(month time.Month) string {
	return ceb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ceb *ceb) MonthsNarrow() []string {
	return ceb.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ceb *ceb) MonthWide(month time.Month) string {
	return ceb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ceb *ceb) MonthsWide() []string {
	return ceb.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ceb *ceb) WeekdayAbbreviated(weekday time.Weekday) string {
	return ceb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ceb *ceb) WeekdaysAbbreviated() []string {
	return ceb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ceb *ceb) WeekdayNarrow(weekday time.Weekday) string {
	return ceb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ceb *ceb) WeekdaysNarrow() []string {
	return ceb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ceb *ceb) WeekdayShort(weekday time.Weekday) string {
	return ceb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ceb *ceb) WeekdaysShort() []string {
	return ceb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ceb *ceb) WeekdayWide(weekday time.Weekday) string {
	return ceb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ceb *ceb) WeekdaysWide() []string {
	return ceb.daysWide
}

// Decimal returns the decimal point of number
func (ceb *ceb) Decimal() string {
	return ceb.decimal
}

// Group returns the group of number
func (ceb *ceb) Group() string {
	return ceb.group
}

// Group returns the minus sign of number
func (ceb *ceb) Minus() string {
	return ceb.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ceb' and handles both Whole and Real numbers based on 'v'
func (ceb *ceb) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ceb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ceb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ceb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ceb' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ceb *ceb) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ceb'
func (ceb *ceb) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ceb.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ceb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ceb.group[0])
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
		b = append(b, ceb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ceb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ceb'
// in accounting notation.
func (ceb *ceb) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ceb.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ceb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ceb.group[0])
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

		b = append(b, ceb.currencyNegativePrefix[0])

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
			b = append(b, ceb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ceb.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ceb'
func (ceb *ceb) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ceb'
func (ceb *ceb) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ceb.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ceb'
func (ceb *ceb) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ceb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ceb'
func (ceb *ceb) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ceb.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ceb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ceb'
func (ceb *ceb) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ceb'
func (ceb *ceb) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ceb'
func (ceb *ceb) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ceb'
func (ceb *ceb) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ceb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ceb.periodsAbbreviated[0]...)
	} else {
		b = append(b, ceb.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ceb.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
