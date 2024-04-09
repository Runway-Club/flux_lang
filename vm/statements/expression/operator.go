package expression

import (
	"github.com/Runway-Club/flux_lang/vm/exception"
	"github.com/Runway-Club/flux_lang/vm/statements"
)

type Operator interface {
	SetLeftExpr(leftExpr *NumericExpression)
	SetRightExpr(rightExpr *NumericExpression)
	GetLeftExpr() *NumericExpression
	GetRightExpr() *NumericExpression
	Execute(ctx *statements.ExecutionContext) *exception.BaseException
}
