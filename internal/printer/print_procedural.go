package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitBeginEndBlock(ctx Context, n *sql.BeginEndBlock) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyLabelDeclaration, true), n.Label())
	p.println(p.keyword("BEGIN"))
	p.incDepth()
	p.accept(ctx, n.StatementListNode())
	if n.HasExceptionHandler() {
		p.decDepth()
		p.println(p.keyword("EXCEPTION WHEN ERROR THEN"))
		p.incDepth()
		p.accept(ctx, n.HandlerList())
	}
	p.movePast(n)
	p.println("")
	p.decDepth()
	// Check if block ends with a label.
	if l := n.Label(); l != nil {
		label := l.Name().GetAsString()
		if strings.HasSuffix(p.viewInput(n.Location()), label) {
			p.print(p.keyword("END"))
			p.accept(ctx, l)
			p.println("")
		} else {
			p.println(p.keyword("END"))
		}
	} else {
		p.println(p.keyword("END"))
	}
}

func (p *Printer) visitBeginStatementNode(_ Context, n *sql.BeginStatement) {
	p.moveBefore(n)
	p.println(p.keyword("BEGIN TRANSACTION"))
	p.movePast(n)
}

func (p *Printer) visitBreakStatement(ctx Context, n *sql.BreakStatement) {
	p.moveBefore(n)
	p.print(p.keyword(n.GetKeywordText()))
	p.accept(ctx, n.Label())
	p.movePast(n)
}

func (p *Printer) visitCallStatement(ctx Context, n *sql.CallStatement) {
	p.moveBefore(n)
	p.print(p.keyword("CALL"))
	p.accept(ctx.WithValue(KeyInFunctionName, true), n.ProcedureName())
	args := n.TVFArguments()
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

func (p *Printer) visitCommitStatement(_ Context, n *sql.CommitStatement) {
	p.moveBefore(n)
	p.print(p.keyword("COMMIT TRANSACTION"))
	p.movePast(n)
}

func (p *Printer) visitContinueStatement(ctx Context, n *sql.ContinueStatement) {
	p.moveBefore(n)
	p.print(p.keyword(n.GetKeywordText()))
	p.accept(ctx, n.Label())
	p.movePast(n)
}

func (p *Printer) visitExceptionHandlerNode(ctx Context, n *sql.ExceptionHandler) {
	p.accept(ctx, n.StatementList())
}

func (p *Printer) visitExceptionHandlerListNode(ctx Context, n *sql.ExceptionHandlerList) {
	for _, node := range n.Handlers() {
		p.accept(ctx, node)
	}
}

func (p *Printer) visitExecuteImmediateStatement(ctx Context, n *sql.ExecuteImmediateStatement) {
	p.moveBefore(n)
	p.println(p.keyword("EXECUTE IMMEDIATE"))
	p.incDepth()

	p.accept(ctx, n.SQL())
	p.println("")
	p.decDepth()
	p.lnaccept(ctx, n.IntoClause())
	p.lnaccept(ctx, n.UsingClause())
}

func (p *Printer) visitExecuteIntoClause(ctx Context, n *sql.ExecuteIntoClause) {
	p.moveBefore(n)
	p.print(p.keyword("INTO"))
	p.accept(ctx, n.Identifiers())
}

func (p *Printer) visitExecuteUsingArgument(ctx Context, n *sql.ExecuteUsingArgument) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Alias())
}

func (p *Printer) visitExecuteUsingClause(ctx Context, n *sql.ExecuteUsingClause) {
	p.moveBefore(n)
	p.println(p.keyword("USING"))
	p.incDepth()
	args := n.Arguments()
	for i, a := range args {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, a)
	}
	p.println("")
	p.decDepth()
}

