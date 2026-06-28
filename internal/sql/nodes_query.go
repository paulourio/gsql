package sql

import (
	"github.com/goccy/go-googlesql"
)

// ─── Query / Script ───────────────────────────────────────────────────────────

// QueryStatement wraps *googlesql.ASTQueryStatement.
type QueryStatement struct {
	baseNode[*googlesql.ASTQueryStatement]
}

func newQueryStatement(r *googlesql.ASTQueryStatement) *QueryStatement {
	if r == nil {
		return nil
	}
	return &QueryStatement{baseNode[*googlesql.ASTQueryStatement]{raw: r}}
}
func (n *QueryStatement) isStatement() {}
func (n *QueryStatement) Query() *Query {
	return newQuery(must(n.raw.Query()))
}

// Query wraps *googlesql.ASTQuery.
type Query struct{ baseNode[*googlesql.ASTQuery] }

func newQuery(r *googlesql.ASTQuery) *Query {
	if r == nil {
		return nil
	}
	return &Query{baseNode[*googlesql.ASTQuery]{raw: r}}
}
func (n *Query) isQueryExpression() {}
func (n *Query) WithClause() *WithClause {
	return newWithClause(must(n.raw.WithClause()))
}

func (n *Query) QueryExpr() QueryExpressionNode {
	return wrapQueryExpr(must(n.raw.QueryExpr()))
}
func (n *Query) OrderBy() *OrderBy { return newOrderBy(must(n.raw.OrderBy())) }
func (n *Query) LimitOffset() *LimitOffset {
	return newLimitOffset(must(n.raw.LimitOffset()))
}
func (n *Query) LockMode() *LockMode { return newLockMode(must(n.raw.LockMode())) }
func (n *Query) IsNested() bool      { return must(n.raw.IsNested()) }

// PipeOperators returns all pipe operators.
func (n *Query) PipeOperators() []PipeOperatorNode {
	var result []PipeOperatorNode
	for _, c := range n.Children() {
		if op, ok := c.(PipeOperatorNode); ok {
			result = append(result, op)
		}
	}
	return result
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

// ─── SELECT ───────────────────────────────────────────────────────────────────

// Select wraps *googlesql.ASTSelect.
type Select struct{ baseNode[*googlesql.ASTSelect] }

func newSelect(r *googlesql.ASTSelect) *Select {
	if r == nil {
		return nil
	}
	return &Select{baseNode[*googlesql.ASTSelect]{raw: r}}
}
func (n *Select) isQueryExpression() {}
func (n *Select) Distinct() bool     { return must(n.raw.Distinct()) }
func (n *Select) Hint() *Hint        { return newHint(must(n.raw.Hint())) }
func (n *Select) SelectAs() *SelectAs {
	return newSelectAs(must(n.raw.SelectAs()))
}

func (n *Select) SelectList() *SelectList {
	return newSelectList(must(n.raw.SelectList()))
}

func (n *Select) FromClause() *FromClause {
	return newFromClause(must(n.raw.FromClause()))
}

func (n *Select) WhereClause() *WhereClause {
	return newWhereClause(must(n.raw.WhereClause()))
}
func (n *Select) GroupBy() *GroupBy { return newGroupBy(must(n.raw.GroupBy())) }
func (n *Select) Having() *Having   { return newHaving(must(n.raw.Having())) }
func (n *Select) Qualify() *Qualify { return newQualify(must(n.raw.Qualify())) }
func (n *Select) WindowClause() *WindowClause {
	return newWindowClause(must(n.raw.WindowClause()))
}

func (n *Select) WithModifier() *WithModifier {
	return newWithModifier(must(n.raw.WithModifier()))
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

// ─── Set operations ───────────────────────────────────────────────────────────

// SetOperation wraps *googlesql.ASTSetOperation.
type SetOperation struct {
	baseNode[*googlesql.ASTSetOperation]
}

func newSetOperation(r *googlesql.ASTSetOperation) *SetOperation {
	if r == nil {
		return nil
	}
	return &SetOperation{baseNode[*googlesql.ASTSetOperation]{raw: r}}
}
func (n *SetOperation) isQueryExpression() {}

func (n *SetOperation) Inputs() []QueryExpressionNode {
	var result []QueryExpressionNode
	for _, c := range n.Children() {
		if qe, ok := c.(QueryExpressionNode); ok {
			result = append(result, qe)
		}
	}
	return result
}

func (n *SetOperation) Metadata() *SetOperationMetadataList {
	return newSetOperationMetadataList(must(n.raw.Metadata()))
}

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

// SetOperationType wraps *googlesql.ASTSetOperationType.
type SetOperationType struct {
	baseNode[*googlesql.ASTSetOperationType]
}

func newSetOperationType(r *googlesql.ASTSetOperationType) *SetOperationType {
	if r == nil {
		return nil
	}
	return &SetOperationType{baseNode[*googlesql.ASTSetOperationType]{raw: r}}
}
func (n *SetOperationType) Value() SetOp { return must(n.raw.Value()) }

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

// ─── WITH ─────────────────────────────────────────────────────────────────────

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
func (n *AliasedQuery) Query() *Query      { return newQuery(must(n.raw.Query())) }

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

// WithExpression wraps *googlesql.ASTWithExpression.
type WithExpression struct {
	baseNode[*googlesql.ASTWithExpression]
}

func newWithExpression(r *googlesql.ASTWithExpression) *WithExpression {
	if r == nil {
		return nil
	}
	return &WithExpression{baseNode[*googlesql.ASTWithExpression]{raw: r}}
}
func (n *WithExpression) isExpression() {}
func (n *WithExpression) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *WithExpression) Variables() *SelectList {
	return newSelectList(must(n.raw.Variables()))
}

// ─── FROM / JOIN ──────────────────────────────────────────────────────────────

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

// Join wraps *googlesql.ASTJoin.
type Join struct{ baseNode[*googlesql.ASTJoin] }

func newJoin(r *googlesql.ASTJoin) *Join {
	if r == nil {
		return nil
	}
	return &Join{baseNode[*googlesql.ASTJoin]{raw: r}}
}
func (n *Join) isTableExpression() {}

// LHS and RHS: raw returns TableExpressionNode (interface) → TableExpressionNode.
func (n *Join) LHS() TableExpressionNode { return wrapTableExpr(must(n.raw.Lhs())) }
func (n *Join) RHS() TableExpressionNode { return wrapTableExpr(must(n.raw.Rhs())) }
func (n *Join) JoinType() JoinType       { return must(n.raw.JoinType()) }
func (n *Join) JoinHint() JoinHint       { return must(n.raw.JoinHint()) }
func (n *Join) Natural() bool            { return must(n.raw.Natural()) }
func (n *Join) Hint() *Hint              { return newHint(must(n.raw.Hint())) }
func (n *Join) OnClause() *OnClause      { return newOnClause(must(n.raw.OnClause())) }
func (n *Join) UsingClause() *UsingClause {
	return newUsingClause(must(n.raw.UsingClause()))
}

func (n *Join) JoinLocation() Node {
	return Wrap(must(n.raw.JoinLocation()))
}

// ParenthesizedJoin wraps *googlesql.ASTParenthesizedJoin.
type ParenthesizedJoin struct {
	baseNode[*googlesql.ASTParenthesizedJoin]
}

func newParenthesizedJoin(r *googlesql.ASTParenthesizedJoin) *ParenthesizedJoin {
	if r == nil {
		return nil
	}
	return &ParenthesizedJoin{baseNode[*googlesql.ASTParenthesizedJoin]{raw: r}}
}
func (n *ParenthesizedJoin) isTableExpression() {}
func (n *ParenthesizedJoin) Join() *Join {
	return newJoin(must(n.raw.Join()))
}

func (n *ParenthesizedJoin) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

// TablePathExpression wraps *googlesql.ASTTablePathExpression.
type TablePathExpression struct {
	baseNode[*googlesql.ASTTablePathExpression]
}

func newTablePathExpression(r *googlesql.ASTTablePathExpression) *TablePathExpression {
	if r == nil {
		return nil
	}
	return &TablePathExpression{baseNode[*googlesql.ASTTablePathExpression]{raw: r}}
}
func (n *TablePathExpression) isTableExpression() {}
func (n *TablePathExpression) PathExpr() *PathExpression {
	return newPathExpression(must(n.raw.PathExpr()))
}
func (n *TablePathExpression) Alias() *Alias { return newAlias(must(n.raw.Alias())) }
func (n *TablePathExpression) Hint() *Hint   { return newHint(must(n.raw.Hint())) }
func (n *TablePathExpression) WithOffset() *WithOffset {
	return newWithOffset(must(n.raw.WithOffset()))
}

func (n *TablePathExpression) ForSystemTime() *ForSystemTime {
	return newForSystemTime(must(n.raw.ForSystemTime()))
}

func (n *TablePathExpression) PivotClause() *PivotClause {
	return newPivotClause(must(n.raw.PivotClause()))
}

func (n *TablePathExpression) UnpivotClause() *UnpivotClause {
	return newUnpivotClause(must(n.raw.UnpivotClause()))
}

func (n *TablePathExpression) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

func (n *TablePathExpression) UnnestExpr() *UnnestExpression {
	return newUnnestExpression(must(n.raw.UnnestExpr()))
}

// TableSubquery wraps *googlesql.ASTTableSubquery.
type TableSubquery struct {
	baseNode[*googlesql.ASTTableSubquery]
}

func newTableSubquery(r *googlesql.ASTTableSubquery) *TableSubquery {
	if r == nil {
		return nil
	}
	return &TableSubquery{baseNode[*googlesql.ASTTableSubquery]{raw: r}}
}
func (n *TableSubquery) isTableExpression() {}
func (n *TableSubquery) Subquery() *Query   { return newQuery(must(n.raw.Subquery())) }
func (n *TableSubquery) Alias() *Alias      { return newAlias(must(n.raw.Alias())) }
func (n *TableSubquery) PivotClause() *PivotClause {
	return newPivotClause(must(n.raw.PivotClause()))
}

func (n *TableSubquery) UnpivotClause() *UnpivotClause {
	return newUnpivotClause(must(n.raw.UnpivotClause()))
}

func (n *TableSubquery) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

// TVF wraps *googlesql.ASTTVF.
type TVF struct {
	baseNode[*googlesql.ASTTVF]
}

func newTVF(r *googlesql.ASTTVF) *TVF {
	if r == nil {
		return nil
	}
	return &TVF{baseNode[*googlesql.ASTTVF]{raw: r}}
}

func (n *TVF) isTableExpression() {}

func (n *TVF) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *TVF) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *TVF) Alias() *Alias {
	return newAlias(must(n.raw.Alias()))
}

