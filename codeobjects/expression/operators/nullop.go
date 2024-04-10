package operators

import (
	"github.com/Runway-Club/flux_lang/codeobjects"
	"github.com/Runway-Club/flux_lang/codeobjects/expression"
	"github.com/Runway-Club/flux_lang/exception"
)

type NullOperator struct {
}

func (n NullOperator) SetLeftExpr(leftExpr *expression.NumericExpression) {}

func (n NullOperator) SetRightExpr(rightExpr *expression.NumericExpression) {}

func (n NullOperator) GetLeftExpr() *expression.NumericExpression {
	return nil
}

func (n NullOperator) GetRightExpr() *expression.NumericExpression {
	return nil
}

func (n NullOperator) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	return nil
}

var NullOp = NullOperator{}
