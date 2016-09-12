package sv

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sv struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'sv' locale
func New() locales.Translator {
	return &sv{
		locale:                 "sv",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{6},
		minus:                  "‏−",
		timeSeparator:          ":",
		currencies:             []string{"ADP ", "AED", "AFA ", "AFN", "ALK ", "ALL", "AMD", "ANG", "AOA", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS", "ATS ", "AUD", "AWG", "AZM ", "AZN", "BAD", "BAM", "BAN ", "Bds$", "BDT", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN", "BGO ", "BHD", "BIF", "BM$", "BND", "BOB", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BR$", "BRN ", "BRR ", "BRZ ", "BS$", "BTN", "BUK ", "BWP", "BYB ", "BYR", "BZ$", "CA$", "CDF", "CHE ", "CHF", "CHW ", "CLE ", "CLF ", "CLP", "CNX", "CNY", "COP", "COU ", "CRC", "CSD ", "CSK ", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM ", "DEM ", "DJF", "Dkr", "RD$", "DZD", "ECS ", "ECV ", "Ekr", "EG£", "ERN", "ESA ", "ESB ", "ESP", "ETB", "€", "FIM ", "FJD", "FKP", "FRF ", "GBP", "GEK ", "GEL", "GHC ", "GHS", "GIP", "GMD", "GNF", "GNS ", "GQE ", "GRD ", "GTQ", "GWE ", "GWP ", "GYD", "HKD", "HNL", "HRD ", "HRK", "HTG", "HUF", "IDR", "IEP ", "ILP ", "ILR ", "₪", "INR", "IQD", "IRR", "ISJ", "Ikr", "ITL", "JM$", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH ", "KRO ", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL ", "LTL", "LTT ", "LUC ", "LUF ", "LUL ", "LVL", "LVR ", "LYD", "MAD", "MAF ", "MCF ", "MDC ", "MDL", "MGA", "MGF ", "MKD", "MKN", "MLF ", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP ", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP ", "MXV ", "MYR", "MZE ", "MZM ", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG ", "Nkr", "NPR", "NZD", "OMR", "PAB", "PEI ", "PEN", "PES ", "PGK", "PHP", "PKR", "PLN", "PLZ ", "PTE ", "PYG", "QAR", "RHD ", "ROL ", "RON", "RSD", "RUB", "RUR ", "RWF", "SAR", "SBD", "SCR", "SDD ", "SDG", "SDP ", "kr", "SGD", "SHP", "SIT", "SKK ", "SLL", "SOS", "SRD", "SRG ", "SSP", "STD", "SUR ", "SVC ", "SYP", "SZL", "THB", "TJR ", "TJS", "TMM ", "TMT", "TND", "TOP", "TPE ", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK ", "UGS ", "UGX", "US$", "USN ", "USS ", "UYI ", "UYP ", "UYU", "UZS", "VEB ", "VEF", "VND", "VNN ", "VUV", "WST", "FCFA", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "EC$", "XDR ", "XEU ", "XFO ", "XFU ", "CFA", "XPD ", "CFPF", "XPT ", "XRE ", "XSU ", "XTS", "XUA ", "XXX ", "YDD ", "YER", "YUD", "YUM", "YUN", "YUR ", "ZAL ", "ZAR", "ZMK ", "ZMW", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mars", "apr.", "maj", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "dec."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"},
		daysAbbreviated:        []string{"sön", "mån", "tis", "ons", "tors", "fre", "lör"},
		daysNarrow:             []string{"S", "M", "T", "O", "T", "F", "L"},
		daysShort:              []string{"sö", "må", "ti", "on", "to", "fr", "lö"},
		daysWide:               []string{"söndag", "måndag", "tisdag", "onsdag", "torsdag", "fredag", "lördag"},
		periodsAbbreviated:     []string{"fm", "em"},
		periodsNarrow:          []string{"fm", "em"},
		periodsWide:            []string{"fm", "em"},
		erasAbbreviated:        []string{"f.Kr.", "e.Kr."},
		erasNarrow:             []string{"f.Kr.", "e.Kr."},
		erasWide:               []string{"före Kristus", "efter Kristus"},
		timezones:              map[string]string{"MEZ": "centraleuropeisk normaltid", "TMST": "turkmensk sommartid", "CLST": "chilensk sommartid", "VET": "venezuelansk tid", "MDT": "Macausommartid", "JDT": "japansk sommartid", "HAT": "Newfoundland, sommartid", "ARST": "östargentinsk sommartid", "HNT": "Newfoundland, normaltid", "MYT": "malaysisk tid", "UYT": "uruguayansk normaltid", "WIT": "östindonesisk tid", "BT": "bhutansk tid", "CHAST": "Chatham, normaltid", "ADT": "nordamerikansk atlantsommartid", "WAT": "västafrikansk normaltid", "TMT": "turkmensk normaltid", "COST": "colombiansk sommartid", "ChST": "Chamorrotid", "LHDT": "Lord Howe, sommartid", "WARST": "västargentinsk sommartid", "BOT": "boliviansk tid", "IST": "indisk tid", "AKDT": "Alaska, sommartid", "OESZ": "östeuropeisk sommartid", "SRT": "Surinamtid", "AWST": "västaustralisk normaltid", "ART": "östargentinsk normaltid", "HKT": "Hongkong, normaltid", "EST": "östnordamerikansk normaltid", "WESZ": "västeuropeisk sommartid", "HAST": "Honolulu, normaltid", "OEZ": "östeuropeisk normaltid", "LHST": "Lord Howe, normaltid", "HKST": "Hongkong, sommartid", "AKST": "Alaska, normaltid", "PDT": "västnordamerikansk sommartid", "∅∅∅": "azorisk sommartid", "CHADT": "Chatham, sommartid", "GYT": "Guyanatid", "AWDT": "västaustralisk sommartid", "WEZ": "västeuropeisk normaltid", "AEDT": "östaustralisk sommartid", "PST": "västnordamerikansk normaltid", "EDT": "östnordamerikansk sommartid", "AEST": "östaustralisk normaltid", "WAST": "västafrikansk sommartid", "ECT": "ecuadoriansk tid", "ACWST": "västcentralaustralisk normaltid", "SAST": "sydafrikansk tid", "MESZ": "centraleuropeisk sommartid", "HADT": "Honolulu, sommartid", "SGT": "Singaporetid", "COT": "colombiansk normaltid", "ACST": "centralaustralisk normaltid", "ACDT": "centralaustralisk sommartid", "MST": "Macaunormaltid", "NZST": "nyzeeländsk normaltid", "NZDT": "nyzeeländsk sommartid", "WIB": "västindonesisk tid", "GMT": "Greenwichtid", "CLT": "chilensk normaltid", "WART": "västargentinsk normaltid", "AST": "nordamerikansk atlantnormaltid", "ACWDT": "västcentralaustralisk sommartid", "EAT": "östafrikansk tid", "WITA": "centralindonesisk tid", "JST": "japansk normaltid", "CAT": "centralafrikansk tid", "UYST": "uruguayansk sommartid", "CST": "centralnordamerikansk normaltid", "CDT": "centralnordamerikansk sommartid", "GFT": "Franska Guyanatid"},
	}
}

