package sql

import "github.com/goccy/go-googlesql"

// ─── Column schema nodes ──────────────────────────────────────────────────────

// SimpleColumnSchema wraps *googlesql.ASTSimpleColumnSchema.
type SimpleColumnSchema struct {
	baseNode[*googlesql.ASTSimpleColumnSchema]
}

func newASTSimpleColumnSchema(r *googlesql.ASTSimpleColumnSchema) *SimpleColumnSchema {
	if r == nil {
		return nil
	}
	return &SimpleColumnSchema{baseNode[*googlesql.ASTSimpleColumnSchema]{raw: r}}
}

func (n *SimpleColumnSchema) TypeName() *PathExpression {
	return newASTPathExpression(must(n.raw.TypeName()))
}

func (n *SimpleColumnSchema) Attributes() *ColumnAttributeList {
	return newASTColumnAttributeList(must(n.raw.Attributes()))
}

func (n *SimpleColumnSchema) Collate() *Collate {
	return newASTCollate(must(n.raw.Collate()))
}

func (n *SimpleColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *SimpleColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newASTGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *SimpleColumnSchema) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

func (n *SimpleColumnSchema) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

// ArrayColumnSchema wraps *googlesql.ASTArrayColumnSchema.
type ArrayColumnSchema struct {
	baseNode[*googlesql.ASTArrayColumnSchema]
}

func newASTArrayColumnSchema(r *googlesql.ASTArrayColumnSchema) *ArrayColumnSchema {
	if r == nil {
		return nil
	}
	return &ArrayColumnSchema{baseNode[*googlesql.ASTArrayColumnSchema]{raw: r}}
}

func (n *ArrayColumnSchema) ElementSchema() Node {
	return Wrap(must(n.raw.ElementSchema()))
}

func (n *ArrayColumnSchema) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

func (n *ArrayColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newASTGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *ArrayColumnSchema) Attributes() *ColumnAttributeList {
	return newASTColumnAttributeList(must(n.raw.Attributes()))
}

func (n *ArrayColumnSchema) Collate() *Collate {
	return newASTCollate(must(n.raw.Collate()))
}

func (n *ArrayColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *ArrayColumnSchema) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// StructColumnSchema wraps *googlesql.ASTStructColumnSchema.
type StructColumnSchema struct {
	baseNode[*googlesql.ASTStructColumnSchema]
}

func newASTStructColumnSchema(r *googlesql.ASTStructColumnSchema) *StructColumnSchema {
	if r == nil {
		return nil
	}
	return &StructColumnSchema{baseNode[*googlesql.ASTStructColumnSchema]{raw: r}}
}

func (n *StructColumnSchema) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

func (n *StructColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newASTGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *StructColumnSchema) Attributes() *ColumnAttributeList {
	return newASTColumnAttributeList(must(n.raw.Attributes()))
}

func (n *StructColumnSchema) Collate() *Collate {
	return newASTCollate(must(n.raw.Collate()))
}

func (n *StructColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *StructColumnSchema) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// StructFields returns []*StructColumnField children.
func (n *StructColumnSchema) StructFields() []*StructColumnField {
	var result []*StructColumnField
	for _, c := range n.Children() {
		if f, ok := c.(*StructColumnField); ok {
			result = append(result, f)
		}
	}
	return result
}

// StructColumnField wraps *googlesql.ASTStructColumnField.
type StructColumnField struct {
	baseNode[*googlesql.ASTStructColumnField]
}

func newASTStructColumnField(r *googlesql.ASTStructColumnField) *StructColumnField {
	if r == nil {
		return nil
	}
	return &StructColumnField{baseNode[*googlesql.ASTStructColumnField]{raw: r}}
}
func (n *StructColumnField) Name() *Identifier { return newASTIdentifier(must(n.raw.Name())) }
func (n *StructColumnField) Schema() Node {
	return Wrap(must(n.raw.Schema()))
}

// ColumnAttributeList wraps *googlesql.ASTColumnAttributeList.
type ColumnAttributeList struct {
	baseNode[*googlesql.ASTColumnAttributeList]
}

func newASTColumnAttributeList(r *googlesql.ASTColumnAttributeList) *ColumnAttributeList {
	if r == nil {
		return nil
	}
	return &ColumnAttributeList{baseNode[*googlesql.ASTColumnAttributeList]{raw: r}}
}

// Values returns all column attributes.
func (n *ColumnAttributeList) Values() []ColumnAttributeNode {
	count := n.NumChildren()
	result := make([]ColumnAttributeNode, 0, count)
	for i := range count {
		c := must(n.raw.Values(int32(i)))
		if c == nil {
			break
		}
		result = append(result, wrapColumnAttribute(c))
	}
	return result
}

// NotNullColumnAttribute wraps *googlesql.ASTNotNullColumnAttribute.
type NotNullColumnAttribute struct {
	baseNode[*googlesql.ASTNotNullColumnAttribute]
}

func newASTNotNullColumnAttribute(r *googlesql.ASTNotNullColumnAttribute) *NotNullColumnAttribute {
	if r == nil {
		return nil
	}
	return &NotNullColumnAttribute{baseNode[*googlesql.ASTNotNullColumnAttribute]{raw: r}}
}
func (n *NotNullColumnAttribute) isColumnAttribute() {}

// PrimaryKeyColumnAttribute wraps *googlesql.ASTPrimaryKeyColumnAttribute.
type PrimaryKeyColumnAttribute struct {
	baseNode[*googlesql.ASTPrimaryKeyColumnAttribute]
}

func newASTPrimaryKeyColumnAttribute(r *googlesql.ASTPrimaryKeyColumnAttribute) *PrimaryKeyColumnAttribute {
	if r == nil {
		return nil
	}
	return &PrimaryKeyColumnAttribute{baseNode[*googlesql.ASTPrimaryKeyColumnAttribute]{raw: r}}
}
func (n *PrimaryKeyColumnAttribute) isColumnAttribute() {}
func (n *PrimaryKeyColumnAttribute) Enforced() bool     { return must(n.raw.Enforced()) }

// ForeignKeyColumnAttribute wraps *googlesql.ASTForeignKeyColumnAttribute.
type ForeignKeyColumnAttribute struct {
	baseNode[*googlesql.ASTForeignKeyColumnAttribute]
}

func newASTForeignKeyColumnAttribute(r *googlesql.ASTForeignKeyColumnAttribute) *ForeignKeyColumnAttribute {
	if r == nil {
		return nil
	}
	return &ForeignKeyColumnAttribute{baseNode[*googlesql.ASTForeignKeyColumnAttribute]{raw: r}}
}
func (n *ForeignKeyColumnAttribute) isColumnAttribute() {}
func (n *ForeignKeyColumnAttribute) ConstraintName() *Identifier {
	return newASTIdentifier(must(n.raw.ConstraintName()))
}

func (n *ForeignKeyColumnAttribute) Reference() *ForeignKeyReference {
	return newASTForeignKeyReference(must(n.raw.Reference()))
}

// HiddenColumnAttribute wraps *googlesql.ASTHiddenColumnAttribute.
type HiddenColumnAttribute struct {
	baseNode[*googlesql.ASTHiddenColumnAttribute]
}

func newASTHiddenColumnAttribute(r *googlesql.ASTHiddenColumnAttribute) *HiddenColumnAttribute {
	if r == nil {
		return nil
	}
	return &HiddenColumnAttribute{baseNode[*googlesql.ASTHiddenColumnAttribute]{raw: r}}
}
func (n *HiddenColumnAttribute) isColumnAttribute() {}

// ColumnDefinition wraps *googlesql.ASTColumnDefinition.
type ColumnDefinition struct {
	baseNode[*googlesql.ASTColumnDefinition]
}

