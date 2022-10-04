package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	sakura := new(person)
	sakura.name = "千尋"
	sakura.age = 14

	sakura.name = "小千"
	fmt.Println(&sakura) //&{小千 14}
	fmt.Println("*******************")
	copySakura := person{
		name: "複製千尋",
		age:  14,
	}
	copySakura.name = "複製小千"
	fmt.Println(&copySakura)

}
