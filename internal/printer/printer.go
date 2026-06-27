package printer

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/tabwriter"

	"github.com/hashicorp/go-multierror"

	"github.com/paulourio/gsql/format"
	"github.com/paulourio/gsql/internal/sql"
)

type Printer struct {
	OriginalInput string
	ErasedInput   string
	Writer        *Writer
	Tracker       *LocationTracker

	err error
}

func (p *Printer) Print(root sql.Node) (string, error) {
	ctx := &emptyCtx{}
	p.visit(ctx, root, false)
	p.Writer.FlushLine()
	// Flush any remaining extensions.
	if len(p.Writer.comments.comments) > 0 {
		p.println("")
		p.Writer.flushCommentsUpTo(len(p.OriginalInput))
	}
	result := p.unnest()
	if p.Writer.opts.AlignTrailingComments {
		result = alignTrailingComments(result)
	}
	result = strings.ReplaceAll(result, "\v", "") + "\n"
	result = strings.ReplaceAll(result, lineBreakPlaceholder, "\n")
	result = rowsTrimRight(result)
	return result, p.err
}

// accept visits a sql.Node on the current line.
// Nil nodes are silently ignored.
func (p *Printer) accept(ctx Context, n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	p.visit(ctx, n, false)
}

// lnaccept visits a sql.Node on a new line.
// If the node is nil, no new line is emitted.
func (p *Printer) lnaccept(ctx Context, n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	p.visit(ctx, n, true)
}