func newASTColumnDefinition(r *googlesql.ASTColumnDefinition) *ColumnDefinition {
	if r == nil {
		return nil
	}
	return &ColumnDefinition{baseNode[*googlesql.ASTColumnDefinition]{raw: r}}
}
func (n *ColumnDefinition) isTableElement()   {}
func (n *ColumnDefinition) Name() *Identifier { return newASTIdentifier(must(n.raw.Name())) }

// Schema returns the column schema as Node.
func (n *ColumnDefinition) Schema() Node { return Wrap(must(n.raw.Schema())) }

// ColumnPosition wraps *googlesql.ASTColumnPosition.
type ColumnPosition struct {
	baseNode[*googlesql.ASTColumnPosition]
}

func newASTColumnPosition(r *googlesql.ASTColumnPosition) *ColumnPosition {
	if r == nil {
		return nil
	}
	return &ColumnPosition{baseNode[*googlesql.ASTColumnPosition]{raw: r}}
}

func (n *ColumnPosition) Identifier() *Identifier {
	return newASTIdentifier(must(n.raw.Identifier()))
}
func (n *ColumnPosition) RelativePosition() RelativePosition { return must(n.raw.Type()) }

// GeneratedColumnInfo wraps *googlesql.ASTGeneratedColumnInfo.
type GeneratedColumnInfo struct {
	baseNode[*googlesql.ASTGeneratedColumnInfo]
}

func newASTGeneratedColumnInfo(r *googlesql.ASTGeneratedColumnInfo) *GeneratedColumnInfo {
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

// TableElementList wraps *googlesql.ASTTableElementList.
type TableElementList struct {
	baseNode[*googlesql.ASTTableElementList]
}

func newASTTableElementList(r *googlesql.ASTTableElementList) *TableElementList {
	if r == nil {
		return nil
	}
	return &TableElementList{baseNode[*googlesql.ASTTableElementList]{raw: r}}
}

// Elements returns all table element nodes.
func (n *TableElementList) Elements() []TableElementNode {
	count := n.NumChildren()
	result := make([]TableElementNode, 0, count)
	for i := range count {
		c := must(n.raw.Elements(int32(i)))
		if c == nil {
			break
		}
		result = append(result, wrapTableElement(c))
	}
	return result
}

// ─── Constraint nodes ─────────────────────────────────────────────────────────

// TableConstraint wraps *googlesql.ASTTableConstraint.
type TableConstraint struct {
	baseNode[*googlesql.ASTTableConstraint]
}

func newASTTableConstraint(r *googlesql.ASTTableConstraint) *TableConstraint {
	if r == nil {
		return nil
	}
	return &TableConstraint{baseNode[*googlesql.ASTTableConstraint]{raw: r}}
}
func (n *TableConstraint) isTableElement() {}

func (n *TableConstraint) ConstraintName() *Identifier {
	return newASTIdentifier(must(n.raw.ConstraintName()))
}

// PrimaryKey wraps *googlesql.ASTPrimaryKey.
type PrimaryKey struct {
	baseNode[*googlesql.ASTPrimaryKey]
}

func newASTPrimaryKey(r *googlesql.ASTPrimaryKey) *PrimaryKey {
	if r == nil {
		return nil
	}
	return &PrimaryKey{baseNode[*googlesql.ASTPrimaryKey]{raw: r}}
}
func (n *PrimaryKey) isTableElement() {}

func (n *PrimaryKey) ConstraintName() *Identifier {
	return newASTIdentifier(must(n.raw.ConstraintName()))
}

func (n *PrimaryKey) ElementList() *PrimaryKeyElementList {
	return newASTPrimaryKeyElementList(must(n.raw.ElementList()))
}
func (n *PrimaryKey) Enforced() bool { return must(n.raw.Enforced()) }
func (n *PrimaryKey) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// PrimaryKeyElementList wraps *googlesql.ASTPrimaryKeyElementList.
type PrimaryKeyElementList struct {
	baseNode[*googlesql.ASTPrimaryKeyElementList]
}

func newASTPrimaryKeyElementList(r *googlesql.ASTPrimaryKeyElementList) *PrimaryKeyElementList {
	if r == nil {
		return nil
	}
	return &PrimaryKeyElementList{baseNode[*googlesql.ASTPrimaryKeyElementList]{raw: r}}
}

// Elements returns all *PrimaryKeyElement children.
func (n *PrimaryKeyElementList) Elements() []*PrimaryKeyElement {
	var result []*PrimaryKeyElement
	for _, c := range n.Children() {
		if e, ok := c.(*PrimaryKeyElement); ok {
			result = append(result, e)
		}
	}
	return result
}

// PrimaryKeyElement wraps *googlesql.ASTPrimaryKeyElement.
type PrimaryKeyElement struct {
	baseNode[*googlesql.ASTPrimaryKeyElement]
}

func newASTPrimaryKeyElement(r *googlesql.ASTPrimaryKeyElement) *PrimaryKeyElement {
	if r == nil {
		return nil
	}
	return &PrimaryKeyElement{baseNode[*googlesql.ASTPrimaryKeyElement]{raw: r}}
}

func (n *PrimaryKeyElement) Column() *Identifier {
	return newASTIdentifier(must(n.raw.Column()))
}
func (n *PrimaryKeyElement) OrderingSpec() OrderingSpec { return must(n.raw.OrderingSpec()) }
func (n *PrimaryKeyElement) NullOrder() *NullOrder {
	return newASTNullOrder(must(n.raw.NullOrder()))
}

// ForeignKey wraps *googlesql.ASTForeignKey.
type ForeignKey struct {
	baseNode[*googlesql.ASTForeignKey]
}

func newASTForeignKey(r *googlesql.ASTForeignKey) *ForeignKey {
	if r == nil {
		return nil
	}
	return &ForeignKey{baseNode[*googlesql.ASTForeignKey]{raw: r}}
}
func (n *ForeignKey) isTableElement() {}

func (n *ForeignKey) ConstraintName() *Identifier {
	return newASTIdentifier(must(n.raw.ConstraintName()))
}

func (n *ForeignKey) ColumnList() *ColumnList {
	return newASTColumnList(must(n.raw.ColumnList()))
}

func (n *ForeignKey) Reference() *ForeignKeyReference {
	return newASTForeignKeyReference(must(n.raw.Reference()))
}

func (n *ForeignKey) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// ForeignKeyReference wraps *googlesql.ASTForeignKeyReference.
type ForeignKeyReference struct {
	baseNode[*googlesql.ASTForeignKeyReference]
}

func newASTForeignKeyReference(r *googlesql.ASTForeignKeyReference) *ForeignKeyReference {
	if r == nil {
		return nil
	}
	return &ForeignKeyReference{baseNode[*googlesql.ASTForeignKeyReference]{raw: r}}
}

func (n *ForeignKeyReference) TableName() *PathExpression {
	return newASTPathExpression(must(n.raw.TableName()))
}

func (n *ForeignKeyReference) ColumnList() *ColumnList {
	return newASTColumnList(must(n.raw.ColumnList()))
}
func (n *ForeignKeyReference) Enforced() bool         { return must(n.raw.Enforced()) }
func (n *ForeignKeyReference) Match() ForeignKeyMatch { return must(n.raw.Match()) }

// ─── ALTER ACTION nodes ───────────────────────────────────────────────────────

// AlterActionList wraps *googlesql.ASTAlterActionList.
type AlterActionList struct {
	baseNode[*googlesql.ASTAlterActionList]
}

func newASTAlterActionList(r *googlesql.ASTAlterActionList) *AlterActionList {
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

// AddColumnAction wraps *googlesql.ASTAddColumnAction.
type AddColumnAction struct {
	baseNode[*googlesql.ASTAddColumnAction]
}

func newASTAddColumnAction(r *googlesql.ASTAddColumnAction) *AddColumnAction {
	if r == nil {
		return nil
	}
	return &AddColumnAction{baseNode[*googlesql.ASTAddColumnAction]{raw: r}}
}
func (n *AddColumnAction) isAlterAction()      {}
func (n *AddColumnAction) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *AddColumnAction) ColumnDefinition() *ColumnDefinition {
	return newASTColumnDefinition(must(n.raw.ColumnDefinition()))
}

// AddConstraintAction wraps *googlesql.ASTAddConstraintAction.
type AddConstraintAction struct {
	baseNode[*googlesql.ASTAddConstraintAction]
}

func newASTAddConstraintAction(r *googlesql.ASTAddConstraintAction) *AddConstraintAction {
	if r == nil {
		return nil
	}
	return &AddConstraintAction{baseNode[*googlesql.ASTAddConstraintAction]{raw: r}}
}
func (n *AddConstraintAction) isAlterAction()      {}
func (n *AddConstraintAction) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *AddConstraintAction) Constraint() Node    { return Wrap(must(n.raw.Constraint())) }

