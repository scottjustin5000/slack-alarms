package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SlackMessage represents the strucure of a Slack Message
type SlackMessage struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

//Attachment represents the structure of a Slack Message Attachment
type Attachment struct {
	Text  string `json:"text"`
	Color string `json:"color"`
	Title string `json:"title"`
}

//BuildSlackMessage builds a Slack Message from the incoming SNS Message
func BuildSlackMessage(message SNSMessage) SlackMessage {
	return SlackMessage{
		Text: fmt.Sprintf("`%s`", message.AlarmName),
		Attachments: []Attachment{
			Attachment{
				Text:  message.NewStateReason,
				Color: "danger",
				Title: "Reason",
			},
		},
	}
}

//PostSlack  is a function to post a message to Slack
func PostSlack(message SlackMessage, webhook string) error {
	client := &http.Client{}
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return err
	}

	return nil
}