func (p *Printer) visit(ctx Context, n sql.Node, newline bool) {
	if !sql.Defined(n) {
		return
	}
	if newline {
		p.println("")
	}
	switch n.Kind() {
	case sql.AddColumnActionKind:
		p.visitAddColumnAction(ctx, n.(*sql.AddColumnAction))
	case sql.AddConstraintActionKind:
		p.visitAddConstraintAction(ctx, n.(*sql.AddConstraintAction))
	case sql.AliasKind:
		p.visitAlias(ctx, n.(*sql.Alias))
	case sql.AliasedGroupRowsKind:
		p.visitAliasedGroupRows(ctx, n.(*sql.AliasedGroupRows))
	case sql.AliasedQueryKind:
		p.visitAliasedQuery(ctx, n.(*sql.AliasedQuery))
	case sql.AlterActionListKind:
		p.visitAlterActionList(ctx, n.(*sql.AlterActionList))
	case sql.AlterAllRowAccessPoliciesStatementKind:
		p.visitAlterAllRowAccessPoliciesStatement(ctx, n.(*sql.AlterAllRowAccessPoliciesStatement))
	case sql.AlterColumnDropDefaultActionKind:
		p.visitAlterColumnDropDefaultAction(ctx, n.(*sql.AlterColumnDropDefaultAction))
	case sql.AlterColumnDropNotNullActionKind:
		p.visitAlterColumnDropNotNullAction(ctx, n.(*sql.AlterColumnDropNotNullAction))
	case sql.AlterColumnOptionsActionKind:
		p.visitAlterColumnOptionsAction(ctx, n.(*sql.AlterColumnOptionsAction))
	case sql.AlterColumnSetDefaultActionKind:
		p.visitAlterColumnSetDefaultAction(ctx, n.(*sql.AlterColumnSetDefaultAction))
	case sql.AlterColumnTypeActionKind:
		p.visitAlterColumnTypeAction(ctx, n.(*sql.AlterColumnTypeAction))
	case sql.AlterConstraintEnforcementActionKind:
		p.visitAlterConstraintEnforcementAction(ctx, n.(*sql.AlterConstraintEnforcementAction))
	case sql.AlterConstraintSetOptionsActionKind:
		p.visitAlterConstraintSetOptionsAction(ctx, n.(*sql.AlterConstraintSetOptionsAction))
	case sql.AlterDatabaseStatementKind:
		p.visitAlterDatabaseStatement(ctx, n.(*sql.AlterDatabaseStatement))
	case sql.AlterEntityStatementKind:
		p.visitAlterEntityStatement(ctx, n.(*sql.AlterEntityStatement))
	case sql.AlterMaterializedViewStatementKind:
		p.visitAlterMaterializedViewStatement(ctx, n.(*sql.AlterMaterializedViewStatement))
	case sql.AlterPrivilegeRestrictionStatementKind:
		p.visitAlterPrivilegeRestrictionStatement(ctx, n.(*sql.AlterPrivilegeRestrictionStatement))
	case sql.AlterRowAccessPolicyStatementKind:
		p.visitAlterRowAccessPolicyStatement(ctx, n.(*sql.AlterRowAccessPolicyStatement))
	case sql.AlterSchemaStatementKind:
		p.visitAlterSchemaStatement(ctx, n.(*sql.AlterSchemaStatement))
	case sql.AlterTableStatementKind:
		p.visitAlterTableStatement(ctx, n.(*sql.AlterTableStatement))
	case sql.AlterViewStatementKind:
		p.visitAlterViewStatement(ctx, n.(*sql.AlterViewStatement))
	case sql.AnalyticFunctionCallKind:
		p.visitAnalyticFunctionCall(ctx, n.(*sql.AnalyticFunctionCall))
	case sql.AndExprKind:
		p.visitAndExpr(ctx, n.(*sql.AndExpr))
	case sql.ArrayConstructorKind:
		p.visitArrayConstructor(ctx, n.(*sql.ArrayConstructor))
	case sql.ArrayColumnSchemaKind:
		if ac, ok := n.(*sql.ArrayColumnSchema); ok {
			p.visitArrayColumnSchema(ctx, ac)
		} else {
			p.visitColumnSchema(ctx, n.(*sql.ColumnSchema))
		}
	case sql.ArrayElementKind:
		if ae, ok := n.(*sql.ArrayElement); ok {
			p.visitArrayElement(ctx, ae)
		} else {
			p.visitGeneralizedPathExpression(ctx, n)
		}
	case sql.ArrayTypeKind:
		p.visitArrayType(ctx, n.(*sql.ArrayType))
	case sql.AssignmentFromStructKind:
		p.visitAssignmentFromStruct(ctx, n.(*sql.AssignmentFromStruct))
	case sql.BeginEndBlockKind:
		p.visitBeginEndBlock(ctx, n.(*sql.BeginEndBlock))
	case sql.BeginStatementKind:
		p.visitBeginStatementNode(ctx, n.(*sql.BeginStatement))
	case sql.BetweenExpressionKind:
		p.visitBetweenExpression(ctx, n.(*sql.BetweenExpression))
	case sql.BignumericLiteralKind:
		p.visitBigNumericLiteral(ctx, n.(*sql.BigNumericLiteral))
	case sql.BinaryExpressionKind:
		p.visitBinaryExpression(ctx, n.(*sql.BinaryExpression))
	case sql.BitwiseShiftExpressionKind:
		p.visitBitwiseShiftExpression(ctx, n.(*sql.BitwiseShiftExpression))
	case sql.BooleanLiteralKind:
		p.visitBoolLiteral(ctx, n.(*sql.BooleanLiteral))
	case sql.BytesLiteralKind:
		p.visitBytesLiteral(ctx, n.(*sql.BytesLiteral))
	case sql.CallStatementKind:
		p.visitCallStatement(ctx, n.(*sql.CallStatement))
	case sql.CaseNoValueExpressionKind:
		p.visitCaseNoValueExpression(ctx, n.(*sql.CaseNoValueExpression))
	case sql.CaseValueExpressionKind:
		p.visitCaseValueExpression(ctx, n.(*sql.CaseValueExpression))
	case sql.CastExpressionKind:
		p.visitCastExpression(ctx, n.(*sql.CastExpression))
	case sql.ClampedBetweenModifierKind:
		p.visitClampedBetweenModifier(ctx, n.(*sql.ClampedBetweenModifier))
	case sql.CloneDataSourceKind:
		p.visitCloneDataSource(ctx, n.(*sql.CloneDataSource))
	case sql.ClusterByKind:
		p.visitClusterBy(ctx, n.(*sql.ClusterBy))
	case sql.CollateKind:
		p.visitCollate(ctx, n.(*sql.Collate))
	case sql.ColumnAttributeListKind:
		p.visitColumnAttributeList(ctx, n.(*sql.ColumnAttributeList))
	case sql.ColumnDefinitionKind:
		p.visitColumnDefinition(ctx, n.(*sql.ColumnDefinition))
	case sql.ColumnListKind:
		p.visitColumnList(ctx, n.(*sql.ColumnList))
	case sql.ColumnWithOptionsKind:
		p.visitColumnWithOptions(ctx, n.(*sql.ColumnWithOptions))
	case sql.ColumnWithOptionsListKind:
		p.visitColumnWithOptionsList(ctx, n.(*sql.ColumnWithOptionsList))
	case sql.CommitStatementKind:
		p.visitCommitStatement(ctx, n.(*sql.CommitStatement))
	case sql.ConnectionClauseKind:
		p.visitConnectionClause(ctx, n.(*sql.ConnectionClause))
	case sql.CopyDataSourceKind:
		p.visitCopyDataSource(ctx, n.(*sql.CopyDataSource))
	case sql.CreateExternalTableStatementKind:
		p.visitCreateExternalTableStatement(ctx, n.(*sql.CreateExternalTableStatement))
	case sql.CreateFunctionStatementKind:
		p.visitCreateFunctionStatement(ctx, n.(*sql.CreateFunctionStatement))
	case sql.CreateMaterializedViewStatementKind:
		p.visitCreateMaterializedViewStatement(ctx, n.(*sql.CreateMaterializedViewStatement))
	case sql.CreateProcedureStatementKind:
		p.visitCreateProcedureStatement(ctx, n.(*sql.CreateProcedureStatement))
	case sql.CreateRowAccessPolicyStatementKind:
		p.visitCreateRowAccessPolicyStatement(ctx, n.(*sql.CreateRowAccessPolicyStatement))
	case sql.CreateSchemaStatementKind:
		p.visitCreateSchemaStatement(ctx, n.(*sql.CreateSchemaStatement))
	case sql.CreateSnapshotTableStatementKind:
		p.visitCreateSnapshotTableStatement(ctx, n.(*sql.CreateSnapshotTableStatement))
	case sql.CreateTableStatementKind:
		p.visitCreateTableStatement(ctx, n.(*sql.CreateTableStatement))
	case sql.CreateTableFunctionStatementKind:
		p.visitCreateTableFunctionStatement(ctx, n.(*sql.CreateTableFunctionStatement))
	case sql.CreateViewStatementKind:
		p.visitCreateViewStatement(ctx, n.(*sql.CreateViewStatement))
	case sql.DateOrTimeLiteralKind:
		p.visitDateOrTimeLiteral(ctx, n.(*sql.DateOrTimeLiteral))
	case sql.DescriptorKind:
		p.visitDescriptor(ctx, n.(*sql.Descriptor))
	case sql.DescriptorColumnKind:
		p.visitDescriptorColumn(ctx, n.(*sql.DescriptorColumn))
	case sql.DescriptorColumnListKind:
		p.visitDescriptorColumnList(ctx, n.(*sql.DescriptorColumnList))
	case sql.DotIdentifierKind:
		p.visitDotIdentifier(ctx, n.(*sql.DotIdentifier))
	case sql.DotGeneralizedFieldKind:
		if dg, ok := n.(*sql.DotGeneralizedField); ok {
			p.visitDotGeneralizedField(ctx, dg)
		} else {
			p.visitGeneralizedPathExpression(ctx, n)
		}
	case sql.DotStarKind:
		p.visitDotStar(ctx, n.(*sql.DotStar))
	case sql.DotStarWithModifiersKind:
		p.visitDotStarWithModifiers(ctx, n.(*sql.DotStarWithModifiers))
	case sql.DropAllRowAccessPoliciesStatementKind:
		p.visitDropAllRowAccessPoliciesStatement(ctx, n.(*sql.DropAllRowAccessPoliciesStatement))
	case sql.DropColumnActionKind:
		p.visitDropColumnAction(ctx, n.(*sql.DropColumnAction))
	case sql.DropConstraintActionKind:
		p.visitDropConstraintAction(ctx, n.(*sql.DropConstraintAction))
	case sql.DropEntityStatementKind:
		p.visitDropEntityStatement(ctx, n.(*sql.DropEntityStatement))
	case sql.DropFunctionStatementKind:
		p.visitDropFunctionStatement(ctx, n.(*sql.DropFunctionStatement))
	case sql.DropMaterializedViewStatementKind:
		p.visitDropMaterializedViewStatement(ctx, n.(*sql.DropMaterializedViewStatement))
	case sql.DropPrimaryKeyActionKind:
		p.visitDropPrimaryKeyAction(ctx, n.(*sql.DropPrimaryKeyAction))
	case sql.DropPrivilegeRestrictionStatementKind:
		p.visitDropPrivilegeRestrictionStatement(ctx, n.(*sql.DropPrivilegeRestrictionStatement))
	case sql.DropRowAccessPolicyStatementKind:
		p.visitDropRowAccessPolicyStatement(ctx, n.(*sql.DropRowAccessPolicyStatement))
	case sql.DropSearchIndexStatementKind:
		p.visitDropSearchIndexStatement(ctx, n.(*sql.DropSearchIndexStatement))
	case sql.DropSnapshotTableStatementKind:
		p.visitDropSnapshotTableStatement(ctx, n.(*sql.DropSnapshotTableStatement))
	case sql.DropTableFunctionStatementKind:
		p.visitDropTableFunctionStatement(ctx, n.(*sql.DropTableFunctionStatement))
	case sql.DropStatementKind:
		p.visitDropStatement(ctx, n.(*sql.DropStatement))
	case sql.ExceptionHandlerListKind:
		p.visitExceptionHandlerListNode(ctx, n.(*sql.ExceptionHandlerList))
	case sql.ExceptionHandlerKind:
		p.visitExceptionHandlerNode(ctx, n.(*sql.ExceptionHandler))
	case sql.ExecuteIntoClauseKind:
		p.visitExecuteIntoClause(ctx, n.(*sql.ExecuteIntoClause))
	case sql.ExecuteImmediateStatementKind:
		p.visitExecuteImmediateStatement(ctx, n.(*sql.ExecuteImmediateStatement))
	case sql.ExecuteUsingArgumentKind:
		p.visitExecuteUsingArgument(ctx, n.(*sql.ExecuteUsingArgument))
	case sql.ExecuteUsingClauseKind:
		p.visitExecuteUsingClause(ctx, n.(*sql.ExecuteUsingClause))
	case sql.ExpressionSubqueryKind:
		p.visitExpressionSubquery(ctx, n.(*sql.ExpressionSubquery))
	case sql.ExpressionWithOptAliasKind:
		p.visitExpressionWithOptAlias(ctx, n.(*sql.ExpressionWithOptAlias))
	case sql.ExtractExpressionKind:
		p.visitExtractExpression(ctx, n.(*sql.ExtractExpression))
	case sql.FloatLiteralKind:
		p.visitFloatLiteral(ctx, n.(*sql.FloatLiteral))
	case sql.FilterUsingClauseKind:
		p.visitFilterUsingClause(ctx, n.(*sql.FilterUsingClause))
	case sql.ForeignKeyKind:
		if fk, ok := n.(*sql.ForeignKey); ok {
			p.visitForeignKey(ctx, fk)
		} else {
			p.visitTableConstraint(ctx, n.(*sql.TableConstraint))
		}
	case sql.ForeignKeyReferenceKind:
		p.visitForeignKeyReference(ctx, n.(*sql.ForeignKeyReference))
	case sql.FormatClauseKind:
		p.visitFormatClause(ctx, n.(*sql.FormatClause))
	case sql.ForSystemTimeKind:
		p.visitForSystemTime(ctx, n.(*sql.ForSystemTime))
	case sql.FromClauseKind:
		p.visitFromClause(ctx, n.(*sql.FromClause))
	case sql.FunctionCallKind:
		p.visitFunctionCall(ctx, n.(*sql.FunctionCall))
	case sql.FunctionDeclarationKind:
		p.visitFunctionDeclaration(ctx, n.(*sql.FunctionDeclaration))
	case sql.FunctionParameterKind:
		p.visitFunctionParameter(ctx, n.(*sql.FunctionParameter))
	case sql.FunctionParametersKind:
		p.visitFunctionParameters(ctx, n.(*sql.FunctionParameters))

	case sql.GranteeListKind:
		p.visitGranteeList(ctx, n.(*sql.GranteeList))
	case sql.GrantToClauseKind:
		p.visitGrantToClause(ctx, n.(*sql.GrantToClause))
	case sql.GroupByKind:
		p.visitGroupBy(ctx, n.(*sql.GroupBy))
	case sql.GroupByAllKind:
		p.visitGroupByAll(ctx, n.(*sql.GroupByAll))
	case sql.GroupingItemKind:
		p.visitGroupingItem(ctx, n.(*sql.GroupingItem))
	case sql.HavingModifierKind:
		p.visitHavingModifier(ctx, n.(*sql.HavingModifier))
	case sql.HavingKind:
		p.visitHaving(ctx, n.(*sql.Having))
	case sql.HintKind:
		p.visitHint(ctx, n.(*sql.Hint))
	case sql.HintedStatementKind:
		p.visitHintedStatement(ctx, n.(*sql.HintedStatement))
	case sql.IdentifierKind:
		p.visitIdentifier(ctx, n.(*sql.Identifier))
	case sql.IdentifierListKind:
		p.visitIdentifierList(ctx, n.(*sql.IdentifierList))
	case sql.IfStatementKind:
		p.visitIfStatement(ctx, n.(*sql.IfStatement))
	case sql.InExpressionKind:
		p.visitInExpression(ctx, n.(*sql.InExpression))
	case sql.InListKind:
		p.visitInList(ctx, n.(*sql.InList))
	case sql.IntervalExprKind:
		p.visitIntervalExpr(ctx, n.(*sql.IntervalExpr))
	case sql.IntLiteralKind:
		p.visitIntLiteral(ctx, n.(*sql.IntLiteral))
	case sql.InsertStatementKind:
		p.visitInsertStatement(ctx, n.(*sql.InsertStatement))
	case sql.InsertValuesRowListKind:
		p.visitInsertValuesRowList(ctx, n.(*sql.InsertValuesRowList))
	case sql.InsertValuesRowKind:
		p.visitInsertValuesRow(ctx, n.(*sql.InsertValuesRow))
	case sql.JoinKind:
		p.visitJoin(ctx, n.(*sql.Join))
	case sql.JsonLiteralKind:
		p.visitJSONLiteral(ctx, n.(*sql.JSONLiteral))
	case sql.LimitKind:
		p.visitLimit(ctx, n.(*sql.Limit))
	case sql.LimitOffsetKind:
		p.visitLimitOffset(ctx, n.(*sql.LimitOffset))
	case sql.MergeActionKind:
		p.visitMergeAction(ctx, n.(*sql.MergeAction))
	case sql.MergeStatementKind:
		p.visitMergeStatement(ctx, n.(*sql.MergeStatement))
	case sql.MergeWhenClauseKind:
		p.visitMergeWhenClause(ctx, n.(*sql.MergeWhenClause))
	case sql.MergeWhenClauseListKind:
		p.visitMergeWhenClauseList(ctx, n.(*sql.MergeWhenClauseList))
	case sql.ModelClauseKind:
		p.visitModelClause(ctx, n.(*sql.ModelClause))
	case sql.NamedArgumentKind:
		p.visitNamedArgument(ctx, n.(*sql.NamedArgument))
	case sql.NotNullColumnAttributeKind:
		p.visitNotNullColumnAttribute(ctx, n.(*sql.NotNullColumnAttribute))
	case sql.NullLiteralKind:
		p.visitNullLiteral(ctx, n.(*sql.NullLiteral))
	case sql.NullOrderKind:
		p.visitNullOrder(ctx, n.(*sql.NullOrder))
	case sql.NumericLiteralKind:
		p.visitNumericLiteral(ctx, n.(*sql.NumericLiteral))
	case sql.OnClauseKind:
		p.visitOnClause(ctx, n.(*sql.OnClause))
	case sql.OptionsListKind:
		p.visitOptionsList(ctx, n.(*sql.OptionsList))
	case sql.OptionsEntryKind:
		p.visitOptionsEntry(ctx, n.(*sql.OptionsEntry))
	case sql.OrExprKind:
		p.visitOrExpr(ctx, n.(*sql.OrExpr))
	case sql.OrderByKind:
		p.visitOrderBy(ctx, n.(*sql.OrderBy))
	case sql.OrderingExpressionKind:
		p.visitOrderingExpression(ctx, n.(*sql.OrderingExpression))
	case sql.ParameterAssignmentKind:
		p.visitParameterAssignment(ctx, n.(*sql.ParameterAssignment))
	case sql.ParameterExprKind:
		p.visitParameterExpr(ctx, n.(*sql.ParameterExpr))
	case sql.ParenthesizedJoinKind:
		p.visitParenthesizedJoin(ctx, n.(*sql.ParenthesizedJoin))
	case sql.PartitionByKind:
		p.visitPartitionBy(ctx, n.(*sql.PartitionBy))
	case sql.PathExpressionListKind:
		p.visitPathExpressionList(ctx, n.(*sql.PathExpressionList))
	case sql.PathExpressionKind:
		if pe, ok := n.(*sql.PathExpression); ok {
			p.visitPathExpression(ctx, pe)
		} else {
			p.visitGeneralizedPathExpression(ctx, n)
		}
	case sql.PivotClauseKind:
		p.visitPivotClause(ctx, n.(*sql.PivotClause))
	case sql.PivotExpressionKind:
		p.visitPivotExpression(ctx, n.(*sql.PivotExpression))
	case sql.PivotExpressionListKind:
		p.visitPivotExpressionList(ctx, n.(*sql.PivotExpressionList))
	case sql.PivotValueKind:
		p.visitPivotValue(ctx, n.(*sql.PivotValue))
	case sql.PivotValueListKind:
		p.visitPivotValueList(ctx, n.(*sql.PivotValueList))
	case sql.PrimaryKeyKind:
		if pk, ok := n.(*sql.PrimaryKey); ok {
			p.visitPrimaryKey(ctx, pk)
		} else {
			p.visitTableConstraint(ctx, n.(*sql.TableConstraint))
		}
	case sql.PrimaryKeyColumnAttributeKind:
		p.visitPrimaryKeyColumnAttribute(ctx, n.(*sql.PrimaryKeyColumnAttribute))
	case sql.PrimaryKeyElementListKind:
		p.visitPrimaryKeyElementList(ctx, n.(*sql.PrimaryKeyElementList))
	case sql.PrimaryKeyElementKind:
		p.visitPrimaryKeyElement(ctx, n.(*sql.PrimaryKeyElement))
	case sql.QualifyKind:
		p.visitQualify(ctx, n.(*sql.Qualify))
	case sql.QueryKind:
		p.visitQuery(ctx, n.(*sql.Query))
	case sql.QueryStatementKind:
		p.visitQueryStatement(ctx, n.(*sql.QueryStatement))
	case sql.RenameColumnActionKind:
		p.visitRenameColumnAction(ctx, n.(*sql.RenameColumnAction))
	case sql.RenameToClauseKind:
		p.visitRenameToClause(ctx, n.(*sql.RenameToClause))
	case sql.RepeatableClauseKind:
		p.visitRepeatableClause(ctx, n.(*sql.RepeatableClause))
	case sql.ReturnStatementKind:
		p.visitReturnStatement(ctx, n.(*sql.ReturnStatement))
	case sql.RollbackStatementKind:
		p.visitRollbackStatementNode(ctx, n.(*sql.RollbackStatement))
	case sql.RollupKind:
		p.visitRollup(ctx, n.(*sql.Rollup))
	case sql.SampleClauseKind:
		p.visitSampleClause(ctx, n.(*sql.SampleClause))
	case sql.SampleSizeKind:
		p.visitSampleSize(ctx, n.(*sql.SampleSize))
	case sql.SampleSuffixKind:
		p.visitSampleSuffix(ctx, n.(*sql.SampleSuffix))
	case sql.SetCollateClauseKind:
		p.visitSetCollateClause(ctx, n.(*sql.SetCollateClause))
	case sql.ScriptKind:
		p.visitScript(ctx, n.(*sql.Script))
	case sql.SelectKind:
		p.visitSelect(ctx, n.(*sql.Select))
	case sql.SelectAsKind:
		p.visitSelectAs(ctx, n.(*sql.SelectAs))
	case sql.SelectColumnKind:
		p.visitSelectColumn(ctx, n.(*sql.SelectColumn))
	case sql.SelectListKind:
		p.visitSelectList(ctx, n.(*sql.SelectList))
	case sql.SetOptionsActionKind:
		p.visitSetOptionsAction(ctx, n.(*sql.SetOptionsAction))
	case sql.SetOperationKind:
		p.visitSetOperation(ctx, n.(*sql.SetOperation))
	case sql.SimpleColumnSchemaKind:
		if sc, ok := n.(*sql.SimpleColumnSchema); ok {
			p.visitSimpleColumnSchema(ctx, sc)
		} else {
			p.visitColumnSchema(ctx, n.(*sql.ColumnSchema))
		}
	case sql.SimpleTypeKind:
		p.visitSimpleType(ctx, n.(*sql.SimpleType))
	case sql.SQLFunctionBodyKind:
		p.visitSQLFunctionBody(ctx, n.(*sql.SQLFunctionBody))
	case sql.StarKind:
		p.visitStar(ctx, n.(*sql.Star))
	case sql.StarModifiersKind:
		p.visitStarModifiers(ctx, n.(*sql.StarModifiers))
	case sql.StarReplaceItemKind:
		p.visitStarReplaceItem(ctx, n.(*sql.StarReplaceItem))
	case sql.StarWithModifiersKind:
		p.visitStarWithModifiers(ctx, n.(*sql.StarWithModifiers))
	case sql.StatementListKind:
		p.visitStatementList(ctx, n.(*sql.StatementList))
	case sql.StringLiteralKind:
		p.visitStringLiteral(ctx, n.(*sql.StringLiteral))
	case sql.StructColumnFieldKind:
		p.visitStructColumnField(ctx, n.(*sql.StructColumnField))
	case sql.StructColumnSchemaKind:
		if sc, ok := n.(*sql.StructColumnSchema); ok {
			p.visitStructColumnSchema(ctx, sc)
		} else {
			p.visitColumnSchema(ctx, n.(*sql.ColumnSchema))
		}
	case sql.StructConstructorArgKind:
		p.visitStructConstructorArg(ctx, n.(*sql.StructConstructorArg))
	case sql.StructConstructorWithKeywordKind:
		p.visitStructConstructorWithKeyword(ctx, n.(*sql.StructConstructorWithKeyword))
	case sql.StructConstructorWithParensKind:
		p.visitStructConstructorWithParens(ctx, n.(*sql.StructConstructorWithParens))
	case sql.StructFieldKind:
		p.visitStructField(ctx, n.(*sql.StructField))
	case sql.StructTypeKind:
		p.visitStructType(ctx, n.(*sql.StructType))
	case sql.SystemVariableAssignmentKind:
		p.visitSystemVariableAssignment(ctx, n.(*sql.SystemVariableAssignment))
	case sql.SystemVariableExprKind:
		p.visitSystemVariableExpr(ctx, n.(*sql.SystemVariableExpr))
	case sql.TableClauseKind:
		p.visitTableClause(ctx, n.(*sql.TableClause))

	case sql.TableElementListKind:
		p.visitTableElementList(ctx, n.(*sql.TableElementList))
	case sql.TablePathExpressionKind:
		p.visitTablePathExpression(ctx, n.(*sql.TablePathExpression))
	case sql.TableSubqueryKind:
		p.visitTableSubquery(ctx, n.(*sql.TableSubquery))
	case sql.TemplatedParameterTypeKind:
		p.visitTemplatedParameterType(ctx, n.(*sql.TemplatedParameterType))
	case sql.TruncateStatementKind:
		p.visitTruncateStatement(ctx, n.(*sql.TruncateStatement))
	case sql.TVFArgumentKind:
		p.visitTVFArgument(ctx, n.(*sql.TVFArgument))
	case sql.TVFKind:
		p.visitTVF(ctx, n.(*sql.TVF))
	case sql.TVFSchemaKind:
		p.visitTVFSchema(ctx, n.(*sql.TVFSchema))
	case sql.TVFSchemaColumnKind:
		p.visitTVFSchemaColumn(ctx, n.(*sql.TVFSchemaColumn))
	case sql.TypeParameterListKind:
		p.visitTypeParameterList(ctx, n.(*sql.TypeParameterList))
	case sql.UnpivotClauseKind:
		p.visitUnpivotClause(ctx, n.(*sql.UnpivotClause))
	case sql.UnaryExpressionKind:
		p.visitUnaryExpression(ctx, n.(*sql.UnaryExpression))
	case sql.UnpivotInItemLabelKind:
		p.visitUnpivotInItemLabel(ctx, n.(*sql.UnpivotInItemLabel))
	case sql.UnpivotInItemListKind:
		p.visitUnpivotInItemList(ctx, n.(*sql.UnpivotInItemList))
	case sql.UnpivotInItemKind:
		p.visitUnpivotInItem(ctx, n.(*sql.UnpivotInItem))
	case sql.UnnestExpressionKind:
		p.visitUnnestExpression(ctx, n.(*sql.UnnestExpression))
	case sql.UnnestExpressionWithOptAliasAndOffsetKind:
		p.visitUnnestExpressionWithOptAliasAndOffset(ctx, n.(*sql.UnnestExpressionWithOptAliasAndOffset))
	case sql.UpdateItemKind:
		p.visitUpdateItem(ctx, n.(*sql.UpdateItem))
	case sql.UpdateItemListKind:
		p.visitUpdateItemList(ctx, n.(*sql.UpdateItemList))
	case sql.UpdateSetValueKind:
		p.visitUpdateSetValue(ctx, n.(*sql.UpdateSetValue))
	case sql.UsingClauseKind:
		p.visitUsingClause(ctx, n.(*sql.UsingClause))
	case sql.VariableDeclarationKind:
		p.visitVariableDeclaration(ctx, n.(*sql.VariableDeclaration))
	case sql.SingleAssignmentKind:
		p.visitSingleAssignment(ctx, n.(*sql.SingleAssignment))
	case sql.WhereClauseKind:
		p.visitWhereClause(ctx, n.(*sql.WhereClause))
	case sql.WindowClauseKind:
		p.visitWindowClause(ctx, n.(*sql.WindowClause))
	case sql.WindowFrameKind:
		p.visitWindowFrame(ctx, n.(*sql.WindowFrame))
	case sql.WindowFrameExprKind:
		p.visitWindowFrameExpr(ctx, n.(*sql.WindowFrameExpr))
	case sql.WindowSpecificationKind:
		p.visitWindowSpecification(ctx, n.(*sql.WindowSpecification))
	case sql.WithClauseKind:
		p.visitWithClause(ctx, n.(*sql.WithClause))
	case sql.WithClauseEntryKind:
		p.visitWithClauseEntry(ctx, n.(*sql.WithClauseEntry))
	case sql.WithConnectionClauseKind:
		p.visitWithConnectionClause(ctx, n.(*sql.WithConnectionClause))
	case sql.WithExpressionKind:
		p.visitWithExpression(ctx, n.(*sql.WithExpression))
	case sql.WithOffsetKind:
		p.visitWithOffset(ctx, n.(*sql.WithOffset))
	case sql.WithPartitionColumnsClauseKind:
		p.visitWithPartitionColumnsClause(ctx, n.(*sql.WithPartitionColumnsClause))
	case sql.WithWeightKind:
		p.visitWithWeight(ctx, n.(*sql.WithWeight))
	case sql.InferredTypeColumnSchemaKind:
		p.visitColumnSchema(ctx, n.(*sql.ColumnSchema))
	default:
		p.addError(&Error{
			Err:  nil,
			Msg:  fmt.Sprintf("not implemented for kind %v", n.Kind()),
			Node: n,
		})
	}
}

