package sql

import "github.com/goccy/go-googlesql"

type DeleteStatement struct {
	baseNode[*googlesql.ASTDeleteStatement]
}

func newDeleteStatement(r *googlesql.ASTDeleteStatement) *DeleteStatement {
	if r == nil {
		return nil
	}
	return &DeleteStatement{baseNode[*googlesql.ASTDeleteStatement]{raw: r}}
}
func (n *DeleteStatement) isStatement() {}
func (n *DeleteStatement) TargetPath() Node {
	return Wrap(must(n.raw.TargetPath()))
}

func (n *DeleteStatement) Alias() *Alias {
	return newAlias(must(n.raw.Alias()))
}

func (n *DeleteStatement) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *DeleteStatement) Offset() *WithOffset {
	return newWithOffset(must(n.raw.Offset()))
}

func (n *DeleteStatement) Where() ExpressionNode {
	return wrapExpr(must(n.raw.Where()))
}

func (n *DeleteStatement) AssertRowsModified() *AssertRowsModified {
	return newAssertRowsModified(must(n.raw.AssertRowsModified()))
}

func (n *DeleteStatement) Returning() *ReturningClause {
	return newReturningClause(must(n.raw.Returning()))
}

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

type ReturningClause struct {
	baseNode[*googlesql.ASTReturningClause]
}

func newReturningClause(r *googlesql.ASTReturningClause) *ReturningClause {
	if r == nil {
		return nil
	}
	return &ReturningClause{baseNode[*googlesql.ASTReturningClause]{raw: r}}
}

func (n *ReturningClause) SelectList() *SelectList {
	return newSelectList(must(n.raw.SelectList()))
}

func (n *ReturningClause) ActionAlias() *Alias {
	return newAlias(must(n.raw.ActionAlias()))
}

// ─── INSERT ───────────────────────────────────────────────────────────────────

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

// InsertStatement wraps *googlesql.ASTInsertStatement.
type InsertStatement struct {
	baseNode[*googlesql.ASTInsertStatement]
}

func newInsertStatement(r *googlesql.ASTInsertStatement) *InsertStatement {
	if r == nil {
		return nil
	}
	return &InsertStatement{baseNode[*googlesql.ASTInsertStatement]{raw: r}}
}
func (n *InsertStatement) isStatement()     {}
func (n *InsertStatement) TargetPath() Node { return Wrap(must(n.raw.TargetPath())) }
func (n *InsertStatement) ColumnList() *ColumnList {
	return newColumnList(must(n.raw.ColumnList()))
}
func (n *InsertStatement) Query() *Query { return newQuery(must(n.raw.Query())) }
func (n *InsertStatement) Rows() *InsertValuesRowList {
	return newInsertValuesRowList(must(n.raw.Rows()))
}

func (n *InsertStatement) AssertRowsModified() *AssertRowsModified {
	return newAssertRowsModified(must(n.raw.AssertRowsModified()))
}

func (n *InsertStatement) Returning() *ReturningClause {
	return newReturningClause(must(n.raw.Returning()))
}

// ─── UPDATE ───────────────────────────────────────────────────────────────────

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
func (n *UpdateSetValue) Path() Node            { return Wrap(must(n.raw.Path())) }
func (n *UpdateSetValue) Value() ExpressionNode { return wrapExpr(must(n.raw.Value())) }

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

type UpdateStatement struct {
	baseNode[*googlesql.ASTUpdateStatement]
}

func newUpdateStatement(r *googlesql.ASTUpdateStatement) *UpdateStatement {
	if r == nil {
		return nil
	}
	return &UpdateStatement{baseNode[*googlesql.ASTUpdateStatement]{raw: r}}
}
func (n *UpdateStatement) isStatement()        {}
func (n *UpdateStatement) TargetPath() Node    { return Wrap(must(n.raw.TargetPath())) }
func (n *UpdateStatement) Alias() *Alias       { return newAlias(must(n.raw.Alias())) }
func (n *UpdateStatement) Offset() *WithOffset { return newWithOffset(must(n.raw.Offset())) }
func (n *UpdateStatement) UpdateItemList() *UpdateItemList {
	return newUpdateItemList(must(n.raw.UpdateItemList()))
}
func (n *UpdateStatement) FromClause() *FromClause { return newFromClause(must(n.raw.FromClause())) }
func (n *UpdateStatement) Where() ExpressionNode   { return wrapExpr(must(n.raw.Where())) }
func (n *UpdateStatement) AssertRowsModified() *AssertRowsModified {
	return newAssertRowsModified(must(n.raw.AssertRowsModified()))
}

