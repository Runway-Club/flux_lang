package common

type Var[T any] struct {
	Name  string
	Value T
}

func (v *Var[T]) GetName() string {
	return v.Name
}

func (v *Var[T]) GetValue() T {
	return v.Value
}

func (v *Var[T]) SetValue(value T) {
	v.Value = value
}
