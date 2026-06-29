package sql

import (
	"github.com/goccy/go-googlesql"
)

// ArrayColumnSchema wraps *googlesql.ASTArrayColumnSchema.
type ArrayColumnSchema struct {
	baseNode[*googlesql.ASTArrayColumnSchema]
}

func newArrayColumnSchema(r *googlesql.ASTArrayColumnSchema) *ArrayColumnSchema {
	if r == nil {
		return nil
	}
	return &ArrayColumnSchema{baseNode[*googlesql.ASTArrayColumnSchema]{raw: r}}
}

func (n *ArrayColumnSchema) ElementSchema() Node {
	return Wrap(must(n.raw.ElementSchema()))
}

func (n *ArrayColumnSchema) TypeParameters() *TypeParameterList {
	return newTypeParameterList(must(n.raw.TypeParameters()))
}

func (n *ArrayColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *ArrayColumnSchema) Attributes() *ColumnAttributeList {
	return newColumnAttributeList(must(n.raw.Attributes()))
}

func (n *ArrayColumnSchema) Collate() *Collate {
	return newCollate(must(n.raw.Collate()))
}

func (n *ArrayColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *ArrayColumnSchema) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// ColumnDefinition wraps *googlesql.ASTColumnDefinition.
type ColumnDefinition struct {
	baseNode[*googlesql.ASTColumnDefinition]
}

func newColumnDefinition(r *googlesql.ASTColumnDefinition) *ColumnDefinition {
	if r == nil {
		return nil
	}
	return &ColumnDefinition{baseNode[*googlesql.ASTColumnDefinition]{raw: r}}
}

func (n *ColumnDefinition) isTableElement() {}

func (n *ColumnDefinition) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

// Schema returns the column schema as Node.
func (n *ColumnDefinition) Schema() Node { return Wrap(must(n.raw.Schema())) }

// ColumnSchema wraps *googlesql.ASTColumnSchema.
type ColumnSchema struct {
	baseNode[*googlesql.ASTColumnSchema]
}

func newColumnSchema(r *googlesql.ASTColumnSchema) *ColumnSchema {
	if r == nil {
		return nil
	}
	return &ColumnSchema{baseNode[*googlesql.ASTColumnSchema]{raw: r}}
}

func (n *ColumnSchema) TypeParameters() *TypeParameterList {
	return newTypeParameterList(must(n.raw.TypeParameters()))
}

func (n *ColumnSchema) Attributes() *ColumnAttributeList {
	return newColumnAttributeList(must(n.raw.Attributes()))
}

func (n *ColumnSchema) Collate() *Collate {
	return newCollate(must(n.raw.Collate()))
}

func (n *ColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *ColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *ColumnSchema) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// ForeignKey wraps *googlesql.ASTForeignKey.
type ForeignKey struct {
	baseNode[*googlesql.ASTForeignKey]
}

func newForeignKey(r *googlesql.ASTForeignKey) *ForeignKey {
	if r == nil {
		return nil
	}
	return &ForeignKey{baseNode[*googlesql.ASTForeignKey]{raw: r}}
}

func (n *ForeignKey) isTableElement() {}

func (n *ForeignKey) ConstraintName() *Identifier {
	return newIdentifier(must(n.raw.ConstraintName()))
}

func (n *ForeignKey) ColumnList() *ColumnList {
	return newColumnList(must(n.raw.ColumnList()))
}

func (n *ForeignKey) Reference() *ForeignKeyReference {
	return newForeignKeyReference(must(n.raw.Reference()))
}

func (n *ForeignKey) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// ForeignKeyReference wraps *googlesql.ASTForeignKeyReference.
type ForeignKeyReference struct {
	baseNode[*googlesql.ASTForeignKeyReference]
}

func newForeignKeyReference(r *googlesql.ASTForeignKeyReference) *ForeignKeyReference {
	if r == nil {
		return nil
	}
	return &ForeignKeyReference{baseNode[*googlesql.ASTForeignKeyReference]{raw: r}}
}

func (n *ForeignKeyReference) TableName() *PathExpression {
	return newPathExpression(must(n.raw.TableName()))
}

func (n *ForeignKeyReference) ColumnList() *ColumnList {
	return newColumnList(must(n.raw.ColumnList()))
}

func (n *ForeignKeyReference) Enforced() bool { return must(n.raw.Enforced()) }

