package statements

import "github.com/Runway-Club/flux_lang/vm/exception"

type Statement interface {
	Execute(ctx *ExecutionContext) *exception.BaseException
	GetLine() int
	GetStartPos() int
	GetEndPos() int
}

type BaseStatement struct {
	Line     int
	StartPos int
	EndPos   int
}

func (b BaseStatement) Execute(ctx *ExecutionContext) *exception.BaseException {
	return nil
}

func (b BaseStatement) GetLine() int {
	return b.Line
}

func (b BaseStatement) GetStartPos() int {
	return b.StartPos
}

func (b BaseStatement) GetEndPos() int {
	return b.EndPos
}
