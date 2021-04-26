package util

type StringSet struct {
	values map[string]bool
}

func (s *StringSet) Init() {
	s.values = make(map[string]bool)
}

func (s *StringSet) Add(value string) {
	s.values[value] = true
}

func (s *StringSet) Array() []string {
	keys := make([]string, 0, len(s.values))
	for k := range s.values {
		keys = append(keys, k)
	}
	return keys
}

type IntSet struct {
	values map[int]bool
}

func (s *IntSet) Add(value int) {
	if s.values == nil {
		s.values = make(map[int]bool)
	}
	s.values[value] = true
}

func (s *IntSet) Array() []int {
	keys := make([]int, 0, len(s.values))
	for k := range s.values {
		keys = append(keys, k)
	}
	return keys
}
