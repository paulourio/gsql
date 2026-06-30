package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitAssertRowsModified(ctx Context, n *sql.AssertRowsModified) {
	p.moveBefore(n)
	p.print(p.keyword("ASSERT_ROWS_MODIFIED"))
	p.acceptNestedLeft(ctx, n.NumRows())
	p.movePast(n)
}

func (p *Printer) visitAssignmentFromStruct(ctx Context, n *sql.AssignmentFromStruct) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("SET") + " (")
	pp.print(pp.toString(ctx, n.Variables()) + ")")
	pp.print("=")
	p.print(pp.unnestLeft())
	p.accept(ctx, n.StructExpression())
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

func (p *Printer) visitExportDataStatement(ctx Context, n *sql.ExportDataStatement) {
	p.moveBefore(n)
	p.print(p.keyword("EXPORT DATA"))
	p.lnaccept(ctx, n.WithConnectionClause())
	p.lnaccept(ctx, n.OptionsList())
	p.println("")
	p.println(p.keyword("AS") + " (")
	p1 := p.nest()
	p1.incDepth()
	p1.accept(ctx, n.Query())
	p.println(p1.unnest())
	p.println(")")
	p.movePast(n)
}

func (p *Printer) visitExportModelStatement(ctx Context, n *sql.ExportModelStatement) {
	p.moveBefore(n)
	p.print(p.keyword("EXPORT MODEL") + " ")
	p.accept(ctx, n.ModelNamePath())
	p.lnaccept(ctx, n.WithConnectionClause())
	p.lnaccept(ctx, n.OptionsList())
	p.movePast(n)
}

func (p *Printer) visitInsertStatement(ctx Context, n *sql.InsertStatement) {
	cl := n.ColumnList()
	pp := p.nest()
	pp.moveBefore(n)
	begin := n.LocationStart()
	end := n.TargetPath().LocationStart()
	input := strings.ToUpper(pp.viewErasedInput(begin, end))
	if strings.Contains(input, "INTO") {
		pp.print(pp.keyword("INSERT INTO"))
	} else {
		pp.print(pp.keyword("INSERT"))
	}
	pp.accept(ctx.WithValue(KeyInTableName, true), n.TargetPath())
	if cl != nil {
		pp.println("")
		pp.incDepth()
		pp.accept(ctx, cl)
		pp.println("")
		pp.decDepth()
	}
	if q := n.Query(); q != nil {
		pp.println("")
		pp.acceptNested(ctx, q)
	} else {
		pp.println("")
		pp.println(pp.keyword("VALUES"))
		pp.incDepth()
		pp.accept(ctx, n.Rows())
		pp.println("")
		pp.decDepth()
	}
	p.print(pp.unnest())
	if a := n.AssertRowsModified(); a != nil {
		p.println("")
		p.acceptNestedLeft(ctx, a)
	}
	if r := n.Returning(); r != nil {
		p.println("")
		p.acceptNestedLeft(ctx, r)
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

func (p *Printer) visitMergeActionDelete(_ Context, _ *sql.MergeAction) {
	p.println(p.keyword("DELETE"))
}

func (p *Printer) visitMergeAction(ctx Context, n *sql.MergeAction) {
	p.moveBefore(n)
	switch n.ActionType() {
	case sql.NotSetMergeAction:

	case sql.InsertAction:
		p.visitMergeActionInsert(ctx, n)
	case sql.UpdateMergeAction:
		p.visitMergeActionUpdate(ctx, n)
	case sql.DeleteAction:
		p.visitMergeActionDelete(ctx, n)
	}
	p.movePast(n)
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

func (p *Printer) visitMergeWhenClause(ctx Context, n *sql.MergeWhenClause) {
	p.moveBefore(n)
	switch n.MatchType() {
	case sql.NotSetMatchType:

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

func (p *Printer) visitReturningClause(ctx Context, n *sql.ReturningClause) {
	p.moveBefore(n)
	p.print(p.keyword("RETURNING"))
	p.accept(ctx, n.SelectList())
	p.accept(ctx, n.ActionAlias())
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

func (p *Printer) visitUpdateItem(ctx Context, n *sql.UpdateItem) {
	p.moveBefore(n)
	set := n.SetValue()
	if set == nil {
		p.println("(")
		p.incDepth()
	}
	p.accept(ctx, set)
	p.accept(ctx, n.InsertStatement())
	p.accept(ctx, n.DeleteStatement())
	p.accept(ctx, n.UpdateStatement())
	if set == nil {
		p.println("")
		p.decDepth()
		p.print(")")
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

func (p *Printer) visitUpdateSetValue(ctx Context, n *sql.UpdateSetValue) {
	p.moveBefore(n)
	p.accept(ctx, n.Path())
	pp := p.nest()
	pp.print("=")
	pp.acceptNested(ctx, n.Value())
	p.print(pp.unnest())
	p.movePast(n)
}

func (p *Printer) visitUpdateStatement(ctx Context, n *sql.UpdateStatement) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.print(pp.keyword("UPDATE"))
	pp.acceptNestedLeft(ctx.WithValue(KeyInTableName, true), n.TargetPath())
	pp.accept(ctx, n.Alias())
	pp.accept(ctx, n.Offset())
	if list := n.UpdateItemList(); list != nil {
		pp.println("")
		pp.print(pp.keyword("SET"))
		pp.acceptNestedLeft(ctx, list)
	}
	if f := n.FromClause(); f != nil {
		pp.println("")
		pp.print(pp.keyword("FROM"))
		pp.acceptNestedLeft(ctx, f)
	}
	if w := n.Where(); w != nil {
		pp.println("")
		pp.print(pp.keyword("WHERE"))
		pp.acceptNestedLeft(ctx, w)
	}
	p.print(pp.unnest())
	if a := n.AssertRowsModified(); a != nil {
		p.println("")
		p.acceptNestedLeft(ctx, a)
	}
	if r := n.Returning(); r != nil {
		p.println("")
		p.acceptNestedLeft(ctx, r)
	}
	p.movePast(n)
}
