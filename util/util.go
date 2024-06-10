package util

import (
	"encoding/csv"
	"os"

	"parse/entities"

	"github.com/microcosm-cc/bluemonday"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CleanFromTags(files *[]entities.FileJSON) {
	p := bluemonday.StripTagsPolicy()
	for _, f := range *files {
		f.ACF.Description = p.Sanitize(f.ACF.Description)
		f.ACF.FutureDescription = p.Sanitize(f.ACF.FutureDescription)
	}
}

func ConvertToCSV(changeFiles []*entities.ChangedFile, nameCSV string) error {
	file, err := os.Create(nameCSV)
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
