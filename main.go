package main

import "fmt"

func main() {
	first := 100
	var second *int

	fmt.Println(*second)

	fmt.Println(first)
	fmt.Println(second)

	second = &first

	fmt.Println(first)
	fmt.Println(second)
}

/* package main

import "fmt"

func main() {
	first := 100
	second := &first

	first++
	*second++

	fmt.Println(first)
	fmt.Println(*second)

	var myNewPointer *int = second
	*myNewPointer++

	fmt.Println(first)
	fmt.Println(*second)
	fmt.Println(*myNewPointer)

} */

/* package main

import "fmt"

func main() {

	price, tax, inStock, _ := 275.00, 27.5, true, true

	var _ = "Alice"

	fmt.Println("Total: ", price+tax)
	fmt.Println("InStock: ", inStock)

	price2, tax := 250.00, 27.5
	fmt.Println("Total: ", price2+tax)
} */

/* package main

import "fmt"

func main() {

	 	var price, tax = 275.00, 27.50

	   	fmt.Println(price + tax)

	var price, tax float64
	price = 275.00
	tax = 27.50
	fmt.Println(price + tax)
} */

/* package main

import "fmt"

func main() {
	var price float32
	fmt.Println(price)
	price = 300
	fmt.Println(price)
} */

/* package main

import "fmt"

func main() {
	var price = 275.00
	var price2 = price

	fmt.Println(price)
	fmt.Println(price2)
} */

/* package main

import "fmt"

func main() {
	var price float32 = 275.00
	var tax float32 = 27.50

	fmt.Println(price + tax)

	price = 300

	fmt.Println(price + tax)
} */

/* package main

import "fmt"

func main() {
	const price, tax float32 = 275.00, 27.50
	const quantity, inStock = 2, true

	fmt.Println("Total: ", 2*quantity*(price+tax))
	fmt.Println("In Stock", inStock)
} */

//////////////////

/* package main

import "fmt"

func main() {
	fmt.Println("Hello, Go")
	fmt.Println(20 + 20)
	fmt.Println(20 + 30)
} */

//////////////////////////

/* package main

import (
	"fmt"
	"math/rand"
)

func main1() {
	fmt.Println("Value: ", rand.Int())
} */
