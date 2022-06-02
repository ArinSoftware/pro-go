package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "100"

	if int1, int1err := strconv.ParseInt(val1, 2, 8); int1err != nil {
		fmt.Println("Can not parse:: ", val1)
	} else {
		smallInt := int8(int1)
		fmt.Println("smallInt:: ", smallInt)
	}
}

/* package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "not false"

	if bool1, b1err := strconv.ParseBool(val1); b1err != nil {
		fmt.Println("Can not parse: ", val1)
	} else {
		fmt.Println("Bool1: ", bool1)
	}

} */

/* package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "false"

	bool1, b1err := strconv.ParseBool(val1)

	if b1err != nil {
		fmt.Println("can not parse: ", val1)
	} else {
		fmt.Println("Bool1: ", bool1)
	}
} */

/* package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "true"
	val2 := "false"
	val3 := "not true"

	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)

	fmt.Println("Bool1", bool1, b1err)
	fmt.Println("Bool2", bool2, b2err)
	fmt.Println("Bool3", bool3, b3err)

} */

/* package main

import (
	"fmt"
	"math"
)

func main() {
	value := 2.7

	fmt.Println(math.Round(value))
	fmt.Println(math.Ceil(value))
	fmt.Println(math.Floor(value))
	fmt.Println(math.RoundToEven(value))
} */

/* package main

import "fmt"

func main() {
	num1 := 275.50
	num2 := 25

	fmt.Println(int(num1) + num2)

	num3 := 275.50
	num4 := 25

	fmt.Println(num3 + float64(num4))

} */

/* package main

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

} */

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
