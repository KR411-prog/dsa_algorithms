package main

import (
	"fmt"
	"sort"
)

func sort_slice() {
	s := []string{"c", "a", "b"}
	sort.Strings(s)
	fmt.Println(s)

	family := []struct {
		Name string
		Age int
	} {
		{
			Name: "Radha",
			Age: 38,
		},
		{
			Name: "Kairav",
			Age: 13,
		},
		{
			Name: "Veera",
			Age: 36,
		},
	}
	sort.SliceStable(family, func(i,j int) bool{
		return family[i].Name > family[j].Name
	})

	fmt.Println(family)


}