func (n *TVF) PivotClause() *PivotClause {
	return newPivotClause(must(n.raw.PivotClause()))
}

func (n *TVF) UnpivotClause() *UnpivotClause {
	return newUnpivotClause(must(n.raw.UnpivotClause()))
}

func (n *TVF) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

// ArgumentEntries returns all argument entries.
func (n *TVF) ArgumentEntries() []*TVFArgument {
	var result []*TVFArgument
	for _, c := range n.Children() {
		if a, ok := c.(*TVFArgument); ok {
			result = append(result, a)
		}
	}
	return result
}

// TableClause wraps *googlesql.ASTTableClause.
type TableClause struct {
	baseNode[*googlesql.ASTTableClause]
}

func newTableClause(r *googlesql.ASTTableClause) *TableClause {
	if r == nil {
		return nil
	}
	return &TableClause{baseNode[*googlesql.ASTTableClause]{raw: r}}
}
func (n *TableClause) isTableExpression() {}
func (n *TableClause) TablePath() *PathExpression {
	return newPathExpression(must(n.raw.TablePath()))
}

func (n *TableClause) Tvf() Node {
	return Wrap(must(n.raw.Tvf()))
}

// UnnestExpression wraps *googlesql.ASTUnnestExpression.
type UnnestExpression struct {
	baseNode[*googlesql.ASTUnnestExpression]
}

func newUnnestExpression(r *googlesql.ASTUnnestExpression) *UnnestExpression {
	if r == nil {
		return nil
	}
	return &UnnestExpression{baseNode[*googlesql.ASTUnnestExpression]{raw: r}}
}
func (n *UnnestExpression) isTableExpression() {}
func (n *UnnestExpression) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// Expressions returns all expressions.
func (n *UnnestExpression) Expressions() []*ExpressionWithOptAlias {
	var result []*ExpressionWithOptAlias
	for _, c := range n.Children() {
		if e, ok := c.(*ExpressionWithOptAlias); ok {
			result = append(result, e)
		}
	}
	return result
}

// UnnestExpressionWithOptAliasAndOffset wraps
// *googlesql.ASTUnnestExpressionWithOptAliasAndOffset.
type UnnestExpressionWithOptAliasAndOffset struct {
	baseNode[*googlesql.ASTUnnestExpressionWithOptAliasAndOffset]
}

func newUnnestExpressionWithOptAliasAndOffset(r *googlesql.ASTUnnestExpressionWithOptAliasAndOffset) *UnnestExpressionWithOptAliasAndOffset {
	if r == nil {
		return nil
	}
	return &UnnestExpressionWithOptAliasAndOffset{baseNode[*googlesql.ASTUnnestExpressionWithOptAliasAndOffset]{raw: r}}
}
func (n *UnnestExpressionWithOptAliasAndOffset) isTableExpression() {}
func (n *UnnestExpressionWithOptAliasAndOffset) UnnestExpression() *UnnestExpression {
	return newUnnestExpression(must(n.raw.UnnestExpression()))
}

