package post

// PForm structure
type PForm struct {
	Content string `json:"content"`
	UserID  int64  `json:"user_id"`
	GroupID int64  `json:"group_id"`
}

// ValidateForm validate form
func (f *PForm) ValidateForm(userID int64) bool {

	if f.UserID != userID {
		return false
	}

	if f.Content == "" || f.UserID < 1 || f.GroupID < 1 {
		return false
	}
	return true
}
