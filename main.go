package main

import "fmt"

func main() {
	var price float32
	fmt.Println(price)
	price = 300
	fmt.Println(price)
}

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
