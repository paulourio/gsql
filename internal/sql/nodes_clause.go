package sql

import (
	"github.com/goccy/go-googlesql"
)

// ClampedBetweenModifier wraps *googlesql.ASTClampedBetweenModifier.
type ClampedBetweenModifier struct {
	baseNode[*googlesql.ASTClampedBetweenModifier]
}

func newClampedBetweenModifier(r *googlesql.ASTClampedBetweenModifier) *ClampedBetweenModifier {
	if r == nil {
		return nil
	}
	return &ClampedBetweenModifier{baseNode[*googlesql.ASTClampedBetweenModifier]{raw: r}}
}

// Low/High: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *ClampedBetweenModifier) Low() ExpressionNode { return wrapExpr(must(n.raw.Low())) }

func (n *ClampedBetweenModifier) High() ExpressionNode { return wrapExpr(must(n.raw.High())) }

// ClusterBy wraps *googlesql.ASTClusterBy.
type ClusterBy struct {
	baseNode[*googlesql.ASTClusterBy]
}

func newClusterBy(r *googlesql.ASTClusterBy) *ClusterBy {
	if r == nil {
		return nil
	}
	return &ClusterBy{baseNode[*googlesql.ASTClusterBy]{raw: r}}
}

// ClusteringExpressions: raw returns ExpressionNode (interface) → []ExpressionNode.
func (n *ClusterBy) ClusteringExpressions() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		e := must(n.raw.ClusteringExpressions(int32(i)))
		if !defined(e) {
			break
		}
		result = append(result, wrapExpr(e))
	}
	return result
}

// Collate wraps *googlesql.ASTCollate.
type Collate struct {
	baseNode[*googlesql.ASTCollate]
}

func newCollate(r *googlesql.ASTCollate) *Collate {
	if r == nil {
		return nil
	}
	return &Collate{baseNode[*googlesql.ASTCollate]{raw: r}}
}

// CollationName: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *Collate) CollationName() ExpressionNode {
	return wrapExpr(must(n.raw.CollationName()))
}

// ColumnList wraps *googlesql.ASTColumnList.
type ColumnList struct {
	baseNode[*googlesql.ASTColumnList]
}

func newColumnList(r *googlesql.ASTColumnList) *ColumnList {
	if r == nil {
		return nil
	}
	return &ColumnList{baseNode[*googlesql.ASTColumnList]{raw: r}}
}

