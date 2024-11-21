// a partir da digitação da altura e largura
// será calculado a área de um trangulo
package main

import "fmt"

func main() {
	// DECALRAÇÃO DAS VARIAVEIS
	var altura, base, area float32

	// ENTRADA DE DADOS

	fmt.Print("Valor da base do triângulo: ")
	if _, err := fmt.Scan(&base); err != nil {
		fmt.Println("Erro ao ler base:", err)
		return
	}

	fmt.Print("Valor da altura do triângulo: ")
	if _, err := fmt.Scan(&altura); err != nil {
		fmt.Println("Erro ao ler a altura: ", err)
		return
	}

	// CALCULO DA AREA

	area = (base * altura) / 2

	// SAIDA

	fmt.Printf("A área do triângulo é : %v\n", area)
}
