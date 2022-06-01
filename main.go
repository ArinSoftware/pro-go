package main

import "fmt"

func main() {
	first := 100
	second := &first
	third := &first

	alfa := 100
	beta := &alfa

	fmt.Println(second == third)
	fmt.Println(second == beta)
	fmt.Println()
	fmt.Println(*second == *third)
	fmt.Println(*second == *beta)

}

/* package main

import (
	"fmt"
	"math"
)

func main() {
	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))

	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)
} */

/* package main

import "fmt"

func main() {
	price, tax := 275.00, 27.40

	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)
} */
