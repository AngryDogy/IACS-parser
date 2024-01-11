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
	writer.Comma = ';'
	defer writer.Flush()
	header := []string{"Name",
		"Description",
		"Notes",
		"Link",
		"Future name",
		"Future description",
		"Future notes",
		"Future link",
		"Changes"}
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
