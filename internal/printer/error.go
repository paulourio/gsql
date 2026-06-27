package printer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

// Error represents and error during printing at a specific node.
// The error message includes the location of the node in the input sql if provided.
type Error struct {
	Msg   string
	Err   error
	Node  sql.Node
	Input *string
}

var (
	ErrInvalidBytesLiteral  = errors.New("invalid bytes literal")
	ErrInvalidStringLiteral = errors.New("invalid string literal")
	ErrInvalidStringStyle   = errors.New("invalid string style")
)

func (e *Error) Error() string {
	parts := []string{"PrinterError"}
	if e.Node != nil && e.Node.Raw() != nil {
		locStr, err := e.Node.Raw().GetLocationString()
		if err != nil {
			panic(err)
		}
		kindStr, err := e.Node.Raw().GetNodeKindString()
		if err != nil {
			panic(err)
		}
		parts = append(parts, fmt.Sprintf("at %s (%s)", locStr, kindStr))
	}
	if e.Msg != "" {
		parts = append(parts, e.Msg)
	}
	if e.Err != nil {
		parts = append(parts, e.Err.Error())
	}
	if e.Node != nil && e.Input != nil {
		begin, end := e.Node.Location()
		s := (*e.Input)[begin:end]
		parts = append(parts, s)
	}
	return strings.Join(parts, ": ")
}
