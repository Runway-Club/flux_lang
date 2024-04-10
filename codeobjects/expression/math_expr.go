package expression

import (
	codeobjects2 "github.com/Runway-Club/flux_lang/codeobjects"
	"github.com/Runway-Club/flux_lang/exception"
)

type MathExpression struct {
	*codeobjects2.BaseStatement
	NumericExpr *NumericExpression
}

func (m MathExpression) Generate(ctx *codeobjects2.GenerateContext) string {
	return m.NumericExpr.Generate(ctx)
}

func (m MathExpression) Execute(ctx *codeobjects2.ExecutionContext) *exception.BaseException {
	if m.NumericExpr != nil {
		return m.NumericExpr.Execute(ctx)
	}
	return nil
}
func NewMathExpression(line int, startPos int, endPos int, numericExpr *NumericExpression) *MathExpression {
	return &MathExpression{BaseStatement: &codeobjects2.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, NumericExpr: numericExpr}
}
