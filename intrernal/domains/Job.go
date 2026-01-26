package domains

// Job - Redis: List, Hash
type Job struct {
	JobID 		string
	ChatID		int64
	FileTypeTo 	FileType
}

// Константы для типов файлов
type FileType string
const (
	Pdf FileType = "pdf"
	Docx FileType = "docx"
	Jpeg FileType = "jpeg"
	Xlsx FileType = "xlsx"
)