func (n *UnnestExpressionWithOptAliasAndOffset) OptionalAlias() *Alias {
	return newAlias(must(n.raw.OptionalAlias()))
}

func (n *UnnestExpressionWithOptAliasAndOffset) OptionalWithOffset() *WithOffset {
	return newWithOffset(must(n.raw.OptionalWithOffset()))
}

// ─── WHERE / GROUP BY / HAVING / QUALIFY / ORDER BY ──────────────────────────

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
func (n *GroupBy) Hint() *Hint      { return newHint(must(n.raw.Hint())) }
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
func (n *GroupingItem) Alias() *Alias   { return newAlias(must(n.raw.Alias())) }
func (n *GroupingItem) Rollup() *Rollup { return newRollup(must(n.raw.Rollup())) }
func (n *GroupingItem) Cube() *Cube     { return newCube(must(n.raw.Cube())) }

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
func (n *Limit) All() *LimitAll             { return newLimitAll(must(n.raw.All())) }

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

// ─── WINDOW ───────────────────────────────────────────────────────────────────

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

// ─── Partition / Cluster ──────────────────────────────────────────────────────

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

// ─── HINT ─────────────────────────────────────────────────────────────────────

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
func (n *HintEntry) Name() *Identifier      { return newIdentifier(must(n.raw.Name())) }

// Value: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *HintEntry) Value() ExpressionNode { return wrapExpr(must(n.raw.Value())) }

// HintedStatement wraps *googlesql.ASTHintedStatement.
type HintedStatement struct {
	baseNode[*googlesql.ASTHintedStatement]
}

func newHintedStatement(r *googlesql.ASTHintedStatement) *HintedStatement {
	if r == nil {
		return nil
	}
	return &HintedStatement{baseNode[*googlesql.ASTHintedStatement]{raw: r}}
}
func (n *HintedStatement) isStatement() {}
func (n *HintedStatement) Hint() *Hint  { return newHint(must(n.raw.Hint())) }
func (n *HintedStatement) Statement() StatementNode {
	return wrapStmt(must(n.raw.Statement()))
}

// ─── ON / USING clauses ────────────────────────────────────────────────────────

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

// ─── HAVING MODIFIER ──────────────────────────────────────────────────────────

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
func (n *ClampedBetweenModifier) Low() ExpressionNode  { return wrapExpr(must(n.raw.Low())) }
func (n *ClampedBetweenModifier) High() ExpressionNode { return wrapExpr(must(n.raw.High())) }

// ─── Misc clause nodes ────────────────────────────────────────────────────────

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

// Alias wraps *googlesql.ASTAlias.
type Alias struct{ baseNode[*googlesql.ASTAlias] }

func newAlias(r *googlesql.ASTAlias) *Alias {
	if r == nil {
		return nil
	}
	return &Alias{baseNode[*googlesql.ASTAlias]{raw: r}}
}
func (n *Alias) Identifier() *Identifier { return newIdentifier(must(n.raw.Identifier())) }

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

// ─── Identifiers / Path expressions ──────────────────────────────────────────

// Identifier wraps *googlesql.ASTIdentifier.
type Identifier struct {
	baseNode[*googlesql.ASTIdentifier]
}

func newIdentifier(r *googlesql.ASTIdentifier) *Identifier {
	if r == nil {
		return nil
	}
	return &Identifier{baseNode[*googlesql.ASTIdentifier]{raw: r}}
}
func (n *Identifier) isExpression()       {}
func (n *Identifier) GetAsString() string { return must(n.raw.GetAsString()) }
func (n *Identifier) IsQuoted() bool      { return must(n.raw.IsQuoted()) }

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

// PathExpression wraps *googlesql.ASTPathExpression.
type PathExpression struct {
	baseNode[*googlesql.ASTPathExpression]
}

func newPathExpression(r *googlesql.ASTPathExpression) *PathExpression {
	if r == nil {
		return nil
	}
	return &PathExpression{baseNode[*googlesql.ASTPathExpression]{raw: r}}
}
func (n *PathExpression) isExpression() {}
func (n *PathExpression) NumNames() int { return int(must(n.raw.NumNames())) }
func (n *PathExpression) FirstName() *Identifier {
	return newIdentifier(must(n.raw.FirstName()))
}

func (n *PathExpression) LastName() *Identifier {
	return newIdentifier(must(n.raw.LastName()))
}

// Names returns all name components.  Raw Name(i) returns *Identifier (concrete).
func (n *PathExpression) Names() []*Identifier {
	count := n.NumNames()
	result := make([]*Identifier, 0, count)
	for i := range count {
		name := must(n.raw.Name(int32(i)))
		if name == nil {
			break
		}
		result = append(result, newIdentifier(name))
	}
	return result
}

func (n *PathExpression) ToIdentifierVector() []string {
	return must(n.raw.ToIdentifierVector())
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

// ─── Star nodes ───────────────────────────────────────────────────────────────

// Star wraps *googlesql.ASTStar.
type Star struct{ baseNode[*googlesql.ASTStar] }

func newStar(r *googlesql.ASTStar) *Star {
	if r == nil {
		return nil
	}
	return &Star{baseNode[*googlesql.ASTStar]{raw: r}}
}
func (n *Star) isExpression() {}
func (n *Star) isLeaf()       {}
func (n *Star) Image() string { return must(n.raw.Image()) }

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

// StarWithModifiers wraps *googlesql.ASTStarWithModifiers.
type StarWithModifiers struct {
	baseNode[*googlesql.ASTStarWithModifiers]
}

func newStarWithModifiers(r *googlesql.ASTStarWithModifiers) *StarWithModifiers {
	if r == nil {
		return nil
	}
	return &StarWithModifiers{baseNode[*googlesql.ASTStarWithModifiers]{raw: r}}
}
func (n *StarWithModifiers) isExpression() {}
func (n *StarWithModifiers) Modifiers() *StarModifiers {
	return newStarModifiers(must(n.raw.Modifiers()))
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

func (n *StarReplaceItem) Alias() *Identifier {
	return newIdentifier(must(n.raw.Alias()))
}

// Expression: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *StarReplaceItem) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// ─── Expression nodes ─────────────────────────────────────────────────────────

// AndExpr wraps *googlesql.ASTAndExpr.
type AndExpr struct {
	baseNode[*googlesql.ASTAndExpr]
}

func newAndExpr(r *googlesql.ASTAndExpr) *AndExpr {
	if r == nil {
		return nil
	}
	return &AndExpr{baseNode[*googlesql.ASTAndExpr]{raw: r}}
}
func (n *AndExpr) isExpression() {}

// Conjuncts: raw Conjuncts(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *AndExpr) Conjuncts() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		c := must(n.raw.Conjuncts(int32(i)))
		if !defined(c) {
			break
		}
		result = append(result, wrapExpr(c))
	}
	return result
}