// AlterColumnDropDefaultAction wraps *googlesql.ASTAlterColumnDropDefaultAction.
type AlterColumnDropDefaultAction struct {
	baseNode[*googlesql.ASTAlterColumnDropDefaultAction]
}

func newASTAlterColumnDropDefaultAction(r *googlesql.ASTAlterColumnDropDefaultAction) *AlterColumnDropDefaultAction {
	if r == nil {
		return nil
	}
	return &AlterColumnDropDefaultAction{baseNode[*googlesql.ASTAlterColumnDropDefaultAction]{raw: r}}
}
func (n *AlterColumnDropDefaultAction) isAlterAction()   {}
func (n *AlterColumnDropDefaultAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *AlterColumnDropDefaultAction) ColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.ColumnName()))
}

// AlterColumnDropNotNullAction wraps *googlesql.ASTAlterColumnDropNotNullAction.
type AlterColumnDropNotNullAction struct {
	baseNode[*googlesql.ASTAlterColumnDropNotNullAction]
}

func newASTAlterColumnDropNotNullAction(r *googlesql.ASTAlterColumnDropNotNullAction) *AlterColumnDropNotNullAction {
	if r == nil {
		return nil
	}
	return &AlterColumnDropNotNullAction{baseNode[*googlesql.ASTAlterColumnDropNotNullAction]{raw: r}}
}
func (n *AlterColumnDropNotNullAction) isAlterAction()   {}
func (n *AlterColumnDropNotNullAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *AlterColumnDropNotNullAction) ColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.ColumnName()))
}

// AlterColumnOptionsAction wraps *googlesql.ASTAlterColumnOptionsAction.
type AlterColumnOptionsAction struct {
	baseNode[*googlesql.ASTAlterColumnOptionsAction]
}

func newASTAlterColumnOptionsAction(r *googlesql.ASTAlterColumnOptionsAction) *AlterColumnOptionsAction {
	if r == nil {
		return nil
	}
	return &AlterColumnOptionsAction{baseNode[*googlesql.ASTAlterColumnOptionsAction]{raw: r}}
}
func (n *AlterColumnOptionsAction) isAlterAction()   {}
func (n *AlterColumnOptionsAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *AlterColumnOptionsAction) ColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.ColumnName()))
}

func (n *AlterColumnOptionsAction) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// AlterColumnSetDefaultAction wraps *googlesql.ASTAlterColumnSetDefaultAction.
type AlterColumnSetDefaultAction struct {
	baseNode[*googlesql.ASTAlterColumnSetDefaultAction]
}

func newASTAlterColumnSetDefaultAction(r *googlesql.ASTAlterColumnSetDefaultAction) *AlterColumnSetDefaultAction {
	if r == nil {
		return nil
	}
	return &AlterColumnSetDefaultAction{baseNode[*googlesql.ASTAlterColumnSetDefaultAction]{raw: r}}
}
func (n *AlterColumnSetDefaultAction) isAlterAction()   {}
func (n *AlterColumnSetDefaultAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *AlterColumnSetDefaultAction) ColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.ColumnName()))
}

func (n *AlterColumnSetDefaultAction) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

// AlterColumnTypeAction wraps *googlesql.ASTAlterColumnTypeAction.
type AlterColumnTypeAction struct {
	baseNode[*googlesql.ASTAlterColumnTypeAction]
}

func newASTAlterColumnTypeAction(r *googlesql.ASTAlterColumnTypeAction) *AlterColumnTypeAction {
	if r == nil {
		return nil
	}
	return &AlterColumnTypeAction{baseNode[*googlesql.ASTAlterColumnTypeAction]{raw: r}}
}
func (n *AlterColumnTypeAction) isAlterAction()   {}
func (n *AlterColumnTypeAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *AlterColumnTypeAction) ColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.ColumnName()))
}
func (n *AlterColumnTypeAction) Schema() Node { return Wrap(must(n.raw.Schema())) }

// AlterConstraintEnforcementAction wraps *googlesql.ASTAlterConstraintEnforcementAction.
type AlterConstraintEnforcementAction struct {
	baseNode[*googlesql.ASTAlterConstraintEnforcementAction]
}

func newASTAlterConstraintEnforcementAction(r *googlesql.ASTAlterConstraintEnforcementAction) *AlterConstraintEnforcementAction {
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

func newASTAlterConstraintSetOptionsAction(r *googlesql.ASTAlterConstraintSetOptionsAction) *AlterConstraintSetOptionsAction {
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

func newASTDropColumnAction(r *googlesql.ASTDropColumnAction) *DropColumnAction {
	if r == nil {
		return nil
	}
	return &DropColumnAction{baseNode[*googlesql.ASTDropColumnAction]{raw: r}}
}
func (n *DropColumnAction) isAlterAction()   {}
func (n *DropColumnAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropColumnAction) ColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.ColumnName()))
}

// DropConstraintAction wraps *googlesql.ASTDropConstraintAction.
type DropConstraintAction struct {
	baseNode[*googlesql.ASTDropConstraintAction]
}

func newASTDropConstraintAction(r *googlesql.ASTDropConstraintAction) *DropConstraintAction {
	if r == nil {
		return nil
	}
	return &DropConstraintAction{baseNode[*googlesql.ASTDropConstraintAction]{raw: r}}
}
func (n *DropConstraintAction) isAlterAction()   {}
func (n *DropConstraintAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropConstraintAction) ConstraintName() *Identifier {
	return newASTIdentifier(must(n.raw.ConstraintName()))
}

// DropPrimaryKeyAction wraps *googlesql.ASTDropPrimaryKeyAction.
type DropPrimaryKeyAction struct {
	baseNode[*googlesql.ASTDropPrimaryKeyAction]
}

func newASTDropPrimaryKeyAction(r *googlesql.ASTDropPrimaryKeyAction) *DropPrimaryKeyAction {
	if r == nil {
		return nil
	}
	return &DropPrimaryKeyAction{baseNode[*googlesql.ASTDropPrimaryKeyAction]{raw: r}}
}
func (n *DropPrimaryKeyAction) isAlterAction()   {}
func (n *DropPrimaryKeyAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }

// RenameColumnAction wraps *googlesql.ASTRenameColumnAction.
type RenameColumnAction struct {
	baseNode[*googlesql.ASTRenameColumnAction]
}

