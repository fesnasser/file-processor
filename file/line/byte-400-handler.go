package line

import (
	"fmt"

	"github.com/fesnasser/file-processor/database"
	"github.com/fesnasser/file-processor/model"
)

type Byte400Handler struct{}

func (h Byte400Handler) Handle(line []byte) {
	db := database.GetCon()

	newLine := model.Line{Valid: len(line) == 400}

	err := db.Create(&newLine).Error

	if err != nil {
		fmt.Println("Erro ao salvar a linha na base", err)
	}
}
