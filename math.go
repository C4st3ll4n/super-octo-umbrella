package main

import "fmt"

func main() {
	fmt.Println(Soma(1, 2))
	fmt.Println(Subtrai(1, 2))
	fmt.Println(Multiplica(1, 2))
}

func Soma(a int, b int) int {
	return a + b
}

func Subtrai(a int, b int) int {
	return a - b
}

func Multiplica(a int, b int) int {
	return a * b
}