func (p *Printer) addError(err error) {
	p.err = multierror.Append(p.err, err)
	log.Println("[ERROR]", err)
}

func (p *Printer) moveBefore(n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	p.Writer.flushCommentsUpTo(n.LocationStart())
}

func (p *Printer) movePast(n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	p.Writer.flushCommentsUpTo(n.LocationEnd())
}

func (p *Printer) moveAt(pos int) {
	p.Writer.flushCommentsUpTo(pos)
}

// movePastLine scans from the end of a node to the end of the line or
// until the next node.
// We do this limited to the end of the parent's end location.
func (p *Printer) movePastLine(n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	e := n.LocationEnd()
	newlinePos := p.Tracker.Lines.NextLineBreak(e)
	b := p.Tracker.MaybeNextPos(e)
	if b == -1 || newlinePos == -1 {
		parent := n.Parent()
		if parent == nil || parent.Kind() == sql.ScriptKind {
			if newlinePos > 0 {
				p.Writer.flushCommentsUpTo(newlinePos)
			} else {
				p.Writer.flushCommentsUpTo(len(p.OriginalInput))
			}
		}
		return
	}
	if newlinePos < b {
		p.Writer.flushCommentsUpTo(newlinePos)
	}
}

// moveBeforeSuccessorOf moves cursor to before the start of the
// succeeding start position.
func (p *Printer) moveBeforeSuccessorOf(n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	e := n.LocationEnd()
	max := e
	parent := n.Parent()
	for parent != nil && parent.Kind() != sql.StatementListKind {
		max = parent.LocationEnd()
		parent = parent.Parent()
	}
	next := min(p.Tracker.MaybeNextPos(e), max)
	if next > 0 {
		p.Writer.flushCommentsUpTo(next)
	}
}

