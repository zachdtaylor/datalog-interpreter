package parser

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

type MatchParameterType func(dp *DatalogParser, predicate *Predicate)

type DatalogParser struct {
	program   *DatalogProgram
	fileName  string
	file      *os.File
	tokenizer DatalogTokenizer
}

func (dp *DatalogParser) Init(program *DatalogProgram, fileName string) {
	dp.program = program
	dp.fileName = fileName
	dp.file = openFile(dp.fileName)

	var tokenizer DatalogTokenizer
	tokenizer.Init(dp.file)
	dp.tokenizer = tokenizer
}

func (dp *DatalogParser) Run() {
	defer dp.file.Close()
	dp.tokenizer.Next()

	dp.Match(SCHEMES)
	dp.Match(COLON)
	dp.MatchScheme() // There must be at least 1 scheme
	dp.MatchSchemeList()
	dp.Match(FACTS)
	dp.Match(COLON)
	dp.MatchFact()
	dp.MatchFactList()
	dp.Match(RULES)
	dp.Match(COLON)
	dp.MatchRuleList()
	dp.Match(QUERIES)
	dp.Match(COLON)
	dp.MatchQuery()
	dp.MatchQueryList()

	fmt.Println("Parse successful!")
	fmt.Println(fmt.Sprintf("  %d schemes", len(dp.program.schemes)))
	fmt.Println(fmt.Sprintf("  %d facts", len(dp.program.facts)))
	fmt.Println(fmt.Sprintf("  %d rules", len(dp.program.rules)))
	fmt.Println(fmt.Sprintf("  %d queries", len(dp.program.queries)))
	fmt.Println(fmt.Sprintf("  Domain: %v", dp.program.domain.Array()))
}

func (dp *DatalogParser) MatchScheme() {
	addScheme(dp.program, dp.MatchPredicate(matchID))
}

func (dp *DatalogParser) MatchSchemeList() {
	if dp.tokenizer.Current().tokenType == ID {
		dp.MatchScheme()
		dp.MatchSchemeList()
	}
}

func (dp *DatalogParser) MatchFact() {
	matchFactString := func(dp *DatalogParser, predicate *Predicate) {
		if dp.tokenizer.Current().tokenType == STRING {
			dp.Match(STRING)
			matchedValue := dp.tokenizer.Prev().value
			predicate.addParameter(Parameter{matchedValue})
			addDomain(dp.program, matchedValue)
		} else {
			dp.Fail(STRING)
		}
	}
	template := dp.MatchPredicate(matchFactString)
	dp.Match(PERIOD)
	addFact(dp.program, template)
}

func (dp *DatalogParser) MatchFactList() {
	if dp.tokenizer.Current().tokenType == ID {
		dp.MatchFact()
		dp.MatchFactList()
	}
}

func (dp *DatalogParser) MatchRule() {
	var template Rule

	template.setHead(dp.MatchPredicate(matchID))
	dp.Match(COLON_DASH)
	predicate := dp.MatchPredicate(matchIDOrString)
	predicateList := dp.MatchPredicateList(matchIDOrString, []Predicate{predicate})
	template.setPredicates(predicateList)
	dp.Match(PERIOD)

	addRule(dp.program, template)
}

func (dp *DatalogParser) MatchRuleList() {
	if dp.tokenizer.Current().tokenType == ID {
		dp.MatchRule()
		dp.MatchRuleList()
	}
}

func (dp *DatalogParser) MatchQuery() {
	template := dp.MatchPredicate(matchIDOrString)
	dp.Match(Q_MARK)
	addQuery(dp.program, template)
}

func (dp *DatalogParser) MatchQueryList() {
	if dp.tokenizer.Current().tokenType == ID {
		dp.MatchQuery()
		dp.MatchQueryList()
	}
}

func (dp *DatalogParser) MatchPredicate(matchParameterType MatchParameterType) Predicate {
	var template Predicate

	dp.Match(ID)
	template.setID(dp.tokenizer.Prev().value)
	dp.Match(LEFT_PAREN)
	matchParameterType(dp, &template)
	dp.MatchParameterList(matchParameterType, &template)
	dp.Match(RIGHT_PAREN)

	return template
}

func (dp *DatalogParser) MatchPredicateList(matchParameterType MatchParameterType, list []Predicate) []Predicate {
	if dp.tokenizer.Current().tokenType == COMMA {
		dp.Match(COMMA)
		predicate := dp.MatchPredicate(matchParameterType)
		return dp.MatchPredicateList(matchParameterType, append(list, predicate))
	}
	return list
}

func (dp *DatalogParser) MatchParameterList(matchParameterType MatchParameterType, predicate *Predicate) {
	if dp.tokenizer.Current().tokenType == COMMA {
		dp.Match(COMMA)
		matchParameterType(dp, predicate)
		dp.MatchParameterList(matchParameterType, predicate)
	}
}

func (dp *DatalogParser) Match(t TokenType) {
	if dp.tokenizer.Current().tokenType == t {
		dp.tokenizer.Next()
	} else {
		dp.Fail(t)
	}
}

func (dp *DatalogParser) Fail(expected TokenType) {
	actual := dp.tokenizer.Current()
	err := errors.New(
		fmt.Sprintf(
			"Invalid syntax: expected %s but found %s (%s %d:%d)",
			expected, actual.value, dp.fileName, actual.lineNumber, actual.column,
		),
	)
	log.Fatal(err)
}

func addScheme(p *DatalogProgram, scheme Predicate) {
	p.schemes = append(p.schemes, scheme)
}

func addFact(p *DatalogProgram, fact Predicate) {
	p.facts = append(p.facts, fact)
}

func addRule(p *DatalogProgram, rule Rule) {
	p.rules = append(p.rules, rule)
}

func addQuery(p *DatalogProgram, query Predicate) {
	p.queries = append(p.queries, query)
}

func addDomain(p *DatalogProgram, value string) {
	p.domain.Add(value)
}

func matchID(dp *DatalogParser, predicate *Predicate) {
	if dp.tokenizer.Current().tokenType == ID {
		dp.Match(ID)
		predicate.addParameter(Parameter{dp.tokenizer.Prev().value})
	} else {
		dp.Fail(ID)
	}
}

func matchIDOrString(dp *DatalogParser, predicate *Predicate) {
	if dp.tokenizer.Current().tokenType == ID {
		dp.Match(ID)
		predicate.addParameter(Parameter{dp.tokenizer.Prev().value})
	} else if dp.tokenizer.Current().tokenType == STRING {
		dp.Match(STRING)
		predicate.addParameter(Parameter{dp.tokenizer.Prev().value})
	}
}
