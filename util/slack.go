package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func SlackNoticeTransaction(text string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// 1. Slackから取得したWebhook URLを設定
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")

	// 2. 送信するメッセージを作成（ライブラリの構造体を使用）
	msg := slack.WebhookMessage{
		Text: text,
	}

	// 3. メッセージを送信
	err = slack.PostWebhook(webhookURL, &msg)
	if err != nil {
		log.Fatalf("Slackへの通知に失敗しました: %v", err)
	}

	fmt.Println("Slackへの通知が成功しました。")
}