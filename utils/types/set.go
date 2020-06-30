package types

var present = struct{}{}

type Set struct {
	data map[interface{}]struct{}
}

// NewSet is a constructor and return an actual set
func NewSet(data ...interface{}) *Set {
	s := &Set{}
	s.data = make(map[interface{}]struct{})
	for _, key := range data {
		s.data[key] = present
	}
	return s
}

func (s *Set) Add(key interface{}) {
	s.data[key] = present
}

func (s Set) Remove(key interface{}) {
	delete(s.data, key)
}

func (s Set) Has(key interface{}) bool {
	_, in := s.data[key]
	return in
}

func (s Set) Length() int {
	return len(s.data)
}

func (s Set) IsEmpty() bool {
	return s.Length() == 0
}

func (s Set) Slice() []interface{} {
	out := []interface{}{}
	for key := range s.data {
		out = append(out, key)
	}
	return out
}
