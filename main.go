package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("CHANNEL_ID", "C05SB92S2NP")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{""}
}
