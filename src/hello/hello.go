package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	Introducao()

	// Menu do sistema

	fmt.Println("----------------------")

	fmt.Println("Iniciando Menu do sistema")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	comando := lerOpcao()
	// Escolha do usuário

	switch comando {
	case 1:
		IniciarMonitoramento()
	case 2:
		fmt.Println("Iniciando Logs")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não reconheço esse comando...")
	}
}

func Introducao() {

	fmt.Println("Digite seu nome: ")
	var nome1 string
	fmt.Scanln(&nome1)

	fmt.Println("Olá," + nome1)
	fmt.Println("Você está executando a versão 1.1 do sistema")
}

func lerOpcao() int {
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func IniciarMonitoramento() {
	fmt.Println("Iniciando Monitoramento")
	site := "https://www.google.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