func (p *Printer) print(s string) {
	p.Writer.Format(s)
}

func (p *Printer) println(s string) {
	p.Writer.FormatLine(s)
}

func (p *Printer) String() string {
	p.Writer.FlushLine()
	return strings.Trim(p.Writer.formatted.String(), "\n")
}

func (p *Printer) incDepth() {
	p.Writer.depth++
}

func (p *Printer) decDepth() {
	p.Writer.depth--
}

// nest returns a new printer with the same options to perform printing
// on a nested section of the tree.
func (p *Printer) nest() *Printer {
	buf := p.Writer.buffer.String()
	currSize := strings.LastIndex(buf, "\n")
	if currSize < 0 {
		currSize = len(p.Writer.buffer.String())
	}
	capacity := p.Writer.opts.SoftMaxColumns - currSize
	// Some scripts with lots of nested printers could lead to very
	// small or even negative maximum length.  We allow at least some
	// characters per-line at any given nested level.
	capacity = max(capacity, 80)
	n := &Printer{
		Writer:        p.Writer.WithCapacity(capacity),
		OriginalInput: p.OriginalInput,
		ErasedInput:   p.ErasedInput,
		Tracker:       p.Tracker,
	}
	return n
}

// unnest flushes the buffer and returns the strings with alignment
// symbols at the beginning of each line.
func (p *Printer) unnest() string {
	trimmed := p.String()
	aligned := alignNested(trimmed)
	aligned = "\v" + aligned
	aligned = strings.ReplaceAll(aligned, "\n", "\n\v")
	return aligned
}

