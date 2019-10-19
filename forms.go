package post

// PForm structure
type PForm struct {
	Content string `json:"content"`
	UserID  int64  `json:"user_id"`
	GroupID int64  `json:"group_id"`
}
