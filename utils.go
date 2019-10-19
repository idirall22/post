package post

// check if id is valid
func checkIfIDIsValid(id int64) error {
	if id < 1 {
		return ErrorID
	}
	return nil
}
