package examples

import "fmt"

func types() {
	fmt.Print("Enter the faringate number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Print("По Цельсию: ", output)
}
