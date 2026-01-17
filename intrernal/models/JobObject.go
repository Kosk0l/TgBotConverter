package models

import "io"

// Для Хранилища сырых файлов
type JobObject struct {
	reader 		io.Reader
	size 		int64
	ContentType string
}