func (p *Printer) visitForInStatement(ctx Context, n *sql.ForInStatement) {
	p.moveBefore(n)
	pp := p.nest()
	pp.accept(ctx.WithValue(KeyLabelDeclaration, true), n.Label())
	pp.print(pp.keyword("FOR"))
	pp.accept(ctx, n.Variable())
	query := n.Query()
	pp.print(pp.keyword("IN") + " (")
	isSimpleQuery := isSimpleQuery(query)
	if !isSimpleQuery {
		pp.println("")
		pp.incDepth()
	}
	pp.acceptNestedLeft(ctx, query)
	if !isSimpleQuery {
		pp.println("")
		pp.decDepth()
	}
	p.print(strings.TrimLeft(pp.unnestLeft(), "\v"))
	if !isSimpleQuery {
		p.println("")
		p.println(")")
	} else {
		p.println(")")
	}
	p.println(p.keyword("DO"))
	p.incDepth()
	p.acceptNestedLeft(ctx, n.Body())
	p.println("")
	p.decDepth()
	// Check if ends with label.
	if l := n.Label(); l != nil {
		label := l.Name().GetAsString()
		if strings.HasSuffix(p.viewInput(n.Location()), label) {
			p.print(p.keyword("END FOR"))
			p.accept(ctx, l)
		} else {
			p.println(p.keyword("END FOR"))
		}
	} else {
		p.println(p.keyword("END FOR"))
	}
	p.movePast(n)
}

