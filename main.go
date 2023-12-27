package main

import "fmt"

func main() {
	fmt.Println("ok")

	// names := []string{"divet", "bambang", "paijo"}
	// for index, value := range names {
	// 	fmt.Printf("absen %v namanya adalah %v \n", index+1, value)
	// }

	names := []string{"divet", "bambang", "paijo"}
	for _, value := range names {
		fmt.Printf("namanya adalah %v \n", value)
	}

}
