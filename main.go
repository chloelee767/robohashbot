package main

import (
	"fmt"
	"github.com/chloelee767/robohashbot/robohash"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)


func main() {
	bot, botCreationErr := tgbotapi.NewBotAPI("1192566547:AAEQ_FH2mCvL2spwloOnAaMneaa7_ya_3f8")
	if botCreationErr != nil {
		fmt.Println(botCreationErr)
		return
	}

	updates, updateChanErr := bot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 60})
	if updateChanErr != nil {
		fmt.Println(updateChanErr)
		return
	}

	fmt.Println("Listening...")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := update.Message
		chatID := message.Chat.ID
		r, _ := robohash.NewRobohash("meow!!", robohash.Cat)
		response := tgbotapi.NewPhotoShare(chatID, r.GetUrl())
		bot.Send(response)
	}
}