// unnest flushes the buffer and returns the strings with alignment
// symbols at the beginning of each line.
func (p *Printer) unnestWithDepth(d int) string {
	trimmed := p.String()
	aligned := alignNested(trimmed)
	aligned = "\v" + aligned
	alignment := strings.Repeat("\v", d)
	aligned = strings.ReplaceAll(aligned, "\n", "\n"+alignment)
	return aligned
}

// printNestedWithSepNode prints sql.Node items separated by sep in a shared
// nested printer.
func printNestedWithSepNode[T sql.Node](ctx Context, p *Printer, items []T, sep string) {
	pp := p.nest()
	for i, item := range items {
		if i > 0 {
			pp.print(sep)
		}
		pp.acceptNested(ctx, item)
	}
	p.print(pp.unnest())
}

// printlnNestedWithSepNode prints sql.Node items separated by sep, one per line.
func printlnNestedWithSepNode[T sql.Node](ctx Context, p *Printer, items []T, sep string) {
	pp := p.nest()
	for i, item := range items {
		if i > 0 {
			pp.println(sep)
		}
		pp.acceptNested(ctx, item)
	}
	s := pp.unnestLeft()
	if strings.Trim(s, "\n\v\t") == "" {
		return
	}
	p.print(s)
}

// acceptNested visits a sql.Node with a nested printer.
func (p *Printer) acceptNested(ctx Context, n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	pp := p.nest()
	pp.accept(ctx, n)
	p.print(pp.unnest())
}

