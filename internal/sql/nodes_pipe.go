package sql

import (
	"github.com/goccy/go-googlesql"
)

// PipeAggregate wraps *googlesql.ASTPipeAggregate.
type PipeAggregate struct {
	baseNode[*googlesql.ASTPipeAggregate]
}

func newPipeAggregate(r *googlesql.ASTPipeAggregate) *PipeAggregate {
	if r == nil {
		return nil
	}
	return &PipeAggregate{baseNode[*googlesql.ASTPipeAggregate]{raw: r}}
}

func (n *PipeAggregate) isPipeOperator() {}

func (n *PipeAggregate) Select() *Select {
	return newSelect(must(n.raw.Select()))
}

func (n *PipeAggregate) WithModifier() *WithModifier {
	return newWithModifier(must(n.raw.WithModifier()))
}

// PipeAs wraps *googlesql.ASTPipeAs.
type PipeAs struct {
	baseNode[*googlesql.ASTPipeAs]
}

func newPipeAs(r *googlesql.ASTPipeAs) *PipeAs {
	if r == nil {
		return nil
	}
	return &PipeAs{baseNode[*googlesql.ASTPipeAs]{raw: r}}
}

func (n *PipeAs) isPipeOperator() {}

func (n *PipeAs) Alias() *Alias {
	return newAlias(must(n.raw.Alias()))
}

// PipeAssert wraps *googlesql.ASTPipeAssert.
type PipeAssert struct {
	baseNode[*googlesql.ASTPipeAssert]
}

func newPipeAssert(r *googlesql.ASTPipeAssert) *PipeAssert {
	if r == nil {
		return nil
	}
	return &PipeAssert{baseNode[*googlesql.ASTPipeAssert]{raw: r}}
}

func (n *PipeAssert) isPipeOperator() {}

func (n *PipeAssert) Condition() ExpressionNode {
	return wrapExpr(must(n.raw.Condition()))
}

// PipeCall wraps *googlesql.ASTPipeCall.
type PipeCall struct {
	baseNode[*googlesql.ASTPipeCall]
}

func newPipeCall(r *googlesql.ASTPipeCall) *PipeCall {
	if r == nil {
		return nil
	}
	return &PipeCall{baseNode[*googlesql.ASTPipeCall]{raw: r}}
}

func (n *PipeCall) isPipeOperator() {}

func (n *PipeCall) TVF() *TVF {
	return newTVF(must(n.raw.Tvf()))
}

// PipeCreateTable wraps *googlesql.ASTPipeCreateTable.
type PipeCreateTable struct {
	baseNode[*googlesql.ASTPipeCreateTable]
}

func newPipeCreateTable(r *googlesql.ASTPipeCreateTable) *PipeCreateTable {
	if r == nil {
		return nil
	}
	return &PipeCreateTable{baseNode[*googlesql.ASTPipeCreateTable]{raw: r}}
}

func (n *PipeCreateTable) isPipeOperator() {}

func (n *PipeCreateTable) CreateTableStatement() *CreateTableStatement {
	return newCreateTableStatement(must(n.raw.CreateTableStatement()))
}

// PipeDescribe wraps *googlesql.ASTPipeDescribe.
type PipeDescribe struct {
	baseNode[*googlesql.ASTPipeDescribe]
}

func newPipeDescribe(r *googlesql.ASTPipeDescribe) *PipeDescribe {
	if r == nil {
		return nil
	}
	return &PipeDescribe{baseNode[*googlesql.ASTPipeDescribe]{raw: r}}
}

func (n *PipeDescribe) isPipeOperator() {}

// PipeDistinct wraps *googlesql.ASTPipeDistinct.
type PipeDistinct struct {
	baseNode[*googlesql.ASTPipeDistinct]
}

func newPipeDistinct(r *googlesql.ASTPipeDistinct) *PipeDistinct {
	if r == nil {
		return nil
	}
	return &PipeDistinct{baseNode[*googlesql.ASTPipeDistinct]{raw: r}}
}

func (n *PipeDistinct) isPipeOperator() {}

// PipeDrop wraps *googlesql.ASTPipeDrop.
type PipeDrop struct {
	baseNode[*googlesql.ASTPipeDrop]
}

func newPipeDrop(r *googlesql.ASTPipeDrop) *PipeDrop {
	if r == nil {
		return nil
	}
	return &PipeDrop{baseNode[*googlesql.ASTPipeDrop]{raw: r}}
}

