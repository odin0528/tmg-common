package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewTelegramBot(botToken string, chatRommId int64) (*TelegramBot, error) {
	tgBot, err := tgbotapi.NewBotAPI(botToken)

	return &TelegramBot{BotAPI: tgBot, ChatRoomId: chatRommId}, err
}
