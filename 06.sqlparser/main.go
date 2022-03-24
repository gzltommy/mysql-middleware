package main

import (
	"github.com/xwb1989/sqlparser"
	"io"
	"strings"
)

func main() {
	//===============================================================================================
	sql := "SELECT * FROM table WHERE a = 'abc'"
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		// Do something with the err
	}

	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		_ = stmt
	case *sqlparser.Insert:
	}

	//===============================从 io.Reader 读取许多查询的替代方法：============================
	r := strings.NewReader("INSERT INTO table1 VALUES (1, 'a'); INSERT INTO table2 VALUES (3, 4);")
	tokens := sqlparser.NewTokenizer(r)
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		// Do something with stmt or err.
		_ = stmt
	}
}
