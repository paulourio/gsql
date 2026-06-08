package main

import (
	"fmt"
	"os"

	"github.com/paulourio/gsql"
)

func main() {
	f, err := gsql.NewSQLFormatter()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	sql, err := f.Format("select 1 from `table` where true")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(sql)
}
