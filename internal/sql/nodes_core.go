package sql

import (
	"github.com/goccy/go-googlesql"
)

// Alias wraps *googlesql.ASTAlias.
type Alias struct{ baseNode[*googlesql.ASTAlias] }

func newAlias(r *googlesql.ASTAlias) *Alias {
	if r == nil {
		return nil
	}
	return &Alias{baseNode[*googlesql.ASTAlias]{raw: r}}
}

func (n *Alias) Identifier() *Identifier { return newIdentifier(must(n.raw.Identifier())) }

// AliasedGroupRows wraps *googlesql.ASTAliasedGroupRows.
type AliasedGroupRows struct {
	baseNode[*googlesql.ASTAliasedGroupRows]
}

func newAliasedGroupRows(r *googlesql.ASTAliasedGroupRows) *AliasedGroupRows {
	if r == nil {
		return nil
	}
	return &AliasedGroupRows{baseNode[*googlesql.ASTAliasedGroupRows]{raw: r}}
}

func (n *AliasedGroupRows) Alias() *Identifier {
	return newIdentifier(must(n.raw.Alias()))
}

// AliasedQuery wraps *googlesql.ASTAliasedQuery.
type AliasedQuery struct {
	baseNode[*googlesql.ASTAliasedQuery]
}

func newAliasedQuery(r *googlesql.ASTAliasedQuery) *AliasedQuery {
	if r == nil {
		return nil
	}
	return &AliasedQuery{baseNode[*googlesql.ASTAliasedQuery]{raw: r}}
}

func (n *AliasedQuery) Alias() *Identifier { return newIdentifier(must(n.raw.Alias())) }

func (n *AliasedQuery) Query() *Query { return newQuery(must(n.raw.Query())) }

type AssertRowsModified struct {
	baseNode[*googlesql.ASTAssertRowsModified]
}

func newAssertRowsModified(r *googlesql.ASTAssertRowsModified) *AssertRowsModified {
	if r == nil {
		return nil
	}
	return &AssertRowsModified{baseNode[*googlesql.ASTAssertRowsModified]{raw: r}}
}

func (n *AssertRowsModified) NumRows() ExpressionNode {
	return wrapExpr(must(n.raw.NumRows()))
}

// CloneDataSource wraps *googlesql.ASTCloneDataSource.
type CloneDataSource struct {
	baseNode[*googlesql.ASTCloneDataSource]
}

func newCloneDataSource(r *googlesql.ASTCloneDataSource) *CloneDataSource {
	if r == nil {
		return nil
	}
	return &CloneDataSource{baseNode[*googlesql.ASTCloneDataSource]{raw: r}}
}

func (n *CloneDataSource) PathExpr() *PathExpression {
	return newPathExpression(must(n.raw.PathExpr()))
}

func (n *CloneDataSource) ForSystemTime() *ForSystemTime {
	return newForSystemTime(must(n.raw.ForSystemTime()))
}

func (n *CloneDataSource) WhereClause() *WhereClause {
	return newWhereClause(must(n.raw.WhereClause()))
}

// ColumnPosition wraps *googlesql.ASTColumnPosition.
type ColumnPosition struct {
	baseNode[*googlesql.ASTColumnPosition]
}

func newColumnPosition(r *googlesql.ASTColumnPosition) *ColumnPosition {
	if r == nil {
		return nil
	}
	return &ColumnPosition{baseNode[*googlesql.ASTColumnPosition]{raw: r}}
}

func (n *ColumnPosition) Identifier() *Identifier {
	return newIdentifier(must(n.raw.Identifier()))
}

func (n *ColumnPosition) RelativePosition() RelativePosition { return must(n.raw.Type()) }

// ColumnWithOptions wraps *googlesql.ASTColumnWithOptions.
type ColumnWithOptions struct {
	baseNode[*googlesql.ASTColumnWithOptions]
}

func newColumnWithOptions(r *googlesql.ASTColumnWithOptions) *ColumnWithOptions {
	if r == nil {
		return nil
	}
	return &ColumnWithOptions{baseNode[*googlesql.ASTColumnWithOptions]{raw: r}}
}

func (n *ColumnWithOptions) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

func (n *ColumnWithOptions) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// ColumnWithOptionsList wraps *googlesql.ASTColumnWithOptionsList.
type ColumnWithOptionsList struct {
	baseNode[*googlesql.ASTColumnWithOptionsList]
}

func newColumnWithOptionsList(r *googlesql.ASTColumnWithOptionsList) *ColumnWithOptionsList {
	if r == nil {
		return nil
	}
	return &ColumnWithOptionsList{baseNode[*googlesql.ASTColumnWithOptionsList]{raw: r}}
}

// Entries returns all *ColumnWithOptions children.
func (n *ColumnWithOptionsList) Entries() []*ColumnWithOptions {
	var result []*ColumnWithOptions
	for _, c := range n.Children() {
		if cw, ok := c.(*ColumnWithOptions); ok {
			result = append(result, cw)
		}
	}
	return result
}

// CopyDataSource wraps *googlesql.ASTCopyDataSource.
type CopyDataSource struct {
	baseNode[*googlesql.ASTCopyDataSource]
}

func newCopyDataSource(r *googlesql.ASTCopyDataSource) *CopyDataSource {
	if r == nil {
		return nil
	}
	return &CopyDataSource{baseNode[*googlesql.ASTCopyDataSource]{raw: r}}
}

func (n *CopyDataSource) PathExpr() *PathExpression {
	return newPathExpression(must(n.raw.PathExpr()))
}

func (n *CopyDataSource) ForSystemTime() *ForSystemTime {
	return newForSystemTime(must(n.raw.ForSystemTime()))
}

func (n *CopyDataSource) WhereClause() *WhereClause {
	return newWhereClause(must(n.raw.WhereClause()))
}

// Cube wraps *googlesql.ASTCube.
type Cube struct{ baseNode[*googlesql.ASTCube] }

