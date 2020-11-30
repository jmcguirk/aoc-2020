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
	this.API = slack.New("xoxp-683360277315-696609298631-709308949490-29d31a0574747242a82ec365e1d9afd2");
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
