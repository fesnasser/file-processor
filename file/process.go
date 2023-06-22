package file

import (
	"bufio"
	"log"
	"os"

	"github.com/fesnasser/file-processor/file/line"
)

const workersQuantity = 20000

func Process(filePath string, lineHandler line.Handler) {
	channel := make(chan string)

	createWorkers(channel, lineHandler)

	readFile(filePath, channel)
}

func readFile(filePath string, channel chan string) {
	readFile, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		channel <- fileScanner.Text()
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createWorkers(channel chan string, lineHandler line.Handler) {
	for i := 0; i < workersQuantity; i++ {
		go readFromChannel(channel, lineHandler)
	}
}

func readFromChannel(data chan string, lineHandler line.Handler) {
	for line := range data {
		lineHandler.Handle(line)
	}
}
