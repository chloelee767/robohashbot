package main

import (
	"fmt"
	"strconv"
	"strings"

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

	list := []robohash.Robohash{} // or make([]robohash.Robohash, 0)

	commandMap := map[string]command{
		"add":    addRobohash,
		"delete": deleteRobohash,
		"list":   listRobohashes,
		"show":   showRobohash,
		"": func(chatID int64, _ *[]robohash.Robohash, _ string) tgbotapi.Chattable {
			return tgbotapi.NewMessage(chatID, "That's not a command!")
		},
	}

	fmt.Println("Listening...")
	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := update.Message
		chatID := message.Chat.ID
		var response tgbotapi.Chattable
		if commandFn, ok := commandMap[message.Command()]; ok {
			response = commandFn(chatID, &list, message.CommandArguments())
		} else {
			response = tgbotapi.NewMessage(chatID, "Unknown command!")
		}
		bot.Send(response)
	}
}

type command func(int64, *[]robohash.Robohash, string) tgbotapi.Chattable

func addRobohash(chatID int64, list *[]robohash.Robohash, args string) tgbotapi.Chattable {
	correctSyntaxMsg := "To add a robohash, use /add TYPE NAME\nTYPE must be one of robot, monster, newRobot, cat or human"
	nameToType := map[string]robohash.Type{
		robohash.Robot.Name():    robohash.Robot,
		robohash.Monster.Name():  robohash.Monster,
		robohash.NewRobot.Name(): robohash.NewRobot,
		robohash.Cat.Name():      robohash.Cat,
		robohash.Human.Name():    robohash.Human,
	}

	split := strings.SplitN(args, " ", 2)
	if len(split) < 2 {
		return tgbotapi.NewMessage(chatID, fmt.Sprint("Your robohash needs a name!\n", correctSyntaxMsg))
	}

	rTypeString, name := split[0], split[1]
	rType, ok := nameToType[rTypeString]
	if !ok {
		return tgbotapi.NewMessage(chatID, fmt.Sprint("Unknown type of robohash!\n", correctSyntaxMsg))
	}
	r, err := robohash.NewRobohash(name, rType)
	if err != nil {
		return tgbotapi.NewMessage(chatID, err.Error())
	}

	*list = append(*list, r)
	return tgbotapi.NewMessage(chatID, "Added robohash!")
}

func listRobohashes(chatID int64, list *[]robohash.Robohash, _ string) tgbotapi.Chattable {
	if len(*list) == 0 {
		return tgbotapi.NewMessage(chatID, "You have no robohashes! Add one using the /add command")
	}
	var sb strings.Builder
	for i, robohash := range *list {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, robohash))
	}
	return tgbotapi.NewMessage(chatID, sb.String())
}

func getListPosition(args string, listSize int) (int, error) {
	listPosition, err := strconv.Atoi(args)
	if err != nil || listPosition < 1 || listPosition > listSize {
		return listPosition, fmt.Errorf("Invalid list position!")
	}
	return listPosition, nil
}

func showRobohash(chatID int64, list *[]robohash.Robohash, args string) tgbotapi.Chattable {
	listPosition, err := getListPosition(args, len(*list))
	if err != nil {
		return tgbotapi.NewMessage(chatID, fmt.Sprint(err))
	}
	robohash := (*list)[listPosition-1]
	photoMessage := tgbotapi.NewPhotoShare(chatID, robohash.GetUrl())
	photoMessage.Caption = robohash.String()
	return photoMessage
}

func deleteRobohash(chatID int64, list *[]robohash.Robohash, args string) tgbotapi.Chattable {
	listPosition, err := getListPosition(args, len(*list))
	if err != nil {
		return tgbotapi.NewMessage(chatID, fmt.Sprint(err))
	}
	i := listPosition - 1
	*list = append((*list)[:i], (*list)[i+1:]...)
	return tgbotapi.NewMessage(chatID, "Removed robohash!")
}