// Identifiers: raw Identifiers(i) returns *Identifier (concrete) → []*Identifier.
func (n *ColumnList) Identifiers() []*Identifier {
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

// ConnectionClause wraps *googlesql.ASTConnectionClause.
type ConnectionClause struct {
	baseNode[*googlesql.ASTConnectionClause]
}

func newConnectionClause(r *googlesql.ASTConnectionClause) *ConnectionClause {
	if r == nil {
		return nil
	}
	return &ConnectionClause{baseNode[*googlesql.ASTConnectionClause]{raw: r}}
}

func (n *ConnectionClause) ConnectionPath() *PathExpression {
	return newPathExpression(must(n.raw.ConnectionPath()).(*googlesql.ASTPathExpression))
}

// DescriptorColumn wraps *googlesql.ASTDescriptorColumn.
type DescriptorColumn struct {
	baseNode[*googlesql.ASTDescriptorColumn]
}

func newDescriptorColumn(r *googlesql.ASTDescriptorColumn) *DescriptorColumn {
	if r == nil {
		return nil
	}
	return &DescriptorColumn{baseNode[*googlesql.ASTDescriptorColumn]{raw: r}}
}

func (n *DescriptorColumn) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

// DescriptorColumnList wraps *googlesql.ASTDescriptorColumnList.
type DescriptorColumnList struct {
	baseNode[*googlesql.ASTDescriptorColumnList]
}

func newDescriptorColumnList(r *googlesql.ASTDescriptorColumnList) *DescriptorColumnList {
	if r == nil {
		return nil
	}
	return &DescriptorColumnList{baseNode[*googlesql.ASTDescriptorColumnList]{raw: r}}
}

// Columns returns all columns.
func (n *DescriptorColumnList) Columns() []*DescriptorColumn {
	count := n.NumChildren()
	result := make([]*DescriptorColumn, 0, count)
	for i := range count {
		c := must(n.raw.DescriptorColumnList(int32(i)))
		if c == nil {
			break
		}
		result = append(result, newDescriptorColumn(c))
	}
	return result
}

// ElseifClause wraps *googlesql.ASTElseifClause.
type ElseifClause struct {
	baseNode[*googlesql.ASTElseifClause]
}

func newElseifClause(r *googlesql.ASTElseifClause) *ElseifClause {
	if r == nil {
		return nil
	}
	return &ElseifClause{baseNode[*googlesql.ASTElseifClause]{raw: r}}
}

func (n *ElseifClause) Condition() ExpressionNode {
	return wrapExpr(must(n.raw.Condition()))
}

func (n *ElseifClause) Body() *StatementList {
	return newStatementList(must(n.raw.Body()))
}

// ExecuteIntoClause wraps *googlesql.ASTExecuteIntoClause.
type ExecuteIntoClause struct {
	baseNode[*googlesql.ASTExecuteIntoClause]
}

func newExecuteIntoClause(r *googlesql.ASTExecuteIntoClause) *ExecuteIntoClause {
	if r == nil {
		return nil
	}
	return &ExecuteIntoClause{baseNode[*googlesql.ASTExecuteIntoClause]{raw: r}}
}

func (n *ExecuteIntoClause) Identifiers() *IdentifierList {
	return newIdentifierList(must(n.raw.Identifiers()))
}

// ExecuteUsingClause wraps *googlesql.ASTExecuteUsingClause.
type ExecuteUsingClause struct {
	baseNode[*googlesql.ASTExecuteUsingClause]
}

func newExecuteUsingClause(r *googlesql.ASTExecuteUsingClause) *ExecuteUsingClause {
	if r == nil {
		return nil
	}
	return &ExecuteUsingClause{baseNode[*googlesql.ASTExecuteUsingClause]{raw: r}}
}

// Arguments returns all *ExecuteUsingArgument children.
func (n *ExecuteUsingClause) Arguments() []*ExecuteUsingArgument {
	var result []*ExecuteUsingArgument
	for _, c := range n.Children() {
		if a, ok := c.(*ExecuteUsingArgument); ok {
			result = append(result, a)
		}
	}
	return result
}

// FilterUsingClause wraps *googlesql.ASTFilterUsingClause.
type FilterUsingClause struct {
	baseNode[*googlesql.ASTFilterUsingClause]
}

func newFilterUsingClause(r *googlesql.ASTFilterUsingClause) *FilterUsingClause {
	if r == nil {
		return nil
	}
	return &FilterUsingClause{baseNode[*googlesql.ASTFilterUsingClause]{raw: r}}
}

func (n *FilterUsingClause) Predicate() ExpressionNode {
	return wrapExpr(must(n.raw.Predicate()))
}

// ForSystemTime wraps *googlesql.ASTForSystemTime.
type ForSystemTime struct {
	baseNode[*googlesql.ASTForSystemTime]
}

func newForSystemTime(r *googlesql.ASTForSystemTime) *ForSystemTime {
	if r == nil {
		return nil
	}
	return &ForSystemTime{baseNode[*googlesql.ASTForSystemTime]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *ForSystemTime) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// FormatClause wraps *googlesql.ASTFormatClause.
type FormatClause struct {
	baseNode[*googlesql.ASTFormatClause]
}

func newFormatClause(r *googlesql.ASTFormatClause) *FormatClause {
	if r == nil {
		return nil
	}
	return &FormatClause{baseNode[*googlesql.ASTFormatClause]{raw: r}}
}

// Format/TimeZoneExpr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *FormatClause) Format() ExpressionNode { return wrapExpr(must(n.raw.Format())) }

func (n *FormatClause) TimeZoneExpr() ExpressionNode {
	return wrapExpr(must(n.raw.TimeZoneExpr()))
}

// FromClause wraps *googlesql.ASTFromClause.
type FromClause struct {
	baseNode[*googlesql.ASTFromClause]
}

func newFromClause(r *googlesql.ASTFromClause) *FromClause {
	if r == nil {
		return nil
	}
	return &FromClause{baseNode[*googlesql.ASTFromClause]{raw: r}}
}

// TableExpression returns the table expression; raw returns TableExpressionNode (interface).
func (n *FromClause) TableExpression() TableExpressionNode {
	return wrapTableExpr(must(n.raw.TableExpression()))
}

// GrantToClause wraps *googlesql.ASTGrantToClause.
type GrantToClause struct {
	baseNode[*googlesql.ASTGrantToClause]
}

func newGrantToClause(r *googlesql.ASTGrantToClause) *GrantToClause {
	if r == nil {
		return nil
	}
	return &GrantToClause{baseNode[*googlesql.ASTGrantToClause]{raw: r}}
}

func (n *GrantToClause) GranteeList() *GranteeList {
	return newGranteeList(must(n.raw.GranteeList()))
}

// GroupBy wraps *googlesql.ASTGroupBy.
type GroupBy struct {
	baseNode[*googlesql.ASTGroupBy]
}

func newGroupBy(r *googlesql.ASTGroupBy) *GroupBy {
	if r == nil {
		return nil
	}
	return &GroupBy{baseNode[*googlesql.ASTGroupBy]{raw: r}}
}

func (n *GroupBy) Hint() *Hint { return newHint(must(n.raw.Hint())) }

func (n *GroupBy) All() *GroupByAll { return newGroupByAll(must(n.raw.All())) }

func (n *GroupBy) AndOrderBy() bool { return must(n.raw.AndOrderBy()) }

// GroupingItems: raw GroupingItems(i) returns *GroupingItem (concrete) → []*GroupingItem.
func (n *GroupBy) GroupingItems() []*GroupingItem {
	var result []*GroupingItem
	for _, c := range n.Children() {
		if item, ok := c.(*GroupingItem); ok {
			result = append(result, item)
		}
	}
	return result
}

// GroupByAll wraps *googlesql.ASTGroupByAll.
type GroupByAll struct {
	baseNode[*googlesql.ASTGroupByAll]
}

func newGroupByAll(r *googlesql.ASTGroupByAll) *GroupByAll {
	if r == nil {
		return nil
	}
	return &GroupByAll{baseNode[*googlesql.ASTGroupByAll]{raw: r}}
}

// GroupingItem wraps *googlesql.ASTGroupingItem.
type GroupingItem struct {
	baseNode[*googlesql.ASTGroupingItem]
}

func newGroupingItem(r *googlesql.ASTGroupingItem) *GroupingItem {
	if r == nil {
		return nil
	}
	return &GroupingItem{baseNode[*googlesql.ASTGroupingItem]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *GroupingItem) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *GroupingItem) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *GroupingItem) Rollup() *Rollup { return newRollup(must(n.raw.Rollup())) }

func (n *GroupingItem) Cube() *Cube { return newCube(must(n.raw.Cube())) }

// HavingModifier wraps *googlesql.ASTHavingModifier.
type HavingModifier struct {
	baseNode[*googlesql.ASTHavingModifier]
}

func newHavingModifier(r *googlesql.ASTHavingModifier) *HavingModifier {
	if r == nil {
		return nil
	}
	return &HavingModifier{baseNode[*googlesql.ASTHavingModifier]{raw: r}}
}

func (n *HavingModifier) ModifierKind() HavingModifierType {
	return must(n.raw.ModifierKind())
}

// Expr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *HavingModifier) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// Hint wraps *googlesql.ASTHint.
type Hint struct{ baseNode[*googlesql.ASTHint] }

func newHint(r *googlesql.ASTHint) *Hint {
	if r == nil {
		return nil
	}
	return &Hint{baseNode[*googlesql.ASTHint]{raw: r}}
}

func (n *Hint) NumShardsHint() *IntLiteral {
	return newIntLiteral(must(n.raw.NumShardsHint()))
}

// HintEntries: raw HintEntries(i) returns *HintEntry (concrete) → []*HintEntry.
func (n *Hint) HintEntries() []*HintEntry {
	var result []*HintEntry
	for _, c := range n.Children() {
		if e, ok := c.(*HintEntry); ok {
			result = append(result, e)
		}
	}
	return result
}

// HintEntry wraps *googlesql.ASTHintEntry.
type HintEntry struct {
	baseNode[*googlesql.ASTHintEntry]
}

func newHintEntry(r *googlesql.ASTHintEntry) *HintEntry {
	if r == nil {
		return nil
	}
	return &HintEntry{baseNode[*googlesql.ASTHintEntry]{raw: r}}
}

func (n *HintEntry) Qualifier() *Identifier { return newIdentifier(must(n.raw.Qualifier())) }

func (n *HintEntry) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

// Value: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *HintEntry) Value() ExpressionNode { return wrapExpr(must(n.raw.Value())) }

// Limit wraps *googlesql.ASTLimit.
type Limit struct{ baseNode[*googlesql.ASTLimit] }

func newLimit(r *googlesql.ASTLimit) *Limit {
	if r == nil {
		return nil
	}
	return &Limit{baseNode[*googlesql.ASTLimit]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *Limit) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }

func (n *Limit) All() *LimitAll { return newLimitAll(must(n.raw.All())) }

// LimitAll wraps *googlesql.ASTLimitAll.
type LimitAll struct {
	baseNode[*googlesql.ASTLimitAll]
}

func newLimitAll(r *googlesql.ASTLimitAll) *LimitAll {
	if r == nil {
		return nil
	}
	return &LimitAll{baseNode[*googlesql.ASTLimitAll]{raw: r}}
}

// LimitOffset wraps *googlesql.ASTLimitOffset.
type LimitOffset struct {
	baseNode[*googlesql.ASTLimitOffset]
}

func newLimitOffset(r *googlesql.ASTLimitOffset) *LimitOffset {
	if r == nil {
		return nil
	}
	return &LimitOffset{baseNode[*googlesql.ASTLimitOffset]{raw: r}}
}

func (n *LimitOffset) Limit() *Limit { return newLimit(must(n.raw.Limit())) }

// Offset: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *LimitOffset) Offset() ExpressionNode { return wrapExpr(must(n.raw.Offset())) }

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

// ModelClause wraps *googlesql.ASTModelClause.
type ModelClause struct {
	baseNode[*googlesql.ASTModelClause]
}

func newModelClause(r *googlesql.ASTModelClause) *ModelClause {
	if r == nil {
		return nil
	}
	return &ModelClause{baseNode[*googlesql.ASTModelClause]{raw: r}}
}

func (n *ModelClause) ModelPath() *PathExpression {
	return newPathExpression(must(n.raw.ModelPath()))
}

// NullOrder wraps *googlesql.ASTNullOrder.
type NullOrder struct {
	baseNode[*googlesql.ASTNullOrder]
}

func newNullOrder(r *googlesql.ASTNullOrder) *NullOrder {
	if r == nil {
		return nil
	}
	return &NullOrder{baseNode[*googlesql.ASTNullOrder]{raw: r}}
}

func (n *NullOrder) NullsFirst() bool { return must(n.raw.NullsFirst()) }

// OnClause wraps *googlesql.ASTOnClause.
type OnClause struct {
	baseNode[*googlesql.ASTOnClause]
}

func newOnClause(r *googlesql.ASTOnClause) *OnClause {
	if r == nil {
		return nil
	}
	return &OnClause{baseNode[*googlesql.ASTOnClause]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *OnClause) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// OptionsEntry wraps *googlesql.ASTOptionsEntry.
type OptionsEntry struct {
	baseNode[*googlesql.ASTOptionsEntry]
}

func newOptionsEntry(r *googlesql.ASTOptionsEntry) *OptionsEntry {
	if r == nil {
		return nil
	}
	return &OptionsEntry{baseNode[*googlesql.ASTOptionsEntry]{raw: r}}
}

func (n *OptionsEntry) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

func (n *OptionsEntry) AssignmentOp() AssignmentOp { return must(n.raw.AssignmentOp()) }

// Value: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *OptionsEntry) Value() ExpressionNode { return wrapExpr(must(n.raw.Value())) }

// OptionsList wraps *googlesql.ASTOptionsList.
type OptionsList struct {
	baseNode[*googlesql.ASTOptionsList]
}

func newOptionsList(r *googlesql.ASTOptionsList) *OptionsList {
	if r == nil {
		return nil
	}
	return &OptionsList{baseNode[*googlesql.ASTOptionsList]{raw: r}}
}

// OptionsEntries: raw OptionsEntries(i) returns *OptionsEntry (concrete).
func (n *OptionsList) OptionsEntries() []*OptionsEntry {
	count := n.NumChildren()
	result := make([]*OptionsEntry, 0, count)
	for i := range count {
		e := must(n.raw.OptionsEntries(int32(i)))
		if e == nil {
			break
		}
		result = append(result, newOptionsEntry(e))
	}
	return result
}

// OrderBy wraps *googlesql.ASTOrderBy.
type OrderBy struct {
	baseNode[*googlesql.ASTOrderBy]
}

func newOrderBy(r *googlesql.ASTOrderBy) *OrderBy {
	if r == nil {
		return nil
	}
	return &OrderBy{baseNode[*googlesql.ASTOrderBy]{raw: r}}
}

func (n *OrderBy) Hint() *Hint { return newHint(must(n.raw.Hint())) }

// OrderingExpressions: raw OrderingExpressions(i) returns *OrderingExpression (concrete).
func (n *OrderBy) OrderingExpressions() []*OrderingExpression {
	var result []*OrderingExpression
	for _, c := range n.Children() {
		if e, ok := c.(*OrderingExpression); ok {
			result = append(result, e)
		}
	}
	return result
}

// OrderingExpression wraps *googlesql.ASTOrderingExpression.
type OrderingExpression struct {
	baseNode[*googlesql.ASTOrderingExpression]
}

func newOrderingExpression(r *googlesql.ASTOrderingExpression) *OrderingExpression {
	if r == nil {
		return nil
	}
	return &OrderingExpression{baseNode[*googlesql.ASTOrderingExpression]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *OrderingExpression) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *OrderingExpression) OrderingSpec() OrderingSpec {
	return must(n.raw.OrderingSpec())
}

func (n *OrderingExpression) Descending() bool { return must(n.raw.Descending()) }

func (n *OrderingExpression) NullOrder() *NullOrder {
	return newNullOrder(must(n.raw.NullOrder()))
}

func (n *OrderingExpression) Collate() *Collate { return newCollate(must(n.raw.Collate())) }

// PartitionBy wraps *googlesql.ASTPartitionBy.
type PartitionBy struct {
	baseNode[*googlesql.ASTPartitionBy]
}

func newPartitionBy(r *googlesql.ASTPartitionBy) *PartitionBy {
	if r == nil {
		return nil
	}
	return &PartitionBy{baseNode[*googlesql.ASTPartitionBy]{raw: r}}
}

func (n *PartitionBy) Hint() *Hint { return newHint(must(n.raw.Hint())) }

// PartitioningExpressions: raw returns ExpressionNode (interface) → []ExpressionNode.
func (n *PartitionBy) PartitioningExpressions() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		e := must(n.raw.PartitioningExpressions(int32(i)))
		if !defined(e) {
			break
		}
		result = append(result, wrapExpr(e))
	}
	return result
}

