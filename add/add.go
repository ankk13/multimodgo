package add

import (
	"fmt"
)

// Add ...
// Method to return sum of ints
func Add(a ...int) int {
	var sum int
	for i := 0; i < len(a); i = i + 1 {
		sum = sum + a[i]
	}
	fmt.Println("Total Sum: ", sum)
	return sum
}