func newCube(r *googlesql.ASTCube) *Cube {
	if r == nil {
		return nil
	}
	return &Cube{baseNode[*googlesql.ASTCube]{raw: r}}
}

// Expressions returns cube expressions.
func (n *Cube) Expressions() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		e := must(n.raw.Expressions(int32(i)))
		if e == nil {
			break
		}
		result = append(result, wrapExpr(e))
	}
	return result
}

// ElseifClauseList wraps *googlesql.ASTElseifClauseList (via raw children).
type ElseifClauseList struct {
	baseNode[*googlesql.ASTElseifClauseList]
}

func newElseifClauseList(r *googlesql.ASTElseifClauseList) *ElseifClauseList {
	if r == nil {
		return nil
	}
	return &ElseifClauseList{baseNode[*googlesql.ASTElseifClauseList]{raw: r}}
}

// Clauses returns all *ElseifClause children.
func (n *ElseifClauseList) Clauses() []*ElseifClause {
	var result []*ElseifClause
	for _, c := range n.Children() {
		if e, ok := c.(*ElseifClause); ok {
			result = append(result, e)
		}
	}
	return result
}

// ExceptionHandler wraps *googlesql.ASTExceptionHandler.
type ExceptionHandler struct {
	baseNode[*googlesql.ASTExceptionHandler]
}

func newExceptionHandler(r *googlesql.ASTExceptionHandler) *ExceptionHandler {
	if r == nil {
		return nil
	}
	return &ExceptionHandler{baseNode[*googlesql.ASTExceptionHandler]{raw: r}}
}

func (n *ExceptionHandler) StatementList() *StatementList {
	return newStatementList(must(n.raw.StatementList()))
}

// ExceptionHandlerList wraps *googlesql.ASTExceptionHandlerList.
type ExceptionHandlerList struct {
	baseNode[*googlesql.ASTExceptionHandlerList]
}

func newExceptionHandlerList(r *googlesql.ASTExceptionHandlerList) *ExceptionHandlerList {
	if r == nil {
		return nil
	}
	return &ExceptionHandlerList{baseNode[*googlesql.ASTExceptionHandlerList]{raw: r}}
}

// Handlers returns all *ExceptionHandler children.
func (n *ExceptionHandlerList) Handlers() []*ExceptionHandler {
	var result []*ExceptionHandler
	for _, c := range n.Children() {
		if h, ok := c.(*ExceptionHandler); ok {
			result = append(result, h)
		}
	}
	return result
}

// ExecuteUsingArgument wraps *googlesql.ASTExecuteUsingArgument.
type ExecuteUsingArgument struct {
	baseNode[*googlesql.ASTExecuteUsingArgument]
}

func newExecuteUsingArgument(r *googlesql.ASTExecuteUsingArgument) *ExecuteUsingArgument {
	if r == nil {
		return nil
	}
	return &ExecuteUsingArgument{baseNode[*googlesql.ASTExecuteUsingArgument]{raw: r}}
}

func (n *ExecuteUsingArgument) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *ExecuteUsingArgument) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

// ExpressionWithOptAlias wraps *googlesql.ASTExpressionWithOptAlias.
type ExpressionWithOptAlias struct {
	baseNode[*googlesql.ASTExpressionWithOptAlias]
}

func newExpressionWithOptAlias(r *googlesql.ASTExpressionWithOptAlias) *ExpressionWithOptAlias {
	if r == nil {
		return nil
	}
	return &ExpressionWithOptAlias{baseNode[*googlesql.ASTExpressionWithOptAlias]{raw: r}}
}

func (n *ExpressionWithOptAlias) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *ExpressionWithOptAlias) OptionalAlias() *Alias {
	return newAlias(must(n.raw.OptionalAlias()))
}

// FunctionDeclaration wraps *googlesql.ASTFunctionDeclaration.
type FunctionDeclaration struct {
	baseNode[*googlesql.ASTFunctionDeclaration]
}

func newFunctionDeclaration(r *googlesql.ASTFunctionDeclaration) *FunctionDeclaration {
	if r == nil {
		return nil
	}
	return &FunctionDeclaration{baseNode[*googlesql.ASTFunctionDeclaration]{raw: r}}
}

func (n *FunctionDeclaration) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *FunctionDeclaration) Parameters() *FunctionParameters {
	return newFunctionParameters(must(n.raw.Parameters()))
}

// FunctionParameter wraps *googlesql.ASTFunctionParameter.
type FunctionParameter struct {
	baseNode[*googlesql.ASTFunctionParameter]
}

func newFunctionParameter(r *googlesql.ASTFunctionParameter) *FunctionParameter {
	if r == nil {
		return nil
	}
	return &FunctionParameter{baseNode[*googlesql.ASTFunctionParameter]{raw: r}}
}

func (n *FunctionParameter) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

func (n *FunctionParameter) ProcedureParameterMode() ParameterMode {
	return must(n.raw.ProcedureParameterMode())
}

func (n *FunctionParameter) IsNotAggregate() bool { return must(n.raw.IsNotAggregate()) }

func (n *FunctionParameter) Type() TypeNode { return wrapType(must(n.raw.Type())) }

func (n *FunctionParameter) TemplatedParameterType() *TemplatedParameterType {
	return newTemplatedParameterType(must(n.raw.TemplatedParameterType()))
}

func (n *FunctionParameter) TvfSchema() Node { return Wrap(must(n.raw.TvfSchema())) }

func (n *FunctionParameter) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *FunctionParameter) DefaultValue() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultValue()))
}

// FunctionParameters wraps *googlesql.ASTFunctionParameters.
type FunctionParameters struct {
	baseNode[*googlesql.ASTFunctionParameters]
}

func newFunctionParameters(r *googlesql.ASTFunctionParameters) *FunctionParameters {
	if r == nil {
		return nil
	}
	return &FunctionParameters{baseNode[*googlesql.ASTFunctionParameters]{raw: r}}
}

