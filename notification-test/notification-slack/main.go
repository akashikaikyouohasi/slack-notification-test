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
	_, _, err := slack_connection.PostMessage(channel_name, createMessage())
	if err != nil {
		panic(err)
	}
}

func main() {
	lambda.Start(handler)
}

func createMessage() slack.MsgOption {
	// sample: 
	// https://blog.zaim.co.jp/n/n9eb41aabf5c6
	// https://pkg.go.dev/github.com/slack-go/slack#SlackErrorResponse
	// https://github.com/slack-go/slack/blob/master/examples/blocks/README.md
	// https://app.slack.com/block-kit-builder/T8XQJV0F8#%7B%22blocks%22:%5B%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22mrkdwn%22,%22text%22:%22You%20have%20a%20new%20request:%5Cn*%3CfakeLink.toEmployeeProfile.com%7CFred%20Enriquez%20-%20New%20device%20request%3E*%22%7D%7D,%7B%22type%22:%22section%22,%22fields%22:%5B%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Type:*%5CnComputer%20(laptop)%22%7D,%7B%22type%22:%22mrkdwn%22,%22text%22:%22*When:*%5CnSubmitted%20Aut%2010%22%7D,%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Last%20Update:*%5CnMar%2010,%202015%20(3%20years,%205%20months)%22%7D,%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Reason:*%5CnAll%20vowel%20keys%20aren't%20working.%22%7D,%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Specs:*%5Cn%5C%22Cheetah%20Pro%2015%5C%22%20-%20Fast,%20really%20fast%5C%22%22%7D%5D%7D,%7B%22type%22:%22actions%22,%22elements%22:%5B%7B%22type%22:%22button%22,%22text%22:%7B%22type%22:%22plain_text%22,%22emoji%22:true,%22text%22:%22Approve%22%7D,%22value%22:%22click_me_123%22%7D,%7B%22type%22:%22button%22,%22text%22:%7B%22type%22:%22plain_text%22,%22emoji%22:true,%22text%22:%22Deny%22%7D,%22value%22:%22click_me_123%22%7D%5D%7D%5D%7D

	// Mention
	mentionText := slack.NewTextBlockObject("mrkdwn", "<@U05PBSSAG06>", false, false)
	mentionSection := slack.NewSectionBlock(mentionText, nil, nil)

	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", "You have a new notification:", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	envField := slack.NewTextBlockObject("mrkdwn", "*Env:*\nStaging", false, false)
	pipelineField := slack.NewTextBlockObject("mrkdwn", "*Pipeline:*\ntest-pipeline", false, false)
	idField := slack.NewTextBlockObject("mrkdwn", "*ID:*\nddda-dfada-dfasf-dadds-test", false, false)
	statusField := slack.NewTextBlockObject("mrkdwn", "*Status:*\nSucceeded", false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, envField)
	fieldSlice = append(fieldSlice, pipelineField)
	fieldSlice = append(fieldSlice, idField)
	fieldSlice = append(fieldSlice, statusField)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	return slack.MsgOptionBlocks(
		mentionSection,
		headerSection,
		fieldsSection,
	)
}