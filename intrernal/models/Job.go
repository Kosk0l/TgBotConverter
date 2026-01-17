package models

// domain Job - Redis, List, Hash
type Job struct {
	JobID 		string
	UserID 		int64
	ChatID		int64
	FileTypeIn	string
	FileTypeTo 	string
	StatusJob   string
}

// Константы для статуса - string
const (
	InQueue = "inQueue" 
	ProcessedNow = "processedNow" 
	GoodConvert = "goodConvert"
	RedisError = "redisError"
)

// Константы для файлов - string
const (
	PDF = "PDF"
	JPG = "JPG"
	DOCX = "DOCX"
)