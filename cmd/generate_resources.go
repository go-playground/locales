package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"

	"github.com/go-playground/locales"

	"golang.org/x/text/unicode/cldr"

	"text/template"
)

const (
	locDir      = "../%s"
	locFilename = locDir + "/%s.go"
)

var (
	prVarFuncs = map[string]string{
		"n": "n := math.Abs(num)\n",
		"i": "i := int64(n)\n",
		// "v": "v := ...", // inherently available as argument
		"w": "w := locales.W(n, v)\n",
		"f": "f := locales.F(n, v)\n",
		"t": "t := locales.T(n, v)\n",
	}

	translators            = make(map[string]*translator)
	baseTranslators        = make(map[string]*translator)
	globalCurrenciesMap    = make(map[string]struct{}) // ["USD"] = "$" currency code, just all currencies for mapping to enum
	globCurrencyIdxMap     = make(map[string]int)      // ["USD"] = 0
	globalCurrencies       = make([]string, 0, 100)    // array of currency codes index maps to enum
	tmpl                   *template.Template
	nModRegex              = regexp.MustCompile("(n%[0-9]+)")
	iModRegex              = regexp.MustCompile("(i%[0-9]+)")
	wModRegex              = regexp.MustCompile("(w%[0-9]+)")
	fModRegex              = regexp.MustCompile("(f%[0-9]+)")
	tModRegex              = regexp.MustCompile("(t%[0-9]+)")
	groupLenRegex          = regexp.MustCompile(",([0-9#]+)\\.")
	groupLenPercentRegex   = regexp.MustCompile(",([0-9#]+)$")
	secondaryGroupLenRegex = regexp.MustCompile(",([0-9#]+),")
	requiredNumRegex       = regexp.MustCompile("([0-9]+)\\.")
	requiredDecimalRegex   = regexp.MustCompile("\\.([0-9]+)")
)

type translator struct {
	Locale         string
	BaseLocale     string
	Plurals        string
	CardinalFunc   string
	PluralsOrdinal string
	OrdinalFunc    string
	RangeFunc      string
	Decimal        string
	DecimalLen     int
	Group          string
	GroupLen       int
	Minus          string
	MinusLen       int
	Percent        string
	PercentLen     int
	PerMille       string
	PerMilleLen    int
	Currencies     string

	// FmtNumber vars
	FmtNumberExists            bool
	FmtNumberGroupLen          int
	FmtNumberSecondaryGroupLen int
	FmtNumberMinDecimalLen     int

	// FmtPercent vars
	FmtPercentExists            bool
	FmtPercentGroupLen          int
	FmtPercentSecondaryGroupLen int
	FmtPercentMinDecimalLen     int
	FmtPercentPrefix            string
	FmtPercentSuffix            string
	FmtPercentInPrefix          bool
	FmtPercentLeft              bool

	// calculation only fields
	DecimalNumberFormat string
	PercentNumberFormat string
}

