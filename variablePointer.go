package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	*numberB = 7 // assign new value

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA address :", &numberA)

	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB address :", &numberB)
}
