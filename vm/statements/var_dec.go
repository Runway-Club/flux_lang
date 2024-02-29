package statements

type VarDeclaration struct {
	*Statement
	Name string
	Type string
}

func NewVarDeclaration(line int, posFrom int, posTo int, name string, type_ string) *VarDeclaration {
	return &VarDeclaration{&Statement{line, posFrom, posTo}, name, type_}
}
