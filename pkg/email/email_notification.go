package email

import (
	"fmt"
	"net/smtp"
	"parse/iternal/config"
	"parse/pkg/entities"
	"strings"
)

/*type Sender struct {
	auth smtp.Auth
}

type Message struct {
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func New() *Sender {
	auth := smtp.PlainAuth("", config.NotificationEmailAddress, config.NotificationEmailPassword, "smtp.gmail.com")
	return &Sender{auth}
}

func (s *Sender) Send(m *Message) error {
	return smtp.SendMail(fmt.Sprintf("%s:%s", "smtp.gmail.com", "587"), s.auth, config.NotificationEmailAddress, m.To, m.ToBytes())
}

func NewMessage(s, b string) *Message {
	return &Message{Subject: s, Body: b, Attachments: make(map[string][]byte)}
}

func (m *Message) AttachFile(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	buf.WriteString(m.Body)
	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}*/

func SendNotificationEmail(messages []*entities.ChangeMessage) {

	/*sender := New()
	m := NewMessage("Test", "Body message.")
	m.To = []string{"vladhober@mail.ru"}
	m.AttachFile("data.txt")
	fmt.Println(sender.Send(m))*/

	to := []string{
		"vladhober@mail.ru",
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := constructMessage(messages)
	//message := "IACS-Ballast-w"
	auth := smtp.PlainAuth("", config.NotificationEmailAddress, config.NotificationEmailPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, config.NotificationEmailAddress, to, []byte(message))
	if err != nil {
		panic(err)
	}
}

func constructMessage(messages []*entities.ChangeMessage) string {
	var builder strings.Builder
	builder.Grow(10000)
	for _, m := range messages {
		builder.WriteString(fmt.Sprintf("%s: %s\n", m.FileName, m.Content))
	}
	return builder.String()
}