// Entries returns all *FunctionParameter children.
func (n *FunctionParameters) Entries() []*FunctionParameter {
	var result []*FunctionParameter
	for _, c := range n.Children() {
		if fp, ok := c.(*FunctionParameter); ok {
			result = append(result, fp)
		}
	}
	return result
}

// GeneratedColumnInfo wraps *googlesql.ASTGeneratedColumnInfo.
type GeneratedColumnInfo struct {
	baseNode[*googlesql.ASTGeneratedColumnInfo]
}

func newGeneratedColumnInfo(r *googlesql.ASTGeneratedColumnInfo) *GeneratedColumnInfo {
	if r == nil {
		return nil
	}
	return &GeneratedColumnInfo{baseNode[*googlesql.ASTGeneratedColumnInfo]{raw: r}}
}

func (n *GeneratedColumnInfo) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *GeneratedColumnInfo) GeneratedMode() GeneratedMode {
	return must(n.raw.GeneratedMode())
}

func (n *GeneratedColumnInfo) StoredMode() StoredMode { return must(n.raw.StoredMode()) }

// GranteeList wraps *googlesql.ASTGranteeList.
type GranteeList struct {
	baseNode[*googlesql.ASTGranteeList]
}

func newGranteeList(r *googlesql.ASTGranteeList) *GranteeList {
	if r == nil {
		return nil
	}
	return &GranteeList{baseNode[*googlesql.ASTGranteeList]{raw: r}}
}

// Grantees returns all grantees as []ExpressionNode.
func (n *GranteeList) Grantees() []ExpressionNode {
	var result []ExpressionNode
	for _, c := range n.Children() {
		if e, ok := c.(ExpressionNode); ok {
			result = append(result, e)
		}
	}
	return result
}

// Having wraps *googlesql.ASTHaving.
type Having struct{ baseNode[*googlesql.ASTHaving] }

func newHaving(r *googlesql.ASTHaving) *Having {
	if r == nil {
		return nil
	}
	return &Having{baseNode[*googlesql.ASTHaving]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *Having) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// IdentifierList wraps *googlesql.ASTIdentifierList.
type IdentifierList struct {
	baseNode[*googlesql.ASTIdentifierList]
}

func newIdentifierList(r *googlesql.ASTIdentifierList) *IdentifierList {
	if r == nil {
		return nil
	}
	return &IdentifierList{baseNode[*googlesql.ASTIdentifierList]{raw: r}}
}

// IdentifierList: raw returns *Identifier (concrete) → []*Identifier.
func (n *IdentifierList) IdentifierList() []*Identifier {
	count := n.NumChildren()
	result := make([]*Identifier, 0, count)
	for i := range count {
		id := must(n.raw.IdentifierList(int32(i)))
		if id == nil {
			break
		}
		result = append(result, newIdentifier(id))
	}
	return result
}

// InList wraps *googlesql.ASTInList.
type InList struct{ baseNode[*googlesql.ASTInList] }

func newInList(r *googlesql.ASTInList) *InList {
	if r == nil {
		return nil
	}
	return &InList{baseNode[*googlesql.ASTInList]{raw: r}}
}

// List: raw List(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *InList) List() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		e := must(n.raw.List(int32(i)))
		if !defined(e) {
			break
		}
		result = append(result, wrapExpr(e))
	}
	return result
}

// InsertValuesRow wraps *googlesql.ASTInsertValuesRow.
type InsertValuesRow struct {
	baseNode[*googlesql.ASTInsertValuesRow]
}

func newInsertValuesRow(r *googlesql.ASTInsertValuesRow) *InsertValuesRow {
	if r == nil {
		return nil
	}
	return &InsertValuesRow{baseNode[*googlesql.ASTInsertValuesRow]{raw: r}}
}

// Values returns all expression children.
func (n *InsertValuesRow) Values() []ExpressionNode {
	var result []ExpressionNode
	for _, c := range n.Children() {
		if e, ok := c.(ExpressionNode); ok {
			result = append(result, e)
		}
	}
	return result
}

// InsertValuesRowList wraps *googlesql.ASTInsertValuesRowList.
type InsertValuesRowList struct {
	baseNode[*googlesql.ASTInsertValuesRowList]
}

func newInsertValuesRowList(r *googlesql.ASTInsertValuesRowList) *InsertValuesRowList {
	if r == nil {
		return nil
	}
	return &InsertValuesRowList{baseNode[*googlesql.ASTInsertValuesRowList]{raw: r}}
}

// Rows returns []*InsertValuesRow children.
func (n *InsertValuesRowList) Rows() []*InsertValuesRow {
	var result []*InsertValuesRow
	for _, c := range n.Children() {
		if r, ok := c.(*InsertValuesRow); ok {
			result = append(result, r)
		}
	}
	return result
}

// Label wraps *googlesql.ASTLabel.
type Label struct {
	baseNode[*googlesql.ASTLabel]
}

func newLabel(r *googlesql.ASTLabel) *Label {
	if r == nil {
		return nil
	}
	return &Label{baseNode[*googlesql.ASTLabel]{raw: r}}
}

func (n *Label) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

// LockMode wraps *googlesql.ASTLockMode.
type LockMode struct {
	baseNode[*googlesql.ASTLockMode]
}

func newLockMode(r *googlesql.ASTLockMode) *LockMode {
	if r == nil {
		return nil
	}
	return &LockMode{baseNode[*googlesql.ASTLockMode]{raw: r}}
}

// MergeWhenClauseList wraps *googlesql.ASTMergeWhenClauseList.
type MergeWhenClauseList struct {
	baseNode[*googlesql.ASTMergeWhenClauseList]
}

func newMergeWhenClauseList(r *googlesql.ASTMergeWhenClauseList) *MergeWhenClauseList {
	if r == nil {
		return nil
	}
	return &MergeWhenClauseList{baseNode[*googlesql.ASTMergeWhenClauseList]{raw: r}}
}

// Clauses returns []*MergeWhenClause children.
func (n *MergeWhenClauseList) Clauses() []*MergeWhenClause {
	var result []*MergeWhenClause
	for _, c := range n.Children() {
		if cl, ok := c.(*MergeWhenClause); ok {
			result = append(result, cl)
		}
	}
	return result
}

