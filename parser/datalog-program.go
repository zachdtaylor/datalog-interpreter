package parser

type DatalogProgram struct {
	schemes []Predicate
	facts   []Predicate
	rules   []Rule
	queries []Predicate
}

func (p *DatalogProgram) Run(fileName string) {
	var parser DatalogParser
	parser.Init(p, fileName)
	parser.Run()
}
