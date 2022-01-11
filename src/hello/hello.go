package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 2
const delay = 5

func main() {

	Introducao()

	for {
		ApresentacaoMenu()

		comando := lerOpcao()
		// Escolha do usuário

		switch comando {
		case 1:
			iniciarMonitoramento()
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
	fmt.Println("Você está executando a versão 1.3 do sistema")

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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