// PathExpressionList wraps *googlesql.ASTPathExpressionList.
type PathExpressionList struct {
	baseNode[*googlesql.ASTPathExpressionList]
}

func newPathExpressionList(r *googlesql.ASTPathExpressionList) *PathExpressionList {
	if r == nil {
		return nil
	}
	return &PathExpressionList{baseNode[*googlesql.ASTPathExpressionList]{raw: r}}
}

// PathExpressionList: raw returns *PathExpression (concrete) → []*PathExpression.
func (n *PathExpressionList) PathExpressionList() []*PathExpression {
	count := n.NumChildren()
	result := make([]*PathExpression, 0, count)
	for i := range count {
		pe := must(n.raw.PathExpressionList(int32(i)))
		if pe == nil {
			break
		}
		result = append(result, newPathExpression(pe))
	}
	return result
}

// Qualify wraps *googlesql.ASTQualify.
type Qualify struct {
	baseNode[*googlesql.ASTQualify]
}

func newQualify(r *googlesql.ASTQualify) *Qualify {
	if r == nil {
		return nil
	}
	return &Qualify{baseNode[*googlesql.ASTQualify]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *Qualify) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// Rollup wraps *googlesql.ASTRollup.
type Rollup struct{ baseNode[*googlesql.ASTRollup] }

func newRollup(r *googlesql.ASTRollup) *Rollup {
	if r == nil {
		return nil
	}
	return &Rollup{baseNode[*googlesql.ASTRollup]{raw: r}}
}

// Expressions returns rollup expressions.
func (n *Rollup) Expressions() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		e := must(n.raw.Expressions(int32(i)))
		if e == nil {
			break
		}
		result = append(result, wrapExpr(e))
	}
	return result
}

// SQLFunctionBody wraps *googlesql.ASTSqlFunctionBody.
//
//nolint:revive
type SQLFunctionBody struct {
	baseNode[*googlesql.ASTSqlFunctionBody]
}

func newSQLFunctionBody(r *googlesql.ASTSqlFunctionBody) *SQLFunctionBody {
	if r == nil {
		return nil
	}
	return &SQLFunctionBody{baseNode[*googlesql.ASTSqlFunctionBody]{raw: r}}
}

func (n *SQLFunctionBody) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// Script wraps *googlesql.ASTScript.
type Script struct{ baseNode[*googlesql.ASTScript] }

func newScript(r *googlesql.ASTScript) *Script {
	if r == nil {
		return nil
	}
	return &Script{baseNode[*googlesql.ASTScript]{raw: r}}
}

func (n *Script) StatementList() *StatementList {
	return newStatementList(must(n.raw.StatementListNode()))
}

// SelectAs wraps *googlesql.ASTSelectAs.
type SelectAs struct {
	baseNode[*googlesql.ASTSelectAs]
}

func newSelectAs(r *googlesql.ASTSelectAs) *SelectAs {
	if r == nil {
		return nil
	}
	return &SelectAs{baseNode[*googlesql.ASTSelectAs]{raw: r}}
}

func (n *SelectAs) AsMode() AsMode { return must(n.raw.AsMode()) }

func (n *SelectAs) TypeName() *PathExpression {
	return newPathExpression(must(n.raw.TypeName()))
}

// SelectColumn wraps *googlesql.ASTSelectColumn.
type SelectColumn struct {
	baseNode[*googlesql.ASTSelectColumn]
}

func newSelectColumn(r *googlesql.ASTSelectColumn) *SelectColumn {
	if r == nil {
		return nil
	}
	return &SelectColumn{baseNode[*googlesql.ASTSelectColumn]{raw: r}}
}

func (n *SelectColumn) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *SelectColumn) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *SelectColumn) GroupingItemOrder() *GroupingItemOrder {
	return newGroupingItemOrder(must(n.raw.GroupingItemOrder()))
}

// SelectList wraps *googlesql.ASTSelectList.
type SelectList struct {
	baseNode[*googlesql.ASTSelectList]
}

func newSelectList(r *googlesql.ASTSelectList) *SelectList {
	if r == nil {
		return nil
	}
	return &SelectList{baseNode[*googlesql.ASTSelectList]{raw: r}}
}

// Columns returns all select columns.
// Raw Columns(i) returns *SelectColumn (concrete) → we return []*SelectColumn.
func (n *SelectList) Columns() []*SelectColumn {
	count := n.NumChildren()
	result := make([]*SelectColumn, 0, count)
	for i := range count {
		c := must(n.raw.Columns(int32(i)))
		if c == nil {
			break
		}
		result = append(result, newSelectColumn(c))
	}
	return result
}

// SetOperationAllOrDistinct wraps *googlesql.ASTSetOperationAllOrDistinct.
type SetOperationAllOrDistinct struct {
	baseNode[*googlesql.ASTSetOperationAllOrDistinct]
}

func newSetOperationAllOrDistinct(r *googlesql.ASTSetOperationAllOrDistinct) *SetOperationAllOrDistinct {
	if r == nil {
		return nil
	}
	return &SetOperationAllOrDistinct{baseNode[*googlesql.ASTSetOperationAllOrDistinct]{raw: r}}
}

func (n *SetOperationAllOrDistinct) Value() AllOrDistinct { return must(n.raw.Value()) }

// SetOperationColumnMatchMode wraps *googlesql.ASTSetOperationColumnMatchMode.
type SetOperationColumnMatchMode struct {
	baseNode[*googlesql.ASTSetOperationColumnMatchMode]
}

func newSetOperationColumnMatchMode(r *googlesql.ASTSetOperationColumnMatchMode) *SetOperationColumnMatchMode {
	if r == nil {
		return nil
	}
	return &SetOperationColumnMatchMode{baseNode[*googlesql.ASTSetOperationColumnMatchMode]{raw: r}}
}

func (n *SetOperationColumnMatchMode) Value() ColumnMatchMode { return must(n.raw.Value()) }

