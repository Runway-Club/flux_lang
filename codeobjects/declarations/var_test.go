package declarations

import (
	"context"
	"github.com/Runway-Club/flux_lang/codeobjects"
	expression2 "github.com/Runway-Club/flux_lang/codeobjects/expression"
	operators2 "github.com/Runway-Club/flux_lang/codeobjects/expression/operators"
	"github.com/Runway-Club/flux_lang/core"
	"testing"
)

func TestVarDeclaration_Execute(t *testing.T) {
	varTable := core.NewVarTable()
	t.Run("Create a new variable", func(t *testing.T) {
		varDeclaration := NewVarDeclaration(1, 1, 1, "testVar", FluxTypeNumber, "",
			expression2.NewMathExpression(1, 1, 1,
				expression2.NewNumericExpression(1, 1, 1,
					expression2.NewNumericExpression(1, 1, 1, nil, operators2.NullOp, nil, float64(2)),
					operators2.NewAdd(1, 1, 1, nil, nil),
					expression2.NewNumericExpression(1, 1, 1, nil, operators2.NullOp, nil, float64(3)), 0),
			),
		)
		ctx := codeobjects.NewExecutionContext(context.Background(), varTable)
		err := varDeclaration.Execute(ctx)
		if err != nil {
			t.Errorf(err.MessageFmt, err.Args...)
		}
		testVar, except := varTable.GetNumValue("testVar")
		if except != nil {
			t.Errorf("Expected nil, got %v", except)
		}
		if testVar != 5 {
			t.Errorf("Expected 5, got %v", testVar)
		}
	})
}
