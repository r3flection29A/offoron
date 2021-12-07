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

const delay = 10
const contador = 5

func main() {

	for {
		log("site-falso", false)
		banner()
		opt := leInput()

		switch opt {
		case 1:
			iniciaMonitoramento()
		case 2:
			outputLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("INVALID.")
		}
	}

}

func banner() {
	fmt.Println("\t+++++++++++++++++++++++++++")
	fmt.Println("\t1 - Iniciar monitoramento")
	fmt.Println("\t2 - Exibir logs")
	fmt.Println("\t0 - Sair do programa")
	fmt.Println("\t+++++++++++++++++++++++++++")
}

func leInput() int {
	var option int
	fmt.Printf("\nDigite uma opção: ")
	fmt.Scanf("%d", &option)
	return option
}

func iniciaMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := leSites()
	for i := 0; i < contador; i++ {
		for i, site := range sites {
			fmt.Println("\nTestando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	r, err := http.Get(site)
	if err != nil {
		fmt.Println("Houve um erro inesperado :(", err)
	}

	if r.StatusCode == 200 {
		fmt.Println("O site", site, "está no ar!")
		log(site, true)
	} else {
		fmt.Println("O site", site, "respondeu fora do status code 200. Status code atual: ", r.StatusCode)
		log(site, false)
	}
}

func leSites() []string {
	var sites []string
	arq, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Houve um erro inesperado :(", err)
	}
	leitor := bufio.NewReader(arq)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err != io.EOF {
			break
		}
	}

	arq.Close()
	return sites
}

func log(site string, statusCode bool) {
	arq, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Houve um erro inesperado :(", err)
	}

	arq.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "online: " + strconv.FormatBool(statusCode) + "\n")

	arq.Close()
}

func outputLog() {
	fmt.Println("Exibindo logs...")
	arq, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Houve um erro inesperado :(", err)
	}

	fmt.Println(string(arq))
}
