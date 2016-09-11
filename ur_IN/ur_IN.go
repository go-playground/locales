package ur_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ur_IN struct {
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
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'ur_IN' locale
func New() locales.Translator {
	return &ur_IN{
		locale:                 "ur_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "‎-‎",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysAbbreviated:        []string{"اتوار", "سوموار", "منگل", "بدھ", "جمعرات", "جمعہ", "ہفتہ"},
		daysNarrow:             []string{"S", "M", "T", "W", "T", "F", "S"},
		daysShort:              []string{"اتوار", "سوموار", "منگل", "بدھ", "جمعرات", "جمعہ", "ہفتہ"},
		daysWide:               []string{"اتوار", "سوموار", "منگل", "بدھ", "جمعرات", "جمعہ", "ہفتہ"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"قبل دوپہر", "بعد دوپہر"},
		erasAbbreviated:        []string{"قبل مسیح", "عیسوی"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"قبل مسیح", "عیسوی"},
		timezones:              map[string]string{"AWST": "آسٹریلیا ویسٹرن اسٹینڈرڈ ٹائم", "PST": "پیسفک اسٹینڈرڈ ٹائم", "MESZ": "وسطی یورپ کا موسم گرما کا وقت", "WITA": "وسطی انڈونیشیا ٹائم", "MDT": "MDT", "EDT": "ایسٹرن ڈے لائٹ ٹائم", "CHAST": "چیتھم اسٹینڈرڈ ٹائم", "COST": "کولمبیا سمر ٹائم", "ACWST": "آسٹریلین سنٹرل ویسٹرن اسٹینڈرڈ ٹائم", "SGT": "سنگاپور سٹینڈرڈ ٹائم", "CAT": "وسطی افریقہ ٹائم", "CHADT": "چیتھم ڈے لائٹ ٹائم", "ART": "ارجنٹینا سٹینڈرڈ ٹائم", "MYT": "ملیشیا ٹائم", "AEDT": "آسٹریلین ایسٹرن ڈے لائٹ ٹائم", "NZST": "نیوزی لینڈ سٹینڈرڈ ٹائم", "TMT": "ترکمانستان سٹینڈرڈ ٹائم", "ECT": "ایکواڈور ٹائم", "ChST": "چامورو سٹینڈرڈ ٹائم", "HKST": "ہانگ کانگ سمر ٹائم", "CDT": "سنٹرل ڈے لائٹ ٹائم", "ADT": "اٹلانٹک ڈے لائٹ ٹائم", "LHST": "لارڈ ہووے اسٹینڈرڈ ٹائم", "MST": "MST", "EST": "ایسٹرن اسٹینڈرڈ ٹائم", "GMT": "گرین وچ مین ٹائم", "WART": "مغربی ارجنٹینا سٹینڈرڈ ٹائم", "ACWDT": "آسٹریلین سنٹرل ویسٹرن ڈے لائٹ ٹائم", "UYT": "یوروگوئے سٹینڈرڈ ٹائم", "BOT": "بولیویا ٹائم", "COT": "کولمبیا سٹینڈرڈ ٹائم", "ACST": "آسٹریلین سنٹرل اسٹینڈرڈ ٹائم", "IST": "انڈیا سٹینڈرڈ ٹائم", "JST": "جاپان سٹینڈرڈ ٹائم", "CST": "سنٹرل اسٹینڈرڈ ٹائم", "AST": "اٹلانٹک اسٹینڈرڈ ٹائم", "MEZ": "وسطی یورپ کا معیاری وقت", "ACDT": "آسٹریلین سنٹرل ڈے لائٹ ٹائم", "ARST": "ارجنٹینا سمر ٹائم", "WARST": "مغربی ارجنٹینا سمر ٹائم", "VET": "وینزوئیلا ٹائم", "AWDT": "آسٹریلین ویسٹرن ڈے لائٹ ٹائم", "WIB": "مغربی انڈونیشیا ٹائم", "TMST": "ترکمانستان سمر ٹائم", "∅∅∅": "ایمیزون سمر ٹائم", "HNT": "نیو فاؤنڈ لینڈ اسٹینڈرڈ ٹائم", "CLST": "چلی سمر ٹائم", "OEZ": "مشرقی یورپ کا معیاری وقت", "GYT": "گیانا ٹائم", "HADT": "ہوائی الیوٹیئن ڈے لائٹ ٹائم", "NZDT": "نیوزی لینڈ ڈے لائٹ ٹائم", "PDT": "پیسفک ڈے لائٹ ٹائم", "CLT": "چلی سٹینڈرڈ ٹائم", "WAST": "مغربی افریقہ سمر ٹائم", "GFT": "فرینچ گیانا ٹائم", "AKST": "الاسکا اسٹینڈرڈ ٹائم", "AKDT": "الاسکا ڈے لائٹ ٹائم", "WEZ": "مغربی یورپ کا معیاری وقت", "WESZ": "مغربی یورپ کا موسم گرما کا وقت", "HAT": "نیو فاؤنڈ لینڈ ڈے لائٹ ٹائم", "SAST": "جنوبی افریقہ سٹینڈرڈ ٹائم", "WAT": "مغربی افریقہ سٹینڈرڈ ٹائم", "SRT": "سورینام ٹائم", "WIT": "مشرقی انڈونیشیا ٹائم", "HKT": "ہانگ کانگ سٹینڈرڈ ٹائم", "EAT": "مشرقی افریقہ ٹائم", "UYST": "یوروگوئے سمر ٹائم", "AEST": "آسٹریلین ایسٹرن اسٹینڈرڈ ٹائم", "BT": "بھوٹان ٹائم", "JDT": "جاپان ڈے لائٹ ٹائم", "LHDT": "لارڈ ہووے ڈے لائٹ ٹائم", "HAST": "ہوائی الیوٹیئن اسٹینڈرڈ ٹائم", "OESZ": "مشرقی یورپ کا موسم گرما کا وقت"},
	}
}

// Locale returns the current translators string locale
func (ur *ur_IN) Locale() string {
	return ur.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ur_IN'
func (ur *ur_IN) PluralsCardinal() []locales.PluralRule {
	return ur.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ur_IN'
func (ur *ur_IN) PluralsOrdinal() []locales.PluralRule {
	return ur.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ur_IN'
func (ur *ur_IN) PluralsRange() []locales.PluralRule {
	return ur.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ur_IN'
func (ur *ur_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ur_IN'
func (ur *ur_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ur_IN'
func (ur *ur_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ur *ur_IN) MonthAbbreviated(month time.Month) string {
	return ur.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ur *ur_IN) MonthsAbbreviated() []string {
	return ur.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ur *ur_IN) MonthNarrow(month time.Month) string {
	return ur.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ur *ur_IN) MonthsNarrow() []string {
	return ur.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ur *ur_IN) MonthWide(month time.Month) string {
	return ur.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ur *ur_IN) MonthsWide() []string {
	return ur.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ur *ur_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return ur.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ur *ur_IN) WeekdaysAbbreviated() []string {
	return ur.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ur *ur_IN) WeekdayNarrow(weekday time.Weekday) string {
	return ur.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ur *ur_IN) WeekdaysNarrow() []string {
	return ur.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ur *ur_IN) WeekdayShort(weekday time.Weekday) string {
	return ur.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ur *ur_IN) WeekdaysShort() []string {
	return ur.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ur *ur_IN) WeekdayWide(weekday time.Weekday) string {
	return ur.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ur *ur_IN) WeekdaysWide() []string {
	return ur.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ur_IN' and handles both Whole and Real numbers based on 'v'
func (ur *ur_IN) FmtNumber(num float64, v uint64) (results string) {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ur.decimal) + len(ur.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ur.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ur.minus) - 1; j >= 0; j-- {
			b = append(b, ur.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	results = string(b)
	return
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ur_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ur *ur_IN) FmtPercent(num float64, v uint64) (results string) {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ur.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ur.minus) - 1; j >= 0; j-- {
			b = append(b, ur.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ur.percent...)

	results = string(b)
	return
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ur_IN'
func (ur *ur_IN) FmtCurrency(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ur.currencies[currency]
	l := len(s) + len(ur.decimal) + len(ur.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, ur.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(ur.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ur.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(ur.minus) - 1; j >= 0; j-- {
			b = append(b, ur.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ur.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	results = string(b)
	return
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ur_IN'
// in accounting notation.
func (ur *ur_IN) FmtAccounting(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ur.currencies[currency]
	l := len(s) + len(ur.decimal) + len(ur.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ur.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, ur.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		for j := len(ur.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ur.currencyNegativePrefix[j])
		}

		for j := len(ur.minus) - 1; j >= 0; j-- {
			b = append(b, ur.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ur.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ur.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ur.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	results = string(b)
	return
}

// FmtDateShort returns the short date representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ur.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ur.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ur.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ur.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ur.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ur.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ur_IN'
func (ur *ur_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ur.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ur.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ur.periodsAbbreviated[0]...)
	} else {
		b = append(b, ur.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ur.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
