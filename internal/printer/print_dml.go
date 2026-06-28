// This file contains printing functions specific for
// Data Definition Language (DDL).
package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitColumnList(ctx Context, n *sql.ColumnList) {
	p.moveBefore(n)
	cols := n.Identifiers()
	simple := len(cols) <= 4
	p.print("(")
	if !simple {
		p.println("")
		p.incDepth()
	}
	for i, c := range cols {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, c)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitDeleteStatement(ctx Context, n *sql.DeleteStatement) {
	pp := p.nest()
	pp.moveBefore(n)
	p1 := pp.nest()
	begin := n.LocationStart()
	end := n.TargetPath().LocationStart()
	input := strings.ToUpper(p1.viewErasedInput(begin, end))
	if strings.Contains(input, "FROM") {
		p1.print(p1.keyword("DELETE \vFROM"))
	} else {
		p1.print(p1.keyword("DELETE") + " \v")
	}
	p1.accept(ctx.WithValue(KeyInTableName, true), n.TargetPath())
	p1.accept(ctx, n.Hint())
	p1.accept(ctx, n.Alias())
	p1.accept(ctx, n.Offset())
	p1.println("")
	if w := n.Where(); w != nil {
		p1.print(p1.keyword("WHERE"))
		p1.acceptNestedLeft(ctx, w)
		p1.println("")
	}
	pp.print(p1.unnest())
	if a := n.AssertRowsModified(); a != nil {
		pp.println("")
		pp.acceptNestedLeft(ctx, a)
	}
	if r := n.Returning(); r != nil {
		pp.println("")
		pp.acceptNestedLeft(ctx, r)
	}
	pp.movePast(n)
	p.print(pp.unnestLeft())
}

func (p *Printer) visitReturningClause(ctx Context, n *sql.ReturningClause) {
	p.moveBefore(n)
	p.print(p.keyword("RETURNING"))
	p.accept(ctx, n.SelectList())
	p.accept(ctx, n.ActionAlias())
	p.movePast(n)
}

func (p *Printer) visitAssertRowsModified(ctx Context, n *sql.AssertRowsModified) {
	p.moveBefore(n)
	p.print(p.keyword("ASSERT_ROWS_MODIFIED"))
	p.acceptNestedLeft(ctx, n.NumRows())
	p.movePast(n)
}

func (p *Printer) visitInsertStatement(ctx Context, n *sql.InsertStatement) {
	cl := n.ColumnList()
	p.moveBefore(n)
	begin := n.LocationStart()
	end := n.TargetPath().LocationStart()
	input := strings.ToUpper(p.viewErasedInput(begin, end))
	if strings.Contains(input, "INTO") {
		p.print(p.keyword("INSERT INTO"))
	} else {
		p.print(p.keyword("INSERT"))
	}
	p.accept(ctx.WithValue(KeyInTableName, true), n.TargetPath())
	if cl != nil {
		p.println("")
		p.incDepth()
		p.accept(ctx, cl)
		p.println("")
		p.decDepth()
	}
	if q := n.Query(); q != nil {
		p.println("")
		p.acceptNested(ctx, q)
	} else {
		p.println("")
		p.println(p.keyword("VALUES"))
		p.incDepth()
		p.accept(ctx, n.Rows())
		p.println("")
		p.decDepth()
	}

	p.movePast(n)
}

func (p *Printer) visitInsertValuesRowList(ctx Context, n *sql.InsertValuesRowList) {
	p.moveBefore(n)
	for i, r := range n.Rows() {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, r)
	}
	p.movePast(n)
}

func (p *Printer) visitInsertValuesRow(ctx Context, n *sql.InsertValuesRow) {
	p.moveBefore(n)
	values := n.Values()
	simple := len(values) <= 4 && allTrue(mapIsSimpleExprs(values))
	p.print("(")
	if !simple {
		p.println("")
		p.incDepth()
	}
	for i, r := range values {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, r)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitMergeStatement(ctx Context, n *sql.MergeStatement) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.print(pp.keyword("MERGE") + " ")
	p1 := pp.nest()
	p1.print(p1.keyword("INTO"))
	p1.acceptNested(ctx.WithValue(KeyInTableName, true), n.TargetPath())
	pp.print(p1.unnest())
	pp.accept(ctx, n.Alias())
	pp.println("")
	pp.print(pp.keyword("USING") + " ")
	pp.acceptNested(ctx.WithValue(KeyInTableName, true), n.TableExpression())
	pp.println("")
	pp.print(pp.keyword("ON") + " ")
	pp.acceptNested(ctx, n.MergeCondition())
	p.println("")
	pp.accept(ctx, n.WhenClauses())
	pp.movePast(n)
	p.print(pp.unnest())
}