// OrExpr wraps *googlesql.ASTOrExpr.
type OrExpr struct{ baseNode[*googlesql.ASTOrExpr] }

func newOrExpr(r *googlesql.ASTOrExpr) *OrExpr {
	if r == nil {
		return nil
	}
	return &OrExpr{baseNode[*googlesql.ASTOrExpr]{raw: r}}
}
func (n *OrExpr) isExpression() {}

// Disjuncts: raw Disjuncts(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *OrExpr) Disjuncts() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		d := must(n.raw.Disjuncts(int32(i)))
		if !defined(d) {
			break
		}
		result = append(result, wrapExpr(d))
	}
	return result
}

// BinaryExpression wraps *googlesql.ASTBinaryExpression.
type BinaryExpression struct {
	baseNode[*googlesql.ASTBinaryExpression]
}

func newBinaryExpression(r *googlesql.ASTBinaryExpression) *BinaryExpression {
	if r == nil {
		return nil
	}
	return &BinaryExpression{baseNode[*googlesql.ASTBinaryExpression]{raw: r}}
}
func (n *BinaryExpression) isExpression() {}
func (n *BinaryExpression) Op() BinaryOp  { return must(n.raw.Op()) }
func (n *BinaryExpression) IsNot() bool   { return must(n.raw.IsNot()) }

// LHS/RHS: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *BinaryExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }
func (n *BinaryExpression) RHS() ExpressionNode { return wrapExpr(must(n.raw.Rhs())) }

// BitwiseShiftExpression wraps *googlesql.ASTBitwiseShiftExpression.
type BitwiseShiftExpression struct {
	baseNode[*googlesql.ASTBitwiseShiftExpression]
}

func newBitwiseShiftExpression(r *googlesql.ASTBitwiseShiftExpression) *BitwiseShiftExpression {
	if r == nil {
		return nil
	}
	return &BitwiseShiftExpression{baseNode[*googlesql.ASTBitwiseShiftExpression]{raw: r}}
}
func (n *BitwiseShiftExpression) isExpression()     {}
func (n *BitwiseShiftExpression) IsLeftShift() bool { return must(n.raw.IsLeftShift()) }

// LHS/RHS: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *BitwiseShiftExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }
func (n *BitwiseShiftExpression) RHS() ExpressionNode { return wrapExpr(must(n.raw.Rhs())) }

// UnaryExpression wraps *googlesql.ASTUnaryExpression.
type UnaryExpression struct {
	baseNode[*googlesql.ASTUnaryExpression]
}

func newUnaryExpression(r *googlesql.ASTUnaryExpression) *UnaryExpression {
	if r == nil {
		return nil
	}
	return &UnaryExpression{baseNode[*googlesql.ASTUnaryExpression]{raw: r}}
}
func (n *UnaryExpression) isExpression() {}
func (n *UnaryExpression) Op() UnaryOp   { return must(n.raw.Op()) }

// Operand: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *UnaryExpression) Operand() ExpressionNode { return wrapExpr(must(n.raw.Operand())) }

// BetweenExpression wraps *googlesql.ASTBetweenExpression.
type BetweenExpression struct {
	baseNode[*googlesql.ASTBetweenExpression]
}

func newBetweenExpression(r *googlesql.ASTBetweenExpression) *BetweenExpression {
	if r == nil {
		return nil
	}
	return &BetweenExpression{baseNode[*googlesql.ASTBetweenExpression]{raw: r}}
}
func (n *BetweenExpression) isExpression() {}
func (n *BetweenExpression) IsNot() bool   { return must(n.raw.IsNot()) }

// LHS/Low/High: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *BetweenExpression) LHS() ExpressionNode  { return wrapExpr(must(n.raw.Lhs())) }
func (n *BetweenExpression) Low() ExpressionNode  { return wrapExpr(must(n.raw.Low())) }
func (n *BetweenExpression) High() ExpressionNode { return wrapExpr(must(n.raw.High())) }

// CaseNoValueExpression wraps *googlesql.ASTCaseNoValueExpression.
type CaseNoValueExpression struct {
	baseNode[*googlesql.ASTCaseNoValueExpression]
}

func newCaseNoValueExpression(r *googlesql.ASTCaseNoValueExpression) *CaseNoValueExpression {
	if r == nil {
		return nil
	}
	return &CaseNoValueExpression{baseNode[*googlesql.ASTCaseNoValueExpression]{raw: r}}
}
func (n *CaseNoValueExpression) isExpression() {}

// Arguments: raw Arguments(i) returns ExpressionNode (interface) → []ExpressionNode.
// Collects all WHEN/THEN/ELSE args by index up to NumChildren().
func (n *CaseNoValueExpression) Arguments() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		arg := must(n.raw.Arguments(int32(i)))
		if !defined(arg) {
			break
		}
		result = append(result, wrapExpr(arg))
	}
	return result
}

// CaseValueExpression wraps *googlesql.ASTCaseValueExpression.
type CaseValueExpression struct {
	baseNode[*googlesql.ASTCaseValueExpression]
}

func newCaseValueExpression(r *googlesql.ASTCaseValueExpression) *CaseValueExpression {
	if r == nil {
		return nil
	}
	return &CaseValueExpression{baseNode[*googlesql.ASTCaseValueExpression]{raw: r}}
}
func (n *CaseValueExpression) isExpression() {}

// Arguments: raw Arguments(i) returns ExpressionNode (interface) → []ExpressionNode.
// First element is the CASE value expression; remaining are WHEN/THEN/ELSE args.
func (n *CaseValueExpression) Arguments() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		arg := must(n.raw.Arguments(int32(i)))
		if !defined(arg) {
			break
		}
		result = append(result, wrapExpr(arg))
	}
	return result
}

// CastExpression wraps *googlesql.ASTCastExpression.
type CastExpression struct {
	baseNode[*googlesql.ASTCastExpression]
}

func newCastExpression(r *googlesql.ASTCastExpression) *CastExpression {
	if r == nil {
		return nil
	}
	return &CastExpression{baseNode[*googlesql.ASTCastExpression]{raw: r}}
}
func (n *CastExpression) isExpression()    {}
func (n *CastExpression) IsSafeCast() bool { return must(n.raw.IsSafeCast()) }

// Expr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *CastExpression) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// Type: raw returns TypeNode (interface) → TypeNode.
func (n *CastExpression) Type() TypeNode { return wrapType(must(n.raw.Type())) }

func (n *CastExpression) Format() *FormatClause {
	return newFormatClause(must(n.raw.Format()))
}

// InExpression wraps *googlesql.ASTInExpression.
type InExpression struct {
	baseNode[*googlesql.ASTInExpression]
}

