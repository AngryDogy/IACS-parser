package email

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"os"

	"parse/config"
	"parse/logger"

	"gopkg.in/gomail.v2"
)

type loginAuth struct {
	username string
	password string
	host     string
}

func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if !server.TLS {
		advertised := false
		for _, mechanism := range server.Auth {
			if mechanism == "LOGIN" {
				advertised = true
				break
			}
		}
		if !advertised {
			return "", nil, errors.New("gomail: unencrypted connection")
		}
	}
	if server.Name != a.host {
		return "", nil, errors.New("gomail: wrong host name")
	}
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		return nil, nil
	}

	switch {
	case bytes.Equal(fromServer, []byte("Username:")):
		return []byte(a.username), nil
	case bytes.Equal(fromServer, []byte("Password:")):
		return []byte(a.password), nil
	default:
		return nil, fmt.Errorf("gomail: unexpected server challenge: %s", fromServer)
	}
}

func SendNotificationEmail(nameTXT, nameCSV string) error {
	d := gomail.NewDialer(config.HostEmailServer, config.PortEmailServer, config.NotificationEmailAddress, config.NotificationEmailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: config.HostEmailServer}
	if config.UsernameEmailServer != "" {
		d.Auth = LoginAuth(config.UsernameEmailServer, config.PasswordEmailServer, config.HostEmailServer)
	}

	for _, e := range config.EmailsToNotify {
		m := gomail.NewMessage()
		m.SetHeader("From", config.NotificationEmailAddress)
		m.SetHeader("To", e)
		m.SetHeader("Subject", "Изменения на сайте МАКО")
		m.SetBody("text/html", "Изменения на сайте МАКО были зафиксированы")
		if _, err := os.Stat(nameCSV); errors.Is(err, os.ErrNotExist) {
			return err
		}
		if _, err := os.Stat(nameTXT); errors.Is(err, os.ErrNotExist) {
			return err
		}
		m.Attach(nameCSV)
		m.Attach(nameTXT)
		if err := d.DialAndSend(m); err != nil {
			logger.ErrorLogger.Printf("An error occurred while sending a message! %s\n", err)
		} else {
			logger.InfoLogger.Printf("Notification message was sent to %s\n", e)
		}

	}
	return nil
}
