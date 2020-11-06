package ff_Adlm

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ff_Adlm struct {
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

// New returns a new instance of translator for the 'ff_Adlm' locale
func New() locales.Translator {
	return &ff_Adlm{
		locale:                 "ff_Adlm",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  "⹁",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "FG", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "𞤐𞤐𞤘", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "𞤑𞤆𞤘", "𞤆𞤆𞤖", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "𞤊𞤅𞤊𞤀", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "𞤅𞤊𞤀", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: " 𞤘𞤵𞤤",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: " 𞤘𞤵𞤤",
		monthsAbbreviated:      []string{"", "𞤅𞤭𞥅𞤤𞤮", "𞤕𞤮𞤤𞤼𞤮", "𞤐𞤦𞤮𞥅𞤴𞤮", "𞤅𞤫𞥅𞤼𞤮", "𞤁𞤵𞥅𞤶𞤮", "𞤑𞤮𞤪𞤧𞤮", "𞤃𞤮𞤪𞤧𞤮", "𞤔𞤵𞤳𞤮", "𞤅𞤭𞤤𞤼𞤮", "𞤒𞤢𞤪𞤳𞤮", "𞤔𞤮𞤤𞤮", "𞤐𞤦𞤮𞤱𞤼𞤮"},
		monthsNarrow:           []string{"", "𞤅", "𞤕", "𞤄", "𞤅", "𞤁", "𞤑", "𞤃", "𞤔", "𞤅", "𞤒", "𞤔", "𞤄"},
		monthsWide:             []string{"", "𞤅𞤭𞥅𞤤𞤮", "𞤕𞤮𞤤𞤼𞤮", "𞤐𞤦𞤮𞥅𞤴𞤮", "𞤅𞤫𞥅𞤼𞤮", "𞤁𞤵𞥅𞤶𞤮", "𞤑𞤮𞤪𞤧𞤮", "𞤃𞤮𞤪𞤧𞤮", "𞤔𞤵𞤳𞤮", "𞤅𞤭𞤤𞤼𞤮", "𞤒𞤢𞤪𞤳𞤮", "𞤔𞤮𞤤𞤮", "𞤐𞤦𞤮𞤱𞤼𞤮"},
		daysAbbreviated:        []string{"𞤈𞤫𞤬", "𞤀𞥄𞤩𞤵", "𞤃𞤢𞤦", "𞤔𞤫𞤧", "𞤐𞤢𞥄𞤧", "𞤃𞤢𞤣", "𞤖𞤮𞤪"},
		daysNarrow:             []string{"𞤈", "𞤀𞥄", "𞤃", "𞤔", "𞤐", "𞤃", "𞤖"},
		daysShort:              []string{"𞤈𞤫𞤬", "𞤀𞥄𞤩𞤵", "𞤃𞤢𞤦", "𞤔𞤫𞤧", "𞤐𞤢𞥄𞤧", "𞤃𞤢𞤣", "𞤖𞤮𞤪"},
		daysWide:               []string{"𞤈𞤫𞤬𞤦𞤭𞤪𞥆𞤫", "𞤀𞥄𞤩𞤵𞤲𞥋𞤣𞤫", "𞤃𞤢𞤱𞤦𞤢𞥄𞤪𞤫", "𞤐𞤶𞤫𞤧𞤤𞤢𞥄𞤪𞤫", "𞤐𞤢𞥄𞤧𞤢𞥄𞤲𞤣𞤫", "𞤃𞤢𞤱𞤲𞤣𞤫", "𞤖𞤮𞤪𞤦𞤭𞤪𞥆𞤫"},
		periodsAbbreviated:     []string{"𞤀𞤎", "𞤇𞤎"},
		periodsNarrow:          []string{"𞤢", "𞤩"},
		periodsWide:            []string{"𞤀𞤎", "𞤇𞤎"},
		erasAbbreviated:        []string{"𞤀𞤀𞤋", "𞤇𞤀𞤋"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"𞤀𞤣𞤮 𞤀𞤲𞥆𞤢𞤦𞤭 𞤋𞥅𞤧𞤢𞥄", "𞤇𞤢𞥄𞤱𞤮 𞤀𞤲𞥆𞤢𞤦𞤭 𞤋𞥅𞤧𞤢𞥄"},
		timezones:              map[string]string{"ACDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "ACST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "ACWDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "ACWST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤥𞤦𞤮 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "ADT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤫𞤳𞤵", "AEDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "AEST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "AKDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤤𞤢𞤧𞤳𞤢𞥄 𞤲𞤣𞤫𞤲", "AKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤀𞤤𞤢𞤧𞤳𞤢𞥄 𞤲𞤣𞤫𞤲", "ARST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "ART": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤭𞤼𞤵𞤲𞥋𞤣𞤫 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "AST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤫𞤳𞤵 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫", "AWDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "AWST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤌𞤧𞤼𞤢𞤪𞤤𞤭𞤴𞤢𞥄", "BOT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤄𞤮𞤤𞤭𞤾𞤭𞤴𞤢𞥄", "BT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤄𞤵𞤼𞤢𞥄𞤲", "CAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "CDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "CHADT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤕𞤢𞥄𞤼𞤢𞤥", "CHAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤕𞤢𞤼𞤢𞥄𞤥", "CLST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤕𞤭𞤤𞤫𞥅", "CLT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤭𞤼𞤵𞤲𞥋𞤣𞤫 𞤕𞤭𞤤𞤫𞥅", "COST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤮𞤤𞤮𞤥𞤦𞤭𞤴𞤢𞥄", "COT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤭𞤼𞤵𞤲𞥋𞤣𞤫 𞤑𞤮𞤤𞤮𞤥𞤦𞤭𞤴𞤢𞥄", "CST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤥𞤦𞤮 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "ChST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤕𞤮𞤥𞤮𞥅𞤪𞤮", "EAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "ECT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤋𞤳𞤵𞤱𞤢𞤣𞤮𞥅𞤪", "EDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "EST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫𞥅𞤪𞤭 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "GFT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤘𞤢𞤴𞤢𞤲𞤢𞥄-𞤊𞤪𞤢𞤲𞤧𞤭", "GMT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤳𞤭𞤲𞤭𞥅𞤲𞥋𞤣𞤫 𞤘𞤪𞤭𞤲𞤱𞤭𞥅𞤧", "GST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤂𞤮𞥅𞤻𞤵", "GYT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤘𞤢𞤴𞤢𞤲𞤢𞥄", "HADT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤢𞤱𞤢𞥄𞤴𞤭𞥅-𞤀𞤤𞤮𞤧𞤭𞤴𞤢𞤲", "HAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤢𞤱𞤢𞥄𞤴𞤭𞥅-𞤀𞤤𞤮𞤧𞤭𞤴𞤢𞤲", "HAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤫𞤱-𞤊𞤵𞤲𞤣𞤵𞤤𞤢𞤲𞤣", "HECU": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤑𞤵𞥅𞤦𞤢𞥄", "HEEG": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "HENOMX": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤢𞤲𞤮-𞤸𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤃𞤫𞤳𞤧𞤭𞤳𞤮𞥅", "HEOG": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "HEPM": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤅𞤼. 𞤆𞤭𞤴𞤫𞥅𞤪 & 𞤃𞤭𞤳𞤫𞤤𞤮𞤲", "HEPMX": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤁𞤫𞤰𞥆𞤮 𞤃𞤫𞤳𞤧𞤭𞤳𞤮𞥅", "HKST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤮𞤲𞤺 𞤑𞤮𞤲𞤺", "HKT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤮𞤲𞤺 𞤑𞤮𞤲𞤺", "HNCU": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤑𞤵𞥅𞤦𞤢𞥄", "HNEG": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤬𞤵𞤯𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "HNNOMX": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤢𞤲𞤮-𞤸𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤃𞤫𞤳𞤧𞤭𞤳𞤮𞥅", "HNOG": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤘𞤭𞤪𞤤𞤢𞤲𞤣", "HNPM": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤅𞤼. 𞤆𞤭𞤴𞤫𞥅𞤪 & 𞤃𞤭𞤳𞤫𞤤𞤮𞤲", "HNPMX": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤁𞤫𞤰𞥆𞤮 𞤃𞤫𞤳𞤧𞤭𞤳𞤮𞥅", "HNT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤫𞤱-𞤊𞤵𞤲𞤣𞤵𞤤𞤢𞤲𞤣", "IST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤋𞤲𞤣𞤭𞤴𞤢", "JDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤭𞤨𞥆𞤮𞤲", "JST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤭𞤨𞥆𞤮𞤲", "LHDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤂𞤮𞤪𞤣𞤵-𞤖𞤮𞤱𞤫", "LHST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤂𞤮𞤪𞤣𞤵-𞤖𞤮𞤱𞤫", "MDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "MESZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤀𞤪𞤮𞥅𞤦𞤢", "MEZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤥𞤦𞤮𞥅𞤪𞤭 𞤀𞤪𞤮𞥅𞤦𞤢", "MST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤆𞤫𞤤𞥆𞤭𞤲𞤳𞤮𞥅𞤪𞤫 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "MYT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤃𞤢𞤤𞤫𞥅𞤧𞤭𞤴𞤢", "NZDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤐𞤫𞤱-𞤟𞤫𞤤𞤢𞤲𞤣𞤭", "NZST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤫𞤱-𞤟𞤫𞤤𞤢𞤲𞤣𞤭", "OESZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "OEZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "PDT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤁𞤫𞤰𞥆𞤮 𞤕𞤫𞥅𞤯𞤵 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "PST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤁𞤫𞤰𞥆𞤮 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤥𞤫𞤪𞤭𞤳𞤢𞥄", "SAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤐𞤢𞤲𞥆𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "SGT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤅𞤭𞤲𞤺𞤢𞤨𞤵𞥅𞤪", "SRT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤅𞤭𞤪𞤭𞤲𞤢𞤥", "TMST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤚𞤵𞤪𞤳𞤵𞤥𞤫𞤲𞤭𞤧𞤼𞤢𞥄𞤲", "TMT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤚𞤵𞤪𞤳𞤵𞤥𞤫𞤲𞤭𞤧𞤼𞤢𞥄𞤲", "UYST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤒𞤵𞥅𞤪𞤺𞤮𞤴", "UYT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤭𞤼𞤵𞤲𞥋𞤣𞤫 𞤒𞤵𞥅𞤪𞤺𞤮𞤴", "VET": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤜𞤫𞤲𞤭𞥅𞥁𞤮𞥅𞤤𞤢", "WARST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "WART": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤭𞤼𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢𞥄", "WAST": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "WAT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢𞥄", "WESZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "WEZ": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤢𞤱𞤪𞤵𞤲𞥋𞤣𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤀𞤪𞤮𞥅𞤦𞤢", "WIB": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤋𞤲𞤣𞤮𞤲𞤭𞥅𞤧𞤭𞤴𞤢", "WIT": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤋𞤲𞤣𞤮𞤲𞤭𞥅𞤧𞤭𞤴𞤢", "WITA": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤚𞤮𞤥𞤦𞤮𞥅𞤪𞤭 𞤋𞤲𞤣𞤮𞤲𞤭𞥅𞤧𞤭𞤴𞤢", "∅∅∅": "𞤑𞤭𞤶𞤮𞥅𞤪𞤫 𞤕𞤫𞥅𞤯𞤵 𞤄𞤪𞤢𞤧𞤭𞤤𞤭𞤴𞤢𞥄"},
	}
}

// Locale returns the current translators string locale
func (ff *ff_Adlm) Locale() string {
	return ff.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ff_Adlm'
func (ff *ff_Adlm) PluralsCardinal() []locales.PluralRule {
	return ff.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ff_Adlm'
func (ff *ff_Adlm) PluralsOrdinal() []locales.PluralRule {
	return ff.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ff_Adlm'
func (ff *ff_Adlm) PluralsRange() []locales.PluralRule {
	return ff.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ff_Adlm'
func (ff *ff_Adlm) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ff_Adlm'
func (ff *ff_Adlm) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ff_Adlm'
func (ff *ff_Adlm) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ff *ff_Adlm) MonthAbbreviated(month time.Month) string {
	return ff.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ff *ff_Adlm) MonthsAbbreviated() []string {
	return ff.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ff *ff_Adlm) MonthNarrow(month time.Month) string {
	return ff.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ff *ff_Adlm) MonthsNarrow() []string {
	return ff.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ff *ff_Adlm) MonthWide(month time.Month) string {
	return ff.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ff *ff_Adlm) MonthsWide() []string {
	return ff.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayAbbreviated(weekday time.Weekday) string {
	return ff.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ff *ff_Adlm) WeekdaysAbbreviated() []string {
	return ff.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayNarrow(weekday time.Weekday) string {
	return ff.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ff *ff_Adlm) WeekdaysNarrow() []string {
	return ff.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayShort(weekday time.Weekday) string {
	return ff.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ff *ff_Adlm) WeekdaysShort() []string {
	return ff.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ff *ff_Adlm) WeekdayWide(weekday time.Weekday) string {
	return ff.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ff *ff_Adlm) WeekdaysWide() []string {
	return ff.daysWide
}

// Decimal returns the decimal point of number
func (ff *ff_Adlm) Decimal() string {
	return ff.decimal
}

// Group returns the group of number
func (ff *ff_Adlm) Group() string {
	return ff.group
}

// Group returns the minus sign of number
func (ff *ff_Adlm) Minus() string {
	return ff.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ff_Adlm' and handles both Whole and Real numbers based on 'v'
func (ff *ff_Adlm) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ff.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ff.group) - 1; j >= 0; j-- {
					b = append(b, ff.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ff.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ff_Adlm' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ff *ff_Adlm) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ff_Adlm'
func (ff *ff_Adlm) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ff.currencies[currency]
	l := len(s) + len(symbol) + 18

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ff.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(ff.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ff.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ff.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ff.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ff_Adlm'
// in accounting notation.
func (ff *ff_Adlm) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ff.currencies[currency]
	l := len(s) + len(symbol) + 18

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ff.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ff.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ff.currencyNegativePrefix[j])
		}

		b = append(b, ff.minus[0])

	} else {

		for j := len(ff.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ff.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, ff.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ff.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ff.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ff.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ff.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ff.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ff_Adlm'
func (ff *ff_Adlm) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ff.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ff.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
