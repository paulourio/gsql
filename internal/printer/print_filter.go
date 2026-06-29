package printer

import (
	"fmt"
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitClampedBetweenModifier(ctx Context, n *sql.ClampedBetweenModifier) {
	p.moveBefore(n)
	p.print(p.keyword("CLAMPED BETWEEN"))
	p.accept(ctx, n.Low())
	p.print(p.keyword("AND"))
	p.accept(ctx, n.High())
}

func (p *Printer) visitClusterBy(ctx Context, n *sql.ClusterBy) {
	p.moveBefore(n)
	p.print(p.keyword("CLUSTER"))
	p1 := p.nest()
	p1.print(p1.keyword("BY"))
	printNestedWithSepNode(ctx, p1, n.ClusteringExpressions(), ",")
	p.print(p1.unnest())
}

func (p *Printer) visitCollate(ctx Context, n *sql.Collate) {
	p.moveBefore(n)
	p.print(p.keyword("COLLATE"))
	p.accept(ctx, n.CollationName())
}

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

func (p *Printer) visitCube(ctx Context, n *sql.Cube) {
	p.moveBefore(n)
	p.print(p.keyword("CUBE") + " (")
	printNestedWithSepNode(ctx, p, n.Expressions(), ",")
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitDescriptor(ctx Context, n *sql.Descriptor) {
	p.moveBefore(n)
	p.print(p.keyword("DESCRIPTOR") + "(")
	p.accept(ctx, n.Columns())
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitDescriptorColumn(ctx Context, n *sql.DescriptorColumn) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
}

func (p *Printer) visitDescriptorColumnList(ctx Context, n *sql.DescriptorColumnList) {
	p.moveBefore(n)
	for i, c := range n.Columns() {
		if i > 0 {
			p.print(",")
		}
		p.accept(ctx, c)
	}
	p.movePast(n)
}

func (p *Printer) visitForSystemTime(ctx Context, n *sql.ForSystemTime) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("FOR SYSTEM_TIME AS OF"))
	p.accept(ctx, n.Expression())
}

func (p *Printer) visitFormatClause(ctx Context, n *sql.FormatClause) {
	p.moveBefore(n)
	p.print(p.keyword("FORMAT"))
	p.accept(ctx, n.Format())
	if tz := n.TimeZoneExpr(); tz != nil {
		p.print(p.keyword("AT TIME ZONE"))
		p.accept(ctx, n.TimeZoneExpr())
	}
}

func (p *Printer) visitGroupBy(ctx Context, n *sql.GroupBy) {
	p.moveBefore(n)
	p.accept(ctx, n.Hint())
	p.print(p.keyword("BY"))
	pp := p.nest()
	p.accept(ctx, n.All())
	printNestedWithSepNode(ctx, pp, n.GroupingItems(), ",")
	p.print(pp.unnest())
	s := sql.ParentAs[*sql.Select](n)
	a, _ := sql.LocationRange(
		s.Having(),
		s.Qualify(),
		s.WindowClause(),
	)
	if a > 0 {
		p.moveAt(a)
	}
}

func (p *Printer) visitGroupByAll(_ Context, n *sql.GroupByAll) {
	p.moveBefore(n)
	p.print(p.keyword("ALL"))
	p.movePast(n)
}

func (p *Printer) visitGroupingItem(ctx Context, n *sql.GroupingItem) {
	p.moveBefore(n)
	if n.NumChildren() == 0 {
		p.print("()")
		return
	}
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Rollup())
	p.accept(ctx, n.Cube())
	p.accept(ctx, n.GroupingSetList())
	p.accept(ctx, n.Alias())
	p.accept(ctx, n.GroupingItemOrder())
}

func (p *Printer) visitGroupingSet(ctx Context, n *sql.GroupingSet) {
	p.moveBefore(n)
	if n.NumChildren() == 0 {
		p.print("()")
		return
	}
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Rollup())
	p.accept(ctx, n.Cube())
}

func (p *Printer) visitGroupingSetList(ctx Context, n *sql.GroupingSetList) {
	p.moveBefore(n)
	p.println(p.keyword("GROUPING SETS") + " (")
	pp := p.nest()
	pp.incDepth()
	printlnNestedWithSepNode(ctx, pp, n.GroupingSets(), ",")
	p.print(pp.unnest())
	p.println("")
	p.println(")")
	p.movePast(n)
}