// PivotClause wraps *googlesql.ASTPivotClause.
type PivotClause struct {
	baseNode[*googlesql.ASTPivotClause]
}

func newPivotClause(r *googlesql.ASTPivotClause) *PivotClause {
	if r == nil {
		return nil
	}
	return &PivotClause{baseNode[*googlesql.ASTPivotClause]{raw: r}}
}

func (n *PivotClause) OutputAlias() *Alias { return newAlias(must(n.raw.OutputAlias())) }

func (n *PivotClause) ForExpression() ExpressionNode {
	return wrapExpr(must(n.raw.ForExpression()))
}

func (n *PivotClause) PivotExpressions() *PivotExpressionList {
	return newPivotExpressionList(must(n.raw.PivotExpressions()))
}

func (n *PivotClause) PivotValues() *PivotValueList {
	return newPivotValueList(must(n.raw.PivotValues()))
}

// PivotExpression wraps *googlesql.ASTPivotExpression.
type PivotExpression struct {
	baseNode[*googlesql.ASTPivotExpression]
}

func newPivotExpression(r *googlesql.ASTPivotExpression) *PivotExpression {
	if r == nil {
		return nil
	}
	return &PivotExpression{baseNode[*googlesql.ASTPivotExpression]{raw: r}}
}

func (n *PivotExpression) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *PivotExpression) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// PivotExpressionList wraps *googlesql.ASTPivotExpressionList.
type PivotExpressionList struct {
	baseNode[*googlesql.ASTPivotExpressionList]
}

