package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"My", "Skill"}
	printMessage("halo", names)
	printMessage("Heso", names)
}

// new function
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