func newASTRenameColumnAction(r *googlesql.ASTRenameColumnAction) *RenameColumnAction {
	if r == nil {
		return nil
	}
	return &RenameColumnAction{baseNode[*googlesql.ASTRenameColumnAction]{raw: r}}
}
func (n *RenameColumnAction) isAlterAction()   {}
func (n *RenameColumnAction) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *RenameColumnAction) ColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.ColumnName()))
}

func (n *RenameColumnAction) NewColumnName() *Identifier {
	return newASTIdentifier(must(n.raw.NewColumnName()))
}

// RenameToClause wraps *googlesql.ASTRenameToClause.
type RenameToClause struct {
	baseNode[*googlesql.ASTRenameToClause]
}

func newASTRenameToClause(r *googlesql.ASTRenameToClause) *RenameToClause {
	if r == nil {
		return nil
	}
	return &RenameToClause{baseNode[*googlesql.ASTRenameToClause]{raw: r}}
}
func (n *RenameToClause) isAlterAction() {}

func (n *RenameToClause) NewName() *PathExpression {
	return newASTPathExpression(must(n.raw.NewName()))
}

// SetCollateClause wraps *googlesql.ASTSetCollateClause.
type SetCollateClause struct {
	baseNode[*googlesql.ASTSetCollateClause]
}

func newASTSetCollateClause(r *googlesql.ASTSetCollateClause) *SetCollateClause {
	if r == nil {
		return nil
	}
	return &SetCollateClause{baseNode[*googlesql.ASTSetCollateClause]{raw: r}}
}
func (n *SetCollateClause) isAlterAction()    {}
func (n *SetCollateClause) Collate() *Collate { return newASTCollate(must(n.raw.Collate())) }

// SetOptionsAction wraps *googlesql.ASTSetOptionsAction.
type SetOptionsAction struct {
	baseNode[*googlesql.ASTSetOptionsAction]
}

func newASTSetOptionsAction(r *googlesql.ASTSetOptionsAction) *SetOptionsAction {
	if r == nil {
		return nil
	}
	return &SetOptionsAction{baseNode[*googlesql.ASTSetOptionsAction]{raw: r}}
}
func (n *SetOptionsAction) isAlterAction() {}

func (n *SetOptionsAction) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// ─── ALTER STATEMENT wrappers ─────────────────────────────────────────────────
// Methods GetDdlTarget, ActionList, IsIfExists are promoted through the
// embedded *AlterStatementBase, so we can call them directly on n.raw.

// AlterAllRowAccessPoliciesStatement wraps *googlesql.ASTAlterAllRowAccessPoliciesStatement.
type AlterAllRowAccessPoliciesStatement struct {
	baseNode[*googlesql.ASTAlterAllRowAccessPoliciesStatement]
}

func newASTAlterAllRowAccessPoliciesStatement(r *googlesql.ASTAlterAllRowAccessPoliciesStatement) *AlterAllRowAccessPoliciesStatement {
	if r == nil {
		return nil
	}
	return &AlterAllRowAccessPoliciesStatement{baseNode[*googlesql.ASTAlterAllRowAccessPoliciesStatement]{raw: r}}
}
func (n *AlterAllRowAccessPoliciesStatement) isStatement() {}

// AlterDatabaseStatement wraps *googlesql.ASTAlterDatabaseStatement.
type AlterDatabaseStatement struct {
	baseNode[*googlesql.ASTAlterDatabaseStatement]
}

func newASTAlterDatabaseStatement(r *googlesql.ASTAlterDatabaseStatement) *AlterDatabaseStatement {
	if r == nil {
		return nil
	}
	return &AlterDatabaseStatement{baseNode[*googlesql.ASTAlterDatabaseStatement]{raw: r}}
}
func (n *AlterDatabaseStatement) isStatement() {}

func (n *AlterDatabaseStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterDatabaseStatement) ActionList() *AlterActionList {
	return newASTAlterActionList(must(n.raw.ActionList()))
}

// AlterEntityStatement wraps *googlesql.ASTAlterEntityStatement.
type AlterEntityStatement struct {
	baseNode[*googlesql.ASTAlterEntityStatement]
}

func newASTAlterEntityStatement(r *googlesql.ASTAlterEntityStatement) *AlterEntityStatement {
	if r == nil {
		return nil
	}
	return &AlterEntityStatement{baseNode[*googlesql.ASTAlterEntityStatement]{raw: r}}
}
func (n *AlterEntityStatement) isStatement() {}

func (n *AlterEntityStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterEntityStatement) ActionList() *AlterActionList {
	return newASTAlterActionList(must(n.raw.ActionList()))
}

// AlterMaterializedViewStatement wraps *googlesql.ASTAlterMaterializedViewStatement.
type AlterMaterializedViewStatement struct {
	baseNode[*googlesql.ASTAlterMaterializedViewStatement]
}

func newASTAlterMaterializedViewStatement(r *googlesql.ASTAlterMaterializedViewStatement) *AlterMaterializedViewStatement {
	if r == nil {
		return nil
	}
	return &AlterMaterializedViewStatement{baseNode[*googlesql.ASTAlterMaterializedViewStatement]{raw: r}}
}
func (n *AlterMaterializedViewStatement) isStatement() {}

func (n *AlterMaterializedViewStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterMaterializedViewStatement) ActionList() *AlterActionList {
	return newASTAlterActionList(must(n.raw.ActionList()))
}

// AlterPrivilegeRestrictionStatement wraps *googlesql.ASTAlterPrivilegeRestrictionStatement.
type AlterPrivilegeRestrictionStatement struct {
	baseNode[*googlesql.ASTAlterPrivilegeRestrictionStatement]
}

func newASTAlterPrivilegeRestrictionStatement(r *googlesql.ASTAlterPrivilegeRestrictionStatement) *AlterPrivilegeRestrictionStatement {
	if r == nil {
		return nil
	}
	return &AlterPrivilegeRestrictionStatement{baseNode[*googlesql.ASTAlterPrivilegeRestrictionStatement]{raw: r}}
}
func (n *AlterPrivilegeRestrictionStatement) isStatement() {}

// AlterRowAccessPolicyStatement wraps *googlesql.ASTAlterRowAccessPolicyStatement.
type AlterRowAccessPolicyStatement struct {
	baseNode[*googlesql.ASTAlterRowAccessPolicyStatement]
}

func newASTAlterRowAccessPolicyStatement(r *googlesql.ASTAlterRowAccessPolicyStatement) *AlterRowAccessPolicyStatement {
	if r == nil {
		return nil
	}
	return &AlterRowAccessPolicyStatement{baseNode[*googlesql.ASTAlterRowAccessPolicyStatement]{raw: r}}
}
func (n *AlterRowAccessPolicyStatement) isStatement() {}

func (n *AlterRowAccessPolicyStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterRowAccessPolicyStatement) ActionList() *AlterActionList {
	return newASTAlterActionList(must(n.raw.ActionList()))
}

// AlterSchemaStatement wraps *googlesql.ASTAlterSchemaStatement.
type AlterSchemaStatement struct {
	baseNode[*googlesql.ASTAlterSchemaStatement]
}

func newASTAlterSchemaStatement(r *googlesql.ASTAlterSchemaStatement) *AlterSchemaStatement {
	if r == nil {
		return nil
	}
	return &AlterSchemaStatement{baseNode[*googlesql.ASTAlterSchemaStatement]{raw: r}}
}
func (n *AlterSchemaStatement) isStatement() {}

func (n *AlterSchemaStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterSchemaStatement) ActionList() *AlterActionList {
	return newASTAlterActionList(must(n.raw.ActionList()))
}

// AlterTableStatement wraps *googlesql.ASTAlterTableStatement.
type AlterTableStatement struct {
	baseNode[*googlesql.ASTAlterTableStatement]
}

