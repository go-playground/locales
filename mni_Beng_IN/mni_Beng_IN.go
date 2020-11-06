package mni_Beng_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type mni_Beng_IN struct {
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

// New returns a new instance of translator for the 'mni_Beng_IN' locale
func New() locales.Translator {
	return &mni_Beng_IN{
		locale:             "mni_Beng_IN",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "জানুৱারি", "ফেব্রুৱারি", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "ওক্টোবর", "নভেম্বর", "ডিসেম্বর"},
		monthsNarrow:       []string{"", "জা", "ফে", "মার", "এপ", "মে", "জুন", "জুল", "আ", "সে", "ওক", "নব", "ডি"},
		monthsWide:         []string{"", "জানুৱারি", "ফেব্রুৱারি", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "ওক্টোবর", "নভেম্বর", "ডিসেম্বর"},
		daysAbbreviated:    []string{"নোংমাইজিং", "নিংথৌকাবা", "লৈবাকপোকপা", "য়ুমশকৈশা", "শগোলশেন", "ইরাই", "থাংজ"},
		daysNarrow:         []string{"নোং", "নিং", "লৈ", "য়ুম", "শগ", "ইরা", "থাং"},
		daysWide:           []string{"নোংমাইজিং", "নিংথৌকাবা", "লৈবাকপোকপা", "য়ুমশকৈশা", "শগোলশেন", "ইরাই", "থাংজ"},
		periodsAbbreviated: []string{"নুমাং", "PM"},
		periodsWide:        []string{"এ এম", "পি এম"},
		erasAbbreviated:    []string{"খৃ: মমাং", "খৃ: মতুং"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"", ""},
		timezones:          map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "অটলান্টিক দেলাইট টাইম", "AEDT": "AEDT", "AEST": "AEST", "AKDT": "AKDT", "AKST": "AKST", "ARST": "ARST", "ART": "ART", "AST": "অটলান্টিক ষ্টেন্দর্দ টাইম", "AWDT": "AWDT", "AWST": "AWST", "BOT": "BOT", "BT": "BT", "CAT": "CAT", "CDT": "নোর্থ অমেরিকান সেন্ত্রেল দেলাইট টাইম", "CHADT": "CHADT", "CHAST": "CHAST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "নোর্থ অমেরিকান সেন্ত্রেল ষ্টেন্দর্দ টাইম", "ChST": "ChST", "EAT": "EAT", "ECT": "ECT", "EDT": "নোর্থ অমেরিকান ইষ্টর্ন দেলাইট টাইম", "EST": "নোর্থ অমেরিকান ইষ্টর্ন ষ্টেন্দর্দ টাইম", "GFT": "GFT", "GMT": "গ্রিনৱিচ মিন টাইম", "GST": "GST", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HAT": "HAT", "HECU": "HECU", "HEEG": "HEEG", "HENOMX": "HENOMX", "HEOG": "HEOG", "HEPM": "HEPM", "HEPMX": "HEPMX", "HKST": "HKST", "HKT": "HKT", "HNCU": "HNCU", "HNEG": "HNEG", "HNNOMX": "HNNOMX", "HNOG": "HNOG", "HNPM": "HNPM", "HNPMX": "HNPMX", "HNT": "HNT", "IST": "IST", "JDT": "JDT", "JST": "JST", "LHDT": "LHDT", "LHST": "LHST", "MDT": "নোর্থ অমেরিকান মাউন্টেন দেলাইট টাইম", "MESZ": "সেন্ত্রেল য়ুরোপিয়ান সমর টাইম", "MEZ": "সেন্ত্রেল য়ুরোপিয়ান ষ্টেন্দর্দ টাইম", "MST": "নোর্থ অমেরিকান মাউন্টেন ষ্টেন্দর্দ টাইম", "MYT": "MYT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "ইষ্টর্ন য়ুরোপিয়ান সমর টাইম", "OEZ": "ইষ্টর্ন য়ুরোপিয়ান ষ্টেন্দর্দ টাইম", "PDT": "নোর্থ অমেরিকান পেসিফিক দেলাইট টাইম", "PST": "নোর্থ অমেরিকান পেসিফিক ষ্টেন্দর্দ টাইম", "SAST": "SAST", "SGT": "SGT", "SRT": "SRT", "TMST": "TMST", "TMT": "TMT", "UYST": "UYST", "UYT": "UYT", "VET": "VET", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "ৱেষ্টর্ন য়ুরোপিয়ান সমর টাইম", "WEZ": "ৱেষ্টর্ন য়ুরোপিয়ান ষ্টেন্দর্দ টাইম", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (mni *mni_Beng_IN) Locale() string {
	return mni.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mni_Beng_IN'
func (mni *mni_Beng_IN) PluralsCardinal() []locales.PluralRule {
	return mni.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mni_Beng_IN'
func (mni *mni_Beng_IN) PluralsOrdinal() []locales.PluralRule {
	return mni.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mni_Beng_IN'
func (mni *mni_Beng_IN) PluralsRange() []locales.PluralRule {
	return mni.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mni *mni_Beng_IN) MonthAbbreviated(month time.Month) string {
	return mni.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mni *mni_Beng_IN) MonthsAbbreviated() []string {
	return mni.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mni *mni_Beng_IN) MonthNarrow(month time.Month) string {
	return mni.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mni *mni_Beng_IN) MonthsNarrow() []string {
	return mni.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mni *mni_Beng_IN) MonthWide(month time.Month) string {
	return mni.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mni *mni_Beng_IN) MonthsWide() []string {
	return mni.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mni *mni_Beng_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return mni.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mni *mni_Beng_IN) WeekdaysAbbreviated() []string {
	return mni.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mni *mni_Beng_IN) WeekdayNarrow(weekday time.Weekday) string {
	return mni.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mni *mni_Beng_IN) WeekdaysNarrow() []string {
	return mni.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mni *mni_Beng_IN) WeekdayShort(weekday time.Weekday) string {
	return mni.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mni *mni_Beng_IN) WeekdaysShort() []string {
	return mni.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mni *mni_Beng_IN) WeekdayWide(weekday time.Weekday) string {
	return mni.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mni *mni_Beng_IN) WeekdaysWide() []string {
	return mni.daysWide
}

// Decimal returns the decimal point of number
func (mni *mni_Beng_IN) Decimal() string {
	return mni.decimal
}

// Group returns the group of number
func (mni *mni_Beng_IN) Group() string {
	return mni.group
}

// Group returns the minus sign of number
func (mni *mni_Beng_IN) Minus() string {
	return mni.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mni_Beng_IN' and handles both Whole and Real numbers based on 'v'
func (mni *mni_Beng_IN) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mni_Beng_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mni *mni_Beng_IN) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mni.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mni_Beng_IN'
// in accounting notation.
func (mni *mni_Beng_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mni.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mni.monthsAbbreviated[t.Month()]...)
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

// FmtDateLong returns the long date representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mni.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mni.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, mni.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mni.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mni.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mni.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mni.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mni.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mni_Beng_IN'
func (mni *mni_Beng_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mni.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mni.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mni.periodsAbbreviated[0]...)
	} else {
		b = append(b, mni.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mni.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
