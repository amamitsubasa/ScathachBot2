package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const token string = "NTg2ODI4Njc2NTYyNDE5NzI4.XPt1gg.s-VWxK2en4SG6niljYLy4rixw8I"
var BotID string 
func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s * discordgo.Session, m *discordgo.MessageCreate) {
	
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}