package vm

import (
	"github.com/Runway-Club/flux_lang/parsing"
	"github.com/Runway-Club/flux_lang/vm/core"
	"github.com/Runway-Club/flux_lang/vm/exception"
	"github.com/Runway-Club/flux_lang/vm/io"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/Runway-Club/flux_lang/vm/statements/declarations"
	"github.com/Runway-Club/flux_lang/vm/statements/expression"
	"github.com/Runway-Club/flux_lang/vm/statements/expression/operators"
	"github.com/antlr4-go/antlr/v4"
)

type FluxProgramParser struct {
	logger           io.Logger
	errorCollector   io.ErrorCollector
	program          *statements.Program
	currentStatement statements.Statement
	executionCtx     *statements.ExecutionContext
}

func (f FluxProgramParser) EnterArgs(c *parsing.ArgsContext) {
	f.logger.Infof("Entering args")
}

func (f FluxProgramParser) ExitArgs(c *parsing.ArgsContext) {
	f.logger.Infof("Exiting args")
}

func NewFluxProgramParser(logger io.Logger, errorCollector io.ErrorCollector, varOperator core.VarOperator) *FluxProgramParser {
	program := &statements.Program{}
	currentStatement := program
	return &FluxProgramParser{
		logger:           logger,
		errorCollector:   errorCollector,
		program:          program,
		currentStatement: currentStatement,
	}
}

func (f FluxProgramParser) GetProgram() *statements.Program {
	return f.program
}

func (f FluxProgramParser) VisitTerminal(node antlr.TerminalNode) {
	f.logger.Infof("Visiting terminal node: %v", node.GetText())
}

func (f FluxProgramParser) VisitErrorNode(node antlr.ErrorNode) {
	f.logger.Errorf("Visiting error node: %v", node.GetText())
}

func (f FluxProgramParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	f.logger.Infof("Entering rule: %v", ctx.GetText())
}

func (f FluxProgramParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	f.logger.Infof("Exiting rule: %v", ctx.GetText())
}

func (f FluxProgramParser) EnterProgram(c *parsing.ProgramContext) {
	f.logger.Infof("Entering program")
}

func (f FluxProgramParser) EnterStatement(c *parsing.StatementContext) {
	f.logger.Infof("Entering statement")
}

func (f FluxProgramParser) EnterExpression(c *parsing.ExpressionContext) {
	f.logger.Infof("Entering expression")
}

func (f FluxProgramParser) EnterBlock(c *parsing.BlockContext) {
	f.logger.Infof("Entering block")
}

func (f FluxProgramParser) EnterLoop_statement(c *parsing.Loop_statementContext) {
	f.logger.Infof("Entering loop statement")
}

func (f FluxProgramParser) EnterIf_statement(c *parsing.If_statementContext) {
	f.logger.Infof("Entering if statement")
}

func (f FluxProgramParser) EnterReturn_statement(c *parsing.Return_statementContext) {
	f.logger.Infof("Entering return statement")
}

func (f FluxProgramParser) EnterData_type(c *parsing.Data_typeContext) {
	f.logger.Infof("Entering data type")
}

func (f FluxProgramParser) EnterFunc_declaration(c *parsing.Func_declarationContext) {
	f.logger.Infof("Entering function declaration: ")
}

func (f FluxProgramParser) EnterVar_type(c *parsing.Var_typeContext) {
	f.logger.Infof("Entering var type")
}

func (f FluxProgramParser) EnterVar_name(c *parsing.Var_nameContext) {
	f.logger.Infof("Entering var name")
}

func (f FluxProgramParser) EnterVar_value(c *parsing.Var_valueContext) {
	f.logger.Infof("Entering var value")
}

func (f *FluxProgramParser) EnterString_var_declaration(c *parsing.String_var_declarationContext) {
	f.logger.Infof("Entering string var declaration")
}

func (f *FluxProgramParser) EnterNumber_var_declaration(c *parsing.Number_var_declarationContext) {
	f.logger.Infof("Entering number var declaration")
}

func (f *FluxProgramParser) EnterBoolean_var_declaration(c *parsing.Boolean_var_declarationContext) {
	f.logger.Infof("Entering boolean var declaration")
}

func (f FluxProgramParser) EnterSingle_var_declaration(c *parsing.Single_var_declarationContext) {
	f.logger.Infof("Entering single var declaration")
}

func (f FluxProgramParser) EnterArray_var_declaration(c *parsing.Array_var_declarationContext) {
	f.logger.Infof("Entering array var declaration")
}

func (f FluxProgramParser) EnterVar_assignment(c *parsing.Var_assignmentContext) {
	f.logger.Infof("Entering var assignment")
}

func (f FluxProgramParser) EnterOp_level1(c *parsing.Op_level1Context) {
	f.logger.Infof("Entering op level 1")
}

