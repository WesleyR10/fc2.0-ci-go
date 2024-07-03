package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func subtracao(a, b int) int {
	return a - b
}

func multiplicacao(a, b int) int {
	return a * b
}

func divisao(a, b int) int {
	return a / b
}

func main() {
	fmt.Println("Soma: ", soma(5, 3))
	fmt.Println("Subtração: ", subtracao(5, 3))
	fmt.Println("Multiplicação: ", multiplicacao(5, 3))
	fmt.Println("Divisão: ", divisao(5, 3))
}