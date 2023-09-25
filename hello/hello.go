package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func showVersion() {
	nome := "Guilherme"
	versao := 1.21
	fmt.Println("Nome:", nome)
	fmt.Println("Versao:", versao)
}

func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		file.Close()
		fmt.Println("Error", err)
	}
	scanner := bufio.NewReader(file)
	for {
		linha, err := scanner.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func getCommand() int {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Verificar os logs")
	fmt.Println("0- Sair")
	var comando int
	fmt.Scan(&comando)
	fmt.Print("\033[H\033[2J") // clear console
	return comando
}

func setLog(site string, status int) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if status != 200 {
		file.WriteString(time.Now().Format("2006-01-02T15:04:05") + " - " + site + " com erro. Error code: " + strconv.Itoa(status) + "\n")
	} else {
		file.WriteString(time.Now().Format("2006-01-02T15:04:05") + " - " + site + " funcional." + "\n")
	}
	file.Close()
}

func getRequest() {
	sites := readFile()
	for _, site := range sites {
		request, err := http.Get(site)
		if err != nil {
			fmt.Println("Error", err)
		}
		setLog(site, request.StatusCode)
		time.Sleep(2 * time.Second)
	}
}

func showLogs() {
	file, err := os.Open("log.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	scanner := bufio.NewReader(file)

	for {
		log, err := scanner.ReadString('\n')
		log = strings.TrimSpace(log)
		fmt.Println(log)
		if err == io.EOF {
			break
		}
	}

	file.Close()
}

func main() {
	showVersion()
	for {
		comando := getCommand()
		switch comando {
		case 1:
			getRequest()
		case 2:
			showLogs()
		default:
			fmt.Println("Saindo")
			os.Exit(0)
		}
	}
}
