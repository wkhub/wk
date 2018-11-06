package mixins

type INamed interface {
	GetName() string
}

type Named struct {
	Name string
}

func (n Named) GetName() string {
	return n.Name
}