func main() {

	var err error

	// load template
	tmpl, err = template.ParseGlob("*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// load CLDR recourses
	var decoder cldr.Decoder

	cldr, err := decoder.DecodePath("data/core")
	if err != nil {
		panic(err)
	}

	preProcess(cldr)
	postProcess(cldr)

	var currencies string

	for i, curr := range globalCurrencies {

		if i == 0 {
			currencies = curr + " Currency = iota\n"
			continue
		}

		currencies += curr + "\n"
	}

	if err = os.MkdirAll(fmt.Sprintf(locDir, "currency"), 0777); err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf(locFilename, "currency", "currency")

	output, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	if err := tmpl.ExecuteTemplate(output, "currencies", currencies); err != nil {
		log.Fatal(err)
	}

	output.Close()

	// after file written run gofmt on file to ensure best formatting
	cmd := exec.Command("goimports", "-w", filename)
	if err = cmd.Run(); err != nil {
		log.Panic(err)
	}

	for _, trans := range translators {

		fmt.Println("Writing Data:", trans.PercentNumberFormat, len(trans.FmtPercentPrefix), trans.FmtPercentPrefix, len(trans.FmtPercentSuffix), trans.FmtPercentSuffix, trans.Locale)

		if err = os.MkdirAll(fmt.Sprintf(locDir, trans.Locale), 0777); err != nil {
			log.Fatal(err)
		}

		filename = fmt.Sprintf(locFilename, trans.Locale, trans.Locale)

		output, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()

		if err := tmpl.ExecuteTemplate(output, "translator", trans); err != nil {
			log.Fatal(err)
		}

		output.Close()

		// after file written run gofmt on file to ensure best formatting
		cmd := exec.Command("goimports", "-w", filename)
		if err = cmd.Run(); err != nil {
			log.Panic(err)
		}
	}
}

func postProcess(cldr *cldr.CLDR) {

	var base *translator
	var found bool

	for _, trans := range translators {

		fmt.Println("Post Processing:", trans.Locale)

		// cardinal plural rules
		trans.CardinalFunc, trans.Plurals = parseCardinalPluralRuleFunc(cldr, trans.BaseLocale)

		//ordinal plural rules
		trans.OrdinalFunc, trans.PluralsOrdinal = parseOrdinalPluralRuleFunc(cldr, trans.BaseLocale)

		// if trans.Locale == "en" {
		trans.RangeFunc = parseRangePluralRuleFunc(cldr, trans.BaseLocale)
		// }

		// ignore base locales
		if trans.BaseLocale == trans.Locale {
			found = false
		} else {

			base, found = baseTranslators[trans.BaseLocale]
		}

		// Numbers

		if len(trans.Decimal) == 0 {

			if found {
				trans.DecimalLen = base.DecimalLen
				trans.Decimal = base.Decimal
			}

			if len(trans.Decimal) == 0 {
				trans.DecimalLen = 0
				trans.Decimal = "[]byte{}"
			}
		}

		if len(trans.Group) == 0 {

			if found {
				trans.GroupLen = base.GroupLen
				trans.Group = base.Group
			}

			if len(trans.Group) == 0 {
				trans.GroupLen = 0
				trans.Group = "[]byte{}"
			}
		}

		if len(trans.Minus) == 0 {

			if found {
				trans.MinusLen = base.MinusLen
				trans.Minus = base.Minus
			}

			if len(trans.Minus) == 0 {
				trans.MinusLen = 0
				trans.Minus = "[]byte{}"
			}
		}

		if len(trans.Percent) == 0 {

			if found {
				trans.PercentLen = base.PercentLen
				trans.Percent = base.Percent
			}

			if len(trans.Percent) == 0 {
				trans.PercentLen = 0
				trans.Percent = "[]byte{}"
			}
		}

		if len(trans.PerMille) == 0 {

			if found {
				trans.PerMilleLen = base.PerMilleLen
				trans.PerMille = base.PerMille
			}

			if len(trans.PerMille) == 0 {
				trans.PerMilleLen = 0
				trans.PerMille = "[]byte{}"
			}
		}

		// Currency

		// number values

		if len(trans.DecimalNumberFormat) == 0 {

			if found {
				trans.DecimalNumberFormat = base.DecimalNumberFormat
			}
		}

		if len(trans.PercentNumberFormat) == 0 {

			if found {
				trans.PercentNumberFormat = base.PercentNumberFormat
			}
		}

		ldml := cldr.RawLDML(trans.Locale)

		currencies := make([][]byte, len(globalCurrencies), len(globalCurrencies))

		for k, v := range globCurrencyIdxMap {
			currencies[v] = []byte(k)
		}

		// some just have no data...
		if ldml.Numbers != nil {

			if ldml.Numbers.Currencies != nil {
				for _, currency := range ldml.Numbers.Currencies.Currency {

					if len(currency.Symbol) == 0 {
						continue
					}

					if len(currency.Symbol[0].Data()) == 0 {
						continue
					}

					if len(currency.Type) == 0 {
						continue
					}

					currencies[globCurrencyIdxMap[currency.Type]] = []byte(currency.Symbol[0].Data())
				}
			}
		}

		trans.Currencies = fmt.Sprintf("%#v", currencies)

		parseDecimalNumberFormat(trans)
		parsePercentNumberFormat(trans)
		// trans.FmtNumberFunc = parseDecimalNumberFormat(trans.DecimalNumberFormat, trans.BaseLocale)
	}
}

// preprocesses maps, array etc... just requires multiple passes no choice....
func preProcess(cldr *cldr.CLDR) {

	for _, l := range cldr.Locales() {

		fmt.Println("Pre Processing:", l)

		split := strings.SplitN(l, "_", 2)
		baseLocale := split[0]

		trans := &translator{
			Locale:     l,
			BaseLocale: baseLocale,
		}

		// if is a base locale
		if len(split) == 1 {
			baseTranslators[baseLocale] = trans
		}

		translators[l] = trans

		// get number, currency and datetime symbols

		// number values
		ldml := cldr.RawLDML(l)

		// some just have no data...
		if ldml.Numbers != nil {

			if len(ldml.Numbers.Symbols) > 0 {

				symbol := ldml.Numbers.Symbols[0]

				if len(symbol.Decimal) > 0 {
					b := []byte(symbol.Decimal[0].Data())
					trans.DecimalLen = len(b)
					trans.Decimal = fmt.Sprintf("%#v", b)
				}
				if len(symbol.Group) > 0 {
					b := []byte(symbol.Group[0].Data())
					trans.GroupLen = len(b)
					trans.Group = fmt.Sprintf("%#v", b)
				}
				if len(symbol.MinusSign) > 0 {
					b := []byte(symbol.MinusSign[0].Data())
					trans.MinusLen = len(b)
					trans.Minus = fmt.Sprintf("%#v", b)
				}
				if len(symbol.PercentSign) > 0 {
					b := []byte(symbol.PercentSign[0].Data())
					trans.PercentLen = len(b)
					trans.Percent = fmt.Sprintf("%#v", b)
				}
				if len(symbol.PerMille) > 0 {
					b := []byte(symbol.PerMille[0].Data())
					trans.PerMilleLen = len(b)
					trans.PerMille = fmt.Sprintf("%#v", b)
				}
			}

			if ldml.Numbers.Currencies != nil {

				for _, currency := range ldml.Numbers.Currencies.Currency {

					if len(strings.TrimSpace(currency.Type)) == 0 {
						continue
					}

					globalCurrenciesMap[currency.Type] = struct{}{}
				}
			}

			if len(ldml.Numbers.DecimalFormats) > 0 && len(ldml.Numbers.DecimalFormats[0].DecimalFormatLength) > 0 {

				for _, dfl := range ldml.Numbers.DecimalFormats[0].DecimalFormatLength {
					if len(dfl.Type) == 0 {
						trans.DecimalNumberFormat = dfl.DecimalFormat[0].Pattern[0].Data()
						break
					}
				}
			}

			if len(ldml.Numbers.PercentFormats) > 0 && len(ldml.Numbers.PercentFormats[0].PercentFormatLength) > 0 {

				for _, dfl := range ldml.Numbers.PercentFormats[0].PercentFormatLength {
					if len(dfl.Type) == 0 {
						trans.PercentNumberFormat = dfl.PercentFormat[0].Pattern[0].Data()
						break
					}
				}
			}

			// var decimalFormat, currencyFormat, currencyAccountingFormat, percentageFormat string

			// if len(ldml.Numbers.DecimalFormats) > 0 && len(ldml.Numbers.DecimalFormats[0].DecimalFormatLength) > 0 {
			// 	decimalFormat = ldml.Numbers.DecimalFormats[0].DecimalFormatLength[0].DecimalFormat[0].Pattern[0].Data()
			// }

			// if len(ldml.Numbers.CurrencyFormats) > 0 && len(ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength) > 0 {

			// 	currencyFormat = ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[0].Pattern[0].Data()
			// 	currencyAccountingFormat = currencyFormat

			// 	if len(ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat) > 1 {
			// 		currencyAccountingFormat = ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[1].Pattern[0].Data()
			// 	}
			// }

			// if len(ldml.Numbers.PercentFormats) > 0 && len(ldml.Numbers.PercentFormats[0].PercentFormatLength) > 0 {
			// 	percentageFormat = ldml.Numbers.PercentFormats[0].PercentFormatLength[0].PercentFormat[0].Pattern[0].Data()
			// }

			// // parse Number values
			// parseNumbers(decimal, group, minus, percent, permille, decimalFormat, currencyFormat, currencyAccountingFormat, percentageFormat)

			// end number values
		}
	}

	for k := range globalCurrenciesMap {
		globalCurrencies = append(globalCurrencies, k)
	}

	sort.Strings(globalCurrencies)

	for i, loc := range globalCurrencies {
		globCurrencyIdxMap[loc] = i
	}
}

func parsePercentNumberFormat(trans *translator) (results string) {

	if len(trans.PercentNumberFormat) == 0 {
		return
	}

	trans.FmtPercentExists = true

	// formats := strings.SplitN(trans.PercentNumberFormat, ";", 2)

	// if len(formats) > 1 {
	// 	trans.FmtNumberHasNegativeFormat = true
	// }

	match := groupLenPercentRegex.FindString(trans.PercentNumberFormat)
	if len(match) > 0 {
		trans.FmtPercentGroupLen = len(match) - 1
	}

	match = requiredDecimalRegex.FindString(trans.PercentNumberFormat)
	if len(match) > 0 {
		trans.FmtPercentMinDecimalLen = len(match) - 1
	}

	match = secondaryGroupLenRegex.FindString(trans.PercentNumberFormat)
	if len(match) > 0 {
		trans.FmtPercentSecondaryGroupLen = len(match) - 2
	}

	// 	FmtPercentPrefix            string
	// FmtPercentSuffix            string
	// FmtPercentInPrefix          bool
	// FmtPercentLeft              bool

	// if formats[0][0] == '%' {
	// 	trans.FmtPercentLeft = true
	// }

	// trans.FmtPercentLeft

	idx := 0

	for idx = 0; idx < len(trans.PercentNumberFormat); idx++ {
		if trans.PercentNumberFormat[idx] == '#' || trans.PercentNumberFormat[idx] == '0' {
			trans.FmtPercentPrefix = trans.PercentNumberFormat[:idx]
			break
		}
	}

	for idx = len(trans.PercentNumberFormat) - 1; idx >= 0; idx-- {
		if trans.PercentNumberFormat[idx] == '#' || trans.PercentNumberFormat[idx] == '0' {
			idx++
			trans.FmtPercentSuffix = trans.PercentNumberFormat[idx:]
			break
		}
	}

	for idx = 0; idx < len(trans.FmtPercentPrefix); idx++ {
		if trans.FmtPercentPrefix[idx] == '%' {
			trans.FmtPercentInPrefix = true

			if idx == 0 {
				trans.FmtPercentLeft = true
				trans.FmtPercentPrefix = trans.FmtPercentPrefix[idx+1:]
			} else {
				trans.FmtPercentLeft = false
				trans.FmtPercentPrefix = trans.FmtPercentPrefix[:idx]
			}

			break
		}
	}

	for idx = 0; idx < len(trans.FmtPercentSuffix); idx++ {
		if trans.FmtPercentSuffix[idx] == '%' {
			trans.FmtPercentInPrefix = false

			if idx == 0 {
				trans.FmtPercentLeft = true
				trans.FmtPercentSuffix = trans.FmtPercentSuffix[idx+1:]
			} else {
				trans.FmtPercentLeft = false
				trans.FmtPercentSuffix = trans.FmtPercentSuffix[:idx]
			}

			break
		}
	}

	// if len(trans.FmtPercentPrefix) == 1 && trans.FmtPercentPrefix[0] == '%' {
	// 	trans.FmtPercentPrefix = ""
	// 	trans.FmtPercentInPrefix = true
	// }

	// if len(trans.FmtPercentSuffix) == 1 && trans.FmtPercentSuffix[0] == '%' {
	// 	trans.FmtPercentSuffix = ""
	// 	trans.FmtPercentSuffix = false
	// }

	// // if start > 0 {
	// // 	prefix = formats[0][:start]
	// // }

	// end := 0

	// // positive prefix
	// for end = len(formats[0]) - 1; end >= 0; end-- {
	// 	if formats[0][end] == '#' || formats[0][end] == '0' {
	// 		end++
	// 		break
	// 	}
	// }

	// fmt.Println(start)
	// fmt.Println(end)

	return
}

func parseDecimalNumberFormat(trans *translator) (results string) {

	if len(trans.DecimalNumberFormat) == 0 {
		return
	}

	trans.FmtNumberExists = true

	formats := strings.SplitN(trans.DecimalNumberFormat, ";", 2)

	// if len(formats) > 1 {
	// 	trans.FmtNumberHasNegativeFormat = true
	// }

	match := groupLenRegex.FindString(formats[0])
	if len(match) > 0 {
		trans.FmtNumberGroupLen = len(match) - 2
	}

	match = requiredDecimalRegex.FindString(formats[0])
	if len(match) > 0 {
		trans.FmtNumberMinDecimalLen = len(match) - 1
	}

	match = secondaryGroupLenRegex.FindString(formats[0])
	if len(match) > 0 {
		trans.FmtNumberSecondaryGroupLen = len(match) - 2
	}

	// start := 0
	// // prefix := ""

	// // positive prefix
	// for start = 0; start < len(formats[0]); start++ {
	// 	if formats[0][start] == '#' || formats[0][start] == '0' {
	// 		break
	// 	}
	// }

	// // if start > 0 {
	// // 	prefix = formats[0][:start]
	// // }

	// end := 0

	// // positive prefix
	// for end = len(formats[0]) - 1; end >= 0; end-- {
	// 	if formats[0][end] == '#' || formats[0][end] == '0' {
	// 		end++
	// 		break
	// 	}
	// }

	// fmt.Println(start)
	// fmt.Println(end)

	return
}

type sortRank struct {
	Rank  uint8
	Value string
}

type ByRank []sortRank

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return a[i].Rank < a[j].Rank }

