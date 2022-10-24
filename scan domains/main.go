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
	fmt.Println("Domain list file: ")
	var fileName string
	fmt.Scanln(&fileName)
	newFileName := fmt.Sprintf("%s-potentially-malicious.txt", fileName)
	fileName += ".txt"

	var regex string
	fmt.Printf("Regex: ")
	fmt.Scanln(&regex)
	r, _ := regexp.Compile(regex)

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
	fmt.Println("Searching for potentially malicious domains...")
	buffer := bufio.NewWriter(file)

	for i, line := range zoneSplit {
		isMatch := r.MatchString(line)
		if isMatch {
			if debug {
				fmt.Println(fmt.Sprintf("%s is potentially malicious!", line))
			}
			_, err = buffer.WriteString(fmt.Sprintf("\n%s", line))
			if err != nil {
				fmt.Println(err)
			}
		}
		if i%10 == 10 {
			buffer.Flush()
		}
	}

	buffer.Flush()

	fmt.Println(fmt.Sprintf("Ported to %s", newFileName))
}
