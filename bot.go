package bot

import (
	"fmt"
	"config"
	"github.com/bwmarrin/discordgo"
)

var BotID string 
var goBat = *discordgo.Session

func Start() {
	gotBot, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := gotBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	gotBot.AddHandler(messageHandler)

	err = gotBot.Open()

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