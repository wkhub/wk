package types

var present = struct{}{}

type set struct {
	data map[interface{}]struct{}
}

// Set is a constructor and return an actual set
func Set(data ...interface{}) *set {
	s := &set{}
	s.data = make(map[interface{}]struct{})
	for _, key := range data {
		s.data[key] = present
	}
	return s
}

func (s *set) Add(key interface{}) {
	s.data[key] = present
}

func (s set) Remove(key interface{}) {
	delete(s.data, key)
}

func (s set) Has(key interface{}) bool {
	_, in := s.data[key]
	return in
}

func (s set) Length() int {
	return len(s.data)
}

func (s set) Slice() []interface{} {
	out := []interface{}{}
	for key := range s.data {
		out = append(out, key)
	}
	return out
}
