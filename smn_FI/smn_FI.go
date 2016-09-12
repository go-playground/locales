package smn_FI

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type smn_FI struct {
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

// New returns a new instance of translator for the 'smn_FI' locale
func New() locales.Translator {
	return &smn_FI{
		locale:                 "smn_FI",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "K",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "K",
		monthsWide:             []string{"", "uđđâivemáánu", "kuovâmáánu", "njuhčâmáánu", "cuáŋuimáánu", "vyesimáánu", "kesimáánu", "syeinimáánu", "porgemáánu", "čohčâmáánu", "roovvâdmáánu", "skammâmáánu", "juovlâmáánu"},
		daysAbbreviated:        []string{"pa", "vu", "ma", "ko", "tu", "vá", "lá"},
		daysNarrow:             []string{"P", "V", "M", "K", "T", "V", "L"},
		daysWide:               []string{"pasepeeivi", "vuossaargâ", "majebaargâ", "koskoho", "tuorâstuv", "vástuppeeivi", "lávurduv"},
		timezones:              map[string]string{"PST": "PST", "NZDT": "NZDT", "WESZ": "WESZ", "HNT": "HNT", "HKST": "HKST", "AKDT": "AKDT", "WEZ": "WEZ", "UYST": "UYST", "ACWDT": "ACWDT", "AKST": "AKST", "PDT": "PDT", "WAST": "WAST", "TMST": "TMST", "LHDT": "LHDT", "EAT": "EAT", "CST": "CST", "CDT": "CDT", "ChST": "ChST", "OESZ": "OESZ", "CLST": "CLST", "WART": "WART", "SRT": "SRT", "BOT": "BOT", "WIB": "WIB", "CAT": "CAT", "GMT": "GMT", "MDT": "MDT", "JST": "JST", "COT": "COT", "MYT": "MYT", "JDT": "JDT", "HADT": "HADT", "LHST": "LHST", "MST": "MST", "HKT": "HKT", "ECT": "ECT", "NZST": "NZST", "SGT": "SGT", "TMT": "TMT", "UYT": "UYT", "CHADT": "CHADT", "AST": "AST", "WAT": "WAT", "OEZ": "OEZ", "ACDT": "ACDT", "IST": "IST", "COST": "COST", "WARST": "WARST", "ADT": "ADT", "EDT": "EDT", "GYT": "GYT", "GFT": "GFT", "AWST": "AWST", "BT": "BT", "ARST": "ARST", "∅∅∅": "∅∅∅", "HAT": "HAT", "AEDT": "AEDT", "CHAST": "CHAST", "VET": "VET", "AWDT": "AWDT", "WITA": "WITA", "EST": "EST", "HAST": "HAST", "CLT": "CLT", "WIT": "WIT", "ART": "ART", "MESZ": "MESZ", "SAST": "SAST", "ACST": "ACST", "ACWST": "ACWST", "MEZ": "MEZ", "AEST": "AEST"},
	}
}

// Locale returns the current translators string locale
func (smn *smn_FI) Locale() string {
	return smn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'smn_FI'
func (smn *smn_FI) PluralsCardinal() []locales.PluralRule {
	return smn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'smn_FI'
func (smn *smn_FI) PluralsOrdinal() []locales.PluralRule {
	return smn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'smn_FI'
func (smn *smn_FI) PluralsRange() []locales.PluralRule {
	return smn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'smn_FI'
func (smn *smn_FI) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'smn_FI'
func (smn *smn_FI) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'smn_FI'
func (smn *smn_FI) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (smn *smn_FI) MonthAbbreviated(month time.Month) string {
	return smn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (smn *smn_FI) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (smn *smn_FI) MonthNarrow(month time.Month) string {
	return smn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (smn *smn_FI) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (smn *smn_FI) MonthWide(month time.Month) string {
	return smn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (smn *smn_FI) MonthsWide() []string {
	return smn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (smn *smn_FI) WeekdayAbbreviated(weekday time.Weekday) string {
	return smn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (smn *smn_FI) WeekdaysAbbreviated() []string {
	return smn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (smn *smn_FI) WeekdayNarrow(weekday time.Weekday) string {
	return smn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (smn *smn_FI) WeekdaysNarrow() []string {
	return smn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (smn *smn_FI) WeekdayShort(weekday time.Weekday) string {
	return smn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (smn *smn_FI) WeekdaysShort() []string {
	return smn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (smn *smn_FI) WeekdayWide(weekday time.Weekday) string {
	return smn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (smn *smn_FI) WeekdaysWide() []string {
	return smn.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'smn_FI' and handles both Whole and Real numbers based on 'v'
func (smn *smn_FI) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'smn_FI' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (smn *smn_FI) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'smn_FI'
func (smn *smn_FI) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := smn.currencies[currency]
	l := len(s) + len(symbol) + 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, smn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(smn.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, smn.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, smn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, smn.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'smn_FI'
// in accounting notation.
func (smn *smn_FI) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := smn.currencies[currency]
	l := len(s) + len(symbol) + 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, smn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(smn.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, smn.currencyNegativePrefix[j])
		}

		b = append(b, smn.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(smn.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, smn.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, smn.currencyNegativeSuffix...)
	} else {

		b = append(b, smn.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'smn_FI'
func (smn *smn_FI) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}
