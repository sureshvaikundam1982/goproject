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

	fmt.Println("Hello", num)
	fmt.Println("Hello", word)
	fmt.Println("Hello", boo)

}