func newInExpression(r *googlesql.ASTInExpression) *InExpression {
	if r == nil {
		return nil
	}
	return &InExpression{baseNode[*googlesql.ASTInExpression]{raw: r}}
}
func (n *InExpression) isExpression() {}
func (n *InExpression) IsNot() bool   { return must(n.raw.IsNot()) }
func (n *InExpression) Hint() *Hint   { return newHint(must(n.raw.Hint())) }

// LHS: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *InExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }
func (n *InExpression) InList() *InList     { return newInList(must(n.raw.InList())) }
func (n *InExpression) Query() *Query       { return newQuery(must(n.raw.Query())) }
func (n *InExpression) UnnestExpr() *UnnestExpression {
	return newUnnestExpression(must(n.raw.UnnestExpr()))
}

func (n *InExpression) InLocation() Node {
	return Wrap(must(n.raw.InLocation()))
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

// ExpressionSubquery wraps *googlesql.ASTExpressionSubquery.
type ExpressionSubquery struct {
	baseNode[*googlesql.ASTExpressionSubquery]
}

func newExpressionSubquery(r *googlesql.ASTExpressionSubquery) *ExpressionSubquery {
	if r == nil {
		return nil
	}
	return &ExpressionSubquery{baseNode[*googlesql.ASTExpressionSubquery]{raw: r}}
}
func (n *ExpressionSubquery) isExpression() {}
func (n *ExpressionSubquery) Modifier() SubqueryModifier {
	return must(n.raw.Modifier())
}
func (n *ExpressionSubquery) Query() *Query { return newQuery(must(n.raw.Query())) }
func (n *ExpressionSubquery) Hint() *Hint   { return newHint(must(n.raw.Hint())) }

// ExtractExpression wraps *googlesql.ASTExtractExpression.
type ExtractExpression struct {
	baseNode[*googlesql.ASTExtractExpression]
}

func newExtractExpression(r *googlesql.ASTExtractExpression) *ExtractExpression {
	if r == nil {
		return nil
	}
	return &ExtractExpression{baseNode[*googlesql.ASTExtractExpression]{raw: r}}
}
func (n *ExtractExpression) isExpression() {}

// LHSExpr/RHSExpr/TimeZoneExpr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *ExtractExpression) LHSExpr() ExpressionNode { return wrapExpr(must(n.raw.LhsExpr())) }
func (n *ExtractExpression) RHSExpr() ExpressionNode { return wrapExpr(must(n.raw.RhsExpr())) }
func (n *ExtractExpression) TimeZoneExpr() ExpressionNode {
	return wrapExpr(must(n.raw.TimeZoneExpr()))
}

// IntervalExpr wraps *googlesql.ASTIntervalExpr.
type IntervalExpr struct {
	baseNode[*googlesql.ASTIntervalExpr]
}

func newIntervalExpr(r *googlesql.ASTIntervalExpr) *IntervalExpr {
	if r == nil {
		return nil
	}
	return &IntervalExpr{baseNode[*googlesql.ASTIntervalExpr]{raw: r}}
}
func (n *IntervalExpr) isExpression() {}

// IntervalValue: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *IntervalExpr) IntervalValue() ExpressionNode {
	return wrapExpr(must(n.raw.IntervalValue()))
}

func (n *IntervalExpr) DatePartName() *Identifier {
	return newIdentifier(must(n.raw.DatePartName()))
}

func (n *IntervalExpr) DatePartNameTo() *Identifier {
	return newIdentifier(must(n.raw.DatePartNameTo()))
}

// ConcatExpr wraps *googlesql.ASTConcatExpr.
type ConcatExpr struct {
	baseNode[*googlesql.ASTConcatExpr]
}

func newConcatExpr(r *googlesql.ASTConcatExpr) *ConcatExpr {
	if r == nil {
		return nil
	}
	return &ConcatExpr{baseNode[*googlesql.ASTConcatExpr]{raw: r}}
}
func (n *ConcatExpr) isExpression() {}

// Operands: raw Operands(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *ConcatExpr) Operands() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		op := must(n.raw.Operands(int32(i)))
		if !defined(op) {
			break
		}
		result = append(result, wrapExpr(op))
	}
	return result
}

// ArrayConstructor wraps *googlesql.ASTArrayConstructor.
type ArrayConstructor struct {
	baseNode[*googlesql.ASTArrayConstructor]
}

func newArrayConstructor(r *googlesql.ASTArrayConstructor) *ArrayConstructor {
	if r == nil {
		return nil
	}
	return &ArrayConstructor{baseNode[*googlesql.ASTArrayConstructor]{raw: r}}
}
func (n *ArrayConstructor) isExpression() {}
func (n *ArrayConstructor) Type() *ArrayType {
	return newArrayType(must(n.raw.Type()))
}

// Elements: raw Elements(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *ArrayConstructor) Elements() []ExpressionNode {
	var elems []ExpressionNode
	for e := range childrenOfType[googlesql.ASTExpressionNode](n) {
		elems = append(elems, wrapExpr(e))
	}
	return elems
}

// ArrayElement wraps *googlesql.ASTArrayElement.
type ArrayElement struct {
	baseNode[*googlesql.ASTArrayElement]
}

func newArrayElement(r *googlesql.ASTArrayElement) *ArrayElement {
	if r == nil {
		return nil
	}
	return &ArrayElement{baseNode[*googlesql.ASTArrayElement]{raw: r}}
}
func (n *ArrayElement) isExpression() {}

// Array/Position: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *ArrayElement) Array() ExpressionNode    { return wrapExpr(must(n.raw.Array())) }
func (n *ArrayElement) Position() ExpressionNode { return wrapExpr(must(n.raw.Position())) }

// DotIdentifier wraps *googlesql.ASTDotIdentifier.
type DotIdentifier struct {
	baseNode[*googlesql.ASTDotIdentifier]
}

func newDotIdentifier(r *googlesql.ASTDotIdentifier) *DotIdentifier {
	if r == nil {
		return nil
	}
	return &DotIdentifier{baseNode[*googlesql.ASTDotIdentifier]{raw: r}}
}
func (n *DotIdentifier) isExpression() {}

// Expr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *DotIdentifier) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }
func (n *DotIdentifier) Name() *Identifier    { return newIdentifier(must(n.raw.Name())) }

// DotStar wraps *googlesql.ASTDotStar.
type DotStar struct {
	baseNode[*googlesql.ASTDotStar]
}

func newDotStar(r *googlesql.ASTDotStar) *DotStar {
	if r == nil {
		return nil
	}
	return &DotStar{baseNode[*googlesql.ASTDotStar]{raw: r}}
}
func (n *DotStar) isExpression()        {}
func (n *DotStar) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// DotStarWithModifiers wraps *googlesql.ASTDotStarWithModifiers.
type DotStarWithModifiers struct {
	baseNode[*googlesql.ASTDotStarWithModifiers]
}

