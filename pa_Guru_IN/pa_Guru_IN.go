package pa_Guru_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type pa_Guru_IN struct {
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

// New returns a new instance of translator for the 'pa_Guru_IN' locale
func New() locales.Translator {
	return &pa_Guru_IN{
		locale:             "pa_Guru_IN",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     []locales.PluralRule{6},
		pluralsRange:       []locales.PluralRule{2, 6},
		decimal:            "٫",
		group:              ",",
		minus:              "-",
		percent:            "٪",
		perMille:           "؉",
		timeSeparator:      ":",
		inifinity:          "∞",
		currencies:         []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		monthsAbbreviated:  []string{"", "ਜਨ", "ਫ਼ਰ", "ਮਾਰਚ", "ਅਪ੍ਰੈ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾ", "ਅਗ", "ਸਤੰ", "ਅਕਤੂ", "ਨਵੰ", "ਦਸੰ"},
		monthsNarrow:       []string{"", "ਜ", "ਫ਼", "ਮਾ", "ਅ", "ਮ", "ਜੂ", "ਜੁ", "ਅ", "ਸ", "ਅ", "ਨ", "ਦ"},
		monthsWide:         []string{"", "ਜਨਵਰੀ", "ਫ਼ਰਵਰੀ", "ਮਾਰਚ", "ਅਪ੍ਰੈਲ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾਈ", "ਅਗਸਤ", "ਸਤੰਬਰ", "ਅਕਤੂਬਰ", "ਨਵੰਬਰ", "ਦਸੰਬਰ"},
		daysAbbreviated:    []string{"ਐਤ", "ਸੋਮ", "ਮੰਗਲ", "ਬੁੱਧ", "ਵੀਰ", "ਸ਼ੁੱਕਰ", "ਸ਼ਨਿੱਚਰ"},
		daysNarrow:         []string{"ਐ", "ਸੋ", "ਮੰ", "ਬੁੱ", "ਵੀ", "ਸ਼ੁੱ", "ਸ਼"},
		daysShort:          []string{"ਐਤ", "ਸੋਮ", "ਮੰਗ", "ਬੁੱਧ", "ਵੀਰ", "ਸ਼ੁੱਕ", "ਸ਼ਨਿੱ"},
		daysWide:           []string{"ਐਤਵਾਰ", "ਸੋਮਵਾਰ", "ਮੰਗਲਵਾਰ", "ਬੁੱਧਵਾਰ", "ਵੀਰਵਾਰ", "ਸ਼ੁੱਕਰਵਾਰ", "ਸ਼ਨਿੱਚਰਵਾਰ"},
		periodsAbbreviated: []string{"ਪੂ.ਦੁ.", "ਬਾ.ਦੁ."},
		periodsNarrow:      []string{"ਸ.", "ਸ਼."},
		periodsWide:        []string{"ਪੂ.ਦੁ.", "ਬਾ.ਦੁ."},
		erasAbbreviated:    []string{"ਈ. ਪੂ.", "ਸੰਨ"},
		erasNarrow:         []string{"ਈ. ਪੂ.", "ਸੰਨ"},
		erasWide:           []string{"ਈਸਵੀ ਪੂਰਵ", "ਈਸਵੀ ਸੰਨ"},
		timezones:          map[string]string{"UYT": "ਉਰੂਗਵੇ ਮਿਆਰੀ ਸਮਾਂ", "GFT": "ਫ੍ਰੈਂਚ ਗੁਏਨਾ ਸਮਾਂ", "VET": "ਵੈਨੇਜ਼ੂਏਲਾ ਸਮਾਂ", "MESZ": "ਮੱਧ ਯੂਰਪੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "SGT": "ਸਿੰਗਾਪੁਰ ਮਿਆਰੀ ਸਮਾਂ", "WAST": "ਪੱਛਮੀ ਅਫਰੀਕਾ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "HNT": "ਨਿਊਫਾਉਂਡਲੈਂਡ ਮਿਆਰੀ ਸਮਾਂ", "LHST": "ਲੌਰਡ ਹੋਵੇ ਮਿਆਰੀ ਸਮਾਂ", "JDT": "ਜਪਾਨ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "AKST": "ਅਲਾਸਕਾ ਮਿਆਰੀ ਸਮਾਂ", "ACWST": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪੱਛਮੀ ਮਿਆਰੀ ਸਮਾਂ", "EST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੂਰਬੀ ਮਿਆਰੀ ਸਮਾਂ", "COT": "ਕੋਲੰਬੀਆ ਮਿਆਰੀ ਸਮਾਂ", "TMT": "ਤੁਰਕਮੇਨਿਸਤਾਨ ਮਿਆਰੀ ਸਮਾਂ", "ACDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "HKST": "ਹਾਂਗ ਕਾਂਗ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "PST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੈਸਿਫਿਕ ਮਿਆਰੀ ਸਮਾਂ", "PDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੈਸਿਫਿਕ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "HKT": "ਹਾਂਗ ਕਾਂਗ ਮਿਆਰੀ ਸਮਾਂ", "ACWDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪੱਛਮੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "UYST": "ਉਰੂਗਵੇ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "BT": "ਭੂਟਾਨ ਸਮਾਂ", "ECT": "ਇਕਵੇਡੋਰ ਸਮਾਂ", "WEZ": "ਪੱਛਮੀ ਯੂਰਪੀ ਮਿਆਰੀ ਸਮਾਂ", "ChST": "ਚਾਮੋਰੋ ਮਿਆਰੀ ਸਮਾਂ", "WARST": "ਪੱਛਮੀ ਅਰਜਨਟੀਨਾ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "SRT": "ਸੂਰੀਨਾਮ ਸਮਾਂ", "MDT": "ਮਕਾਉ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "CHAST": "ਚੈਥਮ ਮਿਆਰੀ ਸਮਾਂ", "AST": "ਅਟਲਾਂਟਿਕ ਮਿਆਰੀ ਸਮਾਂ", "NZST": "ਨਿਊਜ਼ੀਲੈਂਡ ਮਿਆਰੀ ਸਮਾਂ", "EDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੂਰਬੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "CLST": "ਚਿਲੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "∅∅∅": "ਅਜੋਰੇਸ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "MEZ": "ਮੱਧ ਯੂਰਪੀ ਮਿਆਰੀ ਸਮਾਂ", "OEZ": "ਪੂਰਬੀ ਯੂਰਪੀ ਮਿਆਰੀ ਸਮਾਂ", "AEST": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੂਰਬੀ ਮਿਆਰੀ ਸਮਾਂ", "CST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਕੇਂਦਰੀ ਮਿਆਰੀ ਸਮਾਂ", "ARST": "ਅਰਜਨਟੀਨਾ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "COST": "ਕੋਲੰਬੀਆ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "WESZ": "ਪੱਛਮੀ ਯੂਰਪੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "NZDT": "ਨਿਊਜ਼ੀਲੈਂਡ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "TMST": "ਤੁਰਕਮੇਨਿਸਤਾਨ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "WAT": "ਪੱਛਮੀ ਅਫਰੀਕਾ ਮਿਆਰੀ ਸਮਾਂ", "MYT": "ਮਲੇਸ਼ੀਆ ਸਮਾਂ", "CDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਕੇਂਦਰੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "MST": "ਮਕਾਉ ਮਿਆਰੀ ਸਮਾਂ", "SAST": "ਦੱਖਣੀ ਅਫ਼ਰੀਕਾ ਮਿਆਰੀ ਸਮਾਂ", "IST": "ਭਾਰਤੀ ਮਿਆਰੀ ਸਮਾਂ", "EAT": "ਪੂਰਬੀ ਅਫਰੀਕਾ ਸਮਾਂ", "LHDT": "ਲੌਰਡ ਹੋਵੇ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "ACST": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਮਿਆਰੀ ਸਮਾਂ", "ADT": "ਅਟਲਾਂਟਿਕ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "CAT": "ਕੇਂਦਰੀ ਅਫਰੀਕਾ ਸਮਾਂ", "AEDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੂਰਬੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "ART": "ਅਰਜਨਟੀਨਾ ਮਿਆਰੀ ਸਮਾਂ", "HAST": "ਹਵਾਈ-ਅਲੇਯੂਸ਼ਿਅਨ ਮਿਆਰੀ ਸਮਾਂ", "GYT": "ਗੁਯਾਨਾ ਸਮਾਂ", "WART": "ਪੱਛਮੀ ਅਰਜਨਟੀਨਾ ਮਿਆਰੀ ਸਮਾਂ", "WITA": "ਮੱਧ ਇੰਡੋਨੇਸ਼ੀਆਈ ਸਮਾਂ", "AWDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੱਛਮੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "OESZ": "ਪੂਰਬੀ ਯੂਰਪੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "JST": "ਜਪਾਨ ਮਿਆਰੀ ਸਮਾਂ", "AKDT": "ਅਲਾਸਕਾ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "WIB": "ਪੱਛਮੀ ਇੰਡੋਨੇਸ਼ੀਆ ਸਮਾਂ", "HAT": "ਨਿਊਫਾਉਂਡਲੈਂਡ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "GMT": "ਗ੍ਰੀਨਵਿਚ ਮੀਨ ਟਾਈਮ", "CLT": "ਚਿਲੀ ਮਿਆਰੀ ਸਮਾਂ", "WIT": "ਪੂਰਬੀ ਇੰਡੋਨੇਸ਼ੀਆ ਸਮਾਂ", "CHADT": "ਚੈਥਮ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "AWST": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੱਛਮੀ ਮਿਆਰੀ ਸਮਾਂ", "HADT": "ਹਵਾਈ-ਅਲੇਯੂਸ਼ਿਅਨ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "BOT": "ਬੋਲੀਵੀਆ ਸਮਾਂ"},
	}
}

// Locale returns the current translators string locale
func (pa *pa_Guru_IN) Locale() string {
	return pa.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pa_Guru_IN'
func (pa *pa_Guru_IN) PluralsCardinal() []locales.PluralRule {
	return pa.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pa_Guru_IN'
func (pa *pa_Guru_IN) PluralsOrdinal() []locales.PluralRule {
	return pa.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pa_Guru_IN'
func (pa *pa_Guru_IN) PluralsRange() []locales.PluralRule {
	return pa.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := pa.CardinalPluralRule(num1, v1)
	end := pa.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pa *pa_Guru_IN) MonthAbbreviated(month time.Month) string {
	return pa.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pa *pa_Guru_IN) MonthsAbbreviated() []string {
	return pa.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pa *pa_Guru_IN) MonthNarrow(month time.Month) string {
	return pa.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pa *pa_Guru_IN) MonthsNarrow() []string {
	return pa.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pa *pa_Guru_IN) MonthWide(month time.Month) string {
	return pa.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pa *pa_Guru_IN) MonthsWide() []string {
	return pa.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pa *pa_Guru_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return pa.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pa *pa_Guru_IN) WeekdaysAbbreviated() []string {
	return pa.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pa *pa_Guru_IN) WeekdayNarrow(weekday time.Weekday) string {
	return pa.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pa *pa_Guru_IN) WeekdaysNarrow() []string {
	return pa.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pa *pa_Guru_IN) WeekdayShort(weekday time.Weekday) string {
	return pa.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pa *pa_Guru_IN) WeekdaysShort() []string {
	return pa.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pa *pa_Guru_IN) WeekdayWide(weekday time.Weekday) string {
	return pa.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pa *pa_Guru_IN) WeekdaysWide() []string {
	return pa.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pa_Guru_IN' and handles both Whole and Real numbers based on 'v'
func (pa *pa_Guru_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(pa.decimal) - 1; j >= 0; j-- {
				b = append(b, pa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, pa.group[0])
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
		b = append(b, pa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pa_Guru_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pa *pa_Guru_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(pa.decimal) - 1; j >= 0; j-- {
				b = append(b, pa.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pa.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pa.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(pa.decimal) - 1; j >= 0; j-- {
				b = append(b, pa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, pa.group[0])
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

	if num < 0 {
		b = append(b, pa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pa_Guru_IN'
// in accounting notation.
func (pa *pa_Guru_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pa.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(pa.decimal) - 1; j >= 0; j-- {
				b = append(b, pa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, pa.group[0])
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

		b = append(b, pa.minus[0])

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
			b = append(b, pa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, pa.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pa_Guru_IN'
func (pa *pa_Guru_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pa.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