func (f FluxProgramParser) EnterOp_level2(c *parsing.Op_level2Context) {
	f.logger.Infof("Entering op level 2")
}

func (f FluxProgramParser) EnterOp_level3(c *parsing.Op_level3Context) {
	f.logger.Infof("Entering op level 3")
}

func (f FluxProgramParser) EnterOp_level4(c *parsing.Op_level4Context) {
	f.logger.Infof("Entering op level 4")
}

func (f FluxProgramParser) EnterOp_level5(c *parsing.Op_level5Context) {
	f.logger.Infof("Entering op level 5")
}

func (f FluxProgramParser) EnterNumeric_expression(c *parsing.Numeric_expressionContext) {
	f.logger.Infof("Entering numeric expression")
	// check if current statement is a math expression
	if mathExpr, ok := f.currentStatement.(*expression.MathExpression); ok {
		if len(c.AllNumeric_expression()) == 2 {
			numericExpr := &expression.NumericExpression{
				BaseStatement: &statements.BaseStatement{
					Line:     c.GetStart().GetLine(),
					StartPos: c.GetStart().GetStart(),
					EndPos:   c.GetStop().GetStop(),
				},
				LeftExpr:  nil,
				RightExpr: nil,
				Op:        operators.NullOp,
			}
			mathExpr.NumericExpr = numericExpr
			f.currentStatement = numericExpr
		}
	}

	// check if current statement is a numeric expression
	if numericExpr, ok := f.currentStatement.(*expression.NumericExpression); ok {
		if len(c.AllNumeric_expression()) == 2 {
			leftExpr := &expression.NumericExpression{
				BaseStatement: &statements.BaseStatement{
					Line:     c.GetStart().GetLine(),
					StartPos: c.GetStart().GetStart(),
					EndPos:   c.GetStop().GetStop(),
				},
				LeftExpr:  nil,
				RightExpr: nil,
				Op:        operators.NullOp,
			}
			numericExpr.LeftExpr = leftExpr
			f.currentStatement = leftExpr
		} else if c.NUMBER() != nil {
			
		}

	}
}

func (f FluxProgramParser) EnterLogical_expression(c *parsing.Logical_expressionContext) {
	f.logger.Infof("Entering logical expression")
}

func (f FluxProgramParser) EnterComparative_expression(c *parsing.Comparative_expressionContext) {
	f.logger.Infof("Entering comparative expression")
}

func (f FluxProgramParser) EnterMath_expression(c *parsing.Math_expressionContext) {
	f.logger.Infof("Entering math expression")
}

func (f FluxProgramParser) EnterGet_array_element(c *parsing.Get_array_elementContext) {
	f.logger.Infof("Entering get array element")
}

func (f FluxProgramParser) EnterGet_child(c *parsing.Get_childContext) {
	f.logger.Infof("Entering get child")
}

func (f FluxProgramParser) EnterGet_var(c *parsing.Get_varContext) {
	f.logger.Infof("Entering get var")
}

func (f FluxProgramParser) EnterFunction_call(c *parsing.Function_callContext) {
	f.logger.Infof("Entering function call")
}

func (f FluxProgramParser) ExitProgram(c *parsing.ProgramContext) {
	f.logger.Infof("Exiting program")
}

func (f FluxProgramParser) ExitStatement(c *parsing.StatementContext) {
	f.logger.Infof("Exiting statement")
}

func (f FluxProgramParser) ExitExpression(c *parsing.ExpressionContext) {
	f.logger.Infof("Exiting expression")
}

func (f FluxProgramParser) ExitBlock(c *parsing.BlockContext) {
	f.logger.Infof("Exiting block")
}

func (f FluxProgramParser) ExitLoop_statement(c *parsing.Loop_statementContext) {
	f.logger.Infof("Exiting loop statement")
}

func (f FluxProgramParser) ExitIf_statement(c *parsing.If_statementContext) {
	f.logger.Infof("Exiting if statement")
}

func (f FluxProgramParser) ExitReturn_statement(c *parsing.Return_statementContext) {
	f.logger.Infof("Exiting return statement")
}

func (f FluxProgramParser) ExitData_type(c *parsing.Data_typeContext) {
	f.logger.Infof("Exiting data type")
}

func (f FluxProgramParser) ExitFunc_declaration(c *parsing.Func_declarationContext) {
	f.logger.Infof("Exiting function declaration")
}

func (f FluxProgramParser) ExitVar_type(c *parsing.Var_typeContext) {
	f.logger.Infof("Exiting var type")
}

func (f FluxProgramParser) ExitVar_name(c *parsing.Var_nameContext) {
	f.logger.Infof("Exiting var name")
}

func (f FluxProgramParser) ExitVar_value(c *parsing.Var_valueContext) {
	f.logger.Infof("Exiting var value")
}

