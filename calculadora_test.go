package main

import (
	"log"
	"testing"
)

// Aqui estão os testes para as funções da calculadora
// Usando com triple A (Arrange, Act, Assert) para organizar os testes e garantir que eles sejam claros e fáceis de entender
// Usando nomes descritivos para os testes, seguindo o padrão TestFuncion_Correct e TestFuncion_Incorrect
// imprimir mensagens de log para cada etapa do teste, usando o pacote log

func TestSomar_Correct(t *testing.T) {
	// Arrange
	a := 2.0
	b := 3.0
	esperado := 5.0
	// Act
	resultado := Somar(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Somar %f + %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de soma - Teste concluído com sucesso")
}

func TestSubtrair_Correct(t *testing.T) {
	// Arrange
	a := 5.0
	b := 2.0
	esperado := 3.0
	// Act
	resultado := Subtrair(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Subtrair %f - %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de subtração - Teste concluído com sucesso")
}

func TestMultiplicar_Correct(t *testing.T) {
	// Arrange
	a := 4.0
	b := 3.0
	esperado := 12.0
	// Act
	resultado := Multiplicar(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Multiplicar %f * %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de multiplicação - Teste concluído com sucesso")
}

func TestDividir_Correct(t *testing.T) {
	// Arrange
	a := 10.0
	b := 2.0
	esperado := 5.0
	// Act
	resultado := Dividir(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Dividir %f / %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de divisão - Teste concluído com sucesso")
}

func TestDividir_Incorrect(t *testing.T) {
	// Arrange
	a := 10.0
	b := 3.0
	esperado := 5.0
	// Act
	resultado := Dividir(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Dividir %f / %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de divisão incorreta - Teste concluído com sucesso")
}

func TestMultiplicar_Incorrect(t *testing.T) {
	// Arrange
	a := 4.0
	b := 3.0
	esperado := 10.0
	// Act
	resultado := Multiplicar(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Multiplicar %f * %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de multiplicação incorreta - Teste concluído com sucesso")
}

func TestSubtrair_Incorrect(t *testing.T) {
	// Arrange
	a := 5.0
	b := 2.0
	esperado := 4.0
	// Act
	resultado := Subtrair(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Subtrair %f - %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de subtração incorreta - Teste concluído com sucesso")
}

func TestSomar_Incorrect(t *testing.T) {
	// Arrange
	a := 2.0
	b := 3.0
	esperado := 6.0
	// Act
	resultado := Somar(a, b)
	// Assert
	if resultado != esperado {
		t.Errorf("Somar %f + %f) = %f; o valor esperado é %f", a, b, resultado, esperado)
	}
	log.Println("Teste de soma incorreta - Teste concluído com sucesso")
}
