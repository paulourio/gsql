package printer

import (
	"strings"
	"unicode"

	"github.com/paulourio/gsql/format"
	"github.com/paulourio/gsql/internal/extensions"
)

type Writer struct {
	opts                *format.Options
	comments            *CommentsQueue
	buffer              strings.Builder
	formatted           strings.Builder
	maxLength           int
	depth               int
	last                rune
	lastWasUnary        bool // Last was a unary single chararacter.
	noFlushInNextFormat bool // Disables line flushing in the next call to Format()
	lastWasNewLine      bool
}

func NewWriter(fopts *format.Options, comms []*extensions.Comment) *Writer {
	return &Writer{
		opts:      fopts,
		comments:  &CommentsQueue{comments: comms},
		maxLength: fopts.SoftMaxColumns,
	}
}

func (w *Writer) WithCapacity(capacity int) *Writer {
	return &Writer{
		opts:      w.opts,
		comments:  w.comments,
		maxLength: capacity,
	}
}

// Format formats the string automatically according to the context.
//  1. Inserts necessary space between tokens.
//  2. Calls FlushLine() when a line reaches column limit and it is at
//     some point appropriate to break.
//
// The input s should not contain any leading or trailing whitespace,
// such as ' ' and '\n'.
func (w *Writer) Format(s string) {
	if len(s) == 0 {
		return
	}
	defer func() {
		w.lastWasNewLine = false
		w.lastWasUnary = false
		w.noFlushInNextFormat = false
		l := w.buffer.Len() - strings.LastIndex(w.buffer.String(), "\n")
		if l >= w.maxLength && w.lastIsSeparator() {
			w.FlushLine()
		}
	}()
	data := []rune(s)
	if w.buffer.Len() == 0 {
		w.writeIndent()
		w.writeRunes(data)
		return
	}
	switch w.last {
	case '\n':
		w.writeRunes(append([]rune{'\n'}, data...))
	case '(', '[', '@', '.', '~', ' ', '\v', '\b':
		w.writeRunes(data)
	default:
		if w.lastWasUnary {
			w.writeRunes(data)
			return
		}
		curr := data[0]
		if curr == '(' {
			// Inserts a space if last token is a separator, otherwise
			// regards it as a function call.
			if w.lastIsSeparator() {
				w.writeRunes(append([]rune{' '}, data...))
			} else {
				w.writeRunes(data)
			}
			return
		}
		if curr == ';' && w.lastWasNewLine {
			w.writeRunes(append([]rune{'\n'}, data...))
			return
		}
		if curr == ')' ||
			curr == '[' ||
			curr == ']' ||
			// To avoid case like "SELECT 1e10,.1e10"
			(curr == '.' && w.last != ',') ||
			(curr == ';' && w.last != '\n') ||
			curr == ',' {
			w.writeRunes(data)
			return
		}
		if w.last == ' ' && data[0] == ' ' {
			w.writeRunes(data)
		} else {
			w.writeRunes(append([]rune{' '}, data...))
		}
	}
}

// FormatLine is like Format, except always calls FlushLine.
// Use this if you explicitly wants to break the line after this string.
// For example:
//  1. To put a newline after SELECT:
//     FormatLine("SELECT")
//  2. To put close parenthesis on a separate line:
//     FormatLine("")
//     FormatLine(")")
func (w *Writer) FormatLine(s string) {
	w.Format(s)
	w.FlushLine()
}

// FlushLine flushes buffer to formatted, with a line break at the end.
// It will do nothing if it is a new line and buffer is empty, to avoid
// empty lines.
// Remember to call FlushLine once after the whole process is over in
// case some content remains in buffer.
func (w *Writer) FlushLine() {
	sfmt := w.formatted.String()
	sz := len(sfmt)
	if (sz == 0 || sfmt[sz-1] == '\n') && w.buffer.Len() == 0 {
		return
	}
	w.formatted.WriteString(w.buffer.String())
	w.formatted.WriteByte('\n')
	w.buffer.Reset()
	w.lastWasNewLine = true
	w.last = '\n'
}

// flushCommentsUpTo returns the number of comments flushed, and an
// indicator whether the last character is a newline.
func (w *Writer) flushCommentsUpTo(pos int) {
	lhs, rhs := extensions.SplitComments(w.comments.comments, pos)
	w.comments.comments = rhs
	for i, c := range lhs {
		// Contiguous comments will always be rendered on separate
		// lines.
		if i > 0 || c.IsOneline() {
			w.FlushLine()
		}
		image := c.Image
		if c.IsMultiline() {
			image = strings.ReplaceAll(c.Image, "\n", lineBreakPlaceholder)
		}
		if c.AtLineBegin() && w.buffer.Len() > 0 {
			w.FlushLine()
			w.Format(strings.TrimRight(image, "\n"))
		} else if !c.AtLineBegin() && (w.formatted.Len()+w.buffer.Len()) > 0 {
			// Add one additional space between line contents and
			// the comment at the final of current line.
			sp := " "
			if w.buffer.Len() == 0 || w.bufferEndsWithWhitespace() {
				sp = ""
			}
			w.Format(sp + strings.TrimRight(image, "\n"))
		} else {
			w.Format(strings.TrimRight(image, "\n"))
		}
		if c.MustEndLine() || c.AtLineEnd() {
			w.FlushLine()
		}
	}
}

func (w *Writer) bufferEndsWithWhitespace() bool {
	n := w.buffer.Len()
	if n == 0 {
		return false
	}
	s := w.buffer.String()
	last := s[n-1]
	return last == '\n' || unicode.IsSpace(rune(last))
}

func (w *Writer) lastIsSeparator() bool {
	if w.buffer.Len() == 0 {
		return false
	}
	if !isAlphanum(byte(w.last)) {
		return nonWordSeparators[w.last]
	}
	buf := w.buffer.String()
	i := len(buf) - 1
	for i >= 0 && isAlphanum(buf[i]) {
		i--
	}
	lastTok := buf[i+1:]
	return wordSeparators[lastTok]
}

func (w *Writer) addUnary(s string) {
	if w.lastWasUnary && w.last == '-' && s == "-" {
		w.lastWasUnary = false
	}
	w.Format(s)
	w.lastWasUnary = len(s) == 1
}

func (w *Writer) writeIndent() {
	w.buffer.WriteString(strings.Repeat(" ", w.depth*2))
	if w.depth > 0 {
		w.last = ' '
	}
}

func (w *Writer) writeRunes(d []rune) {
	w.buffer.WriteString(string(d))
	w.last = d[len(d)-1]
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isAlphanum(c byte) bool {
	return isLetter(c) || ('0' <= c && c <= '9') || (c == '_')
}

var wordSeparators = map[string]bool{
	"AND": true,
	"OR":  true,
	"ON":  true,
	"IN":  true,
}

var nonWordSeparators = map[rune]bool{
	',': true,
	'<': true,
	'>': true,
	'-': true,
	'+': true,
	'=': true,
	'*': true,
	'/': true,
	'%': true,
}

// lineBreakPlaceholder is used to avoid newline replacement issues
// in multi-line strings.
var lineBreakPlaceholder = string([]byte{33, 26, '1', '0'})