// SetOperationColumnPropagationMode wraps *googlesql.ASTSetOperationColumnPropagationMode.
type SetOperationColumnPropagationMode struct {
	baseNode[*googlesql.ASTSetOperationColumnPropagationMode]
}

func newSetOperationColumnPropagationMode(r *googlesql.ASTSetOperationColumnPropagationMode) *SetOperationColumnPropagationMode {
	if r == nil {
		return nil
	}
	return &SetOperationColumnPropagationMode{baseNode[*googlesql.ASTSetOperationColumnPropagationMode]{raw: r}}
}

func (n *SetOperationColumnPropagationMode) Value() ColumnPropagationMode {
	return must(n.raw.Value())
}

// SetOperationMetadata wraps *googlesql.ASTSetOperationMetadata.
type SetOperationMetadata struct {
	baseNode[*googlesql.ASTSetOperationMetadata]
}

func newSetOperationMetadata(r *googlesql.ASTSetOperationMetadata) *SetOperationMetadata {
	if r == nil {
		return nil
	}
	return &SetOperationMetadata{baseNode[*googlesql.ASTSetOperationMetadata]{raw: r}}
}

func (n *SetOperationMetadata) OpType() *SetOperationType {
	return newSetOperationType(must(n.raw.OpType()))
}

func (n *SetOperationMetadata) AllOrDistinct() *SetOperationAllOrDistinct {
	return newSetOperationAllOrDistinct(must(n.raw.AllOrDistinct()))
}

func (n *SetOperationMetadata) ColumnMatchMode() *SetOperationColumnMatchMode {
	return newSetOperationColumnMatchMode(must(n.raw.ColumnMatchMode()))
}

func (n *SetOperationMetadata) ColumnPropagationMode() *SetOperationColumnPropagationMode {
	return newSetOperationColumnPropagationMode(must(n.raw.ColumnPropagationMode()))
}

func (n *SetOperationMetadata) CorrespondingByColumnList() *ColumnList {
	return newColumnList(must(n.raw.CorrespondingByColumnList()))
}

func (n *SetOperationMetadata) Hint() *Hint { return newHint(must(n.raw.Hint())) }

// SetOperationMetadataList wraps *googlesql.ASTSetOperationMetadataList.
type SetOperationMetadataList struct {
	baseNode[*googlesql.ASTSetOperationMetadataList]
}

func newSetOperationMetadataList(r *googlesql.ASTSetOperationMetadataList) *SetOperationMetadataList {
	if r == nil {
		return nil
	}
	return &SetOperationMetadataList{baseNode[*googlesql.ASTSetOperationMetadataList]{raw: r}}
}

// Items returns all set operation metadata items.
func (n *SetOperationMetadataList) Items() []*SetOperationMetadata {
	count := n.NumChildren()
	result := make([]*SetOperationMetadata, 0, count)
	for i := range count {
		m := must(n.raw.SetOperationMetadataList(int32(i)))
		if m == nil {
			break
		}
		result = append(result, newSetOperationMetadata(m))
	}
	return result
}

// StarExceptList wraps *googlesql.ASTStarExceptList.
type StarExceptList struct {
	baseNode[*googlesql.ASTStarExceptList]
}

func newStarExceptList(r *googlesql.ASTStarExceptList) *StarExceptList {
	if r == nil {
		return nil
	}
	return &StarExceptList{baseNode[*googlesql.ASTStarExceptList]{raw: r}}
}

// Identifiers: raw returns *Identifier (concrete) → []*Identifier.
func (n *StarExceptList) Identifiers() []*Identifier {
	count := n.NumChildren()
	result := make([]*Identifier, 0, count)
	for i := range count {
		id := must(n.raw.Identifiers(int32(i)))
		if id == nil {
			break
		}
		result = append(result, newIdentifier(id))
	}
	return result
}

// StarModifiers wraps *googlesql.ASTStarModifiers.
type StarModifiers struct {
	baseNode[*googlesql.ASTStarModifiers]
}

func newStarModifiers(r *googlesql.ASTStarModifiers) *StarModifiers {
	if r == nil {
		return nil
	}
	return &StarModifiers{baseNode[*googlesql.ASTStarModifiers]{raw: r}}
}

func (n *StarModifiers) ExceptList() *StarExceptList {
	return newStarExceptList(must(n.raw.ExceptList()))
}

// ReplaceItems: raw returns *StarReplaceItem (concrete) → []*StarReplaceItem.
func (n *StarModifiers) ReplaceItems() []*StarReplaceItem {
	var result []*StarReplaceItem
	for _, c := range n.Children() {
		if item, ok := c.(*StarReplaceItem); ok {
			result = append(result, item)
		}
	}
	return result
}

// StarReplaceItem wraps *googlesql.ASTStarReplaceItem.
type StarReplaceItem struct {
	baseNode[*googlesql.ASTStarReplaceItem]
}

func newStarReplaceItem(r *googlesql.ASTStarReplaceItem) *StarReplaceItem {
	if r == nil {
		return nil
	}
	return &StarReplaceItem{baseNode[*googlesql.ASTStarReplaceItem]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *StarReplaceItem) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *StarReplaceItem) Alias() *Identifier {
	return newIdentifier(must(n.raw.Alias()))
}

// StatementList wraps *googlesql.ASTStatementList.
type StatementList struct {
	baseNode[*googlesql.ASTStatementList]
}

func newStatementList(r *googlesql.ASTStatementList) *StatementList {
	if r == nil {
		return nil
	}
	return &StatementList{baseNode[*googlesql.ASTStatementList]{raw: r}}
}

func (n *StatementList) Statements() []StatementNode {
	var result []StatementNode
	for _, c := range n.Children() {
		if stmt, ok := c.(StatementNode); ok {
			result = append(result, stmt)
		}
	}
	return result
}

// StructConstructorArg wraps *googlesql.ASTStructConstructorArg.
type StructConstructorArg struct {
	baseNode[*googlesql.ASTStructConstructorArg]
}

