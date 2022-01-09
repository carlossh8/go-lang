package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	Introducao()

	for {
		ApresentacaoMenu()

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
}

// Função responsável pela captura do nome do usuário
// e exibir a versão do programa

func Introducao() {

	fmt.Println("Digite seu nome: ")
	var nome1 string
	fmt.Scanln(&nome1)
	fmt.Println("-----------------------")

	fmt.Println("Olá," + nome1)
	fmt.Println("Você está executando a versão 1.1 do sistema")

}

// Menu do sistema

func ApresentacaoMenu() {

	fmt.Println("----------------------")

	fmt.Println("Iniciando Menu do sistema")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

// Responsável por ler opção digitada

func lerOpcao() int {
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

// Responsável pelo monitoramente

func IniciarMonitoramento() {
	fmt.Println("Iniciando Monitoramento")
	site := "https://github.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}
