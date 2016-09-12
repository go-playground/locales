package uz_Cyrl

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type uz_Cyrl struct {
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

// New returns a new instance of translator for the 'uz_Cyrl' locale
func New() locales.Translator {
	return &uz_Cyrl{
		locale:                 "uz_Cyrl",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6, 2},
		decimal:                "٫",
		group:                  "٬",
		minus:                  "-",
		percent:                "٪",
		perMille:               "؉",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "A$", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "R$", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CA$", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CN¥", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "€", "FIM ", "FJD ", "FKP ", "FRF ", "£", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HK$", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "₪", "₹", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JP¥", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "₩", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MX$", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZ$", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "฿", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "NT$", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "US$", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "сўм", "VEB ", "VEF ", "₫", "VNN ", "VUV ", "WST ", "FCFA", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "EC$", "XDR ", "XEU ", "XFO ", "XFU ", "CFA", "XPD ", "CFPF", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Янв", "Фев", "Мар", "Апр", "Май", "Июн", "Июл", "Авг", "Сен", "Окт", "Ноя", "Дек"},
		monthsNarrow:           []string{"", "Я", "Ф", "М", "А", "М", "И", "И", "А", "С", "О", "Н", "Д"},
		monthsWide:             []string{"", "Январ", "Феврал", "Март", "Апрел", "Май", "Июн", "Июл", "Август", "Сентябр", "Октябр", "Ноябр", "Декабр"},
		daysAbbreviated:        []string{"Якш", "Душ", "Сеш", "Чор", "Пай", "Жум", "Шан"},
		daysNarrow:             []string{"Я", "Д", "С", "Ч", "П", "Ж", "Ш"},
		daysShort:              []string{"Якш", "Душ", "Сеш", "Чор", "Пай", "Жум", "Шан"},
		daysWide:               []string{"якшанба", "душанба", "сешанба", "чоршанба", "пайшанба", "жума", "шанба"},
		periodsAbbreviated:     []string{"TO", "TK"},
		periodsNarrow:          []string{"TO", "TK"},
		periodsWide:            []string{"TO", "TK"},
		erasAbbreviated:        []string{"М.А.", "Э"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"HKT": "Гонконг стандарт вақти", "ARST": "Аргентина ёзги вақти", "TMST": "Туркманистон ёзги вақти", "UYT": "Уругвай стандарт вақти", "BT": "Бутан вақти", "AWST": "Ғарбий Австралия стандарт вақти", "CLT": "Чили стандарт вақти", "EDT": "Шимолий Америка шарқий кундузги вақти", "WEZ": "Ғарбий Европа стандарт вақти", "COT": "Колумбия стандарт вақти", "LHST": "Лорд Хове стандарт вақти", "JDT": "Япония кундузги вақти", "AWDT": "Ғарбий Австралия кундузги вақти", "BOT": "Боливия вақти", "PST": "Шимолий Америка тинч океани стандарт вақти", "ACDT": "Марказий Австралия кундузги вақти", "WARST": "Ғарбий Аргентина ёзги вақти", "NZDT": "Янги Зеландия кундузги вақти", "TMT": "Туркманистон стандарт вақти", "OEZ": "Шарқий Европа стандарт вақти", "JST": "Япония стандарт вақти", "AKDT": "Аляска кундузги вақти", "ADT": "Атлантика кундузги вақти", "ART": "Аргентина стандарт вақти", "NZST": "Янги Зеландия стандарт вақти", "CAT": "Марказий Африка вақти", "WAT": "Ғарбий Африка стандарт вақти", "AST": "Атлантика стандарт вақти", "WESZ": "Ғарбий Европа ёзги вақти", "ChST": "Каморро вақти", "ACST": "Марказий Австралия стандарт вақти", "SRT": "Суринам вақти", "GFT": "Француз Гвианаси вақти", "ACWDT": "Марказий Австралия Ғарбий кундузги вақти", "WAST": "Ғарбий Африка ёзги вақти", "ACWST": "Марказий Австралия Ғарбий стандарт вақти", "EST": "Шимолий Америка шарқий стандарт вақти", "MEZ": "Марказий Европа стандарт вақти", "HADT": "Гавайи-алеут кундузги вақти", "SGT": "Сингапур вақти", "LHDT": "Лорд Хове кундузги вақти", "∅∅∅": "Азор ёзги вақти", "PDT": "Шимолий Америка тинч океани кундузги вақти", "MESZ": "Марказий Европа ёзги вақти", "OESZ": "Шарқий Европа ёзги вақти", "AEDT": "Шарқий Австралия кундузги вақти", "MYT": "Малайзия вақти", "CST": "Шимолий Америка марказий стандарт вақти", "ECT": "Эквадор вақти", "AKST": "Аляска стандарт вақти", "WART": "Ғарбий Аргентина стандарт вақти", "CHAST": "Чатхам стандарт вақти", "VET": "Венесуэла вақти", "HNT": "Ньюфаундленд стандарт вақти", "UYST": "Уругвай ёзги вақти", "MST": "MST", "HAST": "Гавайи-алеут стандарт вақти", "COST": "Колумбия ёзги вақти", "GMT": "Гринвич вақти", "AEST": "Шарқий Австралия стандарт вақти", "WIT": "Шарқий Индонезия вақти", "HKST": "Гонконг ёзги вақти", "SAST": "Жанубий Африка вақти", "EAT": "Шарқий Африка вақти", "CDT": "Шимолий Америка марказий кундузги вақти", "CHADT": "Чатхам кундузги вақти", "MDT": "MDT", "WITA": "Марказий Индонезия вақти", "IST": "Ҳиндистон вақти", "WIB": "Ғарбий Индонезия вақти", "HAT": "Ньюфаундленд кундузги вақти", "GYT": "Гайана вақти", "CLST": "Чили ёзги вақти"},
	}
}