// TODO: refine generated code a bit, some combinations end up with same plural rule,
// could check all at once; but it works and that's step 1 complete
func parseRangePluralRuleFunc(current *cldr.CLDR, baseLocale string) (results string) {

	var pluralRange *struct {
		cldr.Common
		Locales     string `xml:"locales,attr"`
		PluralRange []*struct {
			cldr.Common
			Start  string `xml:"start,attr"`
			End    string `xml:"end,attr"`
			Result string `xml:"result,attr"`
		} `xml:"pluralRange"`
	}

	for _, pr := range current.Supplemental().Plurals[1].PluralRanges {

		locs := strings.Split(pr.Locales, " ")

		for _, loc := range locs {

			if loc == baseLocale {
				pluralRange = pr
			}
		}
	}

	// no range plural rules for locale
	if pluralRange == nil {
		results = "return locales.PluralRuleUnknown"
		return
	}

	mp := make(map[string]struct{})

	// pre-process if all the same
	for _, rule := range pluralRange.PluralRange {
		mp[rule.Result] = struct{}{}
	}

	if len(mp) == 1 {
		results += "return locales." + pluralStringToString(pluralRange.PluralRange[0].Result)
		return
	}

	multiple := len(pluralRange.PluralRange) > 1

	if multiple {
		results += "start := " + baseLocale + ".CardinalPluralRule(num1, v1)\n"
		results += "end := " + baseLocale + ".CardinalPluralRule(num2, v2)\n\n"
	}

	first := true

	// pre parse for variables
	for i, rule := range pluralRange.PluralRange {

		if i == len(pluralRange.PluralRange)-1 {

			if multiple {
				results += "\n\n"
			}
			results += "return locales." + pluralStringToString(rule.Result)
			continue
		}

		if first {
			results += "if"
			first = false
		} else {
			results += "else if"
		}

		results += " start == locales." + pluralStringToString(rule.Start) + " && end == locales." + pluralStringToString(rule.End) + " {\n return locales." + pluralStringToString(rule.Result) + "\n} "

	}

	if multiple {
		results = "\n" + results + "\n"
	}

	return

}