func (f FluxProgramParser) ExitString_var_declaration(c *parsing.String_var_declarationContext) {
	f.logger.Infof("Exiting string var declaration")
}

func (f FluxProgramParser) ExitNumber_var_declaration(c *parsing.Number_var_declarationContext) {
	f.logger.Infof("Exiting number var declaration")
	// get number value if it exists
	if c.NUMBER() == nil {
		if c.Math_expression() != nil {
			if c.Math_expression().Logical_expression() != nil {
				f.errorCollector.CollectError(&exception.BaseException{
					MessageFmt: "invalid variable value %v for %v as a num",
					Args:       []interface{}{c.Math_expression().GetText(), c.Var_name().GetText()},
					Err:        nil,
					Line:       c.GetStart().GetLine(),
					StartPos:   c.GetStart().GetStart(),
					EndPos:     c.GetStop().GetStop(),
				})
			} else if c.Math_expression().Numeric_expression() != nil {
				mathExpr := &expression.MathExpression{
					BaseStatement: &statements.BaseStatement{
						Line:     c.GetStart().GetLine(),
						StartPos: c.GetStart().GetStart(),
						EndPos:   c.GetStop().GetStop(),
					},
					NumericExpr: nil,
				}
				f.currentStatement = mathExpr
				varDecStmt := declarations.NewVarDeclaration(c.GetStart().GetLine(), c.GetStart().GetStart(), c.GetStop().GetStop(), c.Var_name().GetText(), declarations.FluxTypeNumber, "", mathExpr)
				f.program.Statements = append(f.program.Statements, varDecStmt)
			}
		}
		varDecStmt := declarations.NewVarDeclaration(c.GetStart().GetLine(), c.GetStart().GetStart(), c.GetStop().GetStop(), c.Var_name().GetText(), declarations.FluxTypeNumber, "", nil)
		f.program.Statements = append(f.program.Statements, varDecStmt)
	} else {
		varDecStmt := declarations.NewVarDeclaration(c.GetStart().GetLine(), c.GetStart().GetStart(), c.GetStop().GetStop(), c.Var_name().GetText(), declarations.FluxTypeNumber, c.NUMBER().GetText(), nil)
		f.program.Statements = append(f.program.Statements, varDecStmt)
	}

}

func (f FluxProgramParser) ExitBoolean_var_declaration(c *parsing.Boolean_var_declarationContext) {
	f.logger.Infof("Exiting boolean var declaration")
}

func (f FluxProgramParser) ExitSingle_var_declaration(c *parsing.Single_var_declarationContext) {
	f.logger.Infof("Exiting single var declaration")
}

func (f FluxProgramParser) ExitArray_var_declaration(c *parsing.Array_var_declarationContext) {
	f.logger.Infof("Exiting array var declaration")
}

func (f FluxProgramParser) ExitVar_assignment(c *parsing.Var_assignmentContext) {
	f.logger.Infof("Exiting var assignment")
}

func (f FluxProgramParser) ExitOp_level1(c *parsing.Op_level1Context) {
	f.logger.Infof("Exiting op level 1")
}

func (f FluxProgramParser) ExitOp_level2(c *parsing.Op_level2Context) {
	f.logger.Infof("Exiting op level 2")
}

func (f FluxProgramParser) ExitOp_level3(c *parsing.Op_level3Context) {
	f.logger.Infof("Exiting op level 3")
}

func (f FluxProgramParser) ExitOp_level4(c *parsing.Op_level4Context) {
	f.logger.Infof("Exiting op level 4")
}

func (f FluxProgramParser) ExitOp_level5(c *parsing.Op_level5Context) {
	f.logger.Infof("Exiting op level 5")
}

func (f FluxProgramParser) ExitNumeric_expression(c *parsing.Numeric_expressionContext) {
	f.logger.Infof("Exiting numeric expression")

}

func (f FluxProgramParser) ExitLogical_expression(c *parsing.Logical_expressionContext) {
	f.logger.Infof("Exiting logical expression")
}

func (f FluxProgramParser) ExitComparative_expression(c *parsing.Comparative_expressionContext) {
	f.logger.Infof("Exiting comparative expression")
}

func (f FluxProgramParser) ExitMath_expression(c *parsing.Math_expressionContext) {
	f.logger.Infof("Exiting math expression")
}

func (f FluxProgramParser) ExitGet_array_element(c *parsing.Get_array_elementContext) {
	f.logger.Infof("Exiting get array element")
}

func (f FluxProgramParser) ExitGet_child(c *parsing.Get_childContext) {
	f.logger.Infof("Exiting get child")
}

func (f FluxProgramParser) ExitGet_var(c *parsing.Get_varContext) {
	f.logger.Infof("Exiting get var")
}

func (f FluxProgramParser) ExitFunction_call(c *parsing.Function_callContext) {
	f.logger.Infof("Exiting function call")
}
