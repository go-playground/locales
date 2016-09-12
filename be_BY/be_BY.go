package be_BY

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type be_BY struct {
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

// New returns a new instance of translator for the 'be_BY' locale
func New() locales.Translator {
	return &be_BY{
		locale:                 "be_BY",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{4, 6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ".",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		percentSuffix:          " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "сту", "лют", "сак", "кра", "мая", "чэр", "ліп", "жні", "вер", "кас", "ліс", "сне"},
		monthsNarrow:           []string{"", "с", "л", "с", "к", "м", "ч", "л", "ж", "в", "к", "л", "с"},
		monthsWide:             []string{"", "студзеня", "лютага", "сакавіка", "красавіка", "мая", "чэрвеня", "ліпеня", "жніўня", "верасня", "кастрычніка", "лістапада", "снежня"},
		daysAbbreviated:        []string{"нд", "пн", "аў", "ср", "чц", "пт", "сб"},
		daysNarrow:             []string{"н", "п", "а", "с", "ч", "п", "с"},
		daysShort:              []string{"нд", "пн", "аў", "ср", "чц", "пт", "сб"},
		daysWide:               []string{"нядзеля", "панядзелак", "аўторак", "серада", "чацвер", "пятніца", "субота"},
		periodsAbbreviated:     []string{"раніцы", "вечара"},
		periodsNarrow:          []string{"ран.", "веч."},
		periodsWide:            []string{"да паўдня", "пасля паўдня"},
		erasAbbreviated:        []string{"да н.э.", "н.э."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"HADT": "Гавайска-Алеуцкі летні час", "MDT": "Паўночнаамерыканскі горны летні час", "SRT": "Час Сурынама", "AKST": "Стандартны час Аляскі", "ADT": "Атлантычны летні час", "ACWDT": "Заходні летні час Цэнтральнай Аўстраліі", "PDT": "Ціхаакіянскі летні час", "WIT": "Усходнеінданезійскі час", "CDT": "Паўночнаамерыканскі цэнтральны летні час", "HKT": "Стандартны час Ганконга", "WITA": "Цэнтральнаінданезійскі час", "WIB": "Заходнеінданезійскі час", "UYT": "Уругвайскі стандартны час", "SAST": "Паўднёваафрыканскі час", "HNT": "Ньюфаўндлендскі стандартны час", "ChST": "Час Чамора", "AEST": "Стандартны час усходняй Аўстраліі", "ACDT": "Летні час цэнтральнай Аўстраліі", "AWDT": "Летні час заходняй Аўстраліі", "PST": "Ціхаакіянскі стандартны час", "HAST": "Гавайска-Алеуцкі стандартны час", "ART": "Аргенцінскі стандартны час", "SGT": "Сінгапурскі час", "CLT": "Чылійскі стандартны час", "CAT": "Цэнтральнаафрыканскі час", "WAT": "Заходнеафрыканскі стандартны час", "NZDT": "Летні час Новай Зеландыі", "IST": "Час Індыі", "HAT": "Ньюфаўндлендскі летні час", "AEDT": "Летні час усходняй Аўстраліі", "CLST": "Чылійскі летні час", "WART": "Стандартны час Заходняй Аргенціны", "ECT": "Эквадорскі час", "JST": "Стандартны час Японіі", "NZST": "Стандартны час Новай Зеландыі", "OEZ": "Усходнееўрапейскі стандартны час", "WAST": "Заходнеафрыканскі летні час", "OESZ": "Усходнееўрапейскі летні час", "UYST": "Уругвайскі летні час", "JDT": "Летні час Японіі", "AKDT": "Летні час Аляскі", "TMT": "Стандартны час Туркменістана", "CHAST": "Стандартны час Чатэма", "CHADT": "Летні час Чатэма", "∅∅∅": "Амазонскі летні час", "COST": "Калумбійскі летні час", "LHDT": "Летні час Лорд-Хау", "GYT": "Час Гаяны", "GFT": "Час Французскай Гвіяны", "ARST": "Аргенцінскі летні час", "ACWST": "Заходні стандартны час Цэнтральнай Аўстраліі", "EDT": "Паўночнаамерыканскі ўсходні летні час", "ACST": "Стандартны час цэнтральнай Аўстраліі", "VET": "Венесуэльскі час", "AST": "Атлантычны стандартны час", "COT": "Калумбійскі стандартны час", "LHST": "Стандартны час Лорд-Хау", "GMT": "Час па Грынвічы", "WARST": "Летні час Заходняй Аргенціны", "HKST": "Летні час Ганконга", "MYT": "Час Малайзіі", "BT": "Час Бутана", "CST": "Паўночнаамерыканскі цэнтральны стандартны час", "BOT": "Балівійскі час", "WESZ": "Заходнееўрапейскі летні час", "EAT": "Усходнеафрыканскі час", "AWST": "Стандартны час заходняй Аўстраліі", "MEZ": "Цэнтральнаеўрапейскі стандартны час", "MESZ": "Цэнтральнаеўрапейскі летні час", "MST": "Паўночнаамерыканскі горны стандартны час", "EST": "Паўночнаамерыканскі ўсходні стандартны час", "TMST": "Летні час Туркменістана", "WEZ": "Заходнееўрапейскі стандартны час"},
	}
}

// Locale returns the current translators string locale
func (be *be_BY) Locale() string {
	return be.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'be_BY'
func (be *be_BY) PluralsCardinal() []locales.PluralRule {
	return be.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'be_BY'
func (be *be_BY) PluralsOrdinal() []locales.PluralRule {
	return be.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'be_BY'
func (be *be_BY) PluralsRange() []locales.PluralRule {
	return be.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'be_BY'
func (be *be_BY) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)
	nMod10 := math.Mod(n, 10)

	if nMod10 == 1 && nMod100 != 11 {
		return locales.PluralRuleOne
	} else if nMod10 >= 2 && nMod10 <= 4 && nMod100 < 12 && nMod100 > 14 {
		return locales.PluralRuleFew
	} else if (nMod10 == 0) || (nMod10 >= 5 && nMod10 <= 9) || (nMod100 >= 11 && nMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'be_BY'
func (be *be_BY) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if (nMod10 == 2 || nMod10 == 3) && (nMod100 != 12 && nMod100 != 13) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'be_BY'
func (be *be_BY) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (be *be_BY) MonthAbbreviated(month time.Month) string {
	return be.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (be *be_BY) MonthsAbbreviated() []string {
	return be.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (be *be_BY) MonthNarrow(month time.Month) string {
	return be.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (be *be_BY) MonthsNarrow() []string {
	return be.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (be *be_BY) MonthWide(month time.Month) string {
	return be.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (be *be_BY) MonthsWide() []string {
	return be.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (be *be_BY) WeekdayAbbreviated(weekday time.Weekday) string {
	return be.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (be *be_BY) WeekdaysAbbreviated() []string {
	return be.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (be *be_BY) WeekdayNarrow(weekday time.Weekday) string {
	return be.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (be *be_BY) WeekdaysNarrow() []string {
	return be.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (be *be_BY) WeekdayShort(weekday time.Weekday) string {
	return be.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (be *be_BY) WeekdaysShort() []string {
	return be.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (be *be_BY) WeekdayWide(weekday time.Weekday) string {
	return be.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (be *be_BY) WeekdaysWide() []string {
	return be.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'be_BY' and handles both Whole and Real numbers based on 'v'
func (be *be_BY) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(be.group) - 1; j >= 0; j-- {
					b = append(b, be.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, be.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'be_BY' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (be *be_BY) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, be.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, be.percentSuffix...)

	b = append(b, be.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'be_BY'
func (be *be_BY) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := be.currencies[currency]
	l := len(s) + len(symbol) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(be.group) - 1; j >= 0; j-- {
					b = append(b, be.group[j])
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

	if num < 0 {
		b = append(b, be.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, be.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'be_BY'
// in accounting notation.
func (be *be_BY) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := be.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, be.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(be.group) - 1; j >= 0; j-- {
					b = append(b, be.group[j])
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

		b = append(b, be.currencyNegativePrefix[0])

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
			b = append(b, be.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, be.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'be_BY'
func (be *be_BY) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'be_BY'
func (be *be_BY) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'be_BY'
func (be *be_BY) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, be.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'be_BY'
func (be *be_BY) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, be.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, be.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'be_BY'
func (be *be_BY) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'be_BY'
func (be *be_BY) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'be_BY'
func (be *be_BY) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'be_BY'
func (be *be_BY) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := be.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
