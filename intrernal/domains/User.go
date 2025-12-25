package domains

import "time"

// domain services Пользователя
type User struct {
	ID        int64
	UseNname  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	LastSeen  time.Time
}