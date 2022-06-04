package main

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
}

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