func (n *ForeignKeyReference) Match() ForeignKeyMatch { return must(n.raw.Match()) }

// PrimaryKey wraps *googlesql.ASTPrimaryKey.
type PrimaryKey struct {
	baseNode[*googlesql.ASTPrimaryKey]
}

func newPrimaryKey(r *googlesql.ASTPrimaryKey) *PrimaryKey {
	if r == nil {
		return nil
	}
	return &PrimaryKey{baseNode[*googlesql.ASTPrimaryKey]{raw: r}}
}

func (n *PrimaryKey) isTableElement() {}

func (n *PrimaryKey) ConstraintName() *Identifier {
	return newIdentifier(must(n.raw.ConstraintName()))
}

func (n *PrimaryKey) ElementList() *PrimaryKeyElementList {
	return newPrimaryKeyElementList(must(n.raw.ElementList()))
}

func (n *PrimaryKey) Enforced() bool { return must(n.raw.Enforced()) }

func (n *PrimaryKey) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// PrimaryKeyElement wraps *googlesql.ASTPrimaryKeyElement.
type PrimaryKeyElement struct {
	baseNode[*googlesql.ASTPrimaryKeyElement]
}

func newPrimaryKeyElement(r *googlesql.ASTPrimaryKeyElement) *PrimaryKeyElement {
	if r == nil {
		return nil
	}
	return &PrimaryKeyElement{baseNode[*googlesql.ASTPrimaryKeyElement]{raw: r}}
}

func (n *PrimaryKeyElement) Column() *Identifier {
	return newIdentifier(must(n.raw.Column()))
}

func (n *PrimaryKeyElement) OrderingSpec() OrderingSpec { return must(n.raw.OrderingSpec()) }

func (n *PrimaryKeyElement) NullOrder() *NullOrder {
	return newNullOrder(must(n.raw.NullOrder()))
}

// PrimaryKeyElementList wraps *googlesql.ASTPrimaryKeyElementList.
type PrimaryKeyElementList struct {
	baseNode[*googlesql.ASTPrimaryKeyElementList]
}

func newPrimaryKeyElementList(r *googlesql.ASTPrimaryKeyElementList) *PrimaryKeyElementList {
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

// SimpleColumnSchema wraps *googlesql.ASTSimpleColumnSchema.
type SimpleColumnSchema struct {
	baseNode[*googlesql.ASTSimpleColumnSchema]
}

func newSimpleColumnSchema(r *googlesql.ASTSimpleColumnSchema) *SimpleColumnSchema {
	if r == nil {
		return nil
	}
	return &SimpleColumnSchema{baseNode[*googlesql.ASTSimpleColumnSchema]{raw: r}}
}

func (n *SimpleColumnSchema) TypeName() *PathExpression {
	return newPathExpression(must(n.raw.TypeName()))
}

func (n *SimpleColumnSchema) Attributes() *ColumnAttributeList {
	return newColumnAttributeList(must(n.raw.Attributes()))
}

func (n *SimpleColumnSchema) Collate() *Collate {
	return newCollate(must(n.raw.Collate()))
}

func (n *SimpleColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *SimpleColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *SimpleColumnSchema) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *SimpleColumnSchema) TypeParameters() *TypeParameterList {
	return newTypeParameterList(must(n.raw.TypeParameters()))
}

// StructColumnField wraps *googlesql.ASTStructColumnField.
type StructColumnField struct {
	baseNode[*googlesql.ASTStructColumnField]
}

func newStructColumnField(r *googlesql.ASTStructColumnField) *StructColumnField {
	if r == nil {
		return nil
	}
	return &StructColumnField{baseNode[*googlesql.ASTStructColumnField]{raw: r}}
}

func (n *StructColumnField) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

func (n *StructColumnField) Schema() Node {
	return Wrap(must(n.raw.Schema()))
}

// StructColumnSchema wraps *googlesql.ASTStructColumnSchema.
type StructColumnSchema struct {
	baseNode[*googlesql.ASTStructColumnSchema]
}

func newStructColumnSchema(r *googlesql.ASTStructColumnSchema) *StructColumnSchema {
	if r == nil {
		return nil
	}
	return &StructColumnSchema{baseNode[*googlesql.ASTStructColumnSchema]{raw: r}}
}

func (n *StructColumnSchema) TypeParameters() *TypeParameterList {
	return newTypeParameterList(must(n.raw.TypeParameters()))
}

