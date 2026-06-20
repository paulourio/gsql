// Functions to print types and schemas.
package printer

import (
	"fmt"
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

func (p *Printer) VisitArrayColumnSchema(ctx Context, n *googlesql.ASTArrayColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx.WithValue(KeyInTypeName, true), ast.Must(n.ElementSchema()))
	p1.accept(ctx, ast.Must(n.TypeParameters()))
	typespec := strings.TrimLeft(p1.unnestLeft(), "\v")
	if strings.Contains(typespec, "\n") {
		pp.println(pp.keyword("ARRAY") + "<")
		pp.incDepth()
		pp.println(p1.unnestLeft())
		pp.decDepth()
		pp.println(">")
	} else {
		pp.print(pp.keyword("ARRAY") + "<" + typespec + ">")
	}
	pp.lnaccept(ctx, ast.Must(n.Collate()))
	pp.lnaccept(ctx, ast.Must(n.GeneratedColumnInfo()))
	pp.lnaccept(ctx, ast.Must(n.DefaultExpression()))
	pp.lnaccept(ctx, ast.Must(n.Attributes()))
	pp.lnaccept(ctx, ast.Must(n.OptionsList()))
	p.print(pp.unnestLeft())
}

func (p *Printer) visitArrayColumnSchema(ctx Context, n *googlesql.ASTColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("ARRAY") + "<")
	simpleType, _ := isSimpleColumnSchema(n)
	if !simpleType {
		pp.println("")
		p2 := pp.nest()
		p2.accept(ctx, ast.Must(n.TypeParameters()))
		pp.print(p2.unnestLeft())
	} else {
		pp.accept(ctx, ast.Must(n.TypeParameters()))
	}
	pp.print(">")
	pp.lnaccept(ctx, ast.Must(n.Collate()))
	pp.lnaccept(ctx, ast.Must(n.GeneratedColumnInfo()))
	pp.lnaccept(ctx, ast.Must(n.DefaultExpression()))
	pp.lnaccept(ctx, ast.Must(n.Attributes()))
	pp.lnaccept(ctx, ast.Must(n.OptionsList()))
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitArrayType(ctx Context, n *googlesql.ASTArrayType) {
	pp := p.nest()
	pp.moveBefore(n)
	simple := true
	if et := ast.Must(n.ElementType()); et != nil {
		simple = isSimpleType(et)
	}
	pp2 := pp.nest()
	pp2.accept(ctx, ast.Must(n.ElementType()))
	pp2.accept(ctx, ast.Must(n.TypeParameters()))
	if simple {
		elemType := strings.Trim(pp2.String(), "\n")
		pp.print(pp.keyword("ARRAY") + "<" + elemType + ">")
	} else {
		pp.println(pp.keyword("ARRAY") + "<")
		pp.incDepth()
		pp.print(pp2.unnest())
		pp.println("")
		pp.decDepth()
		pp.print(">")
	}
	pp.accept(ctx, ast.Must(n.Collate()))
	p.print(pp.String())
}

func (p *Printer) VisitColumnDefinition(ctx Context, n *googlesql.ASTColumnDefinition) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Name()))
	p.print("\v")
	p.acceptNestedString(ctx, ast.Must(n.Schema()))
}

func (p *Printer) VisitColumnSchema(ctx Context, n *googlesql.ASTColumnSchema) {
	p.moveBefore(n)
	// In ZetaSQL, ASTColumnSchema is an abstract class, and its
	// extensions are
	//
	//   - ASTArrayColumnSchema
	//   - ASTInferredTypeColumnSchema
	//   - ASTSimpleColumnSchema
	//   - ASTStructColumnSchema
	//
	// However, in Go bindings, we have a struct ast.ColumnSchemaNode
	// and ast.Nodes's of kind ast.ArrayColumnSchema,
	// ast.InferredTypeColumnSchema, ast.StructColumnSchema, and
	// ast.StructColumnSchema are mapped to *googlesql.ASTColumnSchema.
	//
	// Effectively, we cannot reach any of ast.ArrayColumnSchemaNode,
	// ast.InferredTypeColumnSchemaNode, ast.StructColumnSchemaNode,
	// or ast.ArrayColumnSchemaNode by walking with Child() methods.
	//
	// Issue: https://github.com/goccy/go-zetasql/issues/30
	//
	// We circumvent this issue by checking the node's kind and handling
	// children accordingly.
	switch ast.Kind(n) {
	case ast.ArrayColumnSchema:
		p.visitArrayColumnSchema(ctx, n)
	case ast.InferredTypeColumnSchema:
		p.visitInferredTypeColumnSchema(ctx, n)
	case ast.SimpleColumnSchema:
		panic("wtf")
		// p.VisitSimpleColumnSchema(ctx, n.(*googlesql.ASTSimpleColumnSchema))
	case ast.StructColumnSchema:
		p.visitStructColumnSchema(ctx, n)
	default:
		panic(fmt.Errorf("unexpected kind for column schema node")) // Node: n
	}
	p.movePast(n)
}

