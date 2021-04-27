package parser

import (
	"github.com/zachtylr21/datalog-interpreter/graph"
	"github.com/zachtylr21/datalog-interpreter/util"
)

type DatalogProgram struct {
	schemes []Scheme
	facts   []Fact
	rules   []Rule
	queries []Query
	domain  util.StringSet
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

func (p *DatalogProgram) RuleDependencies() graph.Graph {
	var graph graph.Graph
	graph.Init()

	for k, rule := range p.rules {
		graph.AddNode(k)

		for _, pred := range rule.Predicates {
			predID := pred.GetID()
			for j, rule2 := range p.rules {
				ruleID := rule2.ID()
				if ruleID == predID {
					graph.AddDependency(k, j)
				}
			}
		}
	}

	return graph
}
