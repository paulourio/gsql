package printer

import (
	"fmt"
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/format"
	"github.com/paulourio/gsql/internal/ast"
)

func (p *Printer) VisitBigNumericLiteral(ctx Context, n *googlesql.ASTBigNumericLiteral) {
	p.moveBefore(n)
	p.print(p.keyword("BIGNUMERIC"))
	p.accept(ctx, ast.Must(n.StringLiteral()))
}

func (p *Printer) VisitBoolLiteral(ctx Context, n *googlesql.ASTBooleanLiteral) {
	p.moveBefore(n)
	p.print(formatPrintStyle(ast.Must(n.Image()), p.Writer.opts.BoolStyle))
}

func (p *Printer) VisitBytesLiteral(ctx Context, n *googlesql.ASTBytesLiteral) {
	p.moveBefore(n)
	val := p.nodeInput(n)
	s, err := FormatBytes(val, p.Writer.opts.BytesStyle)
	if err != nil {
		p.addError(&Error{
			Msg: fmt.Sprintf("%v: %q", err, val),
		})
	}
	p.print(strings.ReplaceAll(s, "\n", lineBreakPlaceholder))
}

func (p *Printer) VisitDateOrTimeLiteral(ctx Context, n *googlesql.ASTDateOrTimeLiteral) {
	p.moveBefore(n)
	// There's a bug in the mapping of TypeKind to actual type.
	// For instance, TIMESTAMP (11) is being mapped as NUMERIC (19).
	// For now, the safest approach seems to re-tokenize the node input.
	input := p.nodeInput(n)
	pos := strings.Index(input, " ")
	if pos < 0 {
		panic("invalid date time literal")
	}
	switch ast.Must(n.TypeKind()) {
	case ast.Date:
		p.print(p.keyword("DATE"))
	case ast.Datetime:
		p.print(p.keyword("DATETIME"))
	case ast.Time:
		p.print(p.keyword("TIME"))
	case ast.Timestamp:
		p.print(p.keyword("TIMESTAMP"))
	default:
		p.addError(&Error{
			Msg: fmt.Sprintf("failed to parse date time kind: %v", ast.Must(n.TypeKind())),
		})
	}
	p.accept(ctx, ast.Must(n.StringLiteral()))
}

func (p *Printer) VisitFloatLiteral(ctx Context, n *googlesql.ASTFloatLiteral) {
	p.moveBefore(n)
	p.print(strings.ToLower(ast.Must(n.Image())))
}

func (p *Printer) VisitIntLiteral(ctx Context, n *googlesql.ASTIntLiteral) {
	p.moveBefore(n)
	v := ast.Must(n.Image())
	if !maybeHexValue(v) {
		p.print(v)
	} else {
		p.print("0x" + formatPrintStyle(v[2:], p.Writer.opts.HexStyle))
	}
	p.movePast(n)
}

func (p *Printer) VisitJSONLiteral(ctx Context, n *googlesql.ASTJSONLiteral) {
	p.moveBefore(n)
	p.print(p.keyword("JSON"))
	p.accept(ctx, ast.Must(n.StringLiteral()))
	p.movePast(n)
}

func (p *Printer) VisitNullLiteral(ctx Context, n *googlesql.ASTNullLiteral) {
	p.moveBefore(n)
	p.print(formatPrintStyle(ast.Must(n.Image()), p.Writer.opts.NullStyle))
}

func formatPrintStyle(s string, style format.PrintCase) string {
	switch style {
	case format.Unspecified, format.AsIs:
		return s
	case format.LowerCase:
		return strings.ToLower(s)
	case format.UpperCase:
		return strings.ToUpper(s)
	}
	panic(fmt.Sprintf("invalid print style %#v", style))
}

func (p *Printer) VisitNumericLiteral(ctx Context, n *googlesql.ASTNumericLiteral) {
	p.moveBefore(n)
	p.print(p.keyword("NUMERIC"))
	p.accept(ctx, ast.Must(n.StringLiteral()))
}

func (p *Printer) VisitStringLiteral(ctx Context, n *googlesql.ASTStringLiteral) {
	p.moveBefore(n)
	val := p.nodeInput(n)
	s, err := FormatString(val, p.Writer.opts.StringStyle)
	if err != nil {
		p.addError(&Error{
			Msg: fmt.Sprintf("%v: %q", err, val),
		})
	}
	p.print(strings.ReplaceAll(s, "\n", lineBreakPlaceholder))
}