func newDotStarWithModifiers(r *googlesql.ASTDotStarWithModifiers) *DotStarWithModifiers {
	if r == nil {
		return nil
	}
	return &DotStarWithModifiers{baseNode[*googlesql.ASTDotStarWithModifiers]{raw: r}}
}
func (n *DotStarWithModifiers) isExpression() {}
func (n *DotStarWithModifiers) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *DotStarWithModifiers) Modifiers() *StarModifiers {
	return newStarModifiers(must(n.raw.Modifiers()))
}

// DotGeneralizedField wraps *googlesql.ASTDotGeneralizedField.
type DotGeneralizedField struct {
	baseNode[*googlesql.ASTDotGeneralizedField]
}

func newDotGeneralizedField(r *googlesql.ASTDotGeneralizedField) *DotGeneralizedField {
	if r == nil {
		return nil
	}
	return &DotGeneralizedField{baseNode[*googlesql.ASTDotGeneralizedField]{raw: r}}
}
func (n *DotGeneralizedField) isExpression() {}
func (n *DotGeneralizedField) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *DotGeneralizedField) Path() *PathExpression {
	return newPathExpression(must(n.raw.Path()))
}

// ParameterExpr wraps *googlesql.ASTParameterExpr.
type ParameterExpr struct {
	baseNode[*googlesql.ASTParameterExpr]
}

func newParameterExpr(r *googlesql.ASTParameterExpr) *ParameterExpr {
	if r == nil {
		return nil
	}
	return &ParameterExpr{baseNode[*googlesql.ASTParameterExpr]{raw: r}}
}
func (n *ParameterExpr) isExpression()     {}
func (n *ParameterExpr) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }
func (n *ParameterExpr) Position() int     { return int(must(n.raw.Position())) }

// SystemVariableExpr wraps *googlesql.ASTSystemVariableExpr.
type SystemVariableExpr struct {
	baseNode[*googlesql.ASTSystemVariableExpr]
}

func newSystemVariableExpr(r *googlesql.ASTSystemVariableExpr) *SystemVariableExpr {
	if r == nil {
		return nil
	}
	return &SystemVariableExpr{baseNode[*googlesql.ASTSystemVariableExpr]{raw: r}}
}
func (n *SystemVariableExpr) isExpression() {}
func (n *SystemVariableExpr) Path() *PathExpression {
	return newPathExpression(must(n.raw.Path()))
}

// NamedArgument wraps *googlesql.ASTNamedArgument.
type NamedArgument struct {
	baseNode[*googlesql.ASTNamedArgument]
}

func newNamedArgument(r *googlesql.ASTNamedArgument) *NamedArgument {
	if r == nil {
		return nil
	}
	return &NamedArgument{baseNode[*googlesql.ASTNamedArgument]{raw: r}}
}
func (n *NamedArgument) isExpression()        {}
func (n *NamedArgument) Name() *Identifier    { return newIdentifier(must(n.raw.Name())) }
func (n *NamedArgument) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// Lambda wraps *googlesql.ASTLambda.
type Lambda struct{ baseNode[*googlesql.ASTLambda] }

func newLambda(r *googlesql.ASTLambda) *Lambda {
	if r == nil {
		return nil
	}
	return &Lambda{baseNode[*googlesql.ASTLambda]{raw: r}}
}
func (n *Lambda) isExpression()                {}
func (n *Lambda) ArgumentList() ExpressionNode { return wrapExpr(must(n.raw.ArgumentList())) }
func (n *Lambda) Body() ExpressionNode         { return wrapExpr(must(n.raw.Body())) }

// ─── Function call ────────────────────────────────────────────────────────────

// FunctionCall wraps *googlesql.ASTFunctionCall.
type FunctionCall struct {
	baseNode[*googlesql.ASTFunctionCall]
}

func newFunctionCall(r *googlesql.ASTFunctionCall) *FunctionCall {
	if r == nil {
		return nil
	}
	return &FunctionCall{baseNode[*googlesql.ASTFunctionCall]{raw: r}}
}
func (n *FunctionCall) isExpression()       {}
func (n *FunctionCall) Distinct() bool      { return must(n.raw.Distinct()) }
func (n *FunctionCall) IsChainedCall() bool { return must(n.raw.IsChainedCall()) }
func (n *FunctionCall) NullHandlingModifier() NullHandlingModifier {
	return must(n.raw.NullHandlingModifier())
}

// Function: raw returns *PathExpression (concrete) → *PathExpression.
func (n *FunctionCall) Function() *PathExpression {
	return newPathExpression(must(n.raw.Function()))
}

// Arguments returns the call arguments, EXCLUDING the function name node.
// Raw Arguments(i) returns ExpressionNode (interface); i is 0-indexed into
// expression-only children (the function name is NOT included in this index).
// NumChildren() counts all children including the name, so we use count-1.
func (n *FunctionCall) Arguments() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for arg := range childrenOfType[googlesql.ASTExpressionNode](n) {
		result = append(result, wrapExpr(arg))
	}
	if len(result) > 0 {
		return result[1:]
	}
	return nil
}

func (n *FunctionCall) HavingModifier() *HavingModifier {
	return newHavingModifier(must(n.raw.HavingModifier()))
}

func (n *FunctionCall) ClampedBetweenModifier() *ClampedBetweenModifier {
	return newClampedBetweenModifier(must(n.raw.ClampedBetweenModifier()))
}
func (n *FunctionCall) OrderBy() *OrderBy { return newOrderBy(must(n.raw.OrderBy())) }
func (n *FunctionCall) LimitOffset() *LimitOffset {
	return newLimitOffset(must(n.raw.LimitOffset()))
}
func (n *FunctionCall) Hint() *Hint { return newHint(must(n.raw.Hint())) }

// AnalyticFunctionCall wraps *googlesql.ASTAnalyticFunctionCall.
type AnalyticFunctionCall struct {
	baseNode[*googlesql.ASTAnalyticFunctionCall]
}

func newAnalyticFunctionCall(r *googlesql.ASTAnalyticFunctionCall) *AnalyticFunctionCall {
	if r == nil {
		return nil
	}
	return &AnalyticFunctionCall{baseNode[*googlesql.ASTAnalyticFunctionCall]{raw: r}}
}
func (n *AnalyticFunctionCall) isExpression() {}

// Function: raw returns *FunctionCall (concrete) → *FunctionCall.
func (n *AnalyticFunctionCall) Function() *FunctionCall {
	return newFunctionCall(must(n.raw.Function()))
}

func (n *AnalyticFunctionCall) WindowSpec() *WindowSpecification {
	return newWindowSpecification(must(n.raw.WindowSpec()))
}

// ─── SAMPLE / PIVOT / UNPIVOT ─────────────────────────────────────────────────

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

