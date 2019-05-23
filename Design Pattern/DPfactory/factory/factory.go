package factory

type Operation interface {
	GetName() string
}

type A struct{}
type B struct{}

type Factory struct{}

func (a *A) GetName() string {
	return "A"
}
func (b *B) GetName() string {
	return "B"
}

func (f *Factory) Create(name string) Operation {
	switch name {
	case "a":
		return new(A)
	case "b":
		return new(B)
	default:
		panic("name not exists")
	}
	return nil
}