func (p *Printer) visitInferredTypeColumnSchema(ctx Context, n *googlesql.ASTColumnSchema) {
	p.addError(fmt.Errorf("not implemented")) // Node: n
}

func (p *Printer) VisitSimpleColumnSchema(ctx Context, n *googlesql.ASTSimpleColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx.WithValue(KeyInTypeName, true), ast.Must(n.TypeName()))
	p1.accept(ctx, ast.Must(n.TypeParameters()))
	pp.print(p1.unnestLeft())
	pp.lnaccept(ctx, ast.Must(n.Collate()))
	pp.lnaccept(ctx, ast.Must(n.Attributes()))
	pp.lnaccept(ctx, ast.Must(n.OptionsList()))
	p.print(pp.unnest())
	p.movePast(n)
}

func (p *Printer) VisitSimpleType(ctx Context, n *googlesql.ASTSimpleType) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyInTypeName, true), ast.Must(n.TypeName()))
	p.accept(ctx, ast.Must(n.TypeParameters()))
	p.accept(ctx, ast.Must(n.Collate()))
}

func (p *Printer) VisitStructColumnField(ctx Context, n *googlesql.ASTStructColumnField) {
	p.moveBefore(n)
	// When a struct is inside an array, KeyInTypeName is true, so we need to
	// override its value to avoid formatting the struct field name as a type name.
	p.accept(ctx.WithValue(KeyInTypeName, false), ast.Must(n.Name()))
	p.accept(ctx, ast.Must(n.Schema()))
	p.movePast(n)
}

func (p *Printer) VisitStructColumnSchema(ctx Context, n *googlesql.ASTStructColumnSchema) {
	pp := p.nest()
	pp.moveBefore(n)
	fields := ast.ChildrenOfType[*googlesql.ASTStructColumnField](n)
	simple := isSimpleStructColumnSchema(fields)
	if simple {
		p2 := pp.nest()
		for i, c := range fields {
			if i > 0 {
				p2.print(",")
			}
			p2.accept(ctx, c)
		}
		pp.print(pp.keyword("STRUCT") + "<" + p2.unnestLeft() + ">")
	} else {
		pp.println(pp.keyword("STRUCT") + "<")
		pp.incDepth()
		for i, c := range fields {
			if i > 0 {
				pp.println(",")
			}
			pp.moveBefore(c)
			pp.accept(ctx, c)
			pp.movePast(c)
		}
		pp.decDepth()
		pp.println("")
		pp.println(">")
	}
	pp.accept(ctx, ast.Must(n.Collate()))
	pp.accept(ctx, ast.Must(n.Attributes()))
	pp.accept(ctx, ast.Must(n.OptionsList()))
	p.print(pp.unnestLeft())
}

func (p *Printer) visitStructColumnSchema(ctx Context, n *googlesql.ASTColumnSchema) {
	pp := p.nest()
	simpleType, _ := isSimpleColumnSchema(n)
	p1 := pp.nest()
	fields := ast.ChildrenOfType[*googlesql.ASTStructColumnField](n)
	for i, f := range fields {
		if i > 0 {
			p1.println(",")
		}
		p1.accept(ctx, f)
	}
	if simpleType {
		pp.print(pp.keyword("STRUCT") + "<" + p1.unnestLeft() + ">")
		pp.println(pp.keyword("STRUCT") + "<")
		pp.incDepth()
		pp.println(p1.unnestLeft())
		pp.decDepth()
		pp.print(">")
	}
	attrs := ast.ChildrenOfType[*googlesql.ASTColumnAttributeList](n)
	if len(attrs) > 0 {
		printNestedWithSep(ctx, pp, attrs, "")
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitStructField(ctx Context, n *googlesql.ASTStructField) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Name()))
	p.acceptNested(ctx, ast.Must(n.Type()))
}

