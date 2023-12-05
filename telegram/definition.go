package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramBot struct {
	*tgbotapi.BotAPI
	chatRoomId int64
}