func newASTAlterTableStatement(r *googlesql.ASTAlterTableStatement) *AlterTableStatement {
	if r == nil {
		return nil
	}
	return &AlterTableStatement{baseNode[*googlesql.ASTAlterTableStatement]{raw: r}}
}
func (n *AlterTableStatement) isStatement() {}

func (n *AlterTableStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterTableStatement) ActionList() *AlterActionList {
	return newASTAlterActionList(must(n.raw.ActionList()))
}

// AlterViewStatement wraps *googlesql.ASTAlterViewStatement.
type AlterViewStatement struct {
	baseNode[*googlesql.ASTAlterViewStatement]
}

func newASTAlterViewStatement(r *googlesql.ASTAlterViewStatement) *AlterViewStatement {
	if r == nil {
		return nil
	}
	return &AlterViewStatement{baseNode[*googlesql.ASTAlterViewStatement]{raw: r}}
}
func (n *AlterViewStatement) isStatement() {}

func (n *AlterViewStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterViewStatement) ActionList() *AlterActionList {
	return newASTAlterActionList(must(n.raw.ActionList()))
}

// ─── CREATE helpers ───────────────────────────────────────────────────────────

// CloneDataSource wraps *googlesql.ASTCloneDataSource.
type CloneDataSource struct {
	baseNode[*googlesql.ASTCloneDataSource]
}

func newASTCloneDataSource(r *googlesql.ASTCloneDataSource) *CloneDataSource {
	if r == nil {
		return nil
	}
	return &CloneDataSource{baseNode[*googlesql.ASTCloneDataSource]{raw: r}}
}

func (n *CloneDataSource) PathExpr() *PathExpression {
	return newASTPathExpression(must(n.raw.PathExpr()))
}

func (n *CloneDataSource) ForSystemTime() *ForSystemTime {
	return newASTForSystemTime(must(n.raw.ForSystemTime()))
}

func (n *CloneDataSource) WhereClause() *WhereClause {
	return newASTWhereClause(must(n.raw.WhereClause()))
}

// CopyDataSource wraps *googlesql.ASTCopyDataSource.
type CopyDataSource struct {
	baseNode[*googlesql.ASTCopyDataSource]
}

func newASTCopyDataSource(r *googlesql.ASTCopyDataSource) *CopyDataSource {
	if r == nil {
		return nil
	}
	return &CopyDataSource{baseNode[*googlesql.ASTCopyDataSource]{raw: r}}
}

func (n *CopyDataSource) PathExpr() *PathExpression {
	return newASTPathExpression(must(n.raw.PathExpr()))
}

func (n *CopyDataSource) ForSystemTime() *ForSystemTime {
	return newASTForSystemTime(must(n.raw.ForSystemTime()))
}

func (n *CopyDataSource) WhereClause() *WhereClause {
	return newASTWhereClause(must(n.raw.WhereClause()))
}

// WithConnectionClause wraps *googlesql.ASTWithConnectionClause.
type WithConnectionClause struct {
	baseNode[*googlesql.ASTWithConnectionClause]
}

func newASTWithConnectionClause(r *googlesql.ASTWithConnectionClause) *WithConnectionClause {
	if r == nil {
		return nil
	}
	return &WithConnectionClause{baseNode[*googlesql.ASTWithConnectionClause]{raw: r}}
}

func (n *WithConnectionClause) ConnectionClause() *ConnectionClause {
	return newASTConnectionClause(must(n.raw.ConnectionClause()))
}

// ConnectionClause wraps *googlesql.ASTConnectionClause.
type ConnectionClause struct {
	baseNode[*googlesql.ASTConnectionClause]
}

func newASTConnectionClause(r *googlesql.ASTConnectionClause) *ConnectionClause {
	if r == nil {
		return nil
	}
	return &ConnectionClause{baseNode[*googlesql.ASTConnectionClause]{raw: r}}
}

func (n *ConnectionClause) ConnectionPath() *PathExpression {
	return newASTPathExpression(must(n.raw.ConnectionPath()).(*googlesql.ASTPathExpression))
}

// WithPartitionColumnsClause wraps *googlesql.ASTWithPartitionColumnsClause.
type WithPartitionColumnsClause struct {
	baseNode[*googlesql.ASTWithPartitionColumnsClause]
}

func newASTWithPartitionColumnsClause(r *googlesql.ASTWithPartitionColumnsClause) *WithPartitionColumnsClause {
	if r == nil {
		return nil
	}
	return &WithPartitionColumnsClause{baseNode[*googlesql.ASTWithPartitionColumnsClause]{raw: r}}
}

func (n *WithPartitionColumnsClause) TableElementList() *TableElementList {
	return newASTTableElementList(must(n.raw.TableElementList()))
}

// FunctionDeclaration wraps *googlesql.ASTFunctionDeclaration.
type FunctionDeclaration struct {
	baseNode[*googlesql.ASTFunctionDeclaration]
}

func newASTFunctionDeclaration(r *googlesql.ASTFunctionDeclaration) *FunctionDeclaration {
	if r == nil {
		return nil
	}
	return &FunctionDeclaration{baseNode[*googlesql.ASTFunctionDeclaration]{raw: r}}
}

func (n *FunctionDeclaration) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

func (n *FunctionDeclaration) Parameters() *FunctionParameters {
	return newASTFunctionParameters(must(n.raw.Parameters()))
}

// FunctionParameters wraps *googlesql.ASTFunctionParameters.
type FunctionParameters struct {
	baseNode[*googlesql.ASTFunctionParameters]
}

func newASTFunctionParameters(r *googlesql.ASTFunctionParameters) *FunctionParameters {
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

// FunctionParameter wraps *googlesql.ASTFunctionParameter.
type FunctionParameter struct {
	baseNode[*googlesql.ASTFunctionParameter]
}

func newASTFunctionParameter(r *googlesql.ASTFunctionParameter) *FunctionParameter {
	if r == nil {
		return nil
	}
	return &FunctionParameter{baseNode[*googlesql.ASTFunctionParameter]{raw: r}}
}

func (n *FunctionParameter) Name() *Identifier {
	return newASTIdentifier(must(n.raw.Name()))
}

func (n *FunctionParameter) ProcedureParameterMode() ParameterMode {
	return must(n.raw.ProcedureParameterMode())
}
func (n *FunctionParameter) IsNotAggregate() bool { return must(n.raw.IsNotAggregate()) }
func (n *FunctionParameter) Type() TypeNode       { return wrapType(must(n.raw.Type())) }
func (n *FunctionParameter) TemplatedParameterType() *TemplatedParameterType {
	return newASTTemplatedParameterType(must(n.raw.TemplatedParameterType()))
}
func (n *FunctionParameter) TvfSchema() Node { return Wrap(must(n.raw.TvfSchema())) }
func (n *FunctionParameter) Alias() *Alias   { return newASTAlias(must(n.raw.Alias())) }
func (n *FunctionParameter) DefaultValue() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultValue()))
}

// SQLFunctionBody wraps *googlesql.ASTSqlFunctionBody.
//
//nolint:revive
type SQLFunctionBody struct {
	baseNode[*googlesql.ASTSqlFunctionBody]
}

func newASTSQLFunctionBody(r *googlesql.ASTSqlFunctionBody) *SQLFunctionBody {
	if r == nil {
		return nil
	}
	return &SQLFunctionBody{baseNode[*googlesql.ASTSqlFunctionBody]{raw: r}}
}

func (n *SQLFunctionBody) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// GranteeList wraps *googlesql.ASTGranteeList.
type GranteeList struct {
	baseNode[*googlesql.ASTGranteeList]
}

