package main

import (
	"crypto/tls"
    "fmt"
	"github.com/go-mail/mail"
	"os"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

    // Gmail の SMTP サーバー情報
    authEmail := os.Getenv("GMAIL_EMAIL")           // あなたのGmailアドレス
    authPassword := os.Getenv("Email_Password")

    // 送信先のメールアドレス
    to := authEmail

    // メールの内容
    subject := "テストメール"
    body := "これは Go から送信されたテストメールです。"

    // メールメッセージの作成
    m := mail.NewMessage()
    m.SetHeader("From", authEmail)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    // ダイヤラーを作成し、TLS設定を追加
    d := mail.NewDialer("smtp.gmail.com", 587, authEmail, authPassword)

    // TLS設定を行う
    d.TLSConfig = &tls.Config{
        InsecureSkipVerify: false,
        ServerName:         "smtp.gmail.com",
    }

    // メールの送信
    if err := d.DialAndSend(m); err != nil {
        fmt.Println("メールの送信に失敗しました:", err)
        return
    }

    fmt.Println("メールが正常に送信されました。")
}
