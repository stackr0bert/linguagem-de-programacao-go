package main

import "fmt"

func main() {
	// DECLARAÇÃO DE VARIÁVEIS
	var numero1, numero2, resultado float32
	var operacao string = " "
	// ENTRADA DE DADOS
	fmt.Print("Digite o primeiro número: ")
	if _, err := fmt.Scanln(&numero1); err != nil {
		fmt.Print("Erro ao digitar número: ", err)
		return
	}

	fmt.Print("Digite o segundo número: ")
	if _, err := fmt.Scanln(&numero2); err != nil {
		fmt.Print("Erro ao digitar número: ", err)
		return
	}
	// ENTRADA DE DADOS (CARACTERE)
	fmt.Print("Escolha o tipo de operação que será feita (+, -, *, /): ")
	if _, err := fmt.Scan(&operacao); err != nil ||
		(operacao != "+" && operacao != "-" && operacao != "*" && operacao != "/") {
		fmt.Println("Erro! Caractere inválido ou erro ao digitar operação: ", err)
		return
	}

	// CONDIÇÃO E CÁLCULO ESCOLHIDO
	if operacao == "+" {
		resultado = numero1 + numero2
	} else if operacao == "-" {
		resultado = numero1 - numero2
	} else if operacao == "*" {
		resultado = numero1 * numero2
	} else {
		resultado = numero1 / numero2
	}

	// SAÍDA DE DADOS
	fmt.Printf("Resultado da operação: %v\n", resultado)
}
