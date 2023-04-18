package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("hallo")
	os.Exit(1)
	fmt.Println("Welcome")
}