func (n *PipeDrop) isPipeOperator() {}

func (n *PipeDrop) ColumnList() *IdentifierList {
	return newIdentifierList(must(n.raw.ColumnList()))
}

// PipeExportData wraps *googlesql.ASTPipeExportData.
type PipeExportData struct {
	baseNode[*googlesql.ASTPipeExportData]
}

func newPipeExportData(r *googlesql.ASTPipeExportData) *PipeExportData {
	if r == nil {
		return nil
	}
	return &PipeExportData{baseNode[*googlesql.ASTPipeExportData]{raw: r}}
}

func (n *PipeExportData) isPipeOperator() {}

func (n *PipeExportData) ExportDataStatement() StatementNode {
	return wrapStmt(must(n.raw.ExportDataStatement()))
}

// PipeExtend wraps *googlesql.ASTPipeExtend.
type PipeExtend struct {
	baseNode[*googlesql.ASTPipeExtend]
}

func newPipeExtend(r *googlesql.ASTPipeExtend) *PipeExtend {
	if r == nil {
		return nil
	}
	return &PipeExtend{baseNode[*googlesql.ASTPipeExtend]{raw: r}}
}

func (n *PipeExtend) isPipeOperator() {}

func (n *PipeExtend) Select() *Select {
	return newSelect(must(n.raw.Select()))
}

// PipeFork wraps *googlesql.ASTPipeFork.
type PipeFork struct {
	baseNode[*googlesql.ASTPipeFork]
}

func newPipeFork(r *googlesql.ASTPipeFork) *PipeFork {
	if r == nil {
		return nil
	}
	return &PipeFork{baseNode[*googlesql.ASTPipeFork]{raw: r}}
}

func (n *PipeFork) isPipeOperator() {}

func (n *PipeFork) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *PipeFork) SubpipelineList() []*Subpipeline {
	var result []*Subpipeline
	for s := range childrenOfType[*googlesql.ASTSubpipeline](n) {
		result = append(result, newSubpipeline(s))
	}
	return result
}

// PipeIf wraps *googlesql.ASTPipeIf.
type PipeIf struct {
	baseNode[*googlesql.ASTPipeIf]
}

func newPipeIf(r *googlesql.ASTPipeIf) *PipeIf {
	if r == nil {
		return nil
	}
	return &PipeIf{baseNode[*googlesql.ASTPipeIf]{raw: r}}
}

func (n *PipeIf) isPipeOperator() {}

func (n *PipeIf) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *PipeIf) IfCases() []*PipeIfCase {
	var result []*PipeIfCase
	for c := range childrenOfType[*googlesql.ASTPipeIfCase](n) {
		result = append(result, newPipeIfCase(c))
	}
	return result
}

func (n *PipeIf) ElseSubpipeline() *Subpipeline {
	return newSubpipeline(must(n.raw.ElseSubpipeline()))
}

// PipeIfCase wraps *googlesql.ASTPipeIfCase.
type PipeIfCase struct {
	baseNode[*googlesql.ASTPipeIfCase]
}

func newPipeIfCase(r *googlesql.ASTPipeIfCase) *PipeIfCase {
	if r == nil {
		return nil
	}
	return &PipeIfCase{baseNode[*googlesql.ASTPipeIfCase]{raw: r}}
}

func (n *PipeIfCase) Condition() ExpressionNode {
	return wrapExpr(must(n.raw.Condition()))
}

func (n *PipeIfCase) Subpipeline() *Subpipeline {
	return newSubpipeline(must(n.raw.Subpipeline()))
}

// PipeInsert wraps *googlesql.ASTPipeInsert.
type PipeInsert struct {
	baseNode[*googlesql.ASTPipeInsert]
}

func newPipeInsert(r *googlesql.ASTPipeInsert) *PipeInsert {
	if r == nil {
		return nil
	}
	return &PipeInsert{baseNode[*googlesql.ASTPipeInsert]{raw: r}}
}

func (n *PipeInsert) isPipeOperator() {}

func (n *PipeInsert) InsertStatement() *InsertStatement {
	return newInsertStatement(must(n.raw.InsertStatement()))
}

// PipeJoin wraps *googlesql.ASTPipeJoin.
type PipeJoin struct {
	baseNode[*googlesql.ASTPipeJoin]
}

