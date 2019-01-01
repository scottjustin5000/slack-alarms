package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/scottjustin5000/slack-alarms/pkg"
)

func handler(request pkg.Request) error {
	var snsMessage pkg.SNSMessage
	err := json.Unmarshal([]byte(request.Records[0].SNS.SNSMessage), &snsMessage)
	if err != nil {
		return err
	}

	log.Printf("New alarm: %s - Reason: %s", snsMessage.AlarmName, snsMessage.NewStateReason)
	slackMessage := pkg.BuildSlackMessage(snsMessage)
	pkg.PostSlack(slackMessage, os.Getenv("SLACK_WEBHOOK"))
	log.Println("Notification has been sent")
	return nil
}

func main() {
	lambda.Start(handler)
}
