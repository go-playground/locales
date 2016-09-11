package uk_UA

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type uk_UA struct {
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

// New returns a new instance of translator for the 'uk_UA' locale
func New() locales.Translator {
	return &uk_UA{
		locale:                 "uk_UA",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{4, 6},
		pluralsRange:           []locales.PluralRule{4, 5, 6, 2},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "січ.", "лют.", "бер.", "квіт.", "трав.", "черв.", "лип.", "серп.", "вер.", "жовт.", "лист.", "груд."},
		monthsNarrow:           []string{"", "С", "Л", "Б", "К", "Т", "Ч", "Л", "С", "В", "Ж", "Л", "Г"},
		monthsWide:             []string{"", "січня", "лютого", "березня", "квітня", "травня", "червня", "липня", "серпня", "вересня", "жовтня", "листопада", "грудня"},
		daysAbbreviated:        []string{"Нд", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"},
		daysNarrow:             []string{"Н", "П", "В", "С", "Ч", "П", "С"},
		daysShort:              []string{"Нд", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"},
		daysWide:               []string{"неділя", "понеділок", "вівторок", "середа", "четвер", "пʼятниця", "субота"},
		periodsAbbreviated:     []string{"дп", "пп"},
		periodsNarrow:          []string{"дп", "пп"},
		periodsWide:            []string{"дп", "пп"},
		erasAbbreviated:        []string{"до н. е.", "н. е."},
		erasNarrow:             []string{"до н.е.", "н.е."},
		erasWide:               []string{"до нашої ери", "нашої ери"},
		timezones:              map[string]string{"EDT": "за північноамериканським східним літнім часом", "NZST": "за стандартним часом у Новій Зеландії", "ADT": "за атлантичним літнім часом", "IST": "за індійським стандартним часом", "GYT": "за часом у Гаяні", "BT": "за часом у Бутані", "ChST": "за часом на Північних Маріанських островах", "NZDT": "за літнім часом у Новій Зеландії", "EAT": "за східноафриканським часом", "WEZ": "за західноєвропейським стандартним часом", "WART": "за стандартним західноаргентинським часом", "PST": "за північноамериканським тихоокеанським стандартним часом", "HAT": "за літнім часом у Ньюфаундленд", "WAT": "за західноафриканським стандартним часом", "MST": "MST", "CLT": "за стандартним чилійським часом", "CDT": "за північноамериканським центральним літнім часом", "ECT": "за часом в Еквадорі", "LHDT": "за літнім часом на острові Лорд-Хау", "OESZ": "за східноєвропейським літнім часом", "AEST": "за стандартним східноавстралійським часом", "TMST": "за літнім часом у Туркменістані", "HKST": "за літнім часом у Гонконзі", "UYST": "за літнім часом в Уругваї", "COT": "за стандартним колумбійським часом", "WESZ": "за західноєвропейським літнім часом", "MEZ": "за центральноєвропейським стандартним часом", "CLST": "за літнім чилійським часом", "HAST": "за стандартним гавайсько-алеутським часом", "WAST": "за західноафриканським літнім часом", "MDT": "MDT", "ACWST": "за стандартним центральнозахідним австралійським часом", "JDT": "за японським літнім часом", "MYT": "за часом у Малайзії", "AWDT": "за літнім західноавстралійським часом", "HKT": "за стандартним часом у Гонконзі", "GFT": "за часом Французької Гвіани", "ACDT": "за літнім центральноавстралійським часом", "AEDT": "за літнім східноавстралійським часом", "EST": "за північноамериканським східним стандартним часом", "CHADT": "за літнім часом на архіпелазі Чатем", "OEZ": "за східноєвропейським стандартним часом", "BOT": "за болівійським часом", "WIT": "за східноіндонезійським часом", "ART": "за стандартним аргентинським часом", "WITA": "за центральноіндонезійським часом", "SAST": "за південноафриканським часом", "ACWDT": "за літнім центральнозахідним австралійським часом", "WIB": "за західноіндонезійським часом", "MESZ": "за центральноєвропейським літнім часом", "WARST": "за літнім за західноаргентинським часом", "HADT": "за літнім гавайсько-алеутським часом", "AKST": "за стандартним часом на Алясці", "CHAST": "за стандартним часом на архіпелазі Чатем", "AWST": "за стандартним західноавстралійським часом", "HNT": "за стандартним часом у Ньюфаундленд", "SRT": "за часом у Суринамі", "UYT": "за стандартним часом в Уругваї", "GMT": "за Гринвічем", "COST": "за літнім колумбійським часом", "∅∅∅": "за літнім часом на Амазонці", "ACST": "за стандартним центральноавстралійським часом", "VET": "за часом у Венесуелі", "ARST": "за літнім аргентинським часом", "AKDT": "за літнім часом на Алясці", "JST": "за японським стандартним часом", "AST": "за атлантичним стандартним часом", "LHST": "за стандартним часом на острові Лорд-Хау", "PDT": "за північноамериканським тихоокеанським літнім часом", "CST": "за північноамериканським центральним стандартним часом", "SGT": "за часом у Сінгапурі", "CAT": "за центральноафриканським часом", "TMT": "за стандартним часом у Туркменістані"},
	}
}

// Locale returns the current translators string locale
func (uk *uk_UA) Locale() string {
	return uk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uk_UA'
func (uk *uk_UA) PluralsCardinal() []locales.PluralRule {
	return uk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uk_UA'
func (uk *uk_UA) PluralsOrdinal() []locales.PluralRule {
	return uk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uk_UA'
func (uk *uk_UA) PluralsRange() []locales.PluralRule {
	return uk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uk_UA'
func (uk *uk_UA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod100 := i % 100
	iMod10 := i % 10

	if v == 0 && iMod10 == 1 && iMod100 != 11 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && iMod100 < 12 && iMod100 > 14 {
		return locales.PluralRuleFew
	} else if (v == 0 && iMod10 == 0) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 11 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uk_UA'
func (uk *uk_UA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if nMod10 == 3 && nMod100 != 13 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uk_UA'
func (uk *uk_UA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uk.CardinalPluralRule(num1, v1)
	end := uk.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uk *uk_UA) MonthAbbreviated(month time.Month) string {
	return uk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uk *uk_UA) MonthsAbbreviated() []string {
	return uk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uk *uk_UA) MonthNarrow(month time.Month) string {
	return uk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uk *uk_UA) MonthsNarrow() []string {
	return uk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uk *uk_UA) MonthWide(month time.Month) string {
	return uk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uk *uk_UA) MonthsWide() []string {
	return uk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayAbbreviated(weekday time.Weekday) string {
	return uk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uk *uk_UA) WeekdaysAbbreviated() []string {
	return uk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayNarrow(weekday time.Weekday) string {
	return uk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uk *uk_UA) WeekdaysNarrow() []string {
	return uk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayShort(weekday time.Weekday) string {
	return uk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uk *uk_UA) WeekdaysShort() []string {
	return uk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uk *uk_UA) WeekdayWide(weekday time.Weekday) string {
	return uk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uk *uk_UA) WeekdaysWide() []string {
	return uk.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uk_UA' and handles both Whole and Real numbers based on 'v'
func (uk *uk_UA) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(uk.decimal) + len(uk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uk_UA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uk *uk_UA) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(uk.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uk_UA'
func (uk *uk_UA) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uk.currencies[currency]
	l := len(s) + len(uk.decimal) + len(uk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uk_UA'
// in accounting notation.
func (uk *uk_UA) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uk.currencies[currency]
	l := len(s) + len(uk.decimal) + len(uk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, uk.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, uk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateShort(t time.Time) string {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0xd1, 0x80, 0x27, 0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0xd1, 0x80, 0x27, 0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, uk.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0xd1, 0x80, 0x27, 0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'uk_UA'
func (uk *uk_UA) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := uk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
