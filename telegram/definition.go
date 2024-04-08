package telegram

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	*tgbotapi.BotAPI
	chatRoomId     int64
	commandFuncMap map[string]func(string) (string, error)
	funcMutex      sync.RWMutex
}
