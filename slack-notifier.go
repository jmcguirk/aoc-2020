package main

import (
	"fmt"
	"github.com/nlopes/slack"
)




type SlackManager struct {
	API *slack.Client;
}

func (this *SlackManager) Init() {
	Log.Info("Initializing Slack Manager");
	this.API = slack.New("");
}

func (this *SlackManager) SendMessage(text string) {

	_, _, err := this.API.PostMessage("aoc", slack.MsgOptionText(text, false));
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	//fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}


var Slack = &SlackManager{};
