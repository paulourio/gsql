package printer

import (
	"fmt"
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitQuery(ctx Context, n *sql.Query) {
	pp := p.nest()
	nestedWith := withInsideWith(n)
	if nestedWith {
		pp.incDepth()
	}
	pp.moveBefore(n)
	pp.printOpenParenIfNeeded(n)
	pp.accept(ctx, n.WithClause())
	q := n.QueryExpr()
	pp.accept(ctx, q)
	if lock := n.LockMode(); lock != nil {
		pp.accept(ctx, lock)
	}
	alignClauses := q.Kind() == sql.SelectKind
	if ob := n.OrderBy(); ob != nil {
		pp.println("")
		if alignClauses {
			pp.accept(ctx, ob)
		} else {
			p1 := pp.nest()
			p1.acceptNestedLeft(ctx, ob)
			pp.print(strings.TrimLeft(p1.unnest(), "\v"))
		}
	}
	if lo := n.LimitOffset(); lo != nil {
		pp.println("")
		if alignClauses {
			pp.accept(ctx, lo)
		} else {
			p1 := pp.nest()
			p1.acceptNestedLeft(ctx, lo)
			pp.print(strings.TrimLeft(p1.unnest(), "\v"))
		}
	}
	if parent := n.Parent(); parent != nil && parent.Kind() != sql.QueryStatementKind {
		pp.movePast(n)
	}
	if nestedWith {
		pp.decDepth()
	}
	pp.printCloseParenIfNeeded(n)
	p.print(pp.unnest())
	for _, op := range n.PipeOperatorList() {
		p.lnaccept(ctx, op)
	}
}

func (p *Printer) visitLockMode(ctx Context, n *sql.LockMode) {
	p.moveBefore(n)
	// TODO: Support FOR SHARE if/when added to parser.
	p.print(p.keyword("FOR UPDATE"))
	p.movePast(n)
}

func (p *Printer) visitQueryStatement(ctx Context, n *sql.QueryStatement) {
	p.moveBefore(n)
	p1 := p.nest()
	p1.accept(ctx, n.Query())
	p.print(p1.unnest())
}

func (p *Printer) visitScript(ctx Context, n *sql.Script) {
	p.accept(ctx, n.StatementList())
}

func (p *Printer) visitSetOperation(ctx Context, n *sql.SetOperation) {
	p.printOpenParenIfNeeded(n)

	indentUnion := true
	for i, query := range n.Inputs() {
		indentUnion = indentUnion && !query.Parenthesized()
		if i > 0 {
			m := n.Metadata().Child(i - 1).(*sql.SetOperationMetadata)
			switch optype := m.OpType().Value(); optype {
			case sql.NotSetSetOp:
				p.print(p.keyword("<UNKNOWN SET OPERATOR>"))
			case sql.UnionOp:
				if indentUnion {
					p.print(" ")
				}
				p.print(p.keyword("UNION"))
			case sql.ExceptOp:
				p.print(p.keyword("EXCEPT"))
			case sql.IntersectOp:
				p.print(p.keyword("INTERSECT"))
			default:
				p.addError(&Error{
					Msg:   fmt.Sprintf("Unknown set operation with code %d", int(optype)),
					Node:  n,
					Input: &p.OriginalInput,
				})
			}
			p.accept(ctx, m.Hint())
			switch settype := m.AllOrDistinct().Value(); settype {
			case sql.All:
				p.print(p.keyword("ALL"))
			case sql.Distinct:
				p.print(p.keyword("DISTINCT"))
			default:
				p.addError(&Error{
					Msg:   fmt.Sprintf("Unknown all or distinct with code %d", int(settype)),
					Node:  query,
					Input: &p.OriginalInput,
				})
			}
			p.println("")
		}
		p.accept(ctx, query)
		p.println("")
	}
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) visitStatementList(ctx Context, n *sql.StatementList) {
	var (
		prev     sql.Node
		prevKind sql.NodeKind
	)
	p.moveBefore(n)
	for i, c := range n.Statements() {
		currKind := c.Kind()
		if i > 0 {
			p.print(";")
			p.movePastLine(prev)
			p.println("")
			if !canGroupStatements(prevKind, currKind) {
				p.println(" ")
			}
		}
		p.acceptNested(ctx, c)
		p.movePast(c)
		prev, prevKind = c, currKind
	}

	num := n.NumChildren()
	parent := n.Parent()
	topLevel := parent == nil || parent.Kind() == sql.ScriptKind
	if num > 1 || (num > 0 && !topLevel) {
		p.println(";")
		if prev != nil {
			p.movePastLine(prev)
		}
	}
	p.movePastLine(n)
}

func withInsideWith(n *sql.Query) bool {
	if n.WithClause() == nil {
		return false
	}
	parent := n.Parent()
	if parent == nil {
		return false
	}
	_, ok := parent.(*sql.WithClauseEntry)
	return ok
}
