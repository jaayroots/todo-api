package enums

type TodoStatus int

const (
	StatusNew TodoStatus = iota + 1
	StatusPending
	StatusInProgress
	StatusCompleted
)

func (s TodoStatus) String() string {
	switch s {
	case StatusNew:
		return "new"
	case StatusPending:
		return "pending"
	case StatusInProgress:
		return "in_progress"
	case StatusCompleted:
		return "completed"
	default:
		return "unknow"
	}
}

func GetTodoStatusMap() map[int]string {
	return map[int]string{
		int(StatusNew):        StatusNew.String(),
		int(StatusPending):    StatusPending.String(),
		int(StatusInProgress): StatusInProgress.String(),
		int(StatusCompleted):  StatusCompleted.String(),
	}
}

func IsValidTodoStatus(value int) bool {
	switch TodoStatus(value) {
	case StatusNew, StatusPending, StatusInProgress, StatusCompleted:
		return true
	default:
		return false
	}
}
