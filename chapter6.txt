package main

import "fmt"

func main() {

	counter := 0

target:
	fmt.Println("Counter is:", counter)
	counter++
	if counter < 5 {
		goto target
	}
}

/* package main

import "fmt"

func main() {
	product := "Kayak"

	for i, c := range product {
		switch c {
		case 'x':
			fmt.Println("Uppercase character")
			fallthrough
		case 't':
			fmt.Println("x at position", i)
		case 'u':
			fmt.Println("y at index: ", i)
		default:
			fmt.Println("No cased character")
		}

	}
} */

/* package main

import "fmt"

func main() {

	product := "kayak"

	for i, c := range product {
		switch c {
		case 'k':
			fmt.Println("k at position", i)
		case 'y':
			fmt.Println("y at position", i)
		}
	}
} */

/* package main

import "fmt"

func main() {
	product := "kayak"

	for i, c := range product {
		fmt.Println("Index:", i, "Caharacter:", string(c))
	}

	for i := range product {
		fmt.Println("Index:", i)
	}

	for _, c := range product {
		fmt.Println("Caharacter:", string(c))
	}
} */

/* package main

import "fmt"

func main() {
	for counter := 0; counter <= 3; counter++ {
		if counter == 2 {
			continue
		}
		fmt.Println("Counter:", counter)
	}
} */

/* package main

import (
	"fmt"
)

func main() {
	for counter := 0; true; counter++ {
		fmt.Println("Counter:", counter)
		if counter > 3 {
			break
		}
	}
} */

/* package main

import "fmt"

func main() {
	for counter := 0; counter <= 3; counter++ {
		fmt.Println("Counter: ", counter)
	}
} */

/* package main

import "fmt"

func main() {
	counter := 0

	for counter <= 3 {
		fmt.Println("Counter: ", counter)
		counter++
	}
} */

/* package main

import "fmt"

func main() {
	counter := 0

	for {
		fmt.Println("Counter: ", counter)
		counter++
		if counter > 3 {
			break
		}
	}
} */

/* package main

import (
	"fmt"
	"strconv"
)

func main() {
	priceString := "100"

	if kayakPrice, err := strconv.Atoi(priceString); err != nil {
		fmt.Println("Cannot Parse", priceString)
	} else {

		fmt.Println("Parsed value", kayakPrice)
	}
} */

/* package main

import "fmt"

func main() {
	price := 250

	if price > 500 {
		scopedVar := 500
		fmt.Println("Price is higher than ", scopedVar)
	} else if price <= 200 {
		scopedVar := 200
		fmt.Println("Price is lower than or equal ", scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Match", scopedVar)
	}
} */

/* package main

import "fmt"

func main() {
	price := 250

	if price > 500 {
		fmt.Println("Price is higher than 500")
	} else if price <= 200 {
		fmt.Println("Price is lower than or equal 200")
	} else {
		fmt.Println("Price is did not match")
	}

} */
