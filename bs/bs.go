package bs

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type bs struct {
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

// New returns a new instance of translator for the 'bs' locale
func New() locales.Translator {
	return &bs{
		locale:                 "bs",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6, 2, 4},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED", "AFA ", "AFN", "ALK ", "ALL", "AMD", "ANG", "AOA", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS", "ATS ", "AUD", "AWG", "AZM ", "AZN", "BAD ", "KM", "BAN ", "BBD", "BDT", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN", "BGO ", "BHD", "BIF", "BMD", "BND", "BOB", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL", "BRN ", "BRR ", "BRZ ", "BSD", "BTN", "BUK ", "BWP", "BYB ", "BYR", "BZD", "CAD", "CDF", "CHE ", "CHF", "CHW ", "CLE ", "CLF ", "CLP", "CNX ", "CNY", "COP", "COU ", "CRC", "CSD ", "CSK ", "CUC", "CUP", "CVE", "CYP ", "CZK", "DDM ", "DEM ", "DJF", "DKK", "DOP", "DZD", "ECS ", "ECV ", "EEK ", "EGP", "ERN", "ESA ", "ESB ", "ESP ", "ETB", "€", "FIM ", "FJD", "FKP", "FRF ", "GBP", "GEK ", "GEL", "GHC ", "GHS", "GIP", "GMD", "GNF", "GNS ", "GQE ", "GRD ", "GTQ", "GWE ", "GWP ", "GYD", "HKD", "HNL", "HRD ", "kn", "HTG", "HUF", "IDR", "IEP ", "ILP ", "ILR ", "ILS", "₹", "IQD", "IRR", "ISJ ", "ISK", "ITL ", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH ", "KRO ", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL ", "LTL", "LTT ", "LUC ", "LUF ", "LUL ", "LVL", "LVR ", "LYD", "MAD", "MAF ", "MCF ", "MDC ", "MDL", "MGA", "MGF ", "MKD", "MKN ", "MLF ", "MMK", "MNT", "MOP", "MRO", "MTL ", "MTP ", "MUR", "MVP ", "MVR", "MWK", "MXN", "MXP ", "MXV ", "MYR", "MZE ", "MZM ", "MZN", "NAD", "NGN", "NIC ", "NIO", "NLG ", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI ", "PEN", "PES ", "PGK", "PHP", "PKR", "PLN", "PLZ ", "PTE ", "PYG", "QAR", "RHD ", "ROL ", "RON", "din.", "RUB", "RUR ", "RWF", "SAR", "SBD", "SCR", "SDD ", "SDG", "SDP ", "SEK", "SGD", "SHP", "SIT ", "SKK ", "SLL", "SOS", "SRD", "SRG ", "SSP", "STD", "SUR ", "SVC ", "SYP", "SZL", "฿", "TJR ", "TJS", "TMM ", "TMT", "TND", "TOP", "TPE ", "TRL ", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK ", "UGS ", "UGX", "USD", "USN ", "USS ", "UYI ", "UYP ", "UYU", "UZS", "VEB ", "VEF", "₫", "VNN ", "VUV", "WST", "FCFA", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD", "XDR ", "XEU ", "XFO ", "XFU ", "CFA", "XPD ", "XPF", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR", "ZMK ", "ZMW", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januar", "februar", "mart", "april", "maj", "juni", "juli", "august", "septembar", "oktobar", "novembar", "decembar"},
		daysAbbreviated:        []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysNarrow:             []string{"N", "P", "U", "S", "Č", "P", "S"},
		daysShort:              []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysWide:               []string{"nedjelja", "ponedjeljak", "utorak", "srijeda", "četvrtak", "petak", "subota"},
		periodsAbbreviated:     []string{"prijepodne", "popodne"},
		periodsNarrow:          []string{"prijepodne", "popodne"},
		periodsWide:            []string{"prije podne", "popodne"},
		erasAbbreviated:        []string{"p. n. e.", "n. e."},
		erasNarrow:             []string{"pr.n.e.", "AD"},
		erasWide:               []string{"Prije nove ere", "Nove ere"},
		timezones:              map[string]string{"JST": "Standardno japansko vrijeme", "AEST": "Standardno istočnoaustralijsko vrijeme", "GYT": "Gvajansko vrijeme", "WARST": "Zapadnoargentinsko ljetno računanje vremena", "HAST": "Standardno havajsko-aleućansko vrijeme", "MYT": "Malezijsko vrijeme", "GFT": "Vrijeme Francuske Gvajane", "HKST": "Hongkonško ljetno računanje vremena", "SGT": "Standardno singapursko vrijeme", "AEDT": "Istočnoaustralijsko ljetno računanje vremena", "BT": "Butansko vrijeme", "UYT": "Standardno urugvajsko vrijeme", "CHADT": "Čatamsko ljetno računanje vremena", "EDT": "Ljetno računanje istočnog vremena (SAD)", "OEZ": "Standardno istočnoevropsko vrijeme", "ECT": "Ekvadorsko vrijeme", "ACWDT": "Australijsko centralno zapadno ljetno računanje vremena", "PST": "Standardno pacifičko vrijeme (SAD)", "TMT": "Standardno turkmenistansko vrijeme", "LHST": "Standardno vrijeme na Ostrvu Lord Hau", "MDT": "Makao letnje računanje vremena", "WITA": "Centralnoindonezijsko vrijeme", "IST": "Standardno indijsko vrijeme", "AKDT": "Aljaskansko ljetno računanje vremena", "WAST": "Zapadnoafričko ljetno računanje vremena", "WIT": "Istočnoindonezijsko vrijeme", "LHDT": "Ljetno računanje vremena na Ostrvu Lord Hau", "WART": "Standardno zapadnoargentinsko vrijeme", "AKST": "Standardno aljaskansko vrijeme", "ADT": "Ljetno računanje atlantskog vremena", "ACWST": "Standardno australijsko centralno zapadno vrijeme", "AWDT": "Zapadnoaustralijsko ljetno računanje vremena", "COT": "Standardno kolumbijsko vrijeme", "WIB": "Zapadnoindonezijsko vrijeme", "ChST": "Standardno čamorsko vrijeme", "MST": "Makao standardno vreme", "CHAST": "Standardno čatamsko vrijeme", "TMST": "Turkmenistansko ljetno računanje vremena", "CLT": "Standardno čileansko vrijeme", "ACST": "Standardno centralnoaustralijsko vrijeme", "HKT": "Standardno hongkonško vrijeme", "ARST": "Argentinsko ljetno računanje vremena", "SAST": "Standardno južnoafričko vrijeme", "∅∅∅": "Amazonsko ljetno računanje vremena", "WESZ": "Zapadnoevropsko ljetno računanje vremena", "COST": "Kolumbijsko ljetno računanje vremena", "VET": "Venecuelansko vrijeme", "AST": "Standardno atlantsko vrijeme", "HADT": "Havajsko-aleućansko ljetno rečunanje vremena", "CAT": "Centralnoafričko vrijeme", "CLST": "Čileansko ljetno računanje vremena", "UYST": "Urugvajsko ljetno računanje vremena", "MEZ": "Standardno centralnoevropsko vrijeme", "WAT": "Standardno zapadnoafričko vrijeme", "HAT": "Njufaundlendsko ljetno računanje vremena", "GMT": "Vrijeme po Grinviču", "NZST": "Standardno novozelandsko vrijeme", "ACDT": "Centralnoaustralijsko ljetno računanje vremena", "NZDT": "Novozelandsko ljetno računanje vremena", "WEZ": "Standardno zapadnoevropsko vrijeme", "OESZ": "Istočnoevropsko ljetno računanje vremena", "ART": "Standardno argentinsko vrijeme", "EST": "Standardno istočno vrijeme (SAD)", "MESZ": "Centralnoevropsko ljetno računanje vremena", "SRT": "Surinamsko vrijeme", "CDT": "Ljetno računanje centralnog vremena (SAD)", "AWST": "Standardno zapadnoaustralijsko vrijeme", "JDT": "Japansko ljetno računanje vremena", "BOT": "Bolivijsko vrijeme", "PDT": "Ljetno računanje pacifičkog vremena (SAD)", "HNT": "Standardno njufaundlendsko vrijeme", "EAT": "Istočnoafričko vrijeme", "CST": "Standardno centralno vrijeme (SAD)"},
	}
}

