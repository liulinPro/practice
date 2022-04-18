// Author: LiuLin
// Date: 2022/4/14

package main

import (
	"errors"
	"fmt"
	xerror "github.com/pkg/errors"
)

type QueryError struct {
	Query string
	Err   error
}

func (q *QueryError) Error() string {
	return q.Query
}

var qe = &QueryError{
	Query: "hello",
}

func main() {
	err := fmt.Errorf("hello %w", qe)
	if errors.Is(err, qe) {
		fmt.Println(err)
	}
	if xerror.Is(err, qe) {
		fmt.Println(err)
	}
	q := new(QueryError)
	if errors.As(err, &q) {
		fmt.Println(q)
	}
	queryError := QueryError{
		Query: "hello",
	}
	fmt.Println(&queryError == qe)
}
