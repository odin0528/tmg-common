package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(botToken string, chatRoomID int64, message string) error {
	bot, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatRoomID, message)
	_, err = bot.Send(msg)

	if err != nil {
		return err
	}

	return nil
}
