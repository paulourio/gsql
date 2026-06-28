package sql

import (
	"github.com/goccy/go-googlesql"
)

// ColumnAttributeList wraps *googlesql.ASTColumnAttributeList.
type ColumnAttributeList struct {
	baseNode[*googlesql.ASTColumnAttributeList]
}

func newColumnAttributeList(r *googlesql.ASTColumnAttributeList) *ColumnAttributeList {
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

// ForeignKeyColumnAttribute wraps *googlesql.ASTForeignKeyColumnAttribute.
type ForeignKeyColumnAttribute struct {
	baseNode[*googlesql.ASTForeignKeyColumnAttribute]
}

func newForeignKeyColumnAttribute(r *googlesql.ASTForeignKeyColumnAttribute) *ForeignKeyColumnAttribute {
	if r == nil {
		return nil
	}
	return &ForeignKeyColumnAttribute{baseNode[*googlesql.ASTForeignKeyColumnAttribute]{raw: r}}
}

func (n *ForeignKeyColumnAttribute) isColumnAttribute() {}

func (n *ForeignKeyColumnAttribute) ConstraintName() *Identifier {
	return newIdentifier(must(n.raw.ConstraintName()))
}

func (n *ForeignKeyColumnAttribute) Reference() *ForeignKeyReference {
	return newForeignKeyReference(must(n.raw.Reference()))
}

// HiddenColumnAttribute wraps *googlesql.ASTHiddenColumnAttribute.
type HiddenColumnAttribute struct {
	baseNode[*googlesql.ASTHiddenColumnAttribute]
}

func newHiddenColumnAttribute(r *googlesql.ASTHiddenColumnAttribute) *HiddenColumnAttribute {
	if r == nil {
		return nil
	}
	return &HiddenColumnAttribute{baseNode[*googlesql.ASTHiddenColumnAttribute]{raw: r}}
}

func (n *HiddenColumnAttribute) isColumnAttribute() {}

// NotNullColumnAttribute wraps *googlesql.ASTNotNullColumnAttribute.
type NotNullColumnAttribute struct {
	baseNode[*googlesql.ASTNotNullColumnAttribute]
}

func newNotNullColumnAttribute(r *googlesql.ASTNotNullColumnAttribute) *NotNullColumnAttribute {
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

func newPrimaryKeyColumnAttribute(r *googlesql.ASTPrimaryKeyColumnAttribute) *PrimaryKeyColumnAttribute {
	if r == nil {
		return nil
	}
	return &PrimaryKeyColumnAttribute{baseNode[*googlesql.ASTPrimaryKeyColumnAttribute]{raw: r}}
}

func (n *PrimaryKeyColumnAttribute) isColumnAttribute() {}

func (n *PrimaryKeyColumnAttribute) Enforced() bool { return must(n.raw.Enforced()) }
