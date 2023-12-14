package csv

import (
	"encoding/csv"
	"os"
	"parse/pkg/entities"
)

func ConvertToCSV(changeFiles []*entities.ChangedFile) error {
	file, err := os.Create("changes.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"Название",
		"Описание",
		"Доп. описание",
		"Ссылка",
		"Будущее название файла",
		"Будущее описание файла",
		"Будущее доп. описание файла",
		"Будущая ссылка на файл",
		"Изменения"}
	writer.Write(header)
	for _, file := range changeFiles {
		err := writer.Write([]string{file.Name,
			file.Description,
			file.Notes,
			file.Link,
			file.FutureName,
			file.FutureDescription,
			file.FutureNotes,
			file.FutureLink,
			file.Changes})
		if err != nil {
			return err
		}
	}
	return nil
}
