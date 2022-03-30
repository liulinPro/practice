package main

import "fmt"
import "sort"

func main() {
	m := map[string]int{
		"One":   100,
		"Two":   2,
		"Three": 3,
		"Ten":   10,
		"Fifty": 50,
	}
	vs := NewValSorter(m)
	fmt.Printf("%v\n", *vs)
	vs.Sort()
	fmt.Printf("%v\n", *vs)
}

type ValSorter struct {
	Keys []string
	Vals []int
}

func NewValSorter(m map[string]int) *ValSorter {
	vs := &ValSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]int, 0, len(m)),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *ValSorter) Sort() {
	sort.Sort(vs)
}

func (vs *ValSorter) Len() int           { return len(vs.Vals) }
func (vs *ValSorter) Less(i, j int) bool { return vs.Vals[i] < vs.Vals[j] }
func (vs *ValSorter) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}
