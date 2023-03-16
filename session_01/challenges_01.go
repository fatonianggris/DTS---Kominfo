package main

import (
	"fmt"
)

func main() {
	var i int = 21
	var percent string = "%"
	var j bool = true
	var uni uint = '\u042F'
	var charR uint = 'Я'
	var base8 int = 25
	var base10 int = 21
	var base16 int = 0xF
	var base16f13 int = 3859
	var k float64 = 123.456

	fmt.Println("Mohammad Fatoni Anggris (1955617840-17)")
	fmt.Println("-----------------------------------------")
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%v \n", percent)
	fmt.Printf("%t \n", j)
	fmt.Printf("%c \n", uni)
	fmt.Printf("%d \n", base10)
	fmt.Printf("%d \n", base8)
	fmt.Printf("%x \n", base16)
	fmt.Printf("%X \n", base16)
	fmt.Printf("%X \n", base16f13)
	fmt.Printf("%U \n", charR)
	fmt.Printf("%f \n", k)
	fmt.Printf("%e \n", k)
}