// acceptNestedLeft visits a sql.Node with a nested printer, unnested left.
func (p *Printer) acceptNestedLeft(ctx Context, n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	pp := p.nest()
	pp.accept(ctx, n)
	s := pp.unnestLeft()
	if strings.Trim(s, "\n\v\t") == "" {
		return
	}
	p.print(s)
}

// acceptNestedString visits a sql.Node with a nested printer, printing as string.
func (p *Printer) acceptNestedString(ctx Context, n sql.Node) {
	if !sql.Defined(n) {
		return
	}
	pp := p.nest()
	pp.accept(ctx, n)
	p.print(pp.String())
}

// toString visits a sql.Node with a nested printer and returns its string.
func (p *Printer) toString(ctx Context, n sql.Node) string {
	if !sql.Defined(n) {
		return ""
	}
	pp := p.nest()
	pp.accept(ctx, n)
	return pp.String()
}

// toUnnestedString visits a sql.Node with a nested printer and returns the unnested string.
func (p *Printer) toUnnestedString(ctx Context, n sql.Node) string {
	if !sql.Defined(n) {
		return ""
	}
	pp := p.nest()
	pp.accept(ctx, n)
	return pp.unnest()
}

func debugContent(s string) string {
	d := strings.ReplaceAll(s, "\v", "|")
	d = strings.ReplaceAll(d, "\b", "%")
	return d
}

