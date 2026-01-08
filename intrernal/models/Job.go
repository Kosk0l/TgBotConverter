package models

// domain Job - Redis, List, Hash
type Job struct {
	JobID 		int64
	UserID 		int64
	ChatID		int64
	UserName 	string
	status		string
}