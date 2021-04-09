package multiply

import "fmt"

// Multiply ...
// Method to return poduct of integers
func Multiply(a ...int) int {
	var prod int
	for i := 0; i < len(a); i = i + 1 {
		prod = prod * a[i]
	}
	fmt.Println("Total Product: ", prod)
	return prod
}
