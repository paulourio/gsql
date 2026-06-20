// This file contains printing functions specific for
// Data Definition Language (DDL).
package printer

import (
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

func (p *Printer) VisitColumnList(ctx Context, n *googlesql.ASTColumnList) {
	p.moveBefore(n)
	cols := ast.ChildrenOfType[*googlesql.ASTIdentifier](n)
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

func (p *Printer) VisitInsertStatement(ctx Context, n *googlesql.ASTInsertStatement) {
	cl := ast.Must(n.ColumnList())
	p.moveBefore(n)
	begin := ast.GetParseLocationStartOffset(n)
	end := ast.GetParseLocationStartOffset(ast.Must(n.TargetPath()))
	input := strings.ToUpper(p.viewErasedInput(begin, end))
	if strings.Contains(input, "INTO") {
		p.print(p.keyword("INSERT INTO"))
	} else {
		p.print(p.keyword("INSERT"))
	}
	p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.TargetPath()))
	if cl != nil {
		p.println("")
		p.incDepth()
		p.accept(ctx, ast.Must(n.ColumnList()))
		p.println("")
		p.decDepth()
	}
	if q := ast.Must(n.Query()); q != nil {
		p.println("")
		p.acceptNested(ctx, q)
	} else {
		p.println("")
		p.println(p.keyword("VALUES"))
		p.incDepth()
		p.accept(ctx, ast.Must(n.Rows()))
		p.println("")
		p.decDepth()
	}

	p.movePast(n)
}

func (p *Printer) VisitInsertValuesRowList(ctx Context, n *googlesql.ASTInsertValuesRowList) {
	p.moveBefore(n)
	for i, r := range ast.ChildrenOfType[*googlesql.ASTInsertValuesRow](n) {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, r)
	}
	p.movePast(n)
}

func (p *Printer) VisitInsertValuesRow(ctx Context, n *googlesql.ASTInsertValuesRow) {
	p.moveBefore(n)
	values := ast.ChildrenOfType[googlesql.ASTExpressionNode](n)
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

func (p *Printer) VisitMergeStatement(ctx Context, n *googlesql.ASTMergeStatement) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.print(pp.keyword("MERGE") + " ")
	p1 := pp.nest()
	// p1.print(p.keyword("INTO"))
	p1.acceptNested(ctx, ast.Must(n.TargetPath()))
	pp.print(p1.unnest())
	pp.accept(ctx, ast.Must(n.Alias()))
	pp.println("")
	pp.print(pp.keyword("USING") + " ")
	pp.acceptNested(ctx, ast.Must(n.TableExpression()))
	pp.println("")
	pp.print(pp.keyword("ON") + " ")
	pp.acceptNested(ctx, ast.Must(n.MergeCondition()))
	p.println("")
	pp.accept(ctx, ast.Must(n.WhenClauses()))
	pp.movePast(n)
	p.print(pp.unnest())
}

func (p *Printer) VisitMergeAction(ctx Context, n *googlesql.ASTMergeAction) {
	p.moveBefore(n)
	switch ast.Must(n.ActionType()) {
	case googlesql.ASTMergeActionEnums_ActionTypeNotSet:
		// Nothing.
	case googlesql.ASTMergeActionEnums_ActionTypeInsert:
		p.visitMergeActionInsert(ctx, n)
	case googlesql.ASTMergeActionEnums_ActionTypeUpdate:
		p.visitMergeActionUpdate(ctx, n)
	case googlesql.ASTMergeActionEnums_ActionTypeDelete:
		p.visitMergeActionDelete(ctx, n)
	}
	p.movePast(n)
}

func (p *Printer) visitMergeActionDelete(ctx Context, n *googlesql.ASTMergeAction) {
	p.println(p.keyword("DELETE"))
}

