package handlers

import (
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//====================================================================================================

func HandleUpdate(bot *telegram.BotAPI, update telegram.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		handleCommand(bot, update)
		return
	}

	if update.Message.Document != nil {
		handleDocument(bot, update)
		return
	}

	handleText(bot, update)
}

//====================================================================================================

func handleCommand(bot *telegram.BotAPI, update telegram.Update) {
	chatID := update.Message.Chat.ID

	switch update.Message.Command() {
	case "start":
		bot.Send(telegram.NewMessage(
			chatID,
			"Привет! Отправь документ для конвертации.",
		))
	default:
		bot.Send(telegram.NewMessage(
			chatID,
			"Неизвестная команда",
		))
	}
}

func handleDocument(bot *telegram.BotAPI, update telegram.Update) {
	doc := update.Message.Document
	chatID := update.Message.Chat.ID

	filename := strings.ToLower(doc.FileName)

	switch {
	case strings.HasSuffix(filename, ".pdf"):
		handlePDF(bot, chatID)

	case strings.HasSuffix(filename, ".docx"):
		handleDOCX(bot, chatID)

	case strings.HasSuffix(filename, ".xlsx"):
		handleXLSX(bot, chatID)

	default:
		bot.Send(telegram.NewMessage(
			chatID,
			"Этот тип файла пока не поддерживается",
		))
	}
}

//====================================================================================================

func handlePDF(bot *telegram.BotAPI, chatID int64) {
	bot.Send(telegram.NewMessage(
		chatID,
		"PDF получен",
	))
}

func handleDOCX(bot *telegram.BotAPI, chatID int64) {
	bot.Send(telegram.NewMessage(
		chatID,
		"DOCX получен",
	))
}

func handleXLSX(bot *telegram.BotAPI, chatID int64) {
	bot.Send(telegram.NewMessage(
		chatID,
		"XLSX получен",
	))
}

func handleText(bot *telegram.BotAPI, update telegram.Update) {
	bot.Send(telegram.NewMessage(
		update.Message.Chat.ID,
		"Я понимаю команды и документы",
	))
}