func (p *Printer) visitHaving(ctx Context, n *sql.Having) {
	p.moveBefore(n)
	p.visitMaybeClauseAligned(n.Expression(), ctx)
}

func (p *Printer) visitHavingModifier(ctx Context, n *sql.HavingModifier) {
	p.moveBefore(n)
	p.print(p.keyword("HAVING"))
	switch n.ModifierKind() {
	case sql.NotSetHavingModifier:

	case sql.MinHavingModifier:
		p.print(p.keyword("MIN"))
	case sql.MaxHavingModifier:
		p.print(p.keyword("MAX"))
	}

	p.accept(ctx, n.Expr())
}

func (p *Printer) visitHint(ctx Context, n *sql.Hint) {
	// We use a strings builder here because we don't want automatic
	// spaces between token separators.
	var b strings.Builder
	p.moveBefore(n)
	if shards := n.NumShardsHint(); shards != nil {
		p.print("@")
		p.accept(ctx, shards)
	}
	entries := n.HintEntries()
	if len(entries) > 0 {
		b.WriteString("@{")
		for i, h := range entries {
			if i > 0 {
				p.print(",")
			}
			b.WriteString(p.toString(ctx, h.Name()))
			b.WriteString("=")
			b.WriteString(p.toString(ctx, h.Value()))
		}
		b.WriteString("}")
	}
	p.print(b.String())
}

func (p *Printer) visitHintedStatement(ctx Context, n *sql.HintedStatement) {
	p.accept(ctx, n.Hint())
	p.println("")
	p.accept(ctx, n.Statement())
}

func (p *Printer) visitLimit(ctx Context, n *sql.Limit) {
	p.moveBefore(n)
	if all := n.All(); all != nil {
		p.moveBefore(all)
		p.print(p.keyword("ALL"))
		p.movePast(all)
	}
	p.accept(ctx, n.Expression())
	p.movePast(n)
}

func (p *Printer) visitLimitOffset(ctx Context, n *sql.LimitOffset) {
	p.moveBefore(n)
	p.print(p.keyword("LIMIT"))
	pp := p.nest()
	pp.accept(ctx, n.Limit())
	if os := n.Offset(); os != nil {
		pp.print(p.keyword("OFFSET"))
		pp.accept(ctx, os)
	}
	p.print(pp.unnest())
	p.moveBeforeSuccessorOf(n)
}

func (p *Printer) visitNullOrder(_ Context, n *sql.NullOrder) {
	if n.NullsFirst() {
		p.print(p.keyword("NULLS FIRST"))
	} else {
		p.print(p.keyword("NULLS LAST"))
	}
}

func (p *Printer) visitOnClause(ctx Context, n *sql.OnClause) {
	p1 := p.nest()
	p1.printClause(p1.keyword("ON"))
	p1.moveBefore(n)
	p1.accept(ctx, n.Expression())
	p.print(p1.unnestLeft())
}

func (p *Printer) visitOptionsEntry(ctx Context, n *sql.OptionsEntry) {
	keys := KnownOptionKeys(sql.ParentAs[*sql.OptionsList](n))
	pp := p.nest()
	key := keys.Get(pp.toString(ctx, n.Name()))
	value := pp.toUnnestedString(ctx, n.Value())
	if ctx.Bool(KeySimpleOptions) {
		pp.print(key + "=" + value)
	} else {
		pp.print(key + " \v= " + strings.ReplaceAll(value, "\n", "\n\v"))
	}
	p.print(pp.String())
}

func (p *Printer) visitOptionsList(ctx Context, n *sql.OptionsList) {
	entries := n.OptionsEntries()
	simple := len(entries) <= 1 && allTrue(mapIsSimpleOptionsList(n))
	ctx = ctx.WithValue(KeySimpleOptions, simple)
	pp := p.nest()
	pp.print(pp.keyword("OPTIONS") + " (")
	if !simple {
		pp.println("")
		pp.incDepth()
	}
	pp.moveBefore(n)
	p1 := pp.nest()
	for i, e := range entries {
		if i > 0 {
			p1.print(",")
			if !simple {
				p1.println("")
			}
		}
		p1.accept(ctx, e)
	}
	pp.print(p1.unnestLeft())
	if !simple {
		pp.println("")
		pp.decDepth()
	}
	pp.print(")")
	p.print(pp.unnest())
}

