package parser

import "github.com/zachtylr21/datalog-interpreter/util"

type DatalogProgram struct {
	schemes []Scheme
	facts   []Fact
	rules   []Rule
	queries []Query
	domain  util.StringSet
}

func (p *DatalogProgram) Init() {
	p.domain.Init()
}

func (p *DatalogProgram) Run(fileName string) {
	var parser DatalogParser
	parser.Init(p, fileName)
	parser.Run()
}

func (p *DatalogProgram) Schemes() []Scheme {
	return p.schemes
}

func (p *DatalogProgram) Facts() []Fact {
	return p.facts
}
