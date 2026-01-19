package domains

import "time"

// domain services Пользователя
type User struct {
	ID        int64
	UserName  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	LastSeen  time.Time
}