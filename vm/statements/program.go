package statements

import "github.com/Runway-Club/flux_lang/vm/exception"

type Program struct {
	*BaseStatement
	Statements []Statement
	Libraries  []*Program
}

func (p Program) Execute(ctx *ExecutionContext) *exception.BaseException {
	for _, statement := range p.Statements {
		err := statement.Execute(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p Program) GetLine() int {
	return 0
}

func (p Program) GetStartPos() int {
	return 0
}

func (p Program) GetEndPos() int {
	return 0
}

func NewProgram() *Program {
	return &Program{
		Statements: make([]Statement, 0),
	}
}