func newStructConstructorArg(r *googlesql.ASTStructConstructorArg) *StructConstructorArg {
	if r == nil {
		return nil
	}
	return &StructConstructorArg{baseNode[*googlesql.ASTStructConstructorArg]{raw: r}}
}

func (n *StructConstructorArg) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *StructConstructorArg) Alias() *Alias {
	return newAlias(must(n.raw.Alias()))
}

// TVFArgument wraps *googlesql.ASTTVFArgument.
// Script is already declared in nodes_query.go.
type TVFArgument struct {
	baseNode[*googlesql.ASTTVFArgument]
}

func newTVFArgument(r *googlesql.ASTTVFArgument) *TVFArgument {
	if r == nil {
		return nil
	}
	return &TVFArgument{baseNode[*googlesql.ASTTVFArgument]{raw: r}}
}

func (n *TVFArgument) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *TVFArgument) TableClause() Node {
	return Wrap(must(n.raw.TableClause()))
}

func (n *TVFArgument) ModelClause() Node {
	return Wrap(must(n.raw.ModelClause()))
}

func (n *TVFArgument) ConnectionClause() Node {
	return Wrap(must(n.raw.ConnectionClause()))
}

func (n *TVFArgument) Descriptor() Node {
	return Wrap(must(n.raw.Descriptor()))
}

// TVFSchemaColumn wraps *googlesql.ASTTVFSchemaColumn.
type TVFSchemaColumn struct {
	baseNode[*googlesql.ASTTVFSchemaColumn]
}

func newTVFSchemaColumn(r *googlesql.ASTTVFSchemaColumn) *TVFSchemaColumn {
	if r == nil {
		return nil
	}
	return &TVFSchemaColumn{baseNode[*googlesql.ASTTVFSchemaColumn]{raw: r}}
}

func (n *TVFSchemaColumn) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

func (n *TVFSchemaColumn) Type() Node {
	return Wrap(must(n.raw.Type()))
}

// UpdateItem wraps *googlesql.ASTUpdateItem.
type UpdateItem struct {
	baseNode[*googlesql.ASTUpdateItem]
}

func newUpdateItem(r *googlesql.ASTUpdateItem) *UpdateItem {
	if r == nil {
		return nil
	}
	return &UpdateItem{baseNode[*googlesql.ASTUpdateItem]{raw: r}}
}

func (n *UpdateItem) SetValue() *UpdateSetValue {
	return newUpdateSetValue(must(n.raw.SetValue()))
}

func (n *UpdateItem) InsertStatement() *InsertStatement {
	return newInsertStatement(must(n.raw.InsertStatement()))
}

func (n *UpdateItem) DeleteStatement() Node { return Wrap(must(n.raw.DeleteStatement())) }

func (n *UpdateItem) UpdateStatement() Node { return Wrap(must(n.raw.UpdateStatement())) }

// UpdateItemList wraps *googlesql.ASTUpdateItemList.
type UpdateItemList struct {
	baseNode[*googlesql.ASTUpdateItemList]
}

func newUpdateItemList(r *googlesql.ASTUpdateItemList) *UpdateItemList {
	if r == nil {
		return nil
	}
	return &UpdateItemList{baseNode[*googlesql.ASTUpdateItemList]{raw: r}}
}

// Items returns []*UpdateItem children.
func (n *UpdateItemList) Items() []*UpdateItem {
	var result []*UpdateItem
	for _, c := range n.Children() {
		if it, ok := c.(*UpdateItem); ok {
			result = append(result, it)
		}
	}
	return result
}

// UpdateSetValue wraps *googlesql.ASTUpdateSetValue.
type UpdateSetValue struct {
	baseNode[*googlesql.ASTUpdateSetValue]
}

func newUpdateSetValue(r *googlesql.ASTUpdateSetValue) *UpdateSetValue {
	if r == nil {
		return nil
	}
	return &UpdateSetValue{baseNode[*googlesql.ASTUpdateSetValue]{raw: r}}
}

func (n *UpdateSetValue) Path() Node { return Wrap(must(n.raw.Path())) }

func (n *UpdateSetValue) Value() ExpressionNode { return wrapExpr(must(n.raw.Value())) }

// WithClauseEntry wraps *googlesql.ASTWithClauseEntry.
type WithClauseEntry struct {
	baseNode[*googlesql.ASTWithClauseEntry]
}

func newWithClauseEntry(r *googlesql.ASTWithClauseEntry) *WithClauseEntry {
	if r == nil {
		return nil
	}
	return &WithClauseEntry{baseNode[*googlesql.ASTWithClauseEntry]{raw: r}}
}

func (n *WithClauseEntry) AliasedQuery() *AliasedQuery {
	return newAliasedQuery(must(n.raw.AliasedQuery()))
}

func (n *WithClauseEntry) AliasedGroupRows() *AliasedGroupRows {
	return newAliasedGroupRows(must(n.raw.AliasedGroupRows()))
}

// WithModifier wraps *googlesql.ASTWithModifier.
type WithModifier struct {
	baseNode[*googlesql.ASTWithModifier]
}

func newWithModifier(r *googlesql.ASTWithModifier) *WithModifier {
	if r == nil {
		return nil
	}
	return &WithModifier{baseNode[*googlesql.ASTWithModifier]{raw: r}}
}

// Subpipeline wraps *googlesql.ASTSubpipeline.
type Subpipeline struct {
	baseNode[*googlesql.ASTSubpipeline]
}

func newSubpipeline(r *googlesql.ASTSubpipeline) *Subpipeline {
	if r == nil {
		return nil
	}
	return &Subpipeline{baseNode[*googlesql.ASTSubpipeline]{raw: r}}
}

func (n *Subpipeline) PipeOperatorList() []PipeOperatorNode {
	var result []PipeOperatorNode
	for _, c := range n.Children() {
		if op, ok := c.(PipeOperatorNode); ok {
			result = append(result, op)
		}
	}
	return result
}

