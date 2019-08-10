package store

import (
	"sort"
)

type data []int64

func (a data) Len() int           { return len(a) }
func (a data) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a data) Less(i, j int) bool { return a[i] < a[j] }

func sorting(as []int64) {
	sort.Sort(data(as))
}
