package main

import (
	"fmt"
	"sort"
)

func sort_map() {
	m := map[string]int{
		"orange": 5, "apple": 7, "mango": 3, "strawberry": 9,
	}

	keys := make([]string, 0, len(m))

	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, m[key])
	}

}

func sort_map_value() {
	m := map[string]int{
		"orange": 5, "apple": 7, "mango": 3, "strawberry": 9,
	}
	keys := make([]string,0,len(m))
	for key := range m {
		keys = append(keys,key)
	}
	fmt.Println(keys)

	sort.SliceStable(keys, func(i,j int) bool {
		return m[keys[i]] < m[keys[j]]
	})
	for k,v := range keys{
		fmt.Println(k,v)
	}
}