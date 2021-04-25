package parser

type Predicate struct {
	id         string
	parameters []Parameter
}

func (p *Predicate) setID(id string) {
	p.id = id
}

func (p *Predicate) getID() string {
	return p.id
}

func (p *Predicate) addParameter(param Parameter) {
	p.parameters = append(p.parameters, param)
}

func (p *Predicate) getParameters() []Parameter {
	return p.parameters
}

type Scheme struct {
	Predicate
}

type Fact struct {
	Predicate
}

type Query struct {
	Predicate
}
