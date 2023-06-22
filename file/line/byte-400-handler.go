package line

import (
	"log"
	"strings"
	"time"
)

type Byte400Handler struct{}

func (h Byte400Handler) Handle(line string) {
	validate(line)
}

func validate(line string) {
	count := strings.Count(line, "0")

	if count < 400 {
		log.Fatalf("Invalid line: %s", line)
	}

	time.Sleep(1 * time.Second)
}
