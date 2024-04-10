package expression

import (
	"github.com/Runway-Club/flux_lang/codeobjects"
	"github.com/Runway-Club/flux_lang/exception"
)

type Operator interface {
	SetLeftExpr(leftExpr *NumericExpression)
	SetRightExpr(rightExpr *NumericExpression)
	GetLeftExpr() *NumericExpression
	GetRightExpr() *NumericExpression
	Generate(ctx *codeobjects.GenerateContext) string
	Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException
}
