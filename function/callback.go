package main

import "fmt"

func main() {
	var result = filter([]string{"ini", "data"}, func(each string) bool {
		return true
	})
	fmt.Println(result)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
