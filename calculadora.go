package main

import (
	"fmt"
)

// Nesse desafio do curso de GO na DIO.me
// Vou criar uma calculadora simples com as operações de soma, subtração, multiplicação e divisão.
// Onde o usuário pode escolher a operação e os números para realizar o cálculo.
// Fazer interação com o usuário para escolher a operação e os números
// exemplo: ele coloca 1+1 e o programa retorna 2
// ou ele coloca 2*3 e o programa retorna 6
// ou ele coloca 10/2 e o programa retorna 5
// ou ele coloca 5-3 e o programa retorna 2

// Função para somar dois números
func Somar(a, b float64) float64 {
	return a + b
}

// Função para subtrair dois números
func Subtrair(a, b float64) float64 {
	return a - b
}

// Função para multiplicar dois números
func Multiplicar(a, b float64) float64 {
	return a * b
}

// Função para dividir dois números
func Dividir(a, b float64) float64 {
	if b == 0 {
		return 0 // Evitar divisão por zero
	}
	return a / b
}

func main() {
	var num1, num2 float64 // Variáveis para armazenar os números inseridos pelo usuário
	var operador string    // Variável para armazenar o operador inserido pelo usuário
	fmt.Print("Digite a operação (ex: 1 + 1, 2 * 3, 10 / 2, 5 - 3): ")
	fmt.Scanln(&num1, &operador, &num2) // Lendo a entrada do usuário e armazenando os valores nas variáveis correspondentes
	var resultado float64               // Variável para armazenar o resultado do cálculo
	switch operador {                   // Verificando qual operador foi inserido pelo usuário e realizando a operação correspondente
	case "+":
		resultado = Somar(num1, num2)
	case "-":
		resultado = Subtrair(num1, num2)
	case "*":
		resultado = Multiplicar(num1, num2)
	case "/":
		resultado = Dividir(num1, num2)
	default:
		fmt.Println("Operação inválida")
		return
	}
	fmt.Printf("Resultado: %.2f\n", resultado) // Imprimindo o resultado do cálculo
	main()                                     // Chamando a função main novamente para permitir que o usuário realize novos cálculos
}