// Locale returns the current translators string locale
func (uz *uz_Cyrl) Locale() string {
	return uz.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uz_Cyrl'
func (uz *uz_Cyrl) PluralsCardinal() []locales.PluralRule {
	return uz.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uz_Cyrl'
func (uz *uz_Cyrl) PluralsOrdinal() []locales.PluralRule {
	return uz.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uz_Cyrl'
func (uz *uz_Cyrl) PluralsRange() []locales.PluralRule {
	return uz.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Cyrl'
func (uz *uz_Cyrl) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Cyrl'
func (uz *uz_Cyrl) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uz_Cyrl'
func (uz *uz_Cyrl) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uz.CardinalPluralRule(num1, v1)
	end := uz.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uz *uz_Cyrl) MonthAbbreviated(month time.Month) string {
	return uz.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uz *uz_Cyrl) MonthsAbbreviated() []string {
	return uz.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uz *uz_Cyrl) MonthNarrow(month time.Month) string {
	return uz.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uz *uz_Cyrl) MonthsNarrow() []string {
	return uz.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uz *uz_Cyrl) MonthWide(month time.Month) string {
	return uz.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uz *uz_Cyrl) MonthsWide() []string {
	return uz.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayAbbreviated(weekday time.Weekday) string {
	return uz.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uz *uz_Cyrl) WeekdaysAbbreviated() []string {
	return uz.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayNarrow(weekday time.Weekday) string {
	return uz.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uz *uz_Cyrl) WeekdaysNarrow() []string {
	return uz.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayShort(weekday time.Weekday) string {
	return uz.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uz *uz_Cyrl) WeekdaysShort() []string {
	return uz.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uz *uz_Cyrl) WeekdayWide(weekday time.Weekday) string {
	return uz.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uz *uz_Cyrl) WeekdaysWide() []string {
	return uz.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uz_Cyrl' and handles both Whole and Real numbers based on 'v'
func (uz *uz_Cyrl) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uz_Cyrl' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uz *uz_Cyrl) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uz.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 5 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
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

	for j := len(uz.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, uz.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Cyrl'
// in accounting notation.
func (uz *uz_Cyrl) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 5 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
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

		for j := len(uz.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, uz.currencyNegativePrefix[j])
		}

		b = append(b, uz.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(uz.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, uz.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateShort(t time.Time) string {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, uz.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'uz_Cyrl'
func (uz *uz_Cyrl) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := uz.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