func newPipeJoin(r *googlesql.ASTPipeJoin) *PipeJoin {
	if r == nil {
		return nil
	}
	return &PipeJoin{baseNode[*googlesql.ASTPipeJoin]{raw: r}}
}

func (n *PipeJoin) isPipeOperator() {}

func (n *PipeJoin) Join() *Join {
	return newJoin(must(n.raw.Join()))
}

// PipeJoinLHSPlaceholder wraps *googlesql.ASTPipeJoinLhsPlaceholder.
type PipeJoinLHSPlaceholder struct {
	baseNode[*googlesql.ASTPipeJoinLhsPlaceholder]
}

func newPipeJoinLHSPlaceholder(r *googlesql.ASTPipeJoinLhsPlaceholder) *PipeJoinLHSPlaceholder {
	if r == nil {
		return nil
	}
	return &PipeJoinLHSPlaceholder{baseNode[*googlesql.ASTPipeJoinLhsPlaceholder]{raw: r}}
}

// PipeLimitOffset wraps *googlesql.ASTPipeLimitOffset.
type PipeLimitOffset struct {
	baseNode[*googlesql.ASTPipeLimitOffset]
}

func newPipeLimitOffset(r *googlesql.ASTPipeLimitOffset) *PipeLimitOffset {
	if r == nil {
		return nil
	}
	return &PipeLimitOffset{baseNode[*googlesql.ASTPipeLimitOffset]{raw: r}}
}

func (n *PipeLimitOffset) isPipeOperator() {}

func (n *PipeLimitOffset) LimitOffset() *LimitOffset {
	return newLimitOffset(must(n.raw.LimitOffset()))
}

// PipeLog wraps *googlesql.ASTPipeLog.
type PipeLog struct {
	baseNode[*googlesql.ASTPipeLog]
}

func newPipeLog(r *googlesql.ASTPipeLog) *PipeLog {
	if r == nil {
		return nil
	}
	return &PipeLog{baseNode[*googlesql.ASTPipeLog]{raw: r}}
}

func (n *PipeLog) isPipeOperator() {}

func (n *PipeLog) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *PipeLog) Subpipeline() *Subpipeline {
	return newSubpipeline(must(n.raw.Subpipeline()))
}

// PipeMatchRecognize wraps *googlesql.ASTPipeMatchRecognize.
type PipeMatchRecognize struct {
	baseNode[*googlesql.ASTPipeMatchRecognize]
}

func newPipeMatchRecognize(r *googlesql.ASTPipeMatchRecognize) *PipeMatchRecognize {
	if r == nil {
		return nil
	}
	return &PipeMatchRecognize{baseNode[*googlesql.ASTPipeMatchRecognize]{raw: r}}
}

func (n *PipeMatchRecognize) isPipeOperator() {}

// PipeOrderBy wraps *googlesql.ASTPipeOrderBy.
type PipeOrderBy struct {
	baseNode[*googlesql.ASTPipeOrderBy]
}

func newPipeOrderBy(r *googlesql.ASTPipeOrderBy) *PipeOrderBy {
	if r == nil {
		return nil
	}
	return &PipeOrderBy{baseNode[*googlesql.ASTPipeOrderBy]{raw: r}}
}

func (n *PipeOrderBy) isPipeOperator() {}

func (n *PipeOrderBy) OrderBy() *OrderBy {
	return newOrderBy(must(n.raw.OrderBy()))
}

// PipePivot wraps *googlesql.ASTPipePivot.
type PipePivot struct {
	baseNode[*googlesql.ASTPipePivot]
}

func newPipePivot(r *googlesql.ASTPipePivot) *PipePivot {
	if r == nil {
		return nil
	}
	return &PipePivot{baseNode[*googlesql.ASTPipePivot]{raw: r}}
}

func (n *PipePivot) isPipeOperator() {}

func (n *PipePivot) PivotClause() *PivotClause {
	return newPivotClause(must(n.raw.PivotClause()))
}

// PipeRecursiveUnion wraps *googlesql.ASTPipeRecursiveUnion.
type PipeRecursiveUnion struct {
	baseNode[*googlesql.ASTPipeRecursiveUnion]
}

func newPipeRecursiveUnion(r *googlesql.ASTPipeRecursiveUnion) *PipeRecursiveUnion {
	if r == nil {
		return nil
	}
	return &PipeRecursiveUnion{baseNode[*googlesql.ASTPipeRecursiveUnion]{raw: r}}
}

func (n *PipeRecursiveUnion) isPipeOperator() {}

