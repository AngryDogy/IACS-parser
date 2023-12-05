package database

import (
	"errors"
	"parse/iternal/config"
	"parse/pkg/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func FindAllChanges(files []entities.FileJSON) ([]entities.MessageChange, error) {
	db, err := gorm.Open(postgres.Open(config.FilesDatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.FileDB{})
	if err != nil {
		return nil, err
	}
	filesDB := convertToFileDB(files)
	messageChanges := make([]entities.MessageChange, 0)
	for _, f := range filesDB {
		var foundFile entities.FileDB
		result := db.Where("name = ?", f.Name).First(&foundFile)
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			db.Save(&f)
			messageChanges = append(messageChanges, *entities.NewMessageChange(f.Name, f.Description, f.Notes, f.Link, "New file"))
		} else {
			if foundFile.Description != f.Description || foundFile.Notes != f.Notes {
				db.Model(&entities.FileDB{}).Where("name = ?", foundFile.Name).
					Update("description", f.Description).
					Update("notes", f.Notes).
					Update("link", f.Link)
				var changes string
				if foundFile.Notes != f.Notes {
					changes = "edited the file's notes"
				}
				if foundFile.Description != f.Description {
					changes = "edited the file's description"
				}
				messageChanges = append(messageChanges, *entities.NewMessageChange(f.Name, f.Description, f.Notes, f.Link, changes))
			}

		}
	}
	return messageChanges, nil

}

func convertToFileDB(files []entities.FileJSON) []entities.FileDB {
	filesDB := make([]entities.FileDB, 0)

	for _, f := range files {
		filesDB = append(filesDB, *entities.NewFileDB(&f))
	}
	return filesDB
}
