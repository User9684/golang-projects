package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var beemoLogsUrl string
	fmt.Println("Logs URL?")
	fmt.Scanln(&beemoLogsUrl)

	// Get the beemo log
	resp, err := http.Get(beemoLogsUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bodyString := string(body)

	rawIds := strings.Split(bodyString, "Raw IDs:")
	beemoLogsUrlSplit := strings.Split(beemoLogsUrl, "/")
	BeemoLogsID := beemoLogsUrlSplit[len(beemoLogsUrlSplit)-1]

	// Create file
	file, err := os.Create(BeemoLogsID)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Write to file
	file.WriteString(rawIds[len(rawIds)-1])
}
