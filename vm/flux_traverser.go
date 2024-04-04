package vm

import (
	"github.com/Runway-Club/flux_lang/parsing"
	"github.com/Runway-Club/flux_lang/vm/io"
	"github.com/Runway-Club/flux_lang/vm/statements"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type FluxTraverser struct {
	logger   io.Logger
	varTable *statements.VarTable
}

func (f FluxTraverser) EnterArgs(c *parsing.ArgsContext) {
	f.logger.Infof("Entering args")
}

func (f FluxTraverser) ExitArgs(c *parsing.ArgsContext) {
	f.logger.Infof("Exiting args")
}

func NewFluxTraverser(logger io.Logger, varTable *statements.VarTable) *FluxTraverser {
	return &FluxTraverser{
		logger:   logger,
		varTable: varTable,
	}
}

func (f FluxTraverser) VisitTerminal(node antlr.TerminalNode) {
	f.logger.Infof("Visiting terminal node: %v", node.GetText())
}

func (f FluxTraverser) VisitErrorNode(node antlr.ErrorNode) {
	f.logger.Errorf("Visiting error node: %v", node.GetText())
}

func (f FluxTraverser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	f.logger.Infof("Entering rule: %v", ctx.GetText())
}

func (f FluxTraverser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	f.logger.Infof("Exiting rule: %v", ctx.GetText())
}

func (f FluxTraverser) EnterProgram(c *parsing.ProgramContext) {
	f.logger.Infof("Entering program")
}

func (f FluxTraverser) EnterStatement(c *parsing.StatementContext) {
	f.logger.Infof("Entering statement")
}

func (f FluxTraverser) EnterExpression(c *parsing.ExpressionContext) {
	f.logger.Infof("Entering expression")
}

func (f FluxTraverser) EnterBlock(c *parsing.BlockContext) {
	f.logger.Infof("Entering block")
}

func (f FluxTraverser) EnterLoop_statement(c *parsing.Loop_statementContext) {
	f.logger.Infof("Entering loop statement")
}

func (f FluxTraverser) EnterIf_statement(c *parsing.If_statementContext) {
	f.logger.Infof("Entering if statement")
}

func (f FluxTraverser) EnterReturn_statement(c *parsing.Return_statementContext) {
	f.logger.Infof("Entering return statement")
}

func (f FluxTraverser) EnterData_type(c *parsing.Data_typeContext) {
	f.logger.Infof("Entering data type")
}

func (f FluxTraverser) EnterFunc_declaration(c *parsing.Func_declarationContext) {
	f.logger.Infof("Entering function declaration: ")
}

func (f FluxTraverser) EnterVar_type(c *parsing.Var_typeContext) {
	f.logger.Infof("Entering var type")
}

func (f FluxTraverser) EnterVar_name(c *parsing.Var_nameContext) {
	f.logger.Infof("Entering var name")
}

func (f FluxTraverser) EnterVar_value(c *parsing.Var_valueContext) {
	f.logger.Infof("Entering var value")
}

func (f *FluxTraverser) EnterString_var_declaration(c *parsing.String_var_declarationContext) {
	f.logger.Infof("Entering string var declaration")
}

func (f *FluxTraverser) EnterNumber_var_declaration(c *parsing.Number_var_declarationContext) {
	f.logger.Infof("Entering number var declaration")
}

func (f *FluxTraverser) EnterBoolean_var_declaration(c *parsing.Boolean_var_declarationContext) {
	f.logger.Infof("Entering boolean var declaration")
}

func (f FluxTraverser) EnterSingle_var_declaration(c *parsing.Single_var_declarationContext) {
	f.logger.Infof("Entering single var declaration")
}

func (f FluxTraverser) EnterArray_var_declaration(c *parsing.Array_var_declarationContext) {
	f.logger.Infof("Entering array var declaration")
}

func (f FluxTraverser) EnterVar_assignment(c *parsing.Var_assignmentContext) {
	f.logger.Infof("Entering var assignment")
}

func (f FluxTraverser) EnterOp_level1(c *parsing.Op_level1Context) {
	f.logger.Infof("Entering op level 1")
}

func (f FluxTraverser) EnterOp_level2(c *parsing.Op_level2Context) {
	f.logger.Infof("Entering op level 2")
}

func (f FluxTraverser) EnterOp_level3(c *parsing.Op_level3Context) {
	f.logger.Infof("Entering op level 3")
}

func (f FluxTraverser) EnterOp_level4(c *parsing.Op_level4Context) {
	f.logger.Infof("Entering op level 4")
}

func (f FluxTraverser) EnterOp_level5(c *parsing.Op_level5Context) {
	f.logger.Infof("Entering op level 5")
}

func (f FluxTraverser) EnterNumeric_expression(c *parsing.Numeric_expressionContext) {
	f.logger.Infof("Entering numeric expression")
}

func (f FluxTraverser) EnterLogical_expression(c *parsing.Logical_expressionContext) {
	f.logger.Infof("Entering logical expression")
}

func (f FluxTraverser) EnterComparative_expression(c *parsing.Comparative_expressionContext) {
	f.logger.Infof("Entering comparative expression")
}

func (f FluxTraverser) EnterMath_expression(c *parsing.Math_expressionContext) {
	f.logger.Infof("Entering math expression")
}

func (f FluxTraverser) EnterGet_array_element(c *parsing.Get_array_elementContext) {
	f.logger.Infof("Entering get array element")
}

func (f FluxTraverser) EnterGet_child(c *parsing.Get_childContext) {
	f.logger.Infof("Entering get child")
}

func (f FluxTraverser) EnterGet_var(c *parsing.Get_varContext) {
	f.logger.Infof("Entering get var")
}

func (f FluxTraverser) EnterFunction_call(c *parsing.Function_callContext) {
	f.logger.Infof("Entering function call")
}

func (f FluxTraverser) ExitProgram(c *parsing.ProgramContext) {
	f.logger.Infof("Exiting program")
}

func (f FluxTraverser) ExitStatement(c *parsing.StatementContext) {
	f.logger.Infof("Exiting statement")
}

func (f FluxTraverser) ExitExpression(c *parsing.ExpressionContext) {
	f.logger.Infof("Exiting expression")
}

func (f FluxTraverser) ExitBlock(c *parsing.BlockContext) {
	f.logger.Infof("Exiting block")
}

func (f FluxTraverser) ExitLoop_statement(c *parsing.Loop_statementContext) {
	f.logger.Infof("Exiting loop statement")
}

func (f FluxTraverser) ExitIf_statement(c *parsing.If_statementContext) {
	f.logger.Infof("Exiting if statement")
}

func (f FluxTraverser) ExitReturn_statement(c *parsing.Return_statementContext) {
	f.logger.Infof("Exiting return statement")
}

func (f FluxTraverser) ExitData_type(c *parsing.Data_typeContext) {
	f.logger.Infof("Exiting data type")
}

func (f FluxTraverser) ExitFunc_declaration(c *parsing.Func_declarationContext) {
	f.logger.Infof("Exiting function declaration")
}

func (f FluxTraverser) ExitVar_type(c *parsing.Var_typeContext) {
	f.logger.Infof("Exiting var type")
}

func (f FluxTraverser) ExitVar_name(c *parsing.Var_nameContext) {
	f.logger.Infof("Exiting var name")
}

func (f FluxTraverser) ExitVar_value(c *parsing.Var_valueContext) {
	f.logger.Infof("Exiting var value")
}

func (f FluxTraverser) ExitString_var_declaration(c *parsing.String_var_declarationContext) {
	f.logger.Infof("Exiting string var declaration")
}

func (f FluxTraverser) ExitNumber_var_declaration(c *parsing.Number_var_declarationContext) {
	f.logger.Infof("Exiting number var declaration")
	// get number value if it exists
	if c.NUMBER() == nil {
		if c.Math_expression() != nil {
			if c.Math_expression().Logical_expression() != nil {
				f.logger.Errorf("invalid variable value %v for %v as a num", c.Math_expression().GetText(), c.Var_name().GetText())
				return
			}
		}
		f.varTable.SetNum(c.Var_name().GetText(), 0)
	} else {
		numValue, err := strconv.ParseInt(c.NUMBER().GetText(), 10, 64)
		if err != nil {
			f.logger.Errorf("invalid variable value %v for %v as a num", c.NUMBER().GetText(), c.Var_name().GetText())
		}
		f.varTable.SetNum(c.Var_name().GetText(), int(numValue))
	}
}

func (f FluxTraverser) ExitBoolean_var_declaration(c *parsing.Boolean_var_declarationContext) {
	f.logger.Infof("Exiting boolean var declaration")
}

func (f FluxTraverser) ExitSingle_var_declaration(c *parsing.Single_var_declarationContext) {
	f.logger.Infof("Exiting single var declaration")
}

func (f FluxTraverser) ExitArray_var_declaration(c *parsing.Array_var_declarationContext) {
	f.logger.Infof("Exiting array var declaration")
}

func (f FluxTraverser) ExitVar_assignment(c *parsing.Var_assignmentContext) {
	f.logger.Infof("Exiting var assignment")
}

func (f FluxTraverser) ExitOp_level1(c *parsing.Op_level1Context) {
	f.logger.Infof("Exiting op level 1")
}

func (f FluxTraverser) ExitOp_level2(c *parsing.Op_level2Context) {
	f.logger.Infof("Exiting op level 2")
}

func (f FluxTraverser) ExitOp_level3(c *parsing.Op_level3Context) {
	f.logger.Infof("Exiting op level 3")
}

func (f FluxTraverser) ExitOp_level4(c *parsing.Op_level4Context) {
	f.logger.Infof("Exiting op level 4")
}

func (f FluxTraverser) ExitOp_level5(c *parsing.Op_level5Context) {
	f.logger.Infof("Exiting op level 5")
}

func (f FluxTraverser) ExitNumeric_expression(c *parsing.Numeric_expressionContext) {
	f.logger.Infof("Exiting numeric expression")
}

func (f FluxTraverser) ExitLogical_expression(c *parsing.Logical_expressionContext) {
	f.logger.Infof("Exiting logical expression")
}

func (f FluxTraverser) ExitComparative_expression(c *parsing.Comparative_expressionContext) {
	f.logger.Infof("Exiting comparative expression")
}

func (f FluxTraverser) ExitMath_expression(c *parsing.Math_expressionContext) {
	f.logger.Infof("Exiting math expression")
}

func (f FluxTraverser) ExitGet_array_element(c *parsing.Get_array_elementContext) {
	f.logger.Infof("Exiting get array element")
}

func (f FluxTraverser) ExitGet_child(c *parsing.Get_childContext) {
	f.logger.Infof("Exiting get child")
}

func (f FluxTraverser) ExitGet_var(c *parsing.Get_varContext) {
	f.logger.Infof("Exiting get var")
}

func (f FluxTraverser) ExitFunction_call(c *parsing.Function_callContext) {
	f.logger.Infof("Exiting function call")
}
