// Author: LiuLin
// Date: 2022/3/30

package main

import (
	"fmt"
	"io"
)

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, e.err
}

func WriteResponse1(w io.Writer, st Status, headers []Header, body io.Reader) error {
	e := &errWriter{Writer: w}
	fmt.Fprintf(e, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	for _, h := range headers {
		fmt.Fprintf(e, "%s: %s\r\n", h.Key, h.Value)
	}
	fmt.Fprintf(e, "\r\n")
	io.Copy(e, body)
	return e.err
}
