// Methods for procedural language.
package printer

import (
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

func (p *Printer) VisitAssignmentFromStruct(ctx Context, n *googlesql.ASTAssignmentFromStruct) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("SET") + " (")
	pp.print(pp.toString(ctx, ast.Must(n.Variables())) + ")")
	pp.print("=")
	p.print(pp.unnestLeft())
	p.accept(ctx, ast.Must(n.StructExpression()))
	p.movePast(n)
}

func (p *Printer) VisitBeginEndBlock(ctx Context, n *googlesql.ASTBeginEndBlock) {
	p.moveBefore(n)
	p.println(p.keyword("BEGIN"))
	p.incDepth()
	p.accept(ctx, ast.Must(n.StatementListNode()))
	if ast.Must(n.HasExceptionHandler()) {
		p.decDepth()
		p.println(p.keyword("EXCEPTION WHEN ERROR THEN"))
		p.incDepth()
		p.accept(ctx, ast.Must(n.HandlerList()))
	}
	p.movePast(n)
	p.println("")
	p.decDepth()
	p.println(p.keyword("END"))
}

func (p *Printer) VisitBeginStatementNode(ctx Context, n *googlesql.ASTBeginStatement) {
	p.moveBefore(n)
	p.println(p.keyword("BEGIN TRANSACTION"))
	p.movePast(n)
}

func (p *Printer) VisitRollbackStatementNode(ctx Context, n *googlesql.ASTRollbackStatement) {
	p.moveBefore(n)
	p.println(p.keyword("ROLLBACK TRANSACTION"))
	p.movePast(n)
}

func (p *Printer) VisitExceptionHandlerListNode(ctx Context, n *googlesql.ASTExceptionHandlerList) {
	for _, node := range ast.ChildrenOfType[*googlesql.ASTExceptionHandler](n) {
		p.accept(ctx, node)
	}
}

func (p *Printer) VisitExceptionHandlerNode(ctx Context, n *googlesql.ASTExceptionHandler) {
	p.accept(ctx, ast.Must(n.StatementList()))
}

func (p *Printer) VisitCallStatement(ctx Context, n *googlesql.ASTCallStatement) {
	p.moveBefore(n)
	p.print(p.keyword("CALL"))
	p.accept(ctx.WithValue(KeyInFunctionName, true), ast.Must(n.ProcedureName()))
	args := ast.ChildrenOfType[*googlesql.ASTTVFArgument](n)
	simple := len(args) <= 2 && allTrue(mapIsSimpleTVFArguments(args))
	pp := p.nest()
	for i, a := range args {
		if i > 0 {
			pp.print(",")
			if !simple {
				pp.println("")
			}
		}
		pp.accept(ctx, a)
	}
	if simple {
		p.print("(" + pp.unnestLeft() + ")")
	} else {
		p.println("(")
		p.incDepth()
		p.print(pp.unnestLeft())
		p.println("")
		p.decDepth()
		p.print(")")
	}
	p.movePast(n)
}

func (p *Printer) VisitCommitStatement(ctx Context, n *googlesql.ASTCommitStatement) {
	p.moveBefore(n)
	p.print(p.keyword("COMMIT TRANSACTION"))
	p.movePast(n)
}

func (p *Printer) VisitExecuteIntoClause(ctx Context, n *googlesql.ASTExecuteIntoClause) {
	p.moveBefore(n)
	p.print(p.keyword("INTO"))
	p.accept(ctx, ast.Must(n.Identifiers()))
}

func (p *Printer) VisitExecuteImmediateStatement(ctx Context, n *googlesql.ASTExecuteImmediateStatement) {
	p.moveBefore(n)
	p.println(p.keyword("EXECUTE IMMEDIATE"))
	p.incDepth()
	// In the future we may try to format the SQL contents when they're
	// a single string containing a valid SQL.
	p.accept(ctx, ast.Must(n.Sql()))
	p.println("")
	p.decDepth()
	p.lnaccept(ctx, ast.Must(n.IntoClause()))
	p.lnaccept(ctx, ast.Must(n.UsingClause()))
}

func (p *Printer) VisitExecuteUsingArgument(ctx Context, n *googlesql.ASTExecuteUsingArgument) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expression()))
	p.accept(ctx, ast.Must(n.Alias()))
}

