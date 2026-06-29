package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitFromQuery(ctx Context, n *sql.FromQuery) {
	p.moveBefore(n)
	p.printClause("FROM")
	p.acceptNestedLeft(ctx, n.FromClause())
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
}