// WhenThenClause wraps *googlesql.ASTWhenThenClause.
type WhenThenClause struct {
	baseNode[*googlesql.ASTWhenThenClause]
}

func newWhenThenClause(r *googlesql.ASTWhenThenClause) *WhenThenClause {
	if r == nil {
		return nil
	}
	return &WhenThenClause{baseNode[*googlesql.ASTWhenThenClause]{raw: r}}
}

func (n *WhenThenClause) Condition() ExpressionNode {
	return wrapExpr(must(n.raw.Condition()))
}

func (n *WhenThenClause) Body() *StatementList {
	return newStatementList(must(n.raw.Body()))
}

// WhenThenClauseList wraps *googlesql.ASTWhenThenClauseList.
type WhenThenClauseList struct {
	baseNode[*googlesql.ASTWhenThenClauseList]
}

func newWhenThenClauseList(r *googlesql.ASTWhenThenClauseList) *WhenThenClauseList {
	if r == nil {
		return nil
	}
	return &WhenThenClauseList{baseNode[*googlesql.ASTWhenThenClauseList]{raw: r}}
}

func (n *WhenThenClauseList) WhenThenClauses() []*WhenThenClause {
	var result []*WhenThenClause
	for item := range childrenOfType[*googlesql.ASTWhenThenClause](n) {
		result = append(result, newWhenThenClause(item))
	}
	return result
}

// IndexItemList wraps *googlesql.ASTIndexItemList.
type IndexItemList struct {
	baseNode[*googlesql.ASTIndexItemList]
}

func newIndexItemList(r *googlesql.ASTIndexItemList) *IndexItemList {
	if r == nil {
		return nil
	}
	return &IndexItemList{baseNode[*googlesql.ASTIndexItemList]{raw: r}}
}

// IndexStoringExpressionList wraps *googlesql.ASTIndexStoringExpressionList.
type IndexStoringExpressionList struct {
	baseNode[*googlesql.ASTIndexStoringExpressionList]
}

func newIndexStoringExpressionList(r *googlesql.ASTIndexStoringExpressionList) *IndexStoringExpressionList {
	if r == nil {
		return nil
	}
	return &IndexStoringExpressionList{baseNode[*googlesql.ASTIndexStoringExpressionList]{raw: r}}
}

// IndexUnnestExpressionList wraps *googlesql.ASTIndexUnnestExpressionList.
type IndexUnnestExpressionList struct {
	baseNode[*googlesql.ASTIndexUnnestExpressionList]
}

func newIndexUnnestExpressionList(r *googlesql.ASTIndexUnnestExpressionList) *IndexUnnestExpressionList {
	if r == nil {
		return nil
	}
	return &IndexUnnestExpressionList{baseNode[*googlesql.ASTIndexUnnestExpressionList]{raw: r}}
}

// IndexAllColumns wraps *googlesql.ASTIndexAllColumns.
type IndexAllColumns struct {
	baseNode[*googlesql.ASTIndexAllColumns]
}

func newIndexAllColumns(r *googlesql.ASTIndexAllColumns) *IndexAllColumns {
	if r == nil {
		return nil
	}
	return &IndexAllColumns{baseNode[*googlesql.ASTIndexAllColumns]{raw: r}}
}

// TransformClause wraps *googlesql.ASTTransformClause.
type TransformClause struct {
	baseNode[*googlesql.ASTTransformClause]
}

func newTransformClause(r *googlesql.ASTTransformClause) *TransformClause {
	if r == nil {
		return nil
	}
	return &TransformClause{baseNode[*googlesql.ASTTransformClause]{raw: r}}
}

func (n *TransformClause) SelectList() *SelectList {
	return newSelectList(must(n.raw.SelectList()))
}

// InputOutputClause wraps *googlesql.ASTInputOutputClause.
type InputOutputClause struct {
	baseNode[*googlesql.ASTInputOutputClause]
}

func newInputOutputClause(r *googlesql.ASTInputOutputClause) *InputOutputClause {
	if r == nil {
		return nil
	}
	return &InputOutputClause{baseNode[*googlesql.ASTInputOutputClause]{raw: r}}
}

func (n *InputOutputClause) Input() *TableElementList {
	return newTableElementList(must(n.raw.Input()))
}

func (n *InputOutputClause) Output() *TableElementList {
	return newTableElementList(must(n.raw.Output()))
}

// Privilege wraps *googlesql.ASTPrivilege.
type Privilege struct {
	baseNode[*googlesql.ASTPrivilege]
}

func newPrivilege(r *googlesql.ASTPrivilege) *Privilege {
	if r == nil {
		return nil
	}
	return &Privilege{baseNode[*googlesql.ASTPrivilege]{raw: r}}
}

func (n *Privilege) PrivilegeAction() *Identifier {
	return newIdentifier(must(n.raw.PrivilegeAction()))
}
func (n *Privilege) Paths() *PathExpressionList { return newPathExpressionList(must(n.raw.Paths())) }

// Privileges wraps *googlesql.ASTPrivileges.
type Privileges struct {
	baseNode[*googlesql.ASTPrivileges]
}

func newPrivileges(r *googlesql.ASTPrivileges) *Privileges {
	if r == nil {
		return nil
	}
	return &Privileges{baseNode[*googlesql.ASTPrivileges]{raw: r}}
}
func (n *Privileges) IsAllPrivileges() bool { return must(n.raw.IsAllPrivileges()) }
func (n *Privileges) Privileges() []*Privilege {
	var result []*Privilege
	for item := range childrenOfType[*googlesql.ASTPrivilege](n) {
		result = append(result, newPrivilege(item))
	}
	return result
}

// SequenceArg wraps *googlesql.ASTSequenceArg.
type SequenceArg struct {
	baseNode[*googlesql.ASTSequenceArg]
}

func newSequenceArg(r *googlesql.ASTSequenceArg) *SequenceArg {
	if r == nil {
		return nil
	}
	return &SequenceArg{baseNode[*googlesql.ASTSequenceArg]{raw: r}}
}

