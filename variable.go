package main

import "fmt"

func main() {
	var lastName string = "dicky"
	var fisrtName string = "Pratama"
	const age int = 30

	fmt.Print("hallo ", fisrtName, lastName, "!\n")

	fisrtName = "Pratama2" // assign new value
	lastName = "Dicky2"    // assign new value

	fmt.Print("hallo ", fisrtName, lastName, "!\n")

	//age = 14 // can't assign new value
}
