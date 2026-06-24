package printer

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/tabwriter"

	"github.com/goccy/go-googlesql"
	"github.com/hashicorp/go-multierror"

	"github.com/paulourio/gsql/format"
	"github.com/paulourio/gsql/internal/ast"
)

type Printer struct {
	OriginalInput string
	ErasedInput   string
	Writer        *Writer
	Tracker       *LocationTracker

	err error
}

func (p *Printer) Print(root googlesql.ASTNode) (string, error) {
	ctx := &emptyCtx{}
	p.accept(ctx, root)
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

// accept visits a node on current line.
func (p *Printer) accept(ctx Context, n googlesql.ASTNode) {
	p.visit(ctx, n, false)
}

// accept visits a node in a new line. If the node is not defined, no
// line is created.
func (p *Printer) lnaccept(ctx Context, n googlesql.ASTNode) {
	p.visit(ctx, n, true)
}

func (p *Printer) visit(ctx Context, n googlesql.ASTNode, newline bool) {
	if !ast.Defined(n) {
		return
	}
	if newline {
		p.println("")
	}
	switch m := n.(type) {
	case *googlesql.ASTAddColumnAction:
		p.VisitAddColumnAction(ctx, m)
	case *googlesql.ASTAddConstraintAction:
		p.VisitAddConstraintAction(ctx, m)
	case *googlesql.ASTAlias:
		p.VisitAlias(ctx, m)
	case *googlesql.ASTAliasedGroupRows:
		p.VisitAliasedGroupRows(ctx, m)
	case *googlesql.ASTAliasedQuery:
		p.VisitAliasedQuery(ctx, m)
	case *googlesql.ASTAlterActionList:
		p.VisitAlterActionList(ctx, m)
	case *googlesql.ASTAlterAllRowAccessPoliciesStatement:
		p.VisitAlterAllRowAccessPoliciesStatement(ctx, m)
	case *googlesql.ASTAlterColumnDropDefaultAction:
		p.VisitAlterColumnDropDefaultAction(ctx, m)
	case *googlesql.ASTAlterColumnDropNotNullAction:
		p.VisitAlterColumnDropNotNullAction(ctx, m)
	case *googlesql.ASTAlterColumnOptionsAction:
		p.VisitAlterColumnOptionsAction(ctx, m)
	case *googlesql.ASTAlterColumnSetDefaultAction:
		p.VisitAlterColumnSetDefaultAction(ctx, m)
	case *googlesql.ASTAlterColumnTypeAction:
		p.VisitAlterColumnTypeAction(ctx, m)
	case *googlesql.ASTAlterConstraintEnforcementAction:
		p.VisitAlterConstraintEnforcementAction(ctx, m)
	case *googlesql.ASTAlterConstraintSetOptionsAction:
		p.VisitAlterConstraintSetOptionsAction(ctx, m)
	case *googlesql.ASTAlterDatabaseStatement:
		p.VisitAlterDatabaseStatement(ctx, m)
	case *googlesql.ASTAlterEntityStatement:
		p.VisitAlterEntityStatement(ctx, m)
	case *googlesql.ASTAlterMaterializedViewStatement:
		p.VisitAlterMaterializedViewStatement(ctx, m)
	case *googlesql.ASTAlterPrivilegeRestrictionStatement:
		p.VisitAlterPrivilegeRestrictionStatement(ctx, m)
	case *googlesql.ASTAlterRowAccessPolicyStatement:
		p.VisitAlterRowAccessPolicyStatement(ctx, m)
	case *googlesql.ASTAlterSchemaStatement:
		p.VisitAlterSchemaStatement(ctx, m)
	case *googlesql.ASTAlterTableStatement:
		p.VisitAlterTableStatement(ctx, m)
	case *googlesql.ASTAlterViewStatement:
		p.VisitAlterViewStatement(ctx, m)
	case *googlesql.ASTAnalyticFunctionCall:
		p.VisitAnalyticFunctionCall(ctx, m)
	case *googlesql.ASTAndExpr:
		p.VisitAndExpr(ctx, m)
	case *googlesql.ASTArrayConstructor:
		p.VisitArrayConstructor(ctx, m)
	case *googlesql.ASTArrayColumnSchema:
		p.VisitArrayColumnSchema(ctx, m)
	case *googlesql.ASTArrayElement:
		p.VisitArrayElement(ctx, m)
	case *googlesql.ASTArrayType:
		p.VisitArrayType(ctx, m)
	case *googlesql.ASTAssignmentFromStruct:
		p.VisitAssignmentFromStruct(ctx, m)
	case *googlesql.ASTBeginEndBlock:
		p.VisitBeginEndBlock(ctx, m)
	case *googlesql.ASTBeginStatement:
		p.VisitBeginStatementNode(ctx, m)
	case *googlesql.ASTBetweenExpression:
		p.VisitBetweenExpression(ctx, m)
	case *googlesql.ASTBigNumericLiteral:
		p.VisitBigNumericLiteral(ctx, m)
	case *googlesql.ASTBinaryExpression:
		p.VisitBinaryExpression(ctx, m)
	case *googlesql.ASTBitwiseShiftExpression:
		p.VisitBitwiseShiftExpression(ctx, m)
	case *googlesql.ASTBooleanLiteral:
		p.VisitBoolLiteral(ctx, m)
	case *googlesql.ASTBytesLiteral:
		p.VisitBytesLiteral(ctx, m)
	case *googlesql.ASTCallStatement:
		p.VisitCallStatement(ctx, m)
	case *googlesql.ASTCaseNoValueExpression:
		p.VisitCaseNoValueExpression(ctx, m)
	case *googlesql.ASTCaseValueExpression:
		p.VisitCaseValueExpression(ctx, m)
	case *googlesql.ASTCastExpression:
		p.VisitCastExpression(ctx, m)
	case *googlesql.ASTClampedBetweenModifier:
		p.VisitClampedBetweenModifier(ctx, m)
	case *googlesql.ASTCloneDataSource:
		p.VisitCloneDataSource(ctx, m)
	case *googlesql.ASTClusterBy:
		p.VisitClusterBy(ctx, m)
	case *googlesql.ASTCollate:
		p.VisitCollate(ctx, m)
	case *googlesql.ASTColumnAttributeList:
		p.VisitColumnAttributeList(ctx, m)
	case *googlesql.ASTColumnDefinition:
		p.VisitColumnDefinition(ctx, m)
	case *googlesql.ASTColumnList:
		p.VisitColumnList(ctx, m)
	case *googlesql.ASTColumnSchema:
		p.VisitColumnSchema(ctx, m)
	case *googlesql.ASTColumnWithOptions:
		p.VisitColumnWithOptions(ctx, m)
	case *googlesql.ASTColumnWithOptionsList:
		p.VisitColumnWithOptionsList(ctx, m)
	case *googlesql.ASTCommitStatement:
		p.VisitCommitStatement(ctx, m)
	case *googlesql.ASTConnectionClause:
		p.VisitConnectionClause(ctx, m)
	case *googlesql.ASTCopyDataSource:
		p.VisitCopyDataSource(ctx, m)
	case *googlesql.ASTCreateExternalTableStatement:
		p.VisitCreateExternalTableStatement(ctx, m)
	case *googlesql.ASTCreateFunctionStatement:
		p.VisitCreateFunctionStatement(ctx, m)
	case *googlesql.ASTCreateMaterializedViewStatement:
		p.VisitCreateMaterializedViewStatement(ctx, m)
	case *googlesql.ASTCreateProcedureStatement:
		p.VisitCreateProcedureStatement(ctx, m)
	case *googlesql.ASTCreateRowAccessPolicyStatement:
		p.VisitCreateRowAccessPolicyStatement(ctx, m)
	case *googlesql.ASTCreateSchemaStatement:
		p.VisitCreateSchemaStatement(ctx, m)
	case *googlesql.ASTCreateSnapshotTableStatement:
		p.VisitCreateSnapshotTableStatement(ctx, m)
	case *googlesql.ASTCreateTableStatement:
		p.VisitCreateTableStatement(ctx, m)
	case *googlesql.ASTCreateTableFunctionStatement:
		p.VisitCreateTableFunctionStatement(ctx, m)
	case *googlesql.ASTCreateViewStatement:
		p.VisitCreateViewStatement(ctx, m)
	case *googlesql.ASTDateOrTimeLiteral:
		p.VisitDateOrTimeLiteral(ctx, m)
	case *googlesql.ASTDescriptor:
		p.VisitDescriptor(ctx, m)
	case *googlesql.ASTDescriptorColumn:
		p.VisitDescriptorColumn(ctx, m)
	case *googlesql.ASTDescriptorColumnList:
		p.VisitDescriptorColumnList(ctx, m)
	case *googlesql.ASTDotIdentifier:
		p.VisitDotIdentifier(ctx, m)
	case *googlesql.ASTDotGeneralizedField:
		p.VisitDotGeneralizedField(ctx, m)
	case *googlesql.ASTDotStar:
		p.VisitDotStar(ctx, m)
	case *googlesql.ASTDotStarWithModifiers:
		p.VisitDotStarWithModifiers(ctx, m)
	case *googlesql.ASTDropAllRowAccessPoliciesStatement:
		p.VisitDropAllRowAccessPoliciesStatement(ctx, m)
	case *googlesql.ASTDropColumnAction:
		p.VisitDropColumnAction(ctx, m)
	case *googlesql.ASTDropConstraintAction:
		p.VisitDropConstraintAction(ctx, m)
	case *googlesql.ASTDropEntityStatement:
		p.VisitDropEntityStatement(ctx, m)
	case *googlesql.ASTDropFunctionStatement:
		p.VisitDropFunctionStatement(ctx, m)
	case *googlesql.ASTDropMaterializedViewStatement:
		p.VisitDropMaterializedViewStatement(ctx, m)
	case *googlesql.ASTDropPrimaryKeyAction:
		p.VisitDropPrimaryKeyAction(ctx, m)
	case *googlesql.ASTDropPrivilegeRestrictionStatement:
		p.VisitDropPrivilegeRestrictionStatement(ctx, m)
	case *googlesql.ASTDropRowAccessPolicyStatement:
		p.VisitDropRowAccessPolicyStatement(ctx, m)
	case *googlesql.ASTDropSearchIndexStatement:
		p.VisitDropSearchIndexStatement(ctx, m)
	case *googlesql.ASTDropSnapshotTableStatement:
		p.VisitDropSnapshotTableStatement(ctx, m)
	case *googlesql.ASTDropTableFunctionStatement:
		p.VisitDropTableFunctionStatement(ctx, m)
	case *googlesql.ASTDropStatement:
		p.VisitDropStatement(ctx, m)
	case *googlesql.ASTExceptionHandlerList:
		p.VisitExceptionHandlerListNode(ctx, m)
	case *googlesql.ASTExceptionHandler:
		p.VisitExceptionHandlerNode(ctx, m)
	case *googlesql.ASTExecuteIntoClause:
		p.VisitExecuteIntoClause(ctx, m)
	case *googlesql.ASTExecuteImmediateStatement:
		p.VisitExecuteImmediateStatement(ctx, m)
	case *googlesql.ASTExecuteUsingArgument:
		p.VisitExecuteUsingArgument(ctx, m)
	case *googlesql.ASTExecuteUsingClause:
		p.VisitExecuteUsingClause(ctx, m)
	case *googlesql.ASTExpressionSubquery:
		p.VisitExpressionSubquery(ctx, m)
	case *googlesql.ASTExtractExpression:
		p.VisitExtractExpression(ctx, m)
	case *googlesql.ASTFloatLiteral:
		p.VisitFloatLiteral(ctx, m)
	case *googlesql.ASTFilterUsingClause:
		p.VisitFilterUsingClause(ctx, m)
	case *googlesql.ASTForeignKey:
		p.VisitForeignKey(ctx, m)
	case *googlesql.ASTForeignKeyReference:
		p.VisitForeignKeyReference(ctx, m)
	case *googlesql.ASTFormatClause:
		p.VisitFormatClause(ctx, m)
	case *googlesql.ASTForSystemTime:
		p.VisitForSystemTime(ctx, m)
	case *googlesql.ASTFromClause:
		p.VisitFromClause(ctx, m)
	case *googlesql.ASTFunctionCall:
		p.VisitFunctionCall(ctx, m)
	case *googlesql.ASTFunctionDeclaration:
		p.VisitFunctionDeclaration(ctx, m)
	case *googlesql.ASTFunctionParameter:
		p.VisitFunctionParameter(ctx, m)
	case *googlesql.ASTFunctionParameters:
		p.VisitFunctionParameters(ctx, m)
	case *googlesql.ASTGeneralizedPathExpression:
		p.VisitGeneralizedPathExpression(ctx, m)
	case *googlesql.ASTGranteeList:
		p.VisitGranteeList(ctx, m)
	case *googlesql.ASTGrantToClause:
		p.VisitGrantToClause(ctx, m)
	case *googlesql.ASTGroupBy:
		p.VisitGroupBy(ctx, m)
	case *googlesql.ASTGroupingItem:
		p.VisitGroupingItem(ctx, m)
	case *googlesql.ASTHavingModifier:
		p.VisitHavingModifier(ctx, m)
	case *googlesql.ASTHaving:
		p.VisitHaving(ctx, m)
	case *googlesql.ASTHint:
		p.VisitHint(ctx, m)
	case *googlesql.ASTHintedStatement:
		p.VisitHintedStatement(ctx, m)
	case *googlesql.ASTIdentifier:
		p.VisitIdentifier(ctx, m)
	case *googlesql.ASTIdentifierList:
		p.VisitIdentifierList(ctx, m)
	case *googlesql.ASTIfStatement:
		p.VisitIfStatement(ctx, m)
	case *googlesql.ASTInExpression:
		p.VisitInExpression(ctx, m)
	case *googlesql.ASTInList:
		p.VisitInList(ctx, m)
	case *googlesql.ASTIntervalExpr:
		p.VisitIntervalExpr(ctx, m)
	case *googlesql.ASTIntLiteral:
		p.VisitIntLiteral(ctx, m)
	case *googlesql.ASTInsertStatement:
		p.VisitInsertStatement(ctx, m)
	case *googlesql.ASTInsertValuesRowList:
		p.VisitInsertValuesRowList(ctx, m)
	case *googlesql.ASTInsertValuesRow:
		p.VisitInsertValuesRow(ctx, m)
	case *googlesql.ASTJoin:
		p.VisitJoin(ctx, m)
	case *googlesql.ASTJSONLiteral:
		p.VisitJSONLiteral(ctx, m)
	case *googlesql.ASTLimit:
		p.VisitLimit(ctx, m)
	case *googlesql.ASTLimitOffset:
		p.VisitLimitOffset(ctx, m)
	case *googlesql.ASTMergeAction:
		p.VisitMergeAction(ctx, m)
	case *googlesql.ASTMergeStatement:
		p.VisitMergeStatement(ctx, m)
	case *googlesql.ASTMergeWhenClause:
		p.VisitMergeWhenClause(ctx, m)
	case *googlesql.ASTMergeWhenClauseList:
		p.VisitMergeWhenClauseList(ctx, m)
	case *googlesql.ASTModelClause:
		p.VisitModelClause(ctx, m)
	case *googlesql.ASTNamedArgument:
		p.VisitNamedArgument(ctx, m)
	case *googlesql.ASTNotNullColumnAttribute:
		p.VisitNotNullColumnAttribute(ctx, m)
	case *googlesql.ASTNullLiteral:
		p.VisitNullLiteral(ctx, m)
	case *googlesql.ASTNullOrder:
		p.VisitNullOrder(ctx, m)
	case *googlesql.ASTNumericLiteral:
		p.VisitNumericLiteral(ctx, m)
	case *googlesql.ASTOnClause:
		p.VisitOnClause(ctx, m)
	case *googlesql.ASTOptionsList:
		p.VisitOptionsList(ctx, m)
	case *googlesql.ASTOptionsEntry:
		p.VisitOptionsEntry(ctx, m)
	case *googlesql.ASTOrExpr:
		p.VisitOrExpr(ctx, m)
	case *googlesql.ASTOrderBy:
		p.VisitOrderBy(ctx, m)
	case *googlesql.ASTOrderingExpression:
		p.VisitOrderingExpression(ctx, m)
	case *googlesql.ASTParameterAssignment:
		p.VisitParameterAssignment(ctx, m)
	case *googlesql.ASTParameterExpr:
		p.VisitParameterExpr(ctx, m)
	case *googlesql.ASTParenthesizedJoin:
		p.VisitParenthesizedJoin(ctx, m)
	case *googlesql.ASTPartitionBy:
		p.VisitPartitionBy(ctx, m)
	case *googlesql.ASTPathExpressionList:
		p.VisitPathExpressionList(ctx, m)
	case *googlesql.ASTPathExpression:
		p.VisitPathExpression(ctx, m)
	case *googlesql.ASTPivotClause:
		p.VisitPivotClause(ctx, m)
	case *googlesql.ASTPivotExpression:
		p.VisitPivotExpression(ctx, m)
	case *googlesql.ASTPivotExpressionList:
		p.VisitPivotExpressionList(ctx, m)
	case *googlesql.ASTPivotValue:
		p.VisitPivotValue(ctx, m)
	case *googlesql.ASTPivotValueList:
		p.VisitPivotValueList(ctx, m)
	case *googlesql.ASTPrimaryKey:
		p.VisitPrimaryKey(ctx, m)
	case *googlesql.ASTPrimaryKeyColumnAttribute:
		p.VisitPrimaryKeyColumnAttribute(ctx, m)
	case *googlesql.ASTPrimaryKeyElementList:
		p.VisitPrimaryKeyElementList(ctx, m)
	case *googlesql.ASTPrimaryKeyElement:
		p.VisitPrimaryKeyElement(ctx, m)
	case *googlesql.ASTQualify:
		p.VisitQualify(ctx, m)
	case *googlesql.ASTQuery:
		p.VisitQuery(ctx, m)
	case *googlesql.ASTQueryStatement:
		p.VisitQueryStatement(ctx, m)
	case *googlesql.ASTRenameColumnAction:
		p.VisitRenameColumnAction(ctx, m)
	case *googlesql.ASTRenameToClause:
		p.VisitRenameToClause(ctx, m)
	case *googlesql.ASTRepeatableClause:
		p.VisitRepeatableClause(ctx, m)
	case *googlesql.ASTReturnStatement:
		p.VisitReturnStatement(ctx, m)
	case *googlesql.ASTRollbackStatement:
		p.VisitRollbackStatementNode(ctx, m)
	case *googlesql.ASTRollup:
		p.VisitRollup(ctx, m)
	case *googlesql.ASTSampleClause:
		p.VisitSampleClause(ctx, m)
	case *googlesql.ASTSampleSize:
		p.VisitSampleSize(ctx, m)
	case *googlesql.ASTSampleSuffix:
		p.VisitSampleSuffix(ctx, m)
	case *googlesql.ASTSetCollateClause:
		p.VisitSetCollateClause(ctx, m)
	case *googlesql.ASTScript:
		p.VisitScript(ctx, m)
	case *googlesql.ASTSelect:
		p.VisitSelect(ctx, m)
	case *googlesql.ASTSelectAs:
		p.VisitSelectAs(ctx, m)
	case *googlesql.ASTSelectColumn:
		p.VisitSelectColumn(ctx, m)
	case *googlesql.ASTSelectList:
		p.VisitSelectList(ctx, m)
	case *googlesql.ASTSetOptionsAction:
		p.VisitSetOptionsAction(ctx, m)
	case *googlesql.ASTSetOperation:
		p.VisitSetOperation(ctx, m)
	case *googlesql.ASTSimpleColumnSchema:
		p.VisitSimpleColumnSchema(ctx, m)
	case *googlesql.ASTSimpleType:
		p.VisitSimpleType(ctx, m)
	case *googlesql.ASTSqlFunctionBody:
		p.VisitSQLFunctionBody(ctx, m)
	case *googlesql.ASTStar:
		p.VisitStar(ctx, m)
	case *googlesql.ASTStarModifiers:
		p.VisitStarModifiers(ctx, m)
	case *googlesql.ASTStarReplaceItem:
		p.VisitStarReplaceItem(ctx, m)
	case *googlesql.ASTStarWithModifiers:
		p.VisitStarWithModifiers(ctx, m)
	case *googlesql.ASTStatementList:
		p.VisitStatementList(ctx, m)
	case *googlesql.ASTStringLiteral:
		p.VisitStringLiteral(ctx, m)
	case *googlesql.ASTStructColumnField:
		p.VisitStructColumnField(ctx, m)
	case *googlesql.ASTStructColumnSchema:
		p.VisitStructColumnSchema(ctx, m)
	case *googlesql.ASTStructConstructorArg:
		p.VisitStructConstructorArg(ctx, m)
	case *googlesql.ASTStructConstructorWithKeyword:
		p.VisitStructConstructorWithKeyword(ctx, m)
	case *googlesql.ASTStructConstructorWithParens:
		p.VisitStructConstructorWithParens(ctx, m)
	case *googlesql.ASTStructField:
		p.VisitStructField(ctx, m)
	case *googlesql.ASTStructType:
		p.VisitStructType(ctx, m)
	case *googlesql.ASTSystemVariableAssignment:
		p.VisitSystemVariableAssignment(ctx, m)
	case *googlesql.ASTSystemVariableExpr:
		p.VisitSystemVariableExpr(ctx, m)
	case *googlesql.ASTTableClause:
		p.VisitTableClause(ctx, m)
	case *googlesql.ASTTableConstraint:
		p.VisitTableConstraint(ctx, m)
	case *googlesql.ASTTableElementList:
		p.VisitTableElementList(ctx, m)
	case *googlesql.ASTTablePathExpression:
		p.VisitTablePathExpression(ctx, m)
	case *googlesql.ASTTableSubquery:
		p.VisitTableSubquery(ctx, m)
	case *googlesql.ASTTemplatedParameterType:
		p.VisitTemplatedParameterType(ctx, m)
	case *googlesql.ASTTruncateStatement:
		p.VisitTruncateStatement(ctx, m)
	case *googlesql.ASTTVFArgument:
		p.VisitTVFArgument(ctx, m)
	case *googlesql.ASTTVF:
		p.VisitTVF(ctx, m)
	case *googlesql.ASTTVFSchema:
		p.VisitTVFSchema(ctx, m)
	case *googlesql.ASTTVFSchemaColumn:
		p.VisitTVFSchemaColumn(ctx, m)
	case *googlesql.ASTTypeParameterList:
		p.VisitTypeParameterList(ctx, m)
	case *googlesql.ASTUnpivotClause:
		p.VisitUnpivotClause(ctx, m)
	case *googlesql.ASTUnaryExpression:
		p.VisitUnaryExpression(ctx, m)
	case *googlesql.ASTUnpivotInItemLabel:
		p.VisitUnpivotInItemLabel(ctx, m)
	case *googlesql.ASTUnpivotInItemList:
		p.VisitUnpivotInItemList(ctx, m)
	case *googlesql.ASTUnpivotInItem:
		p.VisitUnpivotInItem(ctx, m)
	case *googlesql.ASTUnnestExpression:
		p.VisitUnnestExpression(ctx, m)
	case *googlesql.ASTUnnestExpressionWithOptAliasAndOffset:
		p.VisitUnnestExpressionWithOptAliasAndOffset(ctx, m)
	case *googlesql.ASTUpdateItem:
		p.VisitUpdateItem(ctx, m)
	case *googlesql.ASTUpdateItemList:
		p.VisitUpdateItemList(ctx, m)
	case *googlesql.ASTUpdateSetValue:
		p.VisitUpdateSetValue(ctx, m)
	case *googlesql.ASTUsingClause:
		p.VisitUsingClause(ctx, m)
	case *googlesql.ASTVariableDeclaration:
		p.VisitVariableDeclaration(ctx, m)
	case *googlesql.ASTSingleAssignment:
		p.VisitSingleAssignment(ctx, m)
	case *googlesql.ASTWhereClause:
		p.VisitWhereClause(ctx, m)
	case *googlesql.ASTWindowClause:
		p.VisitWindowClause(ctx, m)
	case *googlesql.ASTWindowFrame:
		p.VisitWindowFrame(ctx, m)
	case *googlesql.ASTWindowFrameExpr:
		p.VisitWindowFrameExpr(ctx, m)
	case *googlesql.ASTWindowSpecification:
		p.VisitWindowSpecification(ctx, m)
	case *googlesql.ASTWithClause:
		p.VisitWithClause(ctx, m)
	case *googlesql.ASTWithClauseEntry:
		p.VisitWithClauseEntry(ctx, m)
	case *googlesql.ASTWithConnectionClause:
		p.VisitWithConnectionClause(ctx, m)
	case *googlesql.ASTWithExpression:
		p.VisitWithExpression(ctx, m)
	case *googlesql.ASTWithOffset:
		p.VisitWithOffset(ctx, m)
	case *googlesql.ASTWithPartitionColumnsClause:
		p.VisitWithPartitionColumnsClause(ctx, m)
	case *googlesql.ASTWithWeight:
		p.VisitWithWeight(ctx, m)

	default:
		p.addError(&Error{
			Err:  nil,
			Msg:  fmt.Sprintf("not implemented for %#v", n),
			Node: n,
		})
	}
}

func (p *Printer) addError(err error) {
	p.err = multierror.Append(p.err, err)
	log.Println("[ERROR]", err)
}

func (p *Printer) moveBefore(n googlesql.ASTNode) {
	p.Writer.flushCommentsUpTo(ast.GetParseLocationStartOffset(n))
}

func (p *Printer) movePast(n googlesql.ASTNode) {
	p.Writer.flushCommentsUpTo(ast.GetParseLocationEndOffset(n))
}

func (p *Printer) moveAt(pos int) {
	p.Writer.flushCommentsUpTo(pos)
}

// movePastLine scans from the end of a node to the end of the line or
// until the next node.
// We do this limited to the end of the parent's end location.
func (p *Printer) movePastLine(n googlesql.ASTNode) {
	e := ast.GetParseLocationEndOffset(n)
	newlinePos := p.Tracker.Lines.NextLineBreak(e)
	b := p.Tracker.MaybeNextPos(e)
	if b == -1 || newlinePos == -1 {
		// Only flush comments if at the top level.
		parent := ast.Parent(n)
		if parent == nil || ast.Kind(parent) == googlesql.ASTNodeKindAstScript {
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

// moveBeforeSuccessorOf move cursor to before the start of the
// succeding start position.
func (p *Printer) moveBeforeSuccessorOf(n googlesql.ASTNode) {
	e := ast.GetParseLocationEndOffset(n)
	// Limit this kind of comment flush up to the statement level.
	// A statement list is flat aligned and we will not have problems
	// with comment indentation at that level or higher.
	max := e
	parent := ast.Parent(n)
	for ast.Defined(parent) && ast.Kind(parent) != ast.StatementList {
		max = ast.GetParseLocationEndOffset(parent)
		parent = ast.Parent(parent)
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

// printNestedWithSep receives a slice of googlesql.ASTNode items and print each
// in a nested printer.
// Since we cannot have generic methods, this is a function that receives
// a printer as the first argument.  Otherwise, it would be a method.
func printNestedWithSep[T googlesql.ASTNode](ctx Context, p *Printer, items []T, sep string) {
	pp := p.nest()
	for i, item := range items {
		if i > 0 {
			pp.print(sep)
		}
		pp.acceptNested(ctx, item)
	}
	p.print(pp.unnest())
}

// printlnNestedWithSep receives a slice of googlesql.ASTNode items and print each
// in a nested printer.
// Since we cannot have generic methods, this is a function that receives
// a printer as the first argument.  Otherwise, it would be a method.
func printlnNestedWithSep[T googlesql.ASTNode](ctx Context, p *Printer, items []T, sep string) {
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

// acceptNested visits a node with a nested printer.
func (p *Printer) acceptNested(ctx Context, n googlesql.ASTNode) {
	pp := p.nest()
	pp.accept(ctx, n)
	p.print(pp.unnest())
}

// acceptNestedLeft visits a node with a nested printer, and unnests
// result with left alignment.
func (p *Printer) acceptNestedLeft(ctx Context, n googlesql.ASTNode) {
	pp := p.nest()
	pp.accept(ctx, n)
	s := pp.unnestLeft()
	if strings.Trim(s, "\n\v\t") == "" {
		return
	}
	p.print(s)
}

// acceptNestedString visits a node with a nested printer, and prints
// result to current printer as a string.
func (p *Printer) acceptNestedString(ctx Context, n googlesql.ASTNode) {
	pp := p.nest()
	pp.accept(ctx, n)
	p.print(pp.String())
}

// toString visits a node with a nested printer and returns its
// string contents instead of writing to current printer.
func (p *Printer) toString(ctx Context, n googlesql.ASTNode) string {
	pp := p.nest()
	pp.accept(ctx, n)
	return pp.String()
}

// toUnnestedString visits a node with a nested printer and returns its
// unnested string contents .
func (p *Printer) toUnnestedString(ctx Context, n googlesql.ASTNode) string {
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

func (p *Printer) printOpenParenIfNeeded(n ast.ParethesizedNode) {
	if p.isParenNeeded(n) {
		p.print("(")
		if ast.Must(n.IsQueryExpression()) {
			p.println("")
			p.incDepth()
		}
	}
}

func (p *Printer) printCloseParenIfNeeded(n ast.ParethesizedNode) {
	if p.isParenNeeded(n) {
		if ast.Must(n.IsQueryExpression()) {
			p.println("")
			p.decDepth()
		}
		p.print(")")
	}
}

func (p *Printer) printOpenParenIfNeededWithDepth(n ast.ParethesizedNode) {
	if p.isParenNeeded(n) {
		p.print("(")
		p.println("")
		p.incDepth()
	}
}

func (p *Printer) printCloseParenIfNeededWithDepth(n ast.ParethesizedNode) {
	if p.isParenNeeded(n) {
		p.println("")
		p.decDepth()
		p.print(")")
	}
}

func (p *Printer) isParenNeeded(n ast.ParethesizedNode) bool {
	parent := ast.Parent(n)
	if ast.Must(n.Parenthesized()) {
		if ast.Kind(n) == ast.Query {
			// We force Create table statement to be wrapped with paranthesis,
			// which is generated by CreateTableStatement visitor.
			// So if we are visiting the query definition of a create table
			// we just return that we do not need to parenthesize.
			if ast.Defined(parent) {
				switch ast.Kind(parent) {
				case ast.CreateTableStatement,
					ast.CreateViewStatement,
					ast.CreateMaterializedViewStatement,
					ast.CreateTableFunctionStatement:
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
func (p *Printer) hasParenAround(n googlesql.ASTNode) bool {
	s := ast.GetParseLocationStartOffset(n)
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

func (p *Printer) nodeInput(n googlesql.ASTNode) string {
	b, e := ast.GetParseLocationByteOffsets(n)
	return p.viewInput(b, e)
}

func (p *Printer) nodeErasedInput(n googlesql.ASTNode) string {
	b, e := ast.GetParseLocationByteOffsets(n)
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

// isInsideOfMergeStatement returns true when the current node is is inside
// of a MERGE statement directly.
func isInsideOfMergeStatement(n googlesql.ASTNode) bool {
	p := ast.Must(n.Parent())
	if !ast.Defined(p) {
		return false
	}
	return ast.Must(p.NodeKind()) == googlesql.ASTNodeKindAstMergeStatement
}

// isInsideOfOnClause returns true when the current node is is inside
// of an ON clause directly. The node can be inside of other AndExpr
// and OrExpr.
func isInsideOfOnClause(n googlesql.ASTNode) bool {
	for p := ast.Must(n.Parent()); ast.Defined(p); p = ast.Must(p.Parent()) {
		kind := ast.Must(p.NodeKind())
		if kind == ast.OnClause {
			return true
		}
		if kind != ast.AndExpr && kind != ast.OrExpr {
			return false
		}
	}
	return false
}

// isInsideOfWhereClause returns true when the current node is inside
// of a WHERE clause directly. The node can be inside of other AndExpr
// and OrExpr.
func isInsideOfWhereClause(n googlesql.ASTNode) bool {
	for p := ast.Must(n.Parent()); ast.Defined(p); p = ast.Must(p.Parent()) {
		switch ast.Must(p.NodeKind()) {
		case ast.WhereClause:
			return true
		case ast.AndExpr, ast.OrExpr:
			continue
		default:
			return false
		}
	}
	return false
}

// hasLowerPrecedence returns whether child has lower precedence than
// parent.  This is mainly used to help on determine the necessity
// for parenthesis when rendering expressions.
func hasLowerPrecedence(parent, child googlesql.ASTNode) (eval bool, ok bool) {
	p := precedenceNum(parent)
	c := lowestPrecedenceBelow(child)
	eval = p > 0 && c > 0 && p > c
	ok = p < 1000 && c < 1000
	return
}

func lowestPrecedenceBelow(n googlesql.ASTNode) int {
	switch t := n.(type) {
	case *googlesql.ASTBinaryExpression:
		var (
			min int = precedenceNum(n)
			lhs int = precedenceNum(ast.Must(t.Lhs()))
			rhs int = precedenceNum(ast.Must(t.Rhs()))
		)
		if lhs < min {
			min = lhs
		}
		if rhs < min {
			min = rhs
		}
		return min
	default:
		return 1000
	}
}

func precedenceNum(n googlesql.ASTNode) int {
	switch ast.Kind(n) {
	case googlesql.ASTNodeKindAstDotStar:
		return 1
	case googlesql.ASTNodeKindAstOrExpr:
		return 2
	case googlesql.ASTNodeKindAstAndExpr:
		return 3
	case googlesql.ASTNodeKindAstUnaryExpression:
		return precedenceUnaryExpr(n.(*googlesql.ASTUnaryExpression))
	case googlesql.ASTNodeKindAstBinaryExpression:
		return precedenceBinExpr(n.(*googlesql.ASTBinaryExpression))
	case googlesql.ASTNodeKindAstBetweenExpression:
		return 5
	}
	return 1000
}

func precedenceBinExpr(n *googlesql.ASTBinaryExpression) int {
	switch ast.Must(n.Op()) {
	case googlesql.ASTBinaryExpressionEnums_OpNotSet:
		return -1
	case googlesql.ASTBinaryExpressionEnums_OpEq,
		googlesql.ASTBinaryExpressionEnums_OpNe,
		googlesql.ASTBinaryExpressionEnums_OpNe2,
		googlesql.ASTBinaryExpressionEnums_OpGt,
		googlesql.ASTBinaryExpressionEnums_OpGe,
		googlesql.ASTBinaryExpressionEnums_OpLt,
		googlesql.ASTBinaryExpressionEnums_OpLe,
		googlesql.ASTBinaryExpressionEnums_OpLike,
		googlesql.ASTBinaryExpressionEnums_OpDistinct,
		googlesql.ASTBinaryExpressionEnums_OpIs:
		return 5
	case googlesql.ASTBinaryExpressionEnums_OpBitwiseOr:
		return 6
	case googlesql.ASTBinaryExpressionEnums_OpBitwiseXor:
		return 7
	case googlesql.ASTBinaryExpressionEnums_OpBitwiseAnd:
		return 8
	case googlesql.ASTBinaryExpressionEnums_OpPlus,
		googlesql.ASTBinaryExpressionEnums_OpMinus:
		return 9
	case googlesql.ASTBinaryExpressionEnums_OpConcatOp:
		return 10
	case googlesql.ASTBinaryExpressionEnums_OpMultiply,
		googlesql.ASTBinaryExpressionEnums_OpDivide:
		return 11
	}
	return 1000
}

func precedenceUnaryExpr(n *googlesql.ASTUnaryExpression) int {
	switch ast.Must(n.Op()) {
	case googlesql.ASTUnaryExpressionEnums_OpNotSet:
		return 0
	case googlesql.ASTUnaryExpressionEnums_OpNot:
		return 4
	case googlesql.ASTUnaryExpressionEnums_OpBitwiseNot,
		googlesql.ASTUnaryExpressionEnums_OpMinus,
		googlesql.ASTUnaryExpressionEnums_OpPlus,
		googlesql.ASTUnaryExpressionEnums_OpIsUnknown,
		googlesql.ASTUnaryExpressionEnums_OpIsNotUnknown:
		return 12
	}
	return -1
}