func (n *SequenceArg) SequencePath() *PathExpression {
	return newPathExpression(must(n.raw.SequencePath()))
}

// IdentityColumnInfo wraps *googlesql.ASTIdentityColumnInfo.
type IdentityColumnInfo struct {
	baseNode[*googlesql.ASTIdentityColumnInfo]
}

func newIdentityColumnInfo(r *googlesql.ASTIdentityColumnInfo) *IdentityColumnInfo {
	if r == nil {
		return nil
	}
	return &IdentityColumnInfo{baseNode[*googlesql.ASTIdentityColumnInfo]{raw: r}}
}
func (n *IdentityColumnInfo) CyclingEnabled() bool   { return must(n.raw.CyclingEnabled()) }
func (n *IdentityColumnInfo) StartWithValue() Node   { return Wrap(must(n.raw.StartWithValue())) }
func (n *IdentityColumnInfo) IncrementByValue() Node { return Wrap(must(n.raw.IncrementByValue())) }
func (n *IdentityColumnInfo) MaxValue() Node         { return Wrap(must(n.raw.MaxValue())) }
func (n *IdentityColumnInfo) MinValue() Node         { return Wrap(must(n.raw.MinValue())) }

// AliasedQueryExpression wraps *googlesql.ASTAliasedQueryExpression.
type AliasedQueryExpression struct {
	baseNode[*googlesql.ASTAliasedQueryExpression]
}

func newAliasedQueryExpression(r *googlesql.ASTAliasedQueryExpression) *AliasedQueryExpression {
	if r == nil {
		return nil
	}
	return &AliasedQueryExpression{baseNode[*googlesql.ASTAliasedQueryExpression]{raw: r}}
}
func (n *AliasedQueryExpression) isQueryExpression() {}
func (n *AliasedQueryExpression) Alias() *Alias { return newAlias(must(n.raw.Alias())) }
func (n *AliasedQueryExpression) Query() *Query { return newQuery(must(n.raw.Query())) }

// AliasedQueryList wraps *googlesql.ASTAliasedQueryList.
type AliasedQueryList struct {
	baseNode[*googlesql.ASTAliasedQueryList]
}

func newAliasedQueryList(r *googlesql.ASTAliasedQueryList) *AliasedQueryList {
	if r == nil {
		return nil
	}
	return &AliasedQueryList{baseNode[*googlesql.ASTAliasedQueryList]{raw: r}}
}

// AliasedQueryModifiers wraps *googlesql.ASTAliasedQueryModifiers.
type AliasedQueryModifiers struct {
	baseNode[*googlesql.ASTAliasedQueryModifiers]
}

func newAliasedQueryModifiers(r *googlesql.ASTAliasedQueryModifiers) *AliasedQueryModifiers {
	if r == nil {
		return nil
	}
	return &AliasedQueryModifiers{baseNode[*googlesql.ASTAliasedQueryModifiers]{raw: r}}
}

// TableAndColumnInfo wraps *googlesql.ASTTableAndColumnInfo.
type TableAndColumnInfo struct {
	baseNode[*googlesql.ASTTableAndColumnInfo]
}

func newTableAndColumnInfo(r *googlesql.ASTTableAndColumnInfo) *TableAndColumnInfo {
	if r == nil {
		return nil
	}
	return &TableAndColumnInfo{baseNode[*googlesql.ASTTableAndColumnInfo]{raw: r}}
}

func (n *TableAndColumnInfo) TableName() *PathExpression {
	return newPathExpression(must(n.raw.TableName()))
}
func (n *TableAndColumnInfo) ColumnList() *ColumnList { return newColumnList(must(n.raw.ColumnList())) }

// TableAndColumnInfoList wraps *googlesql.ASTTableAndColumnInfoList.
type TableAndColumnInfoList struct {
	baseNode[*googlesql.ASTTableAndColumnInfoList]
}

func newTableAndColumnInfoList(r *googlesql.ASTTableAndColumnInfoList) *TableAndColumnInfoList {
	if r == nil {
		return nil
	}
	return &TableAndColumnInfoList{baseNode[*googlesql.ASTTableAndColumnInfoList]{raw: r}}
}

// MacroBody wraps *googlesql.ASTMacroBody.
type MacroBody struct {
	baseNode[*googlesql.ASTMacroBody]
}

func newMacroBody(r *googlesql.ASTMacroBody) *MacroBody {
	if r == nil {
		return nil
	}
	return &MacroBody{baseNode[*googlesql.ASTMacroBody]{raw: r}}
}

// IntOrUnbounded wraps *googlesql.ASTIntOrUnbounded.
type IntOrUnbounded struct {
	baseNode[*googlesql.ASTIntOrUnbounded]
}

func newIntOrUnbounded(r *googlesql.ASTIntOrUnbounded) *IntOrUnbounded {
	if r == nil {
		return nil
	}
	return &IntOrUnbounded{baseNode[*googlesql.ASTIntOrUnbounded]{raw: r}}
}
func (n *IntOrUnbounded) Bound() ExpressionNode { return wrapExpr(must(n.raw.Bound())) }

// RecursionDepthModifier wraps *googlesql.ASTRecursionDepthModifier.
type RecursionDepthModifier struct {
	baseNode[*googlesql.ASTRecursionDepthModifier]
}

func newRecursionDepthModifier(r *googlesql.ASTRecursionDepthModifier) *RecursionDepthModifier {
	if r == nil {
		return nil
	}
	return &RecursionDepthModifier{baseNode[*googlesql.ASTRecursionDepthModifier]{raw: r}}
}
func (n *RecursionDepthModifier) Alias() *Alias { return newAlias(must(n.raw.Alias())) }
func (n *RecursionDepthModifier) LowerBound() *IntOrUnbounded {
	return newIntOrUnbounded(must(n.raw.LowerBound()))
}

func (n *RecursionDepthModifier) UpperBound() *IntOrUnbounded {
	return newIntOrUnbounded(must(n.raw.UpperBound()))
}