// TODO: cleanup function logic perhaps write a lexer... but it's working right now, and
// I'm already farther down the rabbit hole than I'd like and so pulling the chute here.
func parseOrdinalPluralRuleFunc(current *cldr.CLDR, baseLocale string) (results string, plurals string) {

	var prOrdinal *struct {
		cldr.Common
		Locales    string "xml:\"locales,attr\""
		PluralRule []*struct {
			cldr.Common
			Count string "xml:\"count,attr\""
		} "xml:\"pluralRule\""
	}

	var pluralArr []locales.PluralRule

	// idx 0 is ordinal rules
	for _, pr := range current.Supplemental().Plurals[0].PluralRules {

		locs := strings.Split(pr.Locales, " ")

		for _, loc := range locs {

			if loc == baseLocale {

				prOrdinal = pr
				// for _, pl := range pr.PluralRule {
				// 	fmt.Println(pl.Count, pl.Common.Data())
				// }
			}
		}
	}

	// no plural rules for locale
	if prOrdinal == nil {
		plurals = "nil"
		results = "return locales.PluralRuleUnknown"
		return
	}

	vals := make(map[string]struct{})
	first := true

	// pre parse for variables
	for _, rule := range prOrdinal.PluralRule {

		ps1 := pluralStringToString(rule.Count)
		psI := pluralStringToInt(rule.Count)
		pluralArr = append(pluralArr, psI)

		data := strings.Replace(strings.Replace(strings.Replace(strings.TrimSpace(strings.SplitN(rule.Common.Data(), "@", 2)[0]), " = ", " == ", -1), " or ", " || ", -1), " and ", " && ", -1)

		if len(data) == 0 {
			if len(prOrdinal.PluralRule) == 1 {

				results = "return locales." + ps1

			} else {

				results += "\n\nreturn locales." + ps1
				// results += "else {\nreturn locales." + locales.PluralStringToString(rule.Count) + ", nil\n}"
			}

			continue
		}

		// // All need n, so always add
		// if strings.Contains(data, "n") {
		// 	vals[prVarFuncs["n"]] = struct{}{}
		// }

		if strings.Contains(data, "i") {
			vals[prVarFuncs["i"]] = struct{}{}
		}

		// v is inherently avaialable as an argument
		// if strings.Contains(data, "v") {
		// 	vals[prVarFuncs["v"]] = struct{}{}
		// }

		if strings.Contains(data, "w") {
			vals[prVarFuncs["w"]] = struct{}{}
		}

		if strings.Contains(data, "f") {
			vals[prVarFuncs["f"]] = struct{}{}
		}

		if strings.Contains(data, "t") {
			vals[prVarFuncs["t"]] = struct{}{}
		}

		if first {
			results += "if "
			first = false
		} else {
			results += "else if "
		}

		stmt := ""

		// real work here
		//
		// split by 'or' then by 'and' allowing to better
		// determine bracketing for formula

		ors := strings.Split(data, "||")

		for _, or := range ors {

			stmt += "("

			ands := strings.Split(strings.TrimSpace(or), "&&")

			for _, and := range ands {

				inArg := false
				pre := ""
				lft := ""
				preOperator := ""
				args := strings.Split(strings.TrimSpace(and), " ")

				for _, a := range args {

					if inArg {
						// check to see if is a value range 2..9

						multiRange := strings.Count(a, "..") > 1
						cargs := strings.Split(strings.TrimSpace(a), ",")
						hasBracket := len(cargs) > 1
						bracketAdded := false
						lastWasRange := false

						for _, carg := range cargs {

							if rng := strings.Split(carg, ".."); len(rng) > 1 {

								if multiRange {
									pre += " ("
								} else {
									pre += " "
								}

								switch preOperator {
								case "==":
									pre += lft + " >= " + rng[0] + " && " + lft + "<=" + rng[1]
								case "!=":
									pre += lft + " < " + rng[0] + " && " + lft + " > " + rng[1]
								}

								if multiRange {
									pre += ") || "
								} else {
									pre += " || "
								}

								lastWasRange = true
								continue
							}

							if lastWasRange {
								pre = strings.TrimRight(pre, " || ") + " && "
							}

							lastWasRange = false

							if hasBracket && !bracketAdded {
								pre += "("
								bracketAdded = true
							}

							// single comma separated values
							switch preOperator {
							case "==":
								pre += " " + lft + preOperator + carg + " || "
							case "!=":
								pre += " " + lft + preOperator + carg + " && "
							}

						}

						pre = strings.TrimRight(pre, " || ")
						pre = strings.TrimRight(pre, " && ")
						pre = strings.TrimRight(pre, " || ")

						if hasBracket && bracketAdded {
							pre += ")"
						}

						continue
					}

					if strings.Contains(a, "=") || a == ">" || a == "<" {
						inArg = true
						preOperator = a
						continue
					}

					lft += a
				}

				stmt += pre + " && "
			}

			stmt = strings.TrimRight(stmt, " && ") + ") || "
		}

		stmt = strings.TrimRight(stmt, " || ")

		results += stmt

		results += " {\n"

		// return plural rule here
		results += "return locales." + ps1 + "\n"

		results += "}"
	}

	pre := "\n"

	// always needed
	vals[prVarFuncs["n"]] = struct{}{}

	sorted := make([]sortRank, 0, len(vals))

	for k := range vals {
		switch k[:1] {
		case "n":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["n"],
				Rank:  1,
			})
		case "i":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["i"],
				Rank:  2,
			})
		case "w":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["w"],
				Rank:  3,
			})
		case "f":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["f"],
				Rank:  4,
			})
		case "t":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["t"],
				Rank:  5,
			})
		}
	}

	sort.Sort(ByRank(sorted))

	for _, k := range sorted {
		pre += k.Value
	}

	if len(results) == 0 {
		results = "return locales.PluralRuleUnknown"
	} else {

		if !strings.HasPrefix(results, "return") {

			results = manyToSingleVars(results)
			// pre += "\n"
			results = pre + results
		}
	}

	if len(pluralArr) == 0 {
		plurals = "nil"
	} else {
		plurals = fmt.Sprintf("%#v", pluralArr)
	}

	return
}

// TODO: cleanup function logic perhaps write a lexer... but it's working right now, and
// I'm already farther down the rabbit hole than I'd like and so pulling the chute here.
func parseCardinalPluralRuleFunc(current *cldr.CLDR, baseLocale string) (results string, plurals string) {

	var prCardinal *struct {
		cldr.Common
		Locales    string "xml:\"locales,attr\""
		PluralRule []*struct {
			cldr.Common
			Count string "xml:\"count,attr\""
		} "xml:\"pluralRule\""
	}

	var pluralArr []locales.PluralRule

	// idx 2 is cardinal rules
	for _, pr := range current.Supplemental().Plurals[2].PluralRules {

		locs := strings.Split(pr.Locales, " ")

		for _, loc := range locs {

			if loc == baseLocale {
				prCardinal = pr
			}
		}
	}

	// no plural rules for locale
	if prCardinal == nil {
		plurals = "nil"
		results = "return locales.PluralRuleUnknown"
		return
	}

	vals := make(map[string]struct{})
	first := true

	// pre parse for variables
	for _, rule := range prCardinal.PluralRule {

		ps1 := pluralStringToString(rule.Count)
		psI := pluralStringToInt(rule.Count)
		pluralArr = append(pluralArr, psI)

		data := strings.Replace(strings.Replace(strings.Replace(strings.TrimSpace(strings.SplitN(rule.Common.Data(), "@", 2)[0]), " = ", " == ", -1), " or ", " || ", -1), " and ", " && ", -1)

		if len(data) == 0 {
			if len(prCardinal.PluralRule) == 1 {

				results = "return locales." + ps1

			} else {

				results += "\n\nreturn locales." + ps1
				// results += "else {\nreturn locales." + locales.PluralStringToString(rule.Count) + ", nil\n}"
			}

			continue
		}

		// // All need n, so always add
		// if strings.Contains(data, "n") {
		// 	vals[prVarFuncs["n"]] = struct{}{}
		// }

		if strings.Contains(data, "i") {
			vals[prVarFuncs["i"]] = struct{}{}
		}

		// v is inherently avaialable as an argument
		// if strings.Contains(data, "v") {
		// 	vals[prVarFuncs["v"]] = struct{}{}
		// }

		if strings.Contains(data, "w") {
			vals[prVarFuncs["w"]] = struct{}{}
		}

		if strings.Contains(data, "f") {
			vals[prVarFuncs["f"]] = struct{}{}
		}

		if strings.Contains(data, "t") {
			vals[prVarFuncs["t"]] = struct{}{}
		}

		if first {
			results += "if "
			first = false
		} else {
			results += "else if "
		}

		stmt := ""

		// real work here
		//
		// split by 'or' then by 'and' allowing to better
		// determine bracketing for formula

		ors := strings.Split(data, "||")

		for _, or := range ors {

			stmt += "("

			ands := strings.Split(strings.TrimSpace(or), "&&")

			for _, and := range ands {

				inArg := false
				pre := ""
				lft := ""
				preOperator := ""
				args := strings.Split(strings.TrimSpace(and), " ")

				for _, a := range args {

					if inArg {
						// check to see if is a value range 2..9

						multiRange := strings.Count(a, "..") > 1
						cargs := strings.Split(strings.TrimSpace(a), ",")
						hasBracket := len(cargs) > 1
						bracketAdded := false
						lastWasRange := false

						for _, carg := range cargs {

							if rng := strings.Split(carg, ".."); len(rng) > 1 {

								if multiRange {
									pre += " ("
								} else {
									pre += " "
								}

								switch preOperator {
								case "==":
									pre += lft + " >= " + rng[0] + " && " + lft + "<=" + rng[1]
								case "!=":
									pre += lft + " < " + rng[0] + " && " + lft + " > " + rng[1]
								}

								if multiRange {
									pre += ") || "
								} else {
									pre += " || "
								}

								lastWasRange = true
								continue
							}

							if lastWasRange {
								pre = strings.TrimRight(pre, " || ") + " && "
							}

							lastWasRange = false

							if hasBracket && !bracketAdded {
								pre += "("
								bracketAdded = true
							}

							// single comma separated values
							switch preOperator {
							case "==":
								pre += " " + lft + preOperator + carg + " || "
							case "!=":
								pre += " " + lft + preOperator + carg + " && "
							}

						}

						pre = strings.TrimRight(pre, " || ")
						pre = strings.TrimRight(pre, " && ")
						pre = strings.TrimRight(pre, " || ")

						if hasBracket && bracketAdded {
							pre += ")"
						}

						continue
					}

					if strings.Contains(a, "=") || a == ">" || a == "<" {
						inArg = true
						preOperator = a
						continue
					}

					lft += a
				}

				stmt += pre + " && "
			}

			stmt = strings.TrimRight(stmt, " && ") + ") || "
		}

		stmt = strings.TrimRight(stmt, " || ")

		results += stmt

		results += " {\n"

		// return plural rule here
		results += "return locales." + ps1 + "\n"

		results += "}"
	}

	pre := "\n"

	// always needed
	vals[prVarFuncs["n"]] = struct{}{}

	sorted := make([]sortRank, 0, len(vals))

	for k := range vals {
		switch k[:1] {
		case "n":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["n"],
				Rank:  1,
			})
		case "i":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["i"],
				Rank:  2,
			})
		case "w":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["w"],
				Rank:  3,
			})
		case "f":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["f"],
				Rank:  4,
			})
		case "t":
			sorted = append(sorted, sortRank{
				Value: prVarFuncs["t"],
				Rank:  5,
			})
		}
	}

	sort.Sort(ByRank(sorted))

	for _, k := range sorted {
		pre += k.Value
	}

	if len(results) == 0 {
		results = "return locales.PluralRuleUnknown"
	} else {

		if !strings.HasPrefix(results, "return") {

			results = manyToSingleVars(results)
			// pre += "\n"
			results = pre + results
		}
	}

	if len(pluralArr) == 0 {
		plurals = "nil"
	} else {
		plurals = fmt.Sprintf("%#v", pluralArr)
	}

	return
}

