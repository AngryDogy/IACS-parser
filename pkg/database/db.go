package database

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"parse/iternal/config"
	"parse/pkg/entities"
)

func FindAllChanges(files []entities.FileJSON) ([]*entities.ChangeMessage, error) {
	db, err := gorm.Open(postgres.Open(config.FilesDatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.FileDB{})
	if err != nil {
		return nil, err
	}
	filesDB := convertToFileDB(files)
	messages := make([]*entities.ChangeMessage, 0)
	for _, f := range filesDB {
		var foundFile entities.FileDB
		result := db.Where("name = ?", f.Name).First(&foundFile)
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			db.Save(&f)
			messages = append(messages, entities.NewMessageWithContent(f.Name, fmt.Sprintf("New file was added! Link: %s", f.Link)))
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

func findDiffs(prevFile, currentFile *entities.FileDB) *entities.ChangeMessage {
	var builder strings.Builder
	builder.Grow(1000)
	builder.WriteString("file was edited! ")
	if prevFile.Description != currentFile.Description {
		builder.WriteString(fmt.Sprintf("new file's description: %s, ", currentFile.Description))
	}
	if prevFile.Notes != currentFile.Notes {
		builder.WriteString(fmt.Sprintf("new file's notes: %s, ", currentFile.Notes))
	}
	if prevFile.Link != currentFile.Link {
		builder.WriteString(fmt.Sprintf("new file's link: %s, ", currentFile.Link))
	}
	if prevFile.FutureName != currentFile.FutureName {
		builder.WriteString(fmt.Sprintf("new file's future name: %s, ", currentFile.FutureName))
	}
	if prevFile.FutureDescription != currentFile.FutureDescription {
		builder.WriteString(fmt.Sprintf("new file's future description: %s, ", currentFile.FutureDescription))
	}
	if prevFile.FutureNotes != currentFile.FutureNotes {
		builder.WriteString(fmt.Sprintf("new file's future notes: %s, ", currentFile.FutureNotes))
	}
	if prevFile.FutureLink != currentFile.FutureLink {
		builder.WriteString(fmt.Sprintf("new file's future link: %s, ", currentFile.FutureLink))
	}
	return entities.NewMessageWithContent(currentFile.Name, builder.String())
}

func convertToFileDB(files []entities.FileJSON) []entities.FileDB {
	filesDB := make([]entities.FileDB, 0)

	for _, f := range files {
		filesDB = append(filesDB, *entities.NewFileDB(&f))
	}
	return filesDB
}
