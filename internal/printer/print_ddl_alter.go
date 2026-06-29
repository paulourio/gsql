package printer

import (
	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitAddColumnAction(ctx Context, n *sql.AddColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword("ADD COLUMN"))

	if n.IsIfNotExists() {
		p.print(p.keyword("IF NOT EXISTS"))
	}

	p.accept(ctx, n.ColumnDefinition())
}

func (p *Printer) visitAddConstraintAction(ctx Context, n *sql.AddConstraintAction) {
	p.moveBefore(n)
	p.print(p.keyword("ADD"))
	if n.IsIfNotExists() {
		p.print(p.keyword("IF NOT EXISTS"))
	}
	p.accept(ctx, n.Constraint())
}

func (p *Printer) visitAlterActionList(ctx Context, n *sql.AlterActionList) {
	p.moveBefore(n)
	p.println("")
	for i, a := range n.Actions() {
		if i > 0 {
			p.println(",")
		}
		p.acceptNested(ctx, a)
	}
}

func (p *Printer) visitAlterAllRowAccessPoliciesStatement(_ Context, _ *sql.AlterAllRowAccessPoliciesStatement) {
}

func (p *Printer) visitAlterColumnDropDefaultAction(ctx Context, n *sql.AlterColumnDropDefaultAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.println(p.keyword("DROP DEFAULT"))
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitAlterColumnDropNotNullAction(ctx Context, n *sql.AlterColumnDropNotNullAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.println(p.keyword("DROP NOT NULL"))
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitAlterColumnOptionsAction(ctx Context, n *sql.AlterColumnOptionsAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET"))
	p.accept(ctx, n.OptionsList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterColumnSetDefaultAction(ctx Context, n *sql.AlterColumnSetDefaultAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET DEFAULT"))
	p.accept(ctx, n.DefaultExpression())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterColumnTypeAction(ctx Context, n *sql.AlterColumnTypeAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET DATA TYPE"))
	p.accept(ctx, n.Schema())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterConstraintEnforcementAction(_ Context, _ *sql.AlterConstraintEnforcementAction) {
}

func (p *Printer) visitAlterConstraintSetOptionsAction(_ Context, _ *sql.AlterConstraintSetOptionsAction) {
}

func (p *Printer) visitAlterDatabaseStatement(ctx Context, n *sql.AlterDatabaseStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER DATABASE"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterEntityStatement(ctx Context, n *sql.AlterEntityStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER ENTITY"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterMaterializedViewStatement(ctx Context, n *sql.AlterMaterializedViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER MATERIALIZED VIEW"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterPrivilegeRestrictionStatement(_ Context, _ *sql.AlterPrivilegeRestrictionStatement) {
}

func (p *Printer) visitAlterRowAccessPolicyStatement(ctx Context, n *sql.AlterRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER ROW ACCESS POLICY"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterSchemaStatement(ctx Context, n *sql.AlterSchemaStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER SCHEMA"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterTableStatement(ctx Context, n *sql.AlterTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER TABLE"))
	p.accept(ctx.WithValue(KeyInTableName, true), n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitAlterViewStatement(ctx Context, n *sql.AlterViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER VIEW"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitDropColumnAction(ctx Context, n *sql.DropColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "COLUMN")))
	p.print(p.identifier(p.toString(ctx, n.ColumnName())))
	p.movePast(n)
}

func (p *Printer) visitDropConstraintAction(ctx Context, n *sql.DropConstraintAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "CONSTRAINT")))
	p.print(p.identifier(p.toString(ctx, n.ConstraintName())))
	p.movePast(n)
}

func (p *Printer) visitDropPrimaryKeyAction(_ Context, n *sql.DropPrimaryKeyAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "PRIMARY KEY")))
	p.movePast(n)
}

func (p *Printer) visitRenameColumnAction(ctx Context, n *sql.RenameColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword("RENAME COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.print(p.keyword("TO"))
	p.accept(ctx, n.NewColumnName())
}

func (p *Printer) visitRenameStatement(ctx Context, n *sql.RenameStatement) {
	p.moveBefore(n)
	p.print(p.keyword("RENAME"))
	// The identifier is actually an object like TABLE or VIEW, so we print
	// like keyword.
	p.print(p.keyword(n.Identifier().GetAsString()))
	p.accept(ctx, n.OldName())
	p.print(p.keyword("TO"))
	p.accept(ctx, n.NewName())
}

func (p *Printer) visitRenameToClause(ctx Context, n *sql.RenameToClause) {
	p.moveBefore(n)
	p.print(p.keyword("RENAME TO"))
	p.accept(ctx, n.NewName())
}

func (p *Printer) visitSetCollateClause(ctx Context, n *sql.SetCollateClause) {
	p.moveBefore(n)
	p.print(p.keyword("SET DEFAULT"))
	p.accept(ctx, n.Collate())
}

func (p *Printer) visitSetOptionsAction(ctx Context, n *sql.SetOptionsAction) {
	p.moveBefore(n)
	p.print(p.keyword("SET"))
	p.accept(ctx, n.OptionsList())
}