func newASTGranteeList(r *googlesql.ASTGranteeList) *GranteeList {
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

// GrantToClause wraps *googlesql.ASTGrantToClause.
type GrantToClause struct {
	baseNode[*googlesql.ASTGrantToClause]
}

func newASTGrantToClause(r *googlesql.ASTGrantToClause) *GrantToClause {
	if r == nil {
		return nil
	}
	return &GrantToClause{baseNode[*googlesql.ASTGrantToClause]{raw: r}}
}

func (n *GrantToClause) GranteeList() *GranteeList {
	return newASTGranteeList(must(n.raw.GranteeList()))
}

// FilterUsingClause wraps *googlesql.ASTFilterUsingClause.
type FilterUsingClause struct {
	baseNode[*googlesql.ASTFilterUsingClause]
}

func newASTFilterUsingClause(r *googlesql.ASTFilterUsingClause) *FilterUsingClause {
	if r == nil {
		return nil
	}
	return &FilterUsingClause{baseNode[*googlesql.ASTFilterUsingClause]{raw: r}}
}

func (n *FilterUsingClause) Predicate() ExpressionNode {
	return wrapExpr(must(n.raw.Predicate()))
}

// ColumnWithOptions wraps *googlesql.ASTColumnWithOptions.
type ColumnWithOptions struct {
	baseNode[*googlesql.ASTColumnWithOptions]
}

func newASTColumnWithOptions(r *googlesql.ASTColumnWithOptions) *ColumnWithOptions {
	if r == nil {
		return nil
	}
	return &ColumnWithOptions{baseNode[*googlesql.ASTColumnWithOptions]{raw: r}}
}

func (n *ColumnWithOptions) Name() *Identifier {
	return newASTIdentifier(must(n.raw.Name()))
}

func (n *ColumnWithOptions) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// ColumnWithOptionsList wraps *googlesql.ASTColumnWithOptionsList.
type ColumnWithOptionsList struct {
	baseNode[*googlesql.ASTColumnWithOptionsList]
}

func newASTColumnWithOptionsList(r *googlesql.ASTColumnWithOptionsList) *ColumnWithOptionsList {
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

// ─── CREATE statement concrete wrappers ──────────────────────────────────────
// Methods like IsOrReplace, IsIfNotExists, Scope are promoted from
// embedded *CreateStatement, so n.raw.IsOrReplace() works directly.

// CreateExternalTableStatement wraps *googlesql.ASTCreateExternalTableStatement.
type CreateExternalTableStatement struct {
	baseNode[*googlesql.ASTCreateExternalTableStatement]
}

func newASTCreateExternalTableStatement(r *googlesql.ASTCreateExternalTableStatement) *CreateExternalTableStatement {
	if r == nil {
		return nil
	}
	return &CreateExternalTableStatement{baseNode[*googlesql.ASTCreateExternalTableStatement]{raw: r}}
}
func (n *CreateExternalTableStatement) isStatement()        {}
func (n *CreateExternalTableStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateExternalTableStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateExternalTableStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateExternalTableStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateExternalTableStatement) TableElementList() *TableElementList {
	return newASTTableElementList(must(n.raw.TableElementList()))
}

func (n *CreateExternalTableStatement) WithConnectionClause() *WithConnectionClause {
	return newASTWithConnectionClause(must(n.raw.WithConnectionClause()))
}

func (n *CreateExternalTableStatement) WithPartitionColumnsClause() *WithPartitionColumnsClause {
	return newASTWithPartitionColumnsClause(must(n.raw.WithPartitionColumnsClause()))
}

func (n *CreateExternalTableStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// CreateFunctionStatement wraps *googlesql.ASTCreateFunctionStatement.
type CreateFunctionStatement struct {
	baseNode[*googlesql.ASTCreateFunctionStatement]
}

func newASTCreateFunctionStatement(r *googlesql.ASTCreateFunctionStatement) *CreateFunctionStatement {
	if r == nil {
		return nil
	}
	return &CreateFunctionStatement{baseNode[*googlesql.ASTCreateFunctionStatement]{raw: r}}
}
func (n *CreateFunctionStatement) isStatement()        {}
func (n *CreateFunctionStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateFunctionStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateFunctionStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateFunctionStatement) IsAggregate() bool   { return must(n.raw.IsAggregate()) }
func (n *CreateFunctionStatement) IsRemote() bool      { return must(n.raw.IsRemote()) }
func (n *CreateFunctionStatement) FunctionDeclaration() *FunctionDeclaration {
	return newASTFunctionDeclaration(must(n.raw.FunctionDeclaration()))
}
func (n *CreateFunctionStatement) ReturnType() TypeNode { return wrapType(must(n.raw.ReturnType())) }
func (n *CreateFunctionStatement) DeterminismLevel() DeterminismLevel {
	return must(n.raw.DeterminismLevel())
}

func (n *CreateFunctionStatement) Language() *Identifier {
	return newASTIdentifier(must(n.raw.Language()))
}

func (n *CreateFunctionStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateFunctionStatement) Code() *StringLiteral {
	return newASTStringLiteral(must(n.raw.Code()))
}

func (n *CreateFunctionStatement) SQLFunctionBody() *SQLFunctionBody {
	return newASTSQLFunctionBody(must(n.raw.SqlFunctionBody()))
}

func (n *CreateFunctionStatement) WithConnectionClause() *WithConnectionClause {
	return newASTWithConnectionClause(must(n.raw.WithConnectionClause()))
}

// CreateMaterializedViewStatement wraps *googlesql.ASTCreateMaterializedViewStatement.
type CreateMaterializedViewStatement struct {
	baseNode[*googlesql.ASTCreateMaterializedViewStatement]
}

func newASTCreateMaterializedViewStatement(r *googlesql.ASTCreateMaterializedViewStatement) *CreateMaterializedViewStatement {
	if r == nil {
		return nil
	}
	return &CreateMaterializedViewStatement{baseNode[*googlesql.ASTCreateMaterializedViewStatement]{raw: r}}
}
func (n *CreateMaterializedViewStatement) isStatement()        {}
func (n *CreateMaterializedViewStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateMaterializedViewStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateMaterializedViewStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateMaterializedViewStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}
func (n *CreateMaterializedViewStatement) Recursive() bool { return must(n.raw.Recursive()) }
func (n *CreateMaterializedViewStatement) PartitionBy() *PartitionBy {
	return newASTPartitionBy(must(n.raw.PartitionBy()))
}

func (n *CreateMaterializedViewStatement) ClusterBy() *ClusterBy {
	return newASTClusterBy(must(n.raw.ClusterBy()))
}

func (n *CreateMaterializedViewStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateMaterializedViewStatement) Query() *Query {
	return newASTQuery(must(n.raw.Query()))
}

func (n *CreateMaterializedViewStatement) ReplicaSource() *PathExpression {
	return newASTPathExpression(must(n.raw.ReplicaSource()))
}

// CreateProcedureStatement wraps *googlesql.ASTCreateProcedureStatement.
type CreateProcedureStatement struct {
	baseNode[*googlesql.ASTCreateProcedureStatement]
}

func newASTCreateProcedureStatement(r *googlesql.ASTCreateProcedureStatement) *CreateProcedureStatement {
	if r == nil {
		return nil
	}
	return &CreateProcedureStatement{baseNode[*googlesql.ASTCreateProcedureStatement]{raw: r}}
}
func (n *CreateProcedureStatement) isStatement()        {}
func (n *CreateProcedureStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateProcedureStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateProcedureStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateProcedureStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateProcedureStatement) Parameters() *FunctionParameters {
	return newASTFunctionParameters(must(n.raw.Parameters()))
}

func (n *CreateProcedureStatement) ExternalSecurity() SQLSecurity {
	return must(n.raw.ExternalSecurity())
}

func (n *CreateProcedureStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateProcedureStatement) Body() *Script {
	return newASTScript(must(n.raw.Body()))
}

// CreateRowAccessPolicyStatement wraps *googlesql.ASTCreateRowAccessPolicyStatement.
type CreateRowAccessPolicyStatement struct {
	baseNode[*googlesql.ASTCreateRowAccessPolicyStatement]
}

func newASTCreateRowAccessPolicyStatement(r *googlesql.ASTCreateRowAccessPolicyStatement) *CreateRowAccessPolicyStatement {
	if r == nil {
		return nil
	}
	return &CreateRowAccessPolicyStatement{baseNode[*googlesql.ASTCreateRowAccessPolicyStatement]{raw: r}}
}
func (n *CreateRowAccessPolicyStatement) isStatement()        {}
func (n *CreateRowAccessPolicyStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateRowAccessPolicyStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateRowAccessPolicyStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateRowAccessPolicyStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateRowAccessPolicyStatement) Name() *Identifier {
	return newASTIdentifier(must(n.raw.Name()))
}

func (n *CreateRowAccessPolicyStatement) GrantTo() *GrantToClause {
	return newASTGrantToClause(must(n.raw.GrantTo()))
}

func (n *CreateRowAccessPolicyStatement) FilterUsing() *FilterUsingClause {
	return newASTFilterUsingClause(must(n.raw.FilterUsing()))
}

// CreateSchemaStatement wraps *googlesql.ASTCreateSchemaStatement.
type CreateSchemaStatement struct {
	baseNode[*googlesql.ASTCreateSchemaStatement]
}

func newASTCreateSchemaStatement(r *googlesql.ASTCreateSchemaStatement) *CreateSchemaStatement {
	if r == nil {
		return nil
	}
	return &CreateSchemaStatement{baseNode[*googlesql.ASTCreateSchemaStatement]{raw: r}}
}
func (n *CreateSchemaStatement) isStatement()        {}
func (n *CreateSchemaStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateSchemaStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateSchemaStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateSchemaStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

func (n *CreateSchemaStatement) Collate() *Collate {
	return newASTCollate(must(n.raw.Collate()))
}

func (n *CreateSchemaStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// CreateSnapshotTableStatement wraps *googlesql.ASTCreateSnapshotTableStatement.
type CreateSnapshotTableStatement struct {
	baseNode[*googlesql.ASTCreateSnapshotTableStatement]
}

func newASTCreateSnapshotTableStatement(r *googlesql.ASTCreateSnapshotTableStatement) *CreateSnapshotTableStatement {
	if r == nil {
		return nil
	}
	return &CreateSnapshotTableStatement{baseNode[*googlesql.ASTCreateSnapshotTableStatement]{raw: r}}
}
func (n *CreateSnapshotTableStatement) isStatement()        {}
func (n *CreateSnapshotTableStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateSnapshotTableStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateSnapshotTableStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateSnapshotTableStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateSnapshotTableStatement) CloneDataSource() *CloneDataSource {
	return newASTCloneDataSource(must(n.raw.CloneDataSource()))
}

func (n *CreateSnapshotTableStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

// CreateTableStatement wraps *googlesql.ASTCreateTableStatement.
type CreateTableStatement struct {
	baseNode[*googlesql.ASTCreateTableStatement]
}

func newASTCreateTableStatement(r *googlesql.ASTCreateTableStatement) *CreateTableStatement {
	if r == nil {
		return nil
	}
	return &CreateTableStatement{baseNode[*googlesql.ASTCreateTableStatement]{raw: r}}
}
func (n *CreateTableStatement) isStatement()        {}
func (n *CreateTableStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateTableStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateTableStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateTableStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

func (n *CreateTableStatement) TableElementList() *TableElementList {
	return newASTTableElementList(must(n.raw.TableElementList()))
}

func (n *CreateTableStatement) CopyDataSource() *CopyDataSource {
	return newASTCopyDataSource(must(n.raw.CopyDataSource()))
}

func (n *CreateTableStatement) CloneDataSource() *CloneDataSource {
	return newASTCloneDataSource(must(n.raw.CloneDataSource()))
}

func (n *CreateTableStatement) LikeTableName() *PathExpression {
	return newASTPathExpression(must(n.raw.LikeTableName()))
}

func (n *CreateTableStatement) Collate() *Collate {
	return newASTCollate(must(n.raw.Collate()))
}

func (n *CreateTableStatement) PartitionBy() *PartitionBy {
	return newASTPartitionBy(must(n.raw.PartitionBy()))
}

func (n *CreateTableStatement) ClusterBy() *ClusterBy {
	return newASTClusterBy(must(n.raw.ClusterBy()))
}

func (n *CreateTableStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}
func (n *CreateTableStatement) Query() *Query { return newASTQuery(must(n.raw.Query())) }

// CreateTableFunctionStatement wraps *googlesql.ASTCreateTableFunctionStatement.
type CreateTableFunctionStatement struct {
	baseNode[*googlesql.ASTCreateTableFunctionStatement]
}

func newASTCreateTableFunctionStatement(r *googlesql.ASTCreateTableFunctionStatement) *CreateTableFunctionStatement {
	if r == nil {
		return nil
	}
	return &CreateTableFunctionStatement{baseNode[*googlesql.ASTCreateTableFunctionStatement]{raw: r}}
}
func (n *CreateTableFunctionStatement) isStatement()        {}
func (n *CreateTableFunctionStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateTableFunctionStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateTableFunctionStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateTableFunctionStatement) FunctionDeclaration() *FunctionDeclaration {
	return newASTFunctionDeclaration(must(n.raw.FunctionDeclaration()))
}

func (n *CreateTableFunctionStatement) ReturnTvfSchema() Node {
	return Wrap(must(n.raw.ReturnTvfSchema()))
}

func (n *CreateTableFunctionStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateTableFunctionStatement) Query() *Query {
	return newASTQuery(must(n.raw.Query()))
}

// CreateViewStatement wraps *googlesql.ASTCreateViewStatement.
type CreateViewStatement struct {
	baseNode[*googlesql.ASTCreateViewStatement]
}

func newASTCreateViewStatement(r *googlesql.ASTCreateViewStatement) *CreateViewStatement {
	if r == nil {
		return nil
	}
	return &CreateViewStatement{baseNode[*googlesql.ASTCreateViewStatement]{raw: r}}
}
func (n *CreateViewStatement) isStatement()        {}
func (n *CreateViewStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateViewStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateViewStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateViewStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

func (n *CreateViewStatement) ColumnWithOptionsList() *ColumnWithOptionsList {
	return newASTColumnWithOptionsList(must(n.raw.ColumnWithOptionsList()))
}
func (n *CreateViewStatement) SQLSecurity() SQLSecurity { return must(n.raw.SqlSecurity()) }
func (n *CreateViewStatement) Recursive() bool          { return must(n.raw.Recursive()) }
func (n *CreateViewStatement) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}
func (n *CreateViewStatement) Query() *Query { return newASTQuery(must(n.raw.Query())) }

// ─── DROP statement wrappers ──────────────────────────────────────────────────

// DropAllRowAccessPoliciesStatement wraps *googlesql.ASTDropAllRowAccessPoliciesStatement.
type DropAllRowAccessPoliciesStatement struct {
	baseNode[*googlesql.ASTDropAllRowAccessPoliciesStatement]
}

func newASTDropAllRowAccessPoliciesStatement(r *googlesql.ASTDropAllRowAccessPoliciesStatement) *DropAllRowAccessPoliciesStatement {
	if r == nil {
		return nil
	}
	return &DropAllRowAccessPoliciesStatement{baseNode[*googlesql.ASTDropAllRowAccessPoliciesStatement]{raw: r}}
}
func (n *DropAllRowAccessPoliciesStatement) isStatement() {}

func (n *DropAllRowAccessPoliciesStatement) TableName() *PathExpression {
	return newASTPathExpression(must(n.raw.TableName()))
}

// DropEntityStatement wraps *googlesql.ASTDropEntityStatement.
type DropEntityStatement struct {
	baseNode[*googlesql.ASTDropEntityStatement]
}

func newASTDropEntityStatement(r *googlesql.ASTDropEntityStatement) *DropEntityStatement {
	if r == nil {
		return nil
	}
	return &DropEntityStatement{baseNode[*googlesql.ASTDropEntityStatement]{raw: r}}
}
func (n *DropEntityStatement) isStatement()     {}
func (n *DropEntityStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropEntityStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}

// DropFunctionStatement wraps *googlesql.ASTDropFunctionStatement.
type DropFunctionStatement struct {
	baseNode[*googlesql.ASTDropFunctionStatement]
}

func newASTDropFunctionStatement(r *googlesql.ASTDropFunctionStatement) *DropFunctionStatement {
	if r == nil {
		return nil
	}
	return &DropFunctionStatement{baseNode[*googlesql.ASTDropFunctionStatement]{raw: r}}
}
func (n *DropFunctionStatement) isStatement()     {}
func (n *DropFunctionStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropFunctionStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

// DropMaterializedViewStatement wraps *googlesql.ASTDropMaterializedViewStatement.
type DropMaterializedViewStatement struct {
	baseNode[*googlesql.ASTDropMaterializedViewStatement]
}

func newASTDropMaterializedViewStatement(r *googlesql.ASTDropMaterializedViewStatement) *DropMaterializedViewStatement {
	if r == nil {
		return nil
	}
	return &DropMaterializedViewStatement{baseNode[*googlesql.ASTDropMaterializedViewStatement]{raw: r}}
}
func (n *DropMaterializedViewStatement) isStatement()     {}
func (n *DropMaterializedViewStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropMaterializedViewStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

// DropPrivilegeRestrictionStatement wraps *googlesql.ASTDropPrivilegeRestrictionStatement.
type DropPrivilegeRestrictionStatement struct {
	baseNode[*googlesql.ASTDropPrivilegeRestrictionStatement]
}

func newASTDropPrivilegeRestrictionStatement(r *googlesql.ASTDropPrivilegeRestrictionStatement) *DropPrivilegeRestrictionStatement {
	if r == nil {
		return nil
	}
	return &DropPrivilegeRestrictionStatement{baseNode[*googlesql.ASTDropPrivilegeRestrictionStatement]{raw: r}}
}
func (n *DropPrivilegeRestrictionStatement) isStatement() {}

// DropRowAccessPolicyStatement wraps *googlesql.ASTDropRowAccessPolicyStatement.
type DropRowAccessPolicyStatement struct {
	baseNode[*googlesql.ASTDropRowAccessPolicyStatement]
}

func newASTDropRowAccessPolicyStatement(r *googlesql.ASTDropRowAccessPolicyStatement) *DropRowAccessPolicyStatement {
	if r == nil {
		return nil
	}
	return &DropRowAccessPolicyStatement{baseNode[*googlesql.ASTDropRowAccessPolicyStatement]{raw: r}}
}
func (n *DropRowAccessPolicyStatement) isStatement()     {}
func (n *DropRowAccessPolicyStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropRowAccessPolicyStatement) Name() *Identifier {
	return newASTIdentifier(must(n.raw.Name()))
}

func (n *DropRowAccessPolicyStatement) TableName() *PathExpression {
	return newASTPathExpression(must(n.raw.TableName()))
}

// DropSearchIndexStatement wraps *googlesql.ASTDropSearchIndexStatement.
type DropSearchIndexStatement struct {
	baseNode[*googlesql.ASTDropSearchIndexStatement]
}

func newASTDropSearchIndexStatement(r *googlesql.ASTDropSearchIndexStatement) *DropSearchIndexStatement {
	if r == nil {
		return nil
	}
	return &DropSearchIndexStatement{baseNode[*googlesql.ASTDropSearchIndexStatement]{raw: r}}
}
func (n *DropSearchIndexStatement) isStatement()     {}
func (n *DropSearchIndexStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropSearchIndexStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

// DropSnapshotTableStatement wraps *googlesql.ASTDropSnapshotTableStatement.
type DropSnapshotTableStatement struct {
	baseNode[*googlesql.ASTDropSnapshotTableStatement]
}

func newASTDropSnapshotTableStatement(r *googlesql.ASTDropSnapshotTableStatement) *DropSnapshotTableStatement {
	if r == nil {
		return nil
	}
	return &DropSnapshotTableStatement{baseNode[*googlesql.ASTDropSnapshotTableStatement]{raw: r}}
}
func (n *DropSnapshotTableStatement) isStatement()     {}
func (n *DropSnapshotTableStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropSnapshotTableStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

// DropTableFunctionStatement wraps *googlesql.ASTDropTableFunctionStatement.
type DropTableFunctionStatement struct {
	baseNode[*googlesql.ASTDropTableFunctionStatement]
}

func newASTDropTableFunctionStatement(r *googlesql.ASTDropTableFunctionStatement) *DropTableFunctionStatement {
	if r == nil {
		return nil
	}
	return &DropTableFunctionStatement{baseNode[*googlesql.ASTDropTableFunctionStatement]{raw: r}}
}
func (n *DropTableFunctionStatement) isStatement()     {}
func (n *DropTableFunctionStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropTableFunctionStatement) Name() *PathExpression {
	return newASTPathExpression(must(n.raw.Name()))
}

// DropStatement wraps *googlesql.ASTDropStatement.
type DropStatement struct {
	baseNode[*googlesql.ASTDropStatement]
}

func newASTDropStatement(r *googlesql.ASTDropStatement) *DropStatement {
	if r == nil {
		return nil
	}
	return &DropStatement{baseNode[*googlesql.ASTDropStatement]{raw: r}}
}
func (n *DropStatement) isStatement()     {}
func (n *DropStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropStatement) GetDdlTarget() *PathExpression {
	return newASTPathExpression(must(n.raw.GetDdlTarget()))
}
func (n *DropStatement) DropMode() DropMode { return must(n.raw.DropMode()) }
func (n *DropStatement) SchemaObjectKind() SchemaObjectKind {
	return must(n.raw.SchemaObjectKind())
}

// ColumnSchema wraps *googlesql.ASTColumnSchema.
type ColumnSchema struct {
	baseNode[*googlesql.ASTColumnSchema]
}

func newASTColumnSchema(r *googlesql.ASTColumnSchema) *ColumnSchema {
	if r == nil {
		return nil
	}
	return &ColumnSchema{baseNode[*googlesql.ASTColumnSchema]{raw: r}}
}

func (n *ColumnSchema) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

func (n *ColumnSchema) Attributes() *ColumnAttributeList {
	return newASTColumnAttributeList(must(n.raw.Attributes()))
}

func (n *ColumnSchema) Collate() *Collate {
	return newASTCollate(must(n.raw.Collate()))
}

func (n *ColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *ColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newASTGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *ColumnSchema) OptionsList() *OptionsList {
	return newASTOptionsList(must(n.raw.OptionsList()))
}
