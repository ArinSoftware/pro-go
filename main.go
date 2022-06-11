package main

import "fmt"

func main() {
	names := [4]string{"Gurcan", "Gamze", "Arin", "Elis"}

	someNames := names[1:3]
	allNames := names[:]

	fmt.Println(someNames)
	fmt.Println(allNames)
}

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
