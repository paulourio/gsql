package printer_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/goccy/go-googlesql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulourio/gsql/internal/printer"
	"github.com/paulourio/gsql/internal/sql"
)

func TestMain(m *testing.M) {
	// Initialize googlesql
	cacheDir, err := os.MkdirTemp("", "gsql")
	if err != nil {
		log.Fatalf("failed to create cache dir: %v", err)
	}
	defer os.RemoveAll(cacheDir)

	err = googlesql.Init(
	// googlesql.WithCompilationMode(googlesql.CompilationModeCompiler),
	// googlesql.WithCompilationCache(cacheDir),
	)
	if err != nil {
		log.Fatalf("failed to init googlesql: %v", err)
	}
	// defer googlesql.Close()

	os.Exit(m.Run())
}

func parseSQL(t *testing.T, sqlStr string) sql.Node {
	t.Helper()
	popts, err := googlesql.NewParserOptions()
	if err != nil {
		t.Fatalf("unable to create parser options: %v", err)
	}
	lopts, err := googlesql.NewLanguageOptionsMaximumFeatures()
	if err != nil {
		t.Fatalf("unable to create language options: %v", err)
	}
	if err := popts.SetLanguageOptions(lopts); err != nil {
		t.Fatalf("unable to set language options: %v", err)
	}
	pout, err := googlesql.ParseScript(sqlStr, popts, nil)
	if err != nil {
		t.Fatalf("unable to parse script: %v", err)
	}
	root, err := pout.Script()
	if err != nil {
		t.Fatalf("unable to get script: %v", err)
	}
	return sql.Wrap(root)
}

func TestLocationTracker(t *testing.T) {
	t.Parallel()

	input := "select 1 from `table` where true"
	root := parseSQL(t, input)

	lt := printer.NewStartLocationTracker(input, root)

	// Check if Pos is not empty and is sorted
	require.NotEmpty(t, lt.Pos)
	isSorted := true
	for i := 1; i < len(lt.Pos); i++ {
		if lt.Pos[i] < lt.Pos[i-1] {
			isSorted = false
			break
		}
	}
	require.True(t, isSorted, "Pos slice should be sorted")

	// Validate unique elements in Pos
	isUnique := true
	for i := 1; i < len(lt.Pos); i++ {
		if lt.Pos[i] == lt.Pos[i-1] {
			isUnique = false
			break
		}
	}
	require.True(t, isUnique, "Pos slice should contain unique elements")

	// Expected Pos values based on the query "select 1 from `table` where true"
	// Script, StatementList, QueryStatement, Query, Select start at 0.
	// SelectList, SelectColumn, IntLiteral start at 7.
	// FromClause starts at 9.
	// TablePathExpression, PathExpression, Identifier start at 14.
	// WhereClause starts at 22.
	// BooleanLiteral starts at 28.
	expectedPos := []int{0, 7, 9, 14, 22, 28}
	require.Equal(t, expectedPos, lt.Pos)

	// Test MaybeNextPos and NextPos
	testCases := []struct {
		pos      int
		expected int
	}{
		{0, 7},
		{3, 7},
		{7, 9},
		{8, 9},
		{9, 14},
		{14, 22},
		{22, 28},
		{28, -1},
		{100, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Pos_%d", tc.pos), func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.expected, lt.MaybeNextPos(tc.pos))
			if tc.expected == -1 {
				require.Equal(t, tc.pos, lt.NextPos(tc.pos))
			} else {
				require.Equal(t, tc.expected, lt.NextPos(tc.pos))
			}
		})
	}
}

func TestLineTrackerSpan(t *testing.T) {
	t.Parallel()

	input := "select 1\nfrom `table`\nwhere true"
	root := parseSQL(t, input)

	lt := printer.NewLineTracker(input)

	// Script covers whole input [0, 2]
	s, e := lt.Span(root)
	require.Equal(t, 0, s)
	require.Equal(t, 2, e)

	stmtList := root.Child(0)
	queryStmt := stmtList.Child(0)
	query := queryStmt.Child(0)
	sel := query.Child(0)

	numSelChildren := sel.NumChildren()
	require.Equal(t, 3, numSelChildren)

	selectList := sel.Child(0)
	fromClause := sel.Child(1)
	whereClause := sel.Child(2)

	// selectList: "1" at line 0
	slStart, slEnd := lt.Span(selectList)
	require.Equal(t, 0, slStart)
	require.Equal(t, 0, slEnd)

	// fromClause: "from `table`" at line 1
	fcStart, fcEnd := lt.Span(fromClause)
	require.Equal(t, 1, fcStart)
	require.Equal(t, 1, fcEnd)

	// whereClause: "where true" at line 2
	wcStart, wcEnd := lt.Span(whereClause)
	require.Equal(t, 2, wcStart)
	require.Equal(t, 2, wcEnd)
}

func TestLineTracker(t *testing.T) {
	t.Parallel()
	for i, c := range lineLocCases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			t.Parallel()
			lt := printer.NewLineTracker(c.Input)
			cstr := fmt.Sprintf("Input: %#v\nStartPos: %#v\n", c.Input, c.StartPos)
			if assert.Equal(t, c.StartPos, lt.StartPos, cstr) {
				for j, q := range c.Queries {
					msg := fmt.Sprintf("%sTracker: %#v\nQuery %d: %#v",
						cstr, lt, j+1, q)
					require.Equal(t, q.line, lt.LineOf(q.byteOffset), msg)
					require.Equal(t, q.nextLineBreak, lt.NextLineBreak(q.byteOffset), msg)
				}
			}
		})
	}
}

var lineLocCases = []*lineLocCase{
	{
		Input:    "",
		StartPos: []int{},
		Queries: []*lineSpanQuery{
			{0, 0, -1},
			{10, 0, -1},
		},
	},
	{
		Input:    "\n",
		StartPos: []int{0},
		Queries: []*lineSpanQuery{
			{0, 0, 0},
			{1, 1, -1},
			{2, 1, -1},
		},
	},
	{
		Input:    "abc\n\ndef\nbar\n",
		StartPos: []int{3, 4, 8, 12},
		Queries: []*lineSpanQuery{
			{-1, 0, 3},  // out of bounds
			{0, 0, 3},   // "|abc^^def^bar^"
			{1, 0, 3},   // "a|bc^^def^bar^"
			{3, 0, 3},   // "abc|^^def^bar^"
			{4, 1, 4},   // "abc^|^def^bar^"
			{5, 2, 8},   // "abc^^|def^bar^"
			{7, 2, 8},   // "abc^^de|f^bar^"
			{9, 3, 12},  // "abc^^def^|bar^"
			{12, 3, 12}, // "abc^^def^bar|^"
			{13, 4, -1}, // "abc^^def^bar^|"
			{19, 4, -1}, // out of bounds
		},
	},
}

type lineLocCase struct {
	Input    string
	StartPos []int
	Queries  []*lineSpanQuery
}

type lineSpanQuery struct {
	byteOffset    int
	line          int
	nextLineBreak int
}
