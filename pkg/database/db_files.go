package database

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"parse/internal/config"
	"parse/pkg/entities"
)

func FindAllChanges(files []entities.FileJSON, changesInfoFileName string) ([]*entities.ChangedFile, error) {
	changesInfoFile, err := os.Create(changesInfoFileName)
	if err != nil {
		return nil, err
	}
	defer changesInfoFile.Close()
	db, err := gorm.Open(postgres.Open(config.FilesDatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.FileDB{})
	if err != nil {
		return nil, err
	}
	filesDB := convertToFileDB(files)
	changedFiles := make([]*entities.ChangedFile, 0)
	for _, f := range filesDB {
		var foundFile entities.FileDB
		result := db.Where("name = ?", f.Name).First(&foundFile)
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			db.Save(&f)
			changedFiles = append(changedFiles, entities.NewChangedFile(&f, "New file"))
			changesInfoFile.WriteString(fmt.Sprintf("New file was added: %s\n\n", f.Name))
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
				changedFiles = append(changedFiles, entities.NewChangedFile(&f, findDiffs(&foundFile, &f, changesInfoFile)))
			}

		}
	}
	return changedFiles, nil

}

func findDiffs(prevFile, currentFile *entities.FileDB, changesInfoFile *os.File) string {
	var builder strings.Builder
	builder.Grow(1000)
	changesInfoFile.WriteString(fmt.Sprintf("Diffs for %s file\n", currentFile.Name))
	if prevFile.Description != currentFile.Description {
		changesInfoFile.WriteString("Description was changed!\n")
		changesInfoFile.WriteString(fmt.Sprintf("Old: %s\n", prevFile.Description))
		changesInfoFile.WriteString(fmt.Sprintf("New: %s\n", currentFile.Description))
		builder.WriteString("description changed, ")
	}
	if prevFile.Notes != currentFile.Notes {
		changesInfoFile.WriteString("Notes were changed!\n")
		changesInfoFile.WriteString(fmt.Sprintf("Old: %s\n", prevFile.Notes))
		changesInfoFile.WriteString(fmt.Sprintf("New: %s\n", currentFile.Notes))
		builder.WriteString("notes changed, ")
	}
	if prevFile.Link != currentFile.Link {
		changesInfoFile.WriteString("Link was changed!\n")
		changesInfoFile.WriteString(fmt.Sprintf("Old: %s\n", prevFile.Link))
		changesInfoFile.WriteString(fmt.Sprintf("New: %s\n", currentFile.Link))
		builder.WriteString("link changed, ")
	}
	if prevFile.FutureName != currentFile.FutureName {
		changesInfoFile.WriteString("Future name was changed!\n")
		changesInfoFile.WriteString(fmt.Sprintf("Old: %s\n", prevFile.FutureName))
		changesInfoFile.WriteString(fmt.Sprintf("New: %s\n", currentFile.FutureName))
		builder.WriteString("future name changed, ")
	}
	if prevFile.FutureDescription != currentFile.FutureDescription {
		changesInfoFile.WriteString("Future description was changed!\n")
		changesInfoFile.WriteString(fmt.Sprintf("Old: %s\n", prevFile.FutureDescription))
		changesInfoFile.WriteString(fmt.Sprintf("New: %s\n", currentFile.FutureDescription))
		builder.WriteString("future description changed, ")
	}
	if prevFile.FutureNotes != currentFile.FutureNotes {
		changesInfoFile.WriteString("Future notes were changed!\n")
		changesInfoFile.WriteString(fmt.Sprintf("Old: %s\n", prevFile.FutureNotes))
		changesInfoFile.WriteString(fmt.Sprintf("New: %s\n", currentFile.FutureNotes))
		builder.WriteString("future notes changed, ")
	}
	if prevFile.FutureLink != currentFile.FutureLink {
		changesInfoFile.WriteString("Future link was changed!\n")
		changesInfoFile.WriteString(fmt.Sprintf("Old: %s\n", prevFile.FutureLink))
		changesInfoFile.WriteString(fmt.Sprintf("New: %s\n", currentFile.FutureLink))
		builder.WriteString("future link changed, ")
	}
	changesInfoFile.WriteString("\n")
	return builder.String()
}

func convertToFileDB(files []entities.FileJSON) []entities.FileDB {
	filesDB := make([]entities.FileDB, 0)

	for _, f := range files {
		filesDB = append(filesDB, *entities.NewFileDB(&f))
	}
	return filesDB
}
