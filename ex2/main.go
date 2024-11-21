package main

import (
	"fmt"
	"math"
)

func main() {
	// declaração de variável
	var angulo float64
	// entrada de dado

	fmt.Print("Digite o ângulo desejado: ")
	if _, err := fmt.Scan(&angulo); err != nil {
		fmt.Println("Erro ao ler ângulo:", err)
		return
	}
	// calculo e impressão do cosseno
	fmt.Println("Cosseno de", angulo, "é:", math.Cos(angulo))

	// calculo e impressao do seno
	fmt.Println("Seno de", angulo, "é:", math.Sin(angulo))
}