// Locale returns the current translators string locale
func (sv *sv) Locale() string {
	return sv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sv'
func (sv *sv) PluralsCardinal() []locales.PluralRule {
	return sv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sv'
func (sv *sv) PluralsOrdinal() []locales.PluralRule {
	return sv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sv'
func (sv *sv) PluralsRange() []locales.PluralRule {
	return sv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sv'
func (sv *sv) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sv'
func (sv *sv) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)
	nMod10 := math.Mod(n, 10)

	if (nMod10 == 1 || nMod10 == 2) && (nMod100 != 11 && nMod100 != 12) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sv'
func (sv *sv) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sv *sv) MonthAbbreviated(month time.Month) string {
	return sv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sv *sv) MonthsAbbreviated() []string {
	return sv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sv *sv) MonthNarrow(month time.Month) string {
	return sv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sv *sv) MonthsNarrow() []string {
	return sv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sv *sv) MonthWide(month time.Month) string {
	return sv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sv *sv) MonthsWide() []string {
	return sv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sv *sv) WeekdayAbbreviated(weekday time.Weekday) string {
	return sv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sv *sv) WeekdaysAbbreviated() []string {
	return sv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sv *sv) WeekdayNarrow(weekday time.Weekday) string {
	return sv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sv *sv) WeekdaysNarrow() []string {
	return sv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sv *sv) WeekdayShort(weekday time.Weekday) string {
	return sv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sv *sv) WeekdaysShort() []string {
	return sv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sv *sv) WeekdayWide(weekday time.Weekday) string {
	return sv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sv *sv) WeekdaysWide() []string {
	return sv.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sv' and handles both Whole and Real numbers based on 'v'
func (sv *sv) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 6 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sv' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sv *sv) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 8
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sv.percentSuffix...)

	b = append(b, sv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sv'
func (sv *sv) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sv.currencies[currency]
	l := len(s) + len(symbol) + 8 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sv.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sv'
// in accounting notation.
func (sv *sv) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sv.currencies[currency]
	l := len(s) + len(symbol) + 8 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sv.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sv.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sv'
func (sv *sv) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sv'
func (sv *sv) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sv'
func (sv *sv) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sv'
func (sv *sv) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sv'
func (sv *sv) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sv'
func (sv *sv) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sv'
func (sv *sv) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sv'
func (sv *sv) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0x27, 0x6b, 0x6c, 0x27, 0x2e, 0x20}...)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
