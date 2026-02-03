package domains

import "time"

// Структура Пользователя в БД
type User struct {
	ID        int64
	UserName  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	LastSeen  time.Time
}