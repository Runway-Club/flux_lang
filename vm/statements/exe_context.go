package statements

import (
	"context"
	"github.com/Runway-Club/flux_lang/vm/core"
)

type ExecutionContext struct {
	ctx          context.Context
	NumericValue float64
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