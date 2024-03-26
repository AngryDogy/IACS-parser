package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

var (
	FilesDatabasePath         string
	ResolutionsURL            string
	SectionNumberStart        int
	SectionNumberEnd          int
	ProceduresURL             string
	PositionPapersURL         string
	HostEmailServer           string
	PortEmailServer           int
	UsernameEmailServer       string
	PasswordEmailServer       string
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
	SectionNumberStart, err = strconv.Atoi(os.Getenv("SECTION_NUMBER_START"))
	if err != nil {
		return err
	}
	SectionNumberEnd, err = strconv.Atoi(os.Getenv("SECTION_NUMBER_END"))
	if err != nil {
		return err
	}
	ResolutionsURL = os.Getenv("RESOLUTIONS_URL")
	ProceduresURL = os.Getenv("PROCEDURES_URL")
	PositionPapersURL = os.Getenv("POSITION_PAPERS_URL")
	HostEmailServer = os.Getenv("HOST_EMAIL_SERVER")
	PortEmailServer, err = strconv.Atoi(os.Getenv("PORT_EMAIL_SERVER"))
	if err != nil {
		return err
	}
	UsernameEmailServer = os.Getenv("USERNAME_EMAIL_SERVER")
	PasswordEmailServer = os.Getenv("PASSWORD_EMAIL_SERVER")
	NotificationEmailAddress = os.Getenv("NOTIFICATION_EMAIL_ADDRESS")
	NotificationEmailPassword = os.Getenv("NOTIFICATION_EMAIL_PASSWORD")
	EmailsToNotify = strings.Split(os.Getenv("EMAILS_TO_NOTIFY"), ",")

	return nil
}