func (p *Printer) VisitExecuteUsingClause(ctx Context, n *googlesql.ASTExecuteUsingClause) {
	p.moveBefore(n)
	p.println(p.keyword("USING"))
	p.incDepth()
	args := ast.ChildrenOfType[*googlesql.ASTExecuteUsingArgument](n)
	for i, a := range args {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, a)
	}
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitIfStatement(ctx Context, n *googlesql.ASTIfStatement) {
	p.moveBefore(n)
	cond := ast.Must(n.Condition())
	pp := p.nest()
	pp.println("")
	if isSimpleExpr(cond) {
		pp.print(pp.keyword("IF"))
		pp.print(strings.TrimLeft(pp.toUnnestedString(ctx, ast.Must(n.Condition())), "\v"))
		pp.print(pp.keyword("THEN"))
	} else {
		pp.println(pp.keyword("IF"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, ast.Must(n.Condition()))
		pp.println("")
		pp.decDepth()
		pp.print(p.keyword("THEN"))
	}
	p.print(pp.unnestLeft())
	if then := ast.Must(n.ThenList()); ast.Defined(then) && ast.NumChildren(then) > 0 {
		p.println("")
		pp = p.nest()
		pp.incDepth()
		pp.accept(ctx, then)
		pp.println("")
		pp.decDepth()
		p.print(pp.unnestLeft())
	}
	if elseifs := ast.Must(n.ElseifClauses()); ast.Defined(elseifs) {
		p.moveBefore(elseifs)
		for _, e := range ast.ChildrenOfType[*googlesql.ASTElseifClause](elseifs) {
			p.moveBefore(e)
			p.println("")
			pp = p.nest()
			pp.print(pp.keyword("ELSEIF"))
			pp.accept(ctx, ast.Must(e.Condition()))
			pp.print(pp.keyword("THEN"))
			p.println(pp.unnestLeft())
			if body := ast.Must(e.Body()); ast.Defined(body) && ast.NumChildren(body) > 0 {
				pp = p.nest()
				pp.incDepth()
				pp.acceptNestedLeft(ctx, body)
				pp.println("")
				pp.decDepth()
				p.print(pp.unnestLeft())
			}
			p.movePast(e)
		}
		p.movePast(elseifs)
	}
	if e := ast.Must(n.ElseList()); ast.Defined(e) {
		p.println("")
		p.println(p.keyword("ELSE"))
		p.moveBefore(e)
		if ast.NumChildren(e) > 0 {
			pp = p.nest()
			pp.incDepth()
			pp.acceptNestedLeft(ctx, e)
			pp.println("")
			pp.decDepth()
			p.print(pp.unnestLeft())
		}
		p.movePast(e)
	}
	p.println("")
	p.print(p.keyword("END IF"))
}

func (p *Printer) VisitParameterAssignment(ctx Context, n *googlesql.ASTParameterAssignment) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("SET"))
	pp.accept(ctx, ast.Must(n.Parameter()))
	pp.print("=")
	p.print(pp.unnestLeft())
	p.accept(ctx, ast.Must(n.Expression()))
	p.moveBefore(n)
}

func (p *Printer) VisitReturnStatement(ctx Context, n *googlesql.ASTReturnStatement) {
	p.moveBefore(n)
	p.print(p.keyword("RETURN"))
}

func (p *Printer) VisitSystemVariableAssignment(ctx Context, n *googlesql.ASTSystemVariableAssignment) {
	p.moveBefore(n)
	p.print(p.keyword("SET"))
	p.accept(ctx, ast.Must(n.SystemVariable()))
	p.print("=")
	p.accept(ctx, ast.Must(n.Expression()))
	p.movePast(n)
}

func (p *Printer) VisitSingleAssignment(ctx Context, n *googlesql.ASTSingleAssignment) {
	p.moveBefore(n)
	p.print(p.keyword("SET"))
	p.accept(ctx, ast.Must(n.Variable()))
	p.print("=")
	p.accept(ctx.WithValue(KeyInSingleAssignment, true), ast.Must(n.Expression()))
	p.movePast(n)
}

func (p *Printer) VisitVariableDeclaration(ctx Context, n *googlesql.ASTVariableDeclaration) {
	p.moveBefore(n)
	p.print(p.keyword("DECLARE"))
	p.accept(ctx, ast.Must(n.VariableList()))
	p.acceptNested(ctx, ast.Must(n.Type()))
	if dv := ast.Must(n.DefaultValue()); dv != nil {
		p.print(p.keyword("DEFAULT"))
		p.moveBefore(dv)
		p.acceptNested(ctx, ast.Must(n.DefaultValue()))
	}
	p.movePast(n)
}
