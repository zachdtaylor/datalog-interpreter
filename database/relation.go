package database

type Tuple []string

type Schema []string

type Relation struct {
	Name   string
	Schema Schema
	Tuples []Tuple
}
