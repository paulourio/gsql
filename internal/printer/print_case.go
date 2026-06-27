package printer

import (
	"github.com/paulourio/gsql/internal/sql"
)

// Case When are printed in two forms:
//
// Simple form:
//
//	CASE
//	  WHEN expr1 THEN ...
//	             ELSE ...
//	END
//
// General form:
//
//		CASE
//		  WHEN
//	     e1
//		  THEN
//	     ...
//
//	   WHEN
//	     e2
//		  THEN
//	     ...
//
//		  ELSE
//			...
//		END
func visitCaseArgs[T sql.ExpressionNode](p *Printer, ctx Context, args []T) {
	simple, _ := ctx.Bool(KeySimpleCase)
	if len(args) > 0 {
		p.moveBefore(args[0])
	}
	if simple {
		visitSimpleCaseArgs(p, ctx, args)
		return
	}
	visitGeneralCaseArgs(p, ctx, args)
}

func visitSimpleCaseArgs[T sql.ExpressionNode](p *Printer, ctx Context, args []T) {
	var lhs, rhs T
	pp := p.nest()
	for len(args) >= 2 {
		lhs, rhs, args = args[0], args[1], args[2:]
		pp.print(pp.keyword("WHEN"))
		pp.acceptNestedLeft(ctx, lhs)
		pp.print("\v" + pp.keyword("THEN"))
		pp.acceptNestedLeft(ctx, rhs)
		pp.movePastLine(rhs)
		pp.println("")
	}
	if len(args) == 1 {
		pp.print("\v\v" + pp.keyword("ELSE"))
		pp.acceptNestedLeft(ctx, args[0])
		pp.movePastLine(args[0])
		pp.println("")
	}
	p.print(pp.unnestLeft())
}

func visitGeneralCaseArgs[T sql.ExpressionNode](p *Printer, ctx Context, args []T) {
	var lhs, rhs T
	pp := p.nest()
	for len(args) >= 2 {
		lhs, rhs, args = args[0], args[1], args[2:]
		pp.println(pp.keyword("WHEN"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, lhs)
		pp.println("")
		pp.decDepth()
		pp.println(pp.keyword("THEN"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, rhs)
		pp.movePastLine(rhs)
		pp.println("")
		pp.decDepth()
		pp.println(" ")
	}
	if len(args) == 1 {
		pp.println(pp.keyword("ELSE"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, args[0])
		pp.movePastLine(args[0])
		pp.println("")
		pp.println(" ")
		pp.decDepth()
	}
	p.print(pp.unnestLeft())
}
