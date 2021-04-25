package parser

import "github.com/zachtylr21/datalog-interpreter/util"

type DatalogProgram struct {
	schemes []Predicate
	facts   []Predicate
	rules   []Rule
	queries []Predicate
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
