package main

import "fmt"

// use pass by pointer
func addTen(x *int) {
	*x += 10
	fmt.Println("addTen func", *x)
}
func main() {
	var a int = 10
	fmt.Println("before", a)
	addTen(&a)
	fmt.Println("after", a)
}

// use pass by value
// func add(x int) {
// 	x += 10
// 	fmt.Println("add func", x)
// }
// func main() {
// 	var a int = 10
// 	fmt.Println("before", a)
// 	add(a)
// 	fmt.Println("after", a)
// }
