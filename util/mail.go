package util

import (
    "github.com/jordan-wright/email"
	"log"
	"net/smtp"
    "os"
)

func SendEmail(toEmail string, code string) error {
	gmailpass := os.Getenv("GMAIL_APP_PASS")
	senderEmail := os.Getenv("SENDER_EMAIL")

	e := email.NewEmail()
	e.From = senderEmail
	e.To = []string{toEmail}
	e.Subject = "segon_pixの確認コード"
	e.Text = []byte("見覚えのない連絡でしたら無視してください\n確認コード:" + code)
	// GmailのSMTPサーバ情報
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	smtpAddr := smtpServer + ":" + smtpPort

	err := e.Send(smtpAddr, smtp.PlainAuth(
		"",          // identity（通常は空文字）
		senderEmail, // Gmailアドレス
		gmailpass,   // Appパスワード
		smtpServer,  // ホスト名（smtp.gmail.com）
	))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully")
	return nil
}
