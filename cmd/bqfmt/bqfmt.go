package main

import (
	"fmt"
	"io"
	"os"

	"github.com/paulourio/gsql"
)

func main() {
	f, err := gsql.NewSQLFormatter()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	var input []byte
	if len(os.Args) > 1 {
		input, err = os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file:", err)
			os.Exit(1)
		}
	} else {
		// Read from stdin if no file is provided
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
			os.Exit(1)
		}
	}

	sql, err := f.Format(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Formatting error:", err)
		os.Exit(1)
	}
	fmt.Print(sql)
}
