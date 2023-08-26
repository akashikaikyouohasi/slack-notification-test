package main

import (
	"context"
	"fmt"
	"os"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func handler (ctx context.Context, event events.CloudWatchEvent) {
	// アクセストークンを使用してクライアントを生成する
	slack_token := os.Getenv("SLACK_OAUTH_TOKEN")
	channel_name := os.Getenv("SLACK_CHANNEL")

	slack_connection := slack.New(slack_token)

	fmt.Printf("channel_name = %s\n", channel_name)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := slack_connection.PostMessage(channel_name, slack.MsgOptionText("Hello World", true))
	if err != nil {
		panic(err)
	}
}

func main() {
	lambda.Start(handler)
}
