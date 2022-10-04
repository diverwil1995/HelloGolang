package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	account := []struct {
// 		id   string
// 		name string
// 		age  int
// 	}{
// 		{"a001", "Wilson", 26},
// 		{"a002", "Tony", 44},
// 		{"a003", "Eve", 20},
// 		{"a004", "Peggy", 21},
// 	}

// 	fmt.Println(account)
// }

type account struct {
	Id   string
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []account

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	guest := []account{
		{"a001", "Wilson", 26},
		{"a002", "Tony", 42},
		{"a003", "Peggy", 23},
	}
	sort.Sort(ByAge(guest))
	fmt.Println(guest)
}
