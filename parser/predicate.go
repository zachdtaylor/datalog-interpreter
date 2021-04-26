package parser

type Predicate struct {
	id         string
	parameters []Parameter
}

func (p *Predicate) setID(id string) {
	p.id = id
}

func (p *Predicate) GetID() string {
	return p.id
}

func (p *Predicate) addParameter(param Parameter) {
	p.parameters = append(p.parameters, param)
}

func (p *Predicate) GetParameterValues() []string {
	values := []string{}
	for _, parameter := range p.parameters {
		values = append(values, parameter.value)
	}
	return values
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
