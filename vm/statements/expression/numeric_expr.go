package expression

import (
	"github.com/Runway-Club/flux_lang/vm/exception"
	"github.com/Runway-Club/flux_lang/vm/statements"
)

type NumericExpression struct {
	*statements.BaseStatement
	LeftExpr  *NumericExpression
	Op        Operator
	RightExpr *NumericExpression
	Value     float64
}

func (n NumericExpression) Execute(ctx *statements.ExecutionContext) *exception.BaseException {
	if n.LeftExpr == nil && n.RightExpr == nil {
		ctx.NumericValue = n.Value
		return nil
	}
	n.Op.SetLeftExpr(n.LeftExpr)
	n.Op.SetRightExpr(n.RightExpr)
	return n.Op.Execute(ctx)
}

func NewNumericExpression(line int, startPos int, endPos int, leftExpr *NumericExpression, operator Operator, rightExpr *NumericExpression, value float64) *NumericExpression {
	return &NumericExpression{BaseStatement: &statements.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, LeftExpr: leftExpr, Op: operator, RightExpr: rightExpr, Value: value}
}
