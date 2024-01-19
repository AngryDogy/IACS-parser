package email

import (
	"crypto/tls"
	"errors"
	"gopkg.in/gomail.v2"
	"os"
	"parse/iternal/config"
	"parse/iternal/logger"
)

func SendNotificationEmail() error {
	d := gomail.NewDialer("smtp.gmail.com", 587, config.NotificationEmailAddress, config.NotificationEmailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	for _, e := range config.EmailsToNotify {
		m := gomail.NewMessage()
		m.SetHeader("From", config.NotificationEmailAddress)
		m.SetHeader("To", e)
		m.SetHeader("Subject", "Изменения на сайте МАКО")
		m.SetBody("text/html", "Изменения на сайте МАКО были зафиксированы")
		if _, err := os.Stat("changes.csv"); errors.Is(err, os.ErrNotExist) {
			return err
		}
		m.Attach("changes.csv")
		if err := d.DialAndSend(m); err != nil {
			logger.ErrorLogger.Printf("An error occurred while sending a message! %s\n", err)
		}
		logger.InfoLogger.Printf("Notification message was sent to %s", e)
	}

	return nil
}
