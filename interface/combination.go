package main

import "fmt"

type BinLog interface {
	insert()
	update()
	Dispatch()
}

type Base struct {
}

func (b *Base) insert() {
	fmt.Println("base insert")
}

func (b *Base) update() {
	fmt.Println("base update")
}

func (s *Sub) Dispatch() {
	s.insert()
	s.update()
}

type Sub struct {
	*Base
}

func (s *Sub) insert() {
	fmt.Println("sub insert")
}
func main() {
	var h = &Sub{
		&Base{},
	}
	h.Dispatch()
}