func (p *Printer) visitOrderBy(ctx Context, n *sql.OrderBy) {
	p.moveBefore(n)
	if n.Parent().Kind() == sql.QueryKind {
		p.printClause(p.keyword("ORDER"))
	} else {
		p.print(p.keyword("ORDER"))
	}
	p1 := p.nest()
	p1.accept(ctx, n.Hint())
	p1.print(p1.keyword("BY"))
	p1.moveBefore(n)
	printNestedWithSepNode(ctx, p1, n.OrderingExpressions(), ",")
	p1.moveBeforeSuccessorOf(n)
	p.print(p1.unnestLeft())
}

func (p *Printer) visitOrderingExpression(ctx Context, n *sql.OrderingExpression) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Collate())
	switch n.OrderingSpec() {
	case sql.NotSetSpec:

	case sql.Asc:
		p.print(p.keyword("ASC"))
	case sql.Desc:
		p.print(p.keyword("DESC"))
	}
	p.accept(ctx, n.NullOrder())
	p.movePast(n)
}

func (p *Printer) visitPartitionBy(ctx Context, n *sql.PartitionBy) {
	p.moveBefore(n)
	p.print(p.keyword("PARTITION"))
	p1 := p.nest()
	p1.accept(ctx, n.Hint())
	p1.print(p1.keyword("BY"))
	printNestedWithSepNode(ctx, p1, n.PartitioningExpressions(), ",")
	p.print(p1.unnest())
}

func (p *Printer) visitPivotClause(ctx Context, n *sql.PivotClause) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("PIVOT") + " (")
	p.println("")
	p.incDepth()
	p.acceptNestedLeft(ctx, n.PivotExpressions())
	p.println("")
	pp := p.nest()
	pp.visitPivotForExpression(ctx, n)
	p.print(pp.unnestLeft())
	p.println("")
	p.decDepth()
	p.print(")")
	p.accept(ctx, n.OutputAlias())
	p.movePast(n)
}

func (p *Printer) visitPivotForExpression(ctx Context, n *sql.PivotClause) {
	exprsSimple := mapIsSimplePivotForExpression(n)
	simpleLHS := exprsSimple[0]
	simpleRHS := allTrue(exprsSimple[1:])
	simpleValues := simpleLHS && simpleRHS
	ctx = ctx.WithValue(KeySimplePivotFor, simpleLHS).
		WithValue(KeySimplePivotRHS, simpleRHS).
		WithValue(KeySimplePivotValues, simpleValues)
	p.print(p.keyword("FOR"))
	if !simpleLHS {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, n.ForExpression())
	if !simpleLHS {
		p.println("")
		p.decDepth()
	}
	p.print(p.keyword("IN") + " (")
	if !simpleValues {
		p.println("")
		p.incDepth()
	}
	p.acceptNestedLeft(ctx, n.Child(2).(*sql.PivotValueList))
	if !simpleValues {
		p.println("")
		p.decDepth()
	}
	p.print(")")
}

func (p *Printer) visitPivotExpression(ctx Context, n *sql.PivotExpression) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Alias())
	p.movePast(n)
}

func (p *Printer) visitPivotExpressionList(ctx Context, n *sql.PivotExpressionList) {
	p.moveBefore(n)
	simple := allTrue(mapIsSimplePivotExpressionList(n))
	for i, e := range n.Expressions() {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.acceptNested(ctx, e)
	}
	p.movePast(n)
}

func (p *Printer) visitPivotValue(ctx Context, n *sql.PivotValue) {
	p.moveBefore(n)
	p.accept(ctx, n.Value())
	p.accept(ctx, n.Alias())
	p.movePast(n)
}

func (p *Printer) visitPivotValueList(ctx Context, n *sql.PivotValueList) {
	p.moveBefore(n)
	simple := ctx.Bool(KeySimplePivotValues)
	for i, v := range n.Values() {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, v)
	}
	p.movePast(n)
}