func (p *Printer) visitMergeActionInsert(ctx Context, n *googlesql.ASTMergeAction) {
	cl := ast.Must(n.InsertColumnList())
	ir := ast.Must(n.InsertRow())
	if cl == nil && ir != nil && ast.Must(ir.NumChildren()) == 0 {
		p.println(p.keyword("INSERT ROW"))
		return
	}
	p.println(p.keyword("INSERT"))
	if cl != nil {
		p.incDepth()
		p.accept(ctx, ast.Must(n.InsertColumnList()))
		p.println("")
		p.decDepth()
	}
	if ir != nil {
		p.println("")
		p.println(p.keyword("VALUES"))
		p.incDepth()
		p.accept(ctx, ast.Must(n.InsertRow()))
		p.println("")
		p.decDepth()
	}
}

func (p *Printer) visitMergeActionUpdate(ctx Context, n *googlesql.ASTMergeAction) {
	p.println(p.keyword("UPDATE SET"))
	p.incDepth()
	p.accept(ctx, ast.Must(n.UpdateItemList()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitMergeWhenClauseList(ctx Context, n *googlesql.ASTMergeWhenClauseList) {
	p.moveBefore(n)
	for _, c := range ast.ChildrenOfType[*googlesql.ASTMergeWhenClause](n) {
		p.println("")
		p.print(p.keyword("WHEN") + " ")
		p.acceptNested(ctx, c)
	}
	p.println("")
	p.movePast(n)
}

func (p *Printer) VisitMergeWhenClause(ctx Context, n *googlesql.ASTMergeWhenClause) {
	p.moveBefore(n)
	switch ast.Must(n.MatchType()) {
	case googlesql.ASTMergeWhenClauseEnums_MatchTypeNotSet:
		// Nothing.
	case googlesql.ASTMergeWhenClauseEnums_MatchTypeMatched:
		p.print(p.keyword("MATCHED"))
	case googlesql.ASTMergeWhenClauseEnums_MatchTypeNotMatchedBySource:
		p.print(p.keyword("NOT MATCHED BY SOURCE"))
	case googlesql.ASTMergeWhenClauseEnums_MatchTypeNotMatchedByTarget:
		start := ast.GetParseLocationStartOffset(n)
		next := ast.GetParseLocationStartOffset(ast.Must(n.Action()))
		input := strings.ToUpper(p.viewErasedInput(start, next))
		if strings.Contains(input, "TARGET") {
			p.print(p.keyword("NOT MATCHED BY TARGET"))
		} else {
			p.print(p.keyword("NOT MATCHED"))
		}
	}
	if cond := ast.Must(n.SearchCondition()); cond != nil {
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
	p.accept(ctx, ast.Must(n.Action()))
	p.movePast(n)
}

func (p *Printer) VisitTruncateStatement(ctx Context, n *googlesql.ASTTruncateStatement) {
	p.moveBefore(n)
	p.print(p.keyword("TRUNCATE TABLE"))
	p.accept(ctx, ast.Must(n.TargetPath()))
	if w := ast.Must(n.Where()); w != nil {
		p.println("")
		p.print(p.keyword("WHERE"))
		p.accept(ctx, ast.Must(n.Where()))
	}
	p.movePast(n)
}

func (p *Printer) VisitUpdateItemList(ctx Context, n *googlesql.ASTUpdateItemList) {
	p.moveBefore(n)
	pp := p.nest()
	items := ast.ChildrenOfType[*googlesql.ASTUpdateItem](n)
	for i, item := range items {
		if i > 0 {
			pp.println(",")
		}
		pp.accept(ctx, item)
	}
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) VisitUpdateItem(ctx Context, n *googlesql.ASTUpdateItem) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.SetValue()))
	p.accept(ctx, ast.Must(n.InsertStatement()))
	p.accept(ctx, ast.Must(n.DeleteStatement()))
	p.accept(ctx, ast.Must(n.UpdateStatement()))
	p.movePast(n)
}

func (p *Printer) VisitUpdateSetValue(ctx Context, n *googlesql.ASTUpdateSetValue) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Path()))
	pp := p.nest()
	pp.print("=")
	pp.acceptNested(ctx, ast.Must(n.Value()))
	p.print(pp.unnest())
	p.movePast(n)
}
