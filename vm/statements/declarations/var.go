package declarations

import (
	"github.com/Runway-Club/flux_lang/vm/exception"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/Runway-Club/flux_lang/vm/statements/expression"
	"strconv"
)

type FluxType string

const (
	FluxTypeNumber FluxType = "num"
	FluxTypeString FluxType = "text"
	FluxTypeBool   FluxType = "bool"
)

type VarDeclaration struct {
	*statements.BaseStatement
	Name     string
	Type     FluxType
	RawValue string
	Expr     *expression.MathExpression
}

func (v VarDeclaration) Execute(ctx *statements.ExecutionContext) *exception.BaseException {
	exprCtx := ctx.Clone()
	if v.Expr != nil {
		err := v.Expr.Execute(exprCtx)
		if err != nil {
			return err
		}
		if v.Type == FluxTypeNumber {
			err := ctx.VarTable.SetNum(v.Name, exprCtx.NumericValue)
			if err != nil {
				return err
			}
		}
		// TODO: Implement string and bool types
	} else if v.RawValue != "" {
		if v.Type == FluxTypeNumber {
			value, parseErr := strconv.ParseFloat(v.RawValue, 64)
			if parseErr != nil {
				return &exception.BaseException{
					MessageFmt: "Invalid number value: %s",
					Args:       []interface{}{v.RawValue},
					Line:       v.GetLine(),
					StartPos:   v.GetStartPos(),
					EndPos:     v.GetEndPos(),
				}
			}
			err := ctx.VarTable.SetNum(v.Name, value)
			if err != nil {
				return err
			}
		}
		// TODO: Implement string and bool types
	} else {
		// Set default value
		if v.Type == FluxTypeNumber {
			err := ctx.VarTable.SetNum(v.Name, 0)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func NewVarDeclaration(line int, startPos int, endPos int, name string, type_ FluxType, rawValue string, expr *expression.MathExpression) *VarDeclaration {
	return &VarDeclaration{BaseStatement: &statements.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, Name: name, Type: type_, RawValue: rawValue, Expr: expr}
}
