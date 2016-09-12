package ii_CN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ii_CN struct {
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

// New returns a new instance of translator for the 'ii_CN' locale
func New() locales.Translator {
	return &ii_CN{
		locale:                 "ii_CN",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		timeSeparator:          ":",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "K",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "K",
		monthsAbbreviated:      []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "ꋍꆪ", "ꑍꆪ", "ꌕꆪ", "ꇖꆪ", "ꉬꆪ", "ꃘꆪ", "ꏃꆪ", "ꉆꆪ", "ꈬꆪ", "ꊰꆪ", "ꊰꊪꆪ", "ꊰꑋꆪ"},
		daysAbbreviated:        []string{"ꑭꆏ", "ꆏꋍ", "ꆏꑍ", "ꆏꌕ", "ꆏꇖ", "ꆏꉬ", "ꆏꃘ"},
		daysNarrow:             []string{"ꆏ", "ꋍ", "ꑍ", "ꌕ", "ꇖ", "ꉬ", "ꃘ"},
		daysWide:               []string{"ꑭꆏꑍ", "ꆏꊂꋍ", "ꆏꊂꑍ", "ꆏꊂꌕ", "ꆏꊂꇖ", "ꆏꊂꉬ", "ꆏꊂꃘ"},
		periodsAbbreviated:     []string{"ꎸꄑ", "ꁯꋒ"},
		periodsWide:            []string{"ꎸꄑ", "ꁯꋒ"},
		erasAbbreviated:        []string{"ꃅꋊꂿ", "ꃅꋊꊂ"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"COT": "COT", "GMT": "GMT", "AEST": "AEST", "CHAST": "CHAST", "NZDT": "NZDT", "ECT": "ECT", "WAT": "WAT", "HAT": "HAT", "GYT": "GYT", "MYT": "MYT", "CST": "CST", "ACST": "ACST", "CAT": "CAT", "SGT": "SGT", "MST": "MST", "HAST": "HAST", "AKST": "AKST", "AKDT": "AKDT", "AWST": "AWST", "SAST": "SAST", "EAT": "EAT", "UYT": "UYT", "PDT": "PDT", "ARST": "ARST", "TMT": "TMT", "BT": "BT", "UYST": "UYST", "∅∅∅": "∅∅∅", "HKT": "HKT", "ART": "ART", "WAST": "WAST", "OEZ": "OEZ", "AST": "AST", "WITA": "WITA", "OESZ": "OESZ", "SRT": "SRT", "GFT": "GFT", "CHADT": "CHADT", "BOT": "BOT", "LHST": "LHST", "MDT": "MDT", "HKST": "HKST", "ADT": "ADT", "ACWST": "ACWST", "AWDT": "AWDT", "CLT": "CLT", "WARST": "WARST", "ACWDT": "ACWDT", "EDT": "EDT", "MEZ": "MEZ", "WEZ": "WEZ", "HNT": "HNT", "ChST": "ChST", "LHDT": "LHDT", "CDT": "CDT", "IST": "IST", "COST": "COST", "WIB": "WIB", "AEDT": "AEDT", "JST": "JST", "MESZ": "MESZ", "NZST": "NZST", "WIT": "WIT", "WART": "WART", "VET": "VET", "HADT": "HADT", "WESZ": "WESZ", "TMST": "TMST", "CLST": "CLST", "PST": "PST", "ACDT": "ACDT", "JDT": "JDT", "EST": "EST"},
	}
}

// Locale returns the current translators string locale
func (ii *ii_CN) Locale() string {
	return ii.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ii_CN'
func (ii *ii_CN) PluralsCardinal() []locales.PluralRule {
	return ii.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ii_CN'
func (ii *ii_CN) PluralsOrdinal() []locales.PluralRule {
	return ii.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ii_CN'
func (ii *ii_CN) PluralsRange() []locales.PluralRule {
	return ii.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ii_CN'
func (ii *ii_CN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ii_CN'
func (ii *ii_CN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ii_CN'
func (ii *ii_CN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ii *ii_CN) MonthAbbreviated(month time.Month) string {
	return ii.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ii *ii_CN) MonthsAbbreviated() []string {
	return ii.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ii *ii_CN) MonthNarrow(month time.Month) string {
	return ii.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ii *ii_CN) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (ii *ii_CN) MonthWide(month time.Month) string {
	return ii.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ii *ii_CN) MonthsWide() []string {
	return ii.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ii *ii_CN) WeekdayAbbreviated(weekday time.Weekday) string {
	return ii.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ii *ii_CN) WeekdaysAbbreviated() []string {
	return ii.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ii *ii_CN) WeekdayNarrow(weekday time.Weekday) string {
	return ii.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ii *ii_CN) WeekdaysNarrow() []string {
	return ii.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ii *ii_CN) WeekdayShort(weekday time.Weekday) string {
	return ii.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ii *ii_CN) WeekdaysShort() []string {
	return ii.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ii *ii_CN) WeekdayWide(weekday time.Weekday) string {
	return ii.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ii *ii_CN) WeekdaysWide() []string {
	return ii.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ii_CN' and handles both Whole and Real numbers based on 'v'
func (ii *ii_CN) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ii_CN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ii *ii_CN) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ii_CN'
func (ii *ii_CN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ii.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(ii.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ii.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ii.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ii.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ii_CN'
// in accounting notation.
func (ii *ii_CN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ii.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ii.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ii.currencyNegativePrefix[j])
		}

		b = append(b, ii.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ii.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ii.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, ii.currencyNegativeSuffix...)
	} else {

		b = append(b, ii.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ii_CN'
func (ii *ii_CN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ii.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
