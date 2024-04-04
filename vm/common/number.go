package common

type Number struct {
	*Var[int]
}

func NewNumber(name string, value int) *Number {
	return &Number{
		Var: &Var[int]{
			name:  name,
			value: value,
		},
	}
}
