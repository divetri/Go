package main

import "fmt"

func calculator() {
	var menu, n1, n2, value int32
	fmt.Println("Calculator")

	fmt.Println("Choose menu:")
	fmt.Println("1. Sum (+)")
	fmt.Println("2. Sub (-)")
	fmt.Println("3. Ply (*)")
	fmt.Println("4. Div (/)")

	fmt.Println("Choose menu:")
	fmt.Scanln(&menu)

	fmt.Println("Insert first number:")
	fmt.Scanln(&n1)

	fmt.Println("Insert second number:")
	fmt.Scanln(&n2)

	switch menu {
	case 1:
		value = n1 + n2
		fmt.Println(n1, "+", n2, "=", value)
	case 2:
		value = n1 - n2
		fmt.Println(n1, "-", n2, "=", value)
	case 3:
		value = n1 * n2
		fmt.Println(n1, "*", n2, "=", value)
	case 4:
		value = n1 / n2
		fmt.Println(n1, "/", n2, "=", value)
	}
}