func manyToSingleVars(input string) (results string) {

	matches := nModRegex.FindAllString(input, -1)
	mp := make(map[string][]string) // map of formula to variable
	var found bool
	var split []string
	var variable string

	for _, formula := range matches {

		if _, found = mp[formula]; found {
			continue
		}

		split = strings.SplitN(formula, "%", 2)

		mp[formula] = []string{split[1], "math.Mod(" + split[0] + ", " + split[1] + ")"}
	}

	for k, v := range mp {
		variable = "nMod" + v[0]
		results += variable + " := " + v[1] + "\n"
		input = strings.Replace(input, k, variable, -1)
	}

	matches = iModRegex.FindAllString(input, -1)
	mp = make(map[string][]string) // map of formula to variable

	for _, formula := range matches {

		if _, found = mp[formula]; found {
			continue
		}

		split = strings.SplitN(formula, "%", 2)

		mp[formula] = []string{split[1], formula}
	}

	for k, v := range mp {
		variable = "iMod" + v[0]
		results += variable + " := " + v[1] + "\n"
		input = strings.Replace(input, k, variable, -1)
	}

	matches = wModRegex.FindAllString(input, -1)
	mp = make(map[string][]string) // map of formula to variable

	for _, formula := range matches {

		if _, found = mp[formula]; found {
			continue
		}

		split = strings.SplitN(formula, "%", 2)

		mp[formula] = []string{split[1], formula}
	}

	for k, v := range mp {
		variable = "wMod" + v[0]
		results += variable + " := " + v[1] + "\n"
		input = strings.Replace(input, k, variable, -1)
	}

	matches = fModRegex.FindAllString(input, -1)
	mp = make(map[string][]string) // map of formula to variable

	for _, formula := range matches {

		if _, found = mp[formula]; found {
			continue
		}

		split = strings.SplitN(formula, "%", 2)

		mp[formula] = []string{split[1], formula}
	}

	for k, v := range mp {
		variable = "fMod" + v[0]
		results += variable + " := " + v[1] + "\n"
		input = strings.Replace(input, k, variable, -1)
	}

	matches = tModRegex.FindAllString(input, -1)
	mp = make(map[string][]string) // map of formula to variable

	for _, formula := range matches {

		if _, found = mp[formula]; found {
			continue
		}

		split = strings.SplitN(formula, "%", 2)

		mp[formula] = []string{split[1], formula}
	}

	for k, v := range mp {
		variable = "tMod" + v[0]
		results += variable + " := " + v[1] + "\n"
		input = strings.Replace(input, k, variable, -1)
	}

	results = results + "\n" + input

	return
}

// pluralStringToInt returns the enum value of 'plural' provided
func pluralStringToInt(plural string) locales.PluralRule {

	switch plural {
	case "zero":
		return locales.PluralRuleZero
	case "one":
		return locales.PluralRuleOne
	case "two":
		return locales.PluralRuleTwo
	case "few":
		return locales.PluralRuleFew
	case "many":
		return locales.PluralRuleMany
	case "other":
		return locales.PluralRuleOther
	default:
		return locales.PluralRuleUnknown
	}
}

func pluralStringToString(pr string) string {

	pr = strings.TrimSpace(pr)

	switch pr {
	case "zero":
		return "PluralRuleZero"
	case "one":
		return "PluralRuleOne"
	case "two":
		return "PluralRuleTwo"
	case "few":
		return "PluralRuleFew"
	case "many":
		return "PluralRuleMany"
	case "other":
		return "PluralRuleOther"
	default:
		return "PluralRuleUnknown"
	}
}

// //plurals
// rules := "data/rules"
// plurals := map[string]*pluralInfo{}
// basePlurals := map[string]string{}

// err := filepath.Walk(rules, func(path string, info os.FileInfo, err error) error {

// 	if err != nil {
// 		panic(err)
// 	}

// 	if info.IsDir() {
// 		return nil
// 	}

// 	in, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var out yaml.MapSlice
// 	if err = yaml.Unmarshal(in, &out); err != nil {
// 		panic(err)
// 	}

// 	var plural string
// 	for _, item := range out {
// 		if item.Key == "plural" {
// 			plural = item.Value.(string)
// 		}
// 	}

// 	locale := strings.Replace(info.Name(), filepath.Ext(info.Name()), "", 1)
// 	locale = strings.ToLower(strings.Replace(locale, "-", "_", -1))

// 	plurals[locale] = &pluralInfo{
// 		path:   path,
// 		locale: locale,
// 		plural: plural,
// 	}

// 	if plural == "" {
// 		return nil
// 	}

// 	basePlurals[locale] = plural

// 	return nil
// })

// if err != nil {
// 	panic(err)
// }

// for _, p := range plurals {

// 	if p.plural == "" {

// 		var ok bool

// 		fmt.Print("can't find plurals in ", p.path, " attempting to locate base language plural rules...")

// 		base := strings.SplitN(p.locale, "_", 2)

// 		p.plural, ok = basePlurals[base[0]]
// 		if !ok {
// 			fmt.Println("Not Found")
// 			continue
// 		}

// 		fmt.Println("Found")
// 	}
// }

// cldr

// var decoder cldr.Decoder
// cldr, err := decoder.DecodePath("data/core")
// if err != nil {
// 	panic(err)
// }

// 	locs := map[string]string{}
// 	numbers := map[string]i18n.Number{}
// 	calendars := map[string]i18n.Calendar{}
// 	locales := cldr.Locales()
// 	for _, loc := range locales {

// 		ldml := cldr.RawLDML(loc)
// 		if ldml.Numbers == nil {
// 			continue
// 		}
// 		var number i18n.Number

// 		number.Currencies = make(i18n.CurrencyFormatValue)

