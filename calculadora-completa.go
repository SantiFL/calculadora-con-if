package main

import (
	"fmt"
)

func main() {
	inicio()
	var numero1 float32 = 10
	var numero2 float32 = 5
	fmt.Print("La suma es igual a = ")
	fmt.Println(calculadora(numero1, numero2, "+"))

	fmt.Print("La resta es igual a = ")
	fmt.Println(calculadora(numero1, numero2, "-"))

	fmt.Print("La divicion es igual a = ")
	fmt.Println(calculadora(numero1, numero2, "/"))

	fmt.Print("La multiplicacion es igual a = ")
	fmt.Println(calculadora(numero1, numero2, "*"))
}
func inicio() {
	fmt.Println("calculadora basica con funciones")
}

func calculadora(n1 float32, n2 float32, op string) float32 {
	var resultado float32
	if op == "+" {
		resultado = n1 + n2
	}

	if op == "-" {
		resultado = n1 - n2
	}

	if op == "*" {
		resultado = n1 * n2
	}
	if op == "/" {
		resultado = n1 / n2
	}

	return resultado
}
