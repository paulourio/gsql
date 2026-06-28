package printer

import (
	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitAssertStatement(ctx Context, n *sql.AssertStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ASSERT") + " ")
	p.accept(ctx, n.Expr())
	if desc := n.Description(); desc != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, desc)
	}
	p.movePast(n)
}
