package domains

// Структура сообщения пользователя
type State struct {
	FileURL     string
    FileName    string
	ChatId 		int64
    Size        int64
	ContentType string
}