func (p *Printer) visitMergeAction(ctx Context, n *sql.MergeAction) {
	p.moveBefore(n)
	switch n.ActionType() {
	case sql.NotSetMergeAction:
		// Nothing.
	case sql.InsertAction:
		p.visitMergeActionInsert(ctx, n)
	case sql.UpdateMergeAction:
		p.visitMergeActionUpdate(ctx, n)
	case sql.DeleteAction:
		p.visitMergeActionDelete(ctx, n)
	}
	p.movePast(n)
}

func (p *Printer) visitMergeActionDelete(_ Context, _ *sql.MergeAction) {
	p.println(p.keyword("DELETE"))
}

func (p *Printer) visitMergeActionInsert(ctx Context, n *sql.MergeAction) {
	cl := n.InsertColumnList()
	ir := n.InsertRow()
	if cl == nil && ir != nil && ir.NumChildren() == 0 {
		p.println(p.keyword("INSERT ROW"))
		return
	}
	p.println(p.keyword("INSERT"))
	if cl != nil {
		p.incDepth()
		p.accept(ctx, cl)
		p.println("")
		p.decDepth()
	}
	if ir != nil {
		p.println("")
		p.println(p.keyword("VALUES"))
		p.incDepth()
		p.accept(ctx, ir)
		p.println("")
		p.decDepth()
	}
}

func (p *Printer) visitMergeActionUpdate(ctx Context, n *sql.MergeAction) {
	p.println(p.keyword("UPDATE SET"))
	p.incDepth()
	p.accept(ctx, n.UpdateItemList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitMergeWhenClauseList(ctx Context, n *sql.MergeWhenClauseList) {
	p.moveBefore(n)
	for _, c := range n.Clauses() {
		p.println("")
		p.print(p.keyword("WHEN") + " ")
		p.acceptNested(ctx, c)
	}
	p.println("")
	p.movePast(n)
}

func (p *Printer) visitMergeWhenClause(ctx Context, n *sql.MergeWhenClause) {
	p.moveBefore(n)
	switch n.MatchType() {
	case sql.NotSetMatchType:
		// Nothing.
	case sql.Matched:
		p.print(p.keyword("MATCHED"))
	case sql.NotMatchedBySource:
		p.print(p.keyword("NOT MATCHED BY SOURCE"))
	case sql.NotMatchedByTarget:
		start := n.LocationStart()
		next := n.Action().LocationStart()
		input := strings.ToUpper(p.viewErasedInput(start, next))
		if strings.Contains(input, "TARGET") {
			p.print(p.keyword("NOT MATCHED BY TARGET"))
		} else {
			p.print(p.keyword("NOT MATCHED"))
		}
	}
	if cond := n.SearchCondition(); cond != nil {
		if isSimpleExpr(cond) {
			p.print(p.keyword("AND"))
			p.accept(ctx, cond)
		} else {
			p.print(p.keyword("AND"))
			p.println("")
			p.incDepth()
			p.accept(ctx, cond)
			p.println("")
			p.decDepth()
		}
	}
	p.println(p.keyword("THEN"))
	p.accept(ctx, n.Action())
	p.movePast(n)
}

func (p *Printer) visitTruncateStatement(ctx Context, n *sql.TruncateStatement) {
	p.moveBefore(n)
	p.print(p.keyword("TRUNCATE TABLE"))
	p.accept(ctx.WithValue(KeyInTableName, true), n.TargetPath())
	if w := n.Where(); w != nil {
		p.println("")
		p.print(p.keyword("WHERE"))
		p.accept(ctx, w)
	}
	p.movePast(n)
}

func (p *Printer) visitUpdateItemList(ctx Context, n *sql.UpdateItemList) {
	p.moveBefore(n)
	pp := p.nest()
	items := n.Items()
	for i, item := range items {
		if i > 0 {
			pp.println(",")
		}
		pp.accept(ctx, item)
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitUpdateItem(ctx Context, n *sql.UpdateItem) {
	p.moveBefore(n)
	p.accept(ctx, n.SetValue())
	p.accept(ctx, n.InsertStatement())
	p.accept(ctx, n.DeleteStatement())
	p.accept(ctx, n.UpdateStatement())
	p.movePast(n)
}

func (p *Printer) visitUpdateSetValue(ctx Context, n *sql.UpdateSetValue) {
	p.moveBefore(n)
	p.accept(ctx, n.Path())
	pp := p.nest()
	pp.print("=")
	pp.acceptNested(ctx, n.Value())
	p.print(pp.unnest())
	p.movePast(n)
}
