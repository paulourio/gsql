package printer

import (
	"sort"
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

// LineTracker maps AST nodes to line ranges in the original sql text.
type LineTracker struct {
	StartPos []int
}

// LocationTracker track an ordered sequence of nodes.
type LocationTracker struct {
	// Pos is an ordered slice of unique positions of start positions.
	Pos []int
	// Lines tracks the byte offsets of each line begin.
	Lines *LineTracker
}

// NewLineTracker returns a new LineTracker for the given input string.
func NewLineTracker(input string) *LineTracker {
	t := &LineTracker{}
	t.initialize(input)
	return t
}

// Span returns the range of lines a node spans.
func (t *LineTracker) Span(n googlesql.ASTNode) (start, end int) {
	b, e := ast.GetParseLocationByteOffsets(n)
	return t.LineOf(b), t.LineOf(e)
}

// LineOf returns the line a specific byte offset is located in.
func (t *LineTracker) LineOf(b int) int {
	return sort.Search(len(t.StartPos), func(i int) bool {
		return t.StartPos[i] >= b
	})
}

// NextLineBreak returns the byte offset of the next line break or
// -1 if not found.
func (t *LineTracker) NextLineBreak(offset int) int {
	i := t.LineOf(offset)
	if i < len(t.StartPos) {
		return t.StartPos[i]
	}
	return -1
}

func (t *LineTracker) initialize(s string) {
	n := strings.Count(s, "\n")
	t.StartPos = make([]int, n)
	offset := 0
	for i := 0; i < n; i++ {
		pos := strings.Index(s, "\n")
		t.StartPos[i] = offset + pos
		s = s[pos+1:]
		offset += pos + 1
	}
}

// NewStartLocationTracker returns a location tracker for an input and
// the respective parsed AST nodes.
func NewStartLocationTracker(s string, root googlesql.ASTNode) *LocationTracker {
	t := &LocationTracker{}
	t.initNodePos(root)
	t.initLines(s)
	return t
}

func (t *LocationTracker) initNodePos(root googlesql.ASTNode) {
	n := int(float64(countNodes(root)) * .6)
	set := make(map[int]bool, n)
	ast.Walk(root, func(n googlesql.ASTNode) error {
		if !ast.Defined(n) {
			return nil
		}
		b := ast.GetParseLocationStartOffset(n)
		set[b] = true
		return nil
	})
	t.Pos = make([]int, 0, len(set))
	for p := range set {
		t.Pos = append(t.Pos, p)
	}
	sort.Ints(t.Pos)
}

func (t *LocationTracker) initLines(s string) {
	t.Lines = NewLineTracker(s)
}

// NextPos returns the next position in the slice.  If not available,
// returns itself.
func (t *LocationTracker) NextPos(pos int) int {
	j := t.MaybeNextPos(pos)
	if j < 0 {
		return pos
	}
	return j
}

// MaybeNextPos returns the start position of the next node.  If not
// available, returns -1.
func (t *LocationTracker) MaybeNextPos(pos int) int {
	j := sort.Search(len(t.Pos), func(i int) bool { return t.Pos[i] > pos })
	if j == len(t.Pos) {
		return -1
	}
	return t.Pos[j]
}

func countNodes(root googlesql.ASTNode) (count int) {
	ast.Walk(root, func(n googlesql.ASTNode) error {
		count++
		return nil
	})
	return
}
