package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fesnasser/file-processor/file"
	"github.com/fesnasser/file-processor/file/line"
	"github.com/fesnasser/file-processor/utils"
)

func main() {
	start := time.Now()

	utils.PrintMemUsage()

	filePath := os.Args[1]

	file.Process(filePath, line.Byte400Handler{})

	utils.PrintMemUsage()

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed = %s", elapsed)
}
