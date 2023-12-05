package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *TelegramBot) SendMessage(message string) error {
	msg := tgbotapi.NewMessage(bot.chatRoomId, message)
	_, err := bot.Send(msg)

	if err != nil {
		return err
	}

	return nil
}
