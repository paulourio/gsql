package parser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/paulourio/gsql/internal/extensions/ast"
	"github.com/paulourio/gsql/internal/extensions/errors"
	"github.com/paulourio/gsql/internal/extensions/lexer"
	"github.com/paulourio/gsql/internal/extensions/parser"
)

func TestParser(t *testing.T) {
	t.Parallel()
	tests := MustReadTest("testdata/comments.toml")

	for i, tcase := range tests.Cases {
		var name string

		if tcase.Description == "" {
			name = fmt.Sprintf("case %d", i)
		} else {
			name = fmt.Sprintf("case %d: %s", i, tcase.Description)
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var b strings.Builder
			c := tcase

			b.Grow(len(c.Input) * 2)
			fmt.Fprintf(&b, "Input (%d bytes):\n%s\n\n", len(c.Input), c.Input)
			fmt.Fprintf(&b, "Expected Dump (%d bytes):\n%s\n\n", len(c.Dump), c.Dump)

			l := lexer.NewLexer([]byte(c.Input))
			p := parser.NewParser()

			res, err := p.Parse(l)
			if err != nil {
				msg := errors.FormatError(err, c.Input)

				fmt.Fprintf(&b, "Parse error:\n%s\n\n", msg)
				assert.NoError(t, err, b.String())

				return
			}

			var dump string

			if nh, ok := res.(ast.NodeHandler); ok {
				dump = nh.DebugString("")
			} else {
				dump = fmt.Sprintf("%#v", res)
			}

			fmt.Fprintf(&b, "Result Dump (%d bytes):\n%s\n\n", len(dump), dump)

			assert.Equal(t, c.Dump, dump, b.String())
		})
	}
}
