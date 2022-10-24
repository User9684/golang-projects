package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var maxPerFile = 25000

func saveFile(chunk int, ids string) {
	file, err := os.Create(fmt.Sprintf("chunks/Chunk #%d", chunk))
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(ids)
}

func main() {
	var fileNames []string
	var done = false
	for !done {
		var fileName string
		fmt.Println("File name?")
		fmt.Scanln(&fileName)

		fileNames = append(fileNames, fileName)

		var isDone string
		fmt.Println("Add another file? (Y/N)")
		fmt.Scanln(&isDone)

		done = (strings.ToUpper(isDone) != "Y")
	}
	var ids string

	for _, file := range fileNames {
		fileBytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		ids += fmt.Sprintf("%s\n", string(fileBytes))
	}

	var chunkIds string

	chunks := 0
	for i, id := range strings.Split(ids, "\n") {
		if len(id) <= 0 {
			continue
		}

		if i%(maxPerFile) == maxPerFile-1 {
			chunks += 1
			saveFile(chunks, chunkIds)

			chunkIds = ""
			continue
		}

		chunkIds += fmt.Sprintf("%s\n", id)
	}

	saveFile(chunks, chunkIds)
}