func newPivotExpressionList(r *googlesql.ASTPivotExpressionList) *PivotExpressionList {
	if r == nil {
		return nil
	}
	return &PivotExpressionList{baseNode[*googlesql.ASTPivotExpressionList]{raw: r}}
}

// Expressions returns all pivot expressions.
func (n *PivotExpressionList) Expressions() []*PivotExpression {
	count := n.NumChildren()
	result := make([]*PivotExpression, 0, count)
	for i := range count {
		e := must(n.raw.Expressions(int32(i)))
		if e == nil {
			break
		}
		result = append(result, newPivotExpression(e))
	}
	return result
}

// PivotValue wraps *googlesql.ASTPivotValue.
type PivotValue struct {
	baseNode[*googlesql.ASTPivotValue]
}

func newPivotValue(r *googlesql.ASTPivotValue) *PivotValue {
	if r == nil {
		return nil
	}
	return &PivotValue{baseNode[*googlesql.ASTPivotValue]{raw: r}}
}

func (n *PivotValue) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *PivotValue) Value() ExpressionNode {
	return wrapExpr(must(n.raw.Value()))
}

// PivotValueList wraps *googlesql.ASTPivotValueList.
type PivotValueList struct {
	baseNode[*googlesql.ASTPivotValueList]
}

func newPivotValueList(r *googlesql.ASTPivotValueList) *PivotValueList {
	if r == nil {
		return nil
	}
	return &PivotValueList{baseNode[*googlesql.ASTPivotValueList]{raw: r}}
}

