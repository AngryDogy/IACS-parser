package database

import (
	"errors"
	"reflect"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"parse/iternal/config"
	"parse/pkg/entities"
)

func FindAllChanges(files []entities.FileJSON) ([]*entities.ChangedFile, error) {
	db, err := gorm.Open(postgres.Open(config.FilesDatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.FileDB{})
	if err != nil {
		return nil, err
	}
	filesDB := convertToFileDB(files)
	messages := make([]*entities.ChangedFile, 0)
	for _, f := range filesDB {
		var foundFile entities.FileDB
		result := db.Where("name = ?", f.Name).First(&foundFile)
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			db.Save(&f)
			messages = append(messages, entities.NewChangedFile(&f, "Новый файл"))
		} else {
			if !reflect.DeepEqual(foundFile, f) {
				db.Model(&entities.FileDB{}).Where("name = ?", foundFile.Name).
					Update("description", f.Description).
					Update("notes", f.Notes).
					Update("link", f.Link).
					Update("future_name", f.FutureName).
					Update("future_description", f.FutureDescription).
					Update("future_notes", f.FutureNotes).
					Update("future_link", f.FutureLink)
				messages = append(messages, findDiffs(&foundFile, &f))
			}

		}
	}
	return messages, nil

}

func findDiffs(prevFile, currentFile *entities.FileDB) *entities.ChangedFile {
	var builder strings.Builder
	builder.Grow(1000)
	if prevFile.Description != currentFile.Description {
		builder.WriteString("изменено описание файла, ")
	}
	if prevFile.Notes != currentFile.Notes {
		builder.WriteString("изменено дополнительное описание файла, ")
	}
	if prevFile.Link != currentFile.Link {
		builder.WriteString("изменена ссылка на файл, ")
	}
	if prevFile.FutureName != currentFile.FutureName {
		builder.WriteString("изменено название будущего файла, ")
	}
	if prevFile.FutureDescription != currentFile.FutureDescription {
		builder.WriteString("изменено описание будущего файла, ")
	}
	if prevFile.FutureNotes != currentFile.FutureNotes {
		builder.WriteString("изменено дополнительное описание будущего файла, ")
	}
	if prevFile.FutureLink != currentFile.FutureLink {
		builder.WriteString("изменена ссылка на будущий файл, ")
	}
	return entities.NewChangedFile(currentFile, builder.String())
}

func convertToFileDB(files []entities.FileJSON) []entities.FileDB {
	filesDB := make([]entities.FileDB, 0)

	for _, f := range files {
		filesDB = append(filesDB, *entities.NewFileDB(&f))
	}
	return filesDB
}
