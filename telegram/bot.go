package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewTelegramBot(botToken string, chatRommId int64) (*TelegramBot, error) {
	tgBot, err := tgbotapi.NewBotAPI(botToken)

	return &TelegramBot{BotAPI: tgBot, chatRoomId: chatRommId, commandFuncMap: map[string]func(string) (string, error){}}, err
}

func (bot *TelegramBot) SendMessage(message string) error {
	msg := tgbotapi.NewMessage(bot.chatRoomId, message)
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := bot.Send(msg)

	if err != nil {
		return err
	}

	return nil
}

func (bot *TelegramBot) Serve() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			bot.SendMessage(update.Message.Text + " is not command")
			continue
		}

		bot.funcMutex.Lock()

		cFunc, ok := bot.commandFuncMap[update.Message.Command()]
		if !ok {
			bot.SendMessage("command not supported")
			bot.funcMutex.Unlock()
			continue
		}

		repTex, err := cFunc(update.Message.Text)
		if err != nil {
			bot.SendMessage(err.Error())
		} else {
			if repTex == "" {
				bot.SendMessage("success")
			} else {
				bot.SendMessage(repTex)
			}
		}

		bot.funcMutex.Unlock()
	}
}

func (bot *TelegramBot) Handle(key string, pFunc func(string) (string, error)) {
	bot.funcMutex.Lock()
	defer bot.funcMutex.Unlock()

	bot.commandFuncMap[key] = pFunc
}
