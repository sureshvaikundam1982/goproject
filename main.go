package main

import "fmt"

type favNum int
type favWord string
type favBool bool

func main() {
	a := 2
	b := 3
	c := a + b
	fmt.Println("Hello, World!", c)

	var num favNum = 24
	var word favWord = "take care health"
	var boo favBool = true

	//var keyword
	var empty string
	var order int
	var drink string = "milkshake"

	//shorthand method
	quantity := 1
	food := "burrito"

	// const
	const pi = 3.14

	fmt.Println("%T/n", empty)
	fmt.Println("%T/n", order)
	fmt.Println("%v/n", drink)
	fmt.Println("%v/n", quantity)
	fmt.Println("%T/n", quantity)
	fmt.Println("%v/n", food)
	fmt.Println("%T/n", food)
	fmt.Println("%v/n", pi)
	fmt.Println("%T/n", pi)

	fmt.Println("Hello", num)
	fmt.Println("Hello", word)
	fmt.Println("Hello", boo)

}
