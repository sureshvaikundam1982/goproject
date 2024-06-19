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

	fmt.Printf("%T/n", empty)
	fmt.Printf("%T/n", order)
	fmt.Printf("%v/n", drink)
	fmt.Printf("%v/n", quantity)
	fmt.Printf("%T/n", quantity)
	fmt.Printf("%v/n", food)
	fmt.Printf("%T/n", food)
	fmt.Printf("%v/n", pi)
	fmt.Printf("%T/n", pi)

	fmt.Println("Hello", num)
	fmt.Println("Hello", word)
	fmt.Println("Hello", boo)

}