// ─── Literals ─────────────────────────────────────────────────────────────────

// IntLiteral wraps *googlesql.ASTIntLiteral.
type IntLiteral struct {
	baseNode[*googlesql.ASTIntLiteral]
}

func newIntLiteral(r *googlesql.ASTIntLiteral) *IntLiteral {
	if r == nil {
		return nil
	}
	return &IntLiteral{baseNode[*googlesql.ASTIntLiteral]{raw: r}}
}
func (n *IntLiteral) isExpression() {}
func (n *IntLiteral) isLeaf()       {}
func (n *IntLiteral) IsHex() bool   { return must(n.raw.IsHex()) }
func (n *IntLiteral) Image() string { return must(n.raw.Image()) }

// FloatLiteral wraps *googlesql.ASTFloatLiteral.
type FloatLiteral struct {
	baseNode[*googlesql.ASTFloatLiteral]
}

func newFloatLiteral(r *googlesql.ASTFloatLiteral) *FloatLiteral {
	if r == nil {
		return nil
	}
	return &FloatLiteral{baseNode[*googlesql.ASTFloatLiteral]{raw: r}}
}
func (n *FloatLiteral) isExpression() {}
func (n *FloatLiteral) isLeaf()       {}
func (n *FloatLiteral) Image() string { return must(n.raw.Image()) }

// BooleanLiteral wraps *googlesql.ASTBooleanLiteral.
type BooleanLiteral struct {
	baseNode[*googlesql.ASTBooleanLiteral]
}

func newBooleanLiteral(r *googlesql.ASTBooleanLiteral) *BooleanLiteral {
	if r == nil {
		return nil
	}
	return &BooleanLiteral{baseNode[*googlesql.ASTBooleanLiteral]{raw: r}}
}
func (n *BooleanLiteral) isExpression() {}
func (n *BooleanLiteral) isLeaf()       {}
func (n *BooleanLiteral) Value() bool   { return must(n.raw.Value()) }
func (n *BooleanLiteral) Image() string { return must(n.raw.Image()) }

// NullLiteral wraps *googlesql.ASTNullLiteral.
type NullLiteral struct {
	baseNode[*googlesql.ASTNullLiteral]
}

func newNullLiteral(r *googlesql.ASTNullLiteral) *NullLiteral {
	if r == nil {
		return nil
	}
	return &NullLiteral{baseNode[*googlesql.ASTNullLiteral]{raw: r}}
}
func (n *NullLiteral) isExpression() {}
func (n *NullLiteral) isLeaf()       {}
func (n *NullLiteral) Image() string { return must(n.raw.Image()) }

// StringLiteral wraps *googlesql.ASTStringLiteral.
type StringLiteral struct {
	baseNode[*googlesql.ASTStringLiteral]
}

func newStringLiteral(r *googlesql.ASTStringLiteral) *StringLiteral {
	if r == nil {
		return nil
	}
	return &StringLiteral{baseNode[*googlesql.ASTStringLiteral]{raw: r}}
}
func (n *StringLiteral) isExpression()       {}
func (n *StringLiteral) isLeaf()             {}
func (n *StringLiteral) StringValue() string { return must(n.raw.StringValue()) }

// Components: raw Components(i) returns *StringLiteralComponent (concrete).
func (n *StringLiteral) Components() []*StringLiteralComponent {
	count := n.NumChildren()
	result := make([]*StringLiteralComponent, 0, count)
	for i := range count {
		c := must(n.raw.Components(int32(i)))
		if c == nil {
			break
		}
		result = append(result, newStringLiteralComponent(c))
	}
	return result
}

// StringLiteralComponent wraps *googlesql.ASTStringLiteralComponent.
type StringLiteralComponent struct {
	baseNode[*googlesql.ASTStringLiteralComponent]
}

func newStringLiteralComponent(r *googlesql.ASTStringLiteralComponent) *StringLiteralComponent {
	if r == nil {
		return nil
	}
	return &StringLiteralComponent{baseNode[*googlesql.ASTStringLiteralComponent]{raw: r}}
}
func (n *StringLiteralComponent) isExpression()       {}
func (n *StringLiteralComponent) isLeaf()             {}
func (n *StringLiteralComponent) StringValue() string { return must(n.raw.StringValue()) }

// BytesLiteral wraps *googlesql.ASTBytesLiteral.
type BytesLiteral struct {
	baseNode[*googlesql.ASTBytesLiteral]
}

func newBytesLiteral(r *googlesql.ASTBytesLiteral) *BytesLiteral {
	if r == nil {
		return nil
	}
	return &BytesLiteral{baseNode[*googlesql.ASTBytesLiteral]{raw: r}}
}
func (n *BytesLiteral) isExpression()      {}
func (n *BytesLiteral) isLeaf()            {}
func (n *BytesLiteral) BytesValue() string { return must(n.raw.BytesValue()) }
func (n *BytesLiteral) Components() []*BytesLiteralComponent {
	count := n.NumChildren()
	result := make([]*BytesLiteralComponent, 0, count)
	for i := range count {
		c := must(n.raw.Components(int32(i)))
		if c == nil {
			break
		}
		result = append(result, newBytesLiteralComponent(c))
	}
	return result
}

// BytesLiteralComponent wraps *googlesql.ASTBytesLiteralComponent.
type BytesLiteralComponent struct {
	baseNode[*googlesql.ASTBytesLiteralComponent]
}

func newBytesLiteralComponent(r *googlesql.ASTBytesLiteralComponent) *BytesLiteralComponent {
	if r == nil {
		return nil
	}
	return &BytesLiteralComponent{baseNode[*googlesql.ASTBytesLiteralComponent]{raw: r}}
}
func (n *BytesLiteralComponent) isExpression()      {}
func (n *BytesLiteralComponent) isLeaf()            {}
func (n *BytesLiteralComponent) BytesValue() string { return must(n.raw.BytesValue()) }

// NumericLiteral wraps *googlesql.ASTNumericLiteral.
type NumericLiteral struct {
	baseNode[*googlesql.ASTNumericLiteral]
}

