package main

import (
	log "log"
	os "os"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	
	//TODO: перейти на github.com/joho/godotenv
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN is not set")
	}

	// Создаём бота
	bot, err := telegram.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true // можно отключить потом
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем получение апдейтов
	u := telegram.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	// Основной цикл обработки
	for update := range updates {
		if update.Message == nil { // пропускаем не-сообщения
			continue
		}

		// Реакция на текст /start
		if update.Message.Text == "/start" {
			msg := telegram.NewMessage(update.Message.Chat.ID, "Привет! Я базовый бот-конвертер. Отправь мне файл или команду.")
			bot.Send(msg)
			continue
		}

		// Если пришёл документ
		if update.Message.Document != nil {
			msg := telegram.NewMessage(update.Message.Chat.ID, "Файл получен! Скоро сможем конвертировать ")
			bot.Send(msg)
			continue
		}

		// Просто повторю любое сообщение
		msg := telegram.NewMessage(update.Message.Chat.ID, "Ты написал: "+update.Message.Text)
		bot.Send(msg)
	}
}