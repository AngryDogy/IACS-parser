package config

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

var (
	FilesDatabasePath         string
	ResolutionsURL            string
	ProceduresURL             string
	PositionPapersURL         string
	NotificationEmailAddress  string
	NotificationEmailPassword string
	EmailsToNotify            []string
)

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	FilesDatabasePath = os.Getenv("PATH_TO_FILE_DATABASE")
	ResolutionsURL = os.Getenv("RESOLUTIONS_URL")
	ProceduresURL = os.Getenv("PROCEDURES_URL")
	PositionPapersURL = os.Getenv("POSITION_PAPERS_URL")
	NotificationEmailAddress = os.Getenv("NOTIFICATION_EMAIL_ADDRESS")
	NotificationEmailPassword = os.Getenv("NOTIFICATION_EMAIL_PASSWORD")
	EmailsToNotify = strings.Split(os.Getenv("EMAILS_TO_NOTIFY"), ",")
	return nil
}
