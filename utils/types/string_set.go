package types

var stringPresent = struct{}{}

type StringSet struct {
	data map[string]struct{}
}

// Set is a constructor and return an actual set
func NewStringSet(data ...string) *StringSet {
	s := &StringSet{}
	s.data = make(map[string]struct{})
	for _, key := range data {
		s.data[key] = stringPresent
	}
	return s
}

func (s *StringSet) Add(key string) {
	s.data[key] = stringPresent
}

func (s StringSet) Remove(key string) {
	delete(s.data, key)
}

func (s StringSet) Has(key string) bool {
	_, in := s.data[key]
	return in
}

func (s StringSet) Length() int {
	return len(s.data)
}

func (s StringSet) IsEmpty() bool {
	return s.Length() == 0
}

func (s StringSet) Slice() []string {
	out := []string{}
	for key := range s.data {
		out = append(out, key)
	}
	return out
}