func (n *StructColumnSchema) GeneratedColumnInfo() *GeneratedColumnInfo {
	return newGeneratedColumnInfo(must(n.raw.GeneratedColumnInfo()))
}

func (n *StructColumnSchema) Attributes() *ColumnAttributeList {
	return newColumnAttributeList(must(n.raw.Attributes()))
}

func (n *StructColumnSchema) Collate() *Collate {
	return newCollate(must(n.raw.Collate()))
}

func (n *StructColumnSchema) DefaultExpression() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultExpression()))
}

func (n *StructColumnSchema) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
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

// TableConstraint wraps *googlesql.ASTTableConstraint.
type TableConstraint struct {
	baseNode[*googlesql.ASTTableConstraint]
}

func newTableConstraint(r *googlesql.ASTTableConstraint) *TableConstraint {
	if r == nil {
		return nil
	}
	return &TableConstraint{baseNode[*googlesql.ASTTableConstraint]{raw: r}}
}

func (n *TableConstraint) isTableElement() {}

func (n *TableConstraint) ConstraintName() *Identifier {
	return newIdentifier(must(n.raw.ConstraintName()))
}

// TableElementList wraps *googlesql.ASTTableElementList.
type TableElementList struct {
	baseNode[*googlesql.ASTTableElementList]
}

func newTableElementList(r *googlesql.ASTTableElementList) *TableElementList {
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

// CheckConstraint wraps *googlesql.ASTCheckConstraint.
type CheckConstraint struct {
	baseNode[*googlesql.ASTCheckConstraint]
}

func newCheckConstraint(r *googlesql.ASTCheckConstraint) *CheckConstraint {
	if r == nil {
		return nil
	}
	return &CheckConstraint{baseNode[*googlesql.ASTCheckConstraint]{raw: r}}
}
func (n *CheckConstraint) isTableElement() {}
func (n *CheckConstraint) ConstraintName() *Identifier {
	return newIdentifier(must(n.raw.ConstraintName()))
}
func (n *CheckConstraint) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }
func (n *CheckConstraint) IsEnforced() bool           { return must(n.raw.IsEnforced()) }
func (n *CheckConstraint) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// ForeignKeyActions wraps *googlesql.ASTForeignKeyActions.
type ForeignKeyActions struct {
	baseNode[*googlesql.ASTForeignKeyActions]
}

func newForeignKeyActions(r *googlesql.ASTForeignKeyActions) *ForeignKeyActions {
	if r == nil {
		return nil
	}
	return &ForeignKeyActions{baseNode[*googlesql.ASTForeignKeyActions]{raw: r}}
}
func (n *ForeignKeyActions) DeleteAction() ForeignKeyAction { return must(n.raw.DeleteAction()) }
func (n *ForeignKeyActions) UpdateAction() ForeignKeyAction { return must(n.raw.UpdateAction()) }

// MapColumnSchema wraps *googlesql.ASTMapColumnSchema.
type MapColumnSchema struct {
	baseNode[*googlesql.ASTMapColumnSchema]
}

func newMapColumnSchema(r *googlesql.ASTMapColumnSchema) *MapColumnSchema {
	if r == nil {
		return nil
	}
	return &MapColumnSchema{baseNode[*googlesql.ASTMapColumnSchema]{raw: r}}
}
func (n *MapColumnSchema) isTableElement() {}

// RangeColumnSchema wraps *googlesql.ASTRangeColumnSchema.
type RangeColumnSchema struct {
	baseNode[*googlesql.ASTRangeColumnSchema]
}

func newRangeColumnSchema(r *googlesql.ASTRangeColumnSchema) *RangeColumnSchema {
	if r == nil {
		return nil
	}
	return &RangeColumnSchema{baseNode[*googlesql.ASTRangeColumnSchema]{raw: r}}
}
func (n *RangeColumnSchema) isTableElement() {}

// InferredTypeColumnSchema wraps *googlesql.ASTInferredTypeColumnSchema.
type InferredTypeColumnSchema struct {
	baseNode[*googlesql.ASTInferredTypeColumnSchema]
}

func newInferredTypeColumnSchema(r *googlesql.ASTInferredTypeColumnSchema) *InferredTypeColumnSchema {
	if r == nil {
		return nil
	}
	return &InferredTypeColumnSchema{baseNode[*googlesql.ASTInferredTypeColumnSchema]{raw: r}}
}
func (n *InferredTypeColumnSchema) isTableElement() {}