// 		if len(ldml.Numbers.Symbols) > 0 {
// 			symbol := ldml.Numbers.Symbols[0]
// 			if len(symbol.Decimal) > 0 {
// 				number.Symbols.Decimal = symbol.Decimal[0].Data()
// 			}
// 			if len(symbol.Group) > 0 {
// 				number.Symbols.Group = symbol.Group[0].Data()
// 			}
// 			if len(symbol.MinusSign) > 0 {
// 				number.Symbols.Negative = symbol.MinusSign[0].Data()
// 			}
// 			if len(symbol.PercentSign) > 0 {
// 				number.Symbols.Percent = symbol.PercentSign[0].Data()
// 			}
// 			if len(symbol.PerMille) > 0 {
// 				number.Symbols.PerMille = symbol.PerMille[0].Data()
// 			}
// 		}
// 		if len(ldml.Numbers.DecimalFormats) > 0 && len(ldml.Numbers.DecimalFormats[0].DecimalFormatLength) > 0 {
// 			number.Formats.Decimal = ldml.Numbers.DecimalFormats[0].DecimalFormatLength[0].DecimalFormat[0].Pattern[0].Data()
// 		}

// 		if len(ldml.Numbers.CurrencyFormats) > 0 && len(ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength) > 0 {

// 			number.Formats.Currency = ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[0].Pattern[0].Data()
// 			number.Formats.CurrencyAccounting = number.Formats.Currency

// 			if len(ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat) > 1 {
// 				number.Formats.CurrencyAccounting = ldml.Numbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[1].Pattern[0].Data()
// 			}
// 		}
// 		if len(ldml.Numbers.PercentFormats) > 0 && len(ldml.Numbers.PercentFormats[0].PercentFormatLength) > 0 {
// 			number.Formats.Percent = ldml.Numbers.PercentFormats[0].PercentFormatLength[0].PercentFormat[0].Pattern[0].Data()
// 		}
// 		if ldml.Numbers.Currencies != nil {

// 			for _, currency := range ldml.Numbers.Currencies.Currency {

// 				var c i18n.Currency

// 				c.Currency = currency.Type

// 				if len(currency.DisplayName) > 0 {
// 					c.DisplayName = currency.DisplayName[0].Data()
// 				}

// 				if len(currency.Symbol) > 0 {
// 					c.Symbol = currency.Symbol[0].Data()
// 				}

// 				number.Currencies[c.Currency] = c
// 			}
// 		}
// 		numbers[loc] = number
// 		locs[loc] = strings.ToLower(strings.Replace(loc, "_", "", -1))

// 		if ldml.Dates != nil && ldml.Dates.Calendars != nil {

// 			var calendar i18n.Calendar
// 			ldmlCar := ldml.Dates.Calendars.Calendar[0]

// 			for _, cal := range ldml.Dates.Calendars.Calendar {
// 				if cal.Type == "gregorian" {
// 					ldmlCar = cal
// 				}
// 			}

// 			if ldmlCar.DateFormats != nil {
// 				for _, datefmt := range ldmlCar.DateFormats.DateFormatLength {
// 					switch datefmt.Type {
// 					case "full":
// 						calendar.Formats.DateEra.BC.Full = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Full = calendar.Formats.DateEra.BC.Full

// 					case "long":
// 						calendar.Formats.DateEra.BC.Long = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Long = calendar.Formats.DateEra.BC.Long

// 						// Overrides TODO: Split Each Section into separate function, getting to big to maintain

// 					case "medium":
// 						calendar.Formats.DateEra.BC.Medium = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Medium = calendar.Formats.DateEra.BC.Medium
// 					case "short":
// 						calendar.Formats.DateEra.BC.Short = datefmt.DateFormat[0].Pattern[0].Data()
// 						calendar.Formats.DateEra.AD.Short = calendar.Formats.DateEra.BC.Short
// 					}
// 				}

// 				// Overrides TODO: Split Each Section into separate function, getting to big to maintain
// 				switch loc {
// 				case "th":
// 					calendar.Formats.DateEra.BC.Full = "EEEEที่ d MMMM y GGGG"
// 					calendar.Formats.DateEra.AD.Full = "EEEEที่ d MMMM GGGG y"
// 					calendar.Formats.DateEra.BC.Long = "d MMMM y GG"
// 					calendar.Formats.DateEra.AD.Long = "d MMMM GG y"
// 				case "en":
// 					calendar.Formats.DateEra.BC.Full = "EEEE, MMMM d, y GGGG"
// 					calendar.Formats.DateEra.BC.Long = "MMMM d, y GG"
// 					// calendar.Formats.DateEra.BC.Medium = "MMM d, y GG"
// 					// calendar.Formats.DateEra.BC.Short = "M/d/yy G"
// 				}
// 			}

// 			if ldmlCar.TimeFormats != nil {
// 				for _, datefmt := range ldmlCar.TimeFormats.TimeFormatLength {
// 					switch datefmt.Type {
// 					case "full":
// 						calendar.Formats.Time.Full = datefmt.TimeFormat[0].Pattern[0].Data()
// 					case "long":
// 						calendar.Formats.Time.Long = datefmt.TimeFormat[0].Pattern[0].Data()
// 					case "medium":
// 						calendar.Formats.Time.Medium = datefmt.TimeFormat[0].Pattern[0].Data()
// 					case "short":
// 						calendar.Formats.Time.Short = datefmt.TimeFormat[0].Pattern[0].Data()
// 					}
// 				}
// 			}
// 			if ldmlCar.DateTimeFormats != nil {
// 				for _, datefmt := range ldmlCar.DateTimeFormats.DateTimeFormatLength {
// 					switch datefmt.Type {
// 					case "full":
// 						calendar.Formats.DateTime.Full = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					case "long":
// 						calendar.Formats.DateTime.Long = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					case "medium":
// 						calendar.Formats.DateTime.Medium = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					case "short":
// 						calendar.Formats.DateTime.Short = datefmt.DateTimeFormat[0].Pattern[0].Data()
// 					}
// 				}
// 			}
// 			if ldmlCar.Months != nil {

// 				for _, monthctx := range ldmlCar.Months.MonthContext {

// 					for _, months := range monthctx.MonthWidth {

// 						i18nMonth := make(i18n.CalendarMonthFormatNameValue)

// 						for _, m := range months.Month {
// 							switch m.Type {
// 							case "1":
// 								i18nMonth[time.January] = m.Data()
// 							case "2":
// 								i18nMonth[time.February] = m.Data()
// 							case "3":
// 								i18nMonth[time.March] = m.Data()
// 							case "4":
// 								i18nMonth[time.April] = m.Data()
// 							case "5":
// 								i18nMonth[time.May] = m.Data()
// 							case "6":
// 								i18nMonth[time.June] = m.Data()
// 							case "7":
// 								i18nMonth[time.July] = m.Data()
// 							case "8":
// 								i18nMonth[time.August] = m.Data()
// 							case "9":
// 								i18nMonth[time.September] = m.Data()
// 							case "10":
// 								i18nMonth[time.October] = m.Data()
// 							case "11":
// 								i18nMonth[time.November] = m.Data()
// 							case "12":
// 								i18nMonth[time.December] = m.Data()
// 							}
// 						}
// 						switch months.Type {
// 						case "abbreviated":
// 							calendar.FormatNames.Months.Abbreviated = i18nMonth
// 						case "narrow":
// 							calendar.FormatNames.Months.Narrow = i18nMonth
// 						case "short":
// 							calendar.FormatNames.Months.Short = i18nMonth
// 						case "wide":
// 							calendar.FormatNames.Months.Wide = i18nMonth
// 						}
// 					}
// 				}
// 			}
// 			if ldmlCar.Days != nil {
// 				for _, dayctx := range ldmlCar.Days.DayContext {

// 					for _, days := range dayctx.DayWidth {

// 						i18nDay := make(i18n.CalendarDayFormatNameValue)

// 						for _, d := range days.Day {