func (n *PipeRecursiveUnion) Alias() *Alias {
	return newAlias(must(n.raw.Alias()))
}

func (n *PipeRecursiveUnion) Metadata() *SetOperationMetadata {
	return newSetOperationMetadata(must(n.raw.Metadata()))
}

func (n *PipeRecursiveUnion) InputSubpipeline() *Subpipeline {
	return newSubpipeline(must(n.raw.InputSubpipeline()))
}

func (n *PipeRecursiveUnion) InputSubquery() QueryExpressionNode {
	return wrapQueryExpr(must(n.raw.InputSubquery()))
}

// PipeRename wraps *googlesql.ASTPipeRename.
type PipeRename struct {
	baseNode[*googlesql.ASTPipeRename]
}

func newPipeRename(r *googlesql.ASTPipeRename) *PipeRename {
	if r == nil {
		return nil
	}
	return &PipeRename{baseNode[*googlesql.ASTPipeRename]{raw: r}}
}

func (n *PipeRename) isPipeOperator() {}

func (n *PipeRename) RenameItemList() []*PipeRenameItem {
	var result []*PipeRenameItem
	for item := range childrenOfType[*googlesql.ASTPipeRenameItem](n) {
		result = append(result, newPipeRenameItem(item))
	}
	return result
}

// PipeRenameItem wraps *googlesql.ASTPipeRenameItem.
type PipeRenameItem struct {
	baseNode[*googlesql.ASTPipeRenameItem]
}

func newPipeRenameItem(r *googlesql.ASTPipeRenameItem) *PipeRenameItem {
	if r == nil {
		return nil
	}
	return &PipeRenameItem{baseNode[*googlesql.ASTPipeRenameItem]{raw: r}}
}

func (n *PipeRenameItem) OldName() *Identifier {
	return newIdentifier(must(n.raw.OldName()))
}

func (n *PipeRenameItem) NewName() *Identifier {
	return newIdentifier(must(n.raw.NewName()))
}

// PipeSelect wraps *googlesql.ASTPipeSelect.
type PipeSelect struct {
	baseNode[*googlesql.ASTPipeSelect]
}

func newPipeSelect(r *googlesql.ASTPipeSelect) *PipeSelect {
	if r == nil {
		return nil
	}
	return &PipeSelect{baseNode[*googlesql.ASTPipeSelect]{raw: r}}
}

func (n *PipeSelect) isPipeOperator() {}

func (n *PipeSelect) Select() *Select {
	return newSelect(must(n.raw.Select()))
}

// PipeSet wraps *googlesql.ASTPipeSet.
type PipeSet struct {
	baseNode[*googlesql.ASTPipeSet]
}

func newPipeSet(r *googlesql.ASTPipeSet) *PipeSet {
	if r == nil {
		return nil
	}
	return &PipeSet{baseNode[*googlesql.ASTPipeSet]{raw: r}}
}

func (n *PipeSet) isPipeOperator() {}

func (n *PipeSet) SetItemList() []*PipeSetItem {
	var result []*PipeSetItem
	for item := range childrenOfType[*googlesql.ASTPipeSetItem](n) {
		result = append(result, newPipeSetItem(item))
	}
	return result
}

// PipeSetItem wraps *googlesql.ASTPipeSetItem.
type PipeSetItem struct {
	baseNode[*googlesql.ASTPipeSetItem]
}

func newPipeSetItem(r *googlesql.ASTPipeSetItem) *PipeSetItem {
	if r == nil {
		return nil
	}
	return &PipeSetItem{baseNode[*googlesql.ASTPipeSetItem]{raw: r}}
}

func (n *PipeSetItem) Column() *Identifier {
	return newIdentifier(must(n.raw.Column()))
}

func (n *PipeSetItem) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// PipeSetOperation wraps *googlesql.ASTPipeSetOperation.
type PipeSetOperation struct {
	baseNode[*googlesql.ASTPipeSetOperation]
}

func newPipeSetOperation(r *googlesql.ASTPipeSetOperation) *PipeSetOperation {
	if r == nil {
		return nil
	}
	return &PipeSetOperation{baseNode[*googlesql.ASTPipeSetOperation]{raw: r}}
}

func (n *PipeSetOperation) isPipeOperator() {}

func (n *PipeSetOperation) Inputs() []QueryExpressionNode {
	var result []QueryExpressionNode
	for _, c := range n.Children() {
		if qe, ok := c.(QueryExpressionNode); ok {
			result = append(result, qe)
		}
	}
	return result
}

