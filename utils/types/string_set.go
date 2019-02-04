package types

var stringPresent = struct{}{}

type sset struct {
	data map[string]struct{}
}

// Set is a constructor and return an actual set
func SSet(data ...string) *sset {
	s := &sset{}
	s.data = make(map[string]struct{})
	for _, key := range data {
		s.data[key] = stringPresent
	}
	return s
}

func (s *sset) Add(key string) {
	s.data[key] = stringPresent
}

func (s sset) Remove(key string) {
	delete(s.data, key)
}

func (s sset) Has(key string) bool {
	_, in := s.data[key]
	return in
}

func (s sset) Length() int {
	return len(s.data)
}

func (s sset) IsEmpty() bool {
	return s.Length() == 0
}

func (s sset) Slice() []string {
	out := []string{}
	for key := range s.data {
		out = append(out, key)
	}
	return out
}
