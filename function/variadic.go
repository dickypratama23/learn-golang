package main

import "fmt"

func main() {
	var avg = calculateAvg(2, 4, 5, 6, 7, 8, 5, 4, 3, 2, 1, 4, 5, 100)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

// variadic function
func calculateAvg(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))

	return avg
}
