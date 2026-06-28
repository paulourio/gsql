package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

type DropStatementKeyworder interface {
	IsIfExists() bool
}

func dropKeyword(n DropStatementKeyworder, object string) string {
	var b strings.Builder
	b.Grow(20)
	b.WriteString("DROP ")
	b.WriteString(object)
	if n.IsIfExists() {
		b.WriteString(" IF EXISTS")
	}
	return b.String()
}

func (p *Printer) visitDropAllRowAccessPoliciesStatement(ctx Context, n *sql.DropAllRowAccessPoliciesStatement) {
	p.moveBefore(n)
	p.print(p.keyword("DROP ALL ROW ACCESS POLICIES ON "))
	p.print(p.identifier(p.toString(ctx, n.TableName())))
	p.movePast(n)
}

func (p *Printer) visitDropEntityStatement(ctx Context, n *sql.DropEntityStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "ENTITY")))
	p.print(p.identifier(p.toString(ctx, n.GetDdlTarget())))
	p.movePast(n)
}

func (p *Printer) visitDropFunctionStatement(ctx Context, n *sql.DropFunctionStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "FUNCTION")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropMaterializedViewStatement(ctx Context, n *sql.DropMaterializedViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "MATERIALIZED VIEW")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropPrivilegeRestrictionStatement(_ Context, _ *sql.DropPrivilegeRestrictionStatement) {
}

func (p *Printer) visitDropRowAccessPolicyStatement(ctx Context, n *sql.DropRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "ROW ACCESS POLICY")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.print(p.keyword("ON"))
	p.accept(ctx, n.TableName())
	p.movePast(n)
}

func (p *Printer) visitDropSearchIndexStatement(ctx Context, n *sql.DropSearchIndexStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "SEARCH INDEX")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropSnapshotTableStatement(ctx Context, n *sql.DropSnapshotTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "SNAPSHOT TABLE")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropStatement(ctx Context, n *sql.DropStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropStatementKeyword(n)))
	p.accept(ctx.WithValue(KeyInTableName, true), n.GetDdlTarget())
	switch n.DropMode() {
	case sql.RestrictDropMode:
		p.print(p.keyword("RESTRICT"))
	case sql.CascadeDropMode:
		p.print(p.keyword("CASCADE"))
	}
	p.movePast(n)
}

func (p *Printer) visitDropTableFunctionStatement(ctx Context, n *sql.DropTableFunctionStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "TABLE FUNCTION")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}
