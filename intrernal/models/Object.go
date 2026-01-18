package models

import "io"

// Для Хранилища сырых файлов
type Object struct {
	Reader 		io.Reader
	Size 		int64
	ContentType string
}