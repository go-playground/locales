package sd_Deva

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sd_Deva struct {
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

// New returns a new instance of translator for the 'sd_Deva' locale
func New() locales.Translator {
	return &sd_Deva{
		locale:                 "sd_Deva",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "जन", "फर", "मार्च", "अप्रै", "मई", "जून", "जु", "अग", "सितं", "अक्टू", "नवं", "दिसं"},
		monthsNarrow:           []string{"", "ज", "फ़", "म", "अ", "म", "जू", "जु", "अ", "सि", "अ", "न", "द"},
		monthsWide:             []string{"", "जनवरी", "फरवरी", "मार्चु", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्टूबर", "नवंबर", "दिसंबर"},
		daysAbbreviated:        []string{"आर्त", "सू", "मंग", "बुध", "विस", "जुम", "छंछ"},
		daysNarrow:             []string{"आ", "म", "मं", "बु", "वि", "शु", "छं"},
		daysShort:              []string{"آچر", "سومر", "اڱارو", "اربع", "خميس", "جمعو", "ڇنڇر"},
		daysWide:               []string{"आर्तवार", "सूमर", "मंगलु", "बुधर", "विस्पत", "जुमओ", "छंछर"},
		periodsAbbreviated:     []string{"صبح، منجهند", "شام، منجهند"},
		periodsNarrow:          []string{"صبح، منجهند", "منجهند، شام"},
		periodsWide:            []string{"मंझंदि का पहिंरो", "मंझंदि को पोए"},
		erasAbbreviated:        []string{"बीसी", "एडी"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "آسٽريليا جو مرڪزي ڏينهن جو وقت", "ACST": "آسٽريليا جو مرڪزي معياري وقت", "ACWDT": "آسٽريليا جو مرڪزي مغربي ڏينهن جو وقت", "ACWST": "آسٽريليا جو مرڪزي مغربي معياري وقت", "ADT": "अंटलांटिक जे ॾींह वारो भेरो", "AEDT": "آسٽريليا جو مشرقي ڏينهن جو وقت", "AEST": "آسٽريليا جو مشرقي معياري وقت", "AKDT": "الاسڪا جي ڏينهن جو وقت", "AKST": "الاسڪا جو معياري وقت", "ARST": "ارجنٽائن جي اونهاري جو وقت", "ART": "ارجنٽائن معياري وقت", "AST": "अंटलांटिक जे मअयारी वारो भेरो", "AWDT": "آسٽريليا جو مغربي ڏينهن جو وقت", "AWST": "آسٽريليا جو مغربي معياري وقت", "BOT": "بولويائي وقت", "BT": "ڀوٽان جو وقت", "CAT": "مرڪزي آفريقا جو وقت", "CDT": "विचो ॾींह वारो भेरो", "CHADT": "چئٿم جي ڏينهن جو وقت", "CHAST": "چئٿم جو معياري وقت", "CLST": "چلي جي اونهاري جو وقت", "CLT": "چلي جو معياري وقت", "COST": "ڪولمبيا جي اونهاري جو وقت", "COT": "ڪولمبيا جو معياري وقت", "CST": "विचो मअयारी भेरो", "ChST": "چمورو جو معياري وقت", "EAT": "اوڀر آفريڪا جو وقت", "ECT": "ايڪواڊور جو وقت", "EDT": "उभिरंदो विचो वारो भेरो", "EST": "उभिरंदो मअयारी वारो भेरो", "GFT": "فرانسيسي گيانا جو وقت", "GMT": "ग्रीन विचु मीन टाइमु", "GST": "خلج معياري وقت", "GYT": "گيانائي وقت", "HADT": "هوائي اليوٽين جي ڏينهن جو وقت", "HAST": "هوائي اليوٽين جو معياري وقت", "HAT": "نيو فائونڊ لينڊ جي ڏينهن جو وقت", "HECU": "ڪيوبا جي ڏينهن جو وقت", "HEEG": "مشرقي گرين لينڊ جي اونهاري جو وقت", "HENOMX": "شمالي مغربي ميڪسيڪو جي ڏينهن جو وقت", "HEOG": "مغربي گرين لينڊ جي اونهاري جو وقت", "HEPM": "سینٽ پیئر و میڪوئیلون جي ڏينهن جو وقت", "HEPMX": "ميڪسيڪن پيسيفڪ جي ڏينهن جو وقت", "HKST": "هانگ ڪانگ جي اونهاري جو وقت", "HKT": "هانگ ڪانگ جو معياري وقت", "HNCU": "ڪيوبا جو معياري وقت", "HNEG": "مشرقي گرين لينڊ جو معياري وقت", "HNNOMX": "شمالي مغربي ميڪسيڪو جو معياري وقت", "HNOG": "مغربي گرين لينڊ جو معياري وقت", "HNPM": "سینٽ پیئر اَئن میڪوئلون جو مانائتو وقت", "HNPMX": "ميڪسيڪن پيسيفڪ جو معياري وقت", "HNT": "نيو فائونڊ لينڊ جو معياري وقت", "IST": "ڀارت جو معياري وقت", "JDT": "جاپان جي ڏينهن جو وقت", "JST": "جاپان جو معياري وقت", "LHDT": "لورڊ هووي جي ڏينهن جو وقت", "LHST": "لورڊ هووي جو معياري وقت", "MDT": "टकरु जो ॾींह वारो भेरो", "MESZ": "सेंटरलु यूरपी गरमी वारो टाइमु", "MEZ": "सेंटरलु यूरपी मयआरी वारो टाइमु", "MST": "टकरु वारो मअयारी भेरो", "MYT": "ملائيشيا جو وقت", "NZDT": "نيوزيلينڊ جي ڏينهن جو وقت", "NZST": "نيوزيلينڊ جو معياري وقت", "OESZ": "उभिरंदो यूरोपी गरमी वारो वक्तु", "OEZ": "उभिरंदो यूरोपी मअयारी वारो वक्तु", "PDT": "पेसीफिक जो ॾींह वारो भेरो", "PST": "पेसीफिक वारो मअयारी वारो भेरो", "SAST": "ڏکڻ آفريڪا جو معياري وقت", "SGT": "سنگاپور جو معياري وقت", "SRT": "سوري نام جو وقت", "TMST": "ترڪمانستان جي اونهاري جو وقت", "TMT": "ترڪمانستان جو معياري وقت", "UYST": "يوروگائي جي اونهاري جو وقت", "UYT": "يوروگائي جو معياري وقت", "VET": "وينزويلا جو وقت", "WARST": "مغربي ارجنٽائن جي اونهاري جو وقت", "WART": "مغربي ارجنٽائن جو معياري وقت", "WAST": "اولهه آفريقا جي اونهاري جو وقت", "WAT": "اولهه آفريقا جو معياري وقت", "WESZ": "उलहंदो जे यूरोपीअ गरमी वारो वक्तु", "WEZ": "उलहंदो जे मअयारी वारो वक्तु", "WIB": "اولهه انڊونيشيا جو وقت", "WIT": "اوڀر انڊونيشيا جو وقت", "WITA": "مرڪزي انڊونيشيا جو وقت", "∅∅∅": "براسيليا جي اونهاري جو وقت"},
	}
}

// Locale returns the current translators string locale
func (sd *sd_Deva) Locale() string {
	return sd.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sd_Deva'
func (sd *sd_Deva) PluralsCardinal() []locales.PluralRule {
	return sd.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sd_Deva'
func (sd *sd_Deva) PluralsOrdinal() []locales.PluralRule {
	return sd.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sd_Deva'
func (sd *sd_Deva) PluralsRange() []locales.PluralRule {
	return sd.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sd_Deva'
func (sd *sd_Deva) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sd_Deva'
func (sd *sd_Deva) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sd_Deva'
func (sd *sd_Deva) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sd.CardinalPluralRule(num1, v1)
	end := sd.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sd *sd_Deva) MonthAbbreviated(month time.Month) string {
	return sd.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sd *sd_Deva) MonthsAbbreviated() []string {
	return sd.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sd *sd_Deva) MonthNarrow(month time.Month) string {
	return sd.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sd *sd_Deva) MonthsNarrow() []string {
	return sd.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sd *sd_Deva) MonthWide(month time.Month) string {
	return sd.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sd *sd_Deva) MonthsWide() []string {
	return sd.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayAbbreviated(weekday time.Weekday) string {
	return sd.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sd *sd_Deva) WeekdaysAbbreviated() []string {
	return sd.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayNarrow(weekday time.Weekday) string {
	return sd.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sd *sd_Deva) WeekdaysNarrow() []string {
	return sd.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayShort(weekday time.Weekday) string {
	return sd.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sd *sd_Deva) WeekdaysShort() []string {
	return sd.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayWide(weekday time.Weekday) string {
	return sd.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sd *sd_Deva) WeekdaysWide() []string {
	return sd.daysWide
}

// Decimal returns the decimal point of number
func (sd *sd_Deva) Decimal() string {
	return sd.decimal
}

// Group returns the group of number
func (sd *sd_Deva) Group() string {
	return sd.group
}

// Group returns the minus sign of number
func (sd *sd_Deva) Minus() string {
	return sd.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sd_Deva' and handles both Whole and Real numbers based on 'v'
func (sd *sd_Deva) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sd.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sd.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sd_Deva' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sd *sd_Deva) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sd.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sd.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sd_Deva'
func (sd *sd_Deva) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sd.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sd.group[0])
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

	for j := len(sd.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, sd.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, sd.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sd.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sd_Deva'
// in accounting notation.
func (sd *sd_Deva) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sd.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sd.group[0])
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

		for j := len(sd.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sd.currencyNegativePrefix[j])
		}

		b = append(b, sd.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(sd.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, sd.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sd.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sd.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sd.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sd.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sd.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sd.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
