/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import "fmt"

func main() {

	var (
		sum3  int
		sum5  int
		sum15 int
	)

	for i := 0; i < 1000; i++ {
		if i%3 == 0 && i%5 == 0 {
			sum15 += i

		} else if i%5 == 0 {
			sum5 += i
		} else if i%3 == 0 {
			sum3 += i
		}

	}
	fmt.Println("sum 3 and 5:", sum15, "sum 5:", sum5, "sum 3:", sum3, "total:", sum15+sum5+sum3)
}