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

func (p *Printer) visitPipeInsert(ctx Context, n *sql.PipeInsert) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.acceptNestedLeft(ctx, n.InsertStatement())
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

func (p *Printer) visitPipeWith(ctx Context, n *sql.PipeWith) {
	p.moveBefore(n)
	pp := p.nest()
	pp.lnprint("|>")
	pp.acceptNestedLeft(ctx, n.WithClause())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

// printSetOpKeywords emits the keywords for a set operation from its metadata:
// optional propagation modifier (FULL/LEFT/INNER), operation type
// (UNION/INTERSECT/EXCEPT), optional hint, ALL/DISTINCT quantifier, and
// optional column-match suffix (CORRESPONDING [BY (…)] / BY NAME [ON (…)]).
func (p *Printer) printSetOpKeywords(ctx Context, m *sql.SetOperationMetadata, recursive bool) {
	// Leading propagation modifier (not used in RECURSIVE form).
	if !recursive {
		if pm := m.ColumnPropagationMode(); sql.Defined(pm) {
			switch pm.Value() {
			case sql.FullPropagation:
				begin, end := pm.Location()
				origText := strings.ToUpper(p.viewErasedInput(begin, end))
				if strings.Contains(origText, "OUTER") {
					p.print(p.keyword("FULL OUTER"))
				} else {
					p.print(p.keyword("FULL"))
				}
			case sql.LeftPropagation:
				begin, end := pm.Location()
				origText := strings.ToUpper(p.viewErasedInput(begin, end))
				if strings.Contains(origText, "OUTER") {
					p.print(p.keyword("LEFT OUTER"))
				} else {
					p.print(p.keyword("LEFT"))
				}
			case sql.InnerPropagation:
				p.print(p.keyword("INNER"))
			case sql.Strict:
				// STRICT comes after ALL/DISTINCT, handled below.
			}
		}
	}

	// Set operation type.
	if ot := m.OpType(); sql.Defined(ot) {
		switch ot.Value() {
		case sql.UnionOp:
			p.print(p.keyword("UNION"))
		case sql.ExceptOp:
			p.print(p.keyword("EXCEPT"))
		case sql.IntersectOp:
			p.print(p.keyword("INTERSECT"))
		}
	}

	// Optional hint (e.g. @{hint=1}).
	p.accept(ctx, m.Hint())

	// ALL or DISTINCT.
	if ad := m.AllOrDistinct(); sql.Defined(ad) {
		switch ad.Value() {
		case sql.All:
			p.print(p.keyword("ALL"))
		case sql.Distinct:
			p.print(p.keyword("DISTINCT"))
		}
	}

	// STRICT modifier appears after ALL/DISTINCT.
	if pm := m.ColumnPropagationMode(); sql.Defined(pm) && pm.Value() == sql.Strict {
		p.print(p.keyword("STRICT"))
	}

	// Optional column-match suffix.
	if cm := m.ColumnMatchMode(); sql.Defined(cm) {
		switch cm.Value() {
		case sql.Corresponding:
			p.print(p.keyword("CORRESPONDING"))
		case sql.CorrespondingBy:
			p.print(p.keyword("CORRESPONDING BY"))
			p.accept(ctx, m.CorrespondingByColumnList())
		case sql.ByName:
			p.print(p.keyword("BY NAME"))
		case sql.ByNameOn:
			p.print(p.keyword("BY NAME ON"))
			p.accept(ctx, m.CorrespondingByColumnList())
		}
	}
}

// printPipeSetOpInput prints a single input of a pipe set-operation as an
// indented parenthesised block on p:
//
//	(
//	  <content>
//	)
func (p *Printer) printPipeSetOpInput(ctx Context, input sql.Node) {
	p.println("(")
	p.incDepth()
	pp := p.nest()
	pp.accept(ctx, input)
	p.print(pp.unnest())
	p.println("")
	p.decDepth()
	p.print(")")
}

// visitSubpipeline prints a Subpipeline node (a sequence of pipe operators).
// The surrounding parentheses are handled by the caller (printPipeSetOpInput).
func (p *Printer) visitSubpipeline(ctx Context, n *sql.Subpipeline) {
	p.moveBefore(n)
	for _, op := range n.PipeOperatorList() {
		p.lnaccept(ctx, op)
	}
	p.movePast(n)
}

func (p *Printer) visitPipeSetOperation(ctx Context, n *sql.PipeSetOperation) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	p2 := pp.nest()
	p2.printSetOpKeywords(ctx, n.Metadata(), false)
	for i, input := range n.Inputs() {
		if i > 0 {
			p2.print(",")
		}
		p2.println("")
		if tc, ok := input.(*sql.TableClause); ok {
			// Unparenthesized TABLE clauses print inline (TABLE t);
			// parenthesized ones stay inline as (TABLE t).
			if tc.Parenthesized() {
				p2.print("(")
			}
			p2.print(p2.keyword("TABLE"))
			p2.accept(ctx, tc.TablePath())
			if tc.Parenthesized() {
				p2.print(")")
			}
			continue
		}
		// A parenthesized Query whose only content is a TABLE clause
		// (e.g. (TABLE t5)) prints inline as (TABLE t5).
		if q, ok := input.(*sql.Query); ok && q.Parenthesized() {
			if tc, ok := q.QueryExpr().(*sql.TableClause); ok {
				p2.print("(")
				p2.print(p2.keyword("TABLE"))
				p2.accept(ctx, tc.TablePath())
				p2.print(")")
				continue
			}
		}
		p2.printPipeSetOpInput(ctx, input)
	}
	pp.print(p2.unnestLeft())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeRecursiveUnion(ctx Context, n *sql.PipeRecursiveUnion) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	p2 := pp.nest()
	p2.print(p2.keyword("RECURSIVE"))
	p2.printSetOpKeywords(ctx, n.Metadata(), true)
	// Optional depth modifier (WITH DEPTH …).
	p2.accept(ctx, n.RecursionDepthModifier())
	p2.println("")
	if sub := n.InputSubpipeline(); sql.Defined(sub) {
		p2.printPipeSetOpInput(ctx, sub)
	} else {
		p2.printPipeSetOpInput(ctx, n.InputSubquery())
	}
	pp.print(p2.unnestLeft())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeFork(ctx Context, n *sql.PipeFork) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	pp.print(pp.keyword("FORK"))
	if h := n.Hint(); h != nil {
		pp.print(" ")
		pp.accept(ctx, h)
	}
	subs := n.SubpipelineList()
	if len(subs) == 0 {
		pp.print(" ()")
	} else {
		pp.print(" ")
		p2 := pp.nest()
		for i, sub := range subs {
			if i > 0 {
				p2.print(",")
				p2.println("")
			}
			ops := sub.PipeOperatorList()
			if len(ops) == 0 {
				p2.print("()")
			} else {
				p2.print("(")
				p2.println("")
				p3 := p2.nest()
				p3.incDepth()
				for j, op := range ops {
					if j > 0 {
						p3.println("")
					}
					p3.acceptNestedLeft(ctx, op)
				}
				p2.print(p3.unnestLeft())
				p2.println("")
				p2.print(")")
			}
		}
		pp.print(p2.unnestLeft())
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeTee(ctx Context, n *sql.PipeTee) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	pp.print(pp.keyword("TEE"))
	if h := n.Hint(); h != nil {
		pp.print(" ")
		pp.accept(ctx, h)
	}
	subs := n.SubpipelineList()
	if len(subs) == 0 {
		pp.print(" ()")
	} else {
		pp.print(" ")
		p2 := pp.nest()
		for i, sub := range subs {
			if i > 0 {
				p2.print(",")
				p2.println("")
			}
			ops := sub.PipeOperatorList()
			if len(ops) == 0 {
				p2.print("()")
			} else {
				p2.print("(")
				p2.println("")
				p3 := p2.nest()
				p3.incDepth()
				for j, op := range ops {
					if j > 0 {
						p3.println("")
					}
					p3.acceptNestedLeft(ctx, op)
				}
				p2.print(p3.unnestLeft())
				p2.println("")
				p2.print(")")
			}
		}
		pp.print(p2.unnestLeft())
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeIf(ctx Context, n *sql.PipeIf) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")

	cases := n.IfCases()
	for i, c := range cases {
		if i == 0 {
			pp.print(" ")
			pp.print(pp.keyword("IF"))
			if h := n.Hint(); h != nil {
				pp.print(" ")
				pp.accept(ctx, h)
			}
		} else {
			pp.println("")
			pp.print("   ")
			pp.print(pp.keyword("ELSEIF"))
		}

		pp.print(" ")
		pp.accept(ctx, c.Condition())
		pp.print(" ")
		pp.print(pp.keyword("THEN"))

		p.printSubpipelineBlock(ctx, pp, c.Subpipeline(), "     ")
	}

	if el := n.ElseSubpipeline(); el != nil {
		pp.println("")
		pp.print("   ")
		pp.print(pp.keyword("ELSE"))
		p.printSubpipelineBlock(ctx, pp, el, "     ")
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeLog(ctx Context, n *sql.PipeLog) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	pp.print(" ")
	pp.print(pp.keyword("LOG"))
	if h := n.Hint(); h != nil {
		pp.print(" ")
		pp.accept(ctx, h)
	}

	sub := n.Subpipeline()
	if sub == nil {
		p.print(pp.unnestLeft())
		p.movePast(n)
		return
	}
	ops := sub.PipeOperatorList()
	if len(ops) == 0 {
		pp.print(" ()")
	} else {
		pp.print(" ")
		p2 := pp.nest()
		p2.print("(")
		p2.println("")
		p3 := p2.nest()
		p3.incDepth()
		for j, op := range ops {
			if j > 0 {
				p3.println("")
			}
			p3.acceptNestedLeft(ctx, op)
		}
		p2.print(p3.unnestLeft())
		p2.println("")
		p2.print(")")
		pp.print(p2.unnestLeft())
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

// printSubpipelineBlock prints a subpipeline as an indented parenthesized
// block.  The indent parameter controls the column position of ( and )
// relative to the start of the current line in pp.
//
// Alignment is resolved internally so that multiple blocks printed into
// the same parent printer do not interfere with each other through the
// tabwriter's column tracking.
func (p *Printer) printSubpipelineBlock(ctx Context, pp *Printer, sub *sql.Subpipeline, indent string) {
	ops := sub.PipeOperatorList()
	if len(ops) == 0 {
		pp.println("")
		pp.print(indent + "()")
		return
	}
	// Build the subpipeline content using proper nesting, so that the
	// operators inside are aligned correctly relative to each other.
	p2 := p.nest()
	p2.print("(")
	p2.println("")
	p3 := p2.nest()
	p3.incDepth()
	for j, op := range ops {
		if j > 0 {
			p3.println("")
		}
		p3.acceptNestedLeft(ctx, op)
	}
	p2.print(p3.unnestLeft())
	p2.println("")
	p2.print(")")
	// Resolve the internal alignment (processes \v columns within this
	// block only) and prefix each line with the indent string.
	aligned := leftAlignNested(p2.String())
	pp.println("")
	pp.print(prefixLines(aligned, indent))
}

func (p *Printer) visitPipePivot(ctx Context, n *sql.PipePivot) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	pp.acceptNestedLeft(ctx.WithValue(KeyInPipeOperator, true), n.PivotClause())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeUnpivot(ctx Context, n *sql.PipeUnpivot) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	pp.acceptNestedLeft(ctx.WithValue(KeyInPipeOperator, true), n.UnpivotClause())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitPipeMatchRecognize(ctx Context, n *sql.PipeMatchRecognize) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print("|>")
	// Use acceptNestedLeft because MATCH_RECOGNIZE is a block clause
	pp.acceptNestedLeft(ctx, n.MatchRecognizeClause())
	p.print(pp.unnestLeft())
	p.movePast(n)
}
