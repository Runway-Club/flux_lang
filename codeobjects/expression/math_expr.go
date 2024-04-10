package expression

import (
	codeobjects2 "github.com/Runway-Club/flux_lang/codeobjects"
	"github.com/Runway-Club/flux_lang/exception"
)

type MathExpression struct {
	*codeobjects2.BaseStatement
	NumericExpr *NumericExpression
	GetVar      *GetVar
}

func (m MathExpression) Generate(ctx *codeobjects2.GenerateContext) string {
	if m.NumericExpr != nil {
		return m.NumericExpr.Generate(ctx)
	}
	if m.GetVar != nil {
		return m.GetVar.Generate(ctx)
	}
	return ""
}

func (m MathExpression) Execute(ctx *codeobjects2.ExecutionContext) *exception.BaseException {
	if m.NumericExpr != nil {
		return m.NumericExpr.Execute(ctx)
	}
	if m.GetVar != nil {
		return m.GetVar.Execute(ctx)
	}
	return nil
}
func NewMathExpression(line int, startPos int, endPos int, numericExpr *NumericExpression) *MathExpression {
	return &MathExpression{BaseStatement: &codeobjects2.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, NumericExpr: numericExpr}
}
