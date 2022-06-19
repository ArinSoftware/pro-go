package main

import (
	"fmt"
	"strconv"
)

func main() {
	var price []rune = []rune("€48.95")
	var currency string = string(price[0])
	var amountString string = string(price[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

/* package main

import (
	"fmt"
	"sort"
)

func main() {
	products := map[string]float64{"kayak": 10.9, "hat": 2.1, "ceket": 0}

	keys := make([]string, 0, len(products))

	for key := range products {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("Key: ", key, "Value: ", products[key])
	}
} */

/* package main

import "fmt"

func main() {
	products := map[string]float64{"kayak": 10.9, "hat": 2.1}

	delete(products, "hat")

	value, ok := products["hat"]

	if ok {
		fmt.Println("Stored value", value)
	} else {
		fmt.Println("No stored value")
	}
}
*/
/* package main

import "fmt"

func main() {
	 	products := make(map[string]float64, 10)

	   	products["kazak"] = 25.2
	   	products["ceket"] = 44.3

	products := map[string]float64{"kazak": 25.3, "ceket": 44.3}

	fmt.Println("Map size:", len(products))
	fmt.Println("Price: ", products["kazak"])
	fmt.Println("Price: ", products["ceket"])
	fmt.Println("Price: ", products["xyz"])
} */

/* package main

import "fmt"

func main() {
	products := []string{"p1", "p2", "p3", "p4"}
	replacementProducts := []string{"r1", "r2"}

	copy(replacementProducts, products)

	fmt.Println(replacementProducts)
	fmt.Println(products)
} */

/* package main

import "fmt"

func main() {
	products := [4]string{"p1", "p2", "p3", "p4"}

	allNames := products[1:]
	someNames := make([]string, 2)

	copy(someNames, allNames)

	fmt.Println("All Names: ", allNames)
	fmt.Println("Some Names: ", someNames)
} */

/* package main

import "fmt"

func main() {
	names := [4]string{"Gurcan", "Gamze", "Arin", "Elis"}

	someNames := names[1:3]
	allNames := names[:]

	fmt.Println(someNames)
	fmt.Println(allNames)
} */

/* package main

import "fmt"

func main() {
	names := make([]string, 2, 4)

	names[0] = "Gurcan"
	names[1] = "Gamze"

	appendedNames := append(names, "Arin", "Elis")

	names[0] = "Gugu"

	fmt.Println(names)
	fmt.Println(appendedNames)
} */

/* package main

import "fmt"

func main() {
	names := make([]string, 3, 6)

	names[0] = "Gamze"
	names[1] = "Arin"
	names[2] = "Elis"

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))
} */

/* package main

import "fmt"

func main() {
	names := []string{"Gamze", "Arin"}

	apNames := append(names, "Elis")

	names[0] = "Gurcan"

	fmt.Println(names)
	fmt.Println(apNames)
} */

/* package main

import "fmt"

func main() {
	names := []string{"Gamze", "Arin"}
	names = append(names, "Gurcan", "Elis")

	fmt.Println(names)
} */

/* package main

import "fmt"

func main() {
	names := []string{"Gamze", "Arin", "Elis"}

	fmt.Println(names)
} */

/* package main

import "fmt"

func main() {
	names := make([]string, 3)

	names[0] = "Gamze"
	names[1] = "Arin"
	names[2] = "Elis"

	fmt.Println(names)
} */

/* package main

import "fmt"

func main() {
	names := [4]string{"Gamze", "Gurcan", "Arin", "Elis"}

	for index, name := range names {
		fmt.Println("Index: ", index, "Value: ", name)
	}
} */

/* package main

import "fmt"

func main() {
	names := [2]string{"Gamze", "Gurcan"}
	names2 := [2]string{"Gurcan", "Gamze"}

	same := names == names2

	fmt.Println("Same::", same)
} */

/* package main

import "fmt"

func main() {
	names := [2]string{"Gamze", "Gurcan"}

	names2 := &names

	names[0] = "Arin"

	fmt.Println(names)
	fmt.Println(*names2)
} */

/* package main

import "fmt"

func main() {
	names := [2]string{"Gamze", "Gurcan"}

	names2 := names

	names[0] = "Arin"

	fmt.Println(names)
	fmt.Println(names2)
} */

/* package main

import "fmt"

func main() {
	names := [...]string{"Gamze", "Arin", "Elis"}

	fmt.Println(names)
}
*/
/* package main

import "fmt"

func main() {
	names := [4]string{"Gurcan", "Gamze", "Arin", "Elis"}

	fmt.Println(names)
} */

/* package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Gurcan"
	names[1] = "Gamze"
	names[2] = "Arin"
	names[3] = "Elis"

	fmt.Println(names)
} */
