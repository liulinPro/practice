// Author: LiuLin
// Date: 2022/3/29

package main

//
type temporary interface {
	Temporary() bool
}
type Error interface {
	error
	Timeout() bool
	Temporary() bool
}

func IsTemporary(err error) bool {
	t, ok := err.(temporary)
	return ok && t.Temporary()
}

func main() {
	var err error
	//
	if e, ok := err.(Error); ok && e.Temporary() {
		//todo
	}

	if IsTemporary(err) {
		// todo
	}
}
