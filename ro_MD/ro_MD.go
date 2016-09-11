package ro_MD

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ro_MD struct {
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

// New returns a new instance of translator for the 'ro_MD' locale
func New() locales.Translator {
	return &ro_MD{
		locale:                 "ro_MD",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "L", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "ian.", "feb.", "mar.", "apr.", "mai", "iun.", "iul.", "aug.", "sept.", "oct.", "nov.", "dec."},
		monthsNarrow:           []string{"", "I", "F", "M", "A", "M", "I", "I", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "ianuarie", "februarie", "martie", "aprilie", "mai", "iunie", "iulie", "august", "septembrie", "octombrie", "noiembrie", "decembrie"},
		daysAbbreviated:        []string{"Dum", "Lun", "Mar", "Mie", "Joi", "Vin", "Sâm"},
		daysNarrow:             []string{"D", "L", "Ma", "Mi", "J", "V", "S"},
		daysShort:              []string{"Du", "Lu", "Ma", "Mi", "Jo", "Vi", "Sâ"},
		daysWide:               []string{"duminică", "luni", "marți", "miercuri", "joi", "vineri", "sâmbătă"},
		periodsAbbreviated:     []string{"a.m.", "p.m."},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"BOT": "Ora Boliviei", "∅∅∅": "Ora de vară din Peru", "MST": "Ora standard în zona montană nord-americană", "CAT": "Ora Africii Centrale", "CST": "Ora standard centrală nord-americană", "CHAST": "Ora standard din Chatham", "CHADT": "Ora de vară din Chatham", "MYT": "Ora din Malaysia", "MDT": "Ora de vară în zona montană nord-americană", "CLST": "Ora de vară din Chile", "AEDT": "Ora de vară a Australiei Orientale", "EST": "Ora standard orientală nord-americană", "PDT": "Ora de vară în zona Pacific nord-americană", "WARST": "Ora de vară a Argentinei Occidentale", "COST": "Ora de vară a Columbiei", "AEST": "Ora standard a Australiei Orientale", "JST": "Ora standard a Japoniei", "AWST": "Ora standard a Australiei Occidentale", "HADT": "Ora de vară din Hawaii-Aleutine", "GFT": "Ora Guianei Franceze", "WIB": "Ora Indoneziei de Vest", "EDT": "Ora de vară orientală nord-americană", "UYT": "Ora standard a Uruguayului", "SRT": "Ora Surinamului", "ADT": "Ora de vară în zona Atlantic nord-americană", "HAT": "Ora de vară din Newfoundland", "PST": "Ora standard în zona Pacific nord-americană", "WESZ": "Ora de vară a Europei de Vest", "ACWST": "Ora standard a Australiei Central Occidentale", "GYT": "Ora din Guyana", "COT": "Ora standard a Columbiei", "WITA": "Ora Indoneziei Centrale", "BT": "Ora Bhutanului", "WART": "Ora standard a Argentinei Occidentale", "VET": "Ora Venezuelei", "NZST": "Ora standard a Noii Zeelande", "HKT": "Ora standard din Hong Kong", "TMT": "Ora standard din Turkmenistan", "SAST": "Ora Africii Meridionale", "AKST": "Ora standard din Alaska", "ECT": "Ora Ecuadorului", "MEZ": "Ora standard a Europei Centrale", "LHDT": "Ora de vară din Lord Howe", "WAST": "Ora de vară a Africii Occidentale", "JDT": "Ora de vară a Japoniei", "IST": "Ora Indiei", "WEZ": "Ora standard a Europei de Vest", "WIT": "Ora Indoneziei de Est", "HAST": "Ora standard din Hawaii-Aleutine", "EAT": "Ora Africii Orientale", "AWDT": "Ora de vară a Australiei Occidentale", "WAT": "Ora standard a Africii Occidentale", "MESZ": "Ora de vară a Europei Centrale", "OEZ": "Ora standard a Europei de Est", "ACST": "Ora standard a Australiei Centrale", "ACDT": "Ora de vară a Australiei Centrale", "ARST": "Ora de vară a Argentinei", "SGT": "Ora din Singapore", "CDT": "Ora de vară centrală nord-americană", "TMST": "Ora de vară din Turkmenistan", "OESZ": "Ora de vară a Europei de Est", "LHST": "Ora standard din Lord Howe", "NZDT": "Ora de vară a Noii Zeelande", "ART": "Ora standard a Argentinei", "GMT": "Ora de Greenwhich", "UYST": "Ora de vară a Uruguayului", "HNT": "Ora standard din Newfoundland", "HKST": "Ora de vară din Hong Kong", "AKDT": "Ora de vară din Alaska", "ACWDT": "Ora de vară a Australiei Central Occidentale", "ChST": "Ora din Chamorro", "AST": "Ora standard în zona Atlantic nord-americană", "CLT": "Ora standard din Chile"},
	}
}

// Locale returns the current translators string locale
func (ro *ro_MD) Locale() string {
	return ro.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsCardinal() []locales.PluralRule {
	return ro.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsOrdinal() []locales.PluralRule {
	return ro.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsRange() []locales.PluralRule {
	return ro.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	nMod100 := math.Mod(n, 100)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if (v != 0) || (n == 0) || (n != 1 && nMod100 >= 1 && nMod100 <= 19) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ro_MD'
func (ro *ro_MD) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ro.CardinalPluralRule(num1, v1)
	end := ro.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ro *ro_MD) MonthAbbreviated(month time.Month) string {
	return ro.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ro *ro_MD) MonthsAbbreviated() []string {
	return ro.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ro *ro_MD) MonthNarrow(month time.Month) string {
	return ro.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ro *ro_MD) MonthsNarrow() []string {
	return ro.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ro *ro_MD) MonthWide(month time.Month) string {
	return ro.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ro *ro_MD) MonthsWide() []string {
	return ro.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayAbbreviated(weekday time.Weekday) string {
	return ro.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ro *ro_MD) WeekdaysAbbreviated() []string {
	return ro.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayNarrow(weekday time.Weekday) string {
	return ro.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ro *ro_MD) WeekdaysNarrow() []string {
	return ro.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayShort(weekday time.Weekday) string {
	return ro.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ro *ro_MD) WeekdaysShort() []string {
	return ro.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayWide(weekday time.Weekday) string {
	return ro.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ro *ro_MD) WeekdaysWide() []string {
	return ro.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ro_MD' and handles both Whole and Real numbers based on 'v'
func (ro *ro_MD) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ro.decimal) + len(ro.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ro_MD' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ro *ro_MD) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ro.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ro.percentSuffix...)

	b = append(b, ro.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ro.currencies[currency]
	l := len(s) + len(ro.decimal) + len(ro.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ro.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ro.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ro_MD'
// in accounting notation.
func (ro *ro_MD) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ro.currencies[currency]
	l := len(s) + len(ro.decimal) + len(ro.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ro.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ro.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ro.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ro.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ro.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ro.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
