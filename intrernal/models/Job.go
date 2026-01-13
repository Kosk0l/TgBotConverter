package models

// domain Job - Redis, List, Hash
type Job struct {
	JobID 		int64
	UserID 		int64
	ChatID		int64
	FileTypeIn	string
	FileTypeTo 	string
	StatusJob   string
}

// Константы для статуса
const (
	InQueue = "inQueue" 
	ProcessedNow = "processedNow" 
	GoodConvert = "goodConvert"
	RedisError = "redisError"
)

// Константы для файлов
const (
	PDF = "PDF"
	JPG = "JPG"
	DOCX = "DOCX"
)