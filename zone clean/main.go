package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var debug bool = false

func main() {
	r, _ := regexp.Compile(`	`)
	fmt.Println("Zone-file: ")

	var fileName string
	var tld string
	fmt.Scanln(&fileName)

	tld = fileName
	newFileName := fmt.Sprintf("%s-cleaned.txt", fileName)
	fileName += ".txt"

	fmt.Println("Reading file...")
	zoneBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	zone := string(zoneBytes)

	fmt.Println("Splitting file...")
	zoneSplit := strings.Split(zone, "\n")

	fmt.Println("Creating file...")
	file, err := os.Create(newFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Println("Cleaning to just domains...")
	buffer := bufio.NewWriter(file)

	var lastDomain string

	for i, line := range zoneSplit {
		if !r.MatchString(line) {
			continue
		}

		domain := strings.SplitN(line, "	", 2)[0]
		domainNoRoot := domain[:len(domain)-1]
		domainNoTLD := domainNoRoot[0:(strings.LastIndex(domainNoRoot, tld))]

		if (lastDomain == domainNoRoot) || (domainNoRoot == tld) || (strings.Contains(domainNoTLD[:len(domainNoTLD)-1], ".")) {
			continue
		}

		lastDomain = domainNoRoot

		buffer.WriteString(fmt.Sprintf("%s\n", domainNoRoot))

		if i%10 == 10 {
			buffer.Flush()
		}
	}

	buffer.Flush()

	fmt.Println(fmt.Sprintf("Ported to %s", newFileName))
}
