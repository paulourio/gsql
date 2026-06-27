package sql

import "github.com/goccy/go-googlesql"

// SimpleType wraps *googlesql.ASTSimpleType.
type SimpleType struct {
	baseNode[*googlesql.ASTSimpleType]
}

func newASTSimpleType(r *googlesql.ASTSimpleType) *SimpleType {
	if r == nil {
		return nil
	}
	return &SimpleType{baseNode[*googlesql.ASTSimpleType]{raw: r}}
}
func (n *SimpleType) isType() {}
func (n *SimpleType) TypeName() *PathExpression {
	return newASTPathExpression(must(n.raw.TypeName()))
}
func (n *SimpleType) Collate() *Collate { return newASTCollate(must(n.raw.Collate())) }
func (n *SimpleType) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

// ArrayType wraps *googlesql.ASTArrayType.
type ArrayType struct {
	baseNode[*googlesql.ASTArrayType]
}

func newASTArrayType(r *googlesql.ASTArrayType) *ArrayType {
	if r == nil {
		return nil
	}
	return &ArrayType{baseNode[*googlesql.ASTArrayType]{raw: r}}
}
func (n *ArrayType) isType() {}

// ElementType: raw returns TypeNode (interface) → TypeNode.
func (n *ArrayType) ElementType() TypeNode { return wrapType(must(n.raw.ElementType())) }
func (n *ArrayType) Collate() *Collate     { return newASTCollate(must(n.raw.Collate())) }
func (n *ArrayType) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

// StructType wraps *googlesql.ASTStructType.
type StructType struct {
	baseNode[*googlesql.ASTStructType]
}

func newASTStructType(r *googlesql.ASTStructType) *StructType {
	if r == nil {
		return nil
	}
	return &StructType{baseNode[*googlesql.ASTStructType]{raw: r}}
}
func (n *StructType) isType()           {}
func (n *StructType) Collate() *Collate { return newASTCollate(must(n.raw.Collate())) }
func (n *StructType) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

// StructFields: raw StructFields(i) returns *StructField (concrete) → []*StructField.
func (n *StructType) StructFields() []*StructField {
	var fields []*StructField
	for f := range childrenOfType[*googlesql.ASTStructField](n) {
		fields = append(fields, newASTStructField(f))
	}
	return fields
}

// StructField wraps *googlesql.ASTStructField.
type StructField struct {
	baseNode[*googlesql.ASTStructField]
}

func newASTStructField(r *googlesql.ASTStructField) *StructField {
	if r == nil {
		return nil
	}
	return &StructField{baseNode[*googlesql.ASTStructField]{raw: r}}
}
func (n *StructField) Name() *Identifier { return newASTIdentifier(must(n.raw.Name())) }

// Type: raw returns TypeNode (interface) → TypeNode.
func (n *StructField) Type() TypeNode {
	return wrapType(must(n.raw.Type()))
}

// RangeType wraps *googlesql.ASTRangeType.
type RangeType struct {
	baseNode[*googlesql.ASTRangeType]
}

func newASTRangeType(r *googlesql.ASTRangeType) *RangeType {
	if r == nil {
		return nil
	}
	return &RangeType{baseNode[*googlesql.ASTRangeType]{raw: r}}
}
func (n *RangeType) isType() {}

// ElementType: raw returns TypeNode (interface) → TypeNode.
func (n *RangeType) ElementType() TypeNode { return wrapType(must(n.raw.ElementType())) }
func (n *RangeType) Collate() *Collate     { return newASTCollate(must(n.raw.Collate())) }
func (n *RangeType) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

// MapType wraps *googlesql.ASTMapType.
type MapType struct {
	baseNode[*googlesql.ASTMapType]
}

func newASTMapType(r *googlesql.ASTMapType) *MapType {
	if r == nil {
		return nil
	}
	return &MapType{baseNode[*googlesql.ASTMapType]{raw: r}}
}
func (n *MapType) isType() {}

// KeyType/ValueType: raw returns TypeNode (interface) → TypeNode.
func (n *MapType) KeyType() TypeNode   { return wrapType(must(n.raw.KeyType())) }
func (n *MapType) ValueType() TypeNode { return wrapType(must(n.raw.ValueType())) }
func (n *MapType) Collate() *Collate   { return newASTCollate(must(n.raw.Collate())) }
func (n *MapType) TypeParameters() *TypeParameterList {
	return newASTTypeParameterList(must(n.raw.TypeParameters()))
}

// TemplatedParameterType wraps *googlesql.ASTTemplatedParameterType.
type TemplatedParameterType struct {
	baseNode[*googlesql.ASTTemplatedParameterType]
}

func newASTTemplatedParameterType(r *googlesql.ASTTemplatedParameterType) *TemplatedParameterType {
	if r == nil {
		return nil
	}
	return &TemplatedParameterType{baseNode[*googlesql.ASTTemplatedParameterType]{raw: r}}
}
func (n *TemplatedParameterType) isType()                          {}
func (n *TemplatedParameterType) TemplatedKind() TemplatedTypeKind { return must(n.raw.Kind()) }

// TypeParameterList wraps *googlesql.ASTTypeParameterList.
type TypeParameterList struct {
	baseNode[*googlesql.ASTTypeParameterList]
}

func newASTTypeParameterList(r *googlesql.ASTTypeParameterList) *TypeParameterList {
	if r == nil {
		return nil
	}
	return &TypeParameterList{baseNode[*googlesql.ASTTypeParameterList]{raw: r}}
}

// Parameters returns all type parameters.
func (n *TypeParameterList) Parameters() []LeafNode {
	var params []LeafNode
	for p := range childrenOfType[googlesql.ASTLeafNode](n) {
		params = append(params, wrapLeaf(p))
	}
	return params
}