func (n *PipeSetOperation) Metadata() *SetOperationMetadata {
	return newSetOperationMetadata(must(n.raw.Metadata()))
}

// PipeStaticDescribe wraps *googlesql.ASTPipeStaticDescribe.
type PipeStaticDescribe struct {
	baseNode[*googlesql.ASTPipeStaticDescribe]
}

func newPipeStaticDescribe(r *googlesql.ASTPipeStaticDescribe) *PipeStaticDescribe {
	if r == nil {
		return nil
	}
	return &PipeStaticDescribe{baseNode[*googlesql.ASTPipeStaticDescribe]{raw: r}}
}

func (n *PipeStaticDescribe) isPipeOperator() {}

// PipeTablesample wraps *googlesql.ASTPipeTablesample.
type PipeTablesample struct {
	baseNode[*googlesql.ASTPipeTablesample]
}

func newPipeTablesample(r *googlesql.ASTPipeTablesample) *PipeTablesample {
	if r == nil {
		return nil
	}
	return &PipeTablesample{baseNode[*googlesql.ASTPipeTablesample]{raw: r}}
}

func (n *PipeTablesample) isPipeOperator() {}

func (n *PipeTablesample) Sample() *SampleClause {
	return newSampleClause(must(n.raw.Sample()))
}

// PipeTee wraps *googlesql.ASTPipeTee.
type PipeTee struct {
	baseNode[*googlesql.ASTPipeTee]
}

func newPipeTee(r *googlesql.ASTPipeTee) *PipeTee {
	if r == nil {
		return nil
	}
	return &PipeTee{baseNode[*googlesql.ASTPipeTee]{raw: r}}
}

func (n *PipeTee) isPipeOperator() {}

func (n *PipeTee) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *PipeTee) SubpipelineList() []*Subpipeline {
	var result []*Subpipeline
	for s := range childrenOfType[*googlesql.ASTSubpipeline](n) {
		result = append(result, newSubpipeline(s))
	}
	return result
}

// PipeUnpivot wraps *googlesql.ASTPipeUnpivot.
type PipeUnpivot struct {
	baseNode[*googlesql.ASTPipeUnpivot]
}

func newPipeUnpivot(r *googlesql.ASTPipeUnpivot) *PipeUnpivot {
	if r == nil {
		return nil
	}
	return &PipeUnpivot{baseNode[*googlesql.ASTPipeUnpivot]{raw: r}}
}

func (n *PipeUnpivot) isPipeOperator() {}

func (n *PipeUnpivot) UnpivotClause() *UnpivotClause {
	return newUnpivotClause(must(n.raw.UnpivotClause()))
}

// PipeWhere wraps *googlesql.ASTPipeWhere.
type PipeWhere struct {
	baseNode[*googlesql.ASTPipeWhere]
}

func newPipeWhere(r *googlesql.ASTPipeWhere) *PipeWhere {
	if r == nil {
		return nil
	}
	return &PipeWhere{baseNode[*googlesql.ASTPipeWhere]{raw: r}}
}

func (n *PipeWhere) isPipeOperator() {}

func (n *PipeWhere) Where() *WhereClause {
	return newWhereClause(must(n.raw.Where()))
}

// PipeWindow wraps *googlesql.ASTPipeWindow.
type PipeWindow struct {
	baseNode[*googlesql.ASTPipeWindow]
}

func newPipeWindow(r *googlesql.ASTPipeWindow) *PipeWindow {
	if r == nil {
		return nil
	}
	return &PipeWindow{baseNode[*googlesql.ASTPipeWindow]{raw: r}}
}

func (n *PipeWindow) isPipeOperator() {}

func (n *PipeWindow) Select() *Select {
	return newSelect(must(n.raw.Select()))
}

// PipeWith wraps *googlesql.ASTPipeWith.
type PipeWith struct {
	baseNode[*googlesql.ASTPipeWith]
}

func newPipeWith(r *googlesql.ASTPipeWith) *PipeWith {
	if r == nil {
		return nil
	}
	return &PipeWith{baseNode[*googlesql.ASTPipeWith]{raw: r}}
}

func (n *PipeWith) isPipeOperator() {}

func (n *PipeWith) WithClause() *WithClause {
	return newWithClause(must(n.raw.WithClause()))
}
