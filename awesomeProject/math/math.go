package math

import "fmt"

func Sum(a int, b int) int {
	sums := a + b
	fmt.Printf("a soma de %d x %d é: ", a, b)
	return sums
}

func Multiplication(a int, b int) int {
	multi := a * b
	fmt.Printf("a multiplicação de %d x %d é : ", a, b)
	return multi
}