func FormatBytes(s string, style format.StringStyle) (string, error) {
	isBytes := maybeBytesLiteral(s)
	isRaw := maybeRawBytesLiteral(s)
	if !isBytes && !isRaw {
		return "", ErrInvalidStringLiteral
	}
	if style == format.AsIsStringStyle {
		return s, nil
	}
	offset := 0 // Offset to control the error position.
	prefix := ""
	noPrefix := s
	// Strip off the prefix from the raw string content before
	// parsing.
	if isRaw {
		prefix = "rb"
		noPrefix = noPrefix[2:]
		offset = 2
	} else {
		prefix = "b"
		noPrefix = noPrefix[1:]
		offset = 1
	}
	quotesLen := 1
	isTripleQuoted := maybeTripleQuotedStringLiteral(noPrefix)
	isSingleQuote := isSingleQuote(noPrefix)
	if isTripleQuoted {
		quotesLen = 3
	}
	offset += quotesLen
	content := s[offset : len(s)-quotesLen]
	switch style {
	case format.PreferSingleQuote:
		if isSingleQuote || strings.Contains(content, "'") {
			return prefix + s[len(prefix):], nil
		}
		if isTripleQuoted {
			return fmt.Sprintf("%s'''%s'''", prefix, content), nil
		}
		return fmt.Sprintf("%s'%s'", prefix, content), nil
	case format.PreferDoubleQuote:
		if isSingleQuote || strings.Contains(content, `"`) {
			return s, nil
		}
		if isTripleQuoted {
			return fmt.Sprintf(`%s"""%s"""`, prefix, content), nil
		}
		return fmt.Sprintf(`%s"%s"`, prefix, content), nil
	}
	return "", ErrInvalidStringStyle
}

func FormatString(s string, style format.StringStyle) (string, error) {
	isStr := maybeStringLiteral(s)
	isRaw := maybeRawStringLiteral(s)
	if !isStr && !isRaw {
		return "", ErrInvalidStringLiteral
	}
	if style == format.AsIsStringStyle {
		return s, nil
	}
	offset := 0 // Offset to control the error position.
	prefix := ""
	noPrefix := s
	if isRaw {
		// Strip off the prefix 'r' from the raw string content before
		// parsing.
		prefix = "r"
		noPrefix = noPrefix[1:]
		offset = 1
	}
	quotesLen := 1
	isTripleQuoted := maybeTripleQuotedStringLiteral(noPrefix)
	isSingleQuote := isSingleQuote(noPrefix)
	if isTripleQuoted {
		quotesLen = 3
	}
	offset += quotesLen
	content := s[offset : len(s)-quotesLen]
	switch style {
	case format.PreferSingleQuote:
		if isSingleQuote || strings.Contains(content, "'") {
			return prefix + s[len(prefix):], nil
		}
		if isTripleQuoted {
			return fmt.Sprintf("%s'''%s'''", prefix, content), nil
		}
		return fmt.Sprintf("%s'%s'", prefix, content), nil
	case format.PreferDoubleQuote:
		if isSingleQuote || strings.Contains(content, `"`) {
			return s, nil
		}
		if isTripleQuoted {
			return fmt.Sprintf(`%s"""%s"""`, prefix, content), nil
		}
		return fmt.Sprintf(`%s"%s"`, prefix, content), nil
	}
	return "", ErrInvalidStringStyle
}

func isSingleQuote(s string) bool {
	return s[0] == '\''
}

func maybeTripleQuotedStringLiteral(s string) bool {
	if len(s) >= 6 &&
		(strings.HasPrefix(s, "'''") && strings.HasSuffix(s, "'''") ||
			strings.HasPrefix(s, `"""`) && strings.HasSuffix(s, `"""`)) {
		return true
	}
	return false
}

func maybeStringLiteral(s string) bool {
	if (len(s) >= 2) &&
		(s[0] == s[len(s)-1]) &&
		(s[0] == '\'' || s[0] == '"') {
		return true
	}
	return false
}

func maybeRawStringLiteral(s string) bool {
	if (len(s) >= 3) &&
		(s[0] == 'r' || s[0] == 'R') &&
		(s[1] == s[len(s)-1]) &&
		(s[1] == '\'' || s[1] == '"') {
		return true
	}
	return false
}

func maybeBytesLiteral(s string) bool {
	if (len(s) >= 3) &&
		(s[0] == 'b' || s[0] == 'B') &&
		(s[1] == s[len(s)-1]) &&
		(s[1] == '\'' || s[1] == '"') {
		return true
	}
	return false
}

func maybeRawBytesLiteral(s string) bool {
	if len(s) >= 4 {
		low := strings.ToLower(s[:2])
		if (low == "rb" || low == "br") &&
			(s[2] == s[len(s)-1]) &&
			(s[2] == '\'' || s[2] == '"') {
			return true
		}
	}
	return false
}

func maybeHexValue(s string) bool {
	// Note that hex values are always unsigned, and -0xA will be
	// parsed with a unary operator applied to the int literal.
	return len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}