func (p *Printer) VisitStructType(ctx Context, n *googlesql.ASTStructType) {
	pp := p.nest()
	pp.moveBefore(n)
	fields := ast.ChildrenOfType[*googlesql.ASTStructField](n)
	simple := allTrue(mapIsSimpleStructFields(fields))
	pp2 := pp.nest()
	for i, f := range fields {
		if i > 0 {
			pp2.print(",")
			if !simple {
				pp2.println("")
			}
		}
		pp2.accept(ctx, f)
	}
	elemType := pp2.unnestLeft()
	pp3 := pp.nest()
	if simple {
		pp3.print(pp3.keyword("STRUCT") + "<" + elemType + ">")
	} else {
		pp3.print(pp3.keyword("STRUCT") + "<")
		pp3.println("")
		pp3.incDepth()
		pp3.print(elemType)
		pp3.println("")
		pp3.decDepth()
		pp3.print(">")
	}
	pp3.accept(ctx, ast.Must(n.TypeParameters()))
	pp.print(pp3.unnestLeft())
	pp.lnaccept(ctx, ast.Must(n.Collate()))
	p.print(pp.unnest())
}

func (p *Printer) VisitTemplatedParameterType(ctx Context, n *googlesql.ASTTemplatedParameterType) {
	p.moveBefore(n)
	switch ast.Must(n.Kind()) {
	case ast.UninitializedTemplatedTypeKind:
		p.print(p.keyword("<UNINITIALIZED TEMPLATED KIND>"))
	case ast.AnyTypeTemplatedTypeKind:
		p.print(p.keyword("ANY TYPE"))
	case ast.AnyProtoTemplatedTypeKind:
		p.print(p.keyword("ANY PROTO"))
	case ast.AnyEnumTemplatedTypeKind:
		p.print(p.keyword("ANY ENUM"))
	case ast.AnyStructTemplatedTypeKind:
		p.print(p.keyword("ANY STRUCT"))
	case ast.AnyArrayTemplatedTypeKind:
		p.print(p.keyword("ANY ARRAY"))
	case ast.AnyTableTemplatedTypeKind:
		p.print(p.keyword("ANY TABLE"))
	}
	p.movePast(n)
}

// This is a patch to format TemplatedParameterTypes and Table types,
// which are not accessible in go-zetasql for now.
func (p *Printer) patchedVisitTemplatedParameterType(n *googlesql.ASTFunctionParameter) {
	input := p.nodeErasedInput(n)
	inputUpcase := strings.ToUpper(input)
	field := p.toString(nil, ast.Must(n.Name()))
	i := strings.Index(inputUpcase, strings.ToUpper(field))
	// We just print whatever we find after the field name as a typename.
	p.print(p.typename(strings.TrimSpace(input[i+len(field)+1:])))
	// Check if this is one of the expected bug.
	types := []string{
		"ANY TYPE", "ANY PROTO", "ANY ENUM", "ANY STRUCT",
		"ANY ARRAY", "ANY TABLE", "TABLE",
	}
	for _, candidate := range types {
		if strings.Contains(inputUpcase, candidate) {
			return
		}
	}
	panic(fmt.Sprintf("Unsupported type in input %#v", input))
}

func (p *Printer) VisitTVFSchema(ctx Context, n *googlesql.ASTTVFSchema) {
	cols := ast.ChildrenOfType[*googlesql.ASTTVFSchemaColumn](n)
	simple := len(cols) <= 2 && allTrue(mapIsSimpleTVFSchema(cols))
	pp := p.nest()
	pp.moveBefore(n)
	p1 := pp.nest()
	for i, e := range cols {
		if i > 0 {
			p1.print(",")

			if !simple {
				p1.println("")
			}
		}
		p1.accept(ctx, e)
	}
	if simple {
		pp.print(pp.keyword("TABLE") + "<" + p1.unnestLeft() + ">")
	} else {
		pp.println(pp.keyword("TABLE") + "<")
		pp.incDepth()
		pp.println(p1.unnestLeft())
		pp.decDepth()
		pp.print(">")
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitTVFSchemaColumn(ctx Context, n *googlesql.ASTTVFSchemaColumn) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.accept(ctx, ast.Must(n.Name()))
	pp.acceptNested(ctx, ast.Must(n.Type()))
	p.print(pp.String())
}
