package pcm_NG

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type pcm_NG struct {
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

// New returns a new instance of translator for the 'pcm_NG' locale
func New() locales.Translator {
	return &pcm_NG{
		locale:             "pcm_NG",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       []locales.PluralRule{6},
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		perMille:           "‰",
		timeSeparator:      ":",
		inifinity:          "∞",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "Jén", "Fẹ́b", "Mach", "Épr", "Mee", "Jun", "Jul", "Ọgọ", "Sẹp", "Ọkt", "Nọv", "Dis"},
		monthsNarrow:       []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:         []string{"", "Jénúári", "Fẹ́búári", "Mach", "Éprel", "Mee", "Jun", "Julai", "Ọgọst", "Sẹptẹ́mba", "Ọktóba", "Nọvẹ́mba", "Disẹ́mba"},
		daysAbbreviated:    []string{"Sọ́n", "Mọ́n", "Tiú", "Wẹ́n", "Tọ́z", "Fraí", "Sát"},
		daysNarrow:         []string{"S", "M", "T", "W", "T", "F", "S"},
		daysShort:          []string{"Sọ́n", "Mọ́n", "Tiú", "Wẹ́n", "Tọ́z", "Fraí", "Sát"},
		daysWide:           []string{"Sọ́ndè", "Mọ́ndè", "Tiúzdè", "Wẹ́nẹ́zdè", "Tọ́zdè", "Fraídè", "Sátọdè"},
		periodsAbbreviated: []string{"AM", "PM"},
		periodsNarrow:      []string{"AM", "PM"},
		periodsWide:        []string{"Fọ mọ́nin", "Fọ ívnin"},
		erasAbbreviated:    []string{"BK", "KIY"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"Bifọ́ Kraist", "Kraist Im Yiẹ"},
		timezones:          map[string]string{"ACDT": "Ọstrélia Mídúl Délaít Taim", "ACST": "Ọstrélia Mídúl Fíksd Taim", "ACWDT": "Ọstrélia Mídúl Wẹ́stán Délaít Taim", "ACWST": "Ọstrélia Mídúl Wẹ́stán Fíksd Taim", "ADT": "Atlántík Délaít Taim", "AEDT": "Ọstrélia Ístán Délaít Taim", "AEST": "Ọstrélia Ístán Fíksd Taim", "AKDT": "Aláská Délaít Taim", "AKST": "Aláská Fíksd Taim", "ARST": "Ajẹntína Họ́t Sízín Taim", "ART": "Ajẹntína Fíksd Taim", "AST": "Atlántík Fíksd Taim", "AWDT": "Ọstrélia Wẹ́stán Délaít Taim", "AWST": "Ọstrélia Wẹ́stán Fíksd Taim", "BOT": "Bolívia Fíksd Taim", "BT": "Butan Taim", "CAT": "Mídúl Áfríká Taim", "CDT": "Nọ́t Amẹ́ríká Mídúl Ériá Délaít Taim", "CHADT": "Chátam Délaít Taim", "CHAST": "Chátam Fíksd Taim", "CLST": "Chílẹ Họ́t Sízín Taim", "CLT": "Chílẹ Fíksd Taim", "COST": "Kolómbia Họ́t Sízín Taim", "COT": "Kolómbia Fíksd Taim", "CST": "Nọ́t Amẹ́ríká Mídúl Ériá Fíksd Taim", "ChST": "Chamóro Fíksd Taim", "EAT": "Íst Áfríká Taim", "ECT": "Ẹ́kwuádọ Taim", "EDT": "Nọ́t Amẹ́ríká Ístán Ériá Délaít Taim", "EST": "Nọ́t Amẹ́ríká Ístán Ériá Fíksd Taim", "GFT": "Frẹ́nch Giána Taim", "GMT": "Grínwích Mín Taim", "GST": "Gọ́lf Fíksd Taim", "GYT": "Gayána Taim", "HADT": "Hawaií-Elúshián Délaít Taim", "HAST": "Hawaií-Elúshián Fíksd Taim", "HAT": "Niúfaúndlánd Délaít Taim", "HECU": "Kúba Délaít Taim", "HEEG": "Íist Grínlánd Họ́t Sízin Taim", "HENOMX": "Nọ́twẹ́st Mẹ́ksíko Délaít Taim", "HEOG": "Wẹ́st Grínlánd Họ́t Sízin Taim", "HEPM": "Sent Piẹr an Míkẹlọn Délaít Taim", "HEPMX": "Mẹ́ksíkó Pasífík Délaít Taim", "HKST": "Họng Kọng Họ́t Sízin Taim", "HKT": "Họng Kọng Fíksd Taim", "HNCU": "Kúba Fíksd Taim", "HNEG": "Íist Grínlánd Fíksd Taim", "HNNOMX": "Nọ́twẹ́st Mẹ́ksíko Fíksd Taim", "HNOG": "Wẹ́st Grínlánd Fíksd Taim", "HNPM": "Sent Piẹr an Míkẹlọn Fíksd Taim", "HNPMX": "Mẹ́ksíkó Pasífík Fíksd Taim", "HNT": "Niúfaúndlánd Fíksd Taim", "IST": "Índia Fíksd Taim", "JDT": "Japan Délaít Taim", "JST": "Japan Fíksd Taim", "LHDT": "Lọd Haú Délaít Taim", "LHST": "Lọd Haú Fíksd Taim", "MDT": "MDT", "MESZ": "Mídúl Yúrop Họ́t Sízin Taim", "MEZ": "Mídúl Yúrop Fíksd Taim", "MST": "MST", "MYT": "Maléshia Taim", "NZDT": "Niú Ziland Délaít Taim", "NZST": "Niú Ziland Fíksd Taim", "OESZ": "Ístán Yúrop Họ́t Sízin Taim", "OEZ": "Ístán Yúrop Fíksd Taim", "PDT": "Nọ́t Amẹ́ríká Pasífík Ériá Délaít Taim", "PST": "Nọ́t Amẹ́ríká Pasífík Ériá Fíksd Taim", "SAST": "Saút Áfríká Fíksd Taim", "SGT": "Singapọ Taim", "SRT": "Súrínam Taim", "TMST": "Tọkmẹnistan Họ́t Sízin Taim", "TMT": "Tọkmẹnistan Fíksd Taim", "UYST": "Yúrugwue Họ́t Sízín Taim", "UYT": "Yúrugwue Fíksd Taim", "VET": "Vẹnẹzuẹ́la Taim", "WARST": "Wẹ́stán Ajẹntína Họ́t Sízín Taim", "WART": "Wẹ́stán Ajẹntína Fíksd Taim", "WAST": "Wẹ́st Áfríká Họ́t Sízin Taim", "WAT": "Wẹ́st Áfríká Fíksd Taim", "WESZ": "Wẹ́stán Yúrop Họ́t Sízin Taim", "WEZ": "Wẹ́stán Yúrop Fíksd Taim", "WIB": "Wẹ́stán Indonẹ́shia Taim", "WIT": "Ístán Indonẹ́shia Taim", "WITA": "Mídúl Indonẹ́shia Taim", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (pcm *pcm_NG) Locale() string {
	return pcm.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pcm_NG'
func (pcm *pcm_NG) PluralsCardinal() []locales.PluralRule {
	return pcm.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pcm_NG'
func (pcm *pcm_NG) PluralsOrdinal() []locales.PluralRule {
	return pcm.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pcm_NG'
func (pcm *pcm_NG) PluralsRange() []locales.PluralRule {
	return pcm.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pcm_NG'
func (pcm *pcm_NG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pcm_NG'
func (pcm *pcm_NG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pcm_NG'
func (pcm *pcm_NG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pcm *pcm_NG) MonthAbbreviated(month time.Month) string {
	return pcm.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pcm *pcm_NG) MonthsAbbreviated() []string {
	return pcm.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pcm *pcm_NG) MonthNarrow(month time.Month) string {
	return pcm.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pcm *pcm_NG) MonthsNarrow() []string {
	return pcm.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pcm *pcm_NG) MonthWide(month time.Month) string {
	return pcm.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pcm *pcm_NG) MonthsWide() []string {
	return pcm.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pcm *pcm_NG) WeekdayAbbreviated(weekday time.Weekday) string {
	return pcm.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pcm *pcm_NG) WeekdaysAbbreviated() []string {
	return pcm.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pcm *pcm_NG) WeekdayNarrow(weekday time.Weekday) string {
	return pcm.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pcm *pcm_NG) WeekdaysNarrow() []string {
	return pcm.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pcm *pcm_NG) WeekdayShort(weekday time.Weekday) string {
	return pcm.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pcm *pcm_NG) WeekdaysShort() []string {
	return pcm.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pcm *pcm_NG) WeekdayWide(weekday time.Weekday) string {
	return pcm.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pcm *pcm_NG) WeekdaysWide() []string {
	return pcm.daysWide
}

// Decimal returns the decimal point of number
func (pcm *pcm_NG) Decimal() string {
	return pcm.decimal
}

// Group returns the group of number
func (pcm *pcm_NG) Group() string {
	return pcm.group
}

// Group returns the minus sign of number
func (pcm *pcm_NG) Minus() string {
	return pcm.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pcm_NG' and handles both Whole and Real numbers based on 'v'
func (pcm *pcm_NG) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pcm.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pcm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pcm_NG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pcm *pcm_NG) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pcm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pcm.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pcm_NG'
func (pcm *pcm_NG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pcm.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pcm.group[0])
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
		b = append(b, pcm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pcm.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pcm_NG'
// in accounting notation.
func (pcm *pcm_NG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pcm.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pcm.group[0])
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

		b = append(b, pcm.minus[0])

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
			b = append(b, pcm.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pcm.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pcm.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, pcm.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pcm.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pcm_NG'
func (pcm *pcm_NG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pcm.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
