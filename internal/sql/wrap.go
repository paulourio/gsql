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

	// ── nodes_action.go ────────────────────────────────────────────────────────

	case *googlesql.ASTAddColumnAction:
		return newAddColumnAction(m)
	case *googlesql.ASTAddConstraintAction:
		return newAddConstraintAction(m)
	case *googlesql.ASTAlterActionList:
		return newAlterActionList(m)
	case *googlesql.ASTAlterColumnDropDefaultAction:
		return newAlterColumnDropDefaultAction(m)
	case *googlesql.ASTAlterColumnDropNotNullAction:
		return newAlterColumnDropNotNullAction(m)
	case *googlesql.ASTAlterColumnOptionsAction:
		return newAlterColumnOptionsAction(m)
	case *googlesql.ASTAlterColumnSetDefaultAction:
		return newAlterColumnSetDefaultAction(m)
	case *googlesql.ASTAlterColumnTypeAction:
		return newAlterColumnTypeAction(m)
	case *googlesql.ASTAlterConstraintEnforcementAction:
		return newAlterConstraintEnforcementAction(m)
	case *googlesql.ASTAlterConstraintSetOptionsAction:
		return newAlterConstraintSetOptionsAction(m)
	case *googlesql.ASTDropColumnAction:
		return newDropColumnAction(m)
	case *googlesql.ASTDropConstraintAction:
		return newDropConstraintAction(m)
	case *googlesql.ASTDropPrimaryKeyAction:
		return newDropPrimaryKeyAction(m)
	case *googlesql.ASTMergeAction:
		return newMergeAction(m)
	case *googlesql.ASTRenameColumnAction:
		return newRenameColumnAction(m)
	case *googlesql.ASTRenameToClause:
		return newRenameToClause(m)
	case *googlesql.ASTSetCollateClause:
		return newSetCollateClause(m)
	case *googlesql.ASTSetOptionsAction:
		return newSetOptionsAction(m)

	// ── nodes_clause.go ────────────────────────────────────────────────────────

	case *googlesql.ASTClampedBetweenModifier:
		return newClampedBetweenModifier(m)
	case *googlesql.ASTClusterBy:
		return newClusterBy(m)
	case *googlesql.ASTCollate:
		return newCollate(m)
	case *googlesql.ASTColumnList:
		return newColumnList(m)
	case *googlesql.ASTConnectionClause:
		return newConnectionClause(m)
	case *googlesql.ASTDescriptorColumn:
		return newDescriptorColumn(m)
	case *googlesql.ASTDescriptorColumnList:
		return newDescriptorColumnList(m)
	case *googlesql.ASTElseifClause:
		return newElseifClause(m)
	case *googlesql.ASTExecuteIntoClause:
		return newExecuteIntoClause(m)
	case *googlesql.ASTExecuteUsingClause:
		return newExecuteUsingClause(m)
	case *googlesql.ASTFilterUsingClause:
		return newFilterUsingClause(m)
	case *googlesql.ASTForSystemTime:
		return newForSystemTime(m)
	case *googlesql.ASTFormatClause:
		return newFormatClause(m)
	case *googlesql.ASTFromClause:
		return newFromClause(m)
	case *googlesql.ASTGrantToClause:
		return newGrantToClause(m)
	case *googlesql.ASTGroupBy:
		return newGroupBy(m)
	case *googlesql.ASTGroupByAll:
		return newGroupByAll(m)
	case *googlesql.ASTGroupingItem:
		return newGroupingItem(m)
	case *googlesql.ASTHavingModifier:
		return newHavingModifier(m)
	case *googlesql.ASTHint:
		return newHint(m)
	case *googlesql.ASTHintEntry:
		return newHintEntry(m)
	case *googlesql.ASTLimit:
		return newLimit(m)
	case *googlesql.ASTLimitOffset:
		return newLimitOffset(m)
	case *googlesql.ASTMergeWhenClause:
		return newMergeWhenClause(m)
	case *googlesql.ASTModelClause:
		return newModelClause(m)
	case *googlesql.ASTNullOrder:
		return newNullOrder(m)
	case *googlesql.ASTOnClause:
		return newOnClause(m)
	case *googlesql.ASTOptionsEntry:
		return newOptionsEntry(m)
	case *googlesql.ASTOptionsList:
		return newOptionsList(m)
	case *googlesql.ASTOrderBy:
		return newOrderBy(m)
	case *googlesql.ASTOrderingExpression:
		return newOrderingExpression(m)
	case *googlesql.ASTPartitionBy:
		return newPartitionBy(m)
	case *googlesql.ASTPivotClause:
		return newPivotClause(m)
	case *googlesql.ASTPivotExpression:
		return newPivotExpression(m)
	case *googlesql.ASTPivotExpressionList:
		return newPivotExpressionList(m)
	case *googlesql.ASTPivotValue:
		return newPivotValue(m)
	case *googlesql.ASTPivotValueList:
		return newPivotValueList(m)
	case *googlesql.ASTRepeatableClause:
		return newRepeatableClause(m)
	case *googlesql.ASTReturningClause:
		return newReturningClause(m)
	case *googlesql.ASTSampleClause:
		return newSampleClause(m)
	case *googlesql.ASTSampleSize:
		return newSampleSize(m)
	case *googlesql.ASTSampleSuffix:
		return newSampleSuffix(m)
	case *googlesql.ASTUnpivotClause:
		return newUnpivotClause(m)
	case *googlesql.ASTUnpivotInItem:
		return newUnpivotInItem(m)
	case *googlesql.ASTUnpivotInItemLabel:
		return newUnpivotInItemLabel(m)
	case *googlesql.ASTUnpivotInItemList:
		return newUnpivotInItemList(m)
	case *googlesql.ASTUntilClause:
		return newUntilClause(m)
	case *googlesql.ASTUsingClause:
		return newUsingClause(m)
	case *googlesql.ASTWhereClause:
		return newWhereClause(m)
	case *googlesql.ASTWindowClause:
		return newWindowClause(m)
	case *googlesql.ASTWindowDefinition:
		return newWindowDefinition(m)
	case *googlesql.ASTWindowFrame:
		return newWindowFrame(m)
	case *googlesql.ASTWindowFrameExpr:
		return newWindowFrameExpr(m)
	case *googlesql.ASTWindowSpecification:
		return newWindowSpecification(m)
	case *googlesql.ASTWithClause:
		return newWithClause(m)
	case *googlesql.ASTWithConnectionClause:
		return newWithConnectionClause(m)
	case *googlesql.ASTWithOffset:
		return newWithOffset(m)
	case *googlesql.ASTWithPartitionColumnsClause:
		return newWithPartitionColumnsClause(m)
	case *googlesql.ASTWithWeight:
		return newWithWeight(m)

	// ── nodes_column_attribute.go ────────────────────────────────────────────────────────

	case *googlesql.ASTColumnAttributeList:
		return newColumnAttributeList(m)
	case *googlesql.ASTForeignKeyColumnAttribute:
		return newForeignKeyColumnAttribute(m)
	case *googlesql.ASTHiddenColumnAttribute:
		return newHiddenColumnAttribute(m)
	case *googlesql.ASTNotNullColumnAttribute:
		return newNotNullColumnAttribute(m)
	case *googlesql.ASTPrimaryKeyColumnAttribute:
		return newPrimaryKeyColumnAttribute(m)

	// ── nodes_core.go ────────────────────────────────────────────────────────

	case *googlesql.ASTAlias:
		return newAlias(m)
	case *googlesql.ASTAliasedGroupRows:
		return newAliasedGroupRows(m)
	case *googlesql.ASTAliasedQuery:
		return newAliasedQuery(m)
	case *googlesql.ASTAssertRowsModified:
		return newAssertRowsModified(m)
	case *googlesql.ASTCloneDataSource:
		return newCloneDataSource(m)
	case *googlesql.ASTColumnPosition:
		return newColumnPosition(m)
	case *googlesql.ASTColumnWithOptions:
		return newColumnWithOptions(m)
	case *googlesql.ASTColumnWithOptionsList:
		return newColumnWithOptionsList(m)
	case *googlesql.ASTCopyDataSource:
		return newCopyDataSource(m)
	case *googlesql.ASTCube:
		return newCube(m)
	case *googlesql.ASTElseifClauseList:
		return newElseifClauseList(m)
	case *googlesql.ASTExceptionHandler:
		return newExceptionHandler(m)
	case *googlesql.ASTExceptionHandlerList:
		return newExceptionHandlerList(m)
	case *googlesql.ASTExecuteUsingArgument:
		return newExecuteUsingArgument(m)
	case *googlesql.ASTExpressionWithOptAlias:
		return newExpressionWithOptAlias(m)
	case *googlesql.ASTFunctionDeclaration:
		return newFunctionDeclaration(m)
	case *googlesql.ASTFunctionParameter:
		return newFunctionParameter(m)
	case *googlesql.ASTFunctionParameters:
		return newFunctionParameters(m)
	case *googlesql.ASTGeneratedColumnInfo:
		return newGeneratedColumnInfo(m)
	case *googlesql.ASTGranteeList:
		return newGranteeList(m)
	case *googlesql.ASTHaving:
		return newHaving(m)
	case *googlesql.ASTIdentifierList:
		return newIdentifierList(m)
	case *googlesql.ASTInList:
		return newInList(m)
	case *googlesql.ASTInsertValuesRow:
		return newInsertValuesRow(m)
	case *googlesql.ASTInsertValuesRowList:
		return newInsertValuesRowList(m)
	case *googlesql.ASTLabel:
		return newLabel(m)
	case *googlesql.ASTMergeWhenClauseList:
		return newMergeWhenClauseList(m)
	case *googlesql.ASTPathExpressionList:
		return newPathExpressionList(m)
	case *googlesql.ASTQualify:
		return newQualify(m)
	case *googlesql.ASTRollup:
		return newRollup(m)
	case *googlesql.ASTSqlFunctionBody:
		return newSQLFunctionBody(m)
	case *googlesql.ASTScript:
		return newScript(m)
	case *googlesql.ASTSelectAs:
		return newSelectAs(m)
	case *googlesql.ASTSelectColumn:
		return newSelectColumn(m)
	case *googlesql.ASTSelectList:
		return newSelectList(m)
	case *googlesql.ASTSetOperationAllOrDistinct:
		return newSetOperationAllOrDistinct(m)
	case *googlesql.ASTSetOperationColumnMatchMode:
		return newSetOperationColumnMatchMode(m)
	case *googlesql.ASTSetOperationColumnPropagationMode:
		return newSetOperationColumnPropagationMode(m)
	case *googlesql.ASTSetOperationMetadata:
		return newSetOperationMetadata(m)
	case *googlesql.ASTSetOperationMetadataList:
		return newSetOperationMetadataList(m)
	case *googlesql.ASTStarExceptList:
		return newStarExceptList(m)
	case *googlesql.ASTStarModifiers:
		return newStarModifiers(m)
	case *googlesql.ASTStarReplaceItem:
		return newStarReplaceItem(m)
	case *googlesql.ASTStatementList:
		return newStatementList(m)
	case *googlesql.ASTStructConstructorArg:
		return newStructConstructorArg(m)
	case *googlesql.ASTTVFArgument:
		return newTVFArgument(m)
	case *googlesql.ASTTVFSchemaColumn:
		return newTVFSchemaColumn(m)
	case *googlesql.ASTUpdateItem:
		return newUpdateItem(m)
	case *googlesql.ASTUpdateItemList:
		return newUpdateItemList(m)
	case *googlesql.ASTUpdateSetValue:
		return newUpdateSetValue(m)
	case *googlesql.ASTWithClauseEntry:
		return newWithClauseEntry(m)
	case *googlesql.ASTWithModifier:
		return newWithModifier(m)

	// ── nodes_expression.go ────────────────────────────────────────────────────────

	case *googlesql.ASTAnalyticFunctionCall:
		return newAnalyticFunctionCall(m)
	case *googlesql.ASTAndExpr:
		return newAndExpr(m)
	case *googlesql.ASTArrayConstructor:
		return newArrayConstructor(m)
	case *googlesql.ASTArrayElement:
		return newArrayElement(m)
	case *googlesql.ASTBetweenExpression:
		return newBetweenExpression(m)
	case *googlesql.ASTBigNumericLiteral:
		return newBigNumericLiteral(m)
	case *googlesql.ASTBinaryExpression:
		return newBinaryExpression(m)
	case *googlesql.ASTBitwiseShiftExpression:
		return newBitwiseShiftExpression(m)
	case *googlesql.ASTBooleanLiteral:
		return newBooleanLiteral(m)
	case *googlesql.ASTBytesLiteral:
		return newBytesLiteral(m)
	case *googlesql.ASTBytesLiteralComponent:
		return newBytesLiteralComponent(m)
	case *googlesql.ASTCaseNoValueExpression:
		return newCaseNoValueExpression(m)
	case *googlesql.ASTCaseValueExpression:
		return newCaseValueExpression(m)
	case *googlesql.ASTCastExpression:
		return newCastExpression(m)
	case *googlesql.ASTConcatExpr:
		return newConcatExpr(m)
	case *googlesql.ASTDateOrTimeLiteral:
		return newDateOrTimeLiteral(m)
	case *googlesql.ASTDefaultLiteral:
		return newDefaultLiteral(m)
	case *googlesql.ASTDescriptor:
		return newDescriptor(m)
	case *googlesql.ASTDotGeneralizedField:
		return newDotGeneralizedField(m)
	case *googlesql.ASTDotIdentifier:
		return newDotIdentifier(m)
	case *googlesql.ASTDotStar:
		return newDotStar(m)
	case *googlesql.ASTDotStarWithModifiers:
		return newDotStarWithModifiers(m)
	case *googlesql.ASTExpressionSubquery:
		return newExpressionSubquery(m)
	case *googlesql.ASTExtractExpression:
		return newExtractExpression(m)
	case *googlesql.ASTFloatLiteral:
		return newFloatLiteral(m)
	case *googlesql.ASTFunctionCall:
		return newFunctionCall(m)
	case *googlesql.ASTIdentifier:
		return newIdentifier(m)
	case *googlesql.ASTInExpression:
		return newInExpression(m)
	case *googlesql.ASTIntLiteral:
		return newIntLiteral(m)
	case *googlesql.ASTIntervalExpr:
		return newIntervalExpr(m)
	case *googlesql.ASTJSONLiteral:
		return newJSONLiteral(m)
	case *googlesql.ASTLambda:
		return newLambda(m)
	case *googlesql.ASTMaxLiteral:
		return newMaxLiteral(m)
	case *googlesql.ASTNamedArgument:
		return newNamedArgument(m)
	case *googlesql.ASTNullLiteral:
		return newNullLiteral(m)
	case *googlesql.ASTNumericLiteral:
		return newNumericLiteral(m)
	case *googlesql.ASTOrExpr:
		return newOrExpr(m)
	case *googlesql.ASTParameterExpr:
		return newParameterExpr(m)
	case *googlesql.ASTPathExpression:
		return newPathExpression(m)
	case *googlesql.ASTRangeLiteral:
		return newRangeLiteral(m)
	case *googlesql.ASTStar:
		return newStar(m)
	case *googlesql.ASTStarWithModifiers:
		return newStarWithModifiers(m)
	case *googlesql.ASTStringLiteral:
		return newStringLiteral(m)
	case *googlesql.ASTStringLiteralComponent:
		return newStringLiteralComponent(m)
	case *googlesql.ASTStructConstructorWithKeyword:
		return newStructConstructorWithKeyword(m)
	case *googlesql.ASTStructConstructorWithParens:
		return newStructConstructorWithParens(m)
	case *googlesql.ASTSystemVariableExpr:
		return newSystemVariableExpr(m)
	case *googlesql.ASTUnaryExpression:
		return newUnaryExpression(m)
	case *googlesql.ASTWithExpression:
		return newWithExpression(m)

	// ── nodes_query_expression.go ────────────────────────────────────────────────────────

	case *googlesql.ASTQuery:
		return newQuery(m)
	case *googlesql.ASTSelect:
		return newSelect(m)
	case *googlesql.ASTSetOperation:
		return newSetOperation(m)

	// ── nodes_statement.go ────────────────────────────────────────────────────────

	case *googlesql.ASTAlterAllRowAccessPoliciesStatement:
		return newAlterAllRowAccessPoliciesStatement(m)
	case *googlesql.ASTAlterDatabaseStatement:
		return newAlterDatabaseStatement(m)
	case *googlesql.ASTAlterEntityStatement:
		return newAlterEntityStatement(m)
	case *googlesql.ASTAlterMaterializedViewStatement:
		return newAlterMaterializedViewStatement(m)
	case *googlesql.ASTAlterPrivilegeRestrictionStatement:
		return newAlterPrivilegeRestrictionStatement(m)
	case *googlesql.ASTAlterRowAccessPolicyStatement:
		return newAlterRowAccessPolicyStatement(m)
	case *googlesql.ASTAlterSchemaStatement:
		return newAlterSchemaStatement(m)
	case *googlesql.ASTAlterTableStatement:
		return newAlterTableStatement(m)
	case *googlesql.ASTAlterViewStatement:
		return newAlterViewStatement(m)
	case *googlesql.ASTAssertStatement:
		return newAssertStatement(m)
	case *googlesql.ASTAssignmentFromStruct:
		return newAssignmentFromStruct(m)
	case *googlesql.ASTBeginEndBlock:
		return newBeginEndBlock(m)
	case *googlesql.ASTBeginStatement:
		return newBeginStatement(m)
	case *googlesql.ASTBreakStatement:
		return newBreakStatement(m)
	case *googlesql.ASTCallStatement:
		return newCallStatement(m)
	case *googlesql.ASTCommitStatement:
		return newCommitStatement(m)
	case *googlesql.ASTContinueStatement:
		return newContinueStatement(m)
	case *googlesql.ASTCreateExternalTableStatement:
		return newCreateExternalTableStatement(m)
	case *googlesql.ASTCreateFunctionStatement:
		return newCreateFunctionStatement(m)
	case *googlesql.ASTCreateMaterializedViewStatement:
		return newCreateMaterializedViewStatement(m)
	case *googlesql.ASTCreateProcedureStatement:
		return newCreateProcedureStatement(m)
	case *googlesql.ASTCreateRowAccessPolicyStatement:
		return newCreateRowAccessPolicyStatement(m)
	case *googlesql.ASTCreateSchemaStatement:
		return newCreateSchemaStatement(m)
	case *googlesql.ASTCreateSnapshotTableStatement:
		return newCreateSnapshotTableStatement(m)
	case *googlesql.ASTCreateTableFunctionStatement:
		return newCreateTableFunctionStatement(m)
	case *googlesql.ASTCreateTableStatement:
		return newCreateTableStatement(m)
	case *googlesql.ASTCreateViewStatement:
		return newCreateViewStatement(m)
	case *googlesql.ASTDeleteStatement:
		return newDeleteStatement(m)
	case *googlesql.ASTDropAllRowAccessPoliciesStatement:
		return newDropAllRowAccessPoliciesStatement(m)
	case *googlesql.ASTDropEntityStatement:
		return newDropEntityStatement(m)
	case *googlesql.ASTDropFunctionStatement:
		return newDropFunctionStatement(m)
	case *googlesql.ASTDropMaterializedViewStatement:
		return newDropMaterializedViewStatement(m)
	case *googlesql.ASTDropPrivilegeRestrictionStatement:
		return newDropPrivilegeRestrictionStatement(m)
	case *googlesql.ASTDropRowAccessPolicyStatement:
		return newDropRowAccessPolicyStatement(m)
	case *googlesql.ASTDropSearchIndexStatement:
		return newDropSearchIndexStatement(m)
	case *googlesql.ASTDropSnapshotTableStatement:
		return newDropSnapshotTableStatement(m)
	case *googlesql.ASTDropStatement:
		return newDropStatement(m)
	case *googlesql.ASTDropTableFunctionStatement:
		return newDropTableFunctionStatement(m)
	case *googlesql.ASTExecuteImmediateStatement:
		return newExecuteImmediateStatement(m)
	case *googlesql.ASTForInStatement:
		return newForInStatement(m)
	case *googlesql.ASTHintedStatement:
		return newHintedStatement(m)
	case *googlesql.ASTIfStatement:
		return newIfStatement(m)
	case *googlesql.ASTInsertStatement:
		return newInsertStatement(m)
	case *googlesql.ASTMergeStatement:
		return newMergeStatement(m)
	case *googlesql.ASTParameterAssignment:
		return newParameterAssignment(m)
	case *googlesql.ASTQueryStatement:
		return newQueryStatement(m)
	case *googlesql.ASTRaiseStatement:
		return newRaiseStatement(m)
	case *googlesql.ASTReturnStatement:
		return newReturnStatement(m)
	case *googlesql.ASTRepeatStatement:
		return newRepeatStatement(m)
	case *googlesql.ASTRollbackStatement:
		return newRollbackStatement(m)
	case *googlesql.ASTSingleAssignment:
		return newSingleAssignment(m)
	case *googlesql.ASTSystemVariableAssignment:
		return newSystemVariableAssignment(m)
	case *googlesql.ASTTruncateStatement:
		return newTruncateStatement(m)
	case *googlesql.ASTUpdateStatement:
		return newUpdateStatement(m)
	case *googlesql.ASTVariableDeclaration:
		return newVariableDeclaration(m)
	case *googlesql.ASTWhileStatement:
		return newWhileStatement(m)

	// ── nodes_table_element.go ────────────────────────────────────────────────────────

	case *googlesql.ASTArrayColumnSchema:
		return newArrayColumnSchema(m)
	case *googlesql.ASTColumnDefinition:
		return newColumnDefinition(m)
	case *googlesql.ASTColumnSchema:
		return newColumnSchema(m)
	case *googlesql.ASTForeignKey:
		return newForeignKey(m)
	case *googlesql.ASTForeignKeyReference:
		return newForeignKeyReference(m)
	case *googlesql.ASTPrimaryKey:
		return newPrimaryKey(m)
	case *googlesql.ASTPrimaryKeyElement:
		return newPrimaryKeyElement(m)
	case *googlesql.ASTPrimaryKeyElementList:
		return newPrimaryKeyElementList(m)
	case *googlesql.ASTSimpleColumnSchema:
		return newSimpleColumnSchema(m)
	case *googlesql.ASTStructColumnField:
		return newStructColumnField(m)
	case *googlesql.ASTStructColumnSchema:
		return newStructColumnSchema(m)
	case *googlesql.ASTTVFSchema:
		return newTVFSchema(m)
	case *googlesql.ASTTableConstraint:
		return newTableConstraint(m)
	case *googlesql.ASTTableElementList:
		return newTableElementList(m)

	// ── nodes_table_expression.go ────────────────────────────────────────────────────────

	case *googlesql.ASTJoin:
		return newJoin(m)
	case *googlesql.ASTParenthesizedJoin:
		return newParenthesizedJoin(m)
	case *googlesql.ASTTVF:
		return newTVF(m)
	case *googlesql.ASTTableClause:
		return newTableClause(m)
	case *googlesql.ASTTablePathExpression:
		return newTablePathExpression(m)
	case *googlesql.ASTTableSubquery:
		return newTableSubquery(m)
	case *googlesql.ASTUnnestExpression:
		return newUnnestExpression(m)
	case *googlesql.ASTUnnestExpressionWithOptAliasAndOffset:
		return newUnnestExpressionWithOptAliasAndOffset(m)

	// ── nodes_type.go ────────────────────────────────────────────────────────

	case *googlesql.ASTArrayType:
		return newArrayType(m)
	case *googlesql.ASTMapType:
		return newMapType(m)
	case *googlesql.ASTRangeType:
		return newRangeType(m)
	case *googlesql.ASTSetOperationType:
		return newSetOperationType(m)
	case *googlesql.ASTSimpleType:
		return newSimpleType(m)
	case *googlesql.ASTStructField:
		return newStructField(m)
	case *googlesql.ASTStructType:
		return newStructType(m)
	case *googlesql.ASTTemplatedParameterType:
		return newTemplatedParameterType(m)
	case *googlesql.ASTTypeParameterList:
		return newTypeParameterList(m)

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
