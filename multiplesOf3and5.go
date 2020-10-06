package main

import "fmt"

func main() {

	var (
		sum3  int
		sum5  int
		sum15 int
	)

	for i := 0; i <= 1000; i++ {
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
