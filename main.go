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

const (
	monitoramentos = 2
	delay          = 5
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()
		escolherComando(comando)
	}
}

func escolherComando(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Evandro"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := LerSitesDoArquivo("sites.txt")
	for index := 0; index < monitoramentos; index++ {
		for _, site := range sites {
			TestaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func TestaSite(site string) bool {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao abrir site:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		RegistraLog(site, true)
		return true
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		RegistraLog(site, false)
		return false
	}
}

func LerSitesDoArquivo(fileToOpen string) []string {
	var sites []string
	arquivo, err := os.Open(fileToOpen)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return sites
	}
	defer arquivo.Close()

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Erro ao ler linha:", err)
			continue
		}
		linha = strings.TrimSpace(linha)
		fmt.Println("Encontrei a linha:", linha)
		sites = append(sites, linha)
	}
	return sites
}

func RegistraLog(site string, status bool) bool {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return false
	}
	defer arquivo.Close()
	arquivo.WriteString(time.Now().Format("2006-01-02 15:04:05") + " | " + site + " | online: " + strconv.FormatBool(status) + "\n")
	return true
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	fmt.Println(string(arquivo))
}