// unnest flushes the buffer and returns the strings with alignment
// symbols at the beginning of each line.
func (p *Printer) unnestLeft() string {
	aligned := leftAlignNested(p.String())
	return "\v" + strings.ReplaceAll(aligned, "\n", "\n\v")
}

func (p *Printer) printOpenParenIfNeeded(n sql.Node) {
	if p.isParenNeeded(n) {
		p.print("(")
		if qe, ok := n.Raw().(interface{ IsQueryExpression() (bool, error) }); ok && mustBool(qe.IsQueryExpression()) {
			p.println("")
			p.incDepth()
		}
	}
}

func (p *Printer) printCloseParenIfNeeded(n sql.Node) {
	if p.isParenNeeded(n) {
		if qe, ok := n.Raw().(interface{ IsQueryExpression() (bool, error) }); ok && mustBool(qe.IsQueryExpression()) {
			p.println("")
			p.decDepth()
		}
		p.print(")")
	}
}

func (p *Printer) printOpenParenIfNeededWithDepth(n sql.Node) {
	if p.isParenNeeded(n) {
		p.print("(")
		p.println("")
		p.incDepth()
	}
}

func (p *Printer) printCloseParenIfNeededWithDepth(n sql.Node) {
	if p.isParenNeeded(n) {
		p.println("")
		p.decDepth()
		p.print(")")
	}
}

func mustBool(v bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return v
}

func (p *Printer) isParenNeeded(n sql.Node) bool {
	if !sql.Defined(n) {
		return false
	}
	parent := n.Parent()
	if n.Parenthesized() {
		if n.Kind() == sql.QueryKind {
			if parent != nil {
				switch parent.Kind() {
				case sql.CreateTableStatementKind,
					sql.CreateViewStatementKind,
					sql.CreateMaterializedViewStatementKind,
					sql.CreateTableFunctionStatementKind:
					return false
				}
			}
		}
		return true
	}
	if eval, ok := hasLowerPrecedence(parent, n); ok && eval {
		return true
	}
	return false
}

// hasParenAround checks if there is parenthesis just before the start
// location of a node.
func (p *Printer) hasParenAround(n sql.Node) bool {
	if !sql.Defined(n) {
		return false
	}
	s := n.LocationStart()
	if s == 0 {
		return false
	}
	return p.OriginalInput[s-1] == '('
}

func (p *Printer) printClause(s string) {
	if p.Writer.opts.NewlineBeforeClause {
		p.println("")
	}
	p.print(s)
}

func (p *Printer) identifier(s string) string {
	if len(s) > 0 && s[0] == '`' {
		return s
	}
	if p.Writer.opts.IsPseudoColumn(s) {
		return p.identifierWithCase(s, p.Writer.opts.PseudoColumnStyle)
	}
	return p.identifierWithCase(s, p.Writer.opts.IdentifierStyle)
}

func (p *Printer) tableName(s string) string {
	if len(s) > 0 && s[0] == '`' {
		return s
	}
	switch p.Writer.opts.TableNameStyle {
	case format.AsIs:
		return s
	case format.UpperCase:
		return strings.ToUpper(s)
	case format.LowerCase:
		return strings.ToLower(s)
	}
	return ""
}

func (p *Printer) functionName(s string) string {
	if len(s) > 0 && s[0] == '`' {
		return s
	}
	if p.Writer.opts.FunctionCatalog == format.BigQueryCatalog {
		if name := bigqueryFunctions.GetWithFallback(s, ""); name != "" {
			return p.applyStyle(name, p.Writer.opts.BuiltinFunctionNameStyle)
		}
	}
	return p.applyStyle(s, p.Writer.opts.FunctionNameStyle)
}

func (p *Printer) applyStyle(s string, style format.PrintCase) string {
	if s[0] == '`' {
		return s
	}
	switch style {
	case format.AsIs, format.Unspecified:
		return s
	case format.UpperCase:
		return strings.ToUpper(s)
	case format.LowerCase:
		return strings.ToLower(s)
	}
	return ""
}

func (p *Printer) keyword(s string) string {
	switch p.Writer.opts.KeywordStyle {
	case format.AsIs:
		return s
	case format.UpperCase:
		return strings.ToUpper(s)
	case format.LowerCase:
		return strings.ToLower(s)
	}
	return ""
}

