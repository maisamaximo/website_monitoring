package main

import (
	"io/ioutil"
	"strings"
		"fmt"
		"os"
		"net/http"
		"time"
		"bufio"
		"io"
		"strconv"
)

const monitoring = 5
const delay = 5

func main() {

	intro()

	for {
		showMenu()

		command := getCommand()

		switch command {
		case 1: 
			startMonitoring()
		case 2: 
			showLogs()
		case 0: 
			fmt.Println("exit the application")
			os.Exit(0)
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
		}
	}
}

func intro() {
	name := "Maisa"
	version := 1.1

	fmt.Println("Hello Miss.", name)
	fmt.Println("Current version", version)
}

func getCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func showMenu() {
	fmt.Println("1. Start monitoring");
	fmt.Println("2. View logs");
	fmt.Println("0. Exit");
	fmt.Println("")
}

func startMonitoring() {
	fmt.Println("... monitoring")

	sites := readFile()

	for i := 0 ; i < monitoring ; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "não pôde ser carregado. Status Code:", response.StatusCode)
		registerLog(site, false)
	}
}

func readFile() []string {

	var sites []string
	
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	leitor := bufio.NewReader(file)

	for {
		linha, err :=  leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	date := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString(date + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}