package main

type Rule struct {
	head       Predicate
	predicates []Predicate
}

func (r *Rule) setHead(predicate Predicate) {
	r.head = predicate
}

func (r *Rule) setPredicates(predicates []Predicate) {
	r.predicates = predicates
}
