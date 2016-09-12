package sah_RU

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sah_RU struct {
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

// New returns a new instance of translator for the 'sah_RU' locale
func New() locales.Translator {
	return &sah_RU{
		locale:                 "sah_RU",
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
		monthsAbbreviated:      []string{"", "Тохс", "Олун", "Клн_ттр", "Мус_уст", "Ыам_йн", "Бэс_йн", "От_йн", "Атрдь_йн", "Блҕн_йн", "Алт", "Сэт", "Ахс"},
		monthsNarrow:           []string{"", "Т", "О", "К", "М", "Ы", "Б", "О", "А", "Б", "А", "С", "А"},
		monthsWide:             []string{"", "Тохсунньу", "Олунньу", "Кулун тутар", "Муус устар", "Ыам ыйын", "Бэс ыйын", "От ыйын", "Атырдьых ыйын", "Балаҕан ыйын", "Алтынньы", "Сэтинньи", "Ахсынньы"},
		daysAbbreviated:        []string{"Бс", "Бн", "Оп", "Сэ", "Чп", "Бэ", "Сб"},
		daysNarrow:             []string{"Б", "Б", "О", "С", "Ч", "Б", "С"},
		daysWide:               []string{"Баскыһыанньа", "Бэнидиэлинньик", "Оптуорунньук", "Сэрэдэ", "Чэппиэр", "Бээтиҥсэ", "Субуота"},
		periodsAbbreviated:     []string{"ЭИ", "ЭК"},
		periodsWide:            []string{"ЭИ", "ЭК"},
		erasAbbreviated:        []string{"б. э. и.", "б. э"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"GYT": "GYT", "CLST": "CLST", "ACDT": "ACDT", "ECT": "ECT", "IST": "IST", "EST": "EST", "CAT": "CAT", "WEZ": "WEZ", "EDT": "EDT", "WAT": "WAT", "AEDT": "AEDT", "MEZ": "MEZ", "GFT": "GFT", "JST": "JST", "HAST": "HAST", "SAST": "SAST", "PDT": "PDT", "BT": "BT", "MST": "MST", "HKT": "HKT", "JDT": "JDT", "ACWST": "ACWST", "MESZ": "MESZ", "LHST": "LHST", "∅∅∅": "∅∅∅", "VET": "VET", "ARST": "ARST", "AWST": "AWST", "AWDT": "AWDT", "WAST": "WAST", "MYT": "MYT", "ACWDT": "ACWDT", "UYST": "UYST", "HKST": "HKST", "CHADT": "CHADT", "ART": "ART", "AKST": "AKST", "PST": "PST", "ACST": "ACST", "WART": "WART", "AST": "AST", "NZDT": "NZDT", "UYT": "UYT", "AKDT": "AKDT", "AEST": "AEST", "CLT": "CLT", "LHDT": "LHDT", "CST": "CST", "ADT": "ADT", "HADT": "HADT", "COST": "COST", "WESZ": "WESZ", "NZST": "NZST", "WIT": "WIT", "WARST": "WARST", "WITA": "WITA", "SGT": "SGT", "HAT": "HAT", "CHAST": "CHAST", "TMST": "TMST", "OESZ": "OESZ", "GMT": "GMT", "SRT": "SRT", "HNT": "HNT", "OEZ": "OEZ", "ChST": "ChST", "EAT": "EAT", "BOT": "BOT", "COT": "COT", "WIB": "WIB", "CDT": "CDT", "TMT": "TMT", "MDT": "MDT"},
	}
}

// Locale returns the current translators string locale
func (sah *sah_RU) Locale() string {
	return sah.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sah_RU'
func (sah *sah_RU) PluralsCardinal() []locales.PluralRule {
	return sah.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sah_RU'
func (sah *sah_RU) PluralsOrdinal() []locales.PluralRule {
	return sah.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sah_RU'
func (sah *sah_RU) PluralsRange() []locales.PluralRule {
	return sah.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sah_RU'
func (sah *sah_RU) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sah_RU'
func (sah *sah_RU) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sah_RU'
func (sah *sah_RU) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sah *sah_RU) MonthAbbreviated(month time.Month) string {
	return sah.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sah *sah_RU) MonthsAbbreviated() []string {
	return sah.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sah *sah_RU) MonthNarrow(month time.Month) string {
	return sah.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sah *sah_RU) MonthsNarrow() []string {
	return sah.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sah *sah_RU) MonthWide(month time.Month) string {
	return sah.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sah *sah_RU) MonthsWide() []string {
	return sah.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sah *sah_RU) WeekdayAbbreviated(weekday time.Weekday) string {
	return sah.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sah *sah_RU) WeekdaysAbbreviated() []string {
	return sah.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sah *sah_RU) WeekdayNarrow(weekday time.Weekday) string {
	return sah.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sah *sah_RU) WeekdaysNarrow() []string {
	return sah.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sah *sah_RU) WeekdayShort(weekday time.Weekday) string {
	return sah.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sah *sah_RU) WeekdaysShort() []string {
	return sah.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sah *sah_RU) WeekdayWide(weekday time.Weekday) string {
	return sah.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sah *sah_RU) WeekdaysWide() []string {
	return sah.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sah_RU' and handles both Whole and Real numbers based on 'v'
func (sah *sah_RU) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sah_RU' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sah *sah_RU) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sah_RU'
func (sah *sah_RU) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sah.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sah.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(sah.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, sah.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, sah.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sah.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sah_RU'
// in accounting notation.
func (sah *sah_RU) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sah.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sah.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(sah.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sah.currencyNegativePrefix[j])
		}

		b = append(b, sah.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(sah.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, sah.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, sah.currencyNegativeSuffix...)
	} else {

		b = append(b, sah.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, sah.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, sah.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0xd1, 0x81, 0xd1, 0x8b, 0xd0, 0xbb, 0x27, 0x20}...)
	b = append(b, sah.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0xd0, 0xba, 0xd2, 0xaf, 0xd0, 0xbd, 0xd1, 0x8d, 0x27, 0x2c, 0x20}...)
	b = append(b, sah.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sah.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sah.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sah.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sah.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sah_RU'
func (sah *sah_RU) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}