// Values returns all pivot values.
func (n *PivotValueList) Values() []*PivotValue {
	count := n.NumChildren()
	result := make([]*PivotValue, 0, count)
	for i := range count {
		e := must(n.raw.Values(int32(i)))
		if e == nil {
			break
		}
		result = append(result, newPivotValue(e))
	}
	return result
}

// RepeatableClause wraps *googlesql.ASTRepeatableClause.
type RepeatableClause struct {
	baseNode[*googlesql.ASTRepeatableClause]
}

func newRepeatableClause(r *googlesql.ASTRepeatableClause) *RepeatableClause {
	if r == nil {
		return nil
	}
	return &RepeatableClause{baseNode[*googlesql.ASTRepeatableClause]{raw: r}}
}

func (n *RepeatableClause) Argument() ExpressionNode {
	return wrapExpr(must(n.raw.Argument()))
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

// SampleClause wraps *googlesql.ASTSampleClause.
type SampleClause struct {
	baseNode[*googlesql.ASTSampleClause]
}

func newSampleClause(r *googlesql.ASTSampleClause) *SampleClause {
	if r == nil {
		return nil
	}
	return &SampleClause{baseNode[*googlesql.ASTSampleClause]{raw: r}}
}

func (n *SampleClause) SampleMethod() *Identifier {
	return newIdentifier(must(n.raw.SampleMethod()))
}

func (n *SampleClause) SampleSize() *SampleSize {
	return newSampleSize(must(n.raw.SampleSize()))
}

func (n *SampleClause) SampleSuffix() *SampleSuffix {
	return newSampleSuffix(must(n.raw.SampleSuffix()))
}

// SampleSize wraps *googlesql.ASTSampleSize.
type SampleSize struct {
	baseNode[*googlesql.ASTSampleSize]
}

func newSampleSize(r *googlesql.ASTSampleSize) *SampleSize {
	if r == nil {
		return nil
	}
	return &SampleSize{baseNode[*googlesql.ASTSampleSize]{raw: r}}
}

func (n *SampleSize) Size() ExpressionNode {
	return wrapExpr(must(n.raw.Size()))
}

func (n *SampleSize) Unit() SampleSizeUnit {
	return must(n.raw.Unit())
}

func (n *SampleSize) PartitionBy() *PartitionBy {
	return newPartitionBy(must(n.raw.PartitionBy()))
}

// SampleSuffix wraps *googlesql.ASTSampleSuffix.
type SampleSuffix struct {
	baseNode[*googlesql.ASTSampleSuffix]
}

func newSampleSuffix(r *googlesql.ASTSampleSuffix) *SampleSuffix {
	if r == nil {
		return nil
	}
	return &SampleSuffix{baseNode[*googlesql.ASTSampleSuffix]{raw: r}}
}

func (n *SampleSuffix) Weight() *WithWeight {
	return newWithWeight(must(n.raw.Weight()))
}

func (n *SampleSuffix) Repeat() *RepeatableClause {
	return newRepeatableClause(must(n.raw.Repeat()))
}

// UnpivotClause wraps *googlesql.ASTUnpivotClause.
type UnpivotClause struct {
	baseNode[*googlesql.ASTUnpivotClause]
}

func newUnpivotClause(r *googlesql.ASTUnpivotClause) *UnpivotClause {
	if r == nil {
		return nil
	}
	return &UnpivotClause{baseNode[*googlesql.ASTUnpivotClause]{raw: r}}
}

func (n *UnpivotClause) OutputAlias() *Alias {
	return newAlias(must(n.raw.OutputAlias()))
}

func (n *UnpivotClause) NullFilter() NullFilter {
	return must(n.raw.NullFilter())
}

func (n *UnpivotClause) UnpivotInItems() *UnpivotInItemList {
	return newUnpivotInItemList(must(n.raw.UnpivotInItems()))
}

func (n *UnpivotClause) UnpivotOutputNameColumn() *PathExpression {
	return newPathExpression(must(n.raw.UnpivotOutputNameColumn()))
}

func (n *UnpivotClause) UnpivotOutputValueColumns() *PathExpressionList {
	return newPathExpressionList(must(n.raw.UnpivotOutputValueColumns()))
}

// UnpivotInItem wraps *googlesql.ASTUnpivotInItem.
type UnpivotInItem struct {
	baseNode[*googlesql.ASTUnpivotInItem]
}

func newUnpivotInItem(r *googlesql.ASTUnpivotInItem) *UnpivotInItem {
	if r == nil {
		return nil
	}
	return &UnpivotInItem{baseNode[*googlesql.ASTUnpivotInItem]{raw: r}}
}

func (n *UnpivotInItem) Alias() *UnpivotInItemLabel {
	return newUnpivotInItemLabel(must(n.raw.Alias()))
}

func (n *UnpivotInItem) UnpivotColumns() *PathExpressionList {
	return newPathExpressionList(must(n.raw.UnpivotColumns()))
}

// UnpivotInItemLabel wraps *googlesql.ASTUnpivotInItemLabel.
type UnpivotInItemLabel struct {
	baseNode[*googlesql.ASTUnpivotInItemLabel]
}

func newUnpivotInItemLabel(r *googlesql.ASTUnpivotInItemLabel) *UnpivotInItemLabel {
	if r == nil {
		return nil
	}
	return &UnpivotInItemLabel{baseNode[*googlesql.ASTUnpivotInItemLabel]{raw: r}}
}

func (n *UnpivotInItemLabel) Label() ExpressionNode {
	return wrapExpr(must(n.raw.Label()))
}

// UnpivotInItemList wraps *googlesql.ASTUnpivotInItemList.
type UnpivotInItemList struct {
	baseNode[*googlesql.ASTUnpivotInItemList]
}

func newUnpivotInItemList(r *googlesql.ASTUnpivotInItemList) *UnpivotInItemList {
	if r == nil {
		return nil
	}
	return &UnpivotInItemList{baseNode[*googlesql.ASTUnpivotInItemList]{raw: r}}
}

// InItems returns all unpivot IN items.
func (n *UnpivotInItemList) InItems() []*UnpivotInItem {
	count := n.NumChildren()
	result := make([]*UnpivotInItem, 0, count)
	for i := range count {
		item := must(n.raw.InItems(int32(i)))
		if item == nil {
			break
		}
		result = append(result, newUnpivotInItem(item))
	}
	return result
}

// UntilClause wraps *googlesql.ASTUntilClause.
type UntilClause struct {
	baseNode[*googlesql.ASTUntilClause]
}

func newUntilClause(r *googlesql.ASTUntilClause) *UntilClause {
	if r == nil {
		return nil
	}
	return &UntilClause{baseNode[*googlesql.ASTUntilClause]{raw: r}}
}

// Condition: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *UntilClause) Condition() ExpressionNode {
	return wrapExpr(must(n.raw.Condition()))
}

