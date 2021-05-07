package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"econ1100": {"math 0220", "econ100", "econ110"},
	"econ1110": {"math 0220", "econ1100", "econ100", "econ110"},
	"econ1150": {"econ1110"},
	"econ1200": {"econ1110"},

	"math230":  {"math 0220"},
	"math240":  {"math 0230"},
	"math413":  {"math240"},
	"math420":  {"math240", "math413"},
	"math1270": {"math240"},

	"stats1151": {"math240"},
	"stats1152": {"stats1151"},
}

func main() {
	for i, course := range courseSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func courseSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
