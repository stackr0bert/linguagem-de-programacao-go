package main

import "fmt"

// Função para calcular a média das notas
func calcularMedia(total float32) float32 {
	return total / 4
}

// Função para determinar a situação do aluno com base na média
func determinarSituacao(media float32) string {
	switch {
	case media >= 7:
		return "Aprovado"
	case media >= 5:
		return "Recuperação"
	default:
		return "Reprovado"
	}
}

func main() {
	var nota, total float32

	// Loop para coletar as notas dos 4 bimestres
	for bim := 1; bim <= 4; bim++ {
		fmt.Printf("Qual a nota do %dº Bimestre: \n", bim)
		fmt.Scan(&nota)
		total += nota
	}

	// Calcula a média das notas
	media := calcularMedia(total)
	fmt.Printf("O resultado é: %.2f\n", media)
	
	// Determina e imprime a situação do aluno
	fmt.Printf("Situação: %s\n", determinarSituacao(media))
}