// 							switch d.Type {
// 							case "sun":
// 								i18nDay[time.Sunday] = d.Data()
// 							case "mon":
// 								i18nDay[time.Monday] = d.Data()
// 							case "tue":
// 								i18nDay[time.Tuesday] = d.Data()
// 							case "wed":
// 								i18nDay[time.Wednesday] = d.Data()
// 							case "thu":
// 								i18nDay[time.Thursday] = d.Data()
// 							case "fri":
// 								i18nDay[time.Friday] = d.Data()
// 							case "sat":
// 								i18nDay[time.Saturday] = d.Data()
// 							}
// 						}

// 						switch days.Type {
// 						case "abbreviated":
// 							calendar.FormatNames.Days.Abbreviated = i18nDay
// 						case "narrow":
// 							calendar.FormatNames.Days.Narrow = i18nDay
// 						case "short":
// 							calendar.FormatNames.Days.Short = i18nDay
// 						case "wide":
// 							calendar.FormatNames.Days.Wide = i18nDay
// 						}
// 					}
// 				}
// 			}

// 			if ldmlCar.DayPeriods != nil {

// 				for _, ctx := range ldmlCar.DayPeriods.DayPeriodContext {

// 					for _, width := range ctx.DayPeriodWidth {

// 						// var i18nPeriod i18n.CalendarPeriodFormatNameValue
// 						i18nPeriod := make(i18n.CalendarPeriodFormatNameValue)

// 						for _, d := range width.DayPeriod {

// 							if _, ok := i18nPeriod[d.Type]; !ok {
// 								i18nPeriod[d.Type] = d.Data()
// 							}
// 						}

// 						switch width.Type {
// 						case "abbreviated":
// 							calendar.FormatNames.Periods.Abbreviated = i18nPeriod
// 						case "narrow":
// 							calendar.FormatNames.Periods.Narrow = i18nPeriod
// 						case "short":
// 							calendar.FormatNames.Periods.Short = i18nPeriod
// 						case "wide":
// 							calendar.FormatNames.Periods.Wide = i18nPeriod
// 						}
// 					}
// 				}
// 				// var empty i18n.CalendarPeriodFormatNameValue
// 				// if calendar.FormatNames.Periods.Abbreviated == empty {
// 				// 	calendar.FormatNames.Periods.Abbreviated = calendar.FormatNames.Periods.Wide
// 				// }
// 			}

// 			if ldmlCar.Eras != nil {

// 				var eras i18n.Eras

// 				if ldmlCar.Eras.EraNames != nil && len(ldmlCar.Eras.EraNames.Era) > 0 {
// 					eras.BC.Full = ldmlCar.Eras.EraNames.Era[0].Data()
// 				}

// 				if ldmlCar.Eras.EraAbbr != nil && len(ldmlCar.Eras.EraAbbr.Era) > 0 {
// 					eras.BC.Abbrev = ldmlCar.Eras.EraAbbr.Era[0].Data()
// 				}

// 				if ldmlCar.Eras.EraNarrow != nil && len(ldmlCar.Eras.EraNarrow.Era) > 0 {
// 					eras.BC.Narrow = ldmlCar.Eras.EraNarrow.Era[0].Data()
// 				}

// 				if ldmlCar.Eras.EraNames != nil && len(ldmlCar.Eras.EraNames.Era) > 2 {
// 					eras.AD.Full = ldmlCar.Eras.EraNames.Era[2].Data()
// 				}

// 				if ldmlCar.Eras.EraAbbr != nil && len(ldmlCar.Eras.EraAbbr.Era) > 2 {
// 					eras.AD.Abbrev = ldmlCar.Eras.EraAbbr.Era[2].Data()
// 				}

// 				if ldmlCar.Eras.EraNarrow != nil && len(ldmlCar.Eras.EraNarrow.Era) > 2 {
// 					eras.AD.Narrow = ldmlCar.Eras.EraNarrow.Era[2].Data()
// 				}

// 				calendar.FormatNames.Eras = eras
// 			}

// 			calendars[loc] = calendar
// 		}
// 	}

// 	for locale := range locs {

// 		if !strings.Contains(locale, "_") {
// 			continue
// 		}

// 		calendar := calendars[locale]

// 		bString := strings.SplitN(locale, "_", 2)
// 		base := bString[0]

// 		baseCal := calendars[base]

// 		// copy base calendar objects

// 		// Dates
// 		if calendar.Formats.DateEra.AD.Full == "" {
// 			calendar.Formats.DateEra.BC.Full = baseCal.Formats.DateEra.BC.Full
// 			calendar.Formats.DateEra.AD.Full = baseCal.Formats.DateEra.AD.Full
// 		}

// 		if calendar.Formats.DateEra.AD.Long == "" {
// 			calendar.Formats.DateEra.BC.Long = baseCal.Formats.DateEra.BC.Long
// 			calendar.Formats.DateEra.AD.Long = baseCal.Formats.DateEra.AD.Long
// 		}

// 		if calendar.Formats.DateEra.AD.Medium == "" {
// 			calendar.Formats.DateEra.BC.Medium = baseCal.Formats.DateEra.BC.Medium
// 			calendar.Formats.DateEra.AD.Medium = baseCal.Formats.DateEra.AD.Medium
// 		}

// 		if calendar.Formats.DateEra.AD.Short == "" {
// 			calendar.Formats.DateEra.BC.Short = baseCal.Formats.DateEra.BC.Short
// 			calendar.Formats.DateEra.AD.Short = baseCal.Formats.DateEra.AD.Short
// 		}

// 		// times
// 		if calendar.Formats.Time.Full == "" {
// 			calendar.Formats.Time.Full = baseCal.Formats.Time.Full
// 		}

// 		if calendar.Formats.Time.Long == "" {
// 			calendar.Formats.Time.Long = baseCal.Formats.Time.Long
// 		}

// 		if calendar.Formats.Time.Medium == "" {
// 			calendar.Formats.Time.Medium = baseCal.Formats.Time.Medium
// 		}

// 		if calendar.Formats.Time.Short == "" {
// 			calendar.Formats.Time.Short = baseCal.Formats.Time.Short
// 		}

// 		// date & times
// 		if calendar.Formats.DateTime.Full == "" {
// 			calendar.Formats.DateTime.Full = baseCal.Formats.DateTime.Full
// 		}

// 		if calendar.Formats.DateTime.Long == "" {
// 			calendar.Formats.DateTime.Long = baseCal.Formats.DateTime.Long
// 		}

// 		if calendar.Formats.DateTime.Medium == "" {
// 			calendar.Formats.DateTime.Medium = baseCal.Formats.DateTime.Medium
// 		}

// 		if calendar.Formats.DateTime.Short == "" {
// 			calendar.Formats.DateTime.Short = baseCal.Formats.DateTime.Short
// 		}

// 		// months

