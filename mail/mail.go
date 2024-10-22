package mail

import (
    "context"
    "fmt"
    "time"
    "github.com/mailgun/mailgun-go/v4"
    "os"
)

func SendEmail(email string, code string) {
    fmt.Println("メールをこれから送る")

    mailgunDomain := os.Getenv("MAILGUN_DOMAIN")
    mailgunPrivateAPIKey := os.Getenv("MAILGUN_PRIVATE_API_KEY")
    senderEmail := os.Getenv("SENDER_EMAIL")
    recipientEmail := email
    
    // Mailgunクライアントの作成
    mg := mailgun.NewMailgun(mailgunDomain, mailgunPrivateAPIKey)

    // メッセージの作成
    subject := "segon_pixの認証コード"
    body := "認証コード: "+code
    message := mg.NewMessage(
        senderEmail,
        subject,
        body,
        recipientEmail,
    )

    
    // コンテキストの作成（タイムアウト設定）
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()

    // メールの送信
    resp, id, err := mg.Send(ctx, message)
    if err != nil {
        fmt.Println("メールの送信に失敗しました:", err)
        return
    }

    fmt.Printf("メールが正常に送信されました。ID: %s Resp: %s\n", id, resp)
}
