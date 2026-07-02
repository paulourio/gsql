package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitMatchRecognizeClause(ctx Context, n *sql.MatchRecognizeClause) {
	p.moveBefore(n)
	p.print(p.keyword("MATCH_RECOGNIZE") + " (")
	p.println("")
	p.incDepth()

	if n.PartitionBy() != nil {
		p1 := p.nest()
		p1.acceptNestedLeft(ctx, n.PartitionBy())
		p.print(strings.TrimLeft(p1.unnest(), "\v"))
		p.println("")
	}

	if n.OrderBy() != nil {
		p.accept(ctx, n.OrderBy())
		p.println("")
	}

	if n.Measures() != nil {
		p.print(p.keyword("MEASURES") + " ")
		pp := p.nest()
		pp.accept(ctx.WithValue(KeySingleLineCols, true), n.Measures())
		p.print(strings.TrimLeft(pp.unnest(), "\v"))
		p.println("")
	}

	if n.AfterMatchSkipClause() != nil {
		p.accept(ctx, n.AfterMatchSkipClause())
		p.println("")
	}

	p.print(p.keyword("PATTERN") + " (")
	p.accept(ctx, n.Pattern())
	p.print(")")
	p.println("")

	p.print(p.keyword("DEFINE"))
	p.println("")
	p.incDepth()
	p.visitPatternVariableDefinitionList(ctx, n.PatternVariableDefinitionList())
	p.decDepth()

	if n.OptionsList() != nil {
		p1 := p.nest()
		p1.acceptNestedLeft(ctx, n.OptionsList())
		p.println("")
		p.print(strings.TrimLeft(p1.unnest(), "\v"))
	}

	p.decDepth()
	p.println("")
	p.print(")")

	if alias := n.OutputAlias(); alias != nil {
		p.accept(ctx, alias)
	}

	p.movePast(n)
}

func (p *Printer) visitPatternVariableDefinitionList(ctx Context, n *sql.SelectList) {
	if n == nil {
		return
	}
	for i, col := range n.Columns() {
		if i > 0 {
			p.println(",")
		}
		// DEFINE syntax: variable AS expression (reversed from normal SELECT)
		alias := col.Alias()
		if alias != nil {
			p.accept(ctx, alias.Identifier())
			p.print(" " + p.keyword("AS") + " ")
		}
		p.accept(ctx, col.Expression())
	}
}

func (p *Printer) visitAfterMatchSkipClause(ctx Context, n *sql.AfterMatchSkipClause) {
	p.moveBefore(n)
	p.print(p.keyword("AFTER MATCH SKIP "))
	switch n.TargetType() {
	case sql.PastLastRow:
		p.print(p.keyword("PAST LAST ROW"))
	case sql.ToNextRow:
		p.print(p.keyword("TO NEXT ROW"))
	case sql.UnspecifiedSkipTarget:
		p.print(p.keyword("TO "))
		if id := n.SkipToVariable(); id != nil {
			p.accept(ctx, id)
		}
	}
	p.movePast(n)
}

func (p *Printer) visitRowPatternOperation(ctx Context, n *sql.RowPatternOperation) {
	p.moveBefore(n)
	switch n.OpType() {
	case sql.ConcatOp:
		for i, input := range n.Inputs() {
			if i > 0 {
				p.print(" ")
			}
			p.accept(ctx, input)
		}
	case sql.AlternateOp:
		for i, input := range n.Inputs() {
			if i > 0 {
				p.print(" | ")
			}
			p.accept(ctx, input)
		}
	case sql.ExcludeOp:
		p.print("{-")
		for _, input := range n.Inputs() {
			p.accept(ctx, input)
		}
		p.print("-}")
	case sql.PermuteOp:
		p.print(p.keyword("PERMUTE") + "(")
		for i, input := range n.Inputs() {
			if i > 0 {
				p.print(", ")
			}
			p.accept(ctx, input)
		}
		p.print(")")
	case sql.UnspecifiedRowPatternOp:
		for _, input := range n.Inputs() {
			p.accept(ctx, input)
		}
	}
	p.movePast(n)
}

func (p *Printer) visitRowPatternQuantification(ctx Context, n *sql.RowPatternQuantification) {
	p.moveBefore(n)
	p.accept(ctx, n.Operand())
	p.accept(ctx, n.Quantifier())
	p.movePast(n)
}

func (p *Printer) visitSymbolQuantifier(_ Context, n *sql.SymbolQuantifier) {
	p.moveBefore(n)
	switch n.Symbol() {
	case sql.QuestionMark:
		p.print("?")
	case sql.Plus:
		p.print("+")
	case sql.StarSymbol:
		p.print("*")
	case sql.UnspecifiedSymbol:
		// do nothing
	}
	if n.IsReluctant() {
		p.print("?")
	}
	p.movePast(n)
}

func (p *Printer) visitBoundedQuantifier(ctx Context, n *sql.BoundedQuantifier) {
	p.moveBefore(n)
	p.print("{")
	if n.LowerBound() != nil {
		p.accept(ctx, n.LowerBound())
	}
	p.print(",")
	if n.UpperBound() != nil {
		p.print(" ")
		p.accept(ctx, n.UpperBound())
	}
	p.print("}")
	if n.IsReluctant() {
		p.print("?")
	}
	p.movePast(n)
}

func (p *Printer) visitFixedQuantifier(ctx Context, n *sql.FixedQuantifier) {
	p.moveBefore(n)
	p.print("{")
	p.accept(ctx, n.Bound())
	p.print("}")
	if n.IsReluctant() {
		p.print("?")
	}
	p.movePast(n)
}

func (p *Printer) visitQuantifierBound(ctx Context, n *sql.QuantifierBound) {
	p.moveBefore(n)
	p.accept(ctx, n.Bound())
	p.movePast(n)
}

func (p *Printer) visitRowPatternAnchor(_ Context, n *sql.RowPatternAnchor) {
	p.moveBefore(n)
	switch n.Anchor() {
	case sql.Start:
		p.print("^")
	case sql.End:
		p.print("$")
	case sql.UnspecifiedAnchor:
		// do nothing
	}
	p.movePast(n)
}

func (p *Printer) visitRowPatternVariable(ctx Context, n *sql.RowPatternVariable) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
	p.movePast(n)
}