func newNumericLiteral(r *googlesql.ASTNumericLiteral) *NumericLiteral {
	if r == nil {
		return nil
	}
	return &NumericLiteral{baseNode[*googlesql.ASTNumericLiteral]{raw: r}}
}
func (n *NumericLiteral) isExpression() {}
func (n *NumericLiteral) isLeaf()       {}
func (n *NumericLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// BigNumericLiteral wraps *googlesql.ASTBigNumericLiteral.
type BigNumericLiteral struct {
	baseNode[*googlesql.ASTBigNumericLiteral]
}

func newBigNumericLiteral(r *googlesql.ASTBigNumericLiteral) *BigNumericLiteral {
	if r == nil {
		return nil
	}
	return &BigNumericLiteral{baseNode[*googlesql.ASTBigNumericLiteral]{raw: r}}
}
func (n *BigNumericLiteral) isExpression() {}
func (n *BigNumericLiteral) isLeaf()       {}
func (n *BigNumericLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// JSONLiteral wraps *googlesql.ASTJSONLiteral.
type JSONLiteral struct {
	baseNode[*googlesql.ASTJSONLiteral]
}

func newJSONLiteral(r *googlesql.ASTJSONLiteral) *JSONLiteral {
	if r == nil {
		return nil
	}
	return &JSONLiteral{baseNode[*googlesql.ASTJSONLiteral]{raw: r}}
}
func (n *JSONLiteral) isExpression() {}
func (n *JSONLiteral) isLeaf()       {}
func (n *JSONLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// DateOrTimeLiteral wraps *googlesql.ASTDateOrTimeLiteral.
type DateOrTimeLiteral struct {
	baseNode[*googlesql.ASTDateOrTimeLiteral]
}

func newDateOrTimeLiteral(r *googlesql.ASTDateOrTimeLiteral) *DateOrTimeLiteral {
	if r == nil {
		return nil
	}
	return &DateOrTimeLiteral{baseNode[*googlesql.ASTDateOrTimeLiteral]{raw: r}}
}
func (n *DateOrTimeLiteral) isExpression()              {}
func (n *DateOrTimeLiteral) TypeKind() DateTimeTypeKind { return must(n.raw.TypeKind()) }
func (n *DateOrTimeLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// DefaultLiteral wraps *googlesql.ASTDefaultLiteral.
type DefaultLiteral struct {
	baseNode[*googlesql.ASTDefaultLiteral]
}

func newDefaultLiteral(r *googlesql.ASTDefaultLiteral) *DefaultLiteral {
	if r == nil {
		return nil
	}
	return &DefaultLiteral{baseNode[*googlesql.ASTDefaultLiteral]{raw: r}}
}
func (n *DefaultLiteral) isExpression() {}

// MaxLiteral wraps *googlesql.ASTMaxLiteral.
type MaxLiteral struct {
	baseNode[*googlesql.ASTMaxLiteral]
}

func newMaxLiteral(r *googlesql.ASTMaxLiteral) *MaxLiteral {
	if r == nil {
		return nil
	}
	return &MaxLiteral{baseNode[*googlesql.ASTMaxLiteral]{raw: r}}
}
func (n *MaxLiteral) isExpression() {}
func (n *MaxLiteral) isLeaf()       {}

// RangeLiteral wraps *googlesql.ASTRangeLiteral.
type RangeLiteral struct {
	baseNode[*googlesql.ASTRangeLiteral]
}

func newRangeLiteral(r *googlesql.ASTRangeLiteral) *RangeLiteral {
	if r == nil {
		return nil
	}
	return &RangeLiteral{baseNode[*googlesql.ASTRangeLiteral]{raw: r}}
}
func (n *RangeLiteral) isExpression() {}
func (n *RangeLiteral) Type() *RangeType {
	return newRangeType(must(n.raw.Type()))
}

func (n *RangeLiteral) RangeValue() *StringLiteral {
	return newStringLiteral(must(n.raw.RangeValue()))
}

// ─── OPTIONS LIST ─────────────────────────────────────────────────────────────

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
func (n *OptionsEntry) Name() *Identifier          { return newIdentifier(must(n.raw.Name())) }
func (n *OptionsEntry) AssignmentOp() AssignmentOp { return must(n.raw.AssignmentOp()) }

// Value: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *OptionsEntry) Value() ExpressionNode { return wrapExpr(must(n.raw.Value())) }

// ─── COLUMN LIST ─────────────────────────────────────────────────────────────

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

// ─── DESCRIPTOR ───────────────────────────────────────────────────────────────

// Descriptor wraps *googlesql.ASTDescriptor.
type Descriptor struct {
	baseNode[*googlesql.ASTDescriptor]
}

func newDescriptor(r *googlesql.ASTDescriptor) *Descriptor {
	if r == nil {
		return nil
	}
	return &Descriptor{baseNode[*googlesql.ASTDescriptor]{raw: r}}
}
func (n *Descriptor) isExpression() {}
func (n *Descriptor) Columns() *DescriptorColumnList {
	return newDescriptorColumnList(must(n.raw.Columns()))
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

// ─── STATEMENT LIST ───────────────────────────────────────────────────────────

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

// TVFSchema wraps *googlesql.ASTTVFSchema.
type TVFSchema struct {
	baseNode[*googlesql.ASTTVFSchema]
}

func newTVFSchema(r *googlesql.ASTTVFSchema) *TVFSchema {
	if r == nil {
		return nil
	}
	return &TVFSchema{baseNode[*googlesql.ASTTVFSchema]{raw: r}}
}

func (n *TVFSchema) Columns() []*TVFSchemaColumn {
	var result []*TVFSchemaColumn
	for _, c := range n.Children() {
		if col, ok := c.(*TVFSchemaColumn); ok {
			result = append(result, col)
		}
	}
	return result
}

// StructConstructorWithKeyword wraps *googlesql.ASTStructConstructorWithKeyword.
type StructConstructorWithKeyword struct {
	baseNode[*googlesql.ASTStructConstructorWithKeyword]
}

func newStructConstructorWithKeyword(r *googlesql.ASTStructConstructorWithKeyword) *StructConstructorWithKeyword {
	if r == nil {
		return nil
	}
	return &StructConstructorWithKeyword{baseNode[*googlesql.ASTStructConstructorWithKeyword]{raw: r}}
}
func (n *StructConstructorWithKeyword) isExpression() {}
func (n *StructConstructorWithKeyword) StructType() *StructType {
	return newStructType(must(n.raw.StructType()))
}

// Fields returns all struct constructor arguments.
func (n *StructConstructorWithKeyword) Fields() []*StructConstructorArg {
	var result []*StructConstructorArg
	for _, c := range n.Children() {
		if a, ok := c.(*StructConstructorArg); ok {
			result = append(result, a)
		}
	}
	return result
}

// StructConstructorWithParens wraps *googlesql.ASTStructConstructorWithParens.
type StructConstructorWithParens struct {
	baseNode[*googlesql.ASTStructConstructorWithParens]
}

func newStructConstructorWithParens(r *googlesql.ASTStructConstructorWithParens) *StructConstructorWithParens {
	if r == nil {
		return nil
	}
	return &StructConstructorWithParens{baseNode[*googlesql.ASTStructConstructorWithParens]{raw: r}}
}
func (n *StructConstructorWithParens) isExpression() {}

// FieldExpressions returns all field expressions.
func (n *StructConstructorWithParens) FieldExpressions() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		e := must(n.raw.FieldExpressions(int32(i)))
		if e == nil {
			break
		}
		result = append(result, wrapExpr(e))
	}
	return result
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
