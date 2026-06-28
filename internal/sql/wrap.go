package sql

import "github.com/goccy/go-googlesql"

// Wrap returns the appropriate sql.Node wrapper for any raw googlesql.ASTNode.
// Returns nil if n is nil or its underlying pointer is nil.
//
// The switch mirrors the printer's visit() dispatch table, covering all node
// types handled by internal/printer.  Unknown node types fall through to
// genericNode so they remain usable via the base Node interface.
func Wrap(n googlesql.ASTNode) Node {
	if !defined(n) {
		return nil
	}
	switch m := n.(type) {
	// ── Query / Script ──────────────────────────────────────────────────────
	case *googlesql.ASTQueryStatement:
		return newASTQueryStatement(m)
	case *googlesql.ASTQuery:
		return newASTQuery(m)
	case *googlesql.ASTModelClause:
		return newASTModelClause(m)
	case *googlesql.ASTScript:
		return newASTScript(m)

	// ── SELECT ──────────────────────────────────────────────────────────────
	case *googlesql.ASTSelect:
		return newASTSelect(m)
	case *googlesql.ASTSelectList:
		return newASTSelectList(m)
	case *googlesql.ASTSelectColumn:
		return newASTSelectColumn(m)
	case *googlesql.ASTSelectAs:
		return newASTSelectAs(m)

	// ── Set operations ───────────────────────────────────────────────────────
	case *googlesql.ASTSetOperation:
		return newASTSetOperation(m)
	case *googlesql.ASTSetOperationMetadataList:
		return newASTSetOperationMetadataList(m)
	case *googlesql.ASTSetOperationMetadata:
		return newASTSetOperationMetadata(m)
	case *googlesql.ASTSetOperationType:
		return newASTSetOperationType(m)
	case *googlesql.ASTSetOperationAllOrDistinct:
		return newASTSetOperationAllOrDistinct(m)
	case *googlesql.ASTSetOperationColumnMatchMode:
		return newASTSetOperationColumnMatchMode(m)
	case *googlesql.ASTSetOperationColumnPropagationMode:
		return newASTSetOperationColumnPropagationMode(m)

	// ── WITH ─────────────────────────────────────────────────────────────────
	case *googlesql.ASTWithClause:
		return newASTWithClause(m)
	case *googlesql.ASTWithModifier:
		return newASTWithModifier(m)
	case *googlesql.ASTWithClauseEntry:
		return newASTWithClauseEntry(m)
	case *googlesql.ASTAliasedQuery:
		return newASTAliasedQuery(m)
	case *googlesql.ASTAliasedGroupRows:
		return newASTAliasedGroupRows(m)
	case *googlesql.ASTWithExpression:
		return newASTWithExpression(m)

	// ── FROM / JOIN ───────────────────────────────────────────────────────────
	case *googlesql.ASTFromClause:
		return newASTFromClause(m)
	case *googlesql.ASTJoin:
		return newASTJoin(m)
	case *googlesql.ASTParenthesizedJoin:
		return newASTParenthesizedJoin(m)
	case *googlesql.ASTTablePathExpression:
		return newASTTablePathExpression(m)
	case *googlesql.ASTTableSubquery:
		return newASTTableSubquery(m)
	case *googlesql.ASTTableClause:
		return newASTTableClause(m)
	case *googlesql.ASTTVF:
		return newASTTVF(m)

	case *googlesql.ASTUnnestExpression:
		return newASTUnnestExpression(m)
	case *googlesql.ASTUnnestExpressionWithOptAliasAndOffset:
		return newASTUnnestExpressionWithOptAliasAndOffset(m)

	// ── WHERE / GROUP BY / HAVING / QUALIFY ───────────────────────────────────
	case *googlesql.ASTWhereClause:
		return newASTWhereClause(m)
	case *googlesql.ASTGroupBy:
		return newASTGroupBy(m)
	case *googlesql.ASTGroupByAll:
		return newASTGroupByAll(m)
	case *googlesql.ASTGroupingItem:
		return newASTGroupingItem(m)
	case *googlesql.ASTHaving:
		return newASTHaving(m)
	case *googlesql.ASTQualify:
		return newASTQualify(m)

	// ── ORDER BY / LIMIT ──────────────────────────────────────────────────────
	case *googlesql.ASTOrderBy:
		return newASTOrderBy(m)
	case *googlesql.ASTOrderingExpression:
		return newASTOrderingExpression(m)
	case *googlesql.ASTNullOrder:
		return newASTNullOrder(m)
	case *googlesql.ASTLimitOffset:
		return newASTLimitOffset(m)
	case *googlesql.ASTLimit:
		return newASTLimit(m)

	// ── WINDOW ────────────────────────────────────────────────────────────────
	case *googlesql.ASTWindowClause:
		return newASTWindowClause(m)
	case *googlesql.ASTWindowDefinition:
		return newASTWindowDefinition(m)
	case *googlesql.ASTWindowSpecification:
		return newASTWindowSpecification(m)
	case *googlesql.ASTWindowFrame:
		return newASTWindowFrame(m)
	case *googlesql.ASTWindowFrameExpr:
		return newASTWindowFrameExpr(m)

	// ── PARTITION / CLUSTER ───────────────────────────────────────────────────
	case *googlesql.ASTPartitionBy:
		return newASTPartitionBy(m)
	case *googlesql.ASTClusterBy:
		return newASTClusterBy(m)

	// ── HINT ──────────────────────────────────────────────────────────────────
	case *googlesql.ASTHint:
		return newASTHint(m)
	case *googlesql.ASTHintEntry:
		return newASTHintEntry(m)
	case *googlesql.ASTHintedStatement:
		return newASTHintedStatement(m)

	// ── ON / USING ────────────────────────────────────────────────────────────
	case *googlesql.ASTOnClause:
		return newASTOnClause(m)
	case *googlesql.ASTUsingClause:
		return newASTUsingClause(m)

	// ── HAVING MODIFIER / CLAMPED ─────────────────────────────────────────────
	case *googlesql.ASTHavingModifier:
		return newASTHavingModifier(m)
	case *googlesql.ASTClampedBetweenModifier:
		return newASTClampedBetweenModifier(m)

	// ── Misc clauses ──────────────────────────────────────────────────────────
	case *googlesql.ASTCollate:
		return newASTCollate(m)
	case *googlesql.ASTForSystemTime:
		return newASTForSystemTime(m)
	case *googlesql.ASTFormatClause:
		return newASTFormatClause(m)
	case *googlesql.ASTWithOffset:
		return newASTWithOffset(m)
	case *googlesql.ASTWithWeight:
		return newASTWithWeight(m)
	case *googlesql.ASTAlias:
		return newASTAlias(m)

	// ── Identifiers / Paths ───────────────────────────────────────────────────
	case *googlesql.ASTIdentifier:
		return newASTIdentifier(m)
	case *googlesql.ASTIdentifierList:
		return newASTIdentifierList(m)
	case *googlesql.ASTPathExpression:
		return newASTPathExpression(m)
	case *googlesql.ASTPathExpressionList:
		return newASTPathExpressionList(m)

	// ── Star nodes ────────────────────────────────────────────────────────────
	case *googlesql.ASTStar:
		return newASTStar(m)
	case *googlesql.ASTStarWithModifiers:
		return newASTStarWithModifiers(m)
	case *googlesql.ASTStarModifiers:
		return newASTStarModifiers(m)
	case *googlesql.ASTStarExceptList:
		return newASTStarExceptList(m)
	case *googlesql.ASTStarReplaceItem:
		return newASTStarReplaceItem(m)
	case *googlesql.ASTExpressionWithOptAlias:
		return newASTExpressionWithOptAlias(m)

	// ── Expression nodes ──────────────────────────────────────────────────────
	case *googlesql.ASTAndExpr:
		return newASTAndExpr(m)
	case *googlesql.ASTOrExpr:
		return newASTOrExpr(m)
	case *googlesql.ASTBinaryExpression:
		return newASTBinaryExpression(m)
	case *googlesql.ASTBitwiseShiftExpression:
		return newASTBitwiseShiftExpression(m)
	case *googlesql.ASTUnaryExpression:
		return newASTUnaryExpression(m)
	case *googlesql.ASTBetweenExpression:
		return newASTBetweenExpression(m)
	case *googlesql.ASTCaseNoValueExpression:
		return newASTCaseNoValueExpression(m)
	case *googlesql.ASTCaseValueExpression:
		return newASTCaseValueExpression(m)
	case *googlesql.ASTCastExpression:
		return newASTCastExpression(m)
	case *googlesql.ASTInExpression:
		return newASTInExpression(m)
	case *googlesql.ASTInList:
		return newASTInList(m)
	case *googlesql.ASTExpressionSubquery:
		return newASTExpressionSubquery(m)
	case *googlesql.ASTExtractExpression:
		return newASTExtractExpression(m)
	case *googlesql.ASTIntervalExpr:
		return newASTIntervalExpr(m)
	case *googlesql.ASTConcatExpr:
		return newASTConcatExpr(m)
	case *googlesql.ASTArrayConstructor:
		return newASTArrayConstructor(m)
	case *googlesql.ASTArrayElement:
		return newASTArrayElement(m)
	case *googlesql.ASTDotIdentifier:
		return newASTDotIdentifier(m)
	case *googlesql.ASTDotStar:
		return newASTDotStar(m)
	case *googlesql.ASTDotStarWithModifiers:
		return newASTDotStarWithModifiers(m)
	case *googlesql.ASTDotGeneralizedField:
		return newASTDotGeneralizedField(m)
	case *googlesql.ASTParameterExpr:
		return newASTParameterExpr(m)
	case *googlesql.ASTSystemVariableExpr:
		return newASTSystemVariableExpr(m)
	case *googlesql.ASTStructConstructorArg:
		return newASTStructConstructorArg(m)
	case *googlesql.ASTStructConstructorWithKeyword:
		return newASTStructConstructorWithKeyword(m)
	case *googlesql.ASTStructConstructorWithParens:
		return newASTStructConstructorWithParens(m)
	case *googlesql.ASTNamedArgument:
		return newASTNamedArgument(m)
	case *googlesql.ASTLambda:
		return newASTLambda(m)

	// ── Function call ─────────────────────────────────────────────────────────
	case *googlesql.ASTFunctionCall:
		return newASTFunctionCall(m)
	case *googlesql.ASTAnalyticFunctionCall:
		return newASTAnalyticFunctionCall(m)

	// ── SAMPLE / PIVOT / UNPIVOT ──────────────────────────────────────────────
	case *googlesql.ASTSampleClause:
		return newASTSampleClause(m)
	case *googlesql.ASTSampleSize:
		return newASTSampleSize(m)
	case *googlesql.ASTSampleSuffix:
		return newASTSampleSuffix(m)
	case *googlesql.ASTRepeatableClause:
		return newASTRepeatableClause(m)
	case *googlesql.ASTPivotClause:
		return newASTPivotClause(m)
	case *googlesql.ASTPivotExpression:
		return newASTPivotExpression(m)
	case *googlesql.ASTPivotExpressionList:
		return newASTPivotExpressionList(m)
	case *googlesql.ASTPivotValue:
		return newASTPivotValue(m)
	case *googlesql.ASTPivotValueList:
		return newASTPivotValueList(m)
	case *googlesql.ASTUnpivotClause:
		return newASTUnpivotClause(m)
	case *googlesql.ASTUnpivotInItem:
		return newASTUnpivotInItem(m)
	case *googlesql.ASTUnpivotInItemList:
		return newASTUnpivotInItemList(m)
	case *googlesql.ASTUnpivotInItemLabel:
		return newASTUnpivotInItemLabel(m)

	// ── Literals ──────────────────────────────────────────────────────────────
	case *googlesql.ASTIntLiteral:
		return newASTIntLiteral(m)
	case *googlesql.ASTFloatLiteral:
		return newASTFloatLiteral(m)
	case *googlesql.ASTBooleanLiteral:
		return newASTBooleanLiteral(m)
	case *googlesql.ASTNullLiteral:
		return newASTNullLiteral(m)
	case *googlesql.ASTStringLiteral:
		return newASTStringLiteral(m)
	case *googlesql.ASTStringLiteralComponent:
		return newASTStringLiteralComponent(m)
	case *googlesql.ASTBytesLiteral:
		return newASTBytesLiteral(m)
	case *googlesql.ASTBytesLiteralComponent:
		return newASTBytesLiteralComponent(m)
	case *googlesql.ASTNumericLiteral:
		return newASTNumericLiteral(m)
	case *googlesql.ASTBigNumericLiteral:
		return newASTBigNumericLiteral(m)
	case *googlesql.ASTJSONLiteral:
		return newASTJSONLiteral(m)
	case *googlesql.ASTDateOrTimeLiteral:
		return newASTDateOrTimeLiteral(m)
	case *googlesql.ASTDefaultLiteral:
		return newASTDefaultLiteral(m)
	case *googlesql.ASTMaxLiteral:
		return newASTMaxLiteral(m)
	case *googlesql.ASTRangeLiteral:
		return newASTRangeLiteral(m)

	// ── OPTIONS ───────────────────────────────────────────────────────────────
	case *googlesql.ASTOptionsList:
		return newASTOptionsList(m)
	case *googlesql.ASTOptionsEntry:
		return newASTOptionsEntry(m)

	// ── Column list / Descriptor ───────────────────────────────────────────────
	case *googlesql.ASTColumnList:
		return newASTColumnList(m)
	case *googlesql.ASTDescriptor:
		return newASTDescriptor(m)
	case *googlesql.ASTDescriptorColumn:
		return newASTDescriptorColumn(m)
	case *googlesql.ASTDescriptorColumnList:
		return newASTDescriptorColumnList(m)
	case *googlesql.ASTTVFSchemaColumn:
		return newASTTVFSchemaColumn(m)
	case *googlesql.ASTTVFSchema:
		return newASTTVFSchema(m)

	// ── Statement list ────────────────────────────────────────────────────────
	case *googlesql.ASTStatementList:
		return newASTStatementList(m)

	// ── Rollup / Cube ─────────────────────────────────────────────────────────
	case *googlesql.ASTRollup:
		return newASTRollup(m)
	case *googlesql.ASTCube:
		return newASTCube(m)

	// ── Type nodes ────────────────────────────────────────────────────────────
	case *googlesql.ASTSimpleType:
		return newASTSimpleType(m)
	case *googlesql.ASTArrayType:
		return newASTArrayType(m)
	case *googlesql.ASTStructType:
		return newASTStructType(m)
	case *googlesql.ASTStructField:
		return newASTStructField(m)
	case *googlesql.ASTRangeType:
		return newASTRangeType(m)
	case *googlesql.ASTMapType:
		return newASTMapType(m)
	case *googlesql.ASTTemplatedParameterType:
		return newASTTemplatedParameterType(m)
	case *googlesql.ASTTypeParameterList:
		return newASTTypeParameterList(m)

	// ── DDL: Column/Schema nodes ──────────────────────────────────────────────
	case *googlesql.ASTColumnSchema:
		return newASTColumnSchema(m)
	case *googlesql.ASTSimpleColumnSchema:
		return newASTSimpleColumnSchema(m)
	case *googlesql.ASTArrayColumnSchema:
		return newASTArrayColumnSchema(m)
	case *googlesql.ASTStructColumnSchema:
		return newASTStructColumnSchema(m)
	case *googlesql.ASTStructColumnField:
		return newASTStructColumnField(m)
	case *googlesql.ASTColumnAttributeList:
		return newASTColumnAttributeList(m)
	case *googlesql.ASTNotNullColumnAttribute:
		return newASTNotNullColumnAttribute(m)
	case *googlesql.ASTPrimaryKeyColumnAttribute:
		return newASTPrimaryKeyColumnAttribute(m)
	case *googlesql.ASTForeignKeyColumnAttribute:
		return newASTForeignKeyColumnAttribute(m)
	case *googlesql.ASTHiddenColumnAttribute:
		return newASTHiddenColumnAttribute(m)
	case *googlesql.ASTColumnDefinition:

		return newASTColumnDefinition(m)
	case *googlesql.ASTColumnPosition:
		return newASTColumnPosition(m)
	case *googlesql.ASTGeneratedColumnInfo:
		return newASTGeneratedColumnInfo(m)
	case *googlesql.ASTTableElementList:
		return newASTTableElementList(m)

	// ── DDL: Constraints ──────────────────────────────────────────────────────
	case *googlesql.ASTTableConstraint:
		return newASTTableConstraint(m)
	case *googlesql.ASTPrimaryKey:
		return newASTPrimaryKey(m)
	case *googlesql.ASTPrimaryKeyElementList:
		return newASTPrimaryKeyElementList(m)
	case *googlesql.ASTPrimaryKeyElement:
		return newASTPrimaryKeyElement(m)
	case *googlesql.ASTForeignKey:
		return newASTForeignKey(m)
	case *googlesql.ASTForeignKeyReference:
		return newASTForeignKeyReference(m)

	// ── DDL: ALTER actions ────────────────────────────────────────────────────
	case *googlesql.ASTAlterActionList:
		return newASTAlterActionList(m)
	case *googlesql.ASTAddColumnAction:
		return newASTAddColumnAction(m)
	case *googlesql.ASTAddConstraintAction:
		return newASTAddConstraintAction(m)
	case *googlesql.ASTAlterColumnDropDefaultAction:
		return newASTAlterColumnDropDefaultAction(m)
	case *googlesql.ASTAlterColumnDropNotNullAction:
		return newASTAlterColumnDropNotNullAction(m)
	case *googlesql.ASTAlterColumnOptionsAction:
		return newASTAlterColumnOptionsAction(m)
	case *googlesql.ASTAlterColumnSetDefaultAction:
		return newASTAlterColumnSetDefaultAction(m)
	case *googlesql.ASTAlterColumnTypeAction:
		return newASTAlterColumnTypeAction(m)
	case *googlesql.ASTAlterConstraintEnforcementAction:
		return newASTAlterConstraintEnforcementAction(m)
	case *googlesql.ASTAlterConstraintSetOptionsAction:
		return newASTAlterConstraintSetOptionsAction(m)
	case *googlesql.ASTDropColumnAction:
		return newASTDropColumnAction(m)
	case *googlesql.ASTDropConstraintAction:
		return newASTDropConstraintAction(m)
	case *googlesql.ASTDropPrimaryKeyAction:
		return newASTDropPrimaryKeyAction(m)
	case *googlesql.ASTRenameColumnAction:
		return newASTRenameColumnAction(m)
	case *googlesql.ASTRenameToClause:
		return newASTRenameToClause(m)
	case *googlesql.ASTSetCollateClause:
		return newASTSetCollateClause(m)
	case *googlesql.ASTSetOptionsAction:
		return newASTSetOptionsAction(m)

	// ── DDL: ALTER statements ─────────────────────────────────────────────────
	case *googlesql.ASTAlterAllRowAccessPoliciesStatement:
		return newASTAlterAllRowAccessPoliciesStatement(m)
	case *googlesql.ASTAlterDatabaseStatement:
		return newASTAlterDatabaseStatement(m)
	case *googlesql.ASTAlterEntityStatement:
		return newASTAlterEntityStatement(m)
	case *googlesql.ASTAlterMaterializedViewStatement:
		return newASTAlterMaterializedViewStatement(m)
	case *googlesql.ASTAlterPrivilegeRestrictionStatement:
		return newASTAlterPrivilegeRestrictionStatement(m)
	case *googlesql.ASTAlterRowAccessPolicyStatement:
		return newASTAlterRowAccessPolicyStatement(m)
	case *googlesql.ASTAlterSchemaStatement:
		return newASTAlterSchemaStatement(m)
	case *googlesql.ASTAlterTableStatement:
		return newASTAlterTableStatement(m)
	case *googlesql.ASTAlterViewStatement:
		return newASTAlterViewStatement(m)

	// ── DDL: CREATE helpers ───────────────────────────────────────────────────
	case *googlesql.ASTCloneDataSource:
		return newASTCloneDataSource(m)
	case *googlesql.ASTCopyDataSource:
		return newASTCopyDataSource(m)
	case *googlesql.ASTWithConnectionClause:
		return newASTWithConnectionClause(m)
	case *googlesql.ASTConnectionClause:
		return newASTConnectionClause(m)
	case *googlesql.ASTWithPartitionColumnsClause:
		return newASTWithPartitionColumnsClause(m)
	case *googlesql.ASTFunctionDeclaration:
		return newASTFunctionDeclaration(m)
	case *googlesql.ASTFunctionParameters:
		return newASTFunctionParameters(m)
	case *googlesql.ASTFunctionParameter:
		return newASTFunctionParameter(m)
	case *googlesql.ASTSqlFunctionBody:
		return newASTSQLFunctionBody(m)
	case *googlesql.ASTGranteeList:
		return newASTGranteeList(m)
	case *googlesql.ASTGrantToClause:
		return newASTGrantToClause(m)
	case *googlesql.ASTFilterUsingClause:
		return newASTFilterUsingClause(m)
	case *googlesql.ASTColumnWithOptions:
		return newASTColumnWithOptions(m)
	case *googlesql.ASTColumnWithOptionsList:
		return newASTColumnWithOptionsList(m)

	// ── DDL: CREATE statements ────────────────────────────────────────────────
	case *googlesql.ASTCreateExternalTableStatement:
		return newASTCreateExternalTableStatement(m)
	case *googlesql.ASTCreateFunctionStatement:
		return newASTCreateFunctionStatement(m)
	case *googlesql.ASTCreateMaterializedViewStatement:
		return newASTCreateMaterializedViewStatement(m)
	case *googlesql.ASTCreateProcedureStatement:
		return newASTCreateProcedureStatement(m)
	case *googlesql.ASTCreateRowAccessPolicyStatement:
		return newASTCreateRowAccessPolicyStatement(m)
	case *googlesql.ASTCreateSchemaStatement:
		return newASTCreateSchemaStatement(m)
	case *googlesql.ASTCreateSnapshotTableStatement:
		return newASTCreateSnapshotTableStatement(m)
	case *googlesql.ASTCreateTableStatement:
		return newASTCreateTableStatement(m)
	case *googlesql.ASTCreateTableFunctionStatement:
		return newASTCreateTableFunctionStatement(m)
	case *googlesql.ASTCreateViewStatement:
		return newASTCreateViewStatement(m)

	// ── DDL: DROP statements ──────────────────────────────────────────────────
	case *googlesql.ASTDropAllRowAccessPoliciesStatement:
		return newASTDropAllRowAccessPoliciesStatement(m)
	case *googlesql.ASTDropEntityStatement:
		return newASTDropEntityStatement(m)
	case *googlesql.ASTDropFunctionStatement:
		return newASTDropFunctionStatement(m)
	case *googlesql.ASTDropMaterializedViewStatement:
		return newASTDropMaterializedViewStatement(m)
	case *googlesql.ASTDropPrivilegeRestrictionStatement:
		return newASTDropPrivilegeRestrictionStatement(m)
	case *googlesql.ASTDropRowAccessPolicyStatement:
		return newASTDropRowAccessPolicyStatement(m)
	case *googlesql.ASTDropSearchIndexStatement:
		return newASTDropSearchIndexStatement(m)
	case *googlesql.ASTDropSnapshotTableStatement:
		return newASTDropSnapshotTableStatement(m)
	case *googlesql.ASTDropTableFunctionStatement:
		return newASTDropTableFunctionStatement(m)
	case *googlesql.ASTDropStatement:
		return newASTDropStatement(m)

	// ── DML ───────────────────────────────────────────────────────────────────
	case *googlesql.ASTDeleteStatement:
		return newASTDeleteStatement(m)
	case *googlesql.ASTAssertRowsModified:
		return newASTAssertRowsModified(m)
	case *googlesql.ASTReturningClause:
		return newASTReturningClause(m)
	case *googlesql.ASTInsertValuesRow:
		return newASTInsertValuesRow(m)
	case *googlesql.ASTInsertValuesRowList:
		return newASTInsertValuesRowList(m)
	case *googlesql.ASTInsertStatement:
		return newASTInsertStatement(m)
	case *googlesql.ASTUpdateSetValue:
		return newASTUpdateSetValue(m)
	case *googlesql.ASTUpdateItem:
		return newASTUpdateItem(m)
	case *googlesql.ASTUpdateItemList:
		return newASTUpdateItemList(m)
	case *googlesql.ASTUpdateStatement:
		return newASTUpdateStatement(m)
	case *googlesql.ASTMergeAction:
		return newASTMergeAction(m)
	case *googlesql.ASTMergeWhenClause:
		return newASTMergeWhenClause(m)
	case *googlesql.ASTMergeWhenClauseList:
		return newASTMergeWhenClauseList(m)
	case *googlesql.ASTMergeStatement:
		return newASTMergeStatement(m)
	case *googlesql.ASTTruncateStatement:
		return newASTTruncateStatement(m)
	case *googlesql.ASTAssignmentFromStruct:
		return newASTAssignmentFromStruct(m)

	// ── Procedural ────────────────────────────────────────────────────────────
	case *googlesql.ASTTVFArgument:
		return newASTTVFArgument(m)
	case *googlesql.ASTExceptionHandler:
		return newASTExceptionHandler(m)
	case *googlesql.ASTExceptionHandlerList:
		return newASTExceptionHandlerList(m)
	case *googlesql.ASTBeginEndBlock:
		return newASTBeginEndBlock(m)
	case *googlesql.ASTBeginStatement:
		return newASTBeginStatement(m)
	case *googlesql.ASTRollbackStatement:
		return newASTRollbackStatement(m)
	case *googlesql.ASTCallStatement:
		return newASTCallStatement(m)
	case *googlesql.ASTCommitStatement:
		return newASTCommitStatement(m)
	case *googlesql.ASTExecuteIntoClause:
		return newASTExecuteIntoClause(m)
	case *googlesql.ASTExecuteUsingArgument:
		return newASTExecuteUsingArgument(m)
	case *googlesql.ASTExecuteUsingClause:
		return newASTExecuteUsingClause(m)
	case *googlesql.ASTExecuteImmediateStatement:
		return newASTExecuteImmediateStatement(m)
	case *googlesql.ASTElseifClause:
		return newASTElseifClause(m)
	case *googlesql.ASTElseifClauseList:
		return newASTElseifClauseList(m)
	case *googlesql.ASTIfStatement:
		return newASTIfStatement(m)
	case *googlesql.ASTParameterAssignment:
		return newASTParameterAssignment(m)
	case *googlesql.ASTReturnStatement:
		return newASTReturnStatement(m)
	case *googlesql.ASTSystemVariableAssignment:
		return newASTSystemVariableAssignment(m)
	case *googlesql.ASTSingleAssignment:
		return newASTSingleAssignment(m)
	case *googlesql.ASTVariableDeclaration:
		return newASTVariableDeclaration(m)

	// ── Debugging ───────────────────────────────────────────────────────────────

	case *googlesql.ASTAssertStatement:
		return newASTAssertStatement(m)

	// ── Fallback ──────────────────────────────────────────────────────────────
	default:
		if po, ok := n.(googlesql.ASTPipeOperatorNode); ok {
			return newGenericPipeOperatorNode(po)
		}
		if ca, ok := n.(googlesql.ASTColumnAttributeNode); ok {
			return newGenericColumnAttributeNode(ca)
		}
		if te, ok := n.(googlesql.ASTTableElementNode); ok {
			return newGenericTableElementNode(te)
		}
		if aa, ok := n.(googlesql.ASTAlterActionNode); ok {
			return newGenericAlterActionNode(aa)
		}
		if st, ok := n.(googlesql.ASTStatementNode); ok {
			return newGenericStatementNode(st)
		}
		return newGenericNode(n)
	}
}

// WalkNode visits every node in the sub-tree rooted at n in depth-first
// pre-order.  If cb returns a non-nil error, WalkNode stops and returns that
// error.  Returns nil on success.
func WalkNode(n Node, cb func(Node) error) error {
	if n == nil {
		return nil
	}
	if err := cb(n); err != nil {
		return err
	}
	for _, child := range n.Children() {
		if err := WalkNode(child, cb); err != nil {
			return err
		}
	}
	return nil
}