// UsingClause wraps *googlesql.ASTUsingClause.
type UsingClause struct {
	baseNode[*googlesql.ASTUsingClause]
}

func newUsingClause(r *googlesql.ASTUsingClause) *UsingClause {
	if r == nil {
		return nil
	}
	return &UsingClause{baseNode[*googlesql.ASTUsingClause]{raw: r}}
}

// Keys: raw Keys(i) returns *Identifier (concrete) → []*Identifier.
func (n *UsingClause) Keys() []*Identifier {
	count := n.NumChildren()
	result := make([]*Identifier, 0, count)
	for i := range count {
		k := must(n.raw.Keys(int32(i)))
		if k == nil {
			break
		}
		result = append(result, newIdentifier(k))
	}
	return result
}

// WhereClause wraps *googlesql.ASTWhereClause.
type WhereClause struct {
	baseNode[*googlesql.ASTWhereClause]
}

func newWhereClause(r *googlesql.ASTWhereClause) *WhereClause {
	if r == nil {
		return nil
	}
	return &WhereClause{baseNode[*googlesql.ASTWhereClause]{raw: r}}
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *WhereClause) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// WindowClause wraps *googlesql.ASTWindowClause.
type WindowClause struct {
	baseNode[*googlesql.ASTWindowClause]
}

func newWindowClause(r *googlesql.ASTWindowClause) *WindowClause {
	if r == nil {
		return nil
	}
	return &WindowClause{baseNode[*googlesql.ASTWindowClause]{raw: r}}
}

// Windows: raw Windows(i) returns *WindowDefinition (concrete).
func (n *WindowClause) Windows() []*WindowDefinition {
	count := n.NumChildren()
	result := make([]*WindowDefinition, 0, count)
	for i := range count {
		w := must(n.raw.Windows(int32(i)))
		if w == nil {
			break
		}
		result = append(result, newWindowDefinition(w))
	}
	return result
}

// WindowDefinition wraps *googlesql.ASTWindowDefinition.
type WindowDefinition struct {
	baseNode[*googlesql.ASTWindowDefinition]
}

func newWindowDefinition(r *googlesql.ASTWindowDefinition) *WindowDefinition {
	if r == nil {
		return nil
	}
	return &WindowDefinition{baseNode[*googlesql.ASTWindowDefinition]{raw: r}}
}

func (n *WindowDefinition) WindowSpec() *WindowSpecification {
	return newWindowSpecification(must(n.raw.WindowSpec()))
}

func (n *WindowDefinition) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

// WindowFrame wraps *googlesql.ASTWindowFrame.
type WindowFrame struct {
	baseNode[*googlesql.ASTWindowFrame]
}

func newWindowFrame(r *googlesql.ASTWindowFrame) *WindowFrame {
	if r == nil {
		return nil
	}
	return &WindowFrame{baseNode[*googlesql.ASTWindowFrame]{raw: r}}
}

func (n *WindowFrame) FrameUnit() FrameUnit { return must(n.raw.FrameUnit()) }