func (p *Printer) visitQualify(ctx Context, n *sql.Qualify) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.visitMaybeClauseAligned(n.Expression(), ctx)
	pp.moveBeforeSuccessorOf(n)
	p.print(pp.unnest())
}

func (p *Printer) visitRepeatableClause(ctx Context, n *sql.RepeatableClause) {
	p.moveBefore(n)
	p.print(p.keyword("REPEATABLE"))
	p.print("(")
	p.accept(ctx, n.Argument())
	p.print(")")
}

func (p *Printer) visitRollup(ctx Context, n *sql.Rollup) {
	p.moveBefore(n)
	p.print(p.keyword("ROLLUP") + " (")
	for i, expr := range n.Expressions() {
		if i > 0 {
			p.print(",")
		}
		p.accept(ctx, expr)
	}
	p.print(")")
}

func (p *Printer) visitSampleClause(ctx Context, n *sql.SampleClause) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("TABLESAMPLE"))
	p.print(p.keyword(p.toString(ctx, n.SampleMethod())) + " ")
	p.print("(")
	p.accept(ctx, n.SampleSize())
	p.print(")")
	p.accept(ctx, n.SampleSuffix())
}

func (p *Printer) visitSampleSize(ctx Context, n *sql.SampleSize) {
	p.moveBefore(n)
	p.accept(ctx, n.Size())
	switch n.Unit() {
	case sql.NotSetUnit:

	case sql.RowsSampleSize:
		p.print(p.keyword("ROWS"))
	case sql.Percent:
		p.print(p.keyword("PERCENT"))
	}
	p.accept(ctx, n.PartitionBy())
}

func (p *Printer) visitSampleSuffix(ctx Context, n *sql.SampleSuffix) {
	p.moveBefore(n)
	p.accept(ctx, n.Weight())
	p.accept(ctx, n.Repeat())
}

func (p *Printer) visitUnpivotClause(ctx Context, n *sql.UnpivotClause) {
	p.moveBefore(n)
	p.println("")
	switch n.NullFilter() {
	case sql.UnspecifiedNullFilter:
		p.println(p.keyword("UNPIVOT") + " (")
	case sql.IncludeNullFilter:
		p.println(p.keyword("UNPIVOT INCLUDE NULLS") + " (")
	case sql.ExcludeNullFilter:
		p.println(p.keyword("UNPIVOT EXCLUDE NULLS") + " (")
	}
	inItems := n.UnpivotInItems()
	simple := allTrue(mapIsSimpleUnpivotInItemList(inItems))
	ctx = ctx.WithValue(KeySimpleUnpivotInTime, simple)
	p.incDepth()
	p.accept(ctx, n.UnpivotOutputValueColumns())
	p.println("")
	p.print(p.keyword("FOR"))
	p.accept(ctx, n.UnpivotOutputNameColumn())
	p.print(p.keyword("IN") + " (")
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, n.UnpivotInItems())
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.println(")")
	p.decDepth()
	p.print(")")
	p.accept(ctx, n.OutputAlias())
	p.movePast(n)
}

func (p *Printer) visitUnpivotInItem(ctx Context, n *sql.UnpivotInItem) {
	p.moveBefore(n)
	p.accept(ctx, n.UnpivotColumns())
	p.accept(ctx, n.Alias())
	p.movePast(n)
}

func (p *Printer) visitUnpivotInItemLabel(ctx Context, n *sql.UnpivotInItemLabel) {
	p.moveBefore(n)
	if label := n.Label(); label != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, n.Label())
	}
	p.movePast(n)
}

func (p *Printer) visitUnpivotInItemList(ctx Context, n *sql.UnpivotInItemList) {
	p.moveBefore(n)
	simple := ctx.Bool(KeySimpleUnpivotInTime)
	for i, item := range n.InItems() {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, item)
	}
	p.movePast(n)
}

func (p *Printer) visitUsingClause(ctx Context, n *sql.UsingClause) {
	p.moveBefore(n)
	p.printClause(p.keyword("USING") + " (")
	printNestedWithSepNode(ctx, p, n.Keys(), ",")
	p.print(")")
}