// Locale returns the current translators string locale
func (bs *bs) Locale() string {
	return bs.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bs'
func (bs *bs) PluralsCardinal() []locales.PluralRule {
	return bs.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bs'
func (bs *bs) PluralsOrdinal() []locales.PluralRule {
	return bs.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bs'
func (bs *bs) PluralsRange() []locales.PluralRule {
	return bs.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bs'
func (bs *bs) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && iMod100 < 12 && iMod100 > 14) || (fMod10 >= 2 && fMod10 <= 4 && fMod100 < 12 && fMod100 > 14) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bs'
func (bs *bs) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bs'
func (bs *bs) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := bs.CardinalPluralRule(num1, v1)
	end := bs.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (bs *bs) MonthAbbreviated(month time.Month) string {
	return bs.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bs *bs) MonthsAbbreviated() []string {
	return bs.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bs *bs) MonthNarrow(month time.Month) string {
	return bs.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bs *bs) MonthsNarrow() []string {
	return bs.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bs *bs) MonthWide(month time.Month) string {
	return bs.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bs *bs) MonthsWide() []string {
	return bs.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bs *bs) WeekdayAbbreviated(weekday time.Weekday) string {
	return bs.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bs *bs) WeekdaysAbbreviated() []string {
	return bs.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bs *bs) WeekdayNarrow(weekday time.Weekday) string {
	return bs.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bs *bs) WeekdaysNarrow() []string {
	return bs.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bs *bs) WeekdayShort(weekday time.Weekday) string {
	return bs.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bs *bs) WeekdaysShort() []string {
	return bs.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bs *bs) WeekdayWide(weekday time.Weekday) string {
	return bs.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bs *bs) WeekdaysWide() []string {
	return bs.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bs' and handles both Whole and Real numbers based on 'v'
func (bs *bs) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bs' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bs *bs) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bs.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bs'
func (bs *bs) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, bs.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bs'
// in accounting notation.
func (bs *bs) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, bs.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bs.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, bs.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bs'
func (bs *bs) FmtDateShort(t time.Time) string {

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

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'bs'
func (bs *bs) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bs'
func (bs *bs) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bs'
func (bs *bs) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, bs.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bs'
func (bs *bs) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bs'
func (bs *bs) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bs'
func (bs *bs) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bs'
func (bs *bs) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bs.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