func (n *WindowFrame) StartExpr() *WindowFrameExpr {
	return newWindowFrameExpr(must(n.raw.StartExpr()))
}

func (n *WindowFrame) EndExpr() *WindowFrameExpr {
	return newWindowFrameExpr(must(n.raw.EndExpr()))
}

// WindowFrameExpr wraps *googlesql.ASTWindowFrameExpr.
type WindowFrameExpr struct {
	baseNode[*googlesql.ASTWindowFrameExpr]
}

func newWindowFrameExpr(r *googlesql.ASTWindowFrameExpr) *WindowFrameExpr {
	if r == nil {
		return nil
	}
	return &WindowFrameExpr{baseNode[*googlesql.ASTWindowFrameExpr]{raw: r}}
}

func (n *WindowFrameExpr) BoundaryType() BoundaryType { return must(n.raw.BoundaryType()) }

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *WindowFrameExpr) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// WindowSpecification wraps *googlesql.ASTWindowSpecification.
type WindowSpecification struct {
	baseNode[*googlesql.ASTWindowSpecification]
}

func newWindowSpecification(r *googlesql.ASTWindowSpecification) *WindowSpecification {
	if r == nil {
		return nil
	}
	return &WindowSpecification{baseNode[*googlesql.ASTWindowSpecification]{raw: r}}
}

func (n *WindowSpecification) BaseWindowName() *Identifier {
	return newIdentifier(must(n.raw.BaseWindowName()))
}

func (n *WindowSpecification) PartitionBy() *PartitionBy {
	return newPartitionBy(must(n.raw.PartitionBy()))
}

func (n *WindowSpecification) OrderBy() *OrderBy { return newOrderBy(must(n.raw.OrderBy())) }

func (n *WindowSpecification) WindowFrame() *WindowFrame {
	return newWindowFrame(must(n.raw.WindowFrame()))
}

// WithClause wraps *googlesql.ASTWithClause.
type WithClause struct {
	baseNode[*googlesql.ASTWithClause]
}

func newWithClause(r *googlesql.ASTWithClause) *WithClause {
	if r == nil {
		return nil
	}
	return &WithClause{baseNode[*googlesql.ASTWithClause]{raw: r}}
}

func (n *WithClause) Recursive() bool { return must(n.raw.Recursive()) }

// Entries returns all WITH clause entries.
// Raw Entries(i) returns *WithClauseEntry (concrete) → []*WithClauseEntry.
func (n *WithClause) Entries() []*WithClauseEntry {
	count := n.NumChildren()
	result := make([]*WithClauseEntry, 0, count)
	for i := range count {
		e := must(n.raw.Entries(int32(i)))
		if e == nil {
			break
		}
		result = append(result, newWithClauseEntry(e))
	}
	return result
}

// WithConnectionClause wraps *googlesql.ASTWithConnectionClause.
type WithConnectionClause struct {
	baseNode[*googlesql.ASTWithConnectionClause]
}

func newWithConnectionClause(r *googlesql.ASTWithConnectionClause) *WithConnectionClause {
	if r == nil {
		return nil
	}
	return &WithConnectionClause{baseNode[*googlesql.ASTWithConnectionClause]{raw: r}}
}

func (n *WithConnectionClause) ConnectionClause() *ConnectionClause {
	return newConnectionClause(must(n.raw.ConnectionClause()))
}

// WithOffset wraps *googlesql.ASTWithOffset.
type WithOffset struct {
	baseNode[*googlesql.ASTWithOffset]
}

func newWithOffset(r *googlesql.ASTWithOffset) *WithOffset {
	if r == nil {
		return nil
	}
	return &WithOffset{baseNode[*googlesql.ASTWithOffset]{raw: r}}
}

func (n *WithOffset) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

// WithPartitionColumnsClause wraps *googlesql.ASTWithPartitionColumnsClause.
type WithPartitionColumnsClause struct {
	baseNode[*googlesql.ASTWithPartitionColumnsClause]
}

func newWithPartitionColumnsClause(r *googlesql.ASTWithPartitionColumnsClause) *WithPartitionColumnsClause {
	if r == nil {
		return nil
	}
	return &WithPartitionColumnsClause{baseNode[*googlesql.ASTWithPartitionColumnsClause]{raw: r}}
}

func (n *WithPartitionColumnsClause) TableElementList() *TableElementList {
	return newTableElementList(must(n.raw.TableElementList()))
}

// WithWeight wraps *googlesql.ASTWithWeight.
type WithWeight struct {
	baseNode[*googlesql.ASTWithWeight]
}

func newWithWeight(r *googlesql.ASTWithWeight) *WithWeight {
	if r == nil {
		return nil
	}
	return &WithWeight{baseNode[*googlesql.ASTWithWeight]{raw: r}}
}

func (n *WithWeight) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

// GroupingSet wraps *googlesql.ASTGroupingSet.
type GroupingSet struct {
	baseNode[*googlesql.ASTGroupingSet]
}

func newGroupingSet(r *googlesql.ASTGroupingSet) *GroupingSet {
	if r == nil {
		return nil
	}
	return &GroupingSet{baseNode[*googlesql.ASTGroupingSet]{raw: r}}
}
func (n *GroupingSet) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }
func (n *GroupingSet) Rollup() *Rollup            { return newRollup(must(n.raw.Rollup())) }
func (n *GroupingSet) Cube() *Cube                { return newCube(must(n.raw.Cube())) }

