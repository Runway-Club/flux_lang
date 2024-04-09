package operators

import (
	"github.com/Runway-Club/flux_lang/vm/exception"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/Runway-Club/flux_lang/vm/statements/expression"
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

func (n NullOperator) Execute(ctx *statements.ExecutionContext) *exception.BaseException {
	return nil
}

var NullOp = NullOperator{}
