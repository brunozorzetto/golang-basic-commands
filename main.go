package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const VERSION = 1.0

func main() {
	showIntro()
	for {
		showMenu()
		command := readCommand()
		chooseOption(command)
	}
}

func showIntro() {
	fmt.Printf("Introduction to Golang version %f\n", VERSION)
}

func showMenu() {
	fmt.Println(" --- Menu ---")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Have fun")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func chooseOption(option int) {
	switch option {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Showing logs ...")
		printLogs()
	case 3:
		fmt.Println("Have fun with Go ...")
	case 0:
		fmt.Println("Exiting ...")
		os.Exit(0)
	default:
		fmt.Println("Unknown command!")
		os.Exit(-1)
	}
}

func randomStatusCode() int {
	possibleStatusCodes := []int{
		200,
		404,
		500,
	}

	return possibleStatusCodes[rand.Intn(len(possibleStatusCodes))]
}

func startMonitoring() {
	fmt.Println("Monitoring ...")
	fmt.Println("Read from file (Y/N)")
	var option string
	fmt.Scan(&option)

	var chosenWebsite string
	if strings.ToLower(option) == "y" {
		chosenWebsite = showSitesFromFile()
	} else {
		chosenWebsite = showAvailableWebsites()
		fmt.Println("Chosen website:", chosenWebsite)
	}

	fmt.Println(chosenWebsite)
	response, error := http.Get(chosenWebsite)
	if error != nil {
		fmt.Println("Error to request.")
	}

	if response.StatusCode == 200 {
		fmt.Println("Loaded with success")
		log(chosenWebsite, true)
	} else {
		fmt.Println("Error to load website. Stauts code:", response.StatusCode)
		log(chosenWebsite, false)
	}
}

func showAvailableWebsites() string {
	sites := []string{
		fmt.Sprintf("https://httpbin.org/status/%d", randomStatusCode()),
		fmt.Sprintf("https://httpbin.org/status/%d", randomStatusCode()),
		fmt.Sprintf("https://httpbin.org/status/%d", randomStatusCode()),
	}

	return chooseAvailableWebSites(sites)
}

func showSitesFromFile() string {
	var sites []string

	file, error := os.Open("sites.txt")
	if error != nil {
		fmt.Println("File not found.")
	}

	reader := bufio.NewReader(file)
	for {
		line, error := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(line))

		if error == io.EOF {
			break
		}
	}

	file.Close()

	return chooseAvailableWebSites(sites)
}

func chooseAvailableWebSites(sites []string) string {
	fmt.Println("Available websites:")
	for i, v := range sites {
		fmt.Println("Option[", i, "] - ", v)
	}

	var option int
	fmt.Scan(&option)

	if option >= len(sites) {
		println("Error to load website")
		os.Exit(0)
	}

	return sites[option]
}

func log(webSite string, success bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0775)

	if err != nil {
		fmt.Println("Error to open file")
	}

	file.WriteString(time.Now().UTC().Format("2006-01-02 15:04:05 - ") + "Website" + webSite + " - online: " + strconv.FormatBool(success) + "\n")

	file.Close()
}

func printLogs() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error to show logs =(")
	}

	fmt.Println(string(file))
}