func (p *Printer) visitWhereClause(ctx Context, n *sql.WhereClause) {
	p.moveBefore(n)
	p.print(p.keyword("WHERE"))
	e := n.Expression()
	switch e.Kind() {
	case sql.AndExprKind, sql.OrExprKind:
		ctx = ctx.WithValue(KeyAlignBinaryOpBudget, 1)
		p.accept(ctx, e)
	default:
		p.acceptNested(ctx, e)
	}
}

func (p *Printer) visitWindowClause(ctx Context, n *sql.WindowClause) {
	p.moveBefore(n)
	for i, w := range n.Windows() {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, w.Name())
		p.print(p.keyword("AS") + " ")
		ws := w.WindowSpec()
		count := countWindowSpecElems(ws)
		p.print("(")
		if count > 0 {
			p.println("")
			p.incDepth()
			p.acceptNestedLeft(ctx, ws)
			p.println("")
			p.decDepth()
		}
		p.print(")")
	}
	p.moveBeforeSuccessorOf(n)
}

func (p *Printer) visitWindowFrame(ctx Context, n *sql.WindowFrame) {
	p.moveBefore(n)
	switch n.FrameUnit() {
	case sql.RowsFrameUnit:
		p.print(p.keyword("ROWS"))
	case sql.Range:
		p.print(p.keyword("RANGE"))
	default:
		p.addError(&Error{
			Msg: fmt.Sprintf("Unknown frame unit id %d", n.FrameUnit()),
		})
	}
	pp := p.nest()
	if n.EndExpr() != nil {
		pp.print(p.keyword("BETWEEN"))
		pp.accept(ctx, n.StartExpr())
		pp.print(p.keyword("AND"))
		pp.accept(ctx, n.EndExpr())
	} else {
		pp.accept(ctx, n.StartExpr())
	}
	p.print(pp.unnest())
}

func (p *Printer) visitWindowFrameExpr(ctx Context, n *sql.WindowFrameExpr) {
	p.moveBefore(n)

	switch n.BoundaryType() {
	case sql.UnboundedPreceding:
		p.print(p.keyword("UNBOUNDED PRECEDING"))
	case sql.OffsetPreceding:
		p.acceptNestedLeft(ctx, n.Expression())
		p.print(p.keyword("PRECEDING"))
	case sql.CurrentRow:
		p.print(p.keyword("CURRENT ROW"))
	case sql.OffsetFollowing:
		p.acceptNestedLeft(ctx, n.Expression())
		p.print(p.keyword("FOLLOWING"))
	case sql.UnboundedFollowing:
		p.print(p.keyword("UNBOUNDED FOLLOWING"))
	}
}

func (p *Printer) visitWindowSpecification(ctx Context, n *sql.WindowSpecification) {
	forceAcrossLines := true
	p.moveBefore(n)
	pp := p.nest()
	wn := n.BaseWindowName()
	if wn != nil {
		pp.accept(ctx, wn)
		if forceAcrossLines {
			pp.println("")
		}
	}
	pp2 := pp.nest()
	pb := n.PartitionBy()
	if pb != nil {
		pp2.accept(ctx, pb)
	}
	ob := n.OrderBy()
	if ob != nil {
		if forceAcrossLines && pb != nil {
			pp2.println("")
		}
		pp2.accept(ctx, ob)
	}
	if wf := n.WindowFrame(); wf != nil {
		if forceAcrossLines && (pb != nil || ob != nil) {
			pp2.println("")
		}
		pp2.accept(ctx, wf)
	}
	pp.print(pp2.unnest())
	pp.movePast(n)
	p.print(pp.unnest())
}

func (p *Printer) visitWithOffset(ctx Context, n *sql.WithOffset) {
	p.moveBefore(n)
	p.print(p.keyword("WITH OFFSET"))
	switch n.Parent().Kind() {
	case sql.DeleteStatementKind, sql.UpdateStatementKind:
		p.print(p.keyword("AS"))
	}
	p.accept(ctx, n.Alias())
}

func (p *Printer) visitWithWeight(ctx Context, n *sql.WithWeight) {
	p.moveBefore(n)
	p.print(p.keyword("WITH WEIGHT"))
	p.accept(ctx, n.Alias())
}