func (n *UpdateStatement) Returning() *ReturningClause {
	return newReturningClause(must(n.raw.Returning()))
}

// ─── MERGE ────────────────────────────────────────────────────────────────────

// MergeAction wraps *googlesql.ASTMergeAction.
type MergeAction struct {
	baseNode[*googlesql.ASTMergeAction]
}

func newMergeAction(r *googlesql.ASTMergeAction) *MergeAction {
	if r == nil {
		return nil
	}
	return &MergeAction{baseNode[*googlesql.ASTMergeAction]{raw: r}}
}
func (n *MergeAction) ActionType() MergeActionType { return must(n.raw.ActionType()) }
func (n *MergeAction) InsertColumnList() *ColumnList {
	return newColumnList(must(n.raw.InsertColumnList()))
}

func (n *MergeAction) InsertRow() *InsertValuesRow {
	return newInsertValuesRow(must(n.raw.InsertRow()))
}

func (n *MergeAction) UpdateItemList() *UpdateItemList {
	return newUpdateItemList(must(n.raw.UpdateItemList()))
}

// MergeWhenClause wraps *googlesql.ASTMergeWhenClause.
type MergeWhenClause struct {
	baseNode[*googlesql.ASTMergeWhenClause]
}

func newMergeWhenClause(r *googlesql.ASTMergeWhenClause) *MergeWhenClause {
	if r == nil {
		return nil
	}
	return &MergeWhenClause{baseNode[*googlesql.ASTMergeWhenClause]{raw: r}}
}
func (n *MergeWhenClause) MatchType() MergeMatchType { return must(n.raw.MatchType()) }
func (n *MergeWhenClause) SearchCondition() ExpressionNode {
	return wrapExpr(must(n.raw.SearchCondition()))
}

func (n *MergeWhenClause) Action() *MergeAction {
	return newMergeAction(must(n.raw.Action()))
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

// MergeStatement wraps *googlesql.ASTMergeStatement.
type MergeStatement struct {
	baseNode[*googlesql.ASTMergeStatement]
}

func newMergeStatement(r *googlesql.ASTMergeStatement) *MergeStatement {
	if r == nil {
		return nil
	}
	return &MergeStatement{baseNode[*googlesql.ASTMergeStatement]{raw: r}}
}
func (n *MergeStatement) isStatement() {}

func (n *MergeStatement) TargetPath() *PathExpression {
	return newPathExpression(must(n.raw.TargetPath()))
}
func (n *MergeStatement) Alias() *Alias { return newAlias(must(n.raw.Alias())) }
func (n *MergeStatement) TableExpression() TableExpressionNode {
	return wrapTableExpr(must(n.raw.TableExpression()))
}

func (n *MergeStatement) MergeCondition() ExpressionNode {
	return wrapExpr(must(n.raw.MergeCondition()))
}

func (n *MergeStatement) WhenClauses() *MergeWhenClauseList {
	return newMergeWhenClauseList(must(n.raw.WhenClauses()))
}

// ─── TRUNCATE ─────────────────────────────────────────────────────────────────

// TruncateStatement wraps *googlesql.ASTTruncateStatement.
type TruncateStatement struct {
	baseNode[*googlesql.ASTTruncateStatement]
}

func newTruncateStatement(r *googlesql.ASTTruncateStatement) *TruncateStatement {
	if r == nil {
		return nil
	}
	return &TruncateStatement{baseNode[*googlesql.ASTTruncateStatement]{raw: r}}
}
func (n *TruncateStatement) isStatement() {}

func (n *TruncateStatement) TargetPath() *PathExpression {
	return newPathExpression(must(n.raw.TargetPath()))
}
func (n *TruncateStatement) Where() ExpressionNode { return wrapExpr(must(n.raw.Where())) }

// ─── ASSIGNMENT ───────────────────────────────────────────────────────────────

// AssignmentFromStruct wraps *googlesql.ASTAssignmentFromStruct.
type AssignmentFromStruct struct {
	baseNode[*googlesql.ASTAssignmentFromStruct]
}

func newAssignmentFromStruct(r *googlesql.ASTAssignmentFromStruct) *AssignmentFromStruct {
	if r == nil {
		return nil
	}
	return &AssignmentFromStruct{baseNode[*googlesql.ASTAssignmentFromStruct]{raw: r}}
}
func (n *AssignmentFromStruct) isStatement() {}

func (n *AssignmentFromStruct) Variables() *IdentifierList {
	return newIdentifierList(must(n.raw.Variables()))
}

func (n *AssignmentFromStruct) StructExpression() ExpressionNode {
	return wrapExpr(must(n.raw.StructExpression()))
}
