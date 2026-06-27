package sql

import "github.com/goccy/go-googlesql"

// ─── INSERT ───────────────────────────────────────────────────────────────────

// InsertValuesRow wraps *googlesql.ASTInsertValuesRow.
type InsertValuesRow struct {
	baseNode[*googlesql.ASTInsertValuesRow]
}

func newASTInsertValuesRow(r *googlesql.ASTInsertValuesRow) *InsertValuesRow {
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

func newASTInsertValuesRowList(r *googlesql.ASTInsertValuesRowList) *InsertValuesRowList {
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

func newASTInsertStatement(r *googlesql.ASTInsertStatement) *InsertStatement {
	if r == nil {
		return nil
	}
	return &InsertStatement{baseNode[*googlesql.ASTInsertStatement]{raw: r}}
}
func (n *InsertStatement) isStatement()     {}
func (n *InsertStatement) TargetPath() Node { return Wrap(must(n.raw.TargetPath())) }
func (n *InsertStatement) ColumnList() *ColumnList {
	return newASTColumnList(must(n.raw.ColumnList()))
}
func (n *InsertStatement) Query() *Query { return newASTQuery(must(n.raw.Query())) }
func (n *InsertStatement) Rows() *InsertValuesRowList {
	return newASTInsertValuesRowList(must(n.raw.Rows()))
}

// ─── UPDATE ───────────────────────────────────────────────────────────────────

// UpdateSetValue wraps *googlesql.ASTUpdateSetValue.
type UpdateSetValue struct {
	baseNode[*googlesql.ASTUpdateSetValue]
}

func newASTUpdateSetValue(r *googlesql.ASTUpdateSetValue) *UpdateSetValue {
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

func newASTUpdateItem(r *googlesql.ASTUpdateItem) *UpdateItem {
	if r == nil {
		return nil
	}
	return &UpdateItem{baseNode[*googlesql.ASTUpdateItem]{raw: r}}
}

func (n *UpdateItem) SetValue() *UpdateSetValue {
	return newASTUpdateSetValue(must(n.raw.SetValue()))
}

func (n *UpdateItem) InsertStatement() *InsertStatement {
	return newASTInsertStatement(must(n.raw.InsertStatement()))
}
func (n *UpdateItem) DeleteStatement() Node { return Wrap(must(n.raw.DeleteStatement())) }
func (n *UpdateItem) UpdateStatement() Node { return Wrap(must(n.raw.UpdateStatement())) }

// UpdateItemList wraps *googlesql.ASTUpdateItemList.
type UpdateItemList struct {
	baseNode[*googlesql.ASTUpdateItemList]
}

func newASTUpdateItemList(r *googlesql.ASTUpdateItemList) *UpdateItemList {
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

// ─── MERGE ────────────────────────────────────────────────────────────────────

// MergeAction wraps *googlesql.ASTMergeAction.
type MergeAction struct {
	baseNode[*googlesql.ASTMergeAction]
}

func newASTMergeAction(r *googlesql.ASTMergeAction) *MergeAction {
	if r == nil {
		return nil
	}
	return &MergeAction{baseNode[*googlesql.ASTMergeAction]{raw: r}}
}
func (n *MergeAction) ActionType() MergeActionType { return must(n.raw.ActionType()) }
func (n *MergeAction) InsertColumnList() *ColumnList {
	return newASTColumnList(must(n.raw.InsertColumnList()))
}

func (n *MergeAction) InsertRow() *InsertValuesRow {
	return newASTInsertValuesRow(must(n.raw.InsertRow()))
}

func (n *MergeAction) UpdateItemList() *UpdateItemList {
	return newASTUpdateItemList(must(n.raw.UpdateItemList()))
}

// MergeWhenClause wraps *googlesql.ASTMergeWhenClause.
type MergeWhenClause struct {
	baseNode[*googlesql.ASTMergeWhenClause]
}

func newASTMergeWhenClause(r *googlesql.ASTMergeWhenClause) *MergeWhenClause {
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
	return newASTMergeAction(must(n.raw.Action()))
}

// MergeWhenClauseList wraps *googlesql.ASTMergeWhenClauseList.
type MergeWhenClauseList struct {
	baseNode[*googlesql.ASTMergeWhenClauseList]
}

func newASTMergeWhenClauseList(r *googlesql.ASTMergeWhenClauseList) *MergeWhenClauseList {
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

func newASTMergeStatement(r *googlesql.ASTMergeStatement) *MergeStatement {
	if r == nil {
		return nil
	}
	return &MergeStatement{baseNode[*googlesql.ASTMergeStatement]{raw: r}}
}
func (n *MergeStatement) isStatement() {}

func (n *MergeStatement) TargetPath() *PathExpression {
	return newASTPathExpression(must(n.raw.TargetPath()))
}
func (n *MergeStatement) Alias() *Alias { return newASTAlias(must(n.raw.Alias())) }
func (n *MergeStatement) TableExpression() TableExpressionNode {
	return wrapTableExpr(must(n.raw.TableExpression()))
}

func (n *MergeStatement) MergeCondition() ExpressionNode {
	return wrapExpr(must(n.raw.MergeCondition()))
}

func (n *MergeStatement) WhenClauses() *MergeWhenClauseList {
	return newASTMergeWhenClauseList(must(n.raw.WhenClauses()))
}

// ─── TRUNCATE ─────────────────────────────────────────────────────────────────

// TruncateStatement wraps *googlesql.ASTTruncateStatement.
type TruncateStatement struct {
	baseNode[*googlesql.ASTTruncateStatement]
}

func newASTTruncateStatement(r *googlesql.ASTTruncateStatement) *TruncateStatement {
	if r == nil {
		return nil
	}
	return &TruncateStatement{baseNode[*googlesql.ASTTruncateStatement]{raw: r}}
}
func (n *TruncateStatement) isStatement() {}

func (n *TruncateStatement) TargetPath() *PathExpression {
	return newASTPathExpression(must(n.raw.TargetPath()))
}
func (n *TruncateStatement) Where() ExpressionNode { return wrapExpr(must(n.raw.Where())) }

// ─── ASSIGNMENT ───────────────────────────────────────────────────────────────

// AssignmentFromStruct wraps *googlesql.ASTAssignmentFromStruct.
type AssignmentFromStruct struct {
	baseNode[*googlesql.ASTAssignmentFromStruct]
}

func newASTAssignmentFromStruct(r *googlesql.ASTAssignmentFromStruct) *AssignmentFromStruct {
	if r == nil {
		return nil
	}
	return &AssignmentFromStruct{baseNode[*googlesql.ASTAssignmentFromStruct]{raw: r}}
}
func (n *AssignmentFromStruct) isStatement() {}

func (n *AssignmentFromStruct) Variables() *IdentifierList {
	return newASTIdentifierList(must(n.raw.Variables()))
}

func (n *AssignmentFromStruct) StructExpression() ExpressionNode {
	return wrapExpr(must(n.raw.StructExpression()))
}