// 		if calendar.FormatNames.Months.Abbreviated == nil {
// 			calendar.FormatNames.Months.Abbreviated = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		if calendar.FormatNames.Months.Narrow == nil {
// 			calendar.FormatNames.Months.Narrow = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		if calendar.FormatNames.Months.Short == nil {
// 			calendar.FormatNames.Months.Short = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		if calendar.FormatNames.Months.Wide == nil {
// 			calendar.FormatNames.Months.Wide = make(i18n.CalendarMonthFormatNameValue)
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Abbreviated {

// 			val, ok := calendar.FormatNames.Months.Abbreviated[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Abbreviated[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Abbreviated[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Narrow {

// 			val, ok := calendar.FormatNames.Months.Narrow[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Narrow[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Narrow[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Short {

// 			val, ok := calendar.FormatNames.Months.Short[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Short[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Short[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Months.Wide {

// 			val, ok := calendar.FormatNames.Months.Wide[k]
// 			if !ok {
// 				calendar.FormatNames.Months.Wide[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Months.Wide[k] = v
// 			}
// 		}

// 		// days

// 		if calendar.FormatNames.Days.Abbreviated == nil {
// 			calendar.FormatNames.Days.Abbreviated = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		if calendar.FormatNames.Days.Narrow == nil {
// 			calendar.FormatNames.Days.Narrow = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		if calendar.FormatNames.Days.Short == nil {
// 			calendar.FormatNames.Days.Short = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		if calendar.FormatNames.Days.Wide == nil {
// 			calendar.FormatNames.Days.Wide = make(i18n.CalendarDayFormatNameValue)
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Abbreviated {

// 			val, ok := calendar.FormatNames.Days.Abbreviated[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Abbreviated[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Abbreviated[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Narrow {

// 			val, ok := calendar.FormatNames.Days.Narrow[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Narrow[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Narrow[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Short {

// 			val, ok := calendar.FormatNames.Days.Short[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Short[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Short[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Days.Wide {

// 			val, ok := calendar.FormatNames.Days.Wide[k]
// 			if !ok {
// 				calendar.FormatNames.Days.Wide[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Days.Wide[k] = v
// 			}
// 		}

// 		// periods
// 		if calendar.FormatNames.Periods.Abbreviated == nil {
// 			calendar.FormatNames.Periods.Abbreviated = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		if calendar.FormatNames.Periods.Narrow == nil {
// 			calendar.FormatNames.Periods.Narrow = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		if calendar.FormatNames.Periods.Short == nil {
// 			calendar.FormatNames.Periods.Short = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		if calendar.FormatNames.Periods.Wide == nil {
// 			calendar.FormatNames.Periods.Wide = make(i18n.CalendarPeriodFormatNameValue)
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Abbreviated {

// 			val, ok := calendar.FormatNames.Periods.Abbreviated[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Abbreviated[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Abbreviated[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Narrow {

// 			val, ok := calendar.FormatNames.Periods.Narrow[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Narrow[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Narrow[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Short {

// 			val, ok := calendar.FormatNames.Periods.Short[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Short[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Short[k] = v
// 			}
// 		}

// 		for k, v := range baseCal.FormatNames.Periods.Wide {

// 			val, ok := calendar.FormatNames.Periods.Wide[k]
// 			if !ok {
// 				calendar.FormatNames.Periods.Wide[k] = v
// 				continue
// 			}

// 			if val == "" {
// 				calendar.FormatNames.Periods.Wide[k] = v
// 			}
// 		}

// 		calendars[locale] = calendar

// 		number := numbers[locale]
// 		baseNum := numbers[base]

// 		// symbols
// 		if number.Symbols.Decimal == "" {
// 			number.Symbols.Decimal = baseNum.Symbols.Decimal
// 		}

// 		if number.Symbols.Group == "" {
// 			number.Symbols.Group = baseNum.Symbols.Group
// 		}

// 		if number.Symbols.Negative == "" {
// 			number.Symbols.Negative = baseNum.Symbols.Negative
// 		}

// 		if number.Symbols.Percent == "" {
// 			number.Symbols.Percent = baseNum.Symbols.Percent
// 		}

// 		if number.Symbols.PerMille == "" {
// 			number.Symbols.PerMille = baseNum.Symbols.PerMille
// 		}

// 		// formats
// 		if number.Formats.Decimal == "" {
// 			number.Formats.Decimal = baseNum.Formats.Decimal
// 		}

// 		if number.Formats.Currency == "" {
// 			number.Formats.Currency = baseNum.Formats.Currency
// 		}

// 		if number.Formats.CurrencyAccounting == "" {
// 			number.Formats.CurrencyAccounting = baseNum.Formats.CurrencyAccounting
// 		}

// 		if number.Formats.Percent == "" {
// 			number.Formats.Percent = baseNum.Formats.Percent
// 		}

// 		// currency
// 		for k, v := range baseNum.Currencies {

// 			val, ok := number.Currencies[k]
// 			if !ok {
// 				// number.Currencies[k] = v
// 				continue
// 			}

// 			if val.Currency == "" {
// 				val.Currency = v.Currency
// 			}

// 			if val.DisplayName == "" {
// 				val.DisplayName = v.DisplayName
// 			}

// 			if val.Symbol == "" {
// 				val.Symbol = v.Symbol
// 			}

// 			number.Currencies[k] = val
// 		}

// 		numbers[locale] = number
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(len(numbers))
// 	for locale, number := range numbers {
// 		go func(locale string, number i18n.Number) {

// 			localeLowercase := strings.ToLower(locale)

// 			defer func() { wg.Done() }()
// 			path := "../../resources/locales/" + locale

// 			if _, err := os.Stat(path); err != nil {
// 				if err = os.MkdirAll(path, 0777); err != nil {
// 					panic(err)
// 				}
// 			}

// 			path += "/"

// 			mainFile, err := os.Create(path + "main.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer mainFile.Close()

// 			calendar := calendars[locale]

// 			mainCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"

// 			var locale = &ut.Locale{
// 				Locale: %q,
// 				Number: ut.Number{
// 					Symbols: symbols,
// 					Formats: formats,
// 					Currencies: currencies,
// 				},
// 				Calendar: calendar,
// 				PluralRule:   pluralRule,
// 			}

// 			func init() {
// 				ut.RegisterLocale(locale)
// 			}
// 		`, locale, locale)))

// 			if err != nil {
// 				panic(err)
// 			}

// 			fmt.Fprintf(mainFile, "%s", mainCodes)

// 			numberFile, err := os.Create(path + "number.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer numberFile.Close()

// 			numberCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"
// 			var (
// 				symbols = %#v
// 				formats = %#v
// 			)
// 		`, locale, number.Symbols, number.Formats)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(numberFile, "%s", numberCodes)

// 			currencyFile, err := os.Create(path + "currency.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer currencyFile.Close()

// 			currencyCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"
// 			var currencies = %# v
// 		`, locale, number.Currencies)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(currencyFile, "%s", currencyCodes)

// 			calendarFile, err := os.Create(path + "calendar.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer calendarFile.Close()

// 			calendarCodes, err := format.Source([]byte(fmt.Sprintf(`package %s
// 			import "github.com/go-playground/universal-translator"
// 			var calendar = %#v
// 		`, locale, calendar)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(calendarFile, "%s", calendarCodes)

// 			var ok bool
// 			pluralCode := "1"

// 			pInfo, ok := plurals[localeLowercase]
// 			if ok && pInfo.plural != "" {
// 				pluralCode = pInfo.plural
// 			}

// 			pluralFile, err := os.Create(path + "plural.go")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer pluralFile.Close()

// 			pluralCodes, err := format.Source([]byte(fmt.Sprintf(`package %s

// 			var pluralRule = %q
// 		`, locale, pluralCode)))
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(pluralFile, "%s", pluralCodes)
// 		}(locale, number)
// 	}

// 	wg.Wait()

// 	localesFile, err := os.Create("../../resources/locales/all.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer localesFile.Close()

// 	tmpl, err := template.New("").Parse(`package locales

// 		// Imports for all locales
// 		import (
// 			{{range $locale, $_ := .}}// Locale "{{$locale}}" import that automatically registers itslef with the universal-translator package
// 			_ "github.com/go-playground/universal-translator/resources/locales/{{$locale}}"
// 		{{end}})
// 	`)

// 	if err != nil {
// 		panic(err)
// 	}
// 	var buf bytes.Buffer
// 	if err := tmpl.Execute(&buf, locs); err != nil {
// 		panic(err)
// 	}
// 	allCodes, err := format.Source(buf.Bytes())
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = localesFile.Write(allCodes)
// 	if err != nil {
// 		panic(err)
// 	}
// }
