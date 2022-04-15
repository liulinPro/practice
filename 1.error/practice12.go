// Author: LiuLin
// Date: 2022/4/14

package main

type HttpError struct {
	Path string
	User string
}

func (h *HttpError) Is() bool {
	return false
}
