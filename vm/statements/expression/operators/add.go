package operators

import (
	"github.com/Runway-Club/flux_lang/vm/exception"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/Runway-Club/flux_lang/vm/statements/expression"
)

type Add struct {
	*statements.BaseStatement
	Left  *expression.NumericExpression
	Right *expression.NumericExpression
}

func (a *Add) SetLeftExpr(leftExpr *expression.NumericExpression) {
	a.Left = leftExpr
}

func (a *Add) SetRightExpr(rightExpr *expression.NumericExpression) {
	a.Right = rightExpr
}

func (a *Add) GetLeftExpr() *expression.NumericExpression {
	return a.Left
}

func (a *Add) GetRightExpr() *expression.NumericExpression {
	return a.Right
}

func (a *Add) Execute(ctx *statements.ExecutionContext) *exception.BaseException {
	leftCtx := ctx.Clone()
	rightCtx := ctx.Clone()
	if a.Left != nil {
		err := a.Left.Execute(leftCtx)
		if err != nil {
			return err
		}
	} else {
		leftCtx.NumericValue = 0
	}
	if a.Right != nil {
		err := a.Right.Execute(rightCtx)
		if err != nil {
			return err
		}
	} else {
		rightCtx.NumericValue = 0
	}
	ctx.NumericValue = leftCtx.NumericValue + rightCtx.NumericValue
	return nil
}

func NewAdd(line int, startPos int, endPos int, left *expression.NumericExpression, right *expression.NumericExpression) *Add {
	return &Add{BaseStatement: &statements.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, Left: left, Right: right}
}
