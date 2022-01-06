package main

import "fmt"

func main() {

	fmt.Println("Digite seu nome: ")
	var nome1 string
	fmt.Scanln(&nome1)

	fmt.Println("Olá," + nome1)
	fmt.Println("Você está executando a versão 1.0 do sistema")

	fmt.Println("----------------------")

	fmt.Println("Iniciando Menu do sistema")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O comando escolhido foi", comando)
}
