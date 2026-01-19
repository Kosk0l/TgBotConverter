package domains

// Job - Redis: List, Hash
type Job struct {
	JobID 		string
	UserID 		int64
	ChatID		int64
	FileTypeIn	string
	FileTypeTo 	string
	StatusJob   string
}
