package es_MX

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type es_MX struct {
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

// New returns a new instance of translator for the 'es_MX' locale
func New() locales.Translator {
	return &es_MX{
		locale:                 "es_MX",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "Af", "ALK ", "ALL ", "AMD ", "Naf", "Kz", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "$a", "ATS ", "AU$", "Afl.", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "Br", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CN¥", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "RD$", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "$", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TL", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"},
		monthsNarrow:           []string{"", "E", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
		daysAbbreviated:        []string{"dom.", "lun.", "mar.", "mié.", "jue.", "vie.", "sáb."},
		daysNarrow:             []string{"D", "L", "M", "M", "J", "V", "S"},
		daysShort:              []string{"do", "lu", "ma", "mi", "ju", "vi", "sá"},
		daysWide:               []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		periodsAbbreviated:     []string{"a. m.", "p. m."},
		periodsNarrow:          []string{"a. m.", "p. m."},
		periodsWide:            []string{"a. m.", "p. m."},
		erasAbbreviated:        []string{"a. C.", "d. C."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"antes de Cristo", "después de Cristo"},
		timezones:              map[string]string{"SAST": "Hora de Sudáfrica", "MDT": "Hora de verano de la montaña de Norteamérica", "WIT": "Hora de Indonesia Oriental", "HNT": "Hora estándar de Terranova", "WAT": "Hora estándar de África Occidental", "ACST": "Hora estándar de Australia Central", "ACDT": "Hora de verano de Australia Central", "HADT": "Hora de verano de Hawái-Aleutiano", "ART": "Hora estándar de Argentina", "WITA": "Hora de Indonesia Central", "CLST": "Hora de verano de Chile", "AKST": "Hora estándar de Alaska", "HAT": "Hora de verano de Terranova", "BOT": "Hora de Bolivia", "CLT": "Hora estándar de Chile", "VET": "Hora de Venezuela", "ADT": "Hora de verano del Atlántico", "AEDT": "Hora de verano de Australia Oriental", "NZDT": "Hora de verano de Nueva Zelanda", "CAT": "Hora de África Central", "MEZ": "Hora estándar de Europa Central", "SRT": "Hora de Surinam", "WART": "Hora estándar de Argentina Occidental", "GYT": "Hora de Guyana", "AKDT": "Hora de verano de Alaska", "EST": "Hora estándar del este de Norteamérica", "OEZ": "Hora estándar de Europa del Este", "NZST": "Hora estándar de Nueva Zelanda", "PST": "Hora estándar del Pacífico de Norteamérica", "PDT": "Hora de verano del Pacífico de Norteamérica", "ACWDT": "Hora de verano de Australia Centroccidental", "WESZ": "Hora de verano de Europa Occidental", "HAST": "Hora estándar de Hawái-Aleutiano", "HKST": "Hora de verano de Hong Kong", "MESZ": "Hora de verano de Europa Central", "EAT": "Hora de África Oriental", "COST": "Hora de verano de Colombia", "WIB": "Hora de Indonesia Occidental", "LHST": "Hora estándar de Lord Howe", "GMT": "Hora del meridiano de Greenwich", "ChST": "Hora estándar de Chamorro", "ARST": "Hora de verano de Argentina", "SGT": "Hora de Singapur", "HKT": "Hora estándar de Hong Kong", "CHAST": "Hora estándar de Chatham", "UYST": "Hora de verano de Uruguay", "ACWST": "Hora estándar de Australia Centroccidental", "ECT": "Hora de Ecuador", "IST": "Hora estándar de la India", "TMT": "Hora estándar de Turkmenistán", "OESZ": "Hora de verano de Europa del Este", "∅∅∅": "Hora de verano de Brasilia", "BT": "Hora de Bután", "WEZ": "Hora estándar de Europa Occidental", "JDT": "Hora de verano de Japón", "TMST": "Hora de verano de Turkmenistán", "AEST": "Hora estándar de Australia Oriental", "UYT": "Hora estándar de Uruguay", "WARST": "Hora de verano de Argentina Occidental", "CDT": "Hora de verano del centro de Norteamérica", "AWST": "Hora estándar de Australia Occidental", "WAST": "Hora de verano de África Occidental", "COT": "Hora estándar de Colombia", "JST": "Hora estándar de Japón", "AST": "Hora estándar del Atlántico", "AWDT": "Hora de verano de Australia Occidental", "MST": "Hora estándar de la montaña de Norteamérica", "MYT": "Hora de Malasia", "EDT": "Hora de verano del este de Norteamérica", "CHADT": "Hora de verano de Chatham", "LHDT": "Hora de verano de Lord Howe", "GFT": "Hora de la Guayana Francesa", "CST": "Hora estándar del centro de Norteamérica"},
	}
}

// Locale returns the current translators string locale
func (es *es_MX) Locale() string {
	return es.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'es_MX'
func (es *es_MX) PluralsCardinal() []locales.PluralRule {
	return es.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'es_MX'
func (es *es_MX) PluralsOrdinal() []locales.PluralRule {
	return es.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'es_MX'
func (es *es_MX) PluralsRange() []locales.PluralRule {
	return es.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'es_MX'
func (es *es_MX) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'es_MX'
func (es *es_MX) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'es_MX'
func (es *es_MX) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (es *es_MX) MonthAbbreviated(month time.Month) string {
	return es.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (es *es_MX) MonthsAbbreviated() []string {
	return es.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (es *es_MX) MonthNarrow(month time.Month) string {
	return es.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (es *es_MX) MonthsNarrow() []string {
	return es.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (es *es_MX) MonthWide(month time.Month) string {
	return es.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (es *es_MX) MonthsWide() []string {
	return es.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (es *es_MX) WeekdayAbbreviated(weekday time.Weekday) string {
	return es.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (es *es_MX) WeekdaysAbbreviated() []string {
	return es.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (es *es_MX) WeekdayNarrow(weekday time.Weekday) string {
	return es.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (es *es_MX) WeekdaysNarrow() []string {
	return es.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (es *es_MX) WeekdayShort(weekday time.Weekday) string {
	return es.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (es *es_MX) WeekdaysShort() []string {
	return es.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (es *es_MX) WeekdayWide(weekday time.Weekday) string {
	return es.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (es *es_MX) WeekdaysWide() []string {
	return es.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'es_MX' and handles both Whole and Real numbers based on 'v'
func (es *es_MX) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'es_MX' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (es *es_MX) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(es.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, es.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'es_MX'
func (es *es_MX) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, es.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'es_MX'
// in accounting notation.
func (es *es_MX) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, es.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, es.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, es.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'es_MX'
func (es *es_MX) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'es_MX'
func (es *es_MX) FmtDateMedium(t time.Time) string {

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
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'es_MX'
func (es *es_MX) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'es_MX'
func (es *es_MX) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, es.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'es_MX'
func (es *es_MX) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'es_MX'
func (es *es_MX) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'es_MX'
func (es *es_MX) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'es_MX'
func (es *es_MX) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := es.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
