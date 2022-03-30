package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"something": 10,
		"yo":        20,
		"blah":      20,
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value // 降序
		// return ss[i].Value < ss[j].Value  // 升序
	})

	for _, v := range ss {
		fmt.Printf("%s, %d\n", v.Key, v.Value)
	}
}
