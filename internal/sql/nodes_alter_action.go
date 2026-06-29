package sql

import (
	"github.com/goccy/go-googlesql"
)

// AddColumnAction wraps *googlesql.ASTAddColumnAction.
type AddColumnAction struct {
	baseNode[*googlesql.ASTAddColumnAction]
}

func newAddColumnAction(r *googlesql.ASTAddColumnAction) *AddColumnAction {
	if r == nil {
		return nil
	}
	return &AddColumnAction{baseNode[*googlesql.ASTAddColumnAction]{raw: r}}
}

func (n *AddColumnAction) isAlterAction() {}

func (n *AddColumnAction) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *AddColumnAction) ColumnDefinition() *ColumnDefinition {
	return newColumnDefinition(must(n.raw.ColumnDefinition()))
}

// AddConstraintAction wraps *googlesql.ASTAddConstraintAction.
type AddConstraintAction struct {
	baseNode[*googlesql.ASTAddConstraintAction]
}

func newAddConstraintAction(r *googlesql.ASTAddConstraintAction) *AddConstraintAction {
	if r == nil {
		return nil
	}
	return &AddConstraintAction{baseNode[*googlesql.ASTAddConstraintAction]{raw: r}}
}

func (n *AddConstraintAction) isAlterAction() {}

func (n *AddConstraintAction) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *AddConstraintAction) Constraint() Node { return Wrap(must(n.raw.Constraint())) }

// AlterActionList wraps *googlesql.ASTAlterActionList.
type AlterActionList struct {
	baseNode[*googlesql.ASTAlterActionList]
}

func newAlterActionList(r *googlesql.ASTAlterActionList) *AlterActionList {
	if r == nil {
		return nil
	}
	return &AlterActionList{baseNode[*googlesql.ASTAlterActionList]{raw: r}}
}

// Actions returns all AlterActionNode children.
func (n *AlterActionList) Actions() []AlterActionNode {
	count := n.NumChildren()
	result := make([]AlterActionNode, 0, count)
	for i := range count {
		c := must(n.raw.Actions(int32(i)))
		if c == nil {
			break
		}
		result = append(result, wrapAlterAction(c))
	}
	return result
}

// AlterColumnDropDefaultAction wraps *googlesql.ASTAlterColumnDropDefaultAction.
type AlterColumnDropDefaultAction struct {
	baseNode[*googlesql.ASTAlterColumnDropDefaultAction]
}

func newAlterColumnDropDefaultAction(r *googlesql.ASTAlterColumnDropDefaultAction) *AlterColumnDropDefaultAction {
	if r == nil {
		return nil
	}
	return &AlterColumnDropDefaultAction{baseNode[*googlesql.ASTAlterColumnDropDefaultAction]{raw: r}}
}

func (n *AlterColumnDropDefaultAction) isAlterAction() {}

func (n *AlterColumnDropDefaultAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *AlterColumnDropDefaultAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

// AlterColumnDropNotNullAction wraps *googlesql.ASTAlterColumnDropNotNullAction.
type AlterColumnDropNotNullAction struct {
	baseNode[*googlesql.ASTAlterColumnDropNotNullAction]
}

func newAlterColumnDropNotNullAction(r *googlesql.ASTAlterColumnDropNotNullAction) *AlterColumnDropNotNullAction {
	if r == nil {
		return nil
	}
	return &AlterColumnDropNotNullAction{baseNode[*googlesql.ASTAlterColumnDropNotNullAction]{raw: r}}
}

func (n *AlterColumnDropNotNullAction) isAlterAction() {}

func (n *AlterColumnDropNotNullAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *AlterColumnDropNotNullAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

// AlterColumnOptionsAction wraps *googlesql.ASTAlterColumnOptionsAction.
type AlterColumnOptionsAction struct {
	baseNode[*googlesql.ASTAlterColumnOptionsAction]
}

func newAlterColumnOptionsAction(r *googlesql.ASTAlterColumnOptionsAction) *AlterColumnOptionsAction {
	if r == nil {
		return nil
	}
	return &AlterColumnOptionsAction{baseNode[*googlesql.ASTAlterColumnOptionsAction]{raw: r}}
}

func (n *AlterColumnOptionsAction) isAlterAction() {}

func (n *AlterColumnOptionsAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *AlterColumnOptionsAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

func (n *AlterColumnOptionsAction) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// AlterColumnSetDefaultAction wraps *googlesql.ASTAlterColumnSetDefaultAction.
type AlterColumnSetDefaultAction struct {
	baseNode[*googlesql.ASTAlterColumnSetDefaultAction]
}

func newAlterColumnSetDefaultAction(r *googlesql.ASTAlterColumnSetDefaultAction) *AlterColumnSetDefaultAction {
	if r == nil {
		return nil
	}
	return &AlterColumnSetDefaultAction{baseNode[*googlesql.ASTAlterColumnSetDefaultAction]{raw: r}}
}

func (n *AlterColumnSetDefaultAction) isAlterAction() {}

func (n *AlterColumnSetDefaultAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *AlterColumnSetDefaultAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

func (n *AlterColumnSetDefaultAction) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

// AlterColumnTypeAction wraps *googlesql.ASTAlterColumnTypeAction.
type AlterColumnTypeAction struct {
	baseNode[*googlesql.ASTAlterColumnTypeAction]
}

func newAlterColumnTypeAction(r *googlesql.ASTAlterColumnTypeAction) *AlterColumnTypeAction {
	if r == nil {
		return nil
	}
	return &AlterColumnTypeAction{baseNode[*googlesql.ASTAlterColumnTypeAction]{raw: r}}
}

func (n *AlterColumnTypeAction) isAlterAction() {}

func (n *AlterColumnTypeAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *AlterColumnTypeAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

func (n *AlterColumnTypeAction) Schema() Node { return Wrap(must(n.raw.Schema())) }

// AlterConstraintEnforcementAction wraps *googlesql.ASTAlterConstraintEnforcementAction.
type AlterConstraintEnforcementAction struct {
	baseNode[*googlesql.ASTAlterConstraintEnforcementAction]
}

func newAlterConstraintEnforcementAction(r *googlesql.ASTAlterConstraintEnforcementAction) *AlterConstraintEnforcementAction {
	if r == nil {
		return nil
	}
	return &AlterConstraintEnforcementAction{baseNode[*googlesql.ASTAlterConstraintEnforcementAction]{raw: r}}
}

func (n *AlterConstraintEnforcementAction) isAlterAction() {}

// AlterConstraintSetOptionsAction wraps *googlesql.ASTAlterConstraintSetOptionsAction.
type AlterConstraintSetOptionsAction struct {
	baseNode[*googlesql.ASTAlterConstraintSetOptionsAction]
}

func newAlterConstraintSetOptionsAction(r *googlesql.ASTAlterConstraintSetOptionsAction) *AlterConstraintSetOptionsAction {
	if r == nil {
		return nil
	}
	return &AlterConstraintSetOptionsAction{baseNode[*googlesql.ASTAlterConstraintSetOptionsAction]{raw: r}}
}

func (n *AlterConstraintSetOptionsAction) isAlterAction() {}

// DropColumnAction wraps *googlesql.ASTDropColumnAction.
type DropColumnAction struct {
	baseNode[*googlesql.ASTDropColumnAction]
}

func newDropColumnAction(r *googlesql.ASTDropColumnAction) *DropColumnAction {
	if r == nil {
		return nil
	}
	return &DropColumnAction{baseNode[*googlesql.ASTDropColumnAction]{raw: r}}
}

func (n *DropColumnAction) isAlterAction() {}

func (n *DropColumnAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropColumnAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

// DropConstraintAction wraps *googlesql.ASTDropConstraintAction.
type DropConstraintAction struct {
	baseNode[*googlesql.ASTDropConstraintAction]
}

func newDropConstraintAction(r *googlesql.ASTDropConstraintAction) *DropConstraintAction {
	if r == nil {
		return nil
	}
	return &DropConstraintAction{baseNode[*googlesql.ASTDropConstraintAction]{raw: r}}
}

func (n *DropConstraintAction) isAlterAction() {}

func (n *DropConstraintAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropConstraintAction) ConstraintName() *Identifier {
	return newIdentifier(must(n.raw.ConstraintName()))
}

// DropPrimaryKeyAction wraps *googlesql.ASTDropPrimaryKeyAction.
type DropPrimaryKeyAction struct {
	baseNode[*googlesql.ASTDropPrimaryKeyAction]
}

func newDropPrimaryKeyAction(r *googlesql.ASTDropPrimaryKeyAction) *DropPrimaryKeyAction {
	if r == nil {
		return nil
	}
	return &DropPrimaryKeyAction{baseNode[*googlesql.ASTDropPrimaryKeyAction]{raw: r}}
}

func (n *DropPrimaryKeyAction) isAlterAction() {}

func (n *DropPrimaryKeyAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

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

// RenameColumnAction wraps *googlesql.ASTRenameColumnAction.
type RenameColumnAction struct {
	baseNode[*googlesql.ASTRenameColumnAction]
}

func newRenameColumnAction(r *googlesql.ASTRenameColumnAction) *RenameColumnAction {
	if r == nil {
		return nil
	}
	return &RenameColumnAction{baseNode[*googlesql.ASTRenameColumnAction]{raw: r}}
}

func (n *RenameColumnAction) isAlterAction() {}

func (n *RenameColumnAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *RenameColumnAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

func (n *RenameColumnAction) NewColumnName() *Identifier {
	return newIdentifier(must(n.raw.NewColumnName()))
}

// RenameToClause wraps *googlesql.ASTRenameToClause.
type RenameToClause struct {
	baseNode[*googlesql.ASTRenameToClause]
}

func newRenameToClause(r *googlesql.ASTRenameToClause) *RenameToClause {
	if r == nil {
		return nil
	}
	return &RenameToClause{baseNode[*googlesql.ASTRenameToClause]{raw: r}}
}

func (n *RenameToClause) isAlterAction() {}

func (n *RenameToClause) NewName() *PathExpression {
	return newPathExpression(must(n.raw.NewName()))
}

// SetCollateClause wraps *googlesql.ASTSetCollateClause.
type SetCollateClause struct {
	baseNode[*googlesql.ASTSetCollateClause]
}

func newSetCollateClause(r *googlesql.ASTSetCollateClause) *SetCollateClause {
	if r == nil {
		return nil
	}
	return &SetCollateClause{baseNode[*googlesql.ASTSetCollateClause]{raw: r}}
}

func (n *SetCollateClause) isAlterAction() {}

func (n *SetCollateClause) Collate() *Collate { return newCollate(must(n.raw.Collate())) }

// SetOptionsAction wraps *googlesql.ASTSetOptionsAction.
type SetOptionsAction struct {
	baseNode[*googlesql.ASTSetOptionsAction]
}

func newSetOptionsAction(r *googlesql.ASTSetOptionsAction) *SetOptionsAction {
	if r == nil {
		return nil
	}
	return &SetOptionsAction{baseNode[*googlesql.ASTSetOptionsAction]{raw: r}}
}

func (n *SetOptionsAction) isAlterAction() {}

func (n *SetOptionsAction) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// AlterColumnDropGeneratedAction wraps *googlesql.ASTAlterColumnDropGeneratedAction.
type AlterColumnDropGeneratedAction struct {
	baseNode[*googlesql.ASTAlterColumnDropGeneratedAction]
}

func newAlterColumnDropGeneratedAction(r *googlesql.ASTAlterColumnDropGeneratedAction) *AlterColumnDropGeneratedAction {
	if r == nil {
		return nil
	}
	return &AlterColumnDropGeneratedAction{baseNode[*googlesql.ASTAlterColumnDropGeneratedAction]{raw: r}}
}
func (n *AlterColumnDropGeneratedAction) isAlterAction() {}
func (n *AlterColumnDropGeneratedAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}
func (n *AlterColumnDropGeneratedAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

// AlterColumnSetGeneratedAction wraps *googlesql.ASTAlterColumnSetGeneratedAction.
type AlterColumnSetGeneratedAction struct {
	baseNode[*googlesql.ASTAlterColumnSetGeneratedAction]
}

func newAlterColumnSetGeneratedAction(r *googlesql.ASTAlterColumnSetGeneratedAction) *AlterColumnSetGeneratedAction {
	if r == nil {
		return nil
	}
	return &AlterColumnSetGeneratedAction{baseNode[*googlesql.ASTAlterColumnSetGeneratedAction]{raw: r}}
}
func (n *AlterColumnSetGeneratedAction) isAlterAction() {}
func (n *AlterColumnSetGeneratedAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}

func (n *AlterColumnSetGeneratedAction) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}
func (n *AlterColumnSetGeneratedAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

// AlterSubEntityAction wraps *googlesql.ASTAlterSubEntityAction.
type AlterSubEntityAction struct {
	baseNode[*googlesql.ASTAlterSubEntityAction]
}

func newAlterSubEntityAction(r *googlesql.ASTAlterSubEntityAction) *AlterSubEntityAction {
	if r == nil {
		return nil
	}
	return &AlterSubEntityAction{baseNode[*googlesql.ASTAlterSubEntityAction]{raw: r}}
}
func (n *AlterSubEntityAction) isAlterAction()          {}
func (n *AlterSubEntityAction) Type() *Identifier       { return newIdentifier(must(n.raw.Type())) }
func (n *AlterSubEntityAction) Name() *Identifier       { return newIdentifier(must(n.raw.Name())) }
func (n *AlterSubEntityAction) Action() AlterActionNode { return wrapAlterAction(must(n.raw.Action())) }
func (n *AlterSubEntityAction) IsIfExists() bool        { return must(n.raw.IsIfExists()) }

// AddColumnIdentifierAction wraps *googlesql.ASTAddColumnIdentifierAction.
type AddColumnIdentifierAction struct {
	baseNode[*googlesql.ASTAddColumnIdentifierAction]
}

func newAddColumnIdentifierAction(r *googlesql.ASTAddColumnIdentifierAction) *AddColumnIdentifierAction {
	if r == nil {
		return nil
	}
	return &AddColumnIdentifierAction{baseNode[*googlesql.ASTAddColumnIdentifierAction]{raw: r}}
}
func (n *AddColumnIdentifierAction) isAlterAction() {}
func (n *AddColumnIdentifierAction) ColumnName() *Identifier {
	return newIdentifier(must(n.raw.ColumnName()))
}
func (n *AddColumnIdentifierAction) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *AddColumnIdentifierAction) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// AddSubEntityAction wraps *googlesql.ASTAddSubEntityAction.
type AddSubEntityAction struct {
	baseNode[*googlesql.ASTAddSubEntityAction]
}

func newAddSubEntityAction(r *googlesql.ASTAddSubEntityAction) *AddSubEntityAction {
	if r == nil {
		return nil
	}
	return &AddSubEntityAction{baseNode[*googlesql.ASTAddSubEntityAction]{raw: r}}
}
func (n *AddSubEntityAction) isAlterAction()    {}
func (n *AddSubEntityAction) Type() *Identifier { return newIdentifier(must(n.raw.Type())) }
func (n *AddSubEntityAction) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }
func (n *AddSubEntityAction) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}
func (n *AddSubEntityAction) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

// AddToRestricteeListClause wraps *googlesql.ASTAddToRestricteeListClause.
type AddToRestricteeListClause struct {
	baseNode[*googlesql.ASTAddToRestricteeListClause]
}

func newAddToRestricteeListClause(r *googlesql.ASTAddToRestricteeListClause) *AddToRestricteeListClause {
	if r == nil {
		return nil
	}
	return &AddToRestricteeListClause{baseNode[*googlesql.ASTAddToRestricteeListClause]{raw: r}}
}
func (n *AddToRestricteeListClause) isAlterAction() {}
func (n *AddToRestricteeListClause) RestricteeList() *GranteeList {
	return newGranteeList(must(n.raw.RestricteeList()))
}
func (n *AddToRestricteeListClause) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

// AddTTLAction wraps *googlesql.ASTAddTtlAction.
type AddTTLAction struct {
	baseNode[*googlesql.ASTAddTtlAction]
}

func newAddTTLAction(r *googlesql.ASTAddTtlAction) *AddTTLAction {
	if r == nil {
		return nil
	}
	return &AddTTLAction{baseNode[*googlesql.ASTAddTtlAction]{raw: r}}
}
func (n *AddTTLAction) isAlterAction()             {}
func (n *AddTTLAction) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }
func (n *AddTTLAction) IsIfNotExists() bool        { return must(n.raw.IsIfNotExists()) }

// DropSubEntityAction wraps *googlesql.ASTDropSubEntityAction.
type DropSubEntityAction struct {
	baseNode[*googlesql.ASTDropSubEntityAction]
}

func newDropSubEntityAction(r *googlesql.ASTDropSubEntityAction) *DropSubEntityAction {
	if r == nil {
		return nil
	}
	return &DropSubEntityAction{baseNode[*googlesql.ASTDropSubEntityAction]{raw: r}}
}
func (n *DropSubEntityAction) isAlterAction()    {}
func (n *DropSubEntityAction) Type() *Identifier { return newIdentifier(must(n.raw.Type())) }
func (n *DropSubEntityAction) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }
func (n *DropSubEntityAction) IsIfExists() bool  { return must(n.raw.IsIfExists()) }

// DropTTLAction wraps *googlesql.ASTDropTtlAction.
type DropTTLAction struct {
	baseNode[*googlesql.ASTDropTtlAction]
}

func newDropTTLAction(r *googlesql.ASTDropTtlAction) *DropTTLAction {
	if r == nil {
		return nil
	}
	return &DropTTLAction{baseNode[*googlesql.ASTDropTtlAction]{raw: r}}
}
func (n *DropTTLAction) isAlterAction()   {}
func (n *DropTTLAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

// SetAsAction wraps *googlesql.ASTSetAsAction.
type SetAsAction struct {
	baseNode[*googlesql.ASTSetAsAction]
}

func newSetAsAction(r *googlesql.ASTSetAsAction) *SetAsAction {
	if r == nil {
		return nil
	}
	return &SetAsAction{baseNode[*googlesql.ASTSetAsAction]{raw: r}}
}
func (n *SetAsAction) isAlterAction()           {}
func (n *SetAsAction) JSONBody() *JSONLiteral   { return newJSONLiteral(must(n.raw.JsonBody())) }
func (n *SetAsAction) TextBody() *StringLiteral { return newStringLiteral(must(n.raw.TextBody())) }

// RebuildAction wraps *googlesql.ASTRebuildAction.
type RebuildAction struct {
	baseNode[*googlesql.ASTRebuildAction]
}

func newRebuildAction(r *googlesql.ASTRebuildAction) *RebuildAction {
	if r == nil {
		return nil
	}
	return &RebuildAction{baseNode[*googlesql.ASTRebuildAction]{raw: r}}
}
func (n *RebuildAction) isAlterAction() {}

// RemoveFromRestricteeListClause wraps *googlesql.ASTRemoveFromRestricteeListClause.
type RemoveFromRestricteeListClause struct {
	baseNode[*googlesql.ASTRemoveFromRestricteeListClause]
}

func newRemoveFromRestricteeListClause(r *googlesql.ASTRemoveFromRestricteeListClause) *RemoveFromRestricteeListClause {
	if r == nil {
		return nil
	}
	return &RemoveFromRestricteeListClause{baseNode[*googlesql.ASTRemoveFromRestricteeListClause]{raw: r}}
}
func (n *RemoveFromRestricteeListClause) isAlterAction() {}
func (n *RemoveFromRestricteeListClause) RestricteeList() *GranteeList {
	return newGranteeList(must(n.raw.RestricteeList()))
}
func (n *RemoveFromRestricteeListClause) IsIfExists() bool { return must(n.raw.IsIfExists()) }

// ReplaceTTLAction wraps *googlesql.ASTReplaceTtlAction.
type ReplaceTTLAction struct {
	baseNode[*googlesql.ASTReplaceTtlAction]
}

func newReplaceTTLAction(r *googlesql.ASTReplaceTtlAction) *ReplaceTTLAction {
	if r == nil {
		return nil
	}
	return &ReplaceTTLAction{baseNode[*googlesql.ASTReplaceTtlAction]{raw: r}}
}
func (n *ReplaceTTLAction) isAlterAction()             {}
func (n *ReplaceTTLAction) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }
func (n *ReplaceTTLAction) IsIfExists() bool           { return must(n.raw.IsIfExists()) }

// RestrictToClause wraps *googlesql.ASTRestrictToClause.
type RestrictToClause struct {
	baseNode[*googlesql.ASTRestrictToClause]
}

func newRestrictToClause(r *googlesql.ASTRestrictToClause) *RestrictToClause {
	if r == nil {
		return nil
	}
	return &RestrictToClause{baseNode[*googlesql.ASTRestrictToClause]{raw: r}}
}
func (n *RestrictToClause) isAlterAction() {}
func (n *RestrictToClause) RestricteeList() *GranteeList {
	return newGranteeList(must(n.raw.RestricteeList()))
}

// RevokeFromClause wraps *googlesql.ASTRevokeFromClause.
type RevokeFromClause struct {
	baseNode[*googlesql.ASTRevokeFromClause]
}

func newRevokeFromClause(r *googlesql.ASTRevokeFromClause) *RevokeFromClause {
	if r == nil {
		return nil
	}
	return &RevokeFromClause{baseNode[*googlesql.ASTRevokeFromClause]{raw: r}}
}
func (n *RevokeFromClause) isAlterAction() {}
func (n *RevokeFromClause) RevokeFromList() *GranteeList {
	return newGranteeList(must(n.raw.RevokeFromList()))
}
func (n *RevokeFromClause) IsRevokeFromAll() bool { return must(n.raw.IsRevokeFromAll()) }
