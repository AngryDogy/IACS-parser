package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	FilesDatabasePath         string
	ResolutionsURL            string
	ProceduresURL             string
	PositionPapersURL         string
	NotificationEmailAddress  string
	NotificationEmailPassword string
)

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	FilesDatabasePath = os.Getenv("PATH_TO_DATABASE")
	ResolutionsURL = os.Getenv("RESOLUTIONS_URL")
	ProceduresURL = os.Getenv("PROCEDURES_URL")
	PositionPapersURL = os.Getenv("POSITION_PAPERS_URL")
	NotificationEmailAddress = os.Getenv("NOTIFICATION_EMAIL_ADDRESS")
	NotificationEmailPassword = os.Getenv("NOTIFICATION_EMAIL_PASSWORD")
	return nil
}
