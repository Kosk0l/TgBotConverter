package handlers

import (
	"context"
	"io"

	userService "github.com/Kosk0l/TgBotConverter/internal/Services/userService"
	"github.com/Kosk0l/TgBotConverter/internal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// bot - http cliet
// update - http request // Содержит всю информацию

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

//====================================================================================================

type UserServiceRepository interface {
	GetByIdService(ctx context.Context, userId int64) (domains.User, error)
	CreateUserService(ctx context.Context, user domains.User) (error)
	UpdateUserService(ctx context.Context, user domains.User) (error)
	UpdateLastSeenService(ctx context.Context, userId int64) (error)
	DeleteUserService(ctx context.Context, userId int64) (error)
}

type JobServiceRepository interface {
	CreateJob(ctx context.Context, job domains.Job, jobObj domains.Object) (string, error)
	GetJob(ctx context.Context) (domains.Job, io.Reader, error) 
}

type DialogServiceRepository interface {
	SetState(ctx context.Context, state domains.State) (error)
	GetState(ctx context.Context, chatId int64) (domains.State, error)
}

//====================================================================================================

// TODO: Дальше можно разрезать по зонам ответственности: ht *HandlerText
// TODO: добавить инъкцию зависимостей

// основной хендлер сообщений
type Handler struct {
	bot *telegram.BotAPI
	us 	*userService.UserService
	js 	JobServiceRepository
	ds 	DialogServiceRepository
}

// Конструктор
func NewServer(bot *telegram.BotAPI, us *userService.UserService, js JobServiceRepository, ds DialogServiceRepository) (*Handler) {
	return &Handler{
		bot: bot,
		us: us,
		js: js,
		ds: ds,
	}
}

//====================================================================================================

// Распределяет по типам сообщения
func (h *Handler) HandleUpdate(ctx context.Context, update telegram.Update) {
	if update.CallbackQuery != nil {
		h.HandleCallBack(ctx, update)
		return
	}
	
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		h.HandleCommand(ctx, update)
		return
	}	

	if update.Message.Document != nil {
		h.HandleDocument(ctx, update)
		return
	}

	h.HandleText(ctx, update)
}

