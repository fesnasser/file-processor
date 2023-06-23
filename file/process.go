package file

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/fesnasser/file-processor/file/line"
)

// Não pode ser igual ou maior que o máximo de conexões do banco
const workerPoolSize = 50

func Process(filePath string, lineHandler line.Handler) {
	var wg sync.WaitGroup

	channel := make(chan []byte, workerPoolSize)

	createWorkers(channel, lineHandler, &wg)

	readFile(filePath, channel)

	close(channel)
	wg.Wait()
}

func readFile(filePath string, channel chan []byte) {
	readFile, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		channel <- fileScanner.Bytes()
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createWorkers(channel chan []byte, lineHandler line.Handler, wg *sync.WaitGroup) {
	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)

		go readFromChannel(channel, lineHandler, wg)
	}
}

func readFromChannel(channel chan []byte, lineHandler line.Handler, wg *sync.WaitGroup) {
	for line := range channel {
		lineHandler.Handle(line)
	}
	wg.Done()
}
