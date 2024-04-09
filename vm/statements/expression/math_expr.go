package expression

import (
	"github.com/Runway-Club/flux_lang/vm/exception"
	"github.com/Runway-Club/flux_lang/vm/statements"
)

type MathExpression struct {
	*statements.BaseStatement
	NumericExpr *NumericExpression
}

func (m MathExpression) Execute(ctx *statements.ExecutionContext) *exception.BaseException {
	if m.NumericExpr != nil {
		return m.NumericExpr.Execute(ctx)
	}
	return nil
}
func NewMathExpression(line int, startPos int, endPos int, numericExpr *NumericExpression) *MathExpression {
	return &MathExpression{BaseStatement: &statements.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, NumericExpr: numericExpr}
}
