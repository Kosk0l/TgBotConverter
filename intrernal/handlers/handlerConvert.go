package handlers

import (
	"context"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//====================================================================================================

func (h *Handler) handlePDF(ctx context.Context, update telegram.Update) {
	h.bot.Send(telegram.NewMessage(update.Message.Chat.ID, "PDF получен"))
	// TODO: реализовать API бизнес-логику
}

func (h *Handler) HandleText(ctx context.Context, update telegram.Update) {
	h.bot.Send(telegram.NewMessage(update.Message.Chat.ID,"Я принимаю только команды и документы"))
}

// Вид json запроса: (все есть в update)

	/*
	"message":{
		"message_id":19,
		"from":{
			"id":7792217214,
			"is_bot":false,
			"first_name":"Николай",
			"username":"kosk0l",
			"language_code":"ru"
		},
		"chat":{
			"id":7792217214,
			"first_name":"Николай",
			"username":"kosk0l",
			"type":"private"
		},
		"date":1766682293,
		"text":"В"
	}
	*/