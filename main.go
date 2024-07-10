package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	showIntro()
	for {
		showMenu()
		command := readCommand()
		chooseOption(command)
	}
}

func showIntro() {
	version := 1.0
	fmt.Printf("Introduction to Golang version %f\n", version)
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
	chosenOption := showAvailableWebsites()
	fmt.Println("Chosen website:", chosenOption)

	mockUrl := fmt.Sprintf("https://httpbin.org/status/%d", randomStatusCode())
	response, _ := http.Get(mockUrl)

	if response.StatusCode == 200 {
		fmt.Println("Loaded with success")
	} else {
		fmt.Println("Error to load website. Stauts code:", response.StatusCode)
	}
}

func showAvailableWebsites() string {
	sites := []string{
		"www.testing.com",
		"www.testing-again.com",
		"www.last-testing.com",
	}

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
