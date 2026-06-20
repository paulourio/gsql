package sql_test

import (
	"testing"

	"github.com/goccy/go-googlesql"
	"github.com/paulourio/gsql/internal/ast"
	"github.com/paulourio/gsql/internal/sql"
)

func init() {
	err := googlesql.Init(
		googlesql.WithCompilationMode(googlesql.CompilationModeCompiler),
	)
	if err != nil {
		panic(err)
	}
}

func TestConvert(t *testing.T) {
	query := "SELECT a, b FROM t WHERE c = 1"
	opts, err := googlesql.NewParserOptions()
	if err != nil {
		t.Fatalf("failed to create parser options: %v", err)
	}

	errMsgOpts := &googlesql.ErrorMessageOptions{
		AttachErrorLocationPayload: true,
		InputOriginalStartColumn:   1,
		InputOriginalStartLine:     1,
		Mode:                       googlesql.ErrorMessageModeErrorMessageOneLine,
		Stability:                  googlesql.ErrorMessageStabilityTestMinimized,
	}

	z, err := googlesql.ParseScript(query, opts, errMsgOpts)
	if err != nil {
		t.Fatalf("failed to parse script: %v", err)
	}

	astScript := ast.Must(z.Script())

	node := sql.Convert(astScript)
	if node == nil {
		t.Fatalf("Convert returned nil")
	}

	script, ok := node.(*sql.Script)
	if !ok {
		t.Fatalf("expected *sql.Script, got %T", node)
	}

	if script.Kind() != sql.ScriptKind {
		t.Errorf("expected Kind ScriptKind, got %v", script.Kind())
	}

	// Verify LocationRange is set correctly
	loc := script.Loc()
	if loc.Start != 0 || loc.End != len(query) {
		t.Errorf("expected LocationRange [0, %d], got [%d, %d]", len(query), loc.Start, loc.End)
	}
}
