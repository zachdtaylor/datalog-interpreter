package parser

type Rule struct {
	head       Predicate
	Predicates []Predicate
}

func (r *Rule) setHead(predicate Predicate) {
	r.head = predicate
}

func (r *Rule) setPredicates(predicates []Predicate) {
	r.Predicates = predicates
}

func (r *Rule) ID() string {
	return r.head.id
}