func (p *Printer) visitIfStatement(ctx Context, n *sql.IfStatement) {
	p.moveBefore(n)
	cond := n.Condition()
	pp := p.nest()
	pp.println("")
	if isSimpleExpr(cond) {
		pp.print(pp.keyword("IF"))
		pp.print(strings.TrimLeft(pp.toUnnestedString(ctx, n.Condition()), "\v"))
		pp.print(pp.keyword("THEN"))
	} else {
		pp.println(pp.keyword("IF"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, n.Condition())
		pp.println("")
		pp.decDepth()
		pp.print(p.keyword("THEN"))
	}
	p.print(pp.unnestLeft())
	if then := n.ThenList(); then != nil && then.NumChildren() > 0 {
		p.println("")
		pp = p.nest()
		pp.incDepth()
		pp.accept(ctx, then)
		pp.println("")
		pp.decDepth()
		p.print(pp.unnestLeft())
	}
	if elseifs := n.ElseifClauses(); elseifs != nil {
		p.moveBefore(elseifs)
		for _, e := range elseifs.Clauses() {
			p.moveBefore(e)
			p.println("")
			pp = p.nest()
			pp.print(pp.keyword("ELSEIF"))
			pp.accept(ctx, e.Condition())
			pp.print(pp.keyword("THEN"))
			p.println(pp.unnestLeft())
			if body := e.Body(); body != nil && body.NumChildren() > 0 {
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
	if e := n.ElseList(); e != nil {
		p.println("")
		p.println(p.keyword("ELSE"))
		p.moveBefore(e)
		if e.NumChildren() > 0 {
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

func (p *Printer) visitInputTableArgument(ctx Context, n *sql.InputTableArgument) {
	p.moveBefore(n)
	p.print(p.keyword("INPUT TABLE"))
	p.movePast(n)
}

func (p *Printer) visitLabel(ctx Context, n *sql.Label) {
	p.moveBefore(n)
	if ctx.Bool(KeyLabelDeclaration) {
		p.println(p.toString(ctx, n.Name()) + ":")
	} else {
		p.accept(ctx, n.Name())
	}
}

func (p *Printer) visitParameterAssignment(ctx Context, n *sql.ParameterAssignment) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("SET"))
	pp.accept(ctx, n.Parameter())
	pp.print("=")
	p.print(pp.unnestLeft())
	p.accept(ctx, n.Expression())
	p.moveBefore(n)
}

func (p *Printer) visitRaiseStatement(ctx Context, n *sql.RaiseStatement) {
	p.moveBefore(n)

	p.print(p.keyword("RAISE"))
	if m := n.Message(); m != nil {
		p.print("USING MESSAGE =")
		p.accept(ctx, m)
	}
	p.movePast(n)
}

func (p *Printer) visitReturnStatement(_ Context, n *sql.ReturnStatement) {
	p.moveBefore(n)
	p.print(p.keyword("RETURN"))
}

func (p *Printer) visitRepeatStatement(ctx Context, n *sql.RepeatStatement) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyLabelDeclaration, true), n.Label())
	p.println(p.keyword("REPEAT"))
	p.incDepth()
	p.acceptNestedLeft(ctx, n.Body())
	p.println("")
	p.println(" ")
	p.acceptNestedLeft(ctx, n.UntilClause())
	p.println("")
	p.decDepth()
	// Check if ends with label.
	if l := n.Label(); l != nil {
		label := l.Name().GetAsString()
		if strings.HasSuffix(p.viewInput(n.Location()), label) {
			p.print(p.keyword("END REPEAT"))
			p.accept(ctx, l)
			p.println("")
		} else {
			p.println(p.keyword("END REPEAT"))
		}
	} else {
		p.println(p.keyword("END REPEAT"))
	}
}

func (p *Printer) visitUntilClause(ctx Context, n *sql.UntilClause) {
	p.moveBefore(n)
	p.print(p.keyword("UNTIL"))
	p.acceptNestedLeft(ctx, n.Condition())
}

func (p *Printer) visitRollbackStatementNode(_ Context, n *sql.RollbackStatement) {
	p.moveBefore(n)
	p.println(p.keyword("ROLLBACK TRANSACTION"))
	p.movePast(n)
}

func (p *Printer) visitSingleAssignment(ctx Context, n *sql.SingleAssignment) {
	p.moveBefore(n)
	p.print(p.keyword("SET"))
	p.accept(ctx, n.Variable())
	p.print("=")
	p.accept(ctx.WithValue(KeyInSingleAssignment, true), n.Expression())
	p.movePast(n)
}

func (p *Printer) visitSystemVariableAssignment(ctx Context, n *sql.SystemVariableAssignment) {
	p.moveBefore(n)
	p.print(p.keyword("SET"))
	p.accept(ctx, n.SystemVariable())
	p.print("=")
	p.accept(ctx, n.Expression())
	p.movePast(n)
}

func (p *Printer) visitTVFArgument(ctx Context, n *sql.TVFArgument) {
	p.moveBefore(n)
	p.accept(ctx, n.Expr())
	p.accept(ctx, n.TableClause())
	p.accept(ctx, n.ModelClause())
	p.accept(ctx, n.ConnectionClause())
	p.accept(ctx, n.Descriptor())
	p.movePast(n)
}

func (p *Printer) visitVariableDeclaration(ctx Context, n *sql.VariableDeclaration) {
	p.moveBefore(n)
	p.print(p.keyword("DECLARE"))
	p.accept(ctx, n.VariableList())
	p.acceptNested(ctx, n.Type())
	if dv := n.DefaultValue(); dv != nil {
		p.print(p.keyword("DEFAULT"))
		p.moveBefore(dv)
		p.acceptNested(ctx, n.DefaultValue())
	}
	p.movePast(n)
}

func (p *Printer) visitWhileStatement(ctx Context, n *sql.WhileStatement) {
	p.moveBefore(n)
	pp := p.nest()
	pp.accept(ctx.WithValue(KeyLabelDeclaration, true), n.Label())
	whileKeyword := sql.Defined(n.Condition())
	cond := n.Condition()
	simple := isSimpleExpr(cond)
	if whileKeyword {
		p1 := pp.nest()
		p1.print(p1.keyword("WHILE"))
		if !simple {
			p1.println("")
			p1.incDepth()
			p1.acceptNestedLeft(ctx, cond)
			p1.println("")
			p1.decDepth()
		} else {
			p1.acceptNestedLeft(ctx, cond)
		}
		p1.print(p1.keyword("DO"))
		pp.println(strings.TrimLeft(p1.unnestLeft(), "\v"))
	} else {
		pp.println(pp.keyword("LOOP"))
	}
	pp.incDepth()
	pp.acceptNestedLeft(ctx, n.Body())
	pp.println("")
	pp.decDepth()
	var endKeyword string
	if whileKeyword {
		endKeyword = pp.keyword("END WHILE")
	} else {
		endKeyword = pp.keyword("END LOOP")
	}
	// Check if ends with label.
	if l := n.Label(); l != nil {
		label := l.Name().GetAsString()
		if strings.HasSuffix(pp.viewInput(n.Location()), label) {
			pp.print(endKeyword)
			pp.accept(ctx, l)
		} else {
			pp.println(endKeyword)
		}
	} else {
		pp.println(endKeyword)
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}
