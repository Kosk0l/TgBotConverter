package domains

// Структура сообщения пользователя
type State struct {
	ChatId 		int64
	Step 		Step

	FileURL     string
    FileName    string
    Size        int64
	ContentType string
}

type Step string

const (
	WaitingTargetType Step = "waiting_target_type"
	//..
)