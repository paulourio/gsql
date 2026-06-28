// Functions to print types and schemas.
package printer

import (
	"fmt"
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitArrayColumnSchema(ctx Context, n *sql.ArrayColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx.WithValue(KeyInTypeName, true), n.ElementSchema())
	p1.accept(ctx, n.TypeParameters())
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
	pp.lnaccept(ctx, n.Collate())
	pp.lnaccept(ctx, n.GeneratedColumnInfo())
	pp.lnaccept(ctx, n.DefaultExpression())
	pp.lnaccept(ctx, n.Attributes())
	pp.lnaccept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitColumnSchemaAsArrayColumnSchema(ctx Context, n *sql.ColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("ARRAY") + "<")
	if !isSimpleColumnSchema(n) {
		pp.println("")
		p2 := pp.nest()
		p2.accept(ctx, n.TypeParameters())
		pp.print(p2.unnestLeft())
	} else {
		pp.accept(ctx, n.TypeParameters())
	}
	pp.print(">")
	pp.lnaccept(ctx, n.Collate())
	pp.lnaccept(ctx, n.GeneratedColumnInfo())
	pp.lnaccept(ctx, n.DefaultExpression())
	pp.lnaccept(ctx, n.Attributes())
	pp.lnaccept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitArrayType(ctx Context, n *sql.ArrayType) {
	pp := p.nest()
	pp.moveBefore(n)
	simple := true
	if et := n.ElementType(); et != nil {
		simple = isSimpleType(et)
	}
	pp2 := pp.nest()
	pp2.accept(ctx, n.ElementType())
	pp2.accept(ctx, n.TypeParameters())
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
	pp.accept(ctx, n.Collate())
	p.print(pp.String())
}

func (p *Printer) visitColumnDefinition(ctx Context, n *sql.ColumnDefinition) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
	p.print("\v")
	p.acceptNestedString(ctx, n.Schema())
}

func (p *Printer) visitColumnSchema(ctx Context, n *sql.ColumnSchema) {
	p.moveBefore(n)
	switch n.Kind() {
	case sql.ArrayColumnSchemaKind:
		p.visitColumnSchemaAsArrayColumnSchema(ctx, n)
	case sql.InferredTypeColumnSchemaKind:
		p.visitInferredTypeColumnSchema(ctx, n)
	case sql.SimpleColumnSchemaKind:
		panic("not implemented")
	case sql.StructColumnSchemaKind:
		p.visitColumnSchemaAsStructColumnSchema(ctx, n)
	default:
		panic(fmt.Errorf("unexpected kind for column schema node"))
	}
	p.movePast(n)
}

func (p *Printer) visitInferredTypeColumnSchema(_ Context, _ *sql.ColumnSchema) {
	p.addError(fmt.Errorf("not implemented"))
}

func (p *Printer) visitRangeType(ctx Context, n *sql.RangeType) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx, n.ElementType())
	p1.accept(ctx, n.TypeParameters())
	pp.print(pp.keyword("RANGE") + "<" + p1.unnestLeft() + ">")
	pp.accept(ctx, n.Collate())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitSimpleColumnSchema(ctx Context, n *sql.SimpleColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx.WithValue(KeyInTypeName, true), n.TypeName())
	p1.accept(ctx, n.TypeParameters())
	pp.print(p1.unnestLeft())
	pp.lnaccept(ctx, n.Collate())
	pp.lnaccept(ctx, n.Attributes())
	pp.lnaccept(ctx, n.OptionsList())
	p.print(pp.unnest())
	p.movePast(n)
}

func (p *Printer) visitSimpleType(ctx Context, n *sql.SimpleType) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyInTypeName, true), n.TypeName())
	p.accept(ctx, n.TypeParameters())
	p.accept(ctx, n.Collate())
}

func (p *Printer) visitStructColumnField(ctx Context, n *sql.StructColumnField) {
	p.moveBefore(n)
	// When a struct is inside an array, KeyInTypeName is true, so we need to
	// override its value to avoid formatting the struct field name as a type name.
	p.accept(ctx.WithValue(KeyInTypeName, false), n.Name())
	p.accept(ctx, n.Schema())
	p.movePast(n)
}

func (p *Printer) visitStructColumnSchema(ctx Context, n *sql.StructColumnSchema) {
	pp := p.nest()
	pp.moveBefore(n)
	fields := n.StructFields()
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
	pp.accept(ctx, n.Collate())
	pp.accept(ctx, n.Attributes())
	pp.accept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitColumnSchemaAsStructColumnSchema(ctx Context, n *sql.ColumnSchema) {
	pp := p.nest()
	p1 := pp.nest()
	var fields []*sql.StructColumnField
	if sc, ok := sql.Wrap(n.Raw()).(*sql.StructColumnSchema); ok {
		fields = sc.StructFields()
	}
	for i, f := range fields {
		if i > 0 {
			p1.println(",")
		}
		p1.accept(ctx, f)
	}
	if isSimpleColumnSchema(n) {
		pp.print(pp.keyword("STRUCT") + "<" + p1.unnestLeft() + ">")
		pp.println(pp.keyword("STRUCT") + "<")
		pp.incDepth()
		pp.println(p1.unnestLeft())
		pp.decDepth()
		pp.print(">")
	}
	if attrs := n.Attributes(); attrs != nil {
		printNestedWithSepNode(ctx, pp, []*sql.ColumnAttributeList{attrs}, "")
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitStructField(ctx Context, n *sql.StructField) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
	p.acceptNested(ctx, n.Type())
}

func (p *Printer) visitStructType(ctx Context, n *sql.StructType) {
	pp := p.nest()
	pp.moveBefore(n)
	fields := n.StructFields()
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
	pp3.accept(ctx, n.TypeParameters())
	pp.print(pp3.unnestLeft())
	pp.lnaccept(ctx, n.Collate())
	p.print(pp.unnest())
}

func (p *Printer) visitTemplatedParameterType(_ Context, n *sql.TemplatedParameterType) {
	p.moveBefore(n)
	switch n.TemplatedKind() {
	case sql.UninitializedTypeKind:
		p.print(p.keyword("<UNINITIALIZED TEMPLATED KIND>"))
	case sql.AnyType:
		p.print(p.keyword("ANY TYPE"))
	case sql.AnyProto:
		p.print(p.keyword("ANY PROTO"))
	case sql.AnyEnum:
		p.print(p.keyword("ANY ENUM"))
	case sql.AnyStruct:
		p.print(p.keyword("ANY STRUCT"))
	case sql.AnyArray:
		p.print(p.keyword("ANY ARRAY"))
	case sql.AnyTable:
		p.print(p.keyword("ANY TABLE"))
	}
	p.movePast(n)
}

func (p *Printer) visitTVFSchema(ctx Context, n *sql.TVFSchema) {
	cols := n.Columns()
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

func (p *Printer) visitTVFSchemaColumn(ctx Context, n *sql.TVFSchemaColumn) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.accept(ctx, n.Name())
	pp.acceptNested(ctx, n.Type())
	p.print(pp.String())
}