func (p *Printer) typename(s string) string {
	if len(s) == 0 || s[0] == '`' {
		return s
	}
	switch p.Writer.opts.TypeStyle {
	case format.AsIs:
		return s
	case format.UpperCase:
		return strings.ToUpper(s)
	case format.LowerCase:
		return strings.ToLower(s)
	}
	return ""
}

// identifierWithCase prints according to specified case, and falls back
// to default identifier definition.  This is used to render specific
// function arguments.
func (p *Printer) identifierWithCase(s string, c format.PrintCase) string {
	if s[0] == '`' {
		return s
	}
	switch c {
	case format.AsIs:
		return s
	case format.UpperCase:
		return strings.ToUpper(s)
	case format.LowerCase:
		return strings.ToLower(s)
	case format.Unspecified:
		return p.identifier(s)
	}
	return ""
}

func (p *Printer) systemVariable(s string) string {
	if s[0] == '`' {
		return s
	}
	switch p.Writer.opts.SystemVariableStyle {
	case format.AsIs:
		return s
	case format.UpperCase:
		return strings.ToUpper(s)
	case format.LowerCase:
		return strings.ToLower(s)
	case format.Unspecified:
		return p.identifier(s)
	}
	return ""
}

func (p *Printer) queryParameter(s string) string {
	if s[0] == '`' {
		return s
	}
	switch p.Writer.opts.QueryParameterStyle {
	case format.AsIs:
		return s
	case format.UpperCase:
		return strings.ToUpper(s)
	case format.LowerCase:
		return strings.ToLower(s)
	case format.Unspecified:
		return p.identifier(s)
	}
	return ""
}

func (p *Printer) nodeInput(n sql.Node) string {
	b, e := n.Location()
	return p.viewInput(b, e)
}

func (p *Printer) nodeErasedInput(n sql.Node) string {
	b, e := n.Location()
	return p.viewErasedInput(b, e)
}

// viewErasedInput safely returns the input within the interval
// [begin, end). Note that input is not necessarily available, and this
// method may return empty without an error.
func (p *Printer) viewErasedInput(begin, end int) string {
	if end >= len(p.ErasedInput) {
		log.Println("[ERROR] Out of bounds on erased input.")
		return ""
	}
	return p.ErasedInput[begin:end]
}

// viewInput safely returns the input within the interval [begin, end).
// Note that input is not necessarily available, and this method may
// return empty without an error.
func (p *Printer) viewInput(begin, end int) string {
	if end >= len(p.OriginalInput) {
		return ""
	}
	return p.OriginalInput[begin:end]
}

func alignNested(s string) string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 0, ' ', tabwriter.AlignRight)
	fmt.Fprint(w, s)
	w.Flush()
	return strings.Trim(buf.String(), "\n")
}

func leftAlignNested(s string) string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 0, ' ', 0)
	fmt.Fprint(w, s)
	w.Flush()
	return strings.Trim(buf.String(), "\n")
}

// rowsTrimRight returns string where each row has been right-trimmed.
// Note that this procedure applies to all rows in the string,
// disregarding its contents, so in case of a SQL code we will change
// the contents of strings as well.
func rowsTrimRight(s string) string {
	rows := strings.Split(s, "\n")
	r := make([]string, 0, len(rows))
	for _, row := range rows {
		r = append(r, strings.TrimRight(row, " "))
	}
	return strings.Join(r, "\n")
}

func hasLowerPrecedence(parent, child sql.Node) (eval bool, ok bool) {
	p := precedenceNum(parent)
	c := lowestPrecedenceBelow(child)
	eval = p > 0 && c > 0 && p > c
	ok = p < 1000 && c < 1000
	return
}

func lowestPrecedenceBelow(n sql.Node) int {
	if n == nil {
		return 1000
	}
	if n.Kind() != sql.BinaryExpressionKind {
		return 1000
	}
	t := n.(*sql.BinaryExpression)
	min := precedenceNum(n)
	if lhs := precedenceNum(t.LHS()); lhs < min {
		min = lhs
	}
	if rhs := precedenceNum(t.RHS()); rhs < min {
		min = rhs
	}
	return min
}

func precedenceNum(n sql.Node) int {
	if n == nil {
		return 1000
	}
	switch n.Kind() {
	case sql.DotStarKind:
		return 1
	case sql.OrExprKind:
		return 2
	case sql.AndExprKind:
		return 3
	case sql.UnaryExpressionKind:
		return precedenceUnaryExpr(n.(*sql.UnaryExpression))
	case sql.BinaryExpressionKind:
		return precedenceBinExpr(n.(*sql.BinaryExpression))
	case sql.BetweenExpressionKind:
		return 5
	}
	return 1000
}

func precedenceBinExpr(n *sql.BinaryExpression) int {
	switch n.Op() {
	case sql.NotSetBinaryOp:
		return -1
	case sql.EqOp, sql.NEOp, sql.NE2Op, sql.GTOp, sql.GEOp,
		sql.LTOp, sql.LEOp, sql.LikeOp, sql.DistinctOp, sql.IsOp:
		return 5
	case sql.BitwiseOrOp:
		return 6
	case sql.BitwiseXorOp:
		return 7
	case sql.BitwiseAndOp:
		return 8
	case sql.PlusBinaryOp, sql.MinusBinaryOp:
		return 9
	case sql.ConcatOpOp:
		return 10
	case sql.MultiplyOp, sql.DivideOp:
		return 11
	}
	return 1000
}

func precedenceUnaryExpr(n *sql.UnaryExpression) int {
	switch n.Op() {
	case sql.NotSetUnaryOp:
		return 0
	case sql.NotUnaryOp:
		return 4
	case sql.BitwiseNotOp, sql.MinusUnaryOp, sql.PlusUnaryOp,
		sql.IsUnknownOp, sql.IsNotUnknownOp:
		return 12
	}
	return -1
}

func childrenExpressions(n sql.Node) []sql.ExpressionNode {
	if !sql.Defined(n) {
		return nil
	}
	var result []sql.ExpressionNode
	for _, c := range n.Children() {
		if e, ok := c.(sql.ExpressionNode); ok {
			result = append(result, e)
		}
	}
	return result
}
