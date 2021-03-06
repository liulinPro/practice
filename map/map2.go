package main

// sort a map's keys in descending order of its values.

import "sort"
import "fmt"

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	fmt.Println(sm.s)
	sort.Sort(sm)
	return sm.s
}

func main() {
	s := []string{"Python", "Python", "Python", "igor", "igor", "igor", "igor", "go", "go", "Golang", "Golang", "Golang", "Golang", "Py", "Py"}
	count := make(map[string]int)

	for _, v := range s {
		count[v]++
	}

	for _, res := range sortedKeys(count) {
		fmt.Println(res, count[res])
	}
}
