package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"parse/iternal/config"
	"parse/pkg/entities"
)

func GetAllEmails() ([]entities.Email, error) {
	db, err := gorm.Open(postgres.Open(config.EmailsDatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var emails []entities.Email
	result := db.Find(&emails)
	if result.Error != nil {
		return nil, result.Error
	}
	return emails, nil
}
