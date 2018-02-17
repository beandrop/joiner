package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	if _, err := api.AuthTest(); err != nil {
		log.Printf("Failed authentication: %v", err)
		os.Exit(1)
	}
	rtm := api.NewRTM()
	go rtm.ManageConnection()
IncomingEvents:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.SubType == "channel_join" {
					channel, ts, _ := api.DeleteMessage(ev.Channel, ev.Msg.Timestamp)
					log.Println(channel, ts)
				}
			case *slack.RTMError:
				log.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				log.Printf("Invalid credentials")
				break IncomingEvents
			default:
			}
		}
	}
}