// GroupingSetList wraps *googlesql.ASTGroupingSetList.
type GroupingSetList struct {
	baseNode[*googlesql.ASTGroupingSetList]
}

func newGroupingSetList(r *googlesql.ASTGroupingSetList) *GroupingSetList {
	if r == nil {
		return nil
	}
	return &GroupingSetList{baseNode[*googlesql.ASTGroupingSetList]{raw: r}}
}

func (n *GroupingSetList) GroupingSets() []*GroupingSet {
	var result []*GroupingSet
	for item := range childrenOfType[*googlesql.ASTGroupingSet](n) {
		result = append(result, newGroupingSet(item))
	}
	return result
}

// GroupingItemOrder wraps *googlesql.ASTGroupingItemOrder.
type GroupingItemOrder struct {
	baseNode[*googlesql.ASTGroupingItemOrder]
}

func newGroupingItemOrder(r *googlesql.ASTGroupingItemOrder) *GroupingItemOrder {
	if r == nil {
		return nil
	}
	return &GroupingItemOrder{baseNode[*googlesql.ASTGroupingItemOrder]{raw: r}}
}
func (n *GroupingItemOrder) OrderingSpec() OrderingSpec { return must(n.raw.OrderingSpec()) }
func (n *GroupingItemOrder) NullOrder() *NullOrder      { return newNullOrder(must(n.raw.NullOrder())) }

// OnConflictClause wraps *googlesql.ASTOnConflictClause.
type OnConflictClause struct {
	baseNode[*googlesql.ASTOnConflictClause]
}

func newOnConflictClause(r *googlesql.ASTOnConflictClause) *OnConflictClause {
	if r == nil {
		return nil
	}
	return &OnConflictClause{baseNode[*googlesql.ASTOnConflictClause]{raw: r}}
}
func (n *OnConflictClause) ConflictAction() ConflictAction { return must(n.raw.ConflictAction()) }
func (n *OnConflictClause) ConflictTarget() *ColumnList {
	return newColumnList(must(n.raw.ConflictTarget()))
}

func (n *OnConflictClause) UpdateItemList() *UpdateItemList {
	return newUpdateItemList(must(n.raw.UpdateItemList()))
}

// OnOrUsingClauseList wraps *googlesql.ASTOnOrUsingClauseList.
type OnOrUsingClauseList struct {
	baseNode[*googlesql.ASTOnOrUsingClauseList]
}

func newOnOrUsingClauseList(r *googlesql.ASTOnOrUsingClauseList) *OnOrUsingClauseList {
	if r == nil {
		return nil
	}
	return &OnOrUsingClauseList{baseNode[*googlesql.ASTOnOrUsingClauseList]{raw: r}}
}

// WithReportModifier wraps *googlesql.ASTWithReportModifier.
type WithReportModifier struct {
	baseNode[*googlesql.ASTWithReportModifier]
}

func newWithReportModifier(r *googlesql.ASTWithReportModifier) *WithReportModifier {
	if r == nil {
		return nil
	}
	return &WithReportModifier{baseNode[*googlesql.ASTWithReportModifier]{raw: r}}
}

func (n *WithReportModifier) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// IntoAlias wraps *googlesql.ASTIntoAlias.
type IntoAlias struct {
	baseNode[*googlesql.ASTIntoAlias]
}

func newIntoAlias(r *googlesql.ASTIntoAlias) *IntoAlias {
	if r == nil {
		return nil
	}
	return &IntoAlias{baseNode[*googlesql.ASTIntoAlias]{raw: r}}
}
func (n *IntoAlias) Identifier() *Identifier { return newIdentifier(must(n.raw.Identifier())) }

// CloneDataSourceList wraps *googlesql.ASTCloneDataSourceList.
type CloneDataSourceList struct {
	baseNode[*googlesql.ASTCloneDataSourceList]
}

func newCloneDataSourceList(r *googlesql.ASTCloneDataSourceList) *CloneDataSourceList {
	if r == nil {
		return nil
	}
	return &CloneDataSourceList{baseNode[*googlesql.ASTCloneDataSourceList]{raw: r}}
}

// TransactionModeList wraps *googlesql.ASTTransactionModeList.
type TransactionModeList struct {
	baseNode[*googlesql.ASTTransactionModeList]
}

func newTransactionModeList(r *googlesql.ASTTransactionModeList) *TransactionModeList {
	if r == nil {
		return nil
	}
	return &TransactionModeList{baseNode[*googlesql.ASTTransactionModeList]{raw: r}}
}

// TransactionIsolationLevel wraps *googlesql.ASTTransactionIsolationLevel.
type TransactionIsolationLevel struct {
	baseNode[*googlesql.ASTTransactionIsolationLevel]
}

func newTransactionIsolationLevel(r *googlesql.ASTTransactionIsolationLevel) *TransactionIsolationLevel {
	if r == nil {
		return nil
	}
	return &TransactionIsolationLevel{baseNode[*googlesql.ASTTransactionIsolationLevel]{raw: r}}
}

// TransactionReadWriteMode wraps *googlesql.ASTTransactionReadWriteMode.
type TransactionReadWriteMode struct {
	baseNode[*googlesql.ASTTransactionReadWriteMode]
}

func newTransactionReadWriteMode(r *googlesql.ASTTransactionReadWriteMode) *TransactionReadWriteMode {
	if r == nil {
		return nil
	}
	return &TransactionReadWriteMode{baseNode[*googlesql.ASTTransactionReadWriteMode]{raw: r}}
}
