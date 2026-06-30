package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitFromQuery(ctx Context, n *sql.FromQuery) {
	p.moveBefore(n)
	pp := p.nest()
	pp.printClause("FROM")
	pp.acceptNestedLeft(ctx, n.FromClause())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitPipeAggregate(ctx Context, n *sql.PipeAggregate) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	p2 := pp.nest()
	p2.print(p2.keyword("AGGREGATE"))
	p2.accept(ctx, n.WithModifier())
	p2.accept(ctx, n.Select())
	pp.print(p2.unnestLeft())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeAggregateSelect(ctx Context, n *sql.Select) {
	p.moveBefore(n)
	singleLine := p.maybeSingleLineColumns(n)
	ctx = ctx.WithValue(KeySingleLineCols, singleLine)
	if !singleLine {
		p.println("")
		p.incDepth()
	}
	pp := p.nest()
	pp.accept(ctx, n.SelectList())
	p.print(pp.unnest())
	if !singleLine {
		p.decDepth()
	}
	if gb := n.GroupBy(); gb != nil {
		p.moveBefore(gb)
		p.println("")
		pp := p.nest()
		pp.print(pp.keyword("GROUP") + " ")
		pp.acceptNestedLeft(ctx, gb)
		p.print(strings.TrimLeft(pp.unnestLeft(), "\v"))
	}
	p.movePastLine(n)
}

func (p *Printer) visitPipeAs(ctx Context, n *sql.PipeAs) {
	p.moveBefore(n)
	p.lnprint("|>")
	p.accept(ctx, n.Alias())
	p.movePast(n)
}

func (p *Printer) visitPipeAssert(ctx Context, n *sql.PipeAssert) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.print(pp.keyword("ASSERT"))
	pp.acceptNestedLeft(ctx, n.Condition())
	payloads := n.MessageList()
	for _, arg := range payloads {
		pp.print(",")
		pp.accept(ctx, arg)
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeCall(ctx Context, n *sql.PipeCall) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.print(pp.keyword("CALL"))
	pp.acceptNestedLeft(ctx, n.TVF())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeCreateTable(ctx Context, n *sql.PipeCreateTable) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.acceptNestedLeft(ctx, n.CreateTableStatement())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeDrop(ctx Context, n *sql.PipeDrop) {
	p.moveBefore(n)
	p.print("|>")
	p.print(p.keyword("DROP"))
	cols := n.ColumnList().IdentifierList()
	for i, col := range cols {
		if i > 0 {
			p.print(",")
		}
		p.accept(ctx, col)
	}
	p.movePast(n)
}

func (p *Printer) visitPipeJoin(ctx Context, n *sql.PipeJoin) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.visitPipeJoinJoin(ctx, n.Join())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeJoinJoin(ctx Context, n *sql.Join) {
	count, _ := ctx.Int(KeyJoinCounts)
	pp := p.nest()
	switch n.JoinType() {
	case sql.Comma:
		pp.print(",")
	case sql.DefaultJoinType, sql.Cross, sql.FullJoin,
		sql.InnerJoin, sql.LeftJoin, sql.RightJoin:
		if count >= p.Writer.opts.MinJoinsToSeparateInBlocks {
			pp.println("")
		}
		pp.moveBefore(n)
		pp.moveBefore(n.JoinLocation())
		pp.println("\v")
		pp.print(p.keyword(p.joinKeyword(n)))
	}
	pp.accept(ctx, n.Hint())
	pp.println("")
	pp2 := p.nest()
	pp2.acceptNestedLeft(ctx, n.RHS())
	pp2.movePast(n.RHS())
	pp.print(pp2.unnest())
	if oc := n.OnClause(); oc != nil {
		pp.println("")
		pp.acceptNested(ctx, oc)
	}
	if uc := n.UsingClause(); uc != nil {
		pp.println("")
		pp.acceptNested(ctx, uc)
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitPipeSelect(ctx Context, n *sql.PipeSelect) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.acceptNestedLeft(ctx, n.Select())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeDistinct(_ Context, n *sql.PipeDistinct) {
	p.moveBefore(n)
	p.lnprint("|> DISTINCT")
	p.movePast(n)
}

func (p *Printer) visitPipeWhere(ctx Context, n *sql.PipeWhere) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.acceptNestedLeft(ctx, n.Where())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeExtend(ctx Context, n *sql.PipeExtend) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.acceptNestedLeft(ctx, n.Select())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeExtendSelect(ctx Context, n *sql.Select) {
	p.moveBefore(n)
	p.print(p.keyword("EXTEND"))
	singleLine := p.maybeSingleLineColumns(n)
	ctx = ctx.WithValue(KeySingleLineCols, singleLine)
	pp := p.nest()
	pp.accept(ctx, n.SelectList())
	p.print(pp.unnestLeft())
	if win := n.WindowClause(); win != nil {
		p.moveBefore(win)
		p.println("")
		pp := p.nest()
		pp.print(pp.keyword("WINDOW") + " ")
		pp.acceptNestedLeft(ctx, win)
		p.print(strings.TrimLeft(pp.unnestLeft(), "\v"))
	}
	p.movePastLine(n)
}

func (p *Printer) visitPipeSelectSelect(ctx Context, n *sql.Select) {
	p.moveBefore(n)
	p.print(p.keyword("SELECT"))
	
	pp := p.nest()
	if h := n.Hint(); h != nil {
		pp.accept(ctx, h)
	}
	
	singleLine := p.maybeSingleLineColumns(n)
	ctx = ctx.WithValue(KeySingleLineCols, singleLine)

	if n.Distinct() {
		pp.print(pp.keyword("DISTINCT"))
	}
	if sa := n.SelectAs(); sa != nil {
		pp.accept(ctx, sa)
	}
	if wm := n.WithModifier(); wm != nil {
		pp.accept(ctx, wm)
	}
	
	if n.Hint() != nil || n.WithModifier() != nil {
		pp.println("")
	}
	
	pp.accept(ctx, n.SelectList())
	p.print(pp.unnestLeft())
	
	if win := n.WindowClause(); win != nil {
		p.moveBefore(win)
		p.println("")
		pp := p.nest()
		pp.print(pp.keyword("WINDOW") + " ")
		pp.acceptNestedLeft(ctx, win)
		p.print(strings.TrimLeft(pp.unnestLeft(), "\v"))
	}
	p.movePastLine(n)
}

func (p *Printer) visitPipeOrderBy(ctx Context, n *sql.PipeOrderBy) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|> ")
	pp.acceptNestedLeft(ctx, n.OrderBy())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeLimitOffset(ctx Context, n *sql.PipeLimitOffset) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.acceptNestedLeft(ctx, n.LimitOffset())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeStaticDescribe(ctx Context, n *sql.PipeStaticDescribe) {
	p.moveBefore(n)
	p.lnprint("|> " + p.keyword("STATIC_DESCRIBE"))
	p.movePast(n)
}

func (p *Printer) visitPipeDescribe(ctx Context, n *sql.PipeDescribe) {
	p.moveBefore(n)
	p.lnprint("|> " + p.keyword("DESCRIBE"))
	p.movePast(n)
}

func (p *Printer) visitPipeRename(ctx Context, n *sql.PipeRename) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.print(pp.keyword("RENAME"))
	items := n.RenameItemList()
	for i, item := range items {
		if i > 0 {
			pp.print(",")
		}
		pp.acceptNestedLeft(ctx, item)
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeRenameItem(ctx Context, n *sql.PipeRenameItem) {
	p.accept(ctx, n.OldName())
	p.print(" " + p.keyword("AS") + " ")
	p.accept(ctx, n.NewName())
}

func (p *Printer) visitPipeSet(ctx Context, n *sql.PipeSet) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.print(pp.keyword("SET"))
	items := n.SetItemList()
	for i, item := range items {
		if i > 0 {
			pp.print(",")
		}
		pp.acceptNestedLeft(ctx, item)
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeSetItem(ctx Context, n *sql.PipeSetItem) {
	p.accept(ctx, n.Column())
	p.print(" = ")
	p.accept(ctx, n.Expression())
}

func (p *Printer) visitPipeTablesample(ctx Context, n *sql.PipeTablesample) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|> ")
	pp.acceptNestedLeft(ctx, n.Sample())
	p.print(pp.unnestLeft())
	p.movePast(n)
}
