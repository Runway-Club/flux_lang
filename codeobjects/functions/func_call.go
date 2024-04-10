package functions

import (
	"fmt"
	"github.com/Runway-Club/flux_lang/codeobjects"
	"github.com/Runway-Club/flux_lang/exception"
)

type FunctionCall struct {
	*codeobjects.BaseStatement
	Name string
	Args *Args
}

func NewFunctionCall(baseStatement *codeobjects.BaseStatement, name string, args *Args) *FunctionCall {
	return &FunctionCall{BaseStatement: baseStatement, Name: name, Args: args}
}

func (f FunctionCall) Generate(ctx *codeobjects.GenerateContext) string {
	if f.Name == "log" {
		return NewLog(f.Args).Generate(ctx)
	}
	return fmt.Sprintf("%v(%v)", f.Name, f.Args.Generate(ctx))
}

func (f FunctionCall) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	if f.Name == "log" {
		return NewLog(f.Args).Execute(ctx)
	}
	return nil
}
