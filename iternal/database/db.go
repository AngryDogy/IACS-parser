package database

import (
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"parse/iternal/entities"
)

func FindAllChanges(files []entities.File) ([]entities.MessageChange, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("PATH_TO_DATABASE")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.FileDB{})
	if err != nil {
		return nil, err
	}
	filesDB := fromFileToFileDB(files)
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

func fromFileToFileDB(files []entities.File) []entities.FileDB {
	filesDB := make([]entities.FileDB, 0)

	for _, f := range files {
		filesDB = append(filesDB, *entities.NewFileDB(f.RenderedName.Name, f.ReleaseDate, f.ModifiedDate, f.ACF.Description, f.ACF.Notes, f.ACF.Link))
	}
	return filesDB
}
