package codeobjects

import (
	"context"
	"github.com/Runway-Club/flux_lang/core"
)

type ExecutionContext struct {
	ctx          context.Context
	NumericValue float64
	TextValue    string
	BoolValue    bool
	VarTable     *core.VarTable
}

func (e ExecutionContext) Clone() *ExecutionContext {
	return &ExecutionContext{
		ctx:          e.ctx,
		NumericValue: e.NumericValue,
		VarTable:     e.VarTable,
	}
}

func NewExecutionContext(ctx context.Context, varTable *core.VarTable) *ExecutionContext {
	return &ExecutionContext{
		ctx:          ctx,
		NumericValue: 0,
		VarTable:     varTable,
	}
}

type GenerateContext struct {
	Ctx     context.Context
	CodeFmt string
	Args    []interface{}
}

func (e GenerateContext) Clone() *GenerateContext {
	return &GenerateContext{
		Ctx:     e.Ctx,
		CodeFmt: e.CodeFmt,
		Args:    e.Args,
	}
}

func NewGenerateContext(ctx context.Context) *GenerateContext {
	return &GenerateContext{
		Ctx: ctx,
	